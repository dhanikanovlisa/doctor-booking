# Implementation of a Doctor Booking Service Using Istio Service Mesh
This project implements a Doctor Booking Service within a microservices architecture managed by the Istio service mesh.
 It allows users to interact with the service to view available booking slots and book appointments, with traffic routing and observability provided by Istio.

## Requirements
To run the Doctor Booking Service, ensure you have the following installed:

1. Kubernetes: A cluster to deploy the services.
2. Istio: Service mesh for managing traffic and observability.
3. kubectl: CLI tool to manage the Kubernetes cluster.
4. Kiali: Visualization tool for Istio traffic flows.
5. Docker: To build container images for the services.

## How To Run
1. Apply the Kubernetes manifests:
```
kubectl apply -f platform/kube/
kubectl apply -f istio-manifest/
```
2. Access it from external by Using
```
minikube tunnel
```
or
```
kubectl port-forward -n istio-system svc/istio-ingressgateway  8080:80
```