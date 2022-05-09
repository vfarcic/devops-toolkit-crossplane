## Common

```bash
export VERSION=v0.2.7
```

## Publish To Docker Hub

```bash
kubectl crossplane build configuration \
    --name app

kubectl crossplane push configuration \
    vfarcic/crossplane-app:$VERSION
```

## Publish To Upbound

```bash
# Replace `[...]` with the Upbound Cloud account
export UP_ACCOUNT=[...]

# Replace `[...]` with the Upbound Cloud token
export UP_TOKEN=[...]

# Create `dot-kubernetes` repository

up login

up xpkg build --name app.xpkg

up xpkg push \
    --package app.xpkg \
    xpkg.upbound.io/$UP_ACCOUNT/dot-application:$VERSION
```
