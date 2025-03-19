// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"flag"
)

type Options struct {
	Namespace             string
	ServiceName           string
	UpsertSecret          bool
	SecretName            string
	PatchWebhooks         bool
	ValidatingWebhookName string
	MutatingWebhookName   string
}

func (o *Options) AddFlags(fs *flag.FlagSet) {
	fs.StringVar(&o.Namespace, "namespace", "", "The service namespace")
	fs.StringVar(&o.ServiceName, "service-name", "", "The name of the service of the webhook")
	fs.BoolVar(&o.UpsertSecret, "upsert-secret", false, "Create or update a secret")
	fs.StringVar(&o.SecretName, "secret-name", "", "The name of the service of the webhook")
	fs.BoolVar(&o.PatchWebhooks, "patch-webhooks", false, "Patch webhook configuration")
	fs.StringVar(&o.ValidatingWebhookName, "validating-webhook-name", "", "The name of the validating webhook configuration")
	fs.StringVar(&o.MutatingWebhookName, "mutating-webhook-name", "", "The name of the mutating webhook configuration")
}
