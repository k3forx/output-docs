# Argo Rollouts

## What is Argo Rollouts?

Argo Rollouts provides advanced deployment capabilities such as blue-green, canary, canary analysis, experimentation, and progressive delivery features to Kubernetes.

## Concepts

### Rollout

### Progressive Delivery

### Deployment Strategies

Argo Rollouts supports the following types of deployment strategies

#### Rolling Update

#### Recreate

#### Blue-Green

#### Canary

# Demo

## Installation

```bash
kubectl create ns argo-rollouts

kubectl apply -n argo-rollouts -f https://raw.githubusercontent.com/argoproj/argo-rollouts/stable/manifests/install.yaml

kubectl get all -n argo-rollouts

brew install argoproj/tap/kubectl-argo-rollouts
```

## Create project and application

Before you deploy rollout, you can create a project and an application for that.

```bash
kubectl create ns server

kubectl apply -f demo/project.yaml

kubectl apply -f demo/application.yaml
```

## Deploy rollout

With Argo CD, you can sync and check the status of `nginx` rollout.

```bash
kubectl argo rollouts get rollout nginx -n server
Name:            nginx
Namespace:       server
Status:          ✔ Healthy
Strategy:        Canary
  Step:          8/8
  SetWeight:     100
  ActualWeight:  100
Images:          nginx:1.19 (stable)
Replicas:
  Desired:       5
  Current:       5
  Updated:       5
  Ready:         5
  Available:     5

NAME                               KIND        STATUS     AGE    INFO
⟳ nginx                            Rollout     ✔ Healthy  8m53s
└──# revision:1
   └──⧉ nginx-846f597579           ReplicaSet  ✔ Healthy  8m53s  stable
      ├──□ nginx-846f597579-4rg7g  Pod         ✔ Running  8m53s  ready:1/1
      ├──□ nginx-846f597579-b5t4x  Pod         ✔ Running  8m53s  ready:1/1
      ├──□ nginx-846f597579-gmqpc  Pod         ✔ Running  8m53s  ready:1/1
      ├──□ nginx-846f597579-hdtcs  Pod         ✔ Running  8m53s  ready:1/1
      └──□ nginx-846f597579-vmdrb  Pod         ✔ Running  8m53s  ready:1/1
```

## Update image

You can update the image with `kubectl argo rollouts` command.

```bash
kubectl argo rollouts set image nginx nginx=nginx:1.19-alpine -n server

kubectl argo rollouts get rollout nginx -n server
Name:            nginx
Namespace:       server
Status:          ॥ Paused
Message:         CanaryPauseStep
Strategy:        Canary
  Step:          1/8
  SetWeight:     20
  ActualWeight:  20
Images:          nginx:1.19 (stable)
                 nginx:1.19-alpine (canary)
Replicas:
  Desired:       5
  Current:       5
  Updated:       1
  Ready:         5
  Available:     5

NAME                               KIND        STATUS     AGE  INFO
⟳ nginx                            Rollout     ॥ Paused   14m
├──# revision:2
│  └──⧉ nginx-86f7665cc5           ReplicaSet  ✔ Healthy  26s  canary
│     └──□ nginx-86f7665cc5-fxwqm  Pod         ✔ Running  26s  ready:1/1
└──# revision:1
   └──⧉ nginx-846f597579           ReplicaSet  ✔ Healthy  14m  stable
      ├──□ nginx-846f597579-4rg7g  Pod         ✔ Running  14m  ready:1/1
      ├──□ nginx-846f597579-gmqpc  Pod         ✔ Running  14m  ready:1/1
      ├──□ nginx-846f597579-hdtcs  Pod         ✔ Running  14m  ready:1/1
      └──□ nginx-846f597579-vmdrb  Pod         ✔ Running  14m  ready:1/1
```

## Promoting rollout

You can update all of the containers with `promote`.

```bash
kubectl argo rollouts promote nginx -n server

kubectl argo rollouts get rollout nginx -n server
Name:            nginx
Namespace:       server
Status:          ◌ Progressing
Message:         more replicas need to be updated
Strategy:        Canary
  Step:          4/8
  SetWeight:     60
  ActualWeight:  40
Images:          nginx:1.19 (stable)
                 nginx:1.19-alpine (canary)
Replicas:
  Desired:       5
  Current:       6
  Updated:       3
  Ready:         5
  Available:     5

NAME                               KIND        STATUS               AGE    INFO
⟳ nginx                            Rollout     ◌ Progressing        18m
├──# revision:4
│  └──⧉ nginx-86f7665cc5           ReplicaSet  ◌ Progressing        4m54s  canary
│     ├──□ nginx-86f7665cc5-ntm6z  Pod         ✔ Running            2m13s  ready:1/1
│     ├──□ nginx-86f7665cc5-gcw4b  Pod         ✔ Running            12s    ready:1/1
│     └──□ nginx-86f7665cc5-t7d47  Pod         ◌ ContainerCreating  0s     ready:0/1
└──# revision:3
   └──⧉ nginx-846f597579           ReplicaSet  ✔ Healthy            18m    stable
      ├──□ nginx-846f597579-4rg7g  Pod         ✔ Running            18m    ready:1/1
      ├──□ nginx-846f597579-hdtcs  Pod         ✔ Running            18m    ready:1/1
      └──□ nginx-846f597579-vmdrb  Pod         ✔ Running            18m    ready:1/1
```

