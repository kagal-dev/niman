# AGENTS.md

<!-- cspell:words coderabbit fieldalignment golangci GOPATH -->
<!-- cspell:words languagetool -->

This file provides guidance to AI agents working with code in this
repository. For what `niman` is, read [README.md](README.md) first.

## Related Documentation

- [BUILDING.md][building]: the darvaza.org shared build system, its make
  targets, tooling, linting, CI and troubleshooting, kept in
  `darvaza.org/core`.
- [TESTING.md][testing]: testing patterns for darvaza.org projects, kept
  in `darvaza.org/core`.

## Prerequisites

- Go 1.25 or later, the `go` line in `go.mod`.
- `make`, `git`, and a configured `$GOPATH`.

The `Makefile` pins `golangci-lint` and `revive` and fetches them on
demand. It runs markdownlint, cspell and ShellCheck through `pnpm dlx`,
and turns each into a no-op when pnpm or the tool is missing, so a
passing local `make tidy` does not guarantee a passing CI. LanguageTool
runs the same way, under `make check-grammar` rather than `make tidy`.

## Building

```bash
make all          # full cycle: get, generate, tidy, build
make tidy         # format, lint, validate
make test         # run tests (no cache reuse)
make coverage     # tests with coverage
make race         # tests with race detection
```

- revive holds each function to 5 arguments, 3 results, cognitive
  complexity 7 and cyclomatic complexity 10, and fails `make tidy` past
  them. Past 40 statements a function only draws a warning, and
  `make tidy` still passes; treat that warning as an error too.
  Restructure the code; the thresholds stay.
- golangci-lint checks `fieldalignment`, in `make tidy` as in CI. To
  reorder a struct, use the probe-file workflow in
  [BUILDING.md][building-fieldalignment]: running `-fix` over the tree
  strips every comment.

## Testing

- Tests follow [TESTING.md][testing] and use `darvaza.org/core`'s
  assertions.
- Test files use the external `_test` package. Those needing unexported
  access stay in the package under a `*_internal_test.go` name.

## Important Notes

- `.coderabbit.yaml` repeats the LanguageTool rules of
  `internal/build/languagetool.cfg`, so that CodeRabbit's grammar checks
  agree with `make check-grammar`: a change to either belongs in both.
- Spelling is British English. `internal/build/cspell.json` holds words
  used in two or more files; a word in one file goes in an inline
  `cspell:words` directive at its top.

## Pre-commit Checklist

1. Run `make tidy` until it passes, fixing what it reports.
2. Verify tests pass with `make test`.
3. Update this file and `README.md` when the workflow, behaviour or API
   changes.

[building]: https://github.com/darvaza-proxy/core/blob/main/BUILDING.md
[building-fieldalignment]: https://github.com/darvaza-proxy/core/blob/main/BUILDING.md#field-alignment
[testing]: https://github.com/darvaza-proxy/core/blob/main/TESTING.md
