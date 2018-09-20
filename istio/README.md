# <img src="../assets/lab.png" width="32" height="auto"/> Istio Lab

> Let's play Hangman! Deploy a Hangman service to an Istio Cluster

1. Download and provision your minikube cluster using the lesson instructions
1. Define a manifest for a hangman v1 service in a manifest call hangman_v1.yml
   1. Image: k8sland/hangman_svc_go:0.0.1
2. Deploy the hangman v1 service
3. Define a manifest for a hangman v2 service
   1. Same image as above
   2. Change the command to read /app/hangam --file trump


## <img src="../assets/sol.png" width="32" height="auto"/> Solution

1. Download Istio

    ```shell
    mkdir ~/istio && cd istio
    curl -L https://git.io/getLatestIstio | sh -
    cd istio-1.0.2
    # NOTE istio comes bundle with it's own cli aka istioctl.
    export PATH=$PWD/bin:$PATH
    ```

1. Provision Istio in your minikube cluster

    ```shell
    cd ~/istio/istio-1.0.2/install/kubernetes
    kubectl apply -f istio-demo.yaml
    ```

1. Deploy Hangman V1

    ```shell
    kubectl apply -f k8s/hangman_v1.yml
    ```

1. Play the game!

    ```shell
    kubectl run -i --tty --rm hm --image k8sland/hangman-cli-go:0.0.1 --command -- /app/hangman_cli --url hangman:5000
    ```

1. Configure your edge controller and routes

    ```shell
    kubectl apply -f istio/gateway.yml -f istio1/routes.yml -f istio1/subsets.yml
    ```

1. Enable Istio Sidecar injection in your default namespace

    ```shell
    ku label namespace default istio-injection=enabled
    ```

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)
