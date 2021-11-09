package main

type Device struct {
	Extension string `json:"extension"`
	Ip        string `json:"ip"`
	Port      int    `json:"port"`
	Status    string `json:"status"`
}
