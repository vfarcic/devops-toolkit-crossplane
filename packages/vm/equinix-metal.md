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

# Replace `[...]` with your Equinix Project ID token
export EQUINIX_PROJECT_ID=[...]

echo "{
  \"apiKey\":\"$EQUINIX_API_TOKEN\",
  \"projectID\":\"$EQUINIX_PROJECT_ID\"
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
  package: registry.upbound.io/equinix/provider-equinix-metal:v0.0.11
" | kubectl apply --filename -

kubectl get pkgrev

# Wait until all the packages are healthy

echo "
apiVersion: metal.equinix.com/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  projectID: $PROJECT_ID
  credentials:
    source: Secret
    secretRef:
      namespace: crossplane-system
      name: equinix-creds
      key: creds
" | kubectl apply --filename -

kubectl create namespace a-team
```

## Create a cluster

TODO: Create an XRD

TODO: Switch to an XRC

```bash
echo "
apiVersion: server.metal.equinix.com/v1alpha2
kind: Device
metadata:
  name: cp-01
spec:
  forProvider:
    hostname: cp-01
    plan: c3.small.x86
    metro: sv
    operatingSystem: ubuntu_20_04
    billingCycle: hourly
    locked: false
    networkType: hybrid
  writeConnectionSecretToRef:
    name: equinix-cp-01
    namespace: crossplane-system

---

apiVersion: server.metal.equinix.com/v1alpha2
kind: Device
metadata:
  name: worker-01
spec:
  forProvider:
    hostname: worker-01
    plan: c3.small.x86
    metro: sv
    operatingSystem: ubuntu_20_04
    billingCycle: hourly
    locked: false
    networkType: hybrid
  writeConnectionSecretToRef:
    name: equinix-worker-01
    namespace: crossplane-system
" | kubectl --namespace a-team apply --filename -

kubectl get managed

# TODO: Uncomment
# kubectl --namespace a-team get clusterclaims

# Wait until the cluster is ready

# https://github.com/k0sproject/k0sctl#installation

# Move k0sctl commands to a Kubernetes Cronjob

k0sctl init > k0sctl.yaml

# Replace `address` entries in `k0sctl.yaml`

k0sctl apply --config k0sctl.yaml

```

TODO: LB

TODO: Storage

## Use the cluster

```bash
k0sctl kubeconfig \
  | tee kubeconfig.yaml

kubectl --kubeconfig kubeconfig.yaml \
    get nodes
```

## Destroy 

```bash
kubectl delete device.server.metal.equinix.com/cp-01

kubectl delete device.server.metal.equinix.com/worker-01

kubectl get managed

# Wait until all managed AWS resources are removed
```