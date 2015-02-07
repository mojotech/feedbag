package feedbag_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"

	"github.com/google/go-github/github"
	"github.com/mojotech/feedbag/feedbag"
	. "github.com/smartystreets/goconvey/convey"
)

func TestProcessPayload(t *testing.T) {

	Convey("Given a malformed payload", t, func() {
		e := generateEvent("12345678", "IssueCommentEvent", "./testing/malformed_payload.json")
		u := generateUser()

		activityPayload, err := feedbag.ProcessPayload(e, u)

		Convey("The Activity Payload should be Empty", func() {
			So(err, ShouldBeNil)
			So(activityPayload, ShouldBeEmpty)
		})
	})

	Convey("Given an unrecognized Github event", t, func() {
		e := generateEvent("12345678", "BadEvent", "./testing/issue_comment_payload.json")
		u := generateUser()

		activityPayload, err := feedbag.ProcessPayload(e, u)

		Convey("An activity payload should not be created", func() {
			So(err, ShouldBeNil)
			So(activityPayload, ShouldBeEmpty)
		})
	})

	Convey("Given a Github CreateIssueComment event", t, func() {
		e := generateEvent("12345678", "IssueCommentEvent", "./testing/issue_comment_payload.json")
		u := generateUser()

		activityPayload, err := feedbag.ProcessPayload(e, u)
		payload := activityPayload[0]

		Convey("The EventType should equal the Github event type", func() {
			So(err, ShouldBeNil)
			So(activityPayload, ShouldNotBeNil)
		})

		Convey("The Activity Payload CreateIssueComment should be true", func() {
			So(payload, ShouldNotBeNil)
			So(payload.CreateIssueComment, ShouldBeTrue)
		})

		Convey("The github_id should equal the Github event id", func() {
			So(payload.GithubId, ShouldEqual, "12345678")
		})

		Convey("The title should equal the Github event title", func() {
			So(payload.Issue.Title, ShouldEqual, "Title")
		})

		Convey("The body should equal the Github event body", func() {
			So(payload.Comment.Body, ShouldEqual, "Body")
		})

		Convey("The issue labels should not be empty", func() {
			So(payload.Issue.Labels, ShouldNotBeEmpty)
		})
	})

}

/*
	type github.Event struct {
			Type       *string          `json:"type,omitempty"`
			Public     *bool            `json:"public"`
			RawPayload *json.RawMessage `json:"payload,omitempty"`
			Repo       *Repository      `json:"repo,omitempty"`
			Actor      *User            `json:"actor,omitempty"`
			Org        *Organization    `json:"org,omitempty"`
			CreatedAt  *time.Time       `json:"created_at,omitempty"`
			ID         *string          `json:"id,omitempty"`
	}
*/
func generateEvent(id string, eventType string, payloadFile string) []github.Event {

	// Raw Payload
	var rawPayload json.RawMessage
	file, _ := ioutil.ReadFile(payloadFile)
	rawPayload = file
	public := true

	return []github.Event{
		{
			&eventType,
			&public,
			&rawPayload,
			&github.Repository{},
			&github.User{},
			&github.Organization{},
			&time.Time{},
			&id,
		},
	}
}

func generateUser() feedbag.User {
	return feedbag.User{
		1,
		"John Doe",
		"jdoe",
		"http://logo.com/logo.png",
		"http://profile.com/profile",
		"j.doe@gmail.com",
		"12/04/2014",
		feedbag.RawJson{
			"raw": "user",
		},
		"12345678910",
		feedbag.TimeStamp{},
	}
}
