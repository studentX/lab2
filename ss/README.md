# <img src="../assets/lab.png" width="32" height="auto"/> StatefulSet Lab

> Deploy an **Iconoflix** Application using StatefulSets.

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

## <img src="../assets/face.png" class="section"/> Lab Template

+ [Template](./template/tpl.yml)


<br/>

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)
