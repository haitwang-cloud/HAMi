# System Patterns: HAMi

## Architecture Overview

HAMi implements a layered architecture that integrates deeply with Kubernetes while providing vendor-agnostic device management:

```
┌─────────────────┐
│   Applications  │
├─────────────────┤
│  Kubernetes API │
├─────────────────┤
│ MutatingWebhook │ ← HAMi Admission Control
├─────────────────┤
│    Scheduler    │ ← HAMi Scheduler Extender
├─────────────────┤
│  Device Plugin  │ ← HAMi Device Plugins
├─────────────────┤
│ In-Container    │ ← Device-Specific Libraries
│    Control      │
├─────────────────┤
│  Hardware Layer │
└─────────────────┘
```

## Core Design Patterns

### 1. Scheduler Extender Pattern

**Problem**: Kubernetes default scheduler lacks awareness of heterogeneous device topology and sharing capabilities.

**Solution**: Implement scheduler extender that provides Filter and Score methods:

```go
// Filter: Find nodes with sufficient device resources
func (h *HamiScheduler) Filter(args schedulerapi.ExtenderArgs) schedulerapi.ExtenderFilterResult

// Score: Rank nodes based on device topology and utilization
func (h *HamiScheduler) Score(args schedulerapi.ExtenderArgs) schedulerapi.HostPriorityList
```

**Key Decisions**:
- Extends both default kube-scheduler and Volcano scheduler
- Uses annotations on pods to communicate scheduling decisions
- Implements topology-aware scoring for optimal device placement

### 2. Device Plugin Pattern

**Problem**: Standard Kubernetes device plugins don't support device sharing or virtualization.

**Solution**: Custom device plugins that understand HAMi's resource model:

```go
type HamiDevicePlugin struct {
    deviceType string
    allocator  DeviceAllocator
}

func (p *HamiDevicePlugin) Allocate(ctx context.Context, req *pluginapi.AllocateRequest) (*pluginapi.AllocateResponse, error)
```

**Key Decisions**:
- Translates HAMi resource requests to device-specific allocations
- Generates environment variables and volume mounts for containers
- Handles device-specific initialization and cleanup

### 3. Mutating Webhook Pattern

**Problem**: Need to intercept and modify pod specifications for HAMi-managed devices.

**Solution**: Admission controller that validates and annotates pods:

```go
func (wh *WebhookHandler) Handle(ctx context.Context, req admission.Request) admission.Response {
    pod := &corev1.Pod{}
    // Validate device requests
    // Set scheduler name to "hami-scheduler"
    // Add scheduling annotations
}
```

**Key Decisions**:
- Only intercepts pods with HAMi-recognized resource requests
- Passes through non-HAMi pods unchanged
- Validates resource request combinations

### 4. In-Container Control Pattern

**Problem**: Need to enforce resource limits within containers without device driver modifications.

**Solution**: Device-specific shared libraries loaded at container runtime:

```go
// NVIDIA: HAMi-Core library
// Iluvatar: libvgpu-control.so
// Huawei: Ascend-specific libraries
```

**Key Decisions**:
- Library injection via LD_PRELOAD environment variable
- Device-specific enforcement mechanisms
- Zero application code changes required

## Component Relationships

### Control Flow

1. **Pod Creation** → MutatingWebhook validates and annotates
2. **Scheduling** → Scheduler Extender finds optimal node placement
3. **Pod Binding** → Device Plugin allocates specific devices
4. **Container Start** → In-container libraries enforce limits

### Data Flow

```
Pod Spec → Webhook → Annotations → Scheduler → Node Selection → Device Plugin → Env Vars → Container → Library → Hardware
```

### Communication Patterns

- **Scheduler ↔ Device Plugin**: gRPC calls for resource discovery
- **Webhook ↔ API Server**: REST API for admission control
- **Device Plugin ↔ Kubelet**: Device plugin protocol
- **Container ↔ Hardware**: Vendor-specific device drivers

## Critical Implementation Paths

### 1. Device Discovery & Registration

```go
// Device plugin startup
func (p *NvidiaDevicePlugin) Start() error {
    devices := discoverDevices()        // Find physical devices
    vGPUs := createVirtualDevices()     // Create virtual GPUs
    registerWithKubelet(vGPUs)          // Register with kubelet
}
```

