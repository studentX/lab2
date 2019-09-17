<img src="../assets/k8sland.png" align="right" width="128" height="auto"/>


<br/>

# <img src="../assets/lab.png" width="32" height="auto"/> G.O.T Cilium Lab

> Let's play Game of Thrones by leveraging Cilium CNI

If you follow G.O.T you might be aware that Winter has come and CastleBlack is under attack!
In this game, we will be dealing with North men and the NightKing. A North can enter
CastleBlack but not the NightKing. However no one can `melt CastleBlack like a NightKing can!

1. For this lab, you will need to configure minikube for CNI and provision Cilium
   per the commands below.
1. Using the provided K8s manifests deploy into your cluster 3 services. Namely
   castleblack, north and nightking.
1. Each service provide endpoints to enter or melt CastleBlack via /v1/enter, /v1/melt
1. Using a Cilium network policy configure your network to allow a North to enter
   CastleBlack and a NightKing to melt it (Edit k8s/policies.yml).
1. Ensure a NightKing can not just `simply *Enter* CastleBlack.
1. Likewise ensure a North can't `simply *Melt* CastleBack.
1. Delete your application!

## Setup

1. Configure minikube to enable CNI

    ```shell
    # Stop
    minikube stop
    # Delete
    minikube delete
    # Start with new configs
    minikube start --cpus=4 --memory=4096 --vm-driver=hyperkit --network-plugin=cni
    # Mount pbf volume
    minikube ssh -- sudo mount bpffs -t bpf /sys/fs/bpf
    ```

2. Install Cilium on your cluster

    ```shell
    # Install Cilium
    kubectl create -f https://raw.githubusercontent.com/cilium/cilium/1.6.1/install/kubernetes/quick-install.yaml
    ```

## Commands

1. Deploy the G.O.T services

    ```shell
    kubectl apply -f k8s/castle_black.yml -f k8s/north.yml -f k8s/nightking.yml
    ```

2. Deploy your Cilium policy

    ```shell
    kubectl apply -f k8s/policies.yml
    ```

3. Validate your policy is setup correctly

    ```shell
    kubectl exec -it -n kube-system \
    $(kubectl get po -n kube-system -l k8s-app=cilium -o jsonpath='{.items[*].metadata.name}') \
    -- cilium policy get
    ```

4. Check Cilium policy is active

    ```shell
    kubectl exec -it -n kube-system \
    $(kubectl get po -n kube-system -l k8s-app=cilium -o jsonpath='{.items[*].metadata.name}') \
    -- cilium endpoint list
    ```

5. Monitor Cilium Requests

    ```shell
    kubectl exec -it -n kube-system \
    $(kubectl get po -n kube-system -l k8s-app=cilium -o jsonpath='{.items[*].metadata.name}') \
    -- cilium monitor
    ```

6. Check NorthMan endpoints

   ```shell
   http $(minikube ip):30501/v1/enter # Success
   http $(minikube ip):30501/v1/melt  # Fails!
   ```

7. Check NightKing endpoints

   ```shell
   http $(minikube ip):30502/v1/enter # Fails!
   http $(minikube ip):30502/v1/melt  # Success!
   ```

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2019 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)