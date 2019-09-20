#!/bin/bash

hey -c 10 -n 10000 -m POST \
  -d '{"query":"{movies{name}}"}' \
  -H "Content-Type: application/json" \
  http://$(minikube ip):30401/graphql