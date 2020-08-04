package main

type QueueCaller struct {
	MemberId   int    `json:"member_id"`
	Penalty    int    `json:"penalty"`
	RingInUse  string `json:"ring_in_use"`
	MemberType string `json:"member_type"`
	Status     string `json:"status"`
	Calls      int    `json:"calls"`
	LastCall   int    `json:"last_call"`
}
