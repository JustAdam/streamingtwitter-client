// Twitter REST API example - user lookup
// https://dev.twitter.com/docs/api/1.1/get/users/lookup
package main

import (
	"encoding/json"
	"fmt"
	"github.com/JustAdam/streamingtwitter"
	"io/ioutil"
	"log"
	"net/url"
	"os"
)

var (
	// Location to the file containing your app's Twitter API token information
	tokenFile = "tokens.json"
	// File permissions for the token file.
	tokenFilePermission = os.FileMode(0600)
)

func main() {

	// Create new streaming API client
	client := streamingtwitter.NewClient()

	// Token information is saved in a file ..
	cf, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		log.Fatal(err)
	}
	credentials := new(streamingtwitter.ClientTokens)
	if err := json.Unmarshal(cf, &credentials); err != nil {
		log.Fatal(err)
	}

	// Authenicate the client and the user
	userToken, err := client.Authenticate(credentials)
	if err != nil {
		log.Fatal(err)
	}
	// Save the user's token information
	if userToken != nil {
		credentials.User = userToken
		save, err := json.Marshal(credentials)
		if err != nil {
			log.Fatal(err)
		}

		if err := ioutil.WriteFile(tokenFile, save, tokenFilePermission); err != nil {
			log.Fatal(err)
		}
	}

	// Paramaters to send to the API
	args := &url.Values{}
	args.Add("screen_name", "stephenfry,mashable")

	userLookup := &streamingtwitter.TwitterAPIURL{
		AccessMethod: "get",
		URL:          "https://api.twitter.com/1.1/users/lookup.json",
	}

	data := []streamingtwitter.TwitterUser{}
	go client.Rest(&data, userLookup, args)

	select {
	case err := <-client.Errors:
		log.Fatal(err)
	case <-client.Finished:
		fmt.Printf("%+v", data)
	}
}
