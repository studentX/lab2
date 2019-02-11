kubectl create cm prom-ds -n monitoring --from-file grafana/datasource.yml
kubectl create cm prom-dash -n monitoring --from-file grafana/dashboard.yml
kubectl create cm hm-dash -n monitoring --from-file grafana/dashboard.json
