package goesi

import "regexp"

// https://community.eveonline.com/support/policies/naming-policy-en/
func ValidCharacterName(name string) bool {
	if len(name) > 37 {
		return false
	}
	if m, _ := regexp.MatchString("^[\\p{L}\\d' -]+$", name); !m {
		return false
	}
	return true
}

func FactionNameToID(faction string) int32 {
	switch faction {
	case "Caldari":
		return 500001
	case "Minmatar":
		return 500002
	case "Amarr":
		return 500003
	case "Gallente":
		return 500004
	}
	return 0
}
