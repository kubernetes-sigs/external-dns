# AGENTS.md

Guidance for AI coding agents working in this repo.

## What ExternalDNS Does

ExternalDNS = Kubernetes controller. Syncs exposed resources (Services, Ingresses, CRDs, Gateway API routes) with external DNS providers (AWS Route53, Google Cloud DNS, Azure DNS, Cloudflare, 20+ others). Tracks ownership via TXT record markers — safe for non-empty hosted zones.

## Commands

```bash
make build          # Build binary to build/external-dns
make test           # Run tests with race detection: go test -race ./...
make go-lint        # Run golangci-lint --fix ./...
make lint           # Run all linters (license headers + golangci-lint)
make licensecheck   # Check license headers in Go files
make crd            # Regenerate CRD using controller-gen
```

Run single test:

```bash
go test -race ./controller -run TestRunOnce
go test -race ./source -run TestIngressEndpoints
```

All CLI flags also settable as env vars (e.g., `--dry-run` → `EXTERNAL_DNS_DRY_RUN=1`).

## Architecture

Entry point: `main.go` → `controller.Execute()`. Four key abstractions:

**Source** (`source/`) — reads Kubernetes resources, returns desired DNS records as `[]*endpoint.Endpoint`. Implementations: ingress, service, gateway, CRD, pod, connector. `source/store.go` composes + caches multiple sources.

**Provider** (`provider/`) — manages actual DNS records at provider. Each provider in own subdir (`provider/aws/`, `provider/cloudflare/`, etc.). `provider/webhook/` enables out-of-tree providers. `provider/factory/` selects provider by name. `provider/cached_provider.go` wraps any provider with caching.

**Registry** (`registry/`) — wraps Provider to track ownership. TXT registry (`registry/txt/`) marks owned records with TXT records containing owner ID (set via `--txt-owner-id`). Other backends: `registry/awssd/`, `registry/dynamodb/`, `registry/noop/`.

**Controller** (`controller/controller.go`) — orchestrates `RunOnce()`:

1. `Registry.Records()` → current DNS state
2. `Source.Endpoints()` → desired DNS state
3. `Registry.AdjustEndpoints()` → normalize/adapt endpoints
4. `plan.Plan{}` → calc diff (create/update/delete)
5. `Registry.ApplyChanges()` → apply diff

Main loop in `controller.Run()` calls `RunOnce()` on interval (default 1 min), optionally on Kubernetes watch events (`--update-events`), throttled by `--min-event-sync-interval`.

## Key Data Structures

**`endpoint.Endpoint`** (`endpoint/endpoint.go`) — one DNS record: `DNSName`, `RecordType` (A, AAAA, CNAME, TXT, MX, NS, SRV, PTR, NAPTR, DNAME), `Targets`, `TTL`, `SetIdentifier`, `ProviderSpecific`, routing fields (Weights, Latency, Geolocation).

**`plan.Changes`** (`plan/plan.go`) — diff output: `Create`, `UpdateOld`, `UpdateNew`, `Delete` slices of `*endpoint.Endpoint`.

**`plan.Policy`** (`plan/policy.go`) — controls permitted changes: `SyncPolicy` (all), `UpsertOnlyPolicy` (no deletes), `CreateOnlyPolicy` (creates only).

**`pkg/apis/externaldns/types.go`** — `Config` struct with all CLI flags (~300+ fields).

## Providers

**No new in-tree providers accepted.** Use the webhook mechanism (`provider/webhook/`) for new providers — see `docs/tutorials/webhook-provider.md`.

Existing in-tree providers are being actively migrated out-of-tree (see issue #4347), starting with unmaintained ones. Do not add new subdirs under `provider/`.

## Adding a New Source

1. Create `source/<name>.go` implementing `Source` interface (`Endpoints`, `AddEventHandler`).
2. Register in `source/store.go`.

## Comments

Comments capture intent: why the code is shaped this way, what constraint forced
it, what breaks if you change it. The code already shows what it does.

Do not narrate control flow or restate an `if` condition in prose; if the
comment goes stale when the line below it changes, delete it. Keep comments
short. Exported-identifier doc comments are the exception: a concise behavioral
summary is their job.

## Error Handling

Hard errors (a plain `error`) are fine at startup: provider/source construction,
credential loading, and config validation should fail fast. Inside the reconcile
loop a plain `error` from a `Source` or `Provider` is fatal (the process exits
and Kubernetes CrashLoopBackOffs the pod); wrap transient or expected conditions
with `provider.NewSoftErrorf`, and return `nil` for "nothing to do".

Full details in `docs/contributing/sources-and-providers.md` ("Error Handling").
When you learn something non-obvious about a source or provider, update that
guide.

## Commit and PR Titles

PR titles must follow Conventional Commits: `type(scope): imperative summary`
(e.g. `fix(aws): respect --dry-run on record updates`). The title becomes the
squashed commit and feeds the release changelog, so the type and scope matter.
Common types: `feat`, `fix`, `docs`, `chore`, `refactor`, `test`, `build`. See
`CONTRIBUTING.md`.

## Linting Notes

Uses `golangci-lint` with strict `.golangci.yml` (32+ linters). Key rules: `testifylint` (use `assert`/`require` helpers correctly), `errorlint` (wrap errors properly), `gocritic`, `gochecknoinits` (no `init()` functions). All new Go files need Apache 2.0 license header.

## Before You Finish

Self-check the change against this guide before handing it back:

- Comments explain *why*, not *what* — no prose mirroring an `if` condition.
- Provider/source errors: hard only for unrecoverable startup conditions,
  otherwise `provider.NewSoftErrorf` or return `nil`.
- Ran `make go-lint` locally, and `go test -race ./<pkg>/...` for every package
  you touched (`make test` for cross-cutting changes). Both green.
- PR title follows Conventional Commits.
