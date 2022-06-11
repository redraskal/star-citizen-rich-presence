package locations

import (
	"strings"
)

func Find(id string) Location {
	suggestions, _ := suggest.Lookup(strings.ToLower(id))
	if len(suggestions) == 0 {
		return Unknown()
	}
	i := suggestions[0].WordData["index"].(int)
	return locations[i]
}
