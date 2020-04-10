#! /bin/bash

kubectl apply -f ../k8-deployment.yml
kubectl get deployments
kubectl get pods