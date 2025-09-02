# Go Health App

This repository contains a simple Go application built with [Gin](https://github.com/gin-gonic/gin). The app exposes a single `/health` endpoint that returns a JSON response `{ "status": "ok" }`. The project is containerized with Docker and deployable to Kubernetes with Ingress routing.

---

## 1. Start the Local Environment

### Prerequisites

* [Go 1.24+](https://golang.org/doc/install)
* [Docker Desktop](https://www.docker.com/products/docker-desktop) with Kubernetes enabled, or [Minikube](https://minikube.sigs.k8s.io/docs/)
* [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
* (Optional) [Helm](https://helm.sh/) for ingress-nginx installation

### Run Locally (without Kubernetes)

```bash
go run ./cmd/server
```

Test the endpoint:

```bash
curl http://localhost:8080/health
# {"status":"ok"}
```

### Build and Run with Docker

```bash
docker build -t go-health:latest -f deployments/Dockerfile .
docker run -p 8080:8080 go-health:latest
```

Test again:

```bash
curl http://localhost:8080/health
```

---

## 2. Access via Local Kubernetes Ingress

### Step 1. Build Docker Image for Local Cluster

If using Docker Desktop Kubernetes:

```bash
docker build -t go-health:latest -f deployments/Dockerfile .
```

If using Minikube:

```bash
eval $(minikube docker-env)
docker build -t go-health:latest -f deployments/Dockerfile .
```

### Step 2. Apply Manifests

```bash
kubectl apply -f deployments/k8s/deployment.yaml
kubectl apply -f deployments/k8s/service.yaml
kubectl apply -f deployments/k8s/ingress.yaml
```

### Step 3. Verify Resources

```bash
kubectl get pods
kubectl get svc
kubectl get ingress
```

### Step 4. Test the App via Ingress

```bash
curl http://localhost/health
# {"status":"ok"}
```

---

## 3. Mapping a Custom Domain in Production

In production, you’ll typically expose your service using a domain name instead of `localhost`.

### Steps:

1. **Buy/Register a domain** (e.g., from Namecheap, Google Domains, Route53).
2. **Point DNS to your ingress controller’s LoadBalancer IP**:

   * Create an `A` record for your domain (e.g., `api.example.com`) → Ingress controller public IP.
3. **Update Ingress manifest** with the host:

   ```yaml
   spec:
     rules:
     - host: api.example.com
       http:
         paths:
         - path: /health
           pathType: Prefix
           backend:
             service:
               name: go-health-service
               port:
                 number: 80
   ```
4. Re-apply the manifest:

   ```bash
   kubectl apply -f deployments/k8s/ingress.yaml
   ```

Now you can access the service:

```bash
curl http://api.example.com/health
```

---

## 4. Configure SSL/TLS with Certificates

For production, HTTPS should be enabled. The standard tool is [cert-manager](https://cert-manager.io/).

### Step 1. Install cert-manager

```bash
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.15.0/cert-manager.yaml
```

### Step 2. Create a ClusterIssuer (Let’s Encrypt)

```yaml
apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: letsencrypt-prod
spec:
  acme:
    server: https://acme-v02.api.letsencrypt.org/directory
    email: you@example.com
    privateKeySecretRef:
      name: letsencrypt-prod
    solvers:
    - http01:
        ingress:
          class: nginx
```

Apply it:

```bash
kubectl apply -f cluster-issuer.yaml
```

### Step 3. Update Ingress with TLS

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: go-health-ingress
  annotations:
    cert-manager.io/cluster-issuer: "letsencrypt-prod"
spec:
  ingressClassName: nginx
  tls:
  - hosts:
    - api.example.com
    secretName: go-health-tls
  rules:
  - host: api.example.com
    http:
      paths:
      - path: /health
        pathType: Prefix
        backend:
          service:
            name: go-health-service
            port:
              number: 80
```

Re-apply:

```bash
kubectl apply -f deployments/k8s/ingress.yaml
```

### Step 4. Verify HTTPS

```bash
curl https://api.example.com/health --resolve api.example.com:<IP>:443
```

Expected:

```json
{"status":"ok"}
```

---

## Summary

* Local run: `go run ./cmd/server` or Docker.
* Kubernetes: apply manifests under `deployments/k8s`.
* Ingress: accessible at `http://localhost/health` locally.
* Production: map DNS → ingress IP, update Ingress `host`, enable TLS with cert-manager.
