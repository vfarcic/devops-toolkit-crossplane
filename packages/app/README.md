```bash
kubectl crossplane build configuration \
    --name app

kubectl crossplane push configuration \
    vfarcic/crossplane-app:v0.1.3
```