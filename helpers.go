package goesi

import "regexp"

// https://community.eveonline.com/support/policies/naming-policy-en/
func ValidCharacterName(name string) bool {
	if len(name) > 37 {
		return false
	}
	if m, _ := regexp.MatchString("^[a-zA-Z0-9' -]+$", name); !m {
		return false
	}
	return true
}
