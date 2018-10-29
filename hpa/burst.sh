#!/bin/bash
while true; do
  clear
  http --print=b --pretty=colors $(minikube ip):30400/graphql query='{movies{name}}'
  sleep ${1:-1}
done

curl -H "Content-Type: application/json" -d '{"query":"{movies{name}}"}' http://$(minikube ip):30400/graphql
hey -c 1 -n 1 -m POST -H "Content-Type: application/json" -d '{"query":"{movies{name}}"}' http://$(minikube ip):30400/graphql