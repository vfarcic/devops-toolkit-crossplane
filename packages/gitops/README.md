```bash
kubectl crossplane build configuration \
    --name gitops

kubectl crossplane push configuration \
    vfarcic/crossplane-gitops:v0.2.8
```
