// Sample Twitter client using the streaming twitter API.
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

	// Some keywords to track .. see the Twitter Streaming API documentation for more information
	args := &url.Values{}
	args.Add("track", "Norway")

	// Launch the stream
	tweets := make(chan *streamingtwitter.TwitterStatus)
	go client.Stream(tweets, streamingtwitter.Streams["Filter"], args)

	for {
		select {
		// Recieve tweets
		case status := <-tweets:
			fmt.Println(status)
			// Any errors that occured
		case err := <-client.Errors:
			fmt.Printf("ERROR: '%s'\n", err)
			// Stream has finished
		case <-client.Finished:
			return
		}
	}
}
