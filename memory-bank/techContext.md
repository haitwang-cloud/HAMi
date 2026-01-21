# Technical Context: HAMi

## Technology Stack

### Core Language & Framework
- **Go 1.25.5**: Primary programming language
- **Kubernetes APIs**: Native integration with k8s.io/api, k8s.io/client-go, k8s.io/apimachinery
- **Controller Runtime**: Uses sigs.k8s.io/controller-runtime for Kubernetes controller patterns

### Key Dependencies

#### Kubernetes Integration
- `k8s.io/kube-scheduler v0.28.3`: Scheduler extender implementation
- `k8s.io/kubelet v0.32.3`: Integration with kubelet device plugin framework
- `k8s.io/api v0.33.0`: Kubernetes API types and client

#### Device Management
- `github.com/NVIDIA/go-gpuallocator v0.6.0`: GPU allocation algorithms
- `github.com/NVIDIA/go-nvlib v0.9.0`: NVIDIA library bindings
- `github.com/NVIDIA/go-nvml v0.13.0-1`: NVIDIA management library
- `github.com/NVIDIA/k8s-device-plugin v0.18.1`: Device plugin framework
- `github.com/NVIDIA/nvidia-container-toolkit v1.18.1`: Container toolkit integration

#### Observability & Monitoring
- `github.com/prometheus/client_golang v1.23.2`: Metrics collection
- `github.com/sirupsen/logrus v1.9.4`: Structured logging
- `k8s.io/klog/v2 v2.130.1`: Kubernetes logging integration

#### Utilities & Tools
- `github.com/spf13/cobra v1.10.2`: CLI framework
- `github.com/stretchr/testify v1.11.1`: Testing framework
- `github.com/onsi/ginkgo/v2 v2.27.5`: BDD testing
- `gopkg.in/yaml.v2 v2.4.0`: YAML parsing
- `github.com/google/uuid v1.6.0`: UUID generation

## Development Environment

### Prerequisites
- **Go 1.25.5+**: Language runtime
- **Kubernetes 1.18+**: Target cluster version
- **Helm 3.0+**: Package management
- **Docker**: Container build environment

### Hardware Requirements
- **glibc >= 2.17**: System library compatibility
- **glibc < 2.30**: Version constraints
- **kernel >= 3.10**: Minimum kernel version
- **NVIDIA drivers >= 440**: For GPU support
- **nvidia-docker >= 2.0**: Container runtime integration

### Build System
- **Makefile**: Primary build orchestration
- **Docker**: Multi-stage builds for different components
- **Go modules**: Dependency management

## Architecture Components

### Control Plane
- **Mutating Webhook**: Validates and modifies pod specifications
- **Scheduler Extender**: Implements Filter and Score methods for device-aware scheduling
- **Custom Device Plugins**: Vendor-specific device management

### Data Plane
- **In-Container Control**: Device-specific runtime enforcement
  - HAMi-Core: NVIDIA GPU control
  - libvgpu-control.so: Iluvatar GPU control
  - Vendor-specific libraries for other devices

### Monitoring & Observability
- **Metrics Endpoint**: Prometheus-compatible metrics on port 31993
- **Grafana Dashboards**: Pre-built visualization templates
- **Event Logging**: Kubernetes event integration

## Supported Device Types

### Officially Supported
- **NVIDIA GPUs**: Full feature set including MIG, vGPU, topology awareness
- **Huawei Ascend NPUs**: Ascend 910, 910B series support
- **Cambricon MLUs**: MLU device virtualization
- **HYGON DCUs**: DCU device management
- **Iluvatar GPUs**: CoreX series support
- **Moore Threads GPUs**: GPU virtualization
- **MetaX GPUs**: Custom GPU support

### Integration Patterns
- **Device Plugins**: Kubernetes device plugin interface implementation
- **Scheduler Extenders**: Extend default/volcano schedulers
- **Webhook Admission**: Mutating admission controller
- **Container Runtime**: Integration with containerd/docker via device libraries

## Development Workflow

### Local Development
```bash
# Build all components
make build

# Run tests
make test

# Build Docker images
make docker-build

# Deploy to kind/minikube
make deploy-local
```

### Testing Strategy
- **Unit Tests**: Go testing with testify/ginkgo
- **Integration Tests**: Kubernetes API integration
- **E2E Tests**: Full cluster deployment testing
- **Performance Benchmarks**: Device allocation performance

### CI/CD Pipeline
- **GitHub Actions**: Automated testing and releases
- **Code Coverage**: Codecov integration
- **Security Scanning**: Trivy vulnerability scanning
- **License Checking**: Import alias verification

## Deployment Patterns

### Helm Chart Structure
```
charts/hami/
├── Chart.yaml          # Chart metadata
├── values.yaml         # Default configuration
└── templates/          # Kubernetes manifests
    ├── deployment.yaml
    ├── service.yaml
    ├── configmap.yaml
    └── webhook.yaml
```

### Configuration Management
- **Helm Values**: Runtime configuration via values.yaml
- **ConfigMaps**: Device-specific configuration
- **Command Line Flags**: Runtime behavior modification
- **Environment Variables**: Container-specific settings

## Constraints & Limitations

### Technical Constraints
- **Kubernetes Version**: Minimum 1.18, tested up to 1.33
- **Go Version**: Requires Go 1.25.5+ for module support
- **glibc Compatibility**: System library version restrictions
- **Kernel Requirements**: Minimum 3.10 for device access

### Device-Specific Limitations
- **NVIDIA MIG**: Only "none" and "mixed" modes supported
- **Device Affinity**: nodeName field incompatible with scheduling
- **Resource Accounting**: vGPU count vs physical GPU distinction

### Operational Considerations
- **Node Labeling**: Requires `gpu=on` label for management
- **Network Policies**: May require webhook communication
- **RBAC**: Requires cluster-admin for installation
- **Monitoring Setup**: Metrics collection requires task submission

## Tool Usage Patterns

### Development Tools
- **GoLand/VS Code**: Primary IDEs with Go extensions
- **kubectl**: Cluster interaction and debugging
- **helm**: Package management and deployment
- **docker**: Local container development
- **kind/minikube**: Local cluster testing

### Debugging Patterns
- **Logs**: `kubectl logs` for component debugging
- **Events**: `kubectl get events` for scheduling insights
- **Metrics**: Prometheus queries for performance analysis
- **Device Plugins**: Direct API calls for device status

### Maintenance Scripts
- `hack/build.sh`: Cross-platform building
- `hack/unit-test.sh`: Test execution
- `hack/verify-*.sh`: Code quality checks
- `hack/deploy-helm.sh`: Deployment automation
