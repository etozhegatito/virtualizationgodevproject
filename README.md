# virtualizationgodevproject

Welcome to the Virtualization Go Dev Project! 🚀 This repository is tailored for developers (virt) who are keen on virtualization technologies.

## Getting Started

To get this project up and running smoothly, follow these steps:

1. First things first, let's build everything by running the command:

docker-compose build


2. Once everything is built, fire up all the services by running:


docker-compose up


3. You can access the services via the following URLs:
- [http://localhost:8080/](http://localhost:8080/)
- [http://localhost:8081/](http://localhost:8081/)
- [http://localhost:8082/](http://localhost:8082/)

## Kubernetes Exploration

Let's take our journey further into the realms of Kubernetes.

1. Start your Minikube cluster using the following command:

minikube start


2. Once Minikube is up and running, let's check out some essentials:
- To view all running pods:
  ```
  kubectl get pods
  ```
- To list all nodes in the cluster:
  ```
  kubectl get nodes
  ```
- To observe the Horizontal Pod Autoscalers:
  ```
  kubectl get hpa
  ```

Feel free to explore, experiment, and contribute to this project! 🎉