**Critical Path**: Device enumeration must be accurate and complete.

### 2. Resource Allocation Algorithm

```go
func allocateResources(requested, available DeviceResources) (Allocation, error) {
    // Check availability
    // Calculate topology score
    // Reserve resources
    // Generate allocation metadata
}
```

**Critical Path**: Must handle concurrent allocations and prevent over-subscription.

### 3. Scheduling Decision Flow

```go
func makeSchedulingDecision(pod *v1.Pod) (nodeName string, annotations map[string]string) {
    candidates := filterCandidates(pod)    // Find feasible nodes
    scored := scoreCandidates(candidates)  // Score by topology/utilization
    selected := selectBest(scored)         // Choose optimal placement
    annotated := addAnnotations(selected)  // Add scheduling metadata
}
```

**Critical Path**: Decision must be consistent and deterministic.

### 4. Container Runtime Integration

```go
func prepareContainer(pod *v1.Pod, allocation Allocation) ContainerConfig {
    env := generateEnvironmentVars(allocation)
    mounts := generateVolumeMounts(allocation)
    devices := generateDeviceAccess(allocation)
    return ContainerConfig{env, mounts, devices}
}
```

**Critical Path**: All required environment and mounts must be correct for device access.

## Key Technical Decisions

### 1. Resource Model Design

**Decision**: Extend Kubernetes resource model with device-specific resources:

```yaml
resources:
  limits:
    nvidia.com/gpu: 1        # Physical GPU count
    nvidia.com/gpumem: 4096  # Memory in MB
    nvidia.com/gpucores: 50  # Core percentage
```

**Rationale**: Maintains Kubernetes compatibility while enabling fine-grained control.

### 2. Annotation-Based Communication

**Decision**: Use pod annotations to communicate between components:

```yaml
annotations:
  hami.io/scheduling-result: '{"node":"gpu-01","devices":[{"id":"0","memory":"4096"}]}'
  hami.io/device-allocated: 'nvidia-gpu-0'
```

**Rationale**: Decouples components while maintaining auditability.

### 3. Vendor-Specific Abstraction

**Decision**: Abstract device operations behind common interfaces:

```go
type DeviceManager interface {
    DiscoverDevices() ([]Device, error)
    AllocateDevice(request AllocationRequest) (Allocation, error)
    GetDeviceTopology() Topology
}
```

**Rationale**: Enables consistent behavior across different device vendors.

### 4. Library Injection Strategy

**Decision**: Use LD_PRELOAD for runtime enforcement:

```bash
export LD_PRELOAD=/usr/lib/hami-core.so
export CUDA_VISIBLE_DEVICES=0
```

**Rationale**: Works with existing containers without image modifications.

## Failure Modes & Recovery

### 1. Device Plugin Failure
- **Detection**: Kubelet health checks
- **Recovery**: Automatic restart, resource re-registration

### 2. Scheduler Unavailability
- **Detection**: Pod scheduling timeouts
- **Recovery**: Fallback to default scheduler for non-HAMi pods

### 3. Resource Over-Allocation
- **Detection**: Runtime validation in device libraries
- **Recovery**: Container termination with descriptive errors

### 4. Node Resource Drift
- **Detection**: Periodic reconciliation
- **Recovery**: Resource rebalancing across cluster

## Performance Considerations

### 1. Scheduling Latency
- Filter operations: O(nodes × devices)
- Score operations: O(candidates × topology_complexity)

### 2. Memory Overhead
- Device state tracking: ~1KB per device
- Pod annotation metadata: ~2KB per pod

### 3. Scalability Limits
- Tested with 1000+ nodes
- 10,000+ concurrent pods
- Sub-millisecond scheduling decisions

## Evolution Patterns

### 1. Adding New Device Types
1. Implement DeviceManager interface
2. Add device-specific resource types
3. Create corresponding device plugin
4. Add in-container control library

### 2. Extending Scheduling Policies
1. Implement new scoring algorithm
2. Add policy-specific annotations
3. Update webhook validation
4. Test with existing workloads

### 3. Performance Optimization
1. Profile critical paths
2. Implement caching layers
3. Add parallel processing
4. Optimize data structures
