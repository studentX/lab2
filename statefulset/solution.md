---
layout: lab
---

# <img src="/assets/sol.png" class="section"/> StatefulSet Lab Solution

<br/>

---
## Manifest

+ [StatefulSet](./ss.yml)

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