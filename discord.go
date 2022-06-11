package main

import (
	"time"

	"github.com/hugolgst/rich-go/client"
	"github.com/redraskal/star-citizen-rich-presence/rsi"
	"github.com/redraskal/star-citizen-rich-presence/utils"
)

var (
	DefaultActivity = client.Activity{
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

func UpdateProfile(a *client.Activity) error {
	username, err := rsi.Username()
	if err != nil {
		return err
	}
	println("Username:", username)
	a.Buttons = []*client.Button{
		{
			Label: "Profile",
			Url:   rsi.ProfileEndpoint + username,
		},
	}
	return nil
}

func UpdateActivity(a client.Activity, s utils.SessionInfo) error {
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
