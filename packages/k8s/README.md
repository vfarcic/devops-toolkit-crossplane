## Common

```bash
export VERSION=v0.4.19
```

## Publish To Docker Hub

```bash
kubectl crossplane build configuration \
    --name k8s

kubectl crossplane push configuration \
    vfarcic/crossplane-k8s:$VERSION
```

## Publish To Upbound

```bash
# Replace `[...]` with the Upbound Cloud account
export UP_ACCOUNT=[...]

# Replace `[...]` with the Upbound Cloud token
export UP_TOKEN=[...]

# Create `dot-kubernetes` repository

up login

up xpkg build --name k8s.xpkg

up xpkg push \
    --package k8s.xpkg \
    xpkg.upbound.io/$UP_ACCOUNT/dot-kubernetes:$VERSION
```