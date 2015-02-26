package feedbag

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fogcreek/logging"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
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

type tokenSource struct {
	token *oauth2.Token
}

func (t tokenSource) Token() (*oauth2.Token, error) {
	return t.token, nil
}

func userRoutine(u User, c chan<- []Activity) error {

	ts := tokenSource{
		&oauth2.Token{
			AccessToken: u.AccessToken,
		},
	}

	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	//List Options Page, PerPage
	opts := github.ListOptions{1, 50}

	for {
		events, resp, err := client.Activity.ListEventsPerformedByUser(u.Username, false, &opts)
		if err != nil {
			logging.WarnWithTags([]string{"github"}, "Problem retrieving events for user", u.Username, err.Error())
		}

		activityPayload, err := ProcessPayload(events, u)
		if err != nil {
			logging.WarnWithTags([]string{"payload"}, "Failed to process payload", err.Error())
			continue
		}

		activities := ActivityParser(activityPayload.SaveUnique())

		c <- activities

		// Wait as long as the X-Poll-Interval header says to
		interval, err := strconv.ParseInt(resp.Header["X-Poll-Interval"][0], 10, 8)
		if err != nil {
			// if strconv failed for whatever reason, use the default X-Poll-Interval value of 60
			time.Sleep(60 * time.Second)
		} else {
			time.Sleep(time.Duration(interval) * time.Second)
		}
	}

	panic("Shouldn't be here")
}
