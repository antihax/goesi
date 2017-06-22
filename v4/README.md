# Go API client for goesiv4

## Documentation for API Endpoints

All URIs are relative to *https://esi.tech.ccp.is/v4*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CharacterApi* | [**GetCharactersCharacterId**](docs/CharacterApi.md#getcharacterscharacterid) | **Get** /characters/{character_id}/ | Get character&#39;s public information
*SkillsApi* | [**GetCharactersCharacterIdSkills**](docs/SkillsApi.md#getcharacterscharacteridskills) | **Get** /characters/{character_id}/skills/ | Get character skills


## Documentation For Models

 - [Forbidden](docs/Forbidden.md)
 - [GetCharactersCharacterIdNotFound](docs/GetCharactersCharacterIdNotFound.md)
 - [GetCharactersCharacterIdOk](docs/GetCharactersCharacterIdOk.md)
 - [GetCharactersCharacterIdSkillsOk](docs/GetCharactersCharacterIdSkillsOk.md)
 - [GetCharactersCharacterIdSkillsSkill](docs/GetCharactersCharacterIdSkillsSkill.md)
 - [InternalServerError](docs/InternalServerError.md)


## Documentation For Authorization


## evesso

- **Type**: OAuth
- **Flow**: implicit
- **Authorization URL**: https://login.eveonline.com/oauth/authorize
- **Scopes**: 
 - **esi-skills.read_skills.v1**: EVE SSO scope esi-skills.read_skills.v1

Example:
```
  auth, err = oauth2conf.TokenSource(http.Client, token) // Access and Refresh token structure

  client, err := goesiv4.NewClient(nil, "goesiv4 client http://mysite.com")
  ctx := context.WithValue(context.TODO(), goesiv4.ContextOAuth2, auth)
  result, response, err := client.Endpoint.AuthenticatedOperation(  ctx, 
                                                                    requiredParam, 
                                                                    map[string]interface{} { 
                                                                       "optionalParam1": "stringParam",
                                                                       "optionalParam2": 1234.56
                                                                    })
```


## Credits
https://github.com/go-resty/resty (MIT license) Copyright © 2015-2016 Jeevanandam M (jeeva@myjeeva.com)
 - Uses modified setBody and detectContentType

https://github.com/gregjones/httpcache (MIT license) Copyright © 2012 Greg Jones (greg.jones@gmail.com)
  - Uses parseCacheControl and CacheExpires as a helper function


