<img src="../assets/k8sland.png" align="right" width="128" height="auto"/>

<br/>

# <img src="../assets/lab.png" width="32" height="auto"/> Admission Controller Lab

> Provision your cluster with a dynamic admission controller that only admits
> deployment whose pod name is fred into the cluster.

1. Update main.go to check for the correct deployment name and only admits deployments
   named fred
1. Generate certs and keys for your webhook and make sure to provide the correct
   service name in your certificate encryption
1. Generate a ca_bundle to validate api-server callback into your webhook
1. Deploy your admission controller
1. Register your validating admission controller with the api-server
1. Using the provided sample deployment, provision your cluster with the fred deployment
1. Trace the api-server logs and your admission controller logs to ensure all
   are working nominally.
1. Verify your deployment worked
1. Delete your deployment and change the pod names to blee
1. Redeploy and make sure your admission controller rejects the new deployment
1. Delete your application

---
## Commands

### Deploy custom admission controller

    ```shell
    kubectl apply -f k8s/dp.yml
    ```

### Generate your certificates

    ```shell
    ./gen.sh
    ```

### Verify your certificate against your webhook

    ```shell
    openssl s_client -connect $(minikube ip):30443/ -CAfile caCert.pem
    ```

### Base64 encode your certificate

    ```shell
    cat caCert.pem | base64 | tr -d '\n'
    ```

# Register with api-server

    ```shell
    kubectl apply -f k8s/adm.yml
    ```

# Create a new `fred deployment

    ```shell
    kubectl apply -f k8s/fred.yml
    ```


<br/>

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)