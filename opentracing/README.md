<img src="../assets/lab.png" align="right" width="auto" height="128"/>

<br/>
<br/>
<br/>

# OpenTracing G.O.T Lab

<br/>
In this lab, we are going to decorate a web server using the **OpenTracing** API to
decorate your requests. There are 2 services involved: **Castle** and **Knight**. The
Knights want to melt a Castle, but if you're a G.O.T fan, you already
know that only the NightKing can melt a Castle using his undead dragon ;)

<br/>

1. The Knight service surfaces post requests given */v1/melt* url
   1. The service then turns around and post */v1/melt* on the Castle service
      given a Knight name.
1. Instrument the Castle service by tracing incoming *melt* requests
  1. Edit your Castle trace and add the following tags to the trace:
    1. http.method
    2. http.url
    3. knight
1. If the given Knight is *NightKing* add a log to the castle span to indicate
   `the castle is melted`.
1. All other knights should produce a span error.
1. Span errors are indicated by
   1. Setting a span tag error=true
   1. Adding a structured log on the span using
      1. event=error
      2. message=only the nightking can melt
1. Edit the provided Makefile to use your own docker **registry**!
1. Build and push new docker images using the Makefile (see commands below!)
1. Deploy Jaeger, Castle and Knight service to your local cluster.
   1. You will need to modify the K8s manifest to use your image names (castle, knight)
1. Validate that your traces are correctly showing the microservices flow using
   different knights

<br/>

## Commands

1. Setup Jaeger

    ```shell
    kubectl apply -f k8s/jaeger.yml
    # Jaeger UI
    open http://$(minikube ip):30600/
    ```

2. Build your code and publish your own docker images

   > NOTE! You must change the Docker registry to use your own user account!

    ```shell
    make push
    ```

1. Deploy your services

   > NOTE! Edit the manifest and update the Docker image names

    ```shell
    kubectl deploy -f k8s/got
    ```

2. Test your endpoints and traces

   ```shell
   http $(minikube ip):30501/v1/melt knight=fred
   ```

<br/>

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)
