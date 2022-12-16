# Google Cloud PostgreSQL Example

## Setup

```bash
export PROJECT_ID=dot-$(date +%Y%m%d%H%M%S)

gcloud projects create $PROJECT_ID

echo "https://console.cloud.google.com/apis/library/sqladmin.googleapis.com?project=$PROJECT_ID"

# Open the URL and *ENABLE API*

export SA_NAME=devops-toolkit

export SA="${SA_NAME}@${PROJECT_ID}.iam.gserviceaccount.com"

gcloud iam service-accounts create $SA_NAME --project $PROJECT_ID

export ROLE=roles/admin

gcloud projects add-iam-policy-binding --role $ROLE $PROJECT_ID \
    --member serviceAccount:$SA

gcloud iam service-accounts keys create gcp-creds.json \
    --project $PROJECT_ID --iam-account $SA

helm repo add crossplane-stable \
    https://charts.crossplane.io/stable

helm repo update

helm upgrade --install crossplane crossplane-stable/crossplane \
    --namespace crossplane-system --create-namespace --wait

kubectl --namespace crossplane-system \
    create secret generic gcp-creds \
    --from-file creds=./gcp-creds.json

kubectl apply \
    --filename ../../crossplane-config/provider-kubernetes-incluster.yaml

kubectl apply \
    --filename ../../crossplane-config/provider-gcp-official.yaml

kubectl apply \
    --filename ../../crossplane-config/config-sql.yaml

kubectl create namespace infra

kubectl get pkgrev

# Wait until all the packages are healthy

echo "apiVersion: gcp.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  projectID: $PROJECT_ID
  credentials:
    source: Secret
    secretRef:
      namespace: crossplane-system
      name: gcp-creds
      key: creds" \
    | kubectl apply --filename -
```

## Create a PostgreSQL Instance

```bash
cat ../../examples/sql/gcp-official.yaml

kubectl --namespace infra apply \
    --filename ../../examples/sql/gcp-official.yaml

kubectl --namespace infra get sqlclaims

kubectl get managed
```

## Destroy 

```bash
kubectl --namespace infra delete \
    --filename ../../examples/sql/gcp-official.yaml

kubectl get managed

# Wait until all the resources are deleted (ignore `database`)

gcloud projects delete $PROJECT_ID
```
