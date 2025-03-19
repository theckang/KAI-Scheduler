// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package cronjobs

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/NVIDIA/KAI-scheduler/pkg/podgrouper/podgroup"
	defaultpodgrouper "github.com/NVIDIA/KAI-scheduler/pkg/podgrouper/podgrouper/plugins"
)

// +kubebuilder:rbac:groups=batch,resources=cronjobs,verbs=get;list;watch
// +kubebuilder:rbac:groups=batch,resources=cronjobs/finalizers,verbs=patch;update;create

type cronJobGrouper struct {
	client client.Client
}

func NewCronJobGrouper(client client.Client) *cronJobGrouper {
	return &cronJobGrouper{client: client}
}

func (cg *cronJobGrouper) GetPodGroupMetadata(_ *unstructured.Unstructured, pod *v1.Pod, _ ...*metav1.PartialObjectMetadata) (*podgroup.Metadata, error) {
	owner, err := getJobOwnerReference(pod.OwnerReferences)
	if err != nil {
		return nil, err
	}

	job := &unstructured.Unstructured{}
	job.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "batch",
		Version: "v1",
		Kind:    "Job",
	})

	key := types.NamespacedName{
		Namespace: pod.Namespace,
		Name:      owner.Name,
	}
	err = cg.client.Get(context.Background(), key, job)
	if err != nil {
		return nil, err
	}
	return defaultpodgrouper.GetPodGroupMetadata(job, pod)
}

func getJobOwnerReference(references []metav1.OwnerReference) (*metav1.OwnerReference, error) {
	for _, owner := range references {
		if owner.Kind == "Job" {
			return &owner, nil
		}
	}
	return nil, fmt.Errorf("no owner reference of type 'Job'")
}
