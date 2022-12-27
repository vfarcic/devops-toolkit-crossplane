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

kubectl apply \
    --filename ../../crossplane-config/provider-kubernetes-incluster.yaml

kubectl apply \
    --filename ../../crossplane-config/provider-aws-official.yaml

kubectl apply \
    --filename ../../crossplane-config/config-k8s.yaml

kubectl create namespace infra

kubectl get pkgrev

# Wait until all the packages are healthy

kubectl apply \
  --filename ../../crossplane-config/provider-config-aws-official.yaml
```

## Create an EKS Cluster

```bash
# TODO: Continue
kubectl --namespace infra apply \
    --filename ../../examples/k8s/aws-eks-official.yaml
    
kubectl --namespace infra get clusterclaims

kubectl get managed
```

## Destroy 

```bash
kubectl --namespace infra delete \
    --filename ../../examples/k8s/aws-eks-official.yaml

kubectl get managed

# Wait until all the resources are deleted (ignore `database`)

gcloud projects delete $PROJECT_ID
```