package feedbag

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"text/template"
	"time"

	"github.com/fogcreek/logging"
	"github.com/mojotech/feedbag/feedbag/tmpl"
)

type Activity struct {
	Id         int64           `json:"id"`
	TemplateId string          `json:"template_id"`
	Event      ActivityPayload `json:"event"`
	EventTime  time.Time       `json:"event_time"`
	TimeStamp
}

type ActivityList []Activity

type ByPriority struct{ tmpl.TemplateList }

func (b ByPriority) Less(i, j int) bool {
	return b.TemplateList[i].Priority > b.TemplateList[j].Priority
}

func (a *ActivityList) List() error {
	_, err := dbmap.Select(a, "SELECT * FROM ACTIVITIES ORDER BY CreatedAt DESC")
	return err
}

func (a *Activity) Create() error {
	// run the UpdateTime ethod on the user model
	a.UpdateTime()

	// run the DB insert function
	err := dbmap.Insert(a)
	return err
}

func ActivityParser(p ActivityPayloadList) []Activity {
	activities := ActivityList{}

	for _, payload := range p {
		isBreak := make(map[string]bool)

		// Sort templates by priority
		sort.Sort(ByPriority{Templates})

		for _, template := range Templates {
			activity := Activity{}

			//Set the event time on the activity object to the payloads event time
			activity.EventTime = payload.EventTime

			if isBreak[template.Event] {
				continue
			}

			if CheckEventType(template.Event, payload) && ValidateConditional(template.Condition, payload) {
				if template.Break {
					isBreak[template.Event] = true
				}

				logging.InfoWithTags([]string{"activity"}, "Adding activity for event ", activity.Event.EventType)

				//Create the activity for this template and append to list
				activity.TemplateId = template.Id
				activity.Event = payload
				activity.UpdateTime()

				err := activity.Create()
				if err != nil {
					logging.WarnWithTags([]string{"activity"}, "Failed to save activity", err.Error())
				}

				activities = append(activities, activity)

			}
		}
	}

	return activities
}

func CheckEventType(e string, p ActivityPayload) bool {
	if e == "*" {
		return true
	}
	return e == p.EventType
}

func ValidateConditional(c string, p ActivityPayload) bool {
	if c == "" {
		return true
	}

	buf := new(bytes.Buffer)

	tmpl, err := template.New("condition").Parse(fmt.Sprintf("{{%s}}", c))
	if err != nil {
		logging.ErrorWithTags([]string{"condition"}, "Failed to parse template condition", err.Error())
	}

	err = tmpl.Execute(buf, p)
	if err != nil {
		logging.ErrorWithTags([]string{"condition"}, "Failed to execute template condition", err.Error())
	}

	resultStr := buf.String()
	result, err := strconv.ParseBool(resultStr)
	if err != nil {
		logging.ErrorWithTags([]string{"condition"}, "Failed to parse bool from condition return", err.Error())
	}

	return result
}
