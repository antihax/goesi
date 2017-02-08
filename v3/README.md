# Go API client for goesiv3

## Documentation for API Endpoints

All URIs are relative to *https://esi.tech.ccp.is/v3*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CalendarApi* | [**GetCharactersCharacterIdCalendarEventId**](docs/CalendarApi.md#getcharacterscharacteridcalendareventid) | **Get** /characters/{character_id}/calendar/{event_id}/ | Get an event
*CalendarApi* | [**PutCharactersCharacterIdCalendarEventId**](docs/CalendarApi.md#putcharacterscharacteridcalendareventid) | **Put** /characters/{character_id}/calendar/{event_id}/ | Respond to an event
*CharacterApi* | [**GetCharactersCharacterId**](docs/CharacterApi.md#getcharacterscharacterid) | **Get** /characters/{character_id}/ | Get character&#39;s public information
*CharacterApi* | [**PostCharactersCharacterIdCspa**](docs/CharacterApi.md#postcharacterscharacteridcspa) | **Post** /characters/{character_id}/cspa/ | Calculate a CSPA charge cost
*CorporationApi* | [**GetCorporationsCorporationId**](docs/CorporationApi.md#getcorporationscorporationid) | **Get** /corporations/{corporation_id}/ | Get corporation information
*MailApi* | [**GetCharactersCharacterIdMailLabels**](docs/MailApi.md#getcharacterscharacteridmaillabels) | **Get** /characters/{character_id}/mail/labels/ | Get mail labels and unread counts
*SearchApi* | [**GetCharactersCharacterIdSearch**](docs/SearchApi.md#getcharacterscharacteridsearch) | **Get** /characters/{character_id}/search/ | Search on a string
*SkillsApi* | [**GetCharactersCharacterIdSkills**](docs/SkillsApi.md#getcharacterscharacteridskills) | **Get** /characters/{character_id}/skills/ | Get character skills


## Documentation For Models

 - [GetCharactersCharacterIdCalendarEventIdForbidden](docs/GetCharactersCharacterIdCalendarEventIdForbidden.md)
 - [GetCharactersCharacterIdCalendarEventIdInternalServerError](docs/GetCharactersCharacterIdCalendarEventIdInternalServerError.md)
 - [GetCharactersCharacterIdCalendarEventIdOk](docs/GetCharactersCharacterIdCalendarEventIdOk.md)
 - [GetCharactersCharacterIdInternalServerError](docs/GetCharactersCharacterIdInternalServerError.md)
 - [GetCharactersCharacterIdMailLabelsForbidden](docs/GetCharactersCharacterIdMailLabelsForbidden.md)
 - [GetCharactersCharacterIdMailLabelsInternalServerError](docs/GetCharactersCharacterIdMailLabelsInternalServerError.md)
 - [GetCharactersCharacterIdMailLabelsLabel](docs/GetCharactersCharacterIdMailLabelsLabel.md)
 - [GetCharactersCharacterIdMailLabelsOk](docs/GetCharactersCharacterIdMailLabelsOk.md)
 - [GetCharactersCharacterIdOk](docs/GetCharactersCharacterIdOk.md)
 - [GetCharactersCharacterIdSearchForbidden](docs/GetCharactersCharacterIdSearchForbidden.md)
 - [GetCharactersCharacterIdSearchInternalServerError](docs/GetCharactersCharacterIdSearchInternalServerError.md)
 - [GetCharactersCharacterIdSearchOk](docs/GetCharactersCharacterIdSearchOk.md)
 - [GetCharactersCharacterIdSkillsForbidden](docs/GetCharactersCharacterIdSkillsForbidden.md)
 - [GetCharactersCharacterIdSkillsInternalServerError](docs/GetCharactersCharacterIdSkillsInternalServerError.md)
 - [GetCharactersCharacterIdSkillsOk](docs/GetCharactersCharacterIdSkillsOk.md)
 - [GetCharactersCharacterIdSkillsSkill](docs/GetCharactersCharacterIdSkillsSkill.md)
 - [GetCharactersCharacterIdUnprocessableEntity](docs/GetCharactersCharacterIdUnprocessableEntity.md)
 - [GetCorporationsCorporationIdInternalServerError](docs/GetCorporationsCorporationIdInternalServerError.md)
 - [GetCorporationsCorporationIdNotFound](docs/GetCorporationsCorporationIdNotFound.md)
 - [GetCorporationsCorporationIdOk](docs/GetCorporationsCorporationIdOk.md)
 - [PostCharactersCharacterIdCspaCharacters](docs/PostCharactersCharacterIdCspaCharacters.md)
 - [PostCharactersCharacterIdCspaCreated](docs/PostCharactersCharacterIdCspaCreated.md)
 - [PostCharactersCharacterIdCspaForbidden](docs/PostCharactersCharacterIdCspaForbidden.md)
 - [PostCharactersCharacterIdCspaInternalServerError](docs/PostCharactersCharacterIdCspaInternalServerError.md)
 - [PutCharactersCharacterIdCalendarEventIdForbidden](docs/PutCharactersCharacterIdCalendarEventIdForbidden.md)
 - [PutCharactersCharacterIdCalendarEventIdInternalServerError](docs/PutCharactersCharacterIdCalendarEventIdInternalServerError.md)
 - [PutCharactersCharacterIdCalendarEventIdResponse](docs/PutCharactersCharacterIdCalendarEventIdResponse.md)


## Documentation For Authorization


## evesso

- **Type**: OAuth
- **Flow**: implicit
- **Authorization URL**: https://login.eveonline.com/oauth/authorize
- **Scopes**: 
 - **esi-calendar.read_calendar_events.v1**: EVE SSO scope esi-calendar.read_calendar_events.v1
 - **esi-calendar.respond_calendar_events.v1**: EVE SSO scope esi-calendar.respond_calendar_events.v1
 - **esi-characters.read_contacts.v1**: EVE SSO scope esi-characters.read_contacts.v1
 - **esi-mail.read_mail.v1**: EVE SSO scope esi-mail.read_mail.v1
 - **esi-search.search_structures.v1**: EVE SSO scope esi-search.search_structures.v1
 - **esi-skills.read_skills.v1**: EVE SSO scope esi-skills.read_skills.v1

Example:
```
  auth, err = oauth2conf.TokenSource(http.Client, token) // Access and Refresh token structure

  client, err := goesiv3.NewClient(nil, "goesiv3 client http://mysite.com")
  ctx := context.WithValue(context.TODO(), goesiv3.ContextOAuth2, auth)
  result, response, err := client.Endpoint.AuthenticatedOperation(  ctx, 
                                                                    requiredParam, 
                                                                    map[string]interface{} { 
                                                                       "optionalParam1": "stringParam",
                                                                       "optionalParam2": 1234.56
                                                                    })
```


## Credits
https://github.com/go-resty/resty (MIT license) Copyright Â© 2015-2016 Jeevanandam M (jeeva@myjeeva.com)
 - Uses modified setBody and detectContentType

https://github.com/gregjones/httpcache (MIT license) Copyright Â© 2012 Greg Jones (greg.jones@gmail.com)
  - Uses parseCacheControl and CacheExpires as a helper function


