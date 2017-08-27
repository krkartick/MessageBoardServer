package main

import (
	"fmt"
	"net/http"
)

type Route struct {
	Name	string
	Method	string
	Pattern string
	FunName http.HandlerFunc
}

type Routes []Route

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"New Message Board with REST Interface ")
}

var routes = Routes{
	Route{"Root","GET","/",Index},
	Route{"GetTopic","GET","/Topics/{id}",returnOneTopic},
	Route{"PostTopic","POST","/Topics/{id}",updateOneTopic},
	Route{"DeleteTopic","DELETE","/Topics/{id}",deleteOneTopic},
}
