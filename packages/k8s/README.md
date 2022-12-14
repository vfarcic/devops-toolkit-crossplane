## Publish To Upbound

```bash
export VERSION=v0.5.10

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
