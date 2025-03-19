// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package group_mutex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLockSingleGpuGroupMutex(t *testing.T) {
	groupMutex := NewGroupMutex()
	gpuGroup := "gpu-group-1"

	groupMutex.LockMutexForGroup(gpuGroup)

	assert.NotNil(t, groupMutex.mutexMap[gpuGroup])
	assert.Equal(t, 1, groupMutex.mutexRefsMap[gpuGroup])
}

func TestReleaseSingleGpuGroupMutex(t *testing.T) {
	groupMutex := NewGroupMutex()
	gpuGroup := "gpu-group-1"

	groupMutex.LockMutexForGroup(gpuGroup)
	groupMutex.ReleaseMutex(gpuGroup)

	assert.Nil(t, groupMutex.mutexMap[gpuGroup])
	assert.Zero(t, groupMutex.mutexRefsMap[gpuGroup])
}

func TestMultipleGpuGroupMutexes(t *testing.T) {
	groupMutex := NewGroupMutex()
	gpuGroup1 := "gpu-group-1"
	gpuGroup2 := "gpu-group-2"

	groupMutex.LockMutexForGroup(gpuGroup1)
	groupMutex.LockMutexForGroup(gpuGroup2)

	group1Mutex := groupMutex.mutexMap[gpuGroup1]
	group2Mutex := groupMutex.mutexMap[gpuGroup2]
	assert.NotNil(t, group1Mutex)
	assert.NotNil(t, group2Mutex)
	assert.False(t, group1Mutex == group2Mutex)

	assert.Equal(t, 1, groupMutex.mutexRefsMap[gpuGroup1])
	assert.Equal(t, 1, groupMutex.mutexRefsMap[gpuGroup2])

	// release group1 mutex
	groupMutex.ReleaseMutex(gpuGroup1)
	assert.Nil(t, groupMutex.mutexMap[gpuGroup1])
	assert.Zero(t, groupMutex.mutexRefsMap[gpuGroup1])
	assert.NotNil(t, groupMutex.mutexMap[gpuGroup2])
	assert.Equal(t, 1, groupMutex.mutexRefsMap[gpuGroup2])

	// release group2 mutex
	groupMutex.ReleaseMutex(gpuGroup2)
	assert.Nil(t, groupMutex.mutexMap[gpuGroup2])
	assert.Zero(t, groupMutex.mutexRefsMap[gpuGroup2])

	assert.Zero(t, len(groupMutex.mutexMap))
	assert.Zero(t, len(groupMutex.mutexRefsMap))
}
