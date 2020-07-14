package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dghubble/oauth1"
	twitterOAuth1 "github.com/dghubble/oauth1/twitter"
)

func main() {
	config := oauth1.Config{
		ConsumerKey:    os.Getenv("KTBOT_CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("KTBOT_CONSUMER_SECRET"),
		Endpoint:       twitterOAuth1.AuthorizeEndpoint,
	}

	requestToken, requestSecret, err := config.RequestToken()
	fmt.Println(requestToken, requestSecret, err)

	authorizationURL, err := config.AuthorizationURL(requestToken)
	fmt.Println(authorizationURL)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	verifier, _ := reader.ReadString('\n')
	accessToken, accessSecret, err := config.AccessToken(requestToken, requestSecret, verifier)
	fmt.Println(accessToken, accessSecret, err)
}
