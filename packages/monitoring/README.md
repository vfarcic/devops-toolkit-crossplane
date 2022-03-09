## Publish

```bash
kubectl crossplane build configuration \
    --name monitoring

kubectl crossplane push configuration \
    vfarcic/crossplane-monitoring:v0.0.37
```