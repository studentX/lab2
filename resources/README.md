<img src="../assets/k8sland.png" align="right" width="128" height="auto"/>

<br/>

# <img src="../assets/lab.png" width="32" height="auto"/> Resources Lab

> Experiment with resources constraints for an Iconoflix application

1. Edit the resource section in the provided manifest (k8s/iconoflix.yml) and
   initially swag the Iconoflix container quotas
1. Deploy the application
1. Ensure everything is up and running!
1. Pressure the Iconoflix API using the provided hey command and monitor
   the node and pod resources
1. Tune your resources based on your findings
1. Change your cpu request to more than available and redeploy your application
   > REMINDER: Node allocations 4 cores / 8Gb mem
1. What do you notice?
1. Delete your application

<br/>

---

## <img src="../assets/fox.png" width="32" height="auto"/> Commands

### Simulate load

    Install [Hey](https://github.com/rakyll/hey)

    ```shell
    hey -c 1 -n 10000 http://$(minikube ip):30400/graphql?query={movies{name}}
    # Or...
    ./bursh.sh
    ```

<br/>

---

<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)