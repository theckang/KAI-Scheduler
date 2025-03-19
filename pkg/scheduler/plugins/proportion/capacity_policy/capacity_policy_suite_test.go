// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package capacity_policy

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCapacityPolicy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Capacity Policy Suite")
}
