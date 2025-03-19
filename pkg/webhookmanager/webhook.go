// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package webhookmanager

import (
	"context"

	admissionv1 "k8s.io/api/admissionregistration/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations;mutatingwebhookconfigurations,verbs=get;update;patch

func PatchValidatingWebhookCA(ctx context.Context, kubeClient client.Client, name string, certBytes []byte) error {
	webhookConfig := &admissionv1.ValidatingWebhookConfiguration{}
	err := kubeClient.Get(ctx, types.NamespacedName{Name: name}, webhookConfig)
	if err != nil {
		return err
	}

	logger := log.FromContext(ctx)
	logger.Info("Updating validating webhook", "name", name)
	origWebhookConfig := webhookConfig.DeepCopy()
	for index := range len(webhookConfig.Webhooks) {
		webhookConfig.Webhooks[index].ClientConfig.CABundle = certBytes
	}

	return kubeClient.Patch(ctx, webhookConfig, client.MergeFrom(origWebhookConfig))
}

func PatchMutatingWebhookCA(ctx context.Context, kubeClient client.Client, name string, certBytes []byte) error {
	webhookConfig := &admissionv1.MutatingWebhookConfiguration{}
	err := kubeClient.Get(ctx, types.NamespacedName{Name: name}, webhookConfig)
	if err != nil {
		return err
	}

	logger := log.FromContext(ctx)
	logger.Info("Updating mutating webhook", "name", name)
	origWebhookConfig := webhookConfig.DeepCopy()
	for index := range len(webhookConfig.Webhooks) {
		webhookConfig.Webhooks[index].ClientConfig.CABundle = certBytes
	}

	return kubeClient.Patch(ctx, webhookConfig, client.MergeFrom(origWebhookConfig))
}
