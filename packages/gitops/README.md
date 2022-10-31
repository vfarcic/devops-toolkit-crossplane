## Publish To Upbound

```bash
export VERSION=v0.2.14

# Replace `[...]` with the Upbound Cloud account
export UP_ACCOUNT=[...]

# Replace `[...]` with the Upbound Cloud token
export UP_TOKEN=[...]

# Create `dot-gitops` repository

up login

up xpkg build --name gitops.xpkg

up xpkg push \
    --package gitops.xpkg \
    xpkg.upbound.io/$UP_ACCOUNT/dot-gitops:$VERSION
```
