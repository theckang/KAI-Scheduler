// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package discovery

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	gpusDir = "/proc/driver/nvidia/gpus/"
)

func GetGPUDevice(ctx context.Context, podName string, namespace string) (string, error) {
	logger := log.FromContext(ctx)
	logger.Info("Getting GPU device id for pod", "namespace", namespace, "name", podName)

	deviceSubDirs, err := os.ReadDir(gpusDir)
	if err != nil {
		return "", fmt.Errorf("failed to read GPU devices dir: %v", err)
	}

	for _, subDir := range deviceSubDirs {
		infoFilePath := filepath.Join(gpusDir, subDir.Name(), "information")

		logger.Info("Getting GPU device info", "path", infoFilePath)

		data, err := os.ReadFile(infoFilePath)
		if err != nil {
			return "", fmt.Errorf("failed to read GPU device information: %v", err)
		}

		content := string(data)
		logger.Info("GPU device info", "content", content)

		lines := strings.Split(content, "\n")
		for _, line := range lines {
			if strings.Contains(line, "GPU UUID") {
				words := strings.Fields(line)
				return words[len(words)-1], nil
			}
		}
	}

	return "", fmt.Errorf("failed to find GPU UUID")
}
