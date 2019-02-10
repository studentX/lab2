<img src="../assets/k8sland.png" align="right" width="128" height="auto"/>

<br/>

# <img src="../assets/lab.png" width="32" height="auto"/> Admission Controller Lab

> Trick or Treat? Provision your cluster with a dynamic admission controller
> that rejects all Grim Reaper's deployments!

1. Update the code in main.go to reject all grim reaper deployments. ie
   check for deployment resource with a label app=grim reaper.
1. Generate certs and keys for your webhook and make sure to provide the correct
   service name in your certificate encryption
1. Generate a ca_bundle to validate api-server callback into your webhook.
1. Set the ca_bundle in k8s/adm.yml to the output of the previous command
1. Deploy your admission controller
1. Register your validating admission controller with the api-server
1. Update the provided grim deployment with the magic label and provision your cluster
1. Trace the api-server logs and your admission controller logs to ensure all
   are working nominally.
1. Verify your deployment was denied for the right reason
1. Delete your deployment and change/remove the label
1. Redeploy and make sure your admission controller allows the new deployment
1. Delete your application!

---
## Commands

1. Generate your certificates

    ```shell
    ./gen.sh
    ```

1. Build your admission controller Docker image

    ```shell
    make img
    ```

1. Deploy custom admission controller

    ```shell
    kubectl apply -f k8s/dp.yml
    ```

1. Base64 encode your certificate

    ```shell
    # Base64 encode your key
    cat caCert.pem | base64 | tr -d '\n' | pbcopy
    # IMPORTANT! Edit k8s/adm.yml and paste in caBundle
    ```

1. Register with api-server

    ```shell
    kubectl apply -f k8s/adm.yml
    # Verify!
    kubectl get validatingwebhookconfiguration
    ```

1. Verify your certificate against your webhook

    ```shell
    openssl s_client -host $(minikube ip) -port 30443 -CAfile caCert.pem
    ```

1. Create a new `Grim Reaper deployment

    ```shell
    kubectl apply -f k8s/grim.yml
    ```

1. Verify your deployments did not make it!

1. Change the deployment label and redeploy


<br/>

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)