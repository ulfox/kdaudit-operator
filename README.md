# kdaudit-operator

This project highlithts deployments with kubernetes operator. It was bootstraped by using [sample-controller](https://github.com/kubernetes/sample-controller)

## Build the operator (Minikube example)

First access minikube container registry

    eval $(minikube docker-env)

Next build the operator

    bash setup.sh

## Run the operator

The following commands will create an admin service account that operator will use.
You should change `deployment/rbac.yaml` if you wish to restrict the service account

First create the namespace

    kubectl apply -f deployment/namespace.yaml

Deploy kdaudit-operator

    kubectl apply -f deployment/


If you view the operator's pod logs you will notice that the operator complains about missing kdaudit crd.

Apply the crd

    kubectl apply -f deployment/crd/crd.yaml

The missing kdaudit crd error should be gone now.

## Deploy a Kubernetes Controller

Up to now we deployed a custom resource definition  and an operator (controller for that crd). We can now
use our crd to create instruct the operator to do actions for us (e.g. create a new deployment).

Get the following repo 

    git clone https://github.com/ulfox/kdaudit-controller.git
    cd kdaudit-controller

Access minikube container registry

    eval $(minikube docker-env)

Build the controller

    bash setup.sh

To deploy the controller first go back to the kdaudit-operator repo and then run

    kubectl apply -f deployment/crd/kdaudit-controller.yaml

A new deployment should be created soon using the built controller. You can use any image you want however
keep in mind that the crd is built to work with the specific controller. If you wish to deploy and manage a different image,
then make the appropriate chnages in the operator under

- pkg/controllers
- controllers

