package main

import "strings"

func GetLocationName(id string) string {
	switch strings.ToLower(id) {
	case "invalid location 1d":
		return "Main Menu"
	case "stantonstar":
		return "Space"
	case "grindex":
		return "GrimHEX"
	case "grimhex":
		return "GrimHEX"
	default:
		return "Unknown"
	}
}
