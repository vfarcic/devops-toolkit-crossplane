# Azure PostgreSQL Example

## Setup

```bash
export SUBSCRIPTION_ID=$(az account show --query id -o tsv)

az ad sp create-for-rbac \
    --sdk-auth \
    --role Owner \
    --scopes /subscriptions/$SUBSCRIPTION_ID \
    | tee azure-creds.json

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

kubectl apply --filename ../../crossplane-config/config-sql.yaml

kubectl create namespace infra

kubectl get pkgrev

# Wait until all the packages are healthy

kubectl apply \
    --filename ../../crossplane-config/provider-config-azure-official.yaml
```

## Create a PostgreSQL Instance

```bash
cat ../../examples/sql/azure-official.yaml

export NAME_RAND=my-db-$(date +%Y%m%d%H%M%S)

cat ../../examples/sql/azure-official.yaml \
    | sed -e "s@my-db@$NAME_RAND@g" \
    | kubectl --namespace infra apply --filename -

kubectl --namespace infra get sqlclaims

kubectl get managed
```

## Destroy 

```bash
cat ../../examples/sql/azure-official.yaml \
    | sed -e "s@my-db@$NAME_RAND@g" \
    | kubectl --namespace infra delete --filename -

kubectl get managed

# Wait until all the resources are deleted (ignore `database`)
```
