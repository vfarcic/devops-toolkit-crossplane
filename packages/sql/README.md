## Publish To Upbound

```bash
export VERSION=v0.3.3

# Replace `[...]` with the Upbound Cloud account
export UP_ACCOUNT=[...]

# Replace `[...]` with the Upbound Cloud token
export UP_TOKEN=[...]

# Create `dot-sql` repository

up login

up xpkg build --name sql.xpkg

up xpkg push \
    --package sql.xpkg \
    xpkg.upbound.io/$UP_ACCOUNT/dot-sql:$VERSION
```