package eveapi

import (
	"crypto/sha1"
	"fmt"
	"io"
	"strconv"
	"time"
)

// Moved to ESI
// See https://github.com/antihax/swagger-esi-goclient

// Generate the killmail hash using source information.
func GenerateKillMailHash(victimID int64, attackerID int64, shipTypeID int64, killTime time.Time) string {
	v := strconv.FormatInt(victimID, 10)
	if victimID == 0 {
		v = "None"
	}
	a := strconv.FormatInt(attackerID, 10)
	if attackerID == 0 {
		a = "None"
	}
	s := strconv.FormatInt(shipTypeID, 10)

	t := strconv.FormatInt(convertTimeToWindow64(killTime), 10)

	h := sha1.New()
	io.WriteString(h, v)
	io.WriteString(h, a)
	io.WriteString(h, s)
	io.WriteString(h, t)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Convert to killmail time for hashing.
func convertTimeToWindow64(t time.Time) int64 {
	return t.Unix()*10000000 + 116444736000000000
}
