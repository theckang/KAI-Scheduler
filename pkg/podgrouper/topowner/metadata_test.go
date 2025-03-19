// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package topowner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetadataStringYaml(t *testing.T) {
	md := Metadata{
		Name:    "bla",
		UID:     "1",
		Group:   "scheduling.run.ai",
		Version: "v2alpha2",
		Kind:    "hantar",
	}
	expectedString := "name: bla\n" +
		"uid: \"1\"\n" +
		"group: scheduling.run.ai\n" +
		"version: v2alpha2\n" +
		"kind: hantar\n"

	val, err := md.MarshalYAML()
	if err != nil {
		t.Errorf("error marshaling metadata: %v", err)
	}

	assert.Equal(t, val, expectedString)
}
