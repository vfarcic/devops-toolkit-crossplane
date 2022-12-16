# Azure PostgreSQL Example

## Setup

```bash
az ad sp create-for-rbac \
    --sdk-auth \
    --role Owner \
    | tee azure-creds.json

export AZURE_CLIENT_ID=$(\
    cat azure-creds.json \
    | grep clientId \
    | cut -c 16-51)

export AAD_GRAPH_API=00000003-0000-0000-c000-000000000000

az ad app permission add \
    --id "${AZURE_CLIENT_ID}" \
    --api ${AAD_GRAPH_API} \
    --api-permissions \
    e1fe6dd8-ba31-4d61-89e7-88639da4683d=Scope \
    06da0dbc-49e2-44d2-8312-53f166ab848a=Scope \
    7ab1d382-f21e-4acd-a863-ba3e13f7da61=Role

az ad app permission grant \
    --id $AZURE_CLIENT_ID \
    --api $AAD_GRAPH_API \
    --expires never

az ad app permission admin-consent \
    --id "${AZURE_CLIENT_ID}"

helm repo add crossplane-stable \
    https://charts.crossplane.io/stable

helm repo update

helm upgrade --install crossplane crossplane-stable/crossplane \
    --namespace crossplane-system --create-namespace --wait

kubectl --namespace crossplane-system \
    create secret generic azure-creds \
    --from-file creds=./azure-creds.json

kubectl apply \
    --filename ../../crossplane-config/provider-kubernetes-incluster.yaml

kubectl apply \
    --filename ../../crossplane-config/provider-azure-official.yaml

kubectl apply \
    --filename ../../crossplane-config/config-sql.yaml

kubectl create namespace infra

kubectl get pkgrev

# Wait until all the packages are healthy

kubectl apply \
    --filename ../../crossplane-config/provider-config-azure-official.yaml
```

## Create a PostgreSQL Instance

```bash
kubectl --namespace infra apply \
    --filename ../../examples/sql/azure-official.yaml

kubectl --namespace infra get sqlclaims
```

## Destroy 

```bash
kubectl --namespace infra delete \
    --filename ../../examples/sql/azure-official.yaml

kubectl get managed

# Wait until all the resources are deleted
```
