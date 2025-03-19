// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package framework

type ActionType string

const (
	Reclaim           ActionType = "reclaim"
	Preempt           ActionType = "preempt"
	Allocate          ActionType = "allocate"
	Consolidation     ActionType = "consolidation"
	StaleGangEviction ActionType = "stalegangeviction"
)

// Action is the interface of scheduler action.
type Action interface {
	// The unique name of Action.
	Name() ActionType

	// Execute allocates the cluster's resources into each queue.
	Execute(ssn *Session)
}

type Plugin interface {
	// The unique name of Plugin.
	Name() string

	OnSessionOpen(ssn *Session)
	OnSessionClose(ssn *Session)
}
