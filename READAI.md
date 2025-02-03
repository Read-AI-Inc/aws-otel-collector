

# ReadAI customized aws-otel-collector 

This repository is forked from the https://github.com/aws-observability/aws-otel-collector repo. 

We've added our own readai-config.yaml to allow us to customize awscollector settings. 


## Example sidecar usage:

Below is an example of an ECS task definition that runs the aws-otel-collector as a sidecar.

```json
{
  "containerDefinitions": [
    {
      "name": "aws-otel-collector",
      "image": "982262549137.dkr.ecr.us-east-1.amazonaws.com/readai-customized/awscollector:v0.42.0",
      "command": [
        "--config=/etc/ecs/readai-config.yaml"
      ],
      "essential": false,
      "logConfiguration": {
        "logDriver": "awslogs",
        "options": {
          "awslogs-group": "/ecs/ecs-aws-otel-sidecar-collector",
          "awslogs-region": "us-east-1",
          "awslogs-stream-prefix": "ecs",
          "awslogs-create-group": "True"
        }
      },
      "healthCheck": {
        "command": [
          "/healthcheck"
        ],
        "interval": 5,
        "timeout": 6,
        "retries": 5,
        "startPeriod": 1
      }
    },
    {
      "name": $name,
      "cpu": 0,
      "links": [],
      "stopTimeout": $stopTimeout,
      "portMappings": [
        {
          "containerPort": 8000,
          "hostPort": 8000,
          "protocol": "tcp"
        }
      ],
      "essential": true,
      "entryPoint": $entryPoint,
      "command": [],
      "environment": [],
      "mountPoints": [],
      "volumesFrom": [],
      "logConfiguration": {
        "logDriver": "awslogs",
        "options": {
          "mode": "non-blocking",
          "max-buffer-size": "25m",
          "awslogs-group": $logGroup,
          "awslogs-region": "us-east-1",
          "awslogs-stream-prefix": "ecs"
        }
      },
      "dependsOn": [
        {
          "containerName": "aws-otel-collector",
          "condition": "START"
        }
      ]
    }
  ],
  "family": $family,
  "taskRoleArn": $roleARN,
  "executionRoleArn": $roleARN,
  "networkMode": "awsvpc",
  "volumes": [],
  "placementConstraints": [],
  "requiresCompatibilities": [
    $compatibility
  ],
  "cpu": $taskCpu,
  "memory": $taskMemory,
  "tags": [
    {
      "key": "Name",
      "value": "application-backend"
    },
    {
      "key": "Environment",
      "value": $deployEnvironment
    }
  ],
  "runtimePlatform": {
    "operatingSystemFamily": "LINUX",
    "cpuArchitecture": "ARM64"
  }
}

```

## Build + Deploy

We do not currently have a CICD pipeline configured for this project. Building and deploying an image 
must be done locally. 

Images must be published to our root AWS account's ECR repository. 

To build and deploy an image, run 
```shell
git checkout v0.42.0  # or whatever release you want to build & deploy from
git checkout -b v0.42.0-readai
git cherry-pick foobar # cherry pick commits you want to add to our release branch
git push origin v0.42.0-readai

make docker-build-arm
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 982262549137.dkr.ecr.us-east-1.amazonaws.com
docker push 982262549137.dkr.ecr.us-east-1.amazonaws.com/readai-customized/awscollector:v0.42.0 # or whatever tag you want
```