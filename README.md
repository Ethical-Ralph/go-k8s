# Golang k8s

A Golang project that can be deployed on Kubernetes using Minikube. this project includes three services and an API gateway implemented with Ingress. This readme file provides instructions on how to set up and run the project on Minikube, including configuring the `/etc/hosts` file to access the API gateway.

## Prerequisites

Before running the project, make sure you have the following dependencies installed:

- Go programming language
- Docker
- Minikube

## Getting Started

To set up and run the project on Minikube, follow the steps below:

1. Start Minikube by running the following command:

   ```
   minikube start
   ```

2. Set up the Docker environment in the Minikube session using the following command:

   ```
   eval $(minikube docker-env)
   ```

3. Build the Docker containers for each service by navigating to their respective directories and running the following commands:

   - For `auth-service`:
     ```
     docker build -t ethical/auth-api ./auth
     ```
   - For `hello-service`:
     ```
     docker build -t ethical/hello-api ./hello
     ```
   - For `dummy-service`:
     ```
     docker build -t ethical/dummy-api ./dummy
     ```

4. Apply the Kubernetes manifest to create the services, deployments, and ingress by running the following command:

   ```
   kubectl apply -f ./manifest
   ```

5. Configure the `/etc/hosts` file to map the hostname used by the API gateway to the Minikube IP. Open the `/etc/hosts` file using a text editor (you may need administrator/root privileges) and add the following line:

   ```
   <minikube-ip> go-gateway.app
   ```

   Replace `<minikube-ip>` with the IP address of your Minikube cluster. You can get the IP address by running the command `minikube ip`.

6. Access the API gateway and services by opening a web browser and navigating to `http://go-gateway.app`. You should be able to access the following paths:
   - `/login` - Auth service
   - `/hello` - Hello service
   - `/dummy` - Dummy service

- If you are unable to access the services in the browser, you can use the following command to make a curl request from within the Minikube environment:

  ```
  minikube ssh -- curl http://go-gateway.app
  ```

  This command will SSH into the Minikube cluster and then make a curl request to the API gateway.

## Cleaning Up

To clean up the resources created by the project, run the following command:

```
kubectl delete -f ./manifest
```

This will delete all services, deployments, and ingress created for this project.

## Additional Information

- For more information on Go programming language, refer to the official [Go documentation](https://golang.org/doc/).
- For more information on Docker, refer to the official [Docker documentation](https://docs.docker.com/).
- For more information on Minikube, refer to the official [Minikube documentation](https://minikube.sigs.k8s.io/docs/).

Feel free to reach out if you have any questions or encounter any issues.
