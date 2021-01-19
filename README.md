### Branch notes

Goals:
- Allow understanding where output is coming from
- Clean up output to be more consistent with other CLI tools
- Clean up output to be more easily parsable with other CLI tools

I want op references to be human readable, because they can be very long. This
means I need to pass some runtime context into the things printing them.
For this reason, the existing cliOutput structure doesn't work great, since it's
shared with other places that don't need that context, and the plumbing gets
large.

New structure
- Color - semantic coloring utilities
- CliOutput - links writers to semantic coloring utilities
- OpOutputter - intercepts event stream and outputs events to writers

---

[![Build](https://github.com/opctl/opctl/workflows/Build/badge.svg?branch=main)](https://github.com/opctl/opctl/actions?query=workflow%3ABuild+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/opctl/opctl)](https://goreportcard.com/report/github.com/opctl/opctl)
[![Coverage](https://codecov.io/gh/opctl/opctl/branch/master/graph/badge.svg)](https://codecov.io/gh/opctl/opctl)

> *Be advised: this project is currently at Major version zero. Per the
> semantic versioning spec: "Major version zero (0.y.z) is for initial
> development. Anything may change at any time. The public API should
> not be considered stable."*

# Documentation

see [website](https://opctl.io)

# Used By

These awesome companies use opctl. represent by adding yours (or one you know of) to the list!
- [Era](https://helloera.co)
- [Expedia](https://www.expedia.com)
- [Nintex](https://www.nintex.com)
- [ProKarma](https://prokarma.com/)
- [Remitly](https://www.remitly.com)
- [Samsung (SDS)](https://www.samsungsds.com)

# Support

join us on
[![Slack](https://img.shields.io/badge/slack-opctl-E01563.svg)](https://join.slack.com/t/opctl/shared_invite/zt-51zodvjn-Ul_UXfkhqYLWZPQTvNPp5w)
or [open an issue](https://github.com/opctl/opctl/issues)

# Releases

releases are versioned according to
[![semver 2.0.0](https://img.shields.io/badge/semver-2.0.0-brightgreen.svg)](http://semver.org/spec/v2.0.0.html)
and [tagged](https://git-scm.com/book/en/v2/Git-Basics-Tagging); see
[CHANGELOG.md](CHANGELOG.md) for release notes

# Contributing

see [CONTRIBUTING.md](CONTRIBUTING.md)


