// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package framework

import "sync"

var pluginMutex sync.Mutex

type PluginBuilder func(map[string]string) Plugin

// Plugin management
var pluginBuilders = map[string]PluginBuilder{}

func RegisterPluginBuilder(name string, pc func(map[string]string) Plugin) {
	pluginMutex.Lock()
	defer pluginMutex.Unlock()

	pluginBuilders[name] = pc
}

func GetPluginBuilder(name string) (PluginBuilder, bool) {
	pluginMutex.Lock()
	defer pluginMutex.Unlock()

	pb, found := pluginBuilders[name]
	return pb, found
}

// Action management
var actionMap = map[ActionType]Action{}

func RegisterAction(act Action) {
	pluginMutex.Lock()
	defer pluginMutex.Unlock()

	actionMap[act.Name()] = act
}

func GetAction(name string) (Action, bool) {
	pluginMutex.Lock()
	defer pluginMutex.Unlock()

	act, found := actionMap[ActionType(name)]
	return act, found
}
