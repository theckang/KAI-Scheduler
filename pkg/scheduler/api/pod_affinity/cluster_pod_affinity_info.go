// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package pod_affinity

import (
	v1 "k8s.io/api/core/v1"
	k8sframework "k8s.io/kubernetes/pkg/scheduler/framework"
)

type ClusterPodAffinityInfo interface {
	UpdateNodeAffinity(podAffinityInfo NodePodAffinityInfo)
	AddNode(string, *k8sframework.NodeInfo)
}

type NodePodAffinityInfo interface {
	AddPod(*v1.Pod)
	RemovePod(*v1.Pod) error
	HasPodsWithPodAffinity() bool
	HasPodsWithPodAntiAffinity() bool
	Name() string
}
