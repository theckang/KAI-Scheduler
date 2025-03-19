// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package resource_share

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueResourceShare_ResourceShare(t *testing.T) {
	qrs := createQueueResourceShare()
	memoryResourceShare := qrs.ResourceShare(MemoryResource)
	assert.Equal(t, qrs.Memory.Deserved, memoryResourceShare.Deserved)
	assert.Equal(t, qrs.Memory.FairShare, memoryResourceShare.FairShare)
	assert.Equal(t, qrs.Memory.MaxAllowed, memoryResourceShare.MaxAllowed)
	assert.Equal(t, qrs.Memory.OverQuotaWeight, memoryResourceShare.OverQuotaWeight)
	assert.Equal(t, qrs.Memory.Allocated, memoryResourceShare.Allocated)
	assert.Equal(t, qrs.Memory.AllocatedNotPreemptible, memoryResourceShare.AllocatedNotPreemptible)
	assert.Equal(t, qrs.Memory.Request, memoryResourceShare.Request)
}

func TestQueueResourceShare_ResourceShare_Modify(t *testing.T) {
	qrs := createQueueResourceShare()
	share := qrs.ResourceShare(MemoryResource)
	qrs.Memory.Deserved = 123
	assert.Equal(t, qrs.Memory.Deserved, share.Deserved)
}

func TestQueueResourceShare_ResourceShareUnknownResource(t *testing.T) {
	qrs := createQueueResourceShare()
	share := qrs.ResourceShare("unknown")
	assert.Nil(t, share)
}

func TestQueueResourceShare_GetAllocatableShare(t *testing.T) {
	qrs := createQueueResourceShare()
	share := qrs.GetAllocatableShare()
	assert.Equal(t, float64(2), share[CpuResource])
	assert.Equal(t, float64(9), share[MemoryResource])
	assert.Equal(t, float64(16), share[GpuResource])
}

func TestQueueResourceShare_GetFairShare(t *testing.T) {
	qrs := createQueueResourceShare()
	share := qrs.GetFairShare()
	assert.Equal(t, float64(2), share[CpuResource])
	assert.Equal(t, float64(9), share[MemoryResource])
	assert.Equal(t, float64(16), share[GpuResource])
}

func TestQueueResourceShare_GetDeservedShare(t *testing.T) {
	qrs := createQueueResourceShare()
	share := qrs.GetDeservedShare()
	assert.Equal(t, float64(1), share[CpuResource])
	assert.Equal(t, float64(8), share[MemoryResource])
	assert.Equal(t, float64(15), share[GpuResource])
}

func TestQueueResourceShare_GetAllocatedShare(t *testing.T) {
	qrs := createQueueResourceShare()
	share := qrs.GetAllocatedShare()
	assert.Equal(t, float64(5), share[CpuResource])
	assert.Equal(t, float64(12), share[MemoryResource])
	assert.Equal(t, float64(19), share[GpuResource])
}

func TestQueueResourceShare_GetAllocatedNonPreemptible(t *testing.T) {
	qrs := createQueueResourceShare()
	share := qrs.GetAllocatedNonPreemptible()
	assert.Equal(t, float64(6), share[CpuResource])
	assert.Equal(t, float64(13), share[MemoryResource])
	assert.Equal(t, float64(20), share[GpuResource])
}

func TestQueueResourceShare_GetMaxAllowedShare(t *testing.T) {
	qrs := createQueueResourceShare()
	share := qrs.GetMaxAllowedShare()
	assert.Equal(t, float64(3), share[CpuResource])
	assert.Equal(t, float64(10), share[MemoryResource])
	assert.Equal(t, float64(17), share[GpuResource])
}

func TestQueueResourceShare_GetRequestShare(t *testing.T) {
	qrs := createQueueResourceShare()
	share := qrs.GetRequestShare()
	assert.Equal(t, float64(7), share[CpuResource])
	assert.Equal(t, float64(14), share[MemoryResource])
	assert.Equal(t, float64(21), share[GpuResource])
}

func TestQueueResourceShare_GetRequestedShare(t *testing.T) {
	qrs := createQueueResourceShare()
	share := qrs.GetRequestableShare()
	assert.Equal(t, float64(3), share[CpuResource])
	assert.Equal(t, float64(10), share[MemoryResource])
	assert.Equal(t, float64(17), share[GpuResource])
}

func TestQueueResourceShare_GetDominantResourceShare(t *testing.T) {
	qrs := createQueueResourceShare()
	share := qrs.GetDominantResourceShare(ResourceQuantities{})
	assert.Equal(t, 2.5, share)
}

func TestQueueResourceShare_AddResourceShare(t *testing.T) {
	qrs := createQueueResourceShare()
	qrs.AddResourceShare(GpuResource, 1)
	assert.Equal(t, float64(17), qrs.GPU.FairShare)
}

func TestQueueResourceShare_SetQuotaResources(t *testing.T) {
	qrs := createQueueResourceShare()
	qrs.SetQuotaResources(CpuResource, 8.8, 9.9, 1.1)
	assert.Equal(t, 8.8, qrs.CPU.Deserved)
	assert.Equal(t, 9.9, qrs.CPU.MaxAllowed)
	assert.Equal(t, 1.1, qrs.CPU.OverQuotaWeight)
}

func createQueueResourceShare() *QueueResourceShare {
	return &QueueResourceShare{
		CPU: ResourceShare{
			Deserved:                1,
			FairShare:               2,
			MaxAllowed:              3,
			OverQuotaWeight:         4,
			Allocated:               5,
			AllocatedNotPreemptible: 6,
			Request:                 7,
		},
		Memory: ResourceShare{
			Deserved:                8,
			FairShare:               9,
			MaxAllowed:              10,
			OverQuotaWeight:         11,
			Allocated:               12,
			AllocatedNotPreemptible: 13,
			Request:                 14,
		},
		GPU: ResourceShare{
			Deserved:                15,
			FairShare:               16,
			MaxAllowed:              17,
			OverQuotaWeight:         18,
			Allocated:               19,
			AllocatedNotPreemptible: 20,
			Request:                 21,
		},
	}
}
