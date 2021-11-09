package main

type QueueData struct {
	QueueId string   `json:"queue"`
	Members []Agent  `json:"members"`
	Callers []string `json:"callers"`
}
