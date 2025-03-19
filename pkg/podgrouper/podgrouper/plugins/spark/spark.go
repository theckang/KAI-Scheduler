// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package spark

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/NVIDIA/KAI-scheduler/pkg/podgrouper/podgroup"
	"github.com/NVIDIA/KAI-scheduler/pkg/podgrouper/podgrouper/plugins"
	"github.com/NVIDIA/KAI-scheduler/pkg/podgrouper/podgrouper/plugins/constants"
)

const (
	sparkAppLabelName         = "spark-app-name"
	sparkAppSelectorLabelName = "spark-app-selector"

	sparkWorkloadKind = "Spark"
)

func IsSparkPod(pod *v1.Pod) bool {
	_, foundSparkApp := pod.Labels[sparkAppLabelName]
	_, foundSparkAppSelector := pod.Labels[sparkAppSelectorLabelName]
	return foundSparkApp && foundSparkAppSelector
}

func GetPodGroupMetadata(topOwner *unstructured.Unstructured, pod *v1.Pod) (*podgroup.Metadata, error) {
	podGroupMetadata, err := plugins.GetPodGroupMetadata(topOwner, pod)
	if err != nil {
		return nil, err
	}

	podGroupMetadata.Labels[constants.WorkloadKindLabelKey] = sparkWorkloadKind
	podGroupMetadata.Labels[constants.WorkloadNameLabelKey] = fmt.Sprintf(
		"%s-%s", pod.Labels[sparkAppLabelName], pod.Labels[sparkAppSelectorLabelName])
	podGroupMetadata.Name = pod.Labels[sparkAppSelectorLabelName]

	return podGroupMetadata, nil
}
