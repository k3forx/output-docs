# Configure Liveness, Readiness and Startup Probes

## What is for?

### Liveness

> The kubelet uses liveness probes to know when to restart a container. For example, liveness probes could catch a deadlock, where an application is running, but unable to make progress. Restarting a container in such a state can help to make the application more available despite bugs.

It is used for catching deadlock and restarting a pod.

### Readiness

> The kubelet uses readiness probes to know when a container is ready to start accepting traffic. A Pod is considered ready when all of its containers are ready. One use of this signal is to control which Pods are used as backends for Services. When a Pod is not ready, it is removed from Service load balancers.

It is used for checking that a pod is ready to process requests. When a pod isn't ready, it is removed from service load balancer.

### Startup Probe

> The kubelet uses startup probes to know when a container application has started. If such a probe is configured, it disables liveness and readiness checks until it succeeds, making sure those probes don't interfere with the application startup. This can be used to adopt liveness checks on slow starting containers, avoiding them getting killed by the kubelet before they are up and running.

## Demo of liveness

### Define a liveness command

**Success condition**: executed command returns `0`

Here is configuration file for a pod.

```YAML
apiVersion: v1
kind: Pod
metadata:
  labels:
    test: liveness
  name: liveness-exec
spec:
  containers:
  - name: liveness
    image: k8s.gcr.io/busybox
    args:
    - /bin/sh
    - -c
    - touch /tmp/healthy; sleep 30; rm -rf /tmp/healthy; sleep 600
    livenessProbe:
      exec:
        command:
        - cat
        - /tmp/healthy
      initialDelaySeconds: 5
      periodSeconds: 5
```

- The `periodSeconds` field specifies that the kubelet should perform a liveness probe every 5 seconds.
- The `initialDelaySeconds` field tells the kubelet that it should wait 5 seconds before performing the first probe.

The following command is used for liveness.

```bash
/bin/sh -c "touch /tmp/healthy; sleep 30; rm -rf /tmp/healthy; sleep 600"
```

After 30 seconds, the command to remove `/tmp/healthy` is executed and liveness prove will be failed after 30 seconds. The kubelet tries to restart the pod.

Let's check the behavior.

```bash
❯ kubectl apply -f demo/exec-liveness.yaml
pod/liveness-exec created

❯ kubectl describe pod liveness-exec
Name:         liveness-exec
Namespace:    default
Priority:     0
Node:         minikube/192.168.64.9
Start Time:   Wed, 19 May 2021 21:57:52 +0900
Labels:       test=liveness
Annotations:  <none>
Status:       Running
IP:           172.17.0.6
IPs:
  IP:  172.17.0.6
Containers:
  liveness:
    Container ID:  docker://9d4e8d36061cdd71ade571c11f162c9b317273439045d8da4b23c35bc5f766e3
    Image:         k8s.gcr.io/busybox
    Image ID:      docker-pullable://k8s.gcr.io/busybox@sha256:d8d3bc2c183ed2f9f10e7258f84971202325ee6011ba137112e01e30f206de67
    Port:          <none>
    Host Port:     <none>
    Args:
      /bin/sh
      -c
      touch /tmp/healthy; sleep 30; rm -rf /tmp/healthy; sleep 600
    State:          Running
      Started:      Wed, 19 May 2021 21:57:55 +0900
    Ready:          True
    Restart Count:  0
    Liveness:       exec [cat /tmp/healthy] delay=5s timeout=1s period=5s #success=1 #failure=3
    Environment:    <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-m84bx (ro)
Conditions:
  Type              Status
  Initialized       True
  Ready             True
  ContainersReady   True
  PodScheduled      True
Volumes:
  default-token-m84bx:
    Type:        Secret (a volume populated by a Secret)
    SecretName:  default-token-m84bx
    Optional:    false
QoS Class:       BestEffort
Node-Selectors:  <none>
Tolerations:     node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                 node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:
  Type    Reason     Age   From               Message
  ----    ------     ----  ----               -------
  Normal  Scheduled  14s   default-scheduler  Successfully assigned default/liveness-exec to minikube
  Normal  Pulling    13s   kubelet            Pulling image "k8s.gcr.io/busybox"
  Normal  Pulled     11s   kubelet            Successfully pulled image "k8s.gcr.io/busybox" in 2.104322434s
  Normal  Created    11s   kubelet            Created container liveness
  Normal  Started    11s   kubelet            Started container liveness
```

After 30 seconds,

