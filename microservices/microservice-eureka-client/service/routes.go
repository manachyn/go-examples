package service

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
		"Info",
		"GET",
		"/info",
		Info,
	},
	Route{
		"Health",
		"POST",
		"/health",
		Health,
	},
	Route{
		"VendorShow",
		"GET",
		"/vendors/{productId}",
		VendorShow,
	},
}
