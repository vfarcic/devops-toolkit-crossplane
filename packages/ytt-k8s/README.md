## AWS EKS

ytt --file _ytt_lib/resources \
    --file _ytt_lib/aws \
    --data-values-file eks-values.yaml \
    | tee ../k8s/eks.yaml

## Azure AKS

ytt --file _ytt_lib/resources \
    --file _ytt_lib/azure \
    --data-values-file aks-values.yaml \
    | tee ../k8s/aks.yaml

## Civo CK

ytt --file _ytt_lib/resources --file _ytt_lib/civo --data-values-file ck-values.yaml

## DigitalOcean DOK

ytt --file _ytt_lib/resources --file _ytt_lib/do --data-values-file dok-values.yaml

## Google Cloud GKE

ytt --file _ytt_lib/resources \
    --file _ytt_lib/google \
    --data-values-file gke-values.yaml \
    | tee ../k8s/gke.yaml
