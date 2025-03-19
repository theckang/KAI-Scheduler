# Workload Priority
In Kubernetes, assigning different priorities to workloads ensures efficient resource management, minimizes service disruption, and supports better scaling. 
It allows critical applications to receive the resources they need before less important ones, ensuring high availability and compliance with SLAs. 
By prioritizing workloads, KAI Scheduler can preempt lower-priority workloads to make room for higher-priority ones. 
This approach ensures that mission-critical services are always prioritized in resource allocation.

## API
The workload priority is determined as follows:
1. By examining the `priorityClassName` label on the workload object.
2. By checking the `priorityClassName` label on one of the workload's pods.
3. By checking `pod.Spec.PriorityClassName` from one of the pods.

## Preemtibility
KAI Scheduler supports any PriorityClass deployed in the cluster. PriorityClass with a value of 100 or higher will be considered as non-preemptible.

## Workload Default Priority
When priorityClass is not provided, KAI Scheduler will use a predefined default priority based on the workload type:
1. Inference for K8s Deployment or Knative Service
2. Build for Kubeflow Notebook
3. Train as the general default
If pods from the same workload have different priorities, the workload's priority is derived from any of its pods.

## Usability
Workload priorities serve three main purposes:
1. Higher priority workloads are scheduled first within the same queue.
2. In case of insufficient cluster resources, lower priority workloads can be evicted to prioritize higher priority queues.
3. Workloads with build or inference priority are not preemptible (like train workloads), hence they can only run within queue deserved quota.

## Example
To submit a pod with build priority, use the following command:
```
kubectl apply -f pods/build-priority-pod.yaml
```
