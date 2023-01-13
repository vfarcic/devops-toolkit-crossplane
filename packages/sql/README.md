## Publish To Upbound

```bash
# Replace `[...]` with the Upbound Cloud account
export UP_ACCOUNT=[...]

# Replace `[...]` with the Upbound Cloud token
export UP_TOKEN=[...]

# Create `dot-sql` repository

up login

export VERSION=v0.7.14

up xpkg build --name sql.xpkg

up xpkg push \
    --package sql.xpkg \
    xpkg.upbound.io/$UP_ACCOUNT/dot-sql:$VERSION
```
