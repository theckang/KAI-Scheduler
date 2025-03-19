// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package resource_info

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestReousrceInfo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Resource Info suite")
}
