# <img src="../assets/lab.png" width="32" height="auto"/> StatefulSet Lab

> Deploy an **Inspector** Application using StatefulSets.

1. Create a statefulset and service manifest for Inspector
    1. Use image: quay.io/imhotepio/inspector:0.1.0
    2. NOTE: The inspector runs on port 4000
2. Define your StatefulSet to use 5 replicas
3. Make sure the service is external accessible via port 30400
4. Ensure the set is up and running correctly ie 5 instances
5. Ensure you can query the service
   1. watch http $(minikube ip):30400
6. Take note of one of the pod IP address
7. Delete that pod in the set
8. Ensure the pod is up and the ip/host is preserved
9. Scale up the cluster by adding one more replica
10. Ensure all pods are up and running and take note of their IPs
11. Now scale down to one pod
12. Ensure there is only one pod running
13. Delete the last pod and ensure it comes back with the same IP!
14. Delete your application!

<br/>

---
## Templates

+ [Template](./tpl.yml)

<br/>

---
## Commands

### Create the set

```shell
kubectl apply -f ss.yml
```

### Ensure set is running correctly

```shell
kubectl get sts inspect
```

### Check Endpoints

! Note your must access pods individually now

```shell
watch http $(minikube ip):30400
```

### Take note of one of the pods IP address

```shell
kubectl get ep -l app=inspect
```

### Delete on of the pods with that IP address

```shell
kubectl delete po inspect-1
```

### Ensure the pod is back up and the ip/host is preserved

```shell
kubectl get ep -l app=inspect
kubectl exec -it inspect-1 -- printenv | grep HOSTNAME
```

### Scale up the cluster by adding one more replica

```shell
kubectl scale sts inspect --replicas=6
```

### Ensure all pods are up and running

```shell
kubectl get pod,ep
```

### Now scale down to one pod

```shell
kubectl scale sts inspect --replicas=1
```

### Delete the last pod and make sure it comes back up with the right IP!

```shell
kubectl delete po inspect-0
kubectl get ep
```

### Delete the statefulset. Completely!

```shell
kubectl delete -f ss.yml
```

<br/>

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)