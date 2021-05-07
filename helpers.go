package goesi

import (
	"math"
	"regexp"
	"time"
)

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

func ParseTime(input int64) time.Time {
	maxd := time.Duration(math.MaxInt64).Truncate(100 * time.Nanosecond)
	maxdUnits := int64(maxd / 100)
	t := time.Date(1601, 1, 1, 0, 0, 0, 0, time.UTC)
	for input > maxdUnits {
		t = t.Add(maxd)
		input -= maxdUnits
	}
	if input != 0 {
		t = t.Add(time.Duration(input * 100))
	}
	return t
}
