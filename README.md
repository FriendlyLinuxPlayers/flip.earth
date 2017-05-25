# FLiP.Earth - Website at [FriendlyLinuxPlayers.org](https://FriendlyLinuxPlayers.org)

Copyright © 2017 Stephan "DerRidda" Lanfermann and Andrew "HER0" Conrad

FLiP.Earth is the website for the Friendly Linux Players community. Features
include being incomplete and having grand plans implied in the source code.

This code is freely available on
[GitHub](https://github.com/FriendlyLinuxPlayers/flip.earth) under the terms of
the GNU AGPL version 3. See `LICENSE` for details.

## Usage

# Requirements

* [Go](https://golang.org)

In the future you will probably need PostrgeSQL.

We suggest exposing this to the internet via a reverse proxy, such as
[NGINX](https://www.nginx.com).

# Download and Install

First download and install everything you'll need:

`go get -u github.com/friendlylinuxplayers/flip.earth`

Then build the project:

`cd $GOPATH/src/github.com/friendlylinuxplayers/flip.earth`

`go build`

# Configure

Change things in config.json until perfection is reached.

## Help and Contributing

If you have a bug to report, a feature request, or the desire to otherwise
contribute, you may either submit an
[issue on GitHub](https://github.com/FriendlyLinuxPlayers/flip.earth/issues)
or join us in the [Matrix](https://matrix.org) room for our website, at
[#website:flip.earth](https://matrix.to/#/#website:flip.earth).

For fixes that pass all tests, pull requests are accepted, but to avoid any
wasted effort, please contact us before submitting a pull request with other
changes.
