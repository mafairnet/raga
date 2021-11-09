package main

type WsRequest struct {
	Type string `json:"type"`
	Value        string `json:"value"`
	Date      int    `json:"date"`
}