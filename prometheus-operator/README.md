# Prometheus Operator

The Prometheus Operator for Kubernetes provides easy monitoring definitions for Kubernetes services and deployment and management of Prometheus instances.

# Demo

## Start up Kubernetes cluster with minikube

```bash
minikube start --extra-config=kubelet.authentication-token-webhook=true --extra-config=kubelet.authorization-mode=Webhook
```

## Install CRDs

```bash
kubectl apply -f https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/master/bundle.yaml
```

## Deploy sample applications and their service

Deploy k8s resources as an example.

```bash
kubectl apply -f demo/deployment.yaml

kubectl apply -f demo/service
```

## Deploy service monitor

Create a service monitor to monitor the service will the label (`app`: `example-app`) and the port `web`.

The YAML file is like

```YAML
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: example-app
  labels:
    team: frontend
spec:
  selector:
    matchLabels:
      app: example-app
  endpoints:
  - port: web
```

Apply the file with `kubectl`.

```bash
kubectl apply -f demo/service-monitor.yaml
```

## Create RBAC for Prometheus operator

This is used for Prometheus.

The YAML file is

```YAML
apiVersion: v1
kind: ServiceAccount
metadata:
  name: prometheus
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: prometheus
rules:
- apiGroups: [""]
  resources:
  - nodes
  - nodes/metrics
  - services
  - endpoints
  - pods
  verbs: ["get", "list", "watch"]
- apiGroups: [""]
  resources:
  - configmaps
  verbs: ["get"]
- apiGroups:
  - networking.k8s.io
  resources:
  - ingresses
  verbs: ["get", "list", "watch"]
- nonResourceURLs: ["/metrics"]
  verbs: ["get"]
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: prometheus
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: prometheus
subjects:
- kind: ServiceAccount
  name: prometheus
  namespace: default
```

Apply the file.

```bash
kubectl apply -f demo/rbac.yaml
```

## Include the service monitor in Prometheus

As the last step, we need to include the service monitor in Prometheus. With this setting, Prometheus judges which service monitor they manage.
To manager the service monitor `example-app`, we need to specify its label (i.e. `team`: `frontend`).

The YAML file is

```YAML
apiVersion: monitoring.coreos.com/v1
kind: Prometheus
metadata:
  name: prometheus
  labels:
    prometheus: prometheus
spec:
  serviceAccountName: prometheus
  serviceMonitorSelector:
    matchLabels:
      team: frontend
  resources:
    requests:
      memory: 400Mi
  enableAdminAPI: false
```

Apply the YAML.

```bash
kubectl apply -f demo/service-monitors.yaml
```

## Check the status of the applications on UI

Let's expose Prometheus with service and see the status of the applications.

```bash
kubectl apply -f demo/prometheus-service.yaml

minikube service prometheus --url
```

Visit the URL you got the above command.

![image](https://user-images.githubusercontent.com/45956169/115475706-2f1c6800-a27b-11eb-939c-15266c95bdd0.png)
