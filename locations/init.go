package locations

import (
	_ "embed"
	"encoding/json"
	"strings"

	"github.com/agusnavce/ta"
	"github.com/agusnavce/ta/utils"
)

type Location struct {
	IDs    []string `json:"ids"`
	Name   string   `json:"name"`
	Prefix string   `json:"prefix"`
	Image  string   `json:"image"`
}

var (
	//go:embed locations.json
	file      []byte
	locations = make([]Location, 0)
	index     = make(map[string]int)
	suggest   = ta.NewSpellModel()
)

func init() {
	println("Loading locations...")
	if err := json.Unmarshal(file, &locations); err != nil {
		panic(err)
	}
	for i, location := range locations {
		for _, id := range location.IDs {
			id = strings.ToLower(id)
			index[id] = i
			suggest.AddEntry(utils.Entry{
				Frequency: 1,
				Word:      id,
				WordData: utils.WordData{
					"index": i,
				},
			})
		}
		if location.Prefix == "" {
			locations[i].Prefix = "At"
		}
	}
	println("Indexed locations.")
}

func Unknown() Location {
	return Location{
		IDs: []string{
			"unknown",
		},
		Name:   "Radar",
		Prefix: "Off",
		Image:  "offradar",
	}
}
