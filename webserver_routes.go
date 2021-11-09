package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Call",
		"GET",
		"/command/{command}",
		Command,
	},
	Route{
		"Call",
		"GET",
		"/queue/{action}",
		QueueAction,
	},
	Route{
		"Call",
		"GET",
		"/monitor/",
		Monitor,
	},
	Route{
		"Call",
		"GET",
		"/devices/",
		Devices,
	},
	Route{
		"Call",
		"GET",
		"/ws/",
		wsEndpoint,
	},
}
