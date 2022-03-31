# AWS EKS Example

## Setup

```bash
# Create a management Kubernetes cluster manually (e.g., minikube, Rancher Desktop, eksctl, etc.)

helm repo add crossplane-stable \
    https://charts.crossplane.io/stable

helm repo update

helm upgrade --install \
    crossplane crossplane-stable/crossplane \
    --namespace crossplane-system \
    --create-namespace \
    --wait

# Replace `[...]` with your Civo token
export CIVO_TOKEN=[...]

export CIVO_TOKEN_ENCODED=$(\
    echo $CIVO_TOKEN | base64)

echo "apiVersion: v1
kind: Secret
metadata:
  namespace: crossplane-system
  name: civo-creds
type: Opaque
data:
  credentials: $CIVO_TOKEN_ENCODED" \
    | kubectl apply --filename -

echo "
apiVersion: pkg.crossplane.io/v1
kind: Configuration
metadata:
  name: crossplane-k8s
spec:
  package: vfarcic/crossplane-k8s:v0.4.8

---

apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: crossplane-provider-civo
spec:
  package: crossplane/provider-civo:main
" | kubectl apply --filename -

kubectl get pkgrev

# Wait until all the packages are healthy

echo "
apiVersion: civo.crossplane.io/v1alpha1
kind: ProviderConfig
metadata:
  name: crossplane-provider-civo
spec:
  region: nyc1
  credentials:
    source: Secret
    secretRef:
      namespace: crossplane-system
      name: civo-creds
      key: credentials
" | kubectl apply --filename -

kubectl create namespace a-team
```

## Create a cluster

```bash
echo "
apiVersion: devopstoolkitseries.com/v1alpha1
kind: ClusterClaim
metadata:
  name: a-team-ck
spec:
  id: a-team-ck
  compositionSelector:
    matchLabels:
      provider: civo
      cluster: ck
  parameters:
    nodeSize: medium
    minNodeCount: 3
" | kubectl --namespace a-team apply --filename -

kubectl get managed

kubectl --namespace a-team get clusterclaims

# Wait until the cluster is ready
```

## Use the cluster

```bash
kubectl --namespace crossplane-system \
    get secret cluster-civo-a-team-ck \
    --output jsonpath="{.data.kubeconfig}" \
    | base64 -d \
    | tee kubeconfig.yaml

# The credentials in `kubeconfig.yaml` are temporary for security reasons

kubectl --kubeconfig kubeconfig.yaml \
    get nodes
```

## Destroy 

```bash
kubectl --namespace a-team \
    delete clusterclaim a-team-ck

kubectl get managed

# Wait until all managed AWS resources are removed
```