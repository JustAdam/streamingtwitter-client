Streamingtwitter package client example
=======================================

[![Build Status](https://travis-ci.org/JustAdam/streamingtwitter-client.svg?branch=master)](https://travis-ci.org/JustAdam/streamingtwitter-client)

The examples in this repository use a file for storage of the tokens.
When calling Authenticate() without a User token, one will be requested and returned.

Quick start
-----------

	$ go get github.com/garyburd/go-oauth/oauth
 	$ go get github.com/JustAdam/streamingtwitter

	$ mv tokens.json.sample tokens.json

	$ vim tokens.json
		Add your twitter API token and secret under "App".

	$ go run client-simple.go
