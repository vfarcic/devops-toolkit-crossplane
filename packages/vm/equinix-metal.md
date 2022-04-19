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

# Replace `[...]` with your Equinix API token
export EQUINIX_API_TOKEN=[...]

echo "{
  \"auth_token\": \"$EQUINIX_API_TOKEN\",
  \"max_retries\": \"10\",
  \"max_retry_wait_seconds\": \"30\"
}" >equinix-creds.conf

kubectl --namespace crossplane-system \
    create secret generic equinix-creds \
    --from-file creds=./equinix-creds.conf

echo "
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: crossplane-provider-equinix-metal
spec:
  package: crossplane/provider-tf-equinix-metal:v0.2.2
" | kubectl apply --filename -

kubectl get pkgrev

echo "
apiVersion: equinixmetal.jet.crossplane.io/v1alpha1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: equinix-creds
      namespace: crossplane-system
      key: creds
" | kubectl apply --filename -

kubectl create namespace a-team
```

## Create a cluster

```bash
echo "
apiVersion: project.equinixmetal.jet.crossplane.io/v1alpha1
kind: Project
metadata:
 name: devops-toolkit
spec:
 forProvider:
   name: devops-toolkit

---

apiVersion: device.equinixmetal.jet.crossplane.io/v1alpha1
kind: Device
metadata:
  name: my-vm
spec:
  forProvider:
    projectIdRef:
      name: devops-toolkit
    metro: ny
    hostname: my-vm
    plan: c3.small.x86
    operatingSystem: ubuntu_20_04
    billingCycle: hourly
    tags:
    - crossplane
  writeConnectionSecretToRef:
    name: equinix-my-vm
    namespace: crossplane-system
" | kubectl --namespace a-team apply --filename -

kubectl get managed
```

## Destroy 

```bash
kubectl delete device.server.metal.equinix.com/my-vm

kubectl delete project.equinixmetal.jet.crossplane.io/devops-toolkit

kubectl get managed

# Wait until all managed AWS resources are removed
```