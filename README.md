# Go API client for esi

An OpenAPI for EVE Online ESI API

## Overview
A module to allow access to CCP's EVE Online ESI API.
This module offers:
* Versioned Endpoints
* Ability to pass OAuth2 Tokens (See [this module](https://github.com/antihax/eveapi) for authentication)
* Complete API coverage.

## Installation
```
    go get github.com/antihax/goesi
```

## New Client
```
  client, err := goesi.NewClient(*http.Client, userAgent string)
```

One client should be created that will serve as an agent for all requests. This allows http2 multiplexing and keep-alive be used to optimize connections.
It is also good manners to provide a user-agent describing the point of use of the API, allowing CCP to contact you in case of emergencies.

Example
```
  client, err := esi.NewClient(nil, "my esi client http://mysite.com contact <SomeDude> ingame")
  result, response, err := client.V#.Endpoint.Operation(requiredParam, map[string]interface{} { 
                                                                        "optionalParam1": "stringParam",
                                                                        "optionalParam2": 1234.56
                                                                    })
```

## Passing Tokens
OAuth2 tokens are passed to endpoings via contexts. Example:
```
	ctx := context.WithValue(context.TODO(), goesi.ContextOAuth2, ESIPublicToken)
	struc, response, err := c.ctx.ESI.V1.UniverseApi.GetUniverseStructuresStructureId(ctx, s, nil)
```

## Documentation for API Endpoints

[V1 Endpoints](./v1/README.md)
[V2 Endpoints](./v2/README.md)
[V3 Endpoints](./v3/README.md)
[V4 Endpoints](./v4/README.md)

## Author
  croakroach in game.
  antihax on #devfleet slack


## Credits
https://github.com/go-resty/resty (MIT license) Copyright Â© 2015-2016 Jeevanandam M (jeeva@myjeeva.com)
 - Uses modified setBody and detectContentType

https://github.com/gregjones/httpcache (MIT license) Copyright Â© 2012 Greg Jones (greg.jones@gmail.com)
  - Uses parseCacheControl and CacheExpires as a helper function


