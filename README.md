# GoESI "Go Easy" API client for esi
[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/O5O33VK5S)


An OpenAPI for EVE Online ESI API

A module to allow access to CCP's EVE Online ESI API.
This module offers:

* Versioned Endpoints
* OAuth2 authentication to login.eveonline.com
* Handle many tokens, with different scopes.
* 100% ESI API coverage.
* context.Context passthrough (for httptrace, logging, etc).

## Installation

```go
    go get github.com/antihax/goesi
```

## New Client

```go
client := goesi.NewAPIClient(&http.Client, "MyApp (someone@somewhere.com dude on slack)")
```

One client should be created that will serve as an agent for all requests. This allows http2 multiplexing and keep-alive be used to optimize connections.
It is also good manners to provide a user-agent describing the point of use of the API, allowing CCP to contact you in case of emergencies.

Example:

```go
package main

import (
	"context"
	"fmt"

	"github.com/antihax/goesi"
)

func main() {
	// create ESI client
	client := goesi.NewAPIClient(nil, "name@example.com")
	// call Status endpoint
	status, _, err := client.ESI.StatusApi.GetStatus(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	// print current status
	fmt.Println("Players online: ", status.Players)
}
```

## Etiquette

* Create a descriptive user agent so CCP can contact you (preferably on devfleet slack).
* Obey Cache Timers.
* Obey error rate limits: https://developers.eveonline.com/blog/article/error-limiting-imminent

## Obeying the Cache Times

Caching is not implimented by the client and thus it is required to utilize
a caching http client. It is highly recommended to utilize a client capable
of caching the entire cluster of API clients.

An example using gregjones/httpcache and memcache:

```go
import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gregjones/httpcache"
	httpmemcache "github.com/gregjones/httpcache/memcache"
)

func main() {
	// Connect to the memcache server
	cache := memcache.New(MemcachedAddresses...)

	// Create a memcached http client for the CCP APIs.
	transport := httpcache.NewTransport(httpmemcache.NewWithClient(cache))
	transport.Transport = &http.Transport{Proxy: http.ProxyFromEnvironment}
	client = &http.Client{Transport: transport}

	// Get our API Client.
	eve := goesi.NewAPIClient(client, "My user agent, contact somewhere@nowhere")
}
```

## ETags

You should support using ETags if you are requesting data that is frequently not changed.
IF you are using httpcache, it supports etags already. If you are not using a cache
middleware, you will want to create your own middleware like this.

```go
package myetagpackage
type contextKey string

func (c contextKey) String() string {
	return "mylib " + string(c)
}

// ContextETag is the context to pass etags to the transport
var (
	ContextETag = contextKey("etag")
)

// Custom transport to chain into the HTTPClient to gather statistics.
type ETagTransport struct {
	Next *http.Transport
}

// RoundTrip wraps http.DefaultTransport.RoundTrip to provide stats and handle error rates.
func (t *ETagTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if etag, ok := req.Context().Value(ContextETag).(string); ok {
		req.Header.Set("if-none-match", etag)
	}

	// Run the request.
	return t.Next.RoundTrip(req)
}
```

This is then looped in the transport, and passed through a context like so:

```go
func main() {
	// Loop in our middleware
	client := &http.Client{Transport: &ETagTransport{Next: &http.Transport{}}}

	// Make a new client with the middleware
	esiClient := goesi.NewAPIClient(client, "MyApp (someone@somewhere.com dude on slack)")

	// Make a request with the context
	ctx := context.WithValue(context.Background(), myetagpackage.ContextETag, "etag goes here")
	regions, _, err := esiClient.UniverseApi.GetUniverseRegions(ctx, nil)
	if err != nil {
		return err
	}
}
```

## Authenticating

Register your application at https://developers.eveonline.com/ to get your secretKey, clientID, and scopes.

Obtaining tokens for client requires two HTTP handlers. One to generate and redirect
to the SSO URL, and one to receive the response.

It is mandatory to create a random state and compare this state on return to prevent token injection attacks on the application.

pseudocode example:

```go

func main() {
var err error
ctx := appContext.AppContext{}
ctx.ESI = goesi.NewAPIClient(httpClient, "My App, contact someone@nowhere")
ctx.SSOAuthenticator = goesi.NewSSOAuthenticator(httpClient, clientID, secretKey, scopes)
}

func eveSSO(c *appContext.AppContext, w http.ResponseWriter, r *http.Request,
	s *sessions.Session) (int, error) {

	// Generate a random state string
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	// Save the state on the session
	s.Values["state"] = state
	err := s.Save(r, w)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Generate the SSO URL with the state string
	url := c.SSOAuthenticator.AuthorizeURL(state, true)

	// Send the user to the URL
	http.Redirect(w, r, url, 302)
	return http.StatusMovedPermanently, nil
}

func eveSSOAnswer(c *appContext.AppContext, w http.ResponseWriter, r *http.Request,
	s *sessions.Session) (int, error) {

	// get our code and state
	code := r.FormValue("code")
	state := r.FormValue("state")

	// Verify the state matches our randomly generated string from earlier.
	if s.Values["state"] != state {
		return http.StatusInternalServerError, errors.New("Invalid State.")
	}

	// Exchange the code for an Access and Refresh token.
	token, err := c.SSOAuthenticator.TokenExchange(code)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Obtain a token source (automaticlly pulls refresh as needed)
	tokSrc, err := c.SSOAuthenticator.TokenSource(tok)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Assign an auth context to the calls
	auth := context.WithValue(context.TODO(), goesi.ContextOAuth2, tokSrc.Token)

	// Verify the client (returns clientID)
	v, err := c.SSOAuthenticator.Verify(auth)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Save the verification structure on the session for quick access.
	s.Values["character"] = v
	err = s.Save(r, w)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Redirect to the account page.
	http.Redirect(w, r, "/account", 302)
	return http.StatusMovedPermanently, nil
}
```

## Passing Tokens

OAuth2 tokens are passed to endpoints via contexts. Example:

```go
	ctx := context.WithValue(context.Background(), goesi.ContextOAuth2, ESIPublicToken)
	struc, response, err := client.V1.UniverseApi.GetUniverseStructuresStructureId(ctx, structureID, nil)
```

This is done here rather than at the client so you can use one client for many tokens, saving connections.

## Testing

If you would rather not rely on public ESI for testing, a mock ESI server is available for local and CI use.
Information here: https://github.com/antihax/mock-esi

## What about the other stuff?

If you need bleeding edge access, add the endpoint to the generator and rebuild this module.
Generator is here: https://github.com/antihax/swagger-esi-goclient

## Documentation for API Endpoints

[ESI Endpoints](./esi/README.md)

## Author

  antihax on #devfleet slack

## Credits

https://github.com/go-resty/resty (MIT license) Copyright Â© 2015-2016 Jeevanandam M (jeeva@myjeeva.com)
 - Uses modified setBody and detectContentType

https://github.com/gregjones/httpcache (MIT license) Copyright Â© 2012 Greg Jones (greg.jones@gmail.com)
  - Uses parseCacheControl and CacheExpires as a helper function


