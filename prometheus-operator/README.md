# Prometheus Operator

The Prometheus Operator for Kubernetes provides easy monitoring definitions for Kubernetes services and deployment and management of Prometheus instances.

# Demo

In this demonstration, you can see how Prometheus operator works in the real situation.
You deploy an application with deployment and expose it with service of `ClusterIP` type in `test` namespace. Prometheus operator is deployed in `monitoring` namespace and Prometheus monitors the application through service monitor.

## Start up Kubernetes cluster with minikube

```bash
minikube start --extra-config=kubelet.authentication-token-webhook=true --extra-config=kubelet.authorization-mode=Webhook
```

## Install CRDs

```bash
kubectl apply -k .
```

## Deploy sample applications and their service

Deploy k8s resources as an example.

```bash
kubectl create ns test

kubectl apply -f demo/deployment.yaml -n test

kubectl apply -f demo/service -n test

kubectl get all -n test
```

## Create `monitoring` namespace for Prometheus operator

```bash
kubectl apply -f monitoring-namespace.yaml
```

## Deploy service monitor

Create a service monitor to monitor the service will the label (`app`: `example-app`).
The section `endpoints` in `spec` section defines a scrapeable endpoint serving Prometheus metrics. This endpoints should be the same as endpoints of service.

You can refer to the details in the following documentation.

- https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#endpoint

The YAML file is like

```YAML
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: example-app
  labels:
    team: frontend
spec:
  namespaceSelector:
    matchNames:
      - test
  selector:
    matchLabels:
      app: example-app
  endpoints:
    - targetPort: 8080
```

Apply the file with `kubectl`.

```bash
kubectl apply -f demo/service-monitor.yaml -n monitoring
```

## Create RBAC for Prometheus operator

This is used for Prometheus.

The YAML file is

```YAML
apiVersion: v1
kind: ServiceAccount
metadata:
  name: prometheus
  namespace: monitoring
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
    namespace: monitoring
```

Apply the file.

```bash
kubectl apply -f demo/rbac.yaml -n monitoring
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
kubectl apply -f demo/prometheus.yaml -n monitoring
```

## Check the status of the applications on UI

Let's expose Prometheus with service and see the status of the applications.

```bash
kubectl apply -f demo/prometheus-service.yaml -n monitoring

minikube service -n monitoring prometheus --url
```

Visit the URL you got the above command.

![image](https://user-images.githubusercontent.com/45956169/115475706-2f1c6800-a27b-11eb-939c-15266c95bdd0.png)
