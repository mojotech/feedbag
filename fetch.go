package main

import (
	"fmt"
	"time"

	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
)

func init() {
	fmt.Println("Starting go routines")

	err := startExistingUsers()
	if err != nil {
		panic(err)
	}
}

func startExistingUsers() error {
	u := UserList{}

	err := u.List()
	if err != nil {
		return err
	}

	for _, user := range u {
		StartUserRoutine(user)
	}

	return nil
}

func StartUserRoutine(u User) {
	go userRoutine(u)
}

func userRoutine(u User) error {
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: u.AccessToken},
	}

	client := github.NewClient(t.Client())

	//List Options Page, PerPage
	opts := github.ListOptions{1, 50}

	for {
		events, _, err := client.Activity.ListEventsPerformedByUser(u.Username, false, &opts)
		if err != nil {
			panic(err)
		}

		activityPayload, err := ProcessPayload(events, u)
		ActivityParser(activityPayload)

		fmt.Println(activityPayload)

		// Wait 5 seconds after events are recieved and
		// start again
		time.Sleep(5 * time.Second)
	}

	panic("Shouldn't be here")
}
