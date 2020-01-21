<img src="../assets/k8sland.png" align="right" width="128" height="auto"/>

<br/>

# <img src="../assets/lab.png" width="32" height="auto"/> Pod Autoscaler Lab

> Setup an HPA for your Iconoflix application

1. NOTE! You must enable|install [Metrics Server](https://github.com/kubernetes-incubator/metrics-server)
1. Using the provided template, deploy the Iconoflix application
1. Manually scale the application to 2 instances
1. Verify deployment, pods and endpoints
1. Tail the logs for the 2 instances
1. Hit the **Iconoflix** service endpoint and observe the requests logs
1. What do you notice?
1. Edit the provided k8s/hpa.yml template
  1. Autoscale to up to 5 instances once the cpu load reaches 30%
1. Deploy your HPA and make sure its reporting utilization
1. Now load up the application (See commands below)
1. Watch your hpa and pods and make sure scaling is taking place
1. Wait for the load to subsume and verify your cluster did scale back down (May take a while...)
1. Delete the application and HPA!

<br/>

---
## <img src="../assets/fox.png" width="32" height="auto"/> Commands

### Minikube Metrics server

    ```shell
    minikube addons enable metrics-server
    ```

### Go Installed? Use hey

```shell
go get -u github.com/rakyll/hey
```

### Load using Hey!

```shell
hey -c 1 -n 10000 -m POST \
  -H "Content-Type: application/json" \
  -d '{"query":"{movies{name}}"}' \
  http://$(minikube ip):30401/graphql
```

### Or.. Load using script

```shell
./burst.sh
```


---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)