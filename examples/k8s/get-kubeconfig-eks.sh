unset KUBECONFIG

export NAMESPACE=${1:-a-team}

export SECRET=${2:-a-team-eks}

kubectl --namespace $NAMESPACE \
    get secret $SECRET \
    --output jsonpath="{.data.kubeconfig}" \
    | base64 -d | tee kubeconfig.yaml
