# Project Brief: HAMi

## Project Overview

HAMi (Heterogeneous AI Computing Virtualization Middleware), formerly known as 'k8s-vGPU-scheduler', is a Kubernetes middleware designed to manage heterogeneous AI computing devices. The project provides unified device virtualization, sharing, and scheduling capabilities across different types of accelerators including GPUs, NPUs, and other heterogeneous devices.

## Core Mission

To bridge the gap between different heterogeneous devices in Kubernetes clusters, providing a unified interface for device management without requiring application changes. HAMi enables efficient resource utilization through device sharing and intelligent scheduling decisions based on device topology and policies.

## Key Capabilities

- **Device Virtualization**: Supports partial allocation of device cores and memory
- **Device Sharing**: Enables multiple pods to share the same physical device with resource isolation
- **Intelligent Scheduling**: Makes topology-aware scheduling decisions using device plugins and scheduler extenders
- **Multi-Vendor Support**: Works with NVIDIA GPUs, Huawei Ascend NPUs, Cambricon MLUs, and other heterogeneous devices
- **Kubernetes Integration**: Seamlessly integrates with both default kube-scheduler and Volcano scheduler

## Target Users

- Cloud service providers running AI workloads
- Enterprises deploying AI/ML applications at scale
- Research institutions managing heterogeneous computing resources
- Organizations looking to optimize GPU/NPU utilization in Kubernetes clusters

## Success Metrics

- Efficient device utilization across heterogeneous hardware
- Zero application changes required for device migration
- Reliable scheduling decisions based on device topology and availability
- Broad adoption across different industries (finance, telecom, manufacturing, education)

## Project Status

HAMi is a Cloud Native Computing Foundation (CNCF) sandbox project with active community adoption. It has been widely deployed in production environments across internet companies, public/private clouds, and vertical industries, with over 50 organizations as active contributors and users.

## Technical Foundation

- Written in Go with Kubernetes-native architecture
- Uses mutating webhooks, scheduler extenders, and custom device plugins
- Implements device-specific in-container control mechanisms
- Provides monitoring and metrics collection capabilities
