package main

import (
	"bufio"
	"strings"

	"github.com/redraskal/starcitizen/locations"
)

type SessionInfo struct {
	Location locations.Location
	Details  string
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
			id := strings.TrimPrefix(line, "current player location : ")
			s.Location = locations.Find(id)
			s.Details = s.Location.Prefix + " " + s.Location.Name
		}
	}
	return s
}
