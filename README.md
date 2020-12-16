# Branch notes

This branch (mini-opctl) is an exploration into fixing some of my main concerns
with opctl.

## reliability/complexity

I find opctl quite unreliable. The major issues are issues connecting to port
42224 and opctl not properly cleaning up after itself. The core architecture of
opctl starts up a persistent webserver on 42224 that actually manages running
the core business logic of opctl. The CLI then interacts with that webserver
using http api calls and a websocket connection to pipe output back to the user.

Since opctl's primary value to me is a reusable "op" runner, this webserver
feels very unnecessary. It doesn't maintain much (any?) internal state, such as
to allow me to query and interact with running ops. The network interactions
are what I suspect are the primary causes of reliability issues.

Because of this, I've entirely removed this webserver component of opctl in
favor of the CLI directly running core logic. This allows me to pass a
cancellation context properly, which I'm hoping will prevent errors leaving
active containers hanging around.

## line count

This project is _huge_ and can be difficult to work in. This removes checked-in
vendored code, the web UI for opctl, the JS sdk, and the react SDK.

The project also has many layers of abstraction, that I feel could be reduced
to make changes easier. I also think the code could be refactored to be more
idomatic to the go language.

## usability

For complex ops, opctl makes it difficult to understand what's going on. I hope
to improve the output of the CLI tool to allow me to identify what produces
what output.

---

[![Build](https://github.com/opctl/opctl/workflows/Build/badge.svg)](https://github.com/opctl/opctl/actions?query=workflow%3ABuild+branch%3Agithub-actions)
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


