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

	//Events
	OpenPullRequest          bool `json:"open_pull_request"`
	ClosePullRequest         bool `json:"close_pull_request"`
	AssignedPullRequest      bool `json:"assigned_pull_request"`
	UnassignedPullRequest    bool `json:"unassigned_pull_request"`
	LabledPullRequest        bool `json:"labled_pull_request"`
	UnlabeledPullRequest     bool `json:"unlabeled_pull_request"`
	ReopenedPullRequest      bool `json:"repoened_pull_request"`
	SynchronizedPullRequest  bool `json:"synchronize_pull_request"`
	MergedPullRequest        bool `json:"merged_pull_request"`
	CreatePullRequestComment bool `json:"create_pull_request_comment"`
	OpenIssue                bool `json:"open_issue"`
	CloseIssue               bool `json:"close_pull_request"`
	AssignedIssue            bool `json:"assigned_issue"`
	UnassignedIssue          bool `json:"unassigned_issue"`
	LabeledIssue             bool `json:"labeled_issue"`
	UnlabeledIssue           bool `json:"unlabeled_issuse"`
	ReopenedIssue            bool `json:"reopened_issue"`
	Fork                     bool `json:"fork"`
	Push                     bool `json:"push"`
	PushCreated              bool `json:"push_created"`
	PushDeleted              bool `json:"push_deleted"`
	ForcePush                bool `json:"force_push"`
	CreateCommitComment      bool `json:"create_commit_comment"`
	CreateIssueComment       bool `json:"create_issue_comment"`

	//User
	EventOwner User `json:"action_creator"db:"-"`

	//Payload Fields
	Action       string `json:"action,omitempty"`
	Number       int    `json:"number,omitempty"`
	Ref          string `json:"ref,omitempty"`
	Created      bool   `json:"created,omitempty"`
	Deleted      bool   `json:"deleted,omitempty"`
	Forced       bool   `json:"forced,omitempty"`
	Size         int    `json:"size,omitempty"`
	DistinctSize int    `json:"distinct_size,omitempty"`

	//Payload Structs
	PullRequest PullRequest `json:"pull_request"db:"-"`
	Repository  Repository  `json:"repository"db:"-"`
	Issue       Issue       `json:"issue"db:"-"`
	Comment     Comment     `json:"comment"db:"-"`
	Label       Label       `json:"label"db:"-"`
	Sender      GithubUser  `json:"sender"db:"-"`
	Assignee    GithubUser  `json:"assignee"db:"-"`
	Forkee      Forkee      `json:"forkee"db:"-"`
	Commits     Commits     `json:"commits"db:"-"`
	Pusher      GithubUser  `json:"pusher"db:"-"`
	HeadCommit  Commit      `json:"head_commit"db:"-"`
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
		a.EventOwner = u

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

	case "PushEvent":
		if a.Forced {
			a.EventType = "ForcePush"
			a.ForcePush = true
		} else if a.Created {
			a.EventType = "PushCreated"
			a.PushCreated = true
		} else if a.Deleted {
			a.EventType = "PushDeleted"
			a.PushDeleted = true
		} else {
			a.EventType = "Push"
			a.Push = true
		}

	case "PullRequestEvent":
		switch a.Action {
		case "opened":
			a.EventType = "OpenPullRequest"
			a.OpenPullRequest = true
		case "closed":
			if a.PullRequest.Merged {
				a.EventType = "MergedPullRequest"
				a.MergedPullRequest = true
			} else {
				a.EventType = "ClosePullRequest"
				a.ClosePullRequest = true
			}
		case "assigned":
			a.EventType = "AssignedPullRequest"
			a.AssignedPullRequest = true
		case "unassigned":
			a.EventType = "UnassignedPullRequest"
			a.UnassignedPullRequest = true
		case "labled":
			a.EventType = "LabeledPullRequest"
			a.LabledPullRequest = true
		case "unlabeled":
			a.EventType = "UnlabeledPullRequest"
			a.UnlabeledPullRequest = true
		case "reopened":
			a.EventType = "ReopenedPullRequest"
			a.ReopenedPullRequest = true
		case "synchronize":
			a.EventType = "SynchronizedPullRequest"
			a.SynchronizedPullRequest = true
		}

	case "IssueEvent":
		switch a.Action {
		case "opened":
			a.EventType = "OpenIssue"
			a.OpenIssue = true
		case "closed":
			a.EventType = "CloseIssue"
			a.CloseIssue = true
		case "assigned":
			a.EventType = "AssignedIssue"
			a.AssignedIssue = true
		case "unassigned":
			a.EventType = "UnassignedIssue"
			a.UnassignedIssue = true
		case "labeled":
			a.EventType = "LabledIssue"
			a.LabeledIssue = true
		case "unlabeled":
			a.EventType = "UnlabeledIssue"
			a.LabeledIssue = true
		case "reopened":
			a.EventType = "ReopenedIssue"
			a.ReopenedIssue = true
		}

	case "CommitCommentEvent":
		a.EventType = "CreateCommitComment"
		a.CreateCommitComment = true

	case "ForkEvent":
		a.EventType = "Fork"
		a.Fork = true

	case "PullRequestReviewCommentEvent":
		a.EventType = "CreatePullRequestComment"
		a.CreatePullRequestComment = true

	default:
		return false
	}
	return true
}
