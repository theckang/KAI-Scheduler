// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
)

func TestAddConfigMapVolume(t *testing.T) {
	podSpec := &v1.PodSpec{}
	addConfigMapVolume(podSpec, "vol", "conf")
	assert.Equal(t, len(podSpec.Volumes), 1)
	assert.Equal(t, podSpec.Volumes[0].Name, "vol")
	assert.Equal(t, podSpec.Volumes[0].ConfigMap.LocalObjectReference.Name, "conf")
}

func TestOverrideConfigMapVolume(t *testing.T) {
	podSpec := &v1.PodSpec{
		Volumes: []v1.Volume{
			{
				Name: "name",
				VolumeSource: v1.VolumeSource{
					ConfigMap: &v1.ConfigMapVolumeSource{
						LocalObjectReference: v1.LocalObjectReference{
							Name: "conf1",
						},
					},
				},
			},
		},
	}

	addConfigMapVolume(podSpec, "name", "conf2")
	assert.Equal(t, len(podSpec.Volumes), 1)
	assert.Equal(t, podSpec.Volumes[0].Name, "name")
	assert.Equal(t, podSpec.Volumes[0].ConfigMap.LocalObjectReference.Name, "conf2")
}
