export NAMESPACE=${1:-production}

export DB_ENDPOINT=$(kubectl \
    --namespace $NAMESPACE \
    get secret my-db \
    --output jsonpath='{.data.endpoint}' \
    | base64 --decode)

export DB_PORT=$(kubectl \
    --namespace $NAMESPACE \
    get secret my-db \
    --output jsonpath='{.data.port}' \
    | base64 --decode)

export DB_USER=$(kubectl \
    --namespace $NAMESPACE \
    get secret my-db \
    --output jsonpath='{.data.username}' \
    | base64 --decode)

export DB_PASS=$(kubectl \
    --namespace $NAMESPACE \
    get secret my-db \
    --output jsonpath='{.data.password}' \
    | base64 --decode)

export DB_URI=postgresql://$DB_USER:$DB_PASS@$DB_ENDPOINT:$DB_PORT/my-db

env | grep DB_

kubectl --namespace $NAMESPACE \
    create secret generic my-db-uri \
    --from-literal=value=$DB_URI
