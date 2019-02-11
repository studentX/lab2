default: help

cm:       ## Provision ConfigMaps
	@kubectl create cm prom-ds -n monitoring --from-file grafana/datasource.yml
	@kubectl create cm prom-dash -n monitoring --from-file grafana/dashboard.yml
	@kubectl create cm hm-dash -n monitoring --from-file grafana/dashboard.json

up:       ## Provision Hangman on a cluster
	@kubectl apply -f k8s/prom -f k8s/hangman -f k8s/grafana.yml

down:     ## Delete all services
	@kubectl delete -f k8s/prom -f k8s/hangman -f k8s/grafana.yml

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[38;5;69m%-30s\033[38;5;168m %s\n", $$1, $$2}'
