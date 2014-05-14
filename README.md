Streamingtwitter package client example
=======================================

	$ go get github.com/garyburd/go-oauth/oauth
 	$ go get github.com/JustAdam/streamingtwitter

	$ mv tokens.json.sample tokens.json

	$ vim tokens.json
		Add your twitter API token and secret under "App".

	$ go run client-simple.go
