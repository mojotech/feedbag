package main

import (
	"fmt"
	"time"

	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
)

func StartExistingUsers(c chan<- []Activity) error {
	fmt.Println("Starting go routines")
	u := UserList{}

	err := u.List()
	if err != nil {
		return err
	}

	for _, user := range u {
		StartUserRoutine(user, c)
	}

	return nil
}

func StartUserRoutine(u User, c chan<- []Activity) {
	go userRoutine(u, c)
}

func userRoutine(u User, c chan<- []Activity) error {
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
		activities := ActivityParser(activityPayload)

		fmt.Println(activities)
		c <- activities

		// Wait 5 seconds after events are recieved and
		// start again
		time.Sleep(5 * time.Second)
	}

	panic("Shouldn't be here")
}
