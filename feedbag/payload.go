package feedbag

import (
	"encoding/json"

	"github.com/google/go-github/github"
)

type ActivityPayload struct {
	Id         int64   `json:"id"`
	GithubId   string  `json:"github_id"`
	RawPayload RawJson `json:"-"`

	EventType string `json:"event_type"`

	//User
	ActionCreator User `json:"action_creator"db:"-"`

	//Events
	OpenPullRequest    bool `json:"open_pull_request"`
	ClosePullRequest   bool `json:"close_pull_request"`
	OpenIssue          bool `json:"open_issue"`
	CreateIssueComment bool `json:"create_issue_comment"`

	//Payload Structs
	Action      string      `json:"action,omitempty"`
	Number      int         `json:"number,omitempty"`
	PullRequest PullRequest `json:"pull_request"db:"-"`
	Repository  Repository  `json:"repository"db:"-"`
	Issue       Issue       `json:"issue"db:"-"`
	Comment     Comment     `json:"comment"db:"-"`
	Label       Label       `json:"label"db:"-"`
	Sender      GithubUser  `json:"sender"db:"-"`
	Assignee    GithubUser  `json:"assignee"db:"-"`
}

type ActivityPayloadList []ActivityPayload

func (p *ActivityPayload) Create() error {
	err := dbmap.Insert(p)
	return err
}

func (p ActivityPayloadList) SaveUnique() (l ActivityPayloadList) {
	for _, activity := range p {
		err := activity.Create()
		if err == nil {
			l = append(l, activity)
		}
	}
	return l
}

func ProcessPayload(e []github.Event, u User) (ActivityPayloadList, error) {
	activityPayloads := ActivityPayloadList{}

	for _, event := range e {
		a := ActivityPayload{}

		// Marshal raw payload into activity payload struct.
		if err := json.Unmarshal(*event.RawPayload, &a); err != nil {
			// If error unmarshaling payload reset the loop
			continue
		}

		// Set Github Id of event. Has Unique constraint on the database.
		a.GithubId = *event.ID

		// Set Raw Payload. We don't care if this errors out.
		_ = json.Unmarshal(*event.RawPayload, &a.RawPayload)

		// Set the event type and reset the loop if no match is found.
		if !a.setEventType(event) {
			continue
		}

		// Add action creator
		a.ActionCreator = u

		activityPayloads = append(activityPayloads, a)
	}

	return activityPayloads, nil
}

func (a *ActivityPayload) setEventType(e github.Event) bool {
	switch *e.Type {
	case "IssueCommentEvent":
		if a.Action == "created" {
			a.EventType = "CreateIssueComment"
			a.CreateIssueComment = true
		}
	case "PullRequestEvent":
		switch a.Action {
		case "opened":
			a.EventType = "OpenPullRequest"
			a.OpenPullRequest = true
		case "closed":
			a.EventType = "ClosePullRequest"
			a.ClosePullRequest = true
		}
	case "IssueEvent":
		switch a.Action {
		case "opened":
			a.EventType = "OpenIssue"
			a.OpenIssue = true
		}
	default:
		return false
	}
	return true
}
