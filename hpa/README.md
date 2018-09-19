# <img src="../assets/lab.png" width="32" height="auto"/> Pod Autoscaler Lab

> Scale your Iconoflix application.

1. You must enable heapster and metrics-server
    ```shell
    minikube addons enable heapster
    minikube addons enable metrics-server
    ```
1. Create an Iconoflix deployment manifest and service
    1. Use image: quay.io/imhotepio/iconoflix:mem
    1. Ensure the service is accessible locally!
1. Deploy your application
1. Manually scale the application to 2 instances
1. Verify deployment, pods and endpoints
1. Tail the logs for the 2 instances
1. Hit the **Iconoflix** service endpoint and observe the requests logs
1. What do you notice?
1. Define an HPA to autoscale from 1 to 5 instances once the cpu load reaches 30%
    1. export ICX_URL=$(minikube service iconoflix --url)
    1. Simulate load: for i in {1..100}; do wget -qO- $ICX_URL/graphql?query={movies{name}}; done
1. Check you HPA is working
1. Wait for the load to subsume and verify your cluster did scale back down (May take a while...)
1. Delete the application and HPA!


---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)