// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package common_info

import "fmt"

type ObjectKey string

const (
	separator = "/"
)

func NewObjectKey(namespace, name string) ObjectKey {
	if namespace == "" {
		return ObjectKey(name)
	}
	return ObjectKey(fmt.Sprintf("%s%s%s", namespace, separator, name))
}
