# Building from Source
To build and deploy KAI Scheduler from source, follow these steps:

1. Clone the repository:
   ```sh
   git clone git@github.com:NVIDIA/KAI-scheduler.git
   cd KAI-scheduler
   ```

2. Build the container images:
   ```sh
   make build
   ```

3. Package the Helm chart:
   ```sh
   helm package ./deployments/kai-scheduler -d ./charts --version 0.0.0-devel
   ```
   
4. Make sure the images are in loaded to cluster nodes cache

5. Install on your cluster:
   ```sh
   helm upgrade -i kai-scheduler -n kai-scheduler --create-namespace ./charts/kai-scheduler-0.0.0-devel.tgz
   ```