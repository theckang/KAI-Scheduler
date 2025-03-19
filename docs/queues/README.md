# Scheduling Queues
A queue is an object which represents a job queue in the cluster. Queues are an essential scheduling primitive, and can reflect different scheduling guarantees, such as resource quota and priority.
KAI Scheduler offers several features that make it easier to optimize resource distribution among researchers in a way that maximizes resource fairness between queues.

The following attributes can be applied to any scheduling queue for CPU, memory and GPU resources:
* Quota: The designated amount of resources allocated to the scheduling queue upon request.
* Over-Quota Priority: Determines the order in which queues will receive resources that exceed their allocated quota.
* Over-Quota Weight: Dictates how excess resources are distributed among queues within the same priority level.
* Limit: The hard cap on the amount of resources a scheduling queue can consume.