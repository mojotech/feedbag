package main

import (
	"bytes"
	"fmt"
	"strconv"
	"text/template"
)

type Activity struct {
	Template  string                 `json:"template"`
	Event     map[string]interface{} `json:"event"`
	EventTime string                 `json:"event_time"`
	TimeStamp
}

func ActivityParser(p []ActivityPayload) {
	for _, payload := range p {
		for _, template := range templates {
			if ParseConditional(template.Condition, payload) {
				fmt.Println("Condition True")
				//TODO: Add to db and send down the wire to the client
			}
		}
	}
}

func ParseConditional(c string, p ActivityPayload) bool {
	buf := new(bytes.Buffer)

	if c == "" {
		return false
	}

	tmpl, err := template.New("condition").Parse(fmt.Sprintf("{{%s}}", c))
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(buf, p)
	if err != nil {
		panic(err)
	}

	resultStr := buf.String()
	result, err := strconv.ParseBool(resultStr)
	if err != nil {
		panic(err)
	}

	return result
}
