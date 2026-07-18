# Security Policy

## Reporting a vulnerability

Report vulnerabilities privately through [GitHub private vulnerability reporting](https://github.com/tiderops/tider/security/advisories/new) ("Report a vulnerability" on the repository's Security tab).
## Supported versions

Tider is pre-1.0.0 Only the latest release (and the `master` branch) receive security fixes.

## Security posture

Tider is a local-first desktop application. Its security model is intentionally narrow:

- **Read-only Kubernetes access.** Tider observes and diagnoses; it does not expose patch, delete, exec, apply, scale, or port-forward operations. It never requests broader RBAC than its features require (get/list/watch on the supported resource set).
- **Your identity, not ours.** Tider uses your existing kubeconfig and identity chain. There is no TiderOps account, login, or hosted control plane.
- **No data egress by default.** No raw manifests, logs, secrets, or cluster identifiers leave your workstation. Product telemetry, when it ships, will be opt-in and limited to a documented event allowlist with no raw Kubernetes content.
- **Offline capable.** Core functionality requires no network access beyond the target cluster.
- **Supply chain.** Release artifacts are planned to be signed and accompanied by an SBOM; dependencies are pinned and built through GitHub Actions.

Reports that identify gaps between this posture and the actual behavior of the code are especially valuable.

## Scope notes

- Vulnerabilities in Kubernetes itself, client-go, or your cluster configuration are out of scope (report those upstream), but ways Tider could be made to *misuse* them are in scope.