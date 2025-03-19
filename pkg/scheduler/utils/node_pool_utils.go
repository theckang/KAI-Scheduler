// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	commonconstants "github.com/NVIDIA/KAI-scheduler/pkg/common/constants"
)

func GetNodePoolNameFromLabels(labels map[string]string) string {
	nodePoolName, found := labels[commonconstants.NodePoolNameLabel]
	if !found || nodePoolName == "" {
		nodePoolName = commonconstants.DefaultNodePoolName
	}

	return nodePoolName
}
