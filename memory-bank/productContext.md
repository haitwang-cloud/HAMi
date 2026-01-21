# Product Context: HAMi

## Why HAMi Exists

### The Problem

Kubernetes clusters running AI/ML workloads face significant challenges with heterogeneous computing devices:

1. **Device Utilization Inefficiency**: Traditional device allocation requires entire GPUs/NPUs per pod, leading to underutilization
2. **Vendor Lock-in**: Each device vendor (NVIDIA, Huawei, Cambricon, etc.) requires vendor-specific management tools
3. **Complex Scheduling**: No native Kubernetes support for topology-aware scheduling of heterogeneous devices
4. **Resource Fragmentation**: Fixed device allocation prevents fine-grained sharing and optimal resource distribution
5. **Operational Complexity**: Managing different device types requires multiple management interfaces

### Market Context

- AI/ML workloads are exploding, requiring massive parallel computing resources
- Heterogeneous hardware (GPUs, NPUs, TPUs) from multiple vendors is becoming standard
- Kubernetes has become the de facto orchestration platform for AI workloads
- Cloud providers and enterprises need efficient resource utilization to control costs

## What HAMi Solves

### Core Value Proposition

HAMi provides **unified, efficient management of heterogeneous AI devices** in Kubernetes clusters, enabling:

- **70-80% improvement in device utilization** through fine-grained sharing
- **Zero application changes** when migrating between device types
- **Intelligent scheduling** based on device topology and workload requirements
- **Unified management interface** across all supported device vendors

### Key Benefits

1. **Cost Reduction**: Maximize ROI on expensive AI hardware through better utilization
2. **Operational Simplicity**: Single management interface for all heterogeneous devices
3. **Workload Portability**: Run AI applications on any supported device without code changes
4. **Performance Optimization**: Topology-aware scheduling for optimal workload placement

## How HAMi Works (User Experience)

### For AI/ML Practitioners

```yaml
# Simple pod specification - works on any supported device
apiVersion: v1
kind: Pod
metadata:
  name: ai-training-job
spec:
  containers:
  - name: training
    image: my-ai-app:latest
    resources:
      limits:
        nvidia.com/gpu: 1          # Physical GPU count
        nvidia.com/gpumem: 8192    # GPU memory in MB
        nvidia.com/gpucores: 50    # GPU cores percentage
```

**What users see**: Standard Kubernetes resource requests that "just work" across different device types.

### For Platform Administrators

```bash
# Label nodes for HAMi management
kubectl label nodes gpu-node-01 gpu=on

# Install HAMi via Helm
helm install hami hami-charts/hami -n kube-system

# Monitor cluster-wide device utilization
kubectl get nodes -l gpu=on
```

**What admins see**: Seamless integration with existing Kubernetes workflows and tooling.

## User Journey

1. **Discovery**: Learn about device sharing capabilities for cost optimization
2. **Installation**: Deploy HAMi alongside existing Kubernetes cluster
3. **Configuration**: Label nodes and configure device-specific settings
4. **Migration**: Gradually migrate workloads to use HAMi resource requests
5. **Optimization**: Use monitoring and scheduling policies to maximize utilization

## Success Stories

- **Internet Companies**: Improved GPU utilization from 30% to 80%
- **Cloud Providers**: Enabled multi-tenant AI workloads with resource isolation
- **Financial Services**: Cost-effective AI model training and inference
- **Research Institutions**: Shared expensive AI hardware across multiple research teams

## Competitive Landscape

- **Native Kubernetes**: No built-in heterogeneous device management
- **Vendor-specific tools**: Require application changes, single-vendor focus
- **Other schedulers**: Limited to specific device types or orchestration platforms
- **HAMi differentiators**: Multi-vendor support, zero application changes, Kubernetes-native

## Future Vision

HAMi aims to become the standard for heterogeneous device management in Kubernetes, enabling:

- Seamless scaling from single nodes to multi-cluster deployments
- Integration with emerging AI accelerators and specialized hardware
- Advanced scheduling policies based on workload characteristics and SLAs
- Unified observability and troubleshooting across all device types
