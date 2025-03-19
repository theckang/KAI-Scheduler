// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package group_mutex

import (
	"sync"
)

type GroupMutex struct {
	mapMutex     sync.Mutex
	mutexMap     map[string]*sync.Mutex
	mutexRefsMap map[string]int
}

func NewGroupMutex() *GroupMutex {
	return &GroupMutex{
		mapMutex:     sync.Mutex{},
		mutexMap:     map[string]*sync.Mutex{},
		mutexRefsMap: map[string]int{},
	}
}

func (gm *GroupMutex) LockMutexForGroup(group string) {
	mutex := gm.acquireWithRefcountIncrease(group)
	mutex.Lock()
}

func (gm *GroupMutex) ReleaseMutex(group string) {
	mutex := gm.acquireWithRefcountDecrease(group)
	if mutex != nil {
		mutex.Unlock()
	}
}

func (gm *GroupMutex) acquireWithRefcountIncrease(group string) *sync.Mutex {
	gm.mapMutex.Lock()
	defer gm.mapMutex.Unlock()

	mutex, found := gm.mutexMap[group]
	if !found {
		mutex = &sync.Mutex{}
		gm.mutexMap[group] = mutex
	}

	gm.mutexRefsMap[group] = gm.mutexRefsMap[group] + 1

	return mutex
}

func (gm *GroupMutex) acquireWithRefcountDecrease(group string) *sync.Mutex {
	gm.mapMutex.Lock()
	defer gm.mapMutex.Unlock()

	mutex, found := gm.mutexMap[group]
	if !found {
		return nil
	}

	gm.mutexRefsMap[group] = gm.mutexRefsMap[group] - 1
	if gm.mutexRefsMap[group] == 0 {
		delete(gm.mutexMap, group)
		delete(gm.mutexRefsMap, group)
	}

	return mutex
}
