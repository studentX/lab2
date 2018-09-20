# <img src="../assets/lab.png" width="32" height="auto"/> Istio Lab

> Deploy a hangman game app in an Istio Cluster

1. Define a manifest for a hangman v1 service in a manifest call hangman_v1.yml
   1. Image: k8sland/hangman_svc_go:0.0.1
2. Deploy the hangman v1 service
3. Define a manifest for a hangman v2 service
   1. Same image as above
   2. Change the command to read /app/hangam --file trump


## <img src="../assets/lab.png" width="32" height="auto"/> Commands

1. Download and install Istio
2. Deploy Hangman V1

   ```shell
   kubectl apply -f k8s/hangman_v1.yml
   ```
3. Play the game!

  ```shell
  kubectl run -i --tty --rm hm --image k8sland/hangman-cli-go:0.0.1 --command -- /app/hangman_cli --url hangman:5000
  ```
4. Define Istio Resources

  ```shell
  ku apply -f istio/gateway.yml -f istio1/routes.yml -f istio1/subsets.yml
<br/>

ku label namespace default istio-injection=enabled

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)
