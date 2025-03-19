// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package binding

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBinder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "binder tests")
}
