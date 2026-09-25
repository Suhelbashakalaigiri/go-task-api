# Go Task API — CI/CD & Kubernetes

A simple REST API built with Go and integrated with a Jenkins CI/CD pipeline, Docker, Docker Hub, and Kubernetes.

This project demonstrates an end-to-end DevOps workflow:

```text
GitHub
   │
   ▼
Jenkins
   │
   ├── Test
   ├── Vet
   ├── Build
   │
   ▼
Docker Build
   │
   ▼
Docker Hub
   │
   ▼
Kubernetes
   │
   ├── Deployment
   └── Service
```

## Project Overview

The application is a Go-based REST API for managing tasks.

The primary objective of this project is to learn and demonstrate:

* Go REST API development
* Git and GitHub
* Unit testing
* Docker containerization
* Jenkins CI/CD
* Docker Hub image publishing
* Kubernetes Deployment
* Kubernetes Service
* Automated application deployment

## Technology Stack

| Technology                | Purpose                      |
| ------------------------- | ---------------------------- |
| Go                        | Backend REST API             |
| Git                       | Version control              |
| GitHub                    | Source code repository       |
| Jenkins                   | CI/CD automation             |
| Docker                    | Application containerization |
| Docker Hub                | Docker image registry        |
| Kubernetes                | Container orchestration      |
| Docker Desktop Kubernetes | Local Kubernetes environment |

## Project Structure

```text
go-task-api/
│
├── main.go
├── main_test.go
├── go.mod
├── Dockerfile
├── Jenkinsfile
├── deployment.yaml
├── service.yaml
└── README.md
```

### Important Files

**`main.go`**

Contains the Go REST API implementation.

**`main_test.go`**

Contains unit tests for the application.

**`Dockerfile`**

Defines how the Go application is packaged into a Docker image.

**`Jenkinsfile`**

Defines the CI/CD pipeline used by Jenkins.

**`deployment.yaml`**

Defines the Kubernetes Deployment for running the application.

**`service.yaml`**

Defines the Kubernetes Service used to expose the application.

---

# CI/CD Pipeline

The Jenkins pipeline automates the application build and deployment process.

## Pipeline Stages

```text
Checkout
   ↓
Test
   ↓
Vet
   ↓
Build
   ↓
Docker Build
   ↓
Docker Push
   ↓
Kubernetes Deploy
   ↓
Rollout Verification
```

### 1. Checkout

Jenkins checks out the source code from the GitHub repository.

### 2. Test

Runs the Go unit tests:

```bash
go test -v ./...
```

### 3. Vet

Runs Go's static analysis tool:

```bash
go vet ./...
```

### 4. Build

Builds the Go application.

### 5. Docker Build

Jenkins creates a Docker image from the `Dockerfile`.

The image is tagged using the Jenkins build number.

### 6. Docker Push

The generated Docker image is pushed to Docker Hub.

### 7. Kubernetes Deployment

Jenkins updates the Kubernetes Deployment with the newly created Docker image.

### 8. Rollout Verification

Jenkins verifies that the Kubernetes Deployment successfully rolls out:

```bash
kubectl rollout status deployment/go-task-api
```

---

# Docker

Build the Docker image manually:

```bash
docker build -t go-task-api .
```

Run the container:

```bash
docker run -p 8080:8080 go-task-api
```

The application will then be available on:

```text
http://localhost:8080
```

---

# Kubernetes

The application is deployed to a local Kubernetes cluster.

## Deployment

The Kubernetes Deployment manages the application Pod.

```bash
kubectl apply -f deployment.yaml
```

Check the Deployment:

```bash
kubectl get deployments
```

Check Pods:

```bash
kubectl get pods
```

## Service

Apply the Kubernetes Service:

```bash
kubectl apply -f service.yaml
```

Check the Service:

```bash
kubectl get services
```

The current Service uses `NodePort` to expose the application.

---

# Jenkins + Kubernetes Workflow

The Jenkins pipeline connects the complete workflow:

```text
Developer
    │
    │ Git Push
    ▼
GitHub
    │
    │ Checkout
    ▼
Jenkins
    │
    ├── Go Test
    ├── Go Vet
    ├── Go Build
    │
    ▼
Docker Image
    │
    ▼
Docker Hub
    │
    ▼
Kubernetes
    │
    ├── Deployment
    │       │
    │       ▼
    │      Pod
    │
    └── Service
```

This allows a code change to move through the build, containerization, image publishing, and Kubernetes deployment stages automatically.

---

# Running Locally

## Prerequisites

Install the following tools:

* Go
* Git
* Docker
* Kubernetes
* kubectl
* Jenkins

For local Kubernetes deployment, Docker Desktop Kubernetes can be used.

## Run the Go Application

Clone the repository:

```bash
git clone https://github.com/Suhelbashakalaigiri/go-task-api.git
```

Enter the project directory:

```bash
cd go-task-api
```

Run the application:

```bash
go run main.go
```

---

# Kubernetes Deployment

Make sure your Kubernetes cluster is running.

Apply the Deployment:

```bash
kubectl apply -f deployment.yaml
```

Apply the Service:

```bash
kubectl apply -f service.yaml
```

Verify:

```bash
kubectl get pods
kubectl get deployments
kubectl get services
```

Check the rollout:

```bash
kubectl rollout status deployment/go-task-api
```

---

# Learning Objectives

This project was created as a hands-on learning project to understand how a backend application moves through a DevOps pipeline.

The main concepts covered are:

* REST API development with Go
* Git workflow
* GitHub repository management
* CI/CD concepts
* Jenkins pipelines
* Automated testing
* Docker image creation
* Docker image publishing
* Kubernetes Deployments
* Kubernetes Services
* `kubectl` commands
* Jenkins-to-Kubernetes deployment

---

# Current Deployment Environment

The current Kubernetes deployment uses a local Kubernetes environment.

The project is being extended to learn cloud deployment using **AWS and Amazon EKS**.

Planned workflow:

```text
GitHub
   ↓
Jenkins
   ↓
Docker
   ↓
Docker Hub
   ↓
Amazon EKS
   ↓
Go Task API
```

---

# Future Improvements

The following areas can be explored as the project evolves:

* AWS EKS deployment
* Kubernetes ConfigMaps
* Kubernetes Secrets
* Resource requests and limits
* Liveness and readiness probes
* Multiple application replicas
* Rolling update strategies
* Kubernetes Ingress
* Horizontal Pod Autoscaler
* Container/image security scanning
* Application monitoring and observability
* CI/CD security improvements

These features will be added progressively as part of the learning process.

---

# Author

**Suhel Basha Kalaigiri**

GitHub:
https://github.com/Suhelbashakalaigiri

Repository:
https://github.com/Suhelbashakalaigiri/go-task-api
