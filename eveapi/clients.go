package eveapi

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
)

// EVEAPIClient for Public CREST and Public XML API queries.
type EVEAPIClient struct {
	httpClient *http.Client
	base       EveURI
	userAgent  string
}

// ErrorMessage format if a CREST query fails.
type ErrorMessage struct {
	Message string
}

// Executes a request generated with newRequest
func (c *EVEAPIClient) executeRequest(req *http.Request) (*http.Response, error) {
	return c.httpClient.Do(req)
}

// Creates a new http.Request for a public resource.
func (c *EVEAPIClient) newRequest(method, urlStr string, body interface{}, mediaType string) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if body != nil {
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, rel.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("Accept", BASE_API_VERSION)
	req.Header.Add("User-Agent", c.userAgent)

	return req, nil
}

// Calls a resource from the public XML API
func (c *EVEAPIClient) doXML(method, urlStr string, body interface{}, v interface{}) (*http.Response, error) {
	xmlThrottle.throttleRequest()  // Throttle XML requests
	connectionLimit.startRequest() // Limit concurrent requests
	defer connectionLimit.endRequest()

	req, err := c.newRequest(method, urlStr, body, "application/xml")
	if err != nil {
		return nil, err
	}

	res, err := c.executeRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, errors.New(res.Status)
	}

	if err := xml.NewDecoder(res.Body).Decode(v); err != nil {
		return nil, err
	}
	return res, nil
}

// Calls a resource from the public CREST
func (c *EVEAPIClient) doJSON(method, urlStr string, body interface{}, v interface{}, mediaType string, auth oauth2.TokenSource) (*http.Response, error) {
	anonThrottle.throttleRequest() // Throttle Anonymous CREST requests
	connectionLimit.startRequest() // Limit concurrent requests
	defer connectionLimit.endRequest()

	req, err := c.newRequest(method, urlStr, body, mediaType)
	if err != nil {
		return nil, err
	}

	if auth != nil {
		// We were able to grab an oauth2 token from the context
		var latestToken *oauth2.Token
		if latestToken, err = auth.Token(); err != nil {
			return nil, err
		}
		latestToken.SetAuthHeader(req)
	}

	res, err := c.executeRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		e := &ErrorMessage{}
		if err := json.NewDecoder(res.Body).Decode(e); err != nil {
			return nil, err
		}
		return nil, errors.New(e.Message)
	}
	if err := json.NewDecoder(res.Body).Decode(v); err != nil {
		return nil, err
	}

	return res, nil
}

// SetUI set the user agent string of the CREST and XML client.
// It is recommended to change this so that CCP can identify your app.
func (c *EVEAPIClient) SetUA(userAgent string) {
	c.userAgent = userAgent
}

// UseCustomURL allows the base URLs to be changed should the need arise
// for a third party proxy to be used.
func (c *EVEAPIClient) UseCustomURL(custom EveURI) {
	c.base = custom
}

// UseTestServer forces this client to use the test server URLs.
func (c *EVEAPIClient) UseTestServer(testServer bool) {
	if testServer == true {
		c.base = eveSisi
	} else {
		c.base = eveTQ
	}
}

// EVEAPIClient generates a new anonymous client.
// Caller must provide a caching http.Client that obeys all cacheUntil timers
// One Anonymous client per IP address or rate limits will be exceeded resulting in a ban.
func NewEVEAPIClient(client *http.Client, userAgent string) *EVEAPIClient {
	c := &EVEAPIClient{}
	c.base = eveTQ
	c.httpClient = client
	c.userAgent = userAgent
	return c
}

func (c *EVEAPIClient) GetCRESTURI() string {
	return c.base.CREST
}

func (c *EVEAPIClient) GetXMLURI() string {
	return c.base.XML
}

func (c *EVEAPIClient) GetLoginURI() string {
	return c.base.Login
}

func (c *EVEAPIClient) GetImageURI() string {
	return c.base.Images
}

func (c *EVEAPIClient) GetManagementURI() string {
	return c.base.AppManagement
}
