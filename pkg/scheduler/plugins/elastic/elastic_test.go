// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package elastic

import (
	"testing"

	"github.com/NVIDIA/KAI-scheduler/pkg/scheduler/api/common_info"
	"github.com/NVIDIA/KAI-scheduler/pkg/scheduler/api/pod_info"
	"github.com/NVIDIA/KAI-scheduler/pkg/scheduler/api/pod_status"
	"github.com/NVIDIA/KAI-scheduler/pkg/scheduler/api/podgroup_info"
	"github.com/NVIDIA/KAI-scheduler/pkg/scheduler/framework"
)

func TestJobOrderFn(t *testing.T) {
	type args struct {
		l interface{}
		r interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "no pods",
			args: args{
				l: &podgroup_info.PodGroupInfo{
					PodInfos: map[common_info.PodID]*pod_info.PodInfo{},
				},
				r: &podgroup_info.PodGroupInfo{
					PodInfos: map[common_info.PodID]*pod_info.PodInfo{},
				},
			},
			want: 0,
		},
		{
			name: "running single pod",
			args: args{
				l: &podgroup_info.PodGroupInfo{
					MinAvailable: 1,
					PodInfos: map[common_info.PodID]*pod_info.PodInfo{
						"pod-a-1": {
							Name:   "pod-a-1",
							Status: pod_status.Running,
						},
					},
				},
				r: &podgroup_info.PodGroupInfo{
					MinAvailable: 1,
					PodInfos: map[common_info.PodID]*pod_info.PodInfo{
						"pod-a-2": {
							Name:   "pod-a-2",
							Status: pod_status.Running,
						},
					},
				},
			},
			want: 0,
		},
		{
			name: "allocated pod counts as allocated",
			args: args{
				l: &podgroup_info.PodGroupInfo{
					MinAvailable: 1,
					PodInfos: map[common_info.PodID]*pod_info.PodInfo{
						"pod-a-1": {
							Name:   "pod-a-1",
							Status: pod_status.Allocated,
						},
					},
				},
				r: &podgroup_info.PodGroupInfo{
					MinAvailable: 1,
					PodInfos: map[common_info.PodID]*pod_info.PodInfo{
						"pod-a-2": {
							Name:   "pod-a-2",
							Status: pod_status.Running,
						},
					},
				},
			},
			want: 0,
		},
		{
			name: "bound pod counts as allocated",
			args: args{
				l: &podgroup_info.PodGroupInfo{
					MinAvailable: 1,
					PodInfos: map[common_info.PodID]*pod_info.PodInfo{
						"pod-a-1": {
							Name:   "pod-a-1",
							Status: pod_status.Bound,
						},
					},
				},
				r: &podgroup_info.PodGroupInfo{
					MinAvailable: 1,
					PodInfos: map[common_info.PodID]*pod_info.PodInfo{
						"pod-a-2": {
							Name:   "pod-a-2",
							Status: pod_status.Running,
						},
					},
				},
			},
			want: 0,
		},
		{
			name: "releasing pod doesn't count as allocated",
			args: args{
				l: &podgroup_info.PodGroupInfo{
					MinAvailable: 1,
					PodInfos: map[common_info.PodID]*pod_info.PodInfo{
						"pod-a-1": {
							Name:   "pod-a-1",
							Status: pod_status.Releasing,
						},
					},
				},
				r: &podgroup_info.PodGroupInfo{
					MinAvailable: 1,
					PodInfos: map[common_info.PodID]*pod_info.PodInfo{
						"pod-a-2": {
							Name:   "pod-a-2",
							Status: pod_status.Running,
						},
					},
				},
			},
			want: -1,
		},
		{
			name: "pod group with min pods against pod group with no pods",
			args: args{
				l: &podgroup_info.PodGroupInfo{
					MinAvailable: 1,
					PodInfos: map[common_info.PodID]*pod_info.PodInfo{
						"pod-a-1": {
							Name:   "pod-a-1",
							Status: pod_status.Running,
						},
					},
				},
				r: &podgroup_info.PodGroupInfo{
					MinAvailable: 1,
					PodInfos:     map[common_info.PodID]*pod_info.PodInfo{},
				},
			},
			want: 1,
		},
		{
			name: "pod group with less then min pods against pod group with min pods",
			args: args{
				l: &podgroup_info.PodGroupInfo{
					MinAvailable: 2,
					PodInfos: map[common_info.PodID]*pod_info.PodInfo{
						"pod-a-1": {
							Name:   "pod-a-1",
							Status: pod_status.Running,
						},
					},
				},
				r: &podgroup_info.PodGroupInfo{
					MinAvailable: 2,
					PodInfos: map[common_info.PodID]*pod_info.PodInfo{
						"pod-a-2": {
							Name:   "pod-a-2",
							Status: pod_status.Running,
						},
						"pod-b-2": {
							Name:   "pod-b-2",
							Status: pod_status.Running,
						},
					},
				},
			},
			want: -1,
		},
		{
			name: "pod group with less then min pods against pod group with more than min pods",
			args: args{
				l: &podgroup_info.PodGroupInfo{
					MinAvailable: 2,
					PodInfos: map[common_info.PodID]*pod_info.PodInfo{
						"pod-a-1": {
							Name:   "pod-a-1",
							Status: pod_status.Running,
						},
					},
				},
				r: &podgroup_info.PodGroupInfo{
					MinAvailable: 2,
					PodInfos: map[common_info.PodID]*pod_info.PodInfo{
						"pod-a-2": {
							Name:   "pod-a-2",
							Status: pod_status.Running,
						},
						"pod-b-2": {
							Name:   "pod-b-2",
							Status: pod_status.Running,
						},
						"pod-b-3": {
							Name:   "pod-b-3",
							Status: pod_status.Running,
						},
					},
				},
			},
			want: -1,
		},
		{
			name: "pod group with min pods against pod group with more than min pods",
			args: args{
				l: &podgroup_info.PodGroupInfo{
					MinAvailable: 1,
					PodInfos: map[common_info.PodID]*pod_info.PodInfo{
						"pod-a-1": {
							Name:   "pod-a-1",
							Status: pod_status.Running,
						},
					},
				},
				r: &podgroup_info.PodGroupInfo{
					MinAvailable: 1,
					PodInfos: map[common_info.PodID]*pod_info.PodInfo{
						"pod-a-2": {
							Name:   "pod-a-2",
							Status: pod_status.Running,
						},
						"pod-b-2": {
							Name:   "pod-b-2",
							Status: pod_status.Running,
						},
					},
				},
			},
			want: -1,
		},
		{
			name: "pod group with min pods against pod group with less than min pods",
			args: args{
				l: &podgroup_info.PodGroupInfo{
					MinAvailable: 1,
					PodInfos: map[common_info.PodID]*pod_info.PodInfo{
						"pod-a-1": {
							Name:   "pod-a-1",
							Status: pod_status.Running,
						},
					},
				},
				r: &podgroup_info.PodGroupInfo{
					MinAvailable: 3,
					PodInfos: map[common_info.PodID]*pod_info.PodInfo{
						"pod-a-2": {
							Name:   "pod-a-2",
							Status: pod_status.Running,
						},
						"pod-b-2": {
							Name:   "pod-b-2",
							Status: pod_status.Running,
						},
					},
				},
			},
			want: 1,
		},
		{
			name: "pod group with more than min pods against pod group with min pods",
			args: args{
				l: &podgroup_info.PodGroupInfo{
					MinAvailable: 1,
					PodInfos: map[common_info.PodID]*pod_info.PodInfo{
						"pod-a-1": {
							Name:   "pod-a-1",
							Status: pod_status.Running,
						},
						"pod-a-2": {
							Name:   "pod-a-2",
							Status: pod_status.Running,
						},
					},
				},
				r: &podgroup_info.PodGroupInfo{
					MinAvailable: 1,
					PodInfos: map[common_info.PodID]*pod_info.PodInfo{
						"pod-a-2": {
							Name:   "pod-a-2",
							Status: pod_status.Running,
						},
					},
				},
			},
			want: 1,
		},
		{
			name: "pod group with more than min pods against pod group with less than min pods",
			args: args{
				l: &podgroup_info.PodGroupInfo{
					MinAvailable: 1,
					PodInfos: map[common_info.PodID]*pod_info.PodInfo{
						"pod-a-1": {
							Name:   "pod-a-1",
							Status: pod_status.Running,
						},
						"pod-a-2": {
							Name:   "pod-a-2",
							Status: pod_status.Running,
						},
					},
				},
				r: &podgroup_info.PodGroupInfo{
					MinAvailable: 1,
					PodInfos: map[common_info.PodID]*pod_info.PodInfo{
						"pod-a-2": {
							Name:   "pod-a-2",
							Status: pod_status.Running,
						},
					},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JobOrderFn(tt.args.l, tt.args.r); got != tt.want {
				t.Errorf("JobOrderFn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_elasticPlugin_OnSessionOpen(t *testing.T) {
	type fields struct {
		pluginArguments map[string]string
	}
	type args struct {
		ssn *framework.Session
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		jobOrderSetup map[string]common_info.CompareFn
	}{
		{
			name: "Sets the elastic job order func",
			fields: fields{
				pluginArguments: map[string]string{},
			},
			args: args{
				ssn: &framework.Session{},
			},
			jobOrderSetup: map[string]common_info.CompareFn{
				"elastic": JobOrderFn,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pp := &elasticPlugin{
				pluginArguments: tt.fields.pluginArguments,
			}
			pp.OnSessionOpen(tt.args.ssn)
		})
	}
}
