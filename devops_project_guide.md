# DevOps Platform Engineering Practice Project: Cloud-Native E-Commerce Microservice

## Project Theme
Build a **Product Catalog API** for an e-commerce platform - a Go-based REST API that manages product inventory with full cloud-native CI/CD pipeline using Argo Workflows.

## Project Stages

### Stage 1: Go REST API Development
**Objective**: Build a production-ready Go REST API with proper middleware, validation, and observability

**Steps to Complete**:
1. Initialize Go module with proper project structure
2. Implement CRUD operations for products (with categories, pricing, inventory)
3. Add JWT authentication middleware
4. Implement request validation with proper error handling
5. Add health checks and metrics endpoints (Prometheus format)
6. Implement structured logging with correlation IDs
7. Add graceful shutdown handling
8. Create comprehensive unit tests with table-driven tests

**Things to Watch Out For**:
- Go's interface-based architecture - don't over-engineer with too many abstractions
- Proper error handling patterns (avoid panic in production code)
- Context propagation for cancellation and timeouts
- Memory management with large payloads
- HTTP middleware ordering matters

### Stage 2: Local Kubernetes Environment Setup
**Objective**: Configure Docker Desktop K8s cluster that mirrors production environments

**Steps to Complete**:
1. Enable Kubernetes in Docker Desktop
2. Install NGINX Ingress Controller for Docker Desktop
3. Install Argo Workflows with proper RBAC
4. Set up Prometheus and Grafana for observability
5. Configure local container registry (registry:2 container)
6. Create necessary namespaces and base configurations

**Things to Watch Out For**:
- Docker Desktop K8s uses different networking than cloud providers
- Resource limits are important in constrained local environment
- LoadBalancer services work differently than cloud providers
- File system permissions for volume mounts

### Stage 3: Argo Workflows Pipeline Design
**Objective**: Create enterprise-grade workflows for testing, building, and deployment

**Steps to Complete**:
1. Design Test Workflow (unit tests, linting, security scanning)
2. Design Build Workflow (multi-stage Docker builds, image scanning)
3. Design Deploy Workflow (GitOps-style deployment with rollback capabilities)
4. Design Promotion Workflow (environment promotion patterns)
5. Implement proper artifact management between workflow steps
6. Add workflow templates for reusability
7. Configure workflow event triggers

**Things to Watch Out For**:
- Argo uses different YAML syntax than GitLab CI - watch the indentation
- Artifact passing between steps requires proper volume configuration
- Resource requests/limits are critical for workflow pods
- Workflow RBAC can be tricky - start permissive, then lock down
- Error handling in workflows is different from traditional CI/CD

### Stage 4: Container Registry Integration
**Objective**: Implement secure image management with proper tagging strategy

**Steps to Complete**:
1. Set up GitHub Container Registry authentication
2. Implement semantic versioning with Git tags
3. Configure multi-stage Docker builds for optimization
4. Add image vulnerability scanning with Trivy
5. Implement image signing with cosign
6. Create proper image tagging strategy (latest, semver, commit SHA)
7. Configure image cleanup policies

**Things to Watch Out For**:
- GitHub Container Registry has different auth patterns than DockerHub
- Image signing requires careful key management
- Multi-arch builds can be complex in local environment
- Registry rate limits can affect workflows
- Image layer caching strategies for build performance

### Stage 5: Kubernetes Deployment Patterns
**Objective**: Deploy using modern K8s patterns with production-ready configurations

**Steps to Complete**:
1. Create Helm chart with proper templating
2. Configure Horizontal Pod Autoscaler
3. Add Pod Disruption Budgets
4. Implement Network Policies
5. Configure resource quotas and limits
6. Add readiness and liveness probes
7. Configure service discovery and ingress routing
8. Implement configuration management (ConfigMaps/Secrets)

**Things to Watch Out For**:
- HPA requires metrics server to be running
- Network policies can break service communication if misconfigured
- Ingress routing rules depend on your ingress controller type
- Resource requests vs limits - understand the difference
- Probe configuration affects rolling updates

### Stage 6: Observability and Monitoring
**Objective**: Implement comprehensive observability stack

**Steps to Complete**:
1. Add custom Prometheus metrics to your Go application
2. Configure Grafana dashboards for your API
3. Set up distributed tracing with Jaeger
4. Implement structured logging with proper log levels
5. Configure Alertmanager rules for critical metrics
6. Add log aggregation and parsing
7. Create SLI/SLO definitions

**Things to Watch Out For**:
- Prometheus metrics can impact performance if not designed properly
- Grafana dashboard queries need optimization for large datasets
- Jaeger requires proper context propagation in your application
- Log volume can grow quickly - implement proper rotation
- Alert fatigue - start with critical alerts only

