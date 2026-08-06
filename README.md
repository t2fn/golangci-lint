<p align="center">
  <img alt="golangci-lint logo" src="assets/go.png" height="150" />
  <h3 align="center">golangci-lint</h3>
  <p align="center">Fast linters runner for Go</p>
</p>

---

`golangci-lint` is a fast Go linters runner.

It runs linters in parallel, uses caching, supports YAML configuration,
integrates with all major IDEs, and includes over a hundred linters.

**This is the t2fn enhanced fork.** It integrates `github.com/t2fn/godoc-lint` 
with full support for all 9 checkers including the critical `require-func-docs` rule 
for enforcing documented function parameters and return values. This fork also incorporates
community-requested improvements from upstream issues while maintaining backward compatibility.

## Godoc-lint Integration

This fork includes a custom integration with [t2fn/godoc-lint](https://github.com/t2fn/godoc-lint) that provides:
- `broken-doclink`: Detects broken documentation links in Go comments
- `require-func-docs`: Enforces documented function parameters and return values  
- `single-pkg-doc`/`require-pkg-doc`: Package documentation rules
- And more...

### Development Setup

For local development with the forked godoc-lint:

```bash
# Clone this repo
git clone https://github.com/t2fn/golangci-lint.git
cd golangci-lint

# Clone the godoc-lint fork (if not already done)
git clone https://github.com/t2fn/godoc-lint.git /workdir/godoc-lint  # Adjust path as needed

# The go.mod replace directive points to local fork for development
go mod tidy && make build

# Run tests
go test ./pkg/golinters/godoclint/...
```

### Production/Release Setup

For `go install` and CI/CD, the replace directive should point to the published version:

```bash
# Edit go.mod - comment out local replace and use GitHub version:
# replace github.com/t2fn/godoc-lint => /workdir/godoc-lint
replace github.com/t2fn/godoc-lint => github.com/t2fn/godoc-lint v0.11.3

go mod tidy && go install ./cmd/golangci-lint@latest
```

Or publish a tagged release of golangci-lint with the proper replace directive for production use.

## Install `golangci-lint`

- [On my machine](https://golangci-lint.run/docs/welcome/install/local);
- [On CI/CD systems](https://golangci-lint.run/docs/welcome/install/ci).

## Documentation

Documentation is hosted at https://golangci-lint.run.

## Social Networks

[![Join Slack](https://img.shields.io/badge/Slack-4285F4?logo=slack&logoColor=white)](https://gophers.slack.com/archives/CS0TBRKPC)
[![Follow on Mastodon](https://img.shields.io/badge/Mastodon-6364FF?logo=mastodon&logoColor=white)](https://fosstodon.org/@golangcilint)
[![Follow on Bluesky](https://img.shields.io/badge/Bluesky-0a7aff?logo=bluesky&logoColor=white)](https://bsky.app/profile/golangci-lint.run)
[![Follow on Twitter](https://img.shields.io/badge/Twitter-1DA1F2?logo=x&logoColor=white)](https://twitter.com/golangci)

## Support Us

`golangci-lint` is a free and open-source project built by volunteers.

## Badges

![Build Status](https://github.com/golangci/golangci-lint/workflows/CI/badge.svg)
[![License](https://img.shields.io/github/license/golangci/golangci-lint)](/LICENSE)
[![Release](https://img.shields.io/github/release/golangci/golangci-lint.svg)](https://github.com/golangci/golangci-lint/releases/latest)
[![Docker](https://img.shields.io/docker/pulls/golangci/golangci-lint)](https://hub.docker.com/r/golangci/golangci-lint)
[![GitHub Releases Stats of golangci-lint](https://img.shields.io/github/downloads/golangci/golangci-lint/total.svg?logo=github)](https://somsubhra.github.io/github-release-stats/?username=golangci&repository=golangci-lint)

## Contributors

This project exists thanks to all the people who contribute. [How to contribute](https://golangci-lint.run/docs/contributing/).

<a href="https://github.com/golangci/golangci-lint/graphs/contributors">
  <img src="https://opencollective.com/golangci-lint/contributors.svg?width=890&button=false&skip=golangcidev,CLAassistant,renovate,fossabot,golangcibot,kortschak,golangci-releaser,dependabot%5Bbot%5D" />
</a>

## Sponsors

<p>&nbsp;</p>
<p float="left">
  <a href="https://www.jetbrains.com/go/?utm_source=OSS&utm_medium=referral&utm_campaign=golangci" target="_blank">
    <picture>
      <source media="(prefers-color-scheme: dark)" srcset="assets/goland-white.svg">
      <source media="(prefers-color-scheme: light)" srcset="assets/goland.svg">
      <img alt="The complete IDE crafted for professional Go developers." src="assets/goland.svg" width="150" />
    </picture>
  </a>
</p>

## Stargazers over time

[![Stargazers over time](https://starchart.cc/golangci/golangci-lint.svg?variant=adaptive)](https://starchart.cc/golangci/golangci-lint)
