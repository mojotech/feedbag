package main

type Activity struct {
	Template  string                 `json:"template"`
	Event     map[string]interface{} `json:"event"`
	EventTime string                 `json:"event_time"`
	TimeStamp
}
