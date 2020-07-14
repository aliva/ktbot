package main

import (
	"os"

	twitterOAuth1 "github.com/dghubble/oauth1/twitter"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func getClient() *twitter.Client {
	config := oauth1.Config{
		ConsumerKey:    os.Getenv("KTBOT_CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("KTBOT_CONSUMER_SECRET"),
		Endpoint:       twitterOAuth1.AuthorizeEndpoint,
	}
	token := &oauth1.Token{
		Token:       os.Getenv("KTBOT_TOKEN"),
		TokenSecret: os.Getenv("KTBOT_TOKEN_SECRET"),
	}

	httpClient := config.Client(oauth1.NoContext, token)

	return twitter.NewClient(httpClient)
}
