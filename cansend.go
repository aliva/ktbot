package main

import (
	"fmt"
	"time"

	"github.com/dghubble/go-twitter/twitter"
)

func canSend() bool {
	client := getClient()
	params := &twitter.UserTimelineParams{Count: 1}
	tweets, _, err := client.Timelines.UserTimeline(params)

	if err != nil {
		fmt.Println(err)
		return false
	}

	t, err := tweets[0].CreatedAtTime()

	if err != nil {
		fmt.Println(err)
		return false
	}

	d := time.Now().Sub(t)
	dm := int64(d / time.Minute)

	if dm < 60*3 {
		return false
	}

	return true
}
