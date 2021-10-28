package goesi

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"context"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/oauth2"
)

// SSOAuthenticator provides interfacing to the EVE SSO. NewSSOAuthenticator is used to create
// this structure.

// [TODO] lose this mutex and allow scopes to change without conflict.
type SSOAuthenticator struct {
	httpClient *http.Client
	// Hide this...
	oauthConfig *oauth2.Config
	scopeLock   sync.Mutex
}

// EVESSOClaims structure for JWT Claims
type EVESSOClaims struct {
	Name     string   `json:"name,omitempty"`
	Owner    string   `json:"owner,omitempty"`
	Scopes   []string `json:"scp,omitempty"`
	ClientID string   `json:"azp,omitempty"`
	Tenant   string   `json:"tenant,omitempty"`
	Tier     string   `json:"tier,omitempty"`
	Region   string   `json:"region,omitempty"`
	jwt.RegisteredClaims
}

// NewSSOAuthenticator create a new EVE SSO Authenticator.
// Requires your application clientID, clientSecret, and redirectURL.
// RedirectURL must match exactly to what you registered with CCP.
func NewSSOAuthenticator(client *http.Client, clientID string, clientSecret string, redirectURL string, scopes []string) *SSOAuthenticator {
	return newSSOAuthenticator(
		client,
		clientID,
		clientSecret,
		redirectURL,
		scopes,
		oauth2.Endpoint{
			AuthURL:  "https://login.eveonline.com/oauth/authorize",
			TokenURL: "https://login.eveonline.com/oauth/token",
		},
	)
}

// NewSSOAuthenticatorV2 create a new EVE SSO Authenticator with the v2 urls.
// Requires your application clientID, clientSecret, and redirectURL.
// RedirectURL must match exactly to what you registered with CCP.
func NewSSOAuthenticatorV2(client *http.Client, clientID string, clientSecret string, redirectURL string, scopes []string) *SSOAuthenticator {
	return newSSOAuthenticator(
		client,
		clientID,
		clientSecret,
		redirectURL,
		scopes,
		oauth2.Endpoint{
			AuthURL:  "https://login.eveonline.com/v2/oauth/authorize",
			TokenURL: "https://login.eveonline.com/v2/oauth/token",
		},
	)
}

func newSSOAuthenticator(client *http.Client, clientID string, clientSecret string, redirectURL string, scopes []string, endpoint oauth2.Endpoint) *SSOAuthenticator {
	if client == nil {
		return nil
	}

	c := &SSOAuthenticator{}
	c.httpClient = client
	c.oauthConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     endpoint,
		Scopes:       scopes,
		RedirectURL:  redirectURL,
	}

	return c
}

// ChangeAuthURL changes the oauth2 configuration url for authentication
func (c *SSOAuthenticator) ChangeAuthURL(url string) {
	c.oauthConfig.Endpoint.AuthURL = url
}

// ChangeTokenURL changes the oauth2 configuration url for token
func (c *SSOAuthenticator) ChangeTokenURL(url string) {
	c.oauthConfig.Endpoint.TokenURL = url
}

// AuthorizeURL returns a url for an end user to authenticate with EVE SSO
// and return success to the redirectURL.
// It is important to create a significatly unique state for this request
// and verify the state matches when returned to the redirectURL.
func (c *SSOAuthenticator) AuthorizeURL(state string, onlineAccess bool, scopes []string) string {
	var url string

	// Generate the URL
	if onlineAccess == true {
		url = c.oauthConfig.AuthCodeURL(state, oauth2.AccessTypeOnline, oauth2.SetAuthURLParam("scope", strings.Join(scopes, " ")))
	} else {
		url = c.oauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline, oauth2.SetAuthURLParam("scope", strings.Join(scopes, " ")))
	}

	return url
}

// TokenRevoke revokes a refresh token
func (c *SSOAuthenticator) TokenRevoke(refreshToken string) error {
	v := url.Values{
		"token_type_hint": {"refresh_token"},
		"token":           {refreshToken},
	}
	req, err := http.NewRequest("POST", "https://login.eveonline.com/oauth/revoke", strings.NewReader(v.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	req.Header.Set("Authorization",
		"Basic "+base64.URLEncoding.EncodeToString(
			[]byte(c.oauthConfig.ClientID+":"+c.oauthConfig.ClientSecret),
		))

	r, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1<<20))
	if err != nil {
		return fmt.Errorf("oauth2: cannot revoke token: %v", err)
	}
	if code := r.StatusCode; code < 200 || code > 299 {
		return &oauth2.RetrieveError{
			Response: r,
			Body:     body,
		}
	}
	return nil
}

// TokenExchange exchanges the code returned to the redirectURL with
// the CREST server to an access token. A caching client must be passed.
// This client MUST cache per CCP guidelines or face banning.
func (c *SSOAuthenticator) TokenExchange(code string) (*oauth2.Token, error) {
	tok, err := c.oauthConfig.Exchange(createContext(c.httpClient), code)
	if err != nil {
		return nil, err
	}
	return tok, nil
}

// TokenSource creates a refreshable token that can be passed to ESI functions
func (c *SSOAuthenticator) TokenSource(token *oauth2.Token) oauth2.TokenSource {
	return c.oauthConfig.TokenSource(createContext(c.httpClient), token)
}

type VerifyResponse struct {
	CharacterID        int32
	CharacterName      string
	ExpiresOn          string
	Scopes             string
	TokenType          string
	CharacterOwnerHash string
}

// Verify the client and collect user information.
func (c *SSOAuthenticator) Verify(auth oauth2.TokenSource) (*VerifyResponse, error) {
	v := &VerifyResponse{}
	_, err := c.doJSON("GET", "https://login.eveonline.com/oauth/verify", nil, v, "application/json;", auth)

	if err != nil {
		return nil, err
	}
	return v, nil
}

// Creates a new http.Request for a public resource.
func (c *SSOAuthenticator) newRequest(method, urlStr string, body interface{}, mediaType string) (*http.Request, error) {
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

	return req, nil
}

// Calls a resource from the public CREST
func (c *SSOAuthenticator) doJSON(method, urlStr string, body interface{}, v interface{}, mediaType string, auth oauth2.TokenSource) (*http.Response, error) {

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
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(string(buf))
	}
	if err := json.Unmarshal([]byte(buf), v); err != nil {
		return nil, err
	}

	return res, nil
}

// Executes a request generated with newRequest
func (c *SSOAuthenticator) executeRequest(req *http.Request) (*http.Response, error) {
	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusOK ||
		res.StatusCode == http.StatusCreated {
		return res, nil
	}
	return res, errors.New(res.Status)
}

// Add custom clients to the context.
func createContext(httpClient *http.Client) context.Context {
	parent := oauth2.NoContext
	ctx := context.WithValue(parent, oauth2.HTTPClient, httpClient)
	return ctx
}

// TokenToJSON helper function to convert a token to a storable format.
func TokenToJSON(token *oauth2.Token) (string, error) {
	if d, err := json.Marshal(token); err != nil {
		return "", err
	} else {
		return string(d), nil
	}
}

// TokenFromJSON helper function to convert stored JSON to a token.
func TokenFromJSON(jsonStr string) (*oauth2.Token, error) {
	var token oauth2.Token
	if err := json.Unmarshal([]byte(jsonStr), &token); err != nil {
		return nil, err
	}
	return &token, nil
}
