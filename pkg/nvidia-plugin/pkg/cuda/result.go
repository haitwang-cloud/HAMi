/*
* SPDX-License-Identifier: Apache-2.0
*
* The HAMi Contributors require contributions made to
* this file be licensed under the Apache-2.0 license or a
* compatible open source license.
 */

/*
* Licensed to NVIDIA CORPORATION under one or more contributor
* license agreements. See the NOTICE file distributed with
* this work for additional information regarding copyright
* ownership. NVIDIA CORPORATION licenses this file to you under
* the Apache License, Version 2.0 (the "License"); you may
* not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing,
* software distributed under the License is distributed on an
* "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
* KIND, either express or implied.  See the License for the
* specific language governing permissions and limitations
* under the License.
 */

/*
* Modifications Copyright The HAMi Authors. See
* GitHub history for details.
 */

package cuda

import (
	"fmt"
)

// String returns the string representation of a Result
func (r Result) String() string {
	return errorStringFunc(r)
}

// Error returns the string representation of a Result
func (r Result) Error() string {
	return r.String()
}

var errorStringFunc = defaultErrorStringFunc

var defaultErrorStringFunc = func(r Result) string {
	switch r {
	case SUCCESS:
		return "CUDA_SUCCESS"
	case ERROR_INVALID_VALUE:
		return "CUDA_ERROR_INVALID_VALUE"
	case ERROR_OUT_OF_MEMORY:
		return "CUDA_ERROR_OUT_OF_MEMORY"
	case ERROR_NOT_INITIALIZED:
		return "CUDA_ERROR_NOT_INITIALIZED"
	case ERROR_DEINITIALIZED:
		return "CUDA_ERROR_DEINITIALIZED"
	case ERROR_PROFILER_DISABLED:
		return "CUDA_ERROR_PROFILER_DISABLED"
	case ERROR_PROFILER_NOT_INITIALIZED:
		return "CUDA_ERROR_PROFILER_NOT_INITIALIZED"
	case ERROR_PROFILER_ALREADY_STARTED:
		return "CUDA_ERROR_PROFILER_ALREADY_STARTED"
	case ERROR_PROFILER_ALREADY_STOPPED:
		return "CUDA_ERROR_PROFILER_ALREADY_STOPPED"
	case ERROR_NO_DEVICE:
		return "CUDA_ERROR_NO_DEVICE"
	case ERROR_INVALID_DEVICE:
		return "CUDA_ERROR_INVALID_DEVICE"
	case ERROR_INVALID_IMAGE:
		return "CUDA_ERROR_INVALID_IMAGE"
	case ERROR_INVALID_CONTEXT:
		return "CUDA_ERROR_INVALID_CONTEXT"
	case ERROR_CONTEXT_ALREADY_CURRENT:
		return "CUDA_ERROR_CONTEXT_ALREADY_CURRENT"
	case ERROR_MAP_FAILED:
		return "CUDA_ERROR_MAP_FAILED"
	case ERROR_UNMAP_FAILED:
		return "CUDA_ERROR_UNMAP_FAILED"
	case ERROR_ARRAY_IS_MAPPED:
		return "CUDA_ERROR_ARRAY_IS_MAPPED"
	case ERROR_ALREADY_MAPPED:
		return "CUDA_ERROR_ALREADY_MAPPED"
	case ERROR_NO_BINARY_FOR_GPU:
		return "CUDA_ERROR_NO_BINARY_FOR_GPU"
	case ERROR_ALREADY_ACQUIRED:
		return "CUDA_ERROR_ALREADY_ACQUIRED"
	case ERROR_NOT_MAPPED:
		return "CUDA_ERROR_NOT_MAPPED"
	case ERROR_NOT_MAPPED_AS_ARRAY:
		return "CUDA_ERROR_NOT_MAPPED_AS_ARRAY"
	case ERROR_NOT_MAPPED_AS_POINTER:
		return "CUDA_ERROR_NOT_MAPPED_AS_POINTER"
	case ERROR_ECC_UNCORRECTABLE:
		return "CUDA_ERROR_ECC_UNCORRECTABLE"
	case ERROR_UNSUPPORTED_LIMIT:
		return "CUDA_ERROR_UNSUPPORTED_LIMIT"
	case ERROR_CONTEXT_ALREADY_IN_USE:
		return "CUDA_ERROR_CONTEXT_ALREADY_IN_USE"
	case ERROR_PEER_ACCESS_UNSUPPORTED:
		return "CUDA_ERROR_PEER_ACCESS_UNSUPPORTED"
	case ERROR_INVALID_PTX:
		return "CUDA_ERROR_INVALID_PTX"
	case ERROR_INVALID_GRAPHICS_CONTEXT:
		return "CUDA_ERROR_INVALID_GRAPHICS_CONTEXT"
	case ERROR_NVLINK_UNCORRECTABLE:
		return "CUDA_ERROR_NVLINK_UNCORRECTABLE"
	case ERROR_JIT_COMPILER_NOT_FOUND:
		return "CUDA_ERROR_JIT_COMPILER_NOT_FOUND"
	case ERROR_INVALID_SOURCE:
		return "CUDA_ERROR_INVALID_SOURCE"
	case ERROR_FILE_NOT_FOUND:
		return "CUDA_ERROR_FILE_NOT_FOUND"
	case ERROR_SHARED_OBJECT_SYMBOL_NOT_FOUND:
		return "CUDA_ERROR_SHARED_OBJECT_SYMBOL_NOT_FOUND"
	case ERROR_SHARED_OBJECT_INIT_FAILED:
		return "CUDA_ERROR_SHARED_OBJECT_INIT_FAILED"
	case ERROR_OPERATING_SYSTEM:
		return "CUDA_ERROR_OPERATING_SYSTEM"
	case ERROR_INVALID_HANDLE:
		return "CUDA_ERROR_INVALID_HANDLE"
	case ERROR_NOT_FOUND:
		return "CUDA_ERROR_NOT_FOUND"
	case ERROR_NOT_READY:
		return "CUDA_ERROR_NOT_READY"
	case ERROR_ILLEGAL_ADDRESS:
		return "CUDA_ERROR_ILLEGAL_ADDRESS"
	case ERROR_LAUNCH_OUT_OF_RESOURCES:
		return "CUDA_ERROR_LAUNCH_OUT_OF_RESOURCES"
	case ERROR_LAUNCH_TIMEOUT:
		return "CUDA_ERROR_LAUNCH_TIMEOUT"
	case ERROR_LAUNCH_INCOMPATIBLE_TEXTURING:
		return "CUDA_ERROR_LAUNCH_INCOMPATIBLE_TEXTURING"
	case ERROR_PEER_ACCESS_ALREADY_ENABLED:
		return "CUDA_ERROR_PEER_ACCESS_ALREADY_ENABLED"
	case ERROR_PEER_ACCESS_NOT_ENABLED:
		return "CUDA_ERROR_PEER_ACCESS_NOT_ENABLED"
	case ERROR_PRIMARY_CONTEXT_ACTIVE:
		return "CUDA_ERROR_PRIMARY_CONTEXT_ACTIVE"
	case ERROR_CONTEXT_IS_DESTROYED:
		return "CUDA_ERROR_CONTEXT_IS_DESTROYED"
	case ERROR_ASSERT:
		return "CUDA_ERROR_ASSERT"
	case ERROR_TOO_MANY_PEERS:
		return "CUDA_ERROR_TOO_MANY_PEERS"
	case ERROR_HOST_MEMORY_ALREADY_REGISTERED:
		return "CUDA_ERROR_HOST_MEMORY_ALREADY_REGISTERED"
	case ERROR_HOST_MEMORY_NOT_REGISTERED:
		return "CUDA_ERROR_HOST_MEMORY_NOT_REGISTERED"
	case ERROR_HARDWARE_STACK_ERROR:
		return "CUDA_ERROR_HARDWARE_STACK_ERROR"
	case ERROR_ILLEGAL_INSTRUCTION:
		return "CUDA_ERROR_ILLEGAL_INSTRUCTION"
	case ERROR_MISALIGNED_ADDRESS:
		return "CUDA_ERROR_MISALIGNED_ADDRESS"
	case ERROR_INVALID_ADDRESS_SPACE:
		return "CUDA_ERROR_INVALID_ADDRESS_SPACE"
	case ERROR_INVALID_PC:
		return "CUDA_ERROR_INVALID_PC"
	case ERROR_LAUNCH_FAILED:
		return "CUDA_ERROR_LAUNCH_FAILED"
	case ERROR_COOPERATIVE_LAUNCH_TOO_LARGE:
		return "CUDA_ERROR_COOPERATIVE_LAUNCH_TOO_LARGE"
	case ERROR_NOT_PERMITTED:
		return "CUDA_ERROR_NOT_PERMITTED"
	case ERROR_NOT_SUPPORTED:
		return "CUDA_ERROR_NOT_SUPPORTED"
	case ERROR_UNKNOWN:
		return "CUDA_ERROR_UNKNOWN"
	default:
		return fmt.Sprintf("Unknown return value: %d", r)
	}
}
