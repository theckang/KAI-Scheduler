// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package constants

const (
	// Used for statefulsets that are preemptbile
	PriorityInferenceNumber = 125

	// (Build/Interactive)- used for statefulsets that aren't preemptible
	PriorityBuildNumber = 100

	// Used for statefulsets that are preemptbile
	PriorityInteractivePreemptibleNumber = 75

	// Used for batch jobs
	PriorityTrainNumber = 50
)
