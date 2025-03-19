// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package ray

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/NVIDIA/KAI-scheduler/pkg/podgrouper/podgroup"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func (rg *RayGrouper) GetPodGroupMetadataForRayService(
	topOwner *unstructured.Unstructured, pod *v1.Pod, _ ...*metav1.PartialObjectMetadata,
) (*podgroup.Metadata, error) {
	return rg.getPodGroupMetadataWithClusterNamePath(topOwner, pod, [][]string{
		{"status", "activeServiceStatus", "rayClusterName"}, {"status", "pendingServiceStatus", "rayClusterName"},
	})
}
