// Twitter REST API example - user lookup
// https://dev.twitter.com/docs/api/1.1/get/users/lookup
package main

import (
	"fmt"
	"github.com/JustAdam/streamingtwitter"
	"log"
	"net/url"
)

var (
	// Location to the file containing your app's Twitter API token information
	tokenFile = "tokens.json"
)

func main() {

	// Create new streaming API client
	client := streamingtwitter.NewClient()

	// Authenicate the client and the user using the information in the tokenFile
	if err := client.Authenticate(&tokenFile); err != nil {
		log.Fatal(err)
	}

	// Paramaters to send to the API
	args := &url.Values{}
	args.Add("screen_name", "stephenfry,mashable")

	userLookup := &streamingtwitter.TwitterStream{
		AccessMethod: "get",
		Url:          "https://api.twitter.com/1.1/users/lookup.json",
	}

	data := []streamingtwitter.TwitterUser{}
	go client.Rest(userLookup, args, &data)

	select {
	case err := <-client.Errors:
		log.Fatal(err)
	case <-client.Finished:
		fmt.Printf("%+v", data)
	}
}
