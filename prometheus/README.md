# <img src="../assets/lab.png" width="32" height="auto"/> Prometheus Hangman Lab

We're going to play a hangman game. The game consist of a couple of
services hangman and dictionary and a CLI to submit guesses. The code
is already implemented and deployment manifest are in the k8s directory.
To play the game, we leverage Prometheus metrics to track good/bad guesses
as well as the number of game won. We will also display the metrics in a Grafana
dashboard. Sounds cool?

1. Instrument the hangman code base and add 2 prometheus counters to track your
   good and bad guesses.
1. Next define a prometheus gauge to track your game results ie +1 for wins and -1
   for loss.
1. Before you get to play the game, your will need to tell Prometheus to
   track your hangman service by setting the ServiceMonitor CRD.
2. Using the provided deployment templates, deploy Prometheus using the awesome
   CoreOS operator, Grafana and the hangman services namely dictionary and hangman.
1. Launch the Grafana UI and setup your prometheus datasource **Prom** to reference
   http://prometheus.default.svc.cluster.local:9090
2. In the Grafana UI, load the custom dashboard from the grafana/dashboard.json file
3. You can now enjoy the fruits of your labor by firing off the hangman CLI and
   try out your guessing skills while watching your performance in Grafana...


## Commands





curl -X DELETE -u "$user:$pass" https://index.docker.io/v1/repositories/$namespace/$reponame/

ku run -i --tty --rm hm --image k8sland/go-hangman-cli:0.0.1 --command -- /app/hangman_cli --hm hangman:5000

> Leverage Init Containers to provision dictionaries for a Dictionary service.

The dictionary service will load dictionary data from a given asset directory and
dictionary name mounted on a volume. Use an init-container to provision the
volume with a set of dictionaries by cloning a dictionary assets repo.

1. Define a pod using the following Docker image: k8sland/dictionary-svc-go:0.0.2
2. The dictionary service is launched using the following command:
   ```shell
   /app/dictionary -a dictionary_dir -d dictionary_name
   ```
3. This service runs on port 4000 and exposes /words endpoint to list out the words
   contained in the dictionary loaded via *-a/d* options above.
4. Define an init container to providion a volume to be used by the dictionary
   service
5. Your init container will need to clone this repo [Dictionaries](https://github.com/k8sland/dictionaries.git) in order to provision the volume
6. Change the init container command to cause the pod to fail
7. What's happening with your dictionary pod?

## Commands

1. Launch your pod
    ```shell
    kubectl apply -f dictionary.yml
    ```
1. Verify the init container is successful and pod is launched
    ```shell
    kubectl get po
    ```
1. Verify the volume was provisioned correctly
    ```shell
    kubectl exec -it dictionary -- wget -q -O - http://localhost:4000/words
    ```
1. Change git url so that it does not resolve
    ```shell
    kubectl delete -f dictionary.yml --force --grace-period=0
    kubectl apply -f dictionary.yml
    ```

<br/>

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)