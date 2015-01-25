package main

import (
	"encoding/json"

	"github.com/google/go-github/github"
)

type ActivityPayload struct {
	EventType string `json:"event_type"`
	//All possible variables/events/etc
	OpenPullRequest    bool   `json:"open_pull_request"`
	ClosePullRequest   bool   `json:"close_pull_request"`
	CreateIssueComment bool   `json:"create_issue_comment"`
	PRTitle            string `json:"pr_title"`
	PRNumber           int    `json:"pr_number"`
	IssueNumber        int    `json:"issue_number"`
}

func ProcessPayload(e []github.Event, u User) ([]ActivityPayload, error) {
	activityPayloads := []ActivityPayload{}

	for _, event := range e {
		payload := ActivityPayload{}

		rawPayload := make(map[string]interface{})
		if err := json.Unmarshal(*event.RawPayload, &rawPayload); err != nil {
			panic(err.Error())
		}

		switch *event.Type {
		case "PullRequestEvent":
			// Set Pull Request Event type (open or closed)
			if val, ok := rawPayload["action"]; ok && val == "opened" {
				payload.EventType = "OpenPullRequest"
				payload.OpenPullRequest = true
			} else {
				payload.EventType = "ClosePullRequest"
				payload.ClosePullRequest = true
			}

			if pr, ok := rawPayload["pull_request"].(map[string]interface{}); ok {

				// Set title of PR
				payload.PRTitle = pr["title"].(string)

				// Set Pr Number
				payload.PRNumber = int(pr["number"].(float64))
			}
			activityPayloads = append(activityPayloads, payload)

		case "IssueCommentEvent":
			if val, ok := rawPayload["action"]; ok && val == "created" {
				payload.EventType = "CreateIssueComment"
				payload.CreateIssueComment = true
			}

			if issue, ok := rawPayload["issue"].(map[string]interface{}); ok {
				payload.IssueNumber = int(issue["number"].(float64))

			}
			activityPayloads = append(activityPayloads, payload)

		}
	}

	return activityPayloads, nil
}
