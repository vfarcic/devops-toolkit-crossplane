kubectl --namespace a-team \
    get secret a-team-aks \
    --output jsonpath="{.data.kubeconfig}" \
    | base64 -d | tee kubeconfig.yaml
