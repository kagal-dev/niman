# niman

<!-- cspell:words celld kagal Sumerian -->

`niman` is a command-line tool for managing clusters of Workers and
Durable Objects, whether they run on [celld][celld] or on Cloudflare.
It is aware of [kagal][kagal], and needs neither celld nor `wrangler`
on the machine it runs on.

Its name comes from NIM-AN (Sumerian 𒉏 𒀭: NIM, high or fly; AN, sky or
heaven), the system it is part of: [Fly.io][fly], celld and `niman`
running a kagal cluster as if it were on Cloudflare.

The Go module is `kagal.dev/niman`.

## Building

`niman` needs Go 1.25 or later, and uses the darvaza.org shared build
system, described in `darvaza.org/core`'s [BUILDING.md][building]:

```sh
make all     # get, generate, tidy, build
make tidy    # format, lint, validate
make test    # run the tests
```

[building]: https://github.com/darvaza-proxy/core/blob/main/BUILDING.md
[celld]: https://celld.dev/
[fly]: https://fly.io/
[kagal]: https://github.com/kagal-dev/kagal
