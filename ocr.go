package main

import (
	"bufio"
	"strings"
)

type SessionInfo struct {
	LocationID string
	Location   string
	Details    string
}

func ParseSessionInfo(text string) SessionInfo {
	var s = SessionInfo{}
	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		line := scanner.Text()
		start := strings.IndexAny(line, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
		if start == -1 {
			continue
		}
		line = strings.ToLower(line[start:])
		println("DEBUG:", line)
		switch {
		case strings.HasPrefix(line, "current player location : "):
			s.LocationID = strings.TrimPrefix(line, "current player location : ")
			s.Location = GetLocationName(s.LocationID)
		}
	}
	setDetails(&s)
	return s
}

func setDetails(s *SessionInfo) {
	switch s.Location {
	case "":
		break
	case "Main Menu":
		s.Details = "In Main Menu"
	case "Space":
		s.Details = "In Space"
	case "Unknown":
		s.Details = "In Unknown Territory"
	default:
		s.Details = "At " + s.Location
	}
}
