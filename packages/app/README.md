## Publish To Upbound

```bash
export VERSION=v0.3.8

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
