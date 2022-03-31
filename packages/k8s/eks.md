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

# Replace `[...]` with your access key ID`
export AWS_ACCESS_KEY_ID=[...]

# Replace `[...]` with your secret access key
export AWS_SECRET_ACCESS_KEY=[...]

echo "[default]
aws_access_key_id = $AWS_ACCESS_KEY_ID
aws_secret_access_key = $AWS_SECRET_ACCESS_KEY
" >aws-creds.conf

kubectl --namespace crossplane-system \
    create secret generic aws-creds \
    --from-file creds=./aws-creds.conf

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
  name: crossplane-provider-aws
spec:
  package: crossplane/provider-aws:v0.24.1
" | kubectl apply --filename -

kubectl get pkgrev

# Wait until all the packages are healthy

echo "
apiVersion: aws.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: crossplane-system
      name: aws-creds
      key: creds
" | kubectl apply --filename -

kubectl create namespace a-team
```

## Create a cluster

```bash
echo "
apiVersion: devopstoolkitseries.com/v1alpha1
kind: ClusterClaim
metadata:
  name: a-team-eks
spec:
  id: a-team-eks
  compositionSelector:
    matchLabels:
      provider: aws
      cluster: eks
  parameters:
    nodeSize: medium
    minNodeCount: 3
  writeConnectionSecretToRef:
    name: a-team-eks
" | kubectl --namespace a-team apply --filename -

kubectl get managed

kubectl --namespace a-team get clusterclaims

# Wait until the cluster is ready
```

## Use the cluster

```bash
kubectl --namespace a-team \
    get secret a-team-eks \
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
    delete clusterclaim a-team-eks

kubectl get managed

# Wait until all managed AWS resources are removed
```