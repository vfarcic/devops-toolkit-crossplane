## Common

```bash
export VERSION=v0.4.18
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
# Replace `[...]` with the email used to register to Upbound Cloud
export UP_EMAIL=[...]

# Replace `[...]` with the Upbound Cloud account
export UP_ACCOUNT=[...]

# Create `dot-kubernetes` repository

up login --username $UP_EMAIL

up xpkg build --name k8s.xpkg

up xpkg push \
    --package k8s.xpkg \
    xpkg.upbound.io/devops-toolkit/dot-kubernetes:$VERSION
```