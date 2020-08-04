package main

type Queue struct {
	QueueId      int           `json:"queue_id"`
	Calls        int           `json:"calls"`
	Strategy     string        `json:"strategy"`
	HoldTime     int           `json:"hold_time"`
	TalkTime     int           `json:"talk_time"`
	Waiting      int           `json:"waiting"`
	Called       int           `json:"called"`
	Answered     int           `json:"answered"`
	ServiceLevel string        `json:"service_level"`
	Members      []QueueMember `json:"members"`
	Callers      []QueueCaller `json:"callers"`
}
