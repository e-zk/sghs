# stupid go http server

bare-bones and basically useless http server for testing stuff quickly.

## usage

serve the current directory on port 8080:

```console
$ sghs
Serving ./ on :8080
```

serve `/var/www` on `0.0.0.0:7070`:

```console
$ sghs -path . -addr "0.0.0.0:7070"
Serving ./ on 0.0.0.0:7070
```