## Aborting a Rollout

```bash
kubectl argo rollouts set image nginx nginx=nginx:1.19 -n server

kubectl argo rollouts get rollout nginx -n server
Name:            nginx
Namespace:       server
Status:          ◌ Progressing
Message:         more replicas need to be updated
Strategy:        Canary
  Step:          0/8
  SetWeight:     20
  ActualWeight:  0
Images:          nginx:1.19 (canary)
                 nginx:1.19-alpine (stable)
Replicas:
  Desired:       5
  Current:       6
  Updated:       1
  Ready:         5
  Available:     5

NAME                               KIND        STATUS               AGE    INFO
⟳ nginx                            Rollout     ◌ Progressing        22m
├──# revision:5
│  └──⧉ nginx-846f597579           ReplicaSet  ◌ Progressing        22m    canary
│     └──□ nginx-846f597579-ggrsm  Pod         ◌ ContainerCreating  2s     ready:0/1
└──# revision:4
   └──⧉ nginx-86f7665cc5           ReplicaSet  ✔ Healthy            8m51s  stable
      ├──□ nginx-86f7665cc5-ntm6z  Pod         ✔ Running            6m10s  ready:1/1
      ├──□ nginx-86f7665cc5-gcw4b  Pod         ✔ Running            4m9s   ready:1/1
      ├──□ nginx-86f7665cc5-t7d47  Pod         ✔ Running            3m57s  ready:1/1
      ├──□ nginx-86f7665cc5-wtwkz  Pod         ✔ Running            3m45s  ready:1/1
      └──□ nginx-86f7665cc5-j8j4z  Pod         ✔ Running            3m34s  ready:1/1

kubectl argo rollouts abort nginx -n server

kubectl argo rollouts get rollout nginx -n server
Name:            nginx
Namespace:       server
Status:          ✖ Degraded
Message:         RolloutAborted: Rollout is aborted
Strategy:        Canary
  Step:          0/8
  SetWeight:     0
  ActualWeight:  0
Images:          nginx:1.19-alpine (stable)
Replicas:
  Desired:       5
  Current:       5
  Updated:       0
  Ready:         5
  Available:     5

NAME                               KIND        STATUS        AGE    INFO
⟳ nginx                            Rollout     ✖ Degraded    23m
├──# revision:5
│  └──⧉ nginx-846f597579           ReplicaSet  • ScaledDown  23m    canary
└──# revision:4
   └──⧉ nginx-86f7665cc5           ReplicaSet  ✔ Healthy     9m53s  stable
      ├──□ nginx-86f7665cc5-ntm6z  Pod         ✔ Running     7m12s  ready:1/1
      ├──□ nginx-86f7665cc5-gcw4b  Pod         ✔ Running     5m11s  ready:1/1
      ├──□ nginx-86f7665cc5-t7d47  Pod         ✔ Running     4m59s  ready:1/1
      ├──□ nginx-86f7665cc5-wtwkz  Pod         ✔ Running     4m47s  ready:1/1
      └──□ nginx-86f7665cc5-bxtll  Pod         ✔ Running     10s    ready:1/1
```

Status of the application is `Degraded`.

```bash
kubectl argo rollouts set image nginx nginx=nginx:1.19-alpine -n server

kubectl argo rollouts get rollout nginx -n server
Name:            nginx
Namespace:       server
Status:          ✔ Healthy
Strategy:        Canary
  Step:          8/8
  SetWeight:     100
  ActualWeight:  100
Images:          nginx:1.19-alpine (stable)
Replicas:
  Desired:       5
  Current:       5
  Updated:       5
  Ready:         5
  Available:     5

NAME                               KIND        STATUS        AGE   INFO
⟳ nginx                            Rollout     ✔ Healthy     30m
├──# revision:6
│  └──⧉ nginx-86f7665cc5           ReplicaSet  ✔ Healthy     16m   stable
│     ├──□ nginx-86f7665cc5-ntm6z  Pod         ✔ Running     14m   ready:1/1
│     ├──□ nginx-86f7665cc5-gcw4b  Pod         ✔ Running     12m   ready:1/1
│     ├──□ nginx-86f7665cc5-t7d47  Pod         ✔ Running     11m   ready:1/1
│     ├──□ nginx-86f7665cc5-wtwkz  Pod         ✔ Running     11m   ready:1/1
│     └──□ nginx-86f7665cc5-bxtll  Pod         ✔ Running     7m3s  ready:1/1
└──# revision:5
   └──⧉ nginx-846f597579           ReplicaSet  • ScaledDown  30m
```
