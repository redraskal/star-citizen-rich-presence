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

func UpdateActivity(a client.Activity, s SessionInfo) error {
	if len(s.Location.IDs) == 0 {
		return nil
	}
	if s.Details == a.Details {
		return nil
	}
	if s.Location.Image != "" {
		a.LargeImage = s.Location.Image
		a.SmallImage = DefaultActivity.LargeImage
	} else {
		a.LargeImage = DefaultActivity.LargeImage
		a.SmallImage = ""
	}
	println("Activity:", s.Details)
	a.Details = s.Details
	return client.SetActivity(a)
}
