package feedbag_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/mojotech/feedbag/feedbag"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_CheckEventType(t *testing.T) {
	Convey("Given a templates event type", t, func() {
		p := buildPayload("./testing/issue_comment_activity_payload.json")

		Convey("It should pass with '*' event type", func() {
			condition := feedbag.CheckEventType("*", p)
			So(condition, ShouldBeTrue)
		})

		Convey("It should pass with a matching event type", func() {
			condition := feedbag.CheckEventType("CreateIssueComment", p)
			So(condition, ShouldBeTrue)
		})

		Convey("It should not pass with a non matching event type", func() {
			condition := feedbag.CheckEventType("OpenPullRequest", p)
			So(condition, ShouldBeFalse)
		})
	})
}

func Test_ValidateConditional(t *testing.T) {
	Convey("Given conditional statements", t, func() {

		p := buildPayload("./testing/issue_comment_activity_payload.json")

		Convey("It should pass with an empty conditional", func() {
			condition := feedbag.ValidateConditional("", p)
			So(condition, ShouldBeTrue)
		})

		Convey("It should pass with a valid event boolean", func() {
			condition := feedbag.ValidateConditional(".CreateIssueComment", p)
			So(condition, ShouldBeTrue)
		})

		Convey("It should not pass with a false boolean", func() {
			condition := feedbag.ValidateConditional(".OpenPullRequest", p)
			So(condition, ShouldBeFalse)
		})

		Convey("It should pass with a greater than condition", func() {
			// If number is greater than 100
			condition := feedbag.ValidateConditional("gt .Number 100", p)
			So(condition, ShouldBeTrue)
		})

		Convey("It should pass with multiple 'and' statements", func() {
			// If CreateIssueComment and number is less than 150
			condition := feedbag.ValidateConditional("and .CreateIssueComment (lt .Number 150)", p)
			So(condition, ShouldBeTrue)
		})

		Convey("It should pass with multiple 'or' statements", func() {
			condition := feedbag.ValidateConditional("or .CreateIssueComment .OpenPullRequest", p)
			So(condition, ShouldBeTrue)
		})

	})
}

func buildPayload(payloadFile string) (p feedbag.ActivityPayload) {
	file, _ := ioutil.ReadFile(payloadFile)
	_ = json.Unmarshal(file, &p)
	return p
}
