# AWS EKS

ytt --file _ytt_lib/resources --file _ytt_lib/aws --data-values-file eks-values.yaml

# Azure AKS

ytt --file _ytt_lib/resources --file _ytt_lib/azure --data-values-file aks-values.yaml