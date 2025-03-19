// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package spotrequest

import (
	"github.com/NVIDIA/KAI-scheduler/pkg/podgrouper/podgroup"
	"github.com/NVIDIA/KAI-scheduler/pkg/podgrouper/podgrouper/plugins"
	"github.com/NVIDIA/KAI-scheduler/pkg/podgrouper/podgrouper/plugins/constants"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// +kubebuilder:rbac:groups=egx.nvidia.io,resources=spotrequests,verbs=get;list;watch
// +kubebuilder:rbac:groups=egx.nvidia.io,resources=spotrequests/finalizers,verbs=patch;update;create

func GetPodGroupMetadata(topOwner *unstructured.Unstructured, pod *v1.Pod, _ ...*metav1.PartialObjectMetadata) (*podgroup.Metadata, error) {
	pgMetadata, err := plugins.GetPodGroupMetadata(topOwner, pod)
	if err != nil {
		return nil, err
	}

	pgMetadata.PriorityClassName = plugins.CalcPodGroupPriorityClass(topOwner, pod, constants.InferencePriorityClass)

	return pgMetadata, nil

}