```bash
❯ kubectl describe pod liveness-exec
Name:         liveness-exec
Namespace:    default
Priority:     0
Node:         minikube/192.168.64.9
Start Time:   Wed, 19 May 2021 21:57:52 +0900
Labels:       test=liveness
Annotations:  <none>
Status:       Running
IP:           172.17.0.6
IPs:
  IP:  172.17.0.6
Containers:
  liveness:
    Container ID:  docker://9d4e8d36061cdd71ade571c11f162c9b317273439045d8da4b23c35bc5f766e3
    Image:         k8s.gcr.io/busybox
    Image ID:      docker-pullable://k8s.gcr.io/busybox@sha256:d8d3bc2c183ed2f9f10e7258f84971202325ee6011ba137112e01e30f206de67
    Port:          <none>
    Host Port:     <none>
    Args:
      /bin/sh
      -c
      touch /tmp/healthy; sleep 30; rm -rf /tmp/healthy; sleep 600
    State:          Running
      Started:      Wed, 19 May 2021 21:57:55 +0900
    Ready:          True
    Restart Count:  0
    Liveness:       exec [cat /tmp/healthy] delay=5s timeout=1s period=5s #success=1 #failure=3
    Environment:    <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-m84bx (ro)
Conditions:
  Type              Status
  Initialized       True
  Ready             True
  ContainersReady   True
  PodScheduled      True
Volumes:
  default-token-m84bx:
    Type:        Secret (a volume populated by a Secret)
    SecretName:  default-token-m84bx
    Optional:    false
QoS Class:       BestEffort
Node-Selectors:  <none>
Tolerations:     node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                 node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:
  Type     Reason     Age              From               Message
  ----     ------     ----             ----               -------
  Normal   Scheduled  42s              default-scheduler  Successfully assigned default/liveness-exec to minikube
  Normal   Pulling    41s              kubelet            Pulling image "k8s.gcr.io/busybox"
  Normal   Pulled     39s              kubelet            Successfully pulled image "k8s.gcr.io/busybox" in 2.104322434s
  Normal   Created    39s              kubelet            Created container liveness
  Normal   Started    39s              kubelet            Started container liveness
  Warning  Unhealthy  3s (x2 over 8s)  kubelet            Liveness probe failed: cat: can't open '/tmp/healthy': No such file or directory
```

The output shows that `RESTARTS` has been incremented.

```bash
❯ kubectl get pod liveness-exec
NAME            READY   STATUS    RESTARTS   AGE
liveness-exec   1/1     Running   1          82s
```

### Define a liveness HTTP request

**Success condition**: Status code that is greater than `200` and less than `400`

Another kind of liveness probe uses an HTTP GET request. Here is configuration file for a pod.

```YAML
apiVersion: v1
kind: Pod
metadata:
  labels:
    test: liveness
  name: liveness-http
spec:
  containers:
  - name: liveness
    image: k8s.gcr.io/liveness
    args:
    - /server
    livenessProbe:
      httpGet:
        path: /healthz
        port: 8080
        httpHeaders:
        - name: Custom-Header
          value: Awesome
      initialDelaySeconds: 3
      periodSeconds: 3
```

To perform a probe, the kubelet sends an HTTP GET request to the server that is running in the container and listening on port 8080. If the handler for the server's `/healthz` path returns a success code, the kubelet considers the container to be alive and healthy.

Any code greater than or equal to 200 and less than 400 indicates success.

Let's see how a request for the endpoint is processed.

```golang
http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
    duration := time.Now().Sub(started)
    if duration.Seconds() > 10 {
        w.WriteHeader(500)
        w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
    } else {
        w.WriteHeader(200)
        w.Write([]byte("ok"))
    }
})
```

- First 10 seconds, the liveness successes
- After 10 secondes, the endpoint returns status code `500`

Let's check the behavior.

```bash
❯ kubectl apply -f demo/http-liveness.yaml
pod/liveness-http created

❯ kubectl describe pod liveness-http
Name:         liveness-http
Namespace:    default
Priority:     0
Node:         minikube/192.168.64.9
Start Time:   Wed, 19 May 2021 22:12:44 +0900
Labels:       test=liveness
Annotations:  <none>
Status:       Running
IP:           172.17.0.6
IPs:
  IP:  172.17.0.6
Containers:
  liveness:
    Container ID:  docker://d932041fd4b75207d00c939949b52474f0d17ea9a375d4838d570e4c4ec284b8
    Image:         k8s.gcr.io/liveness
    Image ID:      docker-pullable://k8s.gcr.io/liveness@sha256:1aef943db82cf1370d0504a51061fb082b4d351171b304ad194f6297c0bb726a
    Port:          <none>
    Host Port:     <none>
    Args:
      /server
    State:          Running
      Started:      Wed, 19 May 2021 22:12:47 +0900
    Ready:          True
    Restart Count:  0
    Liveness:       http-get http://:8080/healthz delay=3s timeout=1s period=3s #success=1 #failure=3
    Environment:    <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-m84bx (ro)
Conditions:
  Type              Status
  Initialized       True
  Ready             True
  ContainersReady   True
  PodScheduled      True
Volumes:
  default-token-m84bx:
    Type:        Secret (a volume populated by a Secret)
    SecretName:  default-token-m84bx
    Optional:    false
QoS Class:       BestEffort
Node-Selectors:  <none>
Tolerations:     node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                 node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:
  Type     Reason     Age   From               Message
  ----     ------     ----  ----               -------
  Normal   Scheduled  16s   default-scheduler  Successfully assigned default/liveness-http to minikube
  Normal   Pulling    16s   kubelet            Pulling image "k8s.gcr.io/liveness"
  Normal   Pulled     14s   kubelet            Successfully pulled image "k8s.gcr.io/liveness" in 1.878673657s
  Normal   Created    14s   kubelet            Created container liveness
  Normal   Started    14s   kubelet            Started container liveness
  Warning  Unhealthy  1s    kubelet            Liveness probe failed: HTTP probe failed with statuscode: 500

❯ kubectl get pod liveness-http
NAME            READY   STATUS    RESTARTS   AGE
liveness-http   1/1     Running   1          37s
```

