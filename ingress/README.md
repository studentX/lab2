# <img src="../assets/lab.png" width="32" height="auto"/> Ingress Controller Lab

> The dev team has been hard at work and came up with a new set of dictionaries
> for the company flagship product *Hangman*. The hangman game is composed of
> 2 services namely hangman and dictionary. The dictionary service serves a
> collection of words. The hangman service calls the dictionary to get a list
> of words, picks a random word and seeds the guessing game with that word.

In this lab we're going to setup and ingress to allow us to play hangman from
2 separated microservice stacks and leverage a *Traefik* ingress controller to
mutiplex across our 2 hangman services.

1. Provision your cluster with a Traefik ingress via a daemonset
1. Deploy hangman v1 and v2 using the provided manifest
1. Verify you can access the Traefik UI
2. Create an ingress resource that allows to switch between the 2 instances of
   the hangman service
3. Leveraging named hosting, setup the ingress to direct traffic to either the
   trump version (v1) or the halloween version (v2)
4. Play the game using hangman v1
5. Play the game now using hangman v2 and make sure both instanced are serving
   the correct dictionaries!

<br/>

---
## Commands

### Deploy Traefik Ingress + RBAC

    ```shell
    kubectl apply -f k8s/traefik
    ```

### Launch Traefik UI

    ```shell
    open http://traefik-ui.minikube/dashboard/
    ```

### Deploy Hangman v1, v2

    ```shell
    kubectl apply -f k8/hangman1 -f k8s/hangman2
    ```

### Fake a DNS

    ```shell
    echo "$(minikube ip) trump.minikube halloween.minikube" | sudo tee -a /etc/hosts
    ```

### Play!

```shell
bin/hangman -hm trump.minikube
bin/hangman -hm halloween.minikube
```


<br/>

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)