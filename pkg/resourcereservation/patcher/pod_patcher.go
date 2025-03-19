// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package patcher

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	reservedGpuIndexAnnotation = "run.ai/reserve_for_gpu_index"
)

type PodPatcher struct {
	kubeClient client.Client
	namespace  string
	name       string
}

// +kubebuilder:rbac:groups="",resources=pods,verbs=get;list;watch;update;patch

func NewPodPatcher(kubeClient client.Client, namespace, name string) *PodPatcher {
	return &PodPatcher{
		kubeClient: kubeClient,
		namespace:  namespace,
		name:       name,
	}
}

func (pp *PodPatcher) AlreadyPatched(ctx context.Context) (error, bool) {
	pod := &corev1.Pod{}
	err := pp.kubeClient.Get(ctx, types.NamespacedName{Namespace: pp.namespace, Name: pp.name}, pod)
	if err != nil {
		return fmt.Errorf("could not find pod. %v", err), false
	}

	_, found := pod.Annotations[reservedGpuIndexAnnotation]
	return nil, found
}

func (pp *PodPatcher) PatchDeviceInfo(ctx context.Context, uuid string) error {
	originalPod := &corev1.Pod{}
	err := pp.kubeClient.Get(ctx, types.NamespacedName{Namespace: pp.namespace, Name: pp.name}, originalPod)
	if err != nil {
		return fmt.Errorf("could not get pod. %v", err)
	}

	updatedPod := originalPod.DeepCopy()
	if len(updatedPod.Annotations) == 0 {
		updatedPod.Annotations = map[string]string{}
	}
	updatedPod.Annotations[reservedGpuIndexAnnotation] = uuid

	err = pp.kubeClient.Patch(ctx, updatedPod, client.MergeFrom(originalPod))
	if err != nil {
		return fmt.Errorf("could not patch pod. %v", err)
	}

	return nil
}
