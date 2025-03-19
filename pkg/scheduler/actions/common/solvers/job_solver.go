// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package solvers

import (
	"fmt"
	"strings"

	"github.com/NVIDIA/KAI-scheduler/pkg/scheduler/actions/utils"
	"github.com/NVIDIA/KAI-scheduler/pkg/scheduler/api/node_info"
	"github.com/NVIDIA/KAI-scheduler/pkg/scheduler/api/pod_info"
	"github.com/NVIDIA/KAI-scheduler/pkg/scheduler/api/podgroup_info"
	"github.com/NVIDIA/KAI-scheduler/pkg/scheduler/framework"
	"github.com/NVIDIA/KAI-scheduler/pkg/scheduler/log"
	"github.com/NVIDIA/KAI-scheduler/pkg/scheduler/metrics"
)

type GenerateVictimsQueue func() *utils.JobsOrderByQueues

type JobSolver struct {
	feasibleNodes        []*node_info.NodeInfo
	solutionValidator    SolutionValidator
	generateVictimsQueue GenerateVictimsQueue
	actionType           framework.ActionType
}

type solvingState struct {
	pendingTasks   []*pod_info.PodInfo
	numTasksSolved int

	statement            *framework.Statement
	recordedVictimsJobs  []*podgroup_info.PodGroupInfo
	recordedVictimsTasks []*pod_info.PodInfo
}

func NewJobsSolver(
	feasibleNodes []*node_info.NodeInfo,
	solutionValidator SolutionValidator,
	generateVictimsQueue GenerateVictimsQueue,
	action framework.ActionType,
) *JobSolver {
	return &JobSolver{
		feasibleNodes:        feasibleNodes,
		solutionValidator:    solutionValidator,
		generateVictimsQueue: generateVictimsQueue,
		actionType:           action,
	}
}

func (s *JobSolver) Solve(
	ssn *framework.Session, pendingJob *podgroup_info.PodGroupInfo) (bool, *framework.Statement, []string) {
	state := solvingState{
		recordedVictimsJobs: make([]*podgroup_info.PodGroupInfo, 0),
		pendingTasks:        make([]*pod_info.PodInfo, 0),
		numTasksSolved:      0,
	}

	tasksToAllocate := podgroup_info.GetTasksToAllocate(pendingJob, ssn.TaskOrderFn, false)
	for _, nextTaskToSolve := range tasksToAllocate {
		state.pendingTasks = append(state.pendingTasks, nextTaskToSolve)
		partialPendingJob := getPartialJobRepresentative(pendingJob, state.pendingTasks)

		scenarioBuilder := NewPodAccumulatedScenarioBuilder(
			ssn, partialPendingJob, state.recordedVictimsJobs, s.generateVictimsQueue())

		feasibleNodeMap := map[string]*node_info.NodeInfo{}
		for _, node := range s.feasibleNodes {
			feasibleNodeMap[node.Name] = node
		}
		// recorded victim jobs nodes
		for _, task := range state.recordedVictimsTasks {
			node := ssn.Nodes[task.NodeName]
			feasibleNodeMap[task.NodeName] = node
		}

		for scenarioToSolve := scenarioBuilder.GetCurrentScenario(); scenarioToSolve != nil; scenarioToSolve =
			scenarioBuilder.GetNextScenario() {
			returnSolvingStatement := len(state.pendingTasks) == len(tasksToAllocate)
			scenarioSolver := newByPodSolver(feasibleNodeMap, s.solutionValidator, ssn.AllowConsolidatingReclaim(),
				s.actionType, returnSolvingStatement)

			log.InfraLogger.V(5).Infof("Trying to solve scenario: %s", scenarioToSolve.String())
			metrics.IncScenarioSimulatedByAction()

			if result := scenarioSolver.solve(ssn, scenarioToSolve); result.solved {
				state = solvingState{
					state.pendingTasks,
					state.numTasksSolved + 1,
					result.statement,
					result.victimJobs,
					result.victimsTasks,
				}

				log.InfraLogger.V(4).Infof(
					"Scenario solved for %d tasks out of %d tasks to allocate for %s. Victims: %s",
					len(state.pendingTasks), len(tasksToAllocate), pendingJob.Name,
					victimPrintingStruct{result.victimsTasks})
				break
			}
		}
	}

	allTasksSolved := state.numTasksSolved == len(tasksToAllocate)
	return allTasksSolved, state.statement, calcVictimNames(state.recordedVictimsTasks)
}

func getPartialJobRepresentative(
	job *podgroup_info.PodGroupInfo, pendingTasks []*pod_info.PodInfo) *podgroup_info.PodGroupInfo {
	jobRepresentative := job.CloneWithTasks(pendingTasks)
	jobRepresentative.MinAvailable = int32(len(pendingTasks))
	return jobRepresentative
}

func calcVictimNames(victimsTasks []*pod_info.PodInfo) []string {
	var names []string
	for _, victimTask := range victimsTasks {
		names = append(names,
			fmt.Sprintf("<%s/%s>", victimTask.Namespace, victimTask.Name))
	}
	return names
}

type victimPrintingStruct struct {
	victims []*pod_info.PodInfo
}

func (v victimPrintingStruct) String() string {
	if len(v.victims) == 0 {
		return ""
	}
	stringBuilder := strings.Builder{}

	stringBuilder.WriteString(v.victims[0].Namespace)
	stringBuilder.WriteString("/")
	stringBuilder.WriteString(v.victims[0].Name)

	for _, victimTask := range v.victims[1:] {
		stringBuilder.WriteString(", ")
		stringBuilder.WriteString(victimTask.Namespace)
		stringBuilder.WriteString("/")
		stringBuilder.WriteString(victimTask.Name)
	}

	return stringBuilder.String()
}