### Define a TCP liveness probe

**Success condition**: a connection is established

A third type of liveness probe uses a TCP socket. With this configuration, the kubelet will attempt to open a socket to your container on the specified port. If it can establish a connection, the container is considered healthy, if it can't it is considered a failure.

Here is configuration file for a pod.

```YAML
apiVersion: v1
kind: Pod
metadata:
  name: goproxy
  labels:
    app: goproxy
spec:
  containers:
    - name: goproxy
      image: k8s.gcr.io/goproxy:0.1
      ports:
        - containerPort: 8080
      readinessProbe:
        tcpSocket:
          port: 8080
        initialDelaySeconds: 5
        periodSeconds: 10
      livenessProbe:
        tcpSocket:
          port: 8080
        initialDelaySeconds: 15
        periodSeconds: 20
```

The kubelet starts the first liveness after 15 seconds.

Let's check the behavior.

```bash
❯ kubectl apply -f demo/tcp-liveness-readiness.yaml
pod/goproxy created

❯ kubectl describe pod goproxy
Name:         goproxy
Namespace:    default
Priority:     0
Node:         minikube/192.168.64.9
Start Time:   Wed, 19 May 2021 22:29:27 +0900
Labels:       app=goproxy
Annotations:  <none>
Status:       Running
IP:           172.17.0.6
IPs:
  IP:  172.17.0.6
Containers:
  goproxy:
    Container ID:   docker://e4a479cf3977a6dc1b69229ec52b4461e612179af966f9a634c9ef91b528d5dd
    Image:          k8s.gcr.io/goproxy:0.1
    Image ID:       docker-pullable://k8s.gcr.io/goproxy@sha256:5334c7ad43048e3538775cb09aaf184f5e8acf4b0ea60e3bc8f1d93c209865a5
    Port:           8080/TCP
    Host Port:      0/TCP
    State:          Running
      Started:      Wed, 19 May 2021 22:29:30 +0900
    Ready:          True
    Restart Count:  0
    Liveness:       tcp-socket :8080 delay=15s timeout=1s period=20s #success=1 #failure=3
    Readiness:      tcp-socket :8080 delay=5s timeout=1s period=10s #success=1 #failure=3
    Environment:    <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-m84bx (ro)
Conditions:
  Type              Status
  Initialized       True
  Ready             True
  ContainersReady   True
  PodScheduled      True
Volumes:
  default-token-m84bx:
    Type:        Secret (a volume populated by a Secret)
    SecretName:  default-token-m84bx
    Optional:    false
QoS Class:       BestEffort
Node-Selectors:  <none>
Tolerations:     node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                 node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:
  Type    Reason     Age   From               Message
  ----    ------     ----  ----               -------
  Normal  Scheduled  33s   default-scheduler  Successfully assigned default/goproxy to minikube
  Normal  Pulling    32s   kubelet            Pulling image "k8s.gcr.io/goproxy:0.1"
  Normal  Pulled     30s   kubelet            Successfully pulled image "k8s.gcr.io/goproxy:0.1" in 2.252943374s
  Normal  Created    30s   kubelet            Created container goproxy
  Normal  Started    30s   kubelet            Started container goproxy

❯ kubectl get pod goproxy
NAME      READY   STATUS    RESTARTS   AGE
goproxy   1/1     Running   0          55s
```

## Readiness

Sometimes, applications are temporarily unable to serve traffic. For example, an application might

- need to load large data or configuration files during startup
- depend on external services after startup

In such cases, you don't want to kill the application, but you don't want to send it requests either. Kubernetes provides readiness probes to detect and mitigate these situations. A pod with containers reporting that they are not ready does not receive traffic through Kubernetes Services.

Readiness probes are configured similarly to liveness probes. The only difference is that you use the readinessProbe field instead of the livenessProbe field.

```bash
readinessProbe:
  exec:
    command:
    - cat
    - /tmp/healthy
  initialDelaySeconds: 5
  periodSeconds: 5
```

## Configure probes

Probes have a number of fields that you can use to more precisely control the behavior of liveness and readiness checks:

- `initialDelaySeconds`: Number of seconds after the container has started before liveness or readiness probes are initiated. Defaults to 0 seconds. Minimum value is 0.
- `periodSeconds`: How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.
- `timeoutSeconds`: Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1.
- `successThreshold`: Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup Probes. Minimum value is 1.
- `failureThreshold`: When a probe fails, Kubernetes will try failureThreshold times before giving up. Giving up in case of liveness probe means restarting the container. In case of readiness probe the Pod will be marked Unready. Defaults to 3. Minimum value is 1.
