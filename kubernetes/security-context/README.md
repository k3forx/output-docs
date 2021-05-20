# Configure a Security Context for a Pod or Container

- Discretionary Access Control: Permission to access an object, like a file, is based on user ID (UID) and group ID (GID).

## Set the security context for a Pod

Here is a configuration file for a Pod that has a `securityContext` and an `emptyDir` volume.

```YAML
apiVersion: v1
kind: Pod
metadata:
  name: security-context-demo
spec:
  securityContext:
    runAsUser: 1000
    runAsGroup: 3000
    fsGroup: 2000
  volumes:
    - name: sec-ctx-vol
      emptyDir: {}
  containers:
    - name: sec-ctx-demo
      image: busybox
      command: ["sh", "-c", "sleep 1h"]
      volumeMounts:
        - name: sec-ctx-vol
          mountPath: /data/demo
      securityContext:
        allowPrivilegeEscalation: false
```

- The `runAsUser` field specifies that for any Containers in the Pod, all processes run with user ID 1000
- The `runAsGroup` field specifies the primary group ID of 3000 for all processes within any containers of the Pod
  - If this field is empty, the primary group ID of the containers will be root(0)
- The `fsGroup` is for supplementary group ID

Create a pod.

```bash
❯ kubectl apply -f demo/security-context.yaml
pod/security-context-demo created

❯ kubectl get pod security-context-demo
NAME                    READY   STATUS    RESTARTS   AGE
security-context-demo   1/1     Running   0          18s

❯ kubectl exec -it security-context-demo -- sh
/ $ ps
PID   USER     TIME  COMMAND
    1 1000      0:00 sleep 1h
    9 1000      0:00 sh
   17 1000      0:00 ps
/ $ cd data
/data $ ls -l
total 4
drwxrwsrwx    2 root     2000          4096 May 20 13:07 demo
/data $ cd demo/
/data/demo $ echo hello > testfile
/data/demo $ ls -l
total 4
-rw-r--r--    1 1000     2000             6 May 20 13:08 testfile
/data/demo $ id
uid=1000 gid=3000 groups=2000
/data/demo $ exit
```

If the `runAsGroup` was omitted the gid would remain as 0(root) and the process will be able to interact with files that are owned by root(0) group and that have the required group permissions for root(0) group.

## Configure volume permission and ownership change policy for Pods

## Set the security context for a Container

## Set capabilities for a Container

## Set the Seccomp Profile for a Container

## Assign SELinux labels to a Container
