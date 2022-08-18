kubectl --namespace a-team \
    get secret a-team-$CLUSTER_TYPE \
    --output jsonpath="{.data.kubeconfig}" \
    | base64 -d | tee kubeconfig.yaml
