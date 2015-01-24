package main

import (
	"encoding/json"
	"fmt"

	"github.com/google/go-github/github"
)

type ActivityPayload struct {
	//All possible variables/events/etc
	OpenPullRequest  bool   `json:"open_pull_request"`
	ClosePullRequest bool   `json:"close_pull_request"`
	PRTitle          string `json:"pr_title"`
}

func ProcessPayload(e []github.Event, u User) ([]ActivityPayload, error) {
	activityPayloads := []ActivityPayload{}

	for _, event := range e {
		payload := ActivityPayload{}
		rawPayload := make(map[string]interface{})

		if err := json.Unmarshal(*event.RawPayload, &rawPayload); err != nil {
			panic(err.Error())
		}

		fmt.Println(*event.Type)

		switch *event.Type {
		case "PullRequestEvent":
			// Set Pull Request Event type (open or closed)
			if val, ok := rawPayload["action"]; ok && val == "opened" {
				payload.OpenPullRequest = true
			} else {
				payload.ClosePullRequest = true
			}

			if pr, ok := rawPayload["pull_request"].(map[string]interface{}); ok {

				// Set title of PR
				payload.PRTitle = pr["title"].(string)

			}

			activityPayloads = append(activityPayloads, payload)
		}

	}

	return activityPayloads, nil
}
