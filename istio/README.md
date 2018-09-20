# <img src="../assets/lab.png" width="32" height="auto"/> Istio Lab

> Let's play Hangman! Deploy a Hangman service to an Istio Cluster

> The hangman game is composed of 2 K8s services: dictionary and hangman. The hangman
> service fetches a list of words from the dictionary service and pick a random word to
> initialize the guessing game.

1. Download and provision your minikube cluster using the lesson instructions.
1. Ensure all the Istio components are up and running in the *istio-system* namespace.
1. Deploy your Istio gateway, routes and subsets manifests
1. Enable Sidecar injection in your default namespace:

    ```shell
    kubectl label ns default istio-injection=enabled
    ```

1. Define a manifest for a dictionaryV1 deployment in a manifest call dictionary_v1.yml
   1. Image: k8sland/dictionary_svc_go:0.0.2
   1. Change the command to read /app/dictionary -d words.txt
   1. Deploy your dictionaryV1 manifest
   1. Ensure your dictionary deployment is up, running and side-cared! ie Ready=[2/2]
2. Define a dictionary service
   1. The service must watch for pods with label app=dictionary
   2. Make sure the dictionary service is accessible on your node using port 30400
   3. Deploy your dictionary service
   4. Ensure this service found your dictionary pod!
3. Define a manifest for the hangman service in a manifest called hangman.yml
   1. Image: k8sland/hangman_svc_go:0.0.2
   2. Change the command to read /app/hangman --url dictionary:4000
   3. Define a K8s service for hangman to be exposed on nodeport: 30500
   4. Deploy the hangman service
   5. Verify the hangman pod and service are correctly configured!
4. Play the game!

    ```shell
    # NOTE! Press enter once the pod is initialized!
    kubectl run -i --tty --rm hm --image k8sland/hangman-cli-go:0.0.1 --command -- /app/hangman_cli --url hangman:5000
    ```

5. Define a manifest for a dictionaryV2 deployment in a manifest call
   1. Same image as above
   2. Change the command to read /app/dictionary -d trump.txt
   3. Deploy your dictionaryV2
   4. Ensure the dictionary is up and running
6. Using the picker.sh script check the current hangman behavior

    ```shell
    ./picker.sh
    ```

7. Create a new Istio policy to route traffic 80% to v2 and 20% to v1 of the dictionary
   1. Define an istio virtualservice policy to apply weighted routing when traffic is origination from hangman
   2. Provision your new policy
   3. Ensure the virtualservice was created correctly
   4. Check the picker and make sure it produced more v2 words
8. Delete your weighted traffic policy!
   1. Ensure the picker just shows regular dictionary words
9. Next create a new virtual service policy to mirror all traffic coming to v1 to v2
   1. Create a new policy and setup your mirror
   2. Deploy your new policy
   3. Tail both v1 and v2 logs and make sure all traffic destined to v1 also hits the v2 version.


<br/>
<br/>

---
## <img src="../assets/sol.png" width="32" height="auto"/> Solution

1. Download Istio

    ```shell
    mkdir ~/istio && cd istio
    curl -L https://git.io/getLatestIstio | sh -
    cd istio-1.0.2
    # NOTE istio comes bundle with it's own cli aka istioctl.
    export PATH=$PWD/bin:$PATH
    ```

2. Provision Istio in your minikube cluster

    ```shell
    cd ~/istio/istio-1.0.2/install/kubernetes
    kubectl apply -f istio-demo.yaml
    ```

1. Deploy your Istio gateway and routes

    ```shell
    ku apply -f istio/gateway.yml -f istio/routes.yml -f istio/subsets.yml
    ```

1. Deploy DictionaryV1

    ```shell
    kubectl apply -f k8s/dictionary_v1.yml
    kubectl get deploy,rs,po
    ```

1. Deploy the dictionary service

    ```shell
    kubectl apply -f k8s/dictionary.yml
    kubectl get svc,ep
    ```

3. Deploy Hangman V1

    ```shell
    kubectl apply -f k8s/hangman_v1.yml
    kubectl get deploy,rs,po,svc
    ```

4. Play the game!

    ```shell
    kubectl run -i --tty --rm hm --image k8sland/hangman-cli-go:0.0.1 --command -- /app/hangman_cli --url hangman:5000
    ```

5. Deploy dictionary V2

    ```shell
    kubectl apply -f k8s/dictionary_v2.yml
    kubectl get deploy,rs,po
    ```

6. Run the picker

    ```shell
    ./picker.sh
    ```

1. Deploy dictionary weighted traffic policy

    ```shell
    kubectl apply -f istio/dictionary-80-20.yml
    kubectl get virtualservice
    ```

1. Delete your traffic policy

    ```shell
    kubectl delete -f istio/dictionary-80-20.yml
    ```

2. Deploy dictionary mirror traffic policy

    ```shell
    kubectl apply -f istio/mirror.yml
    kubectl get virtualservice
    ```

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)
