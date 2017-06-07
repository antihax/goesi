# Go API client for goesiv2

## Documentation for API Endpoints

All URIs are relative to *https://esi.tech.ccp.is/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AllianceApi* | [**GetAlliancesAllianceId**](docs/AllianceApi.md#getalliancesallianceid) | **Get** /alliances/{alliance_id}/ | Get alliance information
*CalendarApi* | [**GetCharactersCharacterIdCalendarEventId**](docs/CalendarApi.md#getcharacterscharacteridcalendareventid) | **Get** /characters/{character_id}/calendar/{event_id}/ | Get an event
*CalendarApi* | [**PutCharactersCharacterIdCalendarEventId**](docs/CalendarApi.md#putcharacterscharacteridcalendareventid) | **Put** /characters/{character_id}/calendar/{event_id}/ | Respond to an event
*CharacterApi* | [**GetCharactersCharacterIdPortrait**](docs/CharacterApi.md#getcharacterscharacteridportrait) | **Get** /characters/{character_id}/portrait/ | Get character portraits
*ClonesApi* | [**GetCharactersCharacterIdClones**](docs/ClonesApi.md#getcharacterscharacteridclones) | **Get** /characters/{character_id}/clones/ | Get clones
*ContactsApi* | [**DeleteCharactersCharacterIdContacts**](docs/ContactsApi.md#deletecharacterscharacteridcontacts) | **Delete** /characters/{character_id}/contacts/ | Delete contacts
*CorporationApi* | [**GetCorporationsCorporationId**](docs/CorporationApi.md#getcorporationscorporationid) | **Get** /corporations/{corporation_id}/ | Get corporation information
*CorporationApi* | [**GetCorporationsCorporationIdMembers**](docs/CorporationApi.md#getcorporationscorporationidmembers) | **Get** /corporations/{corporation_id}/members/ | Get corporation members
*DogmaApi* | [**GetDogmaEffectsEffectId**](docs/DogmaApi.md#getdogmaeffectseffectid) | **Get** /dogma/effects/{effect_id}/ | Get effect information
*LocationApi* | [**GetCharactersCharacterIdOnline**](docs/LocationApi.md#getcharacterscharacteridonline) | **Get** /characters/{character_id}/online/ | Get character online
*MailApi* | [**GetCharactersCharacterIdMailLabels**](docs/MailApi.md#getcharacterscharacteridmaillabels) | **Get** /characters/{character_id}/mail/labels/ | Get mail labels
*MailApi* | [**PostCharactersCharacterIdMailLabels**](docs/MailApi.md#postcharacterscharacteridmaillabels) | **Post** /characters/{character_id}/mail/labels/ | Create a mail label
*PlanetaryInteractionApi* | [**GetCharactersCharacterIdPlanetsPlanetId**](docs/PlanetaryInteractionApi.md#getcharacterscharacteridplanetsplanetid) | **Get** /characters/{character_id}/planets/{planet_id}/ | Get colony layout
*SearchApi* | [**GetCharactersCharacterIdSearch**](docs/SearchApi.md#getcharacterscharacteridsearch) | **Get** /characters/{character_id}/search/ | Search on a string
*SearchApi* | [**GetSearch**](docs/SearchApi.md#getsearch) | **Get** /search/ | Search on a string
*SkillsApi* | [**GetCharactersCharacterIdSkillqueue**](docs/SkillsApi.md#getcharacterscharacteridskillqueue) | **Get** /characters/{character_id}/skillqueue/ | Get character&#39;s skill queue
*SkillsApi* | [**GetCharactersCharacterIdSkills**](docs/SkillsApi.md#getcharacterscharacteridskills) | **Get** /characters/{character_id}/skills/ | Get character skills
*UniverseApi* | [**GetUniverseStationsStationId**](docs/UniverseApi.md#getuniversestationsstationid) | **Get** /universe/stations/{station_id}/ | Get station information
*UniverseApi* | [**GetUniverseSystemsSystemId**](docs/UniverseApi.md#getuniversesystemssystemid) | **Get** /universe/systems/{system_id}/ | Get solar system information
*UniverseApi* | [**GetUniverseTypesTypeId**](docs/UniverseApi.md#getuniversetypestypeid) | **Get** /universe/types/{type_id}/ | Get type information
*UniverseApi* | [**PostUniverseNames**](docs/UniverseApi.md#postuniversenames) | **Post** /universe/names/ | Get names and categories for a set of ID&#39;s
*UserInterfaceApi* | [**PostUiAutopilotWaypoint**](docs/UserInterfaceApi.md#postuiautopilotwaypoint) | **Post** /ui/autopilot/waypoint/ | Set Autopilot Waypoint


## Documentation For Models

 - [DeleteCharactersCharacterIdContactsForbidden](docs/DeleteCharactersCharacterIdContactsForbidden.md)
 - [DeleteCharactersCharacterIdContactsInternalServerError](docs/DeleteCharactersCharacterIdContactsInternalServerError.md)
 - [GetAlliancesAllianceIdInternalServerError](docs/GetAlliancesAllianceIdInternalServerError.md)
 - [GetAlliancesAllianceIdNotFound](docs/GetAlliancesAllianceIdNotFound.md)
 - [GetAlliancesAllianceIdOk](docs/GetAlliancesAllianceIdOk.md)
 - [GetCharactersCharacterIdCalendarEventIdForbidden](docs/GetCharactersCharacterIdCalendarEventIdForbidden.md)
 - [GetCharactersCharacterIdCalendarEventIdInternalServerError](docs/GetCharactersCharacterIdCalendarEventIdInternalServerError.md)
 - [GetCharactersCharacterIdCalendarEventIdOk](docs/GetCharactersCharacterIdCalendarEventIdOk.md)
 - [GetCharactersCharacterIdClonesForbidden](docs/GetCharactersCharacterIdClonesForbidden.md)
 - [GetCharactersCharacterIdClonesHomeLocation](docs/GetCharactersCharacterIdClonesHomeLocation.md)
 - [GetCharactersCharacterIdClonesInternalServerError](docs/GetCharactersCharacterIdClonesInternalServerError.md)
 - [GetCharactersCharacterIdClonesJumpClone](docs/GetCharactersCharacterIdClonesJumpClone.md)
 - [GetCharactersCharacterIdClonesOk](docs/GetCharactersCharacterIdClonesOk.md)
 - [GetCharactersCharacterIdMailLabels200Ok](docs/GetCharactersCharacterIdMailLabels200Ok.md)
 - [GetCharactersCharacterIdMailLabelsForbidden](docs/GetCharactersCharacterIdMailLabelsForbidden.md)
 - [GetCharactersCharacterIdMailLabelsInternalServerError](docs/GetCharactersCharacterIdMailLabelsInternalServerError.md)
 - [GetCharactersCharacterIdOnlineForbidden](docs/GetCharactersCharacterIdOnlineForbidden.md)
 - [GetCharactersCharacterIdOnlineInternalServerError](docs/GetCharactersCharacterIdOnlineInternalServerError.md)
 - [GetCharactersCharacterIdOnlineOk](docs/GetCharactersCharacterIdOnlineOk.md)
 - [GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails](docs/GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails.md)
 - [GetCharactersCharacterIdPlanetsPlanetIdFactoryDetails](docs/GetCharactersCharacterIdPlanetsPlanetIdFactoryDetails.md)
 - [GetCharactersCharacterIdPlanetsPlanetIdForbidden](docs/GetCharactersCharacterIdPlanetsPlanetIdForbidden.md)
 - [GetCharactersCharacterIdPlanetsPlanetIdHead](docs/GetCharactersCharacterIdPlanetsPlanetIdHead.md)
 - [GetCharactersCharacterIdPlanetsPlanetIdInternalServerError](docs/GetCharactersCharacterIdPlanetsPlanetIdInternalServerError.md)
 - [GetCharactersCharacterIdPlanetsPlanetIdLink](docs/GetCharactersCharacterIdPlanetsPlanetIdLink.md)
 - [GetCharactersCharacterIdPlanetsPlanetIdNotFound](docs/GetCharactersCharacterIdPlanetsPlanetIdNotFound.md)
 - [GetCharactersCharacterIdPlanetsPlanetIdOk](docs/GetCharactersCharacterIdPlanetsPlanetIdOk.md)
 - [GetCharactersCharacterIdPlanetsPlanetIdPin](docs/GetCharactersCharacterIdPlanetsPlanetIdPin.md)
 - [GetCharactersCharacterIdPlanetsPlanetIdRoute](docs/GetCharactersCharacterIdPlanetsPlanetIdRoute.md)
 - [GetCharactersCharacterIdPlanetsPlanetIdWaypoint](docs/GetCharactersCharacterIdPlanetsPlanetIdWaypoint.md)
 - [GetCharactersCharacterIdPortraitInternalServerError](docs/GetCharactersCharacterIdPortraitInternalServerError.md)
 - [GetCharactersCharacterIdPortraitNotFound](docs/GetCharactersCharacterIdPortraitNotFound.md)
 - [GetCharactersCharacterIdPortraitOk](docs/GetCharactersCharacterIdPortraitOk.md)
 - [GetCharactersCharacterIdSearchForbidden](docs/GetCharactersCharacterIdSearchForbidden.md)
 - [GetCharactersCharacterIdSearchInternalServerError](docs/GetCharactersCharacterIdSearchInternalServerError.md)
 - [GetCharactersCharacterIdSearchOk](docs/GetCharactersCharacterIdSearchOk.md)
 - [GetCharactersCharacterIdSkillqueue200Ok](docs/GetCharactersCharacterIdSkillqueue200Ok.md)
 - [GetCharactersCharacterIdSkillqueueForbidden](docs/GetCharactersCharacterIdSkillqueueForbidden.md)
 - [GetCharactersCharacterIdSkillqueueInternalServerError](docs/GetCharactersCharacterIdSkillqueueInternalServerError.md)
 - [GetCharactersCharacterIdSkills200Ok](docs/GetCharactersCharacterIdSkills200Ok.md)
 - [GetCharactersCharacterIdSkillsForbidden](docs/GetCharactersCharacterIdSkillsForbidden.md)
 - [GetCharactersCharacterIdSkillsInternalServerError](docs/GetCharactersCharacterIdSkillsInternalServerError.md)
 - [GetCorporationsCorporationIdInternalServerError](docs/GetCorporationsCorporationIdInternalServerError.md)
 - [GetCorporationsCorporationIdMembers200Ok](docs/GetCorporationsCorporationIdMembers200Ok.md)
 - [GetCorporationsCorporationIdMembersForbidden](docs/GetCorporationsCorporationIdMembersForbidden.md)
 - [GetCorporationsCorporationIdMembersInternalServerError](docs/GetCorporationsCorporationIdMembersInternalServerError.md)
 - [GetCorporationsCorporationIdNotFound](docs/GetCorporationsCorporationIdNotFound.md)
 - [GetCorporationsCorporationIdOk](docs/GetCorporationsCorporationIdOk.md)
 - [GetDogmaEffectsEffectIdInternalServerError](docs/GetDogmaEffectsEffectIdInternalServerError.md)
 - [GetDogmaEffectsEffectIdModifier](docs/GetDogmaEffectsEffectIdModifier.md)
 - [GetDogmaEffectsEffectIdNotFound](docs/GetDogmaEffectsEffectIdNotFound.md)
 - [GetDogmaEffectsEffectIdOk](docs/GetDogmaEffectsEffectIdOk.md)
 - [GetSearchInternalServerError](docs/GetSearchInternalServerError.md)
 - [GetSearchOk](docs/GetSearchOk.md)
 - [GetUniverseStationsStationIdInternalServerError](docs/GetUniverseStationsStationIdInternalServerError.md)
 - [GetUniverseStationsStationIdNotFound](docs/GetUniverseStationsStationIdNotFound.md)
 - [GetUniverseStationsStationIdOk](docs/GetUniverseStationsStationIdOk.md)
 - [GetUniverseStationsStationIdPosition](docs/GetUniverseStationsStationIdPosition.md)
 - [GetUniverseSystemsSystemIdInternalServerError](docs/GetUniverseSystemsSystemIdInternalServerError.md)
 - [GetUniverseSystemsSystemIdNotFound](docs/GetUniverseSystemsSystemIdNotFound.md)
 - [GetUniverseSystemsSystemIdOk](docs/GetUniverseSystemsSystemIdOk.md)
 - [GetUniverseSystemsSystemIdPlanet](docs/GetUniverseSystemsSystemIdPlanet.md)
 - [GetUniverseSystemsSystemIdPosition](docs/GetUniverseSystemsSystemIdPosition.md)
 - [GetUniverseTypesTypeIdDogmaAttribute](docs/GetUniverseTypesTypeIdDogmaAttribute.md)
 - [GetUniverseTypesTypeIdDogmaEffect](docs/GetUniverseTypesTypeIdDogmaEffect.md)
 - [GetUniverseTypesTypeIdInternalServerError](docs/GetUniverseTypesTypeIdInternalServerError.md)
 - [GetUniverseTypesTypeIdNotFound](docs/GetUniverseTypesTypeIdNotFound.md)
 - [GetUniverseTypesTypeIdOk](docs/GetUniverseTypesTypeIdOk.md)
 - [PostCharactersCharacterIdMailLabelsForbidden](docs/PostCharactersCharacterIdMailLabelsForbidden.md)
 - [PostCharactersCharacterIdMailLabelsInternalServerError](docs/PostCharactersCharacterIdMailLabelsInternalServerError.md)
 - [PostCharactersCharacterIdMailLabelsLabel](docs/PostCharactersCharacterIdMailLabelsLabel.md)
 - [PostUiAutopilotWaypointForbidden](docs/PostUiAutopilotWaypointForbidden.md)
 - [PostUiAutopilotWaypointInternalServerError](docs/PostUiAutopilotWaypointInternalServerError.md)
 - [PostUniverseNames200Ok](docs/PostUniverseNames200Ok.md)
 - [PostUniverseNamesInternalServerError](docs/PostUniverseNamesInternalServerError.md)
 - [PostUniverseNamesNotFound](docs/PostUniverseNamesNotFound.md)
 - [PutCharactersCharacterIdCalendarEventIdForbidden](docs/PutCharactersCharacterIdCalendarEventIdForbidden.md)
 - [PutCharactersCharacterIdCalendarEventIdInternalServerError](docs/PutCharactersCharacterIdCalendarEventIdInternalServerError.md)


## Documentation For Authorization


## evesso

- **Type**: OAuth
- **Flow**: implicit
- **Authorization URL**: https://login.eveonline.com/oauth/authorize
- **Scopes**: 
 - **esi-calendar.read_calendar_events.v1**: EVE SSO scope esi-calendar.read_calendar_events.v1
 - **esi-calendar.respond_calendar_events.v1**: EVE SSO scope esi-calendar.respond_calendar_events.v1
 - **esi-characters.write_contacts.v1**: EVE SSO scope esi-characters.write_contacts.v1
 - **esi-clones.read_clones.v1**: EVE SSO scope esi-clones.read_clones.v1
 - **esi-corporations.read_corporation_membership.v1**: EVE SSO scope esi-corporations.read_corporation_membership.v1
 - **esi-location.read_online.v1**: EVE SSO scope esi-location.read_online.v1
 - **esi-mail.organize_mail.v1**: EVE SSO scope esi-mail.organize_mail.v1
 - **esi-mail.read_mail.v1**: EVE SSO scope esi-mail.read_mail.v1
 - **esi-planets.manage_planets.v1**: EVE SSO scope esi-planets.manage_planets.v1
 - **esi-search.search_structures.v1**: EVE SSO scope esi-search.search_structures.v1
 - **esi-skills.read_skillqueue.v1**: EVE SSO scope esi-skills.read_skillqueue.v1
 - **esi-skills.read_skills.v1**: EVE SSO scope esi-skills.read_skills.v1
 - **esi-ui.write_waypoint.v1**: EVE SSO scope esi-ui.write_waypoint.v1

Example:
```
  auth, err = oauth2conf.TokenSource(http.Client, token) // Access and Refresh token structure

  client, err := goesiv2.NewClient(nil, "goesiv2 client http://mysite.com")
  ctx := context.WithValue(context.TODO(), goesiv2.ContextOAuth2, auth)
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


