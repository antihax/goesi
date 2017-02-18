package eveapi

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type crestSimpleFrame struct {
	PageURL    string    // request URI for bookmarking purposes.
	CacheUntil time.Time // reponse cache time
}

type crestPagedFrame struct {
	crestSimpleFrame

	Next struct {
		HRef string
	}
	Previous struct {
		HRef string
	}
	TotalCount int
	PageCount  int
	Page       int
}

// getFrameInfo adds our own information to the responses.
func (c *crestSimpleFrame) getFrameInfo(r *http.Response) error {
	// Save the URL for bookmarking purposes.
	c.PageURL = r.Request.URL.String()

	var iMaxAge int
	var err error

	// Determine the cache duration.
	cacheControl := r.Header.Get("Cache-Control")
	if cacheControl != "" {
		maxAge := strings.Split(cacheControl, "=")[1]
		iMaxAge, err = strconv.Atoi(maxAge)
		if err != nil {
			return err
		}
	}

	// Get the response date.
	date, err := time.Parse(time.RFC1123, r.Header.Get("Date"))
	if err != nil {
		return err
	}

	// Add the cache duration
	c.CacheUntil = date.Add(time.Duration(iMaxAge) * time.Second)

	return nil
}

// getFrameInfo adds additional information on paged frames.
func (c *crestPagedFrame) getFrameInfo(url string, r *http.Response) error {
	// Apply the same procedure as the simple frame.
	if err := c.crestSimpleFrame.getFrameInfo(r); err != nil {
		return err
	}

	// Extract any page numbers.
	page, err := getPageNumberFromURL(url)
	if err != nil {
		return err
	}
	c.Page = page

	return nil
}

// Extract page numbers from the URL.
func getPageNumberFromURL(s string) (int, error) {
	u, err := url.Parse(s)
	if err != nil {
		return 0, err
	}

	m, err := url.ParseQuery(u.RawQuery)
	if m != nil {
		if m["page"] != nil {
			var i int
			i, err = strconv.Atoi(m["page"][0])
			if err != nil {
				return 0, err
			}
			return i, nil
		}
	}
	return 0, err
}
