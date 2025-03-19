// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package common

import (
	"context"

	v1 "k8s.io/api/core/v1"

	k8sframework "k8s.io/kubernetes/pkg/scheduler/framework"

	"github.com/NVIDIA/KAI-scheduler/pkg/apis/scheduling/v1alpha2"
)

type State *k8sframework.CycleState

func NewState() State {
	return k8sframework.NewCycleState()
}

type K8sPlugin interface {
	Name() string
	IsRelevant(pod *v1.Pod) bool
	PreFilter(ctx context.Context, pod *v1.Pod, state State) (error, bool)
	Filter(ctx context.Context, pod *v1.Pod, node *v1.Node, state State) error
	Allocate(ctx context.Context, pod *v1.Pod, hostname string, state State) error
	Bind(ctx context.Context, pod *v1.Pod, request *v1alpha2.BindRequest, state State) error
	PostBind(ctx context.Context, pod *v1.Pod, hostname string, state State)
	UnAllocate(ctx context.Context, pod *v1.Pod, hostname string, state State)
}
