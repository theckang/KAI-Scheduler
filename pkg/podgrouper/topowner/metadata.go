// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package topowner

import (
	"gopkg.in/yaml.v3"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
)

type Metadata struct {
	Name    string    `yaml:"name"`
	UID     types.UID `yaml:"uid"`
	Group   string    `yaml:"group"`
	Version string    `yaml:"version"`
	Kind    string    `yaml:"kind"`
}

func GetTopOwnerMetadata(topOwner *unstructured.Unstructured) Metadata {
	gvk := topOwner.GroupVersionKind()
	return Metadata{
		Name:    topOwner.GetName(),
		UID:     topOwner.GetUID(),
		Group:   gvk.Group,
		Version: gvk.Version,
		Kind:    gvk.Kind,
	}
}

func (md *Metadata) MarshalYAML() (string, error) {
	val, err := yaml.Marshal(md)
	return string(val), err
}
