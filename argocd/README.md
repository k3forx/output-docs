# Argo CD

## How does it work?

[Overview](https://argo-cd.readthedocs.io/en/stable/)

> Argo CD is implemented as a Kubernetes controller which continuously monitors running applications and compares the current, live state against the desired target state (as specified in the Git repo). A deployed application whose live state deviates from the target state is considered `OutOfSync`. Argo CD reports & visualizes the differences, while providing facilities to automatically or manually sync the live state back to the desired target state. Any modifications made to the desired target state in the Git repo can be automatically applied and reflected in the specified target environments.

Summary

- Argo CD is used for deployment for Kubernetes applications
- Argo CD always monitors actual running applications on Kubernetes
- Argo CD always checks that the status of running applications on Kubernetes is the same as YAML files on GitHub repository

# Argo CD for administrators

## Components

### API server

### Repository server

### Application controller

# Argo CD for developers

## Concepts

### Projects

Projects provide a logical grouping of applications, which is useful when Argo CD is used by multiple teams. Projects provide the following features:

- restrict _what_ may be deployed (trusted Git source repositories)

- restrict _where_ apps may be deployed to (destination clusters and namespaces)

- restrict what kinds of objects may or may not be deployed (e.g. RBAC, CRDs, DaemonSets, NetworkPolicy etc...)

- defining project roles to provide application RBAC (bound to OIDC groups and/or JWT tokens)

#### Example of YAML file for project

Here is an example of YAML file for `database` project.

```YAML
apiVersion: argoproj.io/v1alpha1
kind: AppProject
metadata:
  name: database
  namespace: argocd
spec:
  description: database

  sourceRepos:
  - '*'

  destinations:
  - namespace: database
    server: https://kubernetes.default.svc

  clusterResourceWhitelist:
  - group: '*'
    kind: '*'

  namespaceResourceBlacklist:
  - group: ''
    kind: ResourceQuota
  - group: ''
    kind: LimitRange
  - group: ''
    kind: NetworkPolicy

  orphanedResources:
    warn: false
```

The explanations of the above YAML file are described in the following.

| Section                      | Explanation                                                                                |
| ---------------------------- | ------------------------------------------------------------------------------------------ |
| `sourceRepos`                | The repositories that applications within the project can pull manifests from              |
| `destinations`               | Kubernetes cluster and namespace where applications within the project will be deployed on |
| `clusterResourceWhitelist`   | Which cluster-scoped resources are allowed to be created                                   |
| `namespaceResourceBlacklist` | Which namespace-scoped resources are NOT allowed to be created                             |
| `orphanedResources`          | Enable Argo CD to monitor orphaned resources                                               |

References:

- [Projects](https://argo-cd.readthedocs.io/en/stable/user-guide/projects/)
- [Orphaned Resources Monitoring](https://argo-cd.readthedocs.io/en/stable/user-guide/orphaned-resources/)

### Application

The Application CRD is the Kubernetes resource object representing a deployed application instance in an environment. It is defined by two key pieces of information:

- `source` reference to the desired state in Git (repository, revision, path, environment)

- `destination` reference to the target cluster and namespace. For the cluster one of server or name can be used, but not both (which will result in an error). Behind the hood when the server is missing, it is being calculated based on the name and then the server is used for any operations.

#### Example of YAML file for application

Here is an example of YAML file for application `database-mysql` which belongs to `database` project.

```YAML
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: database-mysql
  namespace: argocd
  finalizers:
    - resources-finalizer.argocd.argoproj.io
spec:
  project: database

  source:
    repoURL: https://github.com/k3forx/fastapi-example
    targetRevision: master
    path: k8s/mysql/overlays/database/

  destination:
    server: https://kubernetes.default.svc
    namespace: database

  syncPolicy:
    automated:
      prune: false
      selfHeal: false
```

The explanations of the above YAML file are described in the following.

| Section       | Explanation                                                                                                                                                                                                                                                                                 |
| ------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `project`     | which project the applications belongs to                                                                                                                                                                                                                                                   |
| `source`      | which GitHub repo the application monitors for                                                                                                                                                                                                                                              |
| `destination` | which Kubernetes the application should be synced with                                                                                                                                                                                                                                      |
| `syncPolicy`  | policy for sync. `prune` decides whether the resource that are not managed in GitHub is deleted or not when Argo CD synced the latest codes. `selfHeal` decides whether Argo CD tries to back resources to desired state of GitHub repository when some changes are added in the resources. |

# Demo

In the following sections, you can create a project `server` and a application `nginx` in the project. The project is deployed in `server` namespace and Argo CD is deployed in `argocd` namespace.

## Prerequisite

You can use minikube as Kubernetes cluster where Argo CD works.

```bash
> minikube version --short
minikube version: v1.19.0
```

## Set up for Kubernetes

Before you create a project, make sure that `server` and `argocd` namespaces are created.

```bash
> kubectl create ns argocd

> kubectl create ns server
```

## Deploy Argo CD

```bash
> kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

> kubectl get all -n argocd
```

## Login Argo CD UI

You can login UI for Argo CD.

```bash
> kubectl port-forward svc/argocd-server -n argocd 8080:443
```

Then, you can see the UI with [localhost:8080/login](localhost:8080/login).

![image](https://user-images.githubusercontent.com/45956169/115147360-9bd60d80-a095-11eb-8c93-0d819536ca41.png)

Let's login the UI. Argo CD provide login password for `admin` user. You can get the password by the following command (you need to open a new terminal window).

```bash
> kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
```

You can login with username `admin` and the password you got.

![image](https://user-images.githubusercontent.com/45956169/115147805-caed7e80-a097-11eb-9e99-830cad0635af.png)

## Create a project

It's time to create `server` project.

The YAML file for the project. Applications within this project will be deployed on `server` namespace.

```YAML
apiVersion: argoproj.io/v1alpha1
kind: AppProject
metadata:
  name: server
  namespace: argocd
spec:
  description: Project for server

  sourceRepos:
  - '*'

  destinations:
  - namespace: server
    server: https://kubernetes.default.svc

  clusterResourceWhitelist:
  - group: '*'
    kind: '*'

  namespaceResourceBlacklist:
  - group: ''
    kind: ResourceQuota
  - group: ''
    kind: LimitRange
  - group: ''
    kind: NetworkPolicy
```

Create the project with `kubectl` command.

```bash
> kubectl apply -f demo/project.yaml
```

On the UI, you can see that the project is successfully created.

![image](https://user-images.githubusercontent.com/45956169/115147963-62eb6800-a098-11eb-87cc-7c1bef113acd.png)

## Create a application

As a next step, you can create `nginx` application within the project.

The YAML file for the application.

```YAML
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: server-nginx
  namespace: argocd
spec:
  project: server

  source:
    repoURL: https://github.com/k3forx/output-docs
    targetRevision: master
    path: argocd/demo/nginx

  destination:
    server: https://kubernetes.default.svc
    namespace: server
```

Create the application with `kubectl` command.
