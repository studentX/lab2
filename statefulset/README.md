<img src="../assets/k8sland.png" align="right" width="128" height="auto"/>

<br/>

# <img src="../assets/lab.png" width="32" height="auto"/> StatefulSet Lab

> Deploy an **Inspector** Application using StatefulSets.

1. Update the provided StatefulSet + Service manifests use 5 replicas
1. Deploy your stateful set
1. Ensure the set is up and running correctly ie 5 instances
1. Ensure you can query the service
   1. watch http $(minikube ip):30400
1. Delete inspector-4
1. Ensure the pod is up and the host is preserved
1. Scale up the cluster to 6 replicas
1.  Ensure all pods are up and running
1.  Now scale down to one pod
1. Ensure there is only one pod running
1. Delete the last pod and ensure it comes back with the same name!
1. Delete your application!

<br/>

---
## <img src="../assets/face.png" width="32" height="auto"/> Templates

```yaml
---
apiVersion: v1
kind:       Service
metadata:
  name: inspect
  labels:
    CHANGE_ME
spec:
  # Indicate headless service
  !!CHANGE_ME!!
  selector:
    CHANGE_ME
  ports:
  - port: CHANGE_ME
    targetPor: CHANGE_ME

---
apiVersion: apps/v1
kind:       StatefulSet
metadata:
  name: inspect
spec:
  serviceName: inspect
  replicas:    CHANGE_ME
  selector:
    matchLabels:
      !!CHANGE_ME!!
  template:
    metadata:
      labels:
        !!CHANGE_ME!!
    spec:
      containers:
      - name:            inspect
        image:           quay.io/imhotepio/inspector:0.1.0
        imagePullPolicy: IfNotPresent
        ports:
        - name:          http
          containerPort: 4000
```

<br/>

---
## <img src="../assets/fox.png" width="32" height="auto"/> Commands

### Ensure set is running correctly

```shell
kubectl get sts inspect
```

### Check Endpoints

! Note your must access pods individually now

```shell
watch http $(minikube ip):30400
```

### Ensure the pod is back up and the host is preserved

```shell
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

<br/>

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)