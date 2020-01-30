#!/bin/bash
for i in {1..300}; do \
  curl -H "Content-Type: application/json" -d '{"query":"{movies{name}}"}' http://$(minikube ip):30400/graphql; \
done
