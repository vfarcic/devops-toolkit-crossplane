## Publish To Upbound

```bash
export VERSION=v0.0.41

# Replace `[...]` with the Upbound Cloud account
export UP_ACCOUNT=[...]

# Replace `[...]` with the Upbound Cloud token
export UP_TOKEN=[...]

# Create `dot-monitoring` repository

up login

up xpkg build --name monitoring.xpkg

up xpkg push \
    --package monitoring.xpkg \
    xpkg.upbound.io/$UP_ACCOUNT/dot-monitoring:$VERSION
```