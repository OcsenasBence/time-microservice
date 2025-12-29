# Time Microservice & Corteza Integration

This project implements a Go-based microservice that provides the current time and integrates with the Corteza Low-Code platform via a Corredor server automation script.

## Project Structure

- **/server**: Contains the source code for the Go HTTP server and the Dockerfile.
- **/k8s**: Kubernetes configuration files (Deployment and Service).
- **/corteza**: Corteza Corredor automation scripts.

## Deployment & Installation

### 1. Deploying the Microservice (Kubernetes)

The server runs on port 8080 and exposes a Service named `time-microservice-service` within the cluster.

```bash
# Apply the Kubernetes manifests
kubectl apply -f k8s/deployment.yaml

# Verify the deployment
kubectl get pods

2. Installing the Corteza Script
Copy the corteza/scripts/time_workflow.js file to your Corredor server's script directory.

Trigger: Manual button in the TimeRecord module list view.

Function: Calls the internal API and saves the timestamp into a new record.

API Endpoints
The microservice exposes the following endpoints:

Method  Path     Description
GET     /time    Returns the current time in ISO 8601 format.
GET     /health  Kubernetes Liveness Probe endpoint. Returns 200 OK.

Tech Stack
Backend: Go (Golang) with Gin Framework

Containerization: Docker

Orchestration: Kubernetes

Automation: Corteza Corredor (JavaScript)

Usage
Open Corteza Compose.

Navigate to the TimeMicroservice namespace and the TimeRecord module.

Click the automation button (e.g., "Fetch Time from Microservice") in the toolbar.

Refresh the list to see the newly created record with the current timestamp.