---

## Installation Instructions

### Prerequisites Setup

#### 1. Go Development Environment
```bash
# Install Go 1.21+ from https://golang.org/dl/
# Verify installation
go version

# Install essential Go tools
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install github.com/swaggo/swag/cmd/swag@latest
```

#### 2. Docker Desktop Kubernetes Setup
```bash
# Enable Kubernetes in Docker Desktop Settings
# Verify cluster is running
kubectl cluster-info

# Install NGINX Ingress Controller for Docker Desktop
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/cloud/deploy.yaml

# Wait for ingress controller to be ready
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s
```

#### 3. Argo Workflows Installation
```bash
# Install Argo Workflows
kubectl create namespace argo
kubectl apply -n argo -f https://github.com/argoproj/argo-workflows/releases/download/v3.4.4/install.yaml

# Patch server auth mode for local development
kubectl patch deployment \
  argo-server \
  --namespace argo \
  --type='json' \
  -p='[{"op": "replace", "path": "/spec/template/spec/containers/0/args", "value": [
  "server",
  "--auth-mode=server"
]}]'

# Install Argo CLI
# Download from https://github.com/argoproj/argo-workflows/releases
# Add to PATH
```

#### 4. Local Container Registry
```bash
# Run local registry for development
docker run -d -p 5000:5000 --name registry registry:2

# Configure Docker Desktop to use insecure registry
# Add "localhost:5000" to Docker Desktop Settings > Docker Engine > insecure-registries
```

#### 5. Observability Stack
```bash
# Install Prometheus Operator
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update

# Install with values suitable for local development
helm install prometheus prometheus-community/kube-prometheus-stack \
  --set prometheus.prometheusSpec.storageSpec.volumeClaimTemplate.spec.resources.requests.storage=1Gi \
  --set alertmanager.alertmanagerSpec.storage.volumeClaimTemplate.spec.resources.requests.storage=1Gi

# Install Jaeger
kubectl create namespace observability
kubectl apply -n observability -f https://github.com/jaegertracing/jaeger-operator/releases/download/v1.41.0/jaeger-operator.yaml

# Create Jaeger instance
kubectl apply -f - <<EOF
apiVersion: jaegertracing.io/v1
kind: Jaeger
metadata:
  name: jaeger
  namespace: observability
spec:
  strategy: allInOne
  allInOne:
    image: jaegertracing/all-in-one:latest
    options:
      log-level: debug
  storage:
    type: memory
    options:
      memory:
        max-traces: 10000
  ui:
    options:
      dependencies:
        menuEnabled: false
EOF
```

### Authentication Setup

#### GitHub Container Registry
```bash
# Create GitHub Personal Access Token with packages:write scope
# Login to GHCR
echo $GITHUB_TOKEN | docker login ghcr.io -u $GITHUB_USERNAME --password-stdin

# Create Kubernetes secret for image pulling
kubectl create secret docker-registry ghcr-secret \
  --docker-server=ghcr.io \
  --docker-username=$GITHUB_USERNAME \
  --docker-password=$GITHUB_TOKEN \
  --docker-email=$GITHUB_EMAIL
```

#### Image Signing Setup
```bash
# Install cosign
go install github.com/sigstore/cosign/cmd/cosign@latest

# Generate signing key pair
cosign generate-key-pair

# Store private key as Kubernetes secret
kubectl create secret generic cosign-key --from-file=cosign.key
```
## Common Pitfalls & Solutions

### Docker Desktop Kubernetes Limitations
- **Issue**: LoadBalancer services don't get external IPs
- **Solution**: Use kubectl port-forward or ingress controllers

### Argo Workflows Common Issues
- **Issue**: Workflows stuck in pending state
- **Solution**: Check resource requests and node capacity

### Container Registry Authentication
- **Issue**: ImagePullBackOff errors
- **Solution**: Verify secret creation and service account binding

### Observability Stack Resource Usage
- **Issue**: Prometheus consuming too much memory
- **Solution**: Configure retention policies and storage limits

## Success Metrics

By project completion, you should be able to:
1. **Demonstrate** a complete commit-to-production pipeline
2. **Explain** the trade-offs between different deployment strategies
3. **Implement** proper observability and alerting
4. **Troubleshoot** common Kubernetes and Argo issues
5. **Design** workflows that handle failures gracefully

## Next Steps for Production Readiness

1. **Multi-environment setup** (dev/staging/prod)
2. **GitOps patterns** with ArgoCD
3. **Security scanning** integration
4. **Performance testing** in pipelines
5. **Disaster recovery** procedures