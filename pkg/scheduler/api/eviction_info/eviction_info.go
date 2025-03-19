// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package eviction_info

import (
	"k8s.io/apimachinery/pkg/types"
)

type EvictionMetadata struct {
	EvictionGangSize int
	Action           string
	Preemptor        *types.NamespacedName
}
