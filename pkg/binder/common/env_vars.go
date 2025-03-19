// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package common

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/utils/ptr"
)

func AddDirectEnvVarsConfigMapSource(container *v1.Container, directEnvVarsMapName string) {
	for _, env := range container.EnvFrom {
		if env.ConfigMapRef != nil && env.ConfigMapRef.Name == directEnvVarsMapName {
			return
		}
	}

	container.EnvFrom = append(container.EnvFrom, v1.EnvFromSource{
		ConfigMapRef: &v1.ConfigMapEnvSource{
			LocalObjectReference: v1.LocalObjectReference{
				Name: directEnvVarsMapName,
			},
			Optional: ptr.To(false),
		},
	})
}

func AddEnvVarToContainer(container *v1.Container, envVar v1.EnvVar) {
	for i, env := range container.Env {
		if env.Name == envVar.Name {
			container.Env[i].Value = envVar.Value
			return
		}
	}

	container.Env = append(container.Env, envVar)
}
