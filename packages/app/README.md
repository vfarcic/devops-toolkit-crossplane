## Publish To Upbound

```bash
# Replace `[...]` with the Upbound Cloud account
export UP_ACCOUNT=[...]

# Replace `[...]` with the Upbound Cloud token
export UP_TOKEN=[...]

# Create `dot-application` repository

up login

export VERSION=v0.4.5

up xpkg build --name app.xpkg

up xpkg push --package app.xpkg \
    xpkg.upbound.io/$UP_ACCOUNT/dot-application:$VERSION
```
