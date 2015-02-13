package tmpl_test

import (
	"testing"

	"github.com/mojotech/feedbag/feedbag/tmpl"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ParseTemplatesDir(t *testing.T) {
	testDir := "./test_templates"
	templates, err := tmpl.ParseDir(testDir)

	Convey("Given the test template directory", t, func() {
		So(err, ShouldBeNil)
		So(len(templates), ShouldEqual, 4)

		var (
			newPrTmpl         tmpl.Template
			missingNameTmpl   tmpl.Template
			newCommentTmpl    tmpl.Template
			longerMarkersTmpl tmpl.Template
		)

		for _, tmpl := range templates {
			switch tmpl.Id {
			case "new_comment":
				newCommentTmpl = *tmpl
			case "new_pr":
				newPrTmpl = *tmpl
			case "missing_name":
				missingNameTmpl = *tmpl
			case "longer_markers":
				longerMarkersTmpl = *tmpl
			}
		}

		Convey("The New Comment template should exist", func() {
			t := newCommentTmpl
			So(t, ShouldNotBeNil)
			So(t.Id, ShouldEqual, "new_comment")
			So(t.Name, ShouldEqual, "New Comment")
			So(t.Event, ShouldEqual, "comment_pull_request")
		})

		Convey("The New PR template should exist", func() {
			t := newPrTmpl
			So(t, ShouldNotBeNil)
			So(t.Id, ShouldEqual, "new_pr")
			So(t.Name, ShouldEqual, "New Pull Request")
			So(t.Event, ShouldEqual, "new_pull_request")
		})

		Convey("The Missing Name template should exist", func() {
			t := missingNameTmpl
			So(t, ShouldNotBeNil)
			So(t.Id, ShouldEqual, "missing_name")
			So(t.Name, ShouldEqual, "")
			So(t.Event, ShouldEqual, "comment_pull_request")
		})

		Convey("The Longer Than Default Markers template should exist", func() {
			t := longerMarkersTmpl
			So(t, ShouldNotBeNil)
			So(t.Id, ShouldEqual, "longer_markers")
			So(t.Name, ShouldEqual, "Longer Than Default Markers")
			So(t.Event, ShouldEqual, "comment_pull_request")
		})

	})
}
