<img src="../assets/k8sland.png" align="right" width="auto" height="128"/>

<br/>
<br/>

# <img src="../assets/lab.png" width="32" height="auto"/> Prometheus Hangman Lab

> We're going to play a hangman game 🎉🎊🥳 The game consists of a couple of
> services namely hangman and dictionary and a CLI to submit guesses. The hangman
> service queries the dictionary service to get a list of words for the guess
> word. To play the game, we are going to leverage Prometheus metrics to
> track good/bad guess counts as well as tracking win rates. Next, we are going
> to display the tally metrics in a Grafana dashboard. Sounds cool?

NOTE: Skip the first 3 steps if no GO chops!

1. Instrument the hangman code base and add 2 prometheus counters to track your
   good and bad guesses (see game.go).
2. Next define a prometheus gauge to track your game results:
   ie +1 for wins and -1 for loss (see tally.go)
3. Build your new game images and push to DockerHub using the provided makefile.
   The target will build and push 2 images for the hangman CLI and the service.
   You will need to modify the Makefile and change the docker registry account to
   you own!
4. Before you get to play the game, your will need to tell Prometheus to
   track your hangman service by setting the ServiceMonitor CRD (k8s/prom/crd.yml)
5. Using the provided deployment templates (k8s directory), deploy Prometheus using the awesome
   CoreOS operator, Grafana and the hangman services namely dictionary and hangman.
6. Launch the Grafana UI.
7. You can now enjoy the fruits of your hard labor by firing off the hangman CLI and
   try out your guessing skills while watching your game performance in Grafana...
8. Delete all resources when done!

<br/>

---

## Commands

### [IF GO CODER!] Build and Push your Docker images

```shell
  cd hangman
  # Be sure to edit the Makefile REGISTRY to your own!
  make push
```

### Deploy Prometheus Operator

```shell
kubectl apply -f k8s/prom
```

### Deploy Grafana

```shell
# Create ConfigMaps for datasource and dashboard
make cm
# NOTE! Grafana default credentials: admin/admin
kubectl apply -f k8s/grafana.yml
```

### Open Grafana

  ```shell
  # Creds admin/admin
  minikube service -n monitoring grafana
  ```

### Play!

```shell
kubectl run -i --tty --rm hm \
  --image k8sland/go-hangman-prom-cli:0.0.2 \
  --generator=run-pod/v1 \
  --command -- /app/hangman_cli --hm hangman:5000
```

<br/>

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)
