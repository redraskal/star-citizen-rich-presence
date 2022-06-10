package main

import (
	"time"

	"github.com/hugolgst/rich-go/client"
)

var (
	DefaultActivity = client.Activity{
		State:      "No Active Mission",
		Details:    "Starting Game",
		LargeImage: "logo",
		Timestamps: &client.Timestamps{},
	}
)

func UpdateStartTimestamp(a *client.Activity) {
	time := time.Now()
	a.Timestamps = &client.Timestamps{
		Start: &time,
	}
}

func UpdateActivity(a client.Activity, s SessionInfo) {
	if s.LocationID == "" {
		return
	}
	if s.Details == a.Details {
		return
	}
	println("Activity:", s.Details)
	a.Details = s.Details
	client.SetActivity(a)
}
