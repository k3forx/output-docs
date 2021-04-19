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
> kubectl create ns argo-rollouts

> kubectl apply -n argo-rollouts -f https://raw.githubusercontent.com/argoproj/argo-rollouts/stable/manifests/install.yaml

> kubectl get all -n argo-rollouts
```

## Deploy rollout

Create `server` namespace

```bash
> kubectl create ns server
```
