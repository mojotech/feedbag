package main

import (
	"fmt"
	"time"

	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
)

func init() {
	fmt.Println("Starting go routines")
}

func StartGoRoutine(u *User) {
	go userRoutine(u)
}

func userRoutine(u *User) error {
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: u.AccessToken},
	}

	client := github.NewClient(t.Client())

	//List Options Page, PerPage
	opts := github.ListOptions{1, 50}

	for {
		event, _, err := client.Activity.ListEvents(&opts)
		if err != nil {
			panic(err)
		}

		//DEBUG: print out events
		fmt.Println(event)

		// Wait 5 seconds after events are recieved and
		// start again
		time.Sleep(5 * time.Second)
	}

	panic("Shouldn't be here")
}
