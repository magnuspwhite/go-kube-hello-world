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
		"EmailIndex",
		"POST",
		"/email",
		EmailIndex,
	},
	Route{
		"HealthcheckIndex",
		"GET",
		"/healthcheck",
		HealthcheckIndex,
	},
}
