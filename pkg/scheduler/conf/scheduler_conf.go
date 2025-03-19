// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package conf

import (
	"time"

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
)

type SchedulingNodePoolParams struct {
	NodePoolLabelKey   string
	NodePoolLabelValue string
}

type SchedulerParams struct {
	SchedulerName                     string
	RestrictSchedulingNodes           bool
	PartitionParams                   *SchedulingNodePoolParams
	IsInferencePreemptible            bool
	MaxNumberConsolidationPreemptees  int
	ScheduleCSIStorage                bool
	UseSchedulingSignatures           bool
	FullHierarchyFairness             bool
	NodeLevelScheduler                bool
	AllowConsolidatingReclaim         bool
	NumOfStatusRecordingWorkers       int
	GlobalDefaultStalenessGracePeriod time.Duration
	SchedulePeriod                    time.Duration
	DetailedFitErrors                 bool
}

// SchedulerConfiguration defines the configuration of scheduler.
type SchedulerConfiguration struct {
	// Actions defines the actions list of scheduler in order
	Actions string `yaml:"actions"`
	// Tiers defines plugins in different tiers
	Tiers []Tier `yaml:"tiers,omitempty"`
	// QueueDepthPerAction max number of jobs to try for action per queue
	QueueDepthPerAction map[string]int `yaml:"queueDepthPerAction,omitempty"`
}

// Tier defines plugin tier
type Tier struct {
	Plugins []PluginOption `yaml:"plugins"`
}

// PluginOption defines the options of plugin
type PluginOption struct {
	// The name of Plugin
	Name string `yaml:"name"`
	// JobOrderDisabled defines whether jobOrderFn is disabled
	JobOrderDisabled bool `yaml:"disableJobOrder"`
	// TaskOrderDisabled defines whether taskOrderFn is disabled
	TaskOrderDisabled bool `yaml:"disableTaskOrder"`
	// PreemptableDisabled defines whether preemptableFn is disabled
	PreemptableDisabled bool `yaml:"disablePreemptable"`
	// ReclaimableDisabled defines whether reclaimableFn is disabled
	ReclaimableDisabled bool `yaml:"disableReclaimable"`
	// QueueOrderDisabled defines whether queueOrderFn is disabled
	QueueOrderDisabled bool `yaml:"disableQueueOrder"`
	// PredicateDisabled defines whether predicateFn is disabled
	PredicateDisabled bool `yaml:"disablePredicate"`
	// NodeOrderDisabled defines whether NodeOrderFn is disabled
	NodeOrderDisabled bool `yaml:"disableNodeOrder"`
	// Arguments defines the different arguments that can be given to different plugins
	Arguments map[string]string `yaml:"arguments"`
}

func (s *SchedulingNodePoolParams) GetLabelSelector() (labels.Selector, error) {
	if s.NodePoolLabelKey == "" {
		return labels.Everything(), nil
	}
	operator := selection.DoesNotExist
	var vals []string
	if len(s.NodePoolLabelValue) > 0 {
		operator = selection.Equals
		vals = []string{s.NodePoolLabelValue}
	}

	requirement, err := labels.NewRequirement(s.NodePoolLabelKey, operator, vals)
	if err != nil {
		return nil, err
	}
	selector := labels.NewSelector().Add(*requirement)
	return selector, nil
}

func (s *SchedulingNodePoolParams) GetLabels() map[string]string {
	if s.NodePoolLabelKey == "" || s.NodePoolLabelValue == "" {
		return map[string]string{}
	}
	return map[string]string{s.NodePoolLabelKey: s.NodePoolLabelValue}
}
