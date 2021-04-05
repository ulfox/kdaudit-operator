#!/usr/bin/env bash


kubectl delete  -f deployment/deployment.yaml

docker build -f Dockerfile.build . -t local/build-kdaudit-operator
docker build -f Dockerfile.kdaudit . -t local/kdaudit-operator

kubectl apply  -f deployment/deployment.yaml
