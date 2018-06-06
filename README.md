# FLiP.Earth - Website at [FriendlyLinuxPlayers.org](https://FriendlyLinuxPlayers.org)

Copyright © 2017 Stephan "DerRidda" Lanfermann and Andrew "HER0" Conrad

FLiP.Earth is the website for the Friendly Linux Players community. Features
include being incomplete and having grand plans implied in the source code.

This code is freely available on
[GitLab](https://gitlab.com/FriendlyLinuxPlayers/flip.earth) under the terms of
the GNU AGPL version 3. See `LICENSE` for details.

## Usage

# Requirements

* [Go](https://golang.org)
* [Dep](https://github.com/golang/dep) (this should be part of Go soon)

In the future you will probably need PostrgeSQL.

We suggest exposing this to the internet via a reverse proxy, such as
[NGINX](https://www.nginx.com).

# Download and Install

First download and install everything you'll need:

`go get -u gitlab.com/friendlylinuxplayers/flip.earth`

`cd $GOPATH/src/gitlab.com/friendlylinuxplayers/flip.earth`

`dep ensure`

Then build the project:

`go build`

# Configure

Change things in config/config.json until perfection is reached.

## Help and Contributing

By participating in this community, you agree to abide by the terms of our
[Code of Conduct](https://FriendlyLinuxPlayers.org/conduct).

If you have a bug to report, a feature request, or the desire to otherwise
contribute, you may either submit an
[issue on GitLab](https://gitlab.com/FriendlyLinuxPlayers/flip.earth/issues) or
join us in the [Matrix](https://matrix.org) room for our website, at
[#website:flip.earth](https://matrix.to/#/#website:flip.earth) (our web client
is available [here](https://riot.flip.earth/#/room/#website:flip.earth)).

For fixes that pass all tests, pull requests are accepted, but to avoid any
wasted effort, please contact us before submitting a pull request with other
changes.
