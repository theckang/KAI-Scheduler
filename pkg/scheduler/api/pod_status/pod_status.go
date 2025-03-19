// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package pod_status

import "golang.org/x/exp/slices"

// PodStatus defines the status of a task/pod.
type PodStatus int

const (
	// Pending means the task is pending in the apiserver.
	Pending PodStatus = 1 << iota

	// Allocated means the scheduler assigns a host to it.
	Allocated

	// Pipelined means the scheduler assigns a host to wait for releasing resource.
	Pipelined

	// Binding means the scheduler send Bind request to apiserver.
	Binding

	// Bound means the task/Pod bounds to a host.
	Bound

	// Running means a task is running on the host.
	Running

	// Releasing means a task/pod is deleted.
	Releasing

	// Succeeded means that all containers in the pod have voluntarily terminated
	// with a container exit code of 0, and the system is not going to restart any of these containers.
	Succeeded

	// Failed means that all containers in the pod have terminated, and at least one container has
	// terminated in a failure (exited with a non-zero exit code or was stopped by the system).
	Failed

	// Unknown means the status of task/pod is unknown to the scheduler.
	Unknown

	Deleted
)

func (ps PodStatus) String() string {
	switch ps {
	case Pending:
		return "Pending"
	case Allocated:
		return "Allocated"
	case Pipelined:
		return "Pipelined"
	case Binding:
		return "Binding"
	case Bound:
		return "Bound"
	case Running:
		return "Running"
	case Releasing:
		return "Releasing"
	case Succeeded:
		return "Succeeded"
	case Failed:
		return "Failed"
	case Deleted:
		return "Deleted"
	default:
		return "Unknown"
	}
}

var activeUsedStatuses = []PodStatus{Allocated, Pipelined, Binding, Bound, Running, Releasing}
var activeAllocatedStatuses = []PodStatus{Allocated, Pipelined, Binding, Bound, Running}
var aliveStatuses = []PodStatus{Allocated, Pipelined, Binding, Bound, Running, Pending}
var boundStatuses = []PodStatus{Allocated, Bound, Running, Releasing}
var allocatedStatuses = []PodStatus{Allocated, Bound, Binding, Running}

func IsAliveStatus(statusInput PodStatus) bool {
	return slices.Contains(aliveStatuses, statusInput)
}

func IsActiveUsedStatus(statusInput PodStatus) bool {
	return slices.Contains(activeUsedStatuses, statusInput)
}

func IsActiveAllocatedStatus(statusInput PodStatus) bool {
	return slices.Contains(activeAllocatedStatuses, statusInput)
}

func IsPodBound(statusInput PodStatus) bool {
	return slices.Contains(boundStatuses, statusInput)
}

func AllocatedStatus(status PodStatus) bool {
	return slices.Contains(allocatedStatuses, status)
}
