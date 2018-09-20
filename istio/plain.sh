#!/bin/bash
while true; do
  clear
  echo "Hangman"
  http --print=b --pretty=colors $(minikube ip):30500/new_game
  sleep ${1:-2}
done
