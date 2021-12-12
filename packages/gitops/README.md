```bash
kubectl crossplane build configuration \
    --name gitops

kubectl crossplane push configuration \
    vfarcic/crossplane-gitops:v0.1.4
```
