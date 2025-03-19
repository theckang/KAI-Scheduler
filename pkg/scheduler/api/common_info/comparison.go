// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package common_info

// LessFn is the func declaration used by sort or priority queue.
type LessFn func(interface{}, interface{}) bool

// CompareFn is the func declaration used by sort or priority queue.
type CompareFn func(interface{}, interface{}) int

// CompareTwoFactors is the func declaration used by sort or priority queue for 2 factors.
type CompareTwoFactors func(interface{}, interface{}, interface{}, interface{}) int
