Streamingtwitter package client example
=======================================

[![Build Status](https://travis-ci.org/JustAdam/streamingtwitter-client.svg?branch=master)](https://travis-ci.org/JustAdam/streamingtwitter-client)

	$ go get github.com/garyburd/go-oauth/oauth
 	$ go get github.com/JustAdam/streamingtwitter

	$ mv tokens.json.sample tokens.json

	$ vim tokens.json
		Add your twitter API token and secret under "App".

	$ go run client-simple.go
