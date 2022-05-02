## Common

```bash
export VERSION=v0.2.4
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
# Replace `[...]` with the email used to register to Upbound Cloud
export UP_EMAIL=[...]

# Replace `[...]` with the Upbound Cloud account
export UP_ACCOUNT=[...]

# Create `dot-kubernetes` repository

up login --username $UP_EMAIL

up xpkg build --name app.xpkg

up xpkg push \
    --package app.xpkg \
    xpkg.upbound.io/devops-toolkit/dot-application:$VERSION
```
