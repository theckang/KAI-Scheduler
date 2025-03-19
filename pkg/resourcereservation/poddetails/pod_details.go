// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package poddetails

import (
	"fmt"
	"os"
)

const (
	podNameEnvVar      = "POD_NAME"
	podNamespaceEnvVar = "POD_NAMESPACE"
)

func CurrentPodDetails() (string, string, error) {
	namespace, found := os.LookupEnv(podNamespaceEnvVar)
	if !found {
		return "", "", fmt.Errorf("could not find pod namespace")
	}

	name, found := os.LookupEnv(podNameEnvVar)
	if !found {
		return "", "", fmt.Errorf("could not find pod name")
	}

	return namespace, name, nil
}
