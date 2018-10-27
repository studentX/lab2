<img src="../assets/k8sland.png" align="right" width="128" height="auto"/>

<br/>

# <img src="../assets/lab.png" width="32" height="auto"/> Resources Lab

> Deploy an **Iconoflix** application with resource constraints

1. Create a deployment and service for Iconoflix
    1. Image: quay.io/imhotepio/iconoflix:mem
2. Configure resources for cpu and memory using request and limit
3. Deploy your application
4. Ensure everything is up and running!
5. Change your cpu request more than available and redeploy your application
   > REMINDER: Node allocations 4 cores / 2Gb mem
6. What do you notice?
7. Delete your application

## <img src="../assets/face.png" width="32" height="auto"/> Lab Template

```yaml
---
# Deployment
apiVersion: apps/v1
kind:       Deployment
metadata:
  name: iconoflix
spec:
  replicas: 1
  selector:
    matchLabels:
      app: iconoflix
  template:
    metadata:
      labels:
        app: iconoflix
    spec:
      containers:
      - name:  iconoflix
        image: quay.io/imhotepio/iconoflix:mem
        resources:
          !!CHANGE_ME!!
        ports:
        - containerPort: 4000
---
# Service
apiVersion: v1
kind:       Service
metadata:
  name: iconoflix
spec:
  type:  NodePort
  selector:
    app: iconoflix
  ports:
    - name:     http
      port:     4000
      nodePort: 30400
```

<br/>

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)