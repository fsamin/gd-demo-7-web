# gd-demo-7-web — step 7: web exposures

The step-1 application, published through git-deploy's web exposures: a
hostname of your choice, then a sign-in in front of it, with no code change in
the application.

Behind an authenticated exposure, the page shows who is signed in. It reads the
headers the platform's authenticating proxy sets:

| Header | Shown as |
|---|---|
| `X-Forwarded-Preferred-Username` | the username of the signed-in account |
| `X-Forwarded-Email` | the identity provider used (the platform puts it there) |
| `X-Forwarded-Groups` | the account's groups |

Behind the proxy the application is reachable through it only, so these headers
come from it. On a public exposure anyone can send them: the page only displays
them, it never decides anything from them.

## Demo

```sh
git-deploy -n demo init --name portal
git-deploy -n demo expose set --hostname portal.<base-domain> web portal
git-deploy -n demo expose set --hostname portal.<base-domain> --proxy --idp <provider> web portal
```

## Locally

```sh
go run .                                                   # http://localhost:8080
curl -H 'X-Forwarded-Preferred-Username: alice' -H 'X-Forwarded-Email: partner' localhost:8080
```
