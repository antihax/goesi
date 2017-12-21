# Go API client for esi

## Documentation for API Endpoints

All URIs are relative to *https://esi.tech.ccp.is*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AllianceApi* | [**GetAlliances**](docs/AllianceApi.md#getalliances) | **Get** /v1/alliances/ | List all alliances
*AllianceApi* | [**GetAlliancesAllianceId**](docs/AllianceApi.md#getalliancesallianceid) | **Get** /v3/alliances/{alliance_id}/ | Get alliance information
*AllianceApi* | [**GetAlliancesAllianceIdCorporations**](docs/AllianceApi.md#getalliancesallianceidcorporations) | **Get** /v1/alliances/{alliance_id}/corporations/ | List alliance&#39;s corporations
*AllianceApi* | [**GetAlliancesAllianceIdIcons**](docs/AllianceApi.md#getalliancesallianceidicons) | **Get** /v1/alliances/{alliance_id}/icons/ | Get alliance icon
*AllianceApi* | [**GetAlliancesNames**](docs/AllianceApi.md#getalliancesnames) | **Get** /v2/alliances/names/ | Get alliance names
*AssetsApi* | [**GetCharactersCharacterIdAssets**](docs/AssetsApi.md#getcharacterscharacteridassets) | **Get** /v3/characters/{character_id}/assets/ | Get character assets
*AssetsApi* | [**GetCorporationsCorporationIdAssets**](docs/AssetsApi.md#getcorporationscorporationidassets) | **Get** /v2/corporations/{corporation_id}/assets/ | Get corporation assets
*AssetsApi* | [**PostCharactersCharacterIdAssetsLocations**](docs/AssetsApi.md#postcharacterscharacteridassetslocations) | **Post** /v2/characters/{character_id}/assets/locations/ | Get character asset locations
*AssetsApi* | [**PostCharactersCharacterIdAssetsNames**](docs/AssetsApi.md#postcharacterscharacteridassetsnames) | **Post** /v1/characters/{character_id}/assets/names/ | Get character asset names
*AssetsApi* | [**PostCorporationsCorporationIdAssetsLocations**](docs/AssetsApi.md#postcorporationscorporationidassetslocations) | **Post** /v2/corporations/{corporation_id}/assets/locations/ | Get corporation asset locations
*AssetsApi* | [**PostCorporationsCorporationIdAssetsNames**](docs/AssetsApi.md#postcorporationscorporationidassetsnames) | **Post** /v1/corporations/{corporation_id}/assets/names/ | Get coporation asset names
*BookmarksApi* | [**GetCharactersCharacterIdBookmarks**](docs/BookmarksApi.md#getcharacterscharacteridbookmarks) | **Get** /v2/characters/{character_id}/bookmarks/ | List bookmarks
*BookmarksApi* | [**GetCharactersCharacterIdBookmarksFolders**](docs/BookmarksApi.md#getcharacterscharacteridbookmarksfolders) | **Get** /v2/characters/{character_id}/bookmarks/folders/ | List bookmark folders
*BookmarksApi* | [**GetCorporationsCorporationIdBookmarks**](docs/BookmarksApi.md#getcorporationscorporationidbookmarks) | **Get** /v1/corporations/{corporation_id}/bookmarks/ | List corporation bookmarks
*BookmarksApi* | [**GetCorporationsCorporationIdBookmarksFolders**](docs/BookmarksApi.md#getcorporationscorporationidbookmarksfolders) | **Get** /v1/corporations/{corporation_id}/bookmarks/folders/ | List corporation bookmark folders
*CalendarApi* | [**GetCharactersCharacterIdCalendar**](docs/CalendarApi.md#getcharacterscharacteridcalendar) | **Get** /v1/characters/{character_id}/calendar/ | List calendar event summaries
*CalendarApi* | [**GetCharactersCharacterIdCalendarEventId**](docs/CalendarApi.md#getcharacterscharacteridcalendareventid) | **Get** /v3/characters/{character_id}/calendar/{event_id}/ | Get an event
*CalendarApi* | [**GetCharactersCharacterIdCalendarEventIdAttendees**](docs/CalendarApi.md#getcharacterscharacteridcalendareventidattendees) | **Get** /v1/characters/{character_id}/calendar/{event_id}/attendees/ | Get attendees
*CalendarApi* | [**PutCharactersCharacterIdCalendarEventId**](docs/CalendarApi.md#putcharacterscharacteridcalendareventid) | **Put** /v3/characters/{character_id}/calendar/{event_id}/ | Respond to an event
*CharacterApi* | [**GetCharactersCharacterId**](docs/CharacterApi.md#getcharacterscharacterid) | **Get** /v4/characters/{character_id}/ | Get character&#39;s public information
*CharacterApi* | [**GetCharactersCharacterIdAgentsResearch**](docs/CharacterApi.md#getcharacterscharacteridagentsresearch) | **Get** /v1/characters/{character_id}/agents_research/ | Get agents research
*CharacterApi* | [**GetCharactersCharacterIdBlueprints**](docs/CharacterApi.md#getcharacterscharacteridblueprints) | **Get** /v2/characters/{character_id}/blueprints/ | Get blueprints
*CharacterApi* | [**GetCharactersCharacterIdChatChannels**](docs/CharacterApi.md#getcharacterscharacteridchatchannels) | **Get** /v1/characters/{character_id}/chat_channels/ | Get chat channels
*CharacterApi* | [**GetCharactersCharacterIdCorporationhistory**](docs/CharacterApi.md#getcharacterscharacteridcorporationhistory) | **Get** /v1/characters/{character_id}/corporationhistory/ | Get corporation history
*CharacterApi* | [**GetCharactersCharacterIdFatigue**](docs/CharacterApi.md#getcharacterscharacteridfatigue) | **Get** /v1/characters/{character_id}/fatigue/ | Get jump fatigue
*CharacterApi* | [**GetCharactersCharacterIdMedals**](docs/CharacterApi.md#getcharacterscharacteridmedals) | **Get** /v1/characters/{character_id}/medals/ | Get medals
*CharacterApi* | [**GetCharactersCharacterIdNotifications**](docs/CharacterApi.md#getcharacterscharacteridnotifications) | **Get** /v1/characters/{character_id}/notifications/ | Get character notifications
*CharacterApi* | [**GetCharactersCharacterIdNotificationsContacts**](docs/CharacterApi.md#getcharacterscharacteridnotificationscontacts) | **Get** /v1/characters/{character_id}/notifications/contacts/ | Get new contact notifications
*CharacterApi* | [**GetCharactersCharacterIdPortrait**](docs/CharacterApi.md#getcharacterscharacteridportrait) | **Get** /v2/characters/{character_id}/portrait/ | Get character portraits
*CharacterApi* | [**GetCharactersCharacterIdRoles**](docs/CharacterApi.md#getcharacterscharacteridroles) | **Get** /v2/characters/{character_id}/roles/ | Get character corporation roles
*CharacterApi* | [**GetCharactersCharacterIdStandings**](docs/CharacterApi.md#getcharacterscharacteridstandings) | **Get** /v1/characters/{character_id}/standings/ | Get standings
*CharacterApi* | [**GetCharactersCharacterIdStats**](docs/CharacterApi.md#getcharacterscharacteridstats) | **Get** /v1/characters/{character_id}/stats/ | Yearly aggregate stats
*CharacterApi* | [**GetCharactersCharacterIdTitles**](docs/CharacterApi.md#getcharacterscharacteridtitles) | **Get** /v1/characters/{character_id}/titles/ | Get character corporation titles
*CharacterApi* | [**GetCharactersNames**](docs/CharacterApi.md#getcharactersnames) | **Get** /v1/characters/names/ | Get character names
*CharacterApi* | [**PostCharactersAffiliation**](docs/CharacterApi.md#postcharactersaffiliation) | **Post** /v1/characters/affiliation/ | Character affiliation
*CharacterApi* | [**PostCharactersCharacterIdCspa**](docs/CharacterApi.md#postcharacterscharacteridcspa) | **Post** /v4/characters/{character_id}/cspa/ | Calculate a CSPA charge cost
*ClonesApi* | [**GetCharactersCharacterIdClones**](docs/ClonesApi.md#getcharacterscharacteridclones) | **Get** /v3/characters/{character_id}/clones/ | Get clones
*ClonesApi* | [**GetCharactersCharacterIdImplants**](docs/ClonesApi.md#getcharacterscharacteridimplants) | **Get** /v1/characters/{character_id}/implants/ | Get active implants
*ContactsApi* | [**DeleteCharactersCharacterIdContacts**](docs/ContactsApi.md#deletecharacterscharacteridcontacts) | **Delete** /v2/characters/{character_id}/contacts/ | Delete contacts
*ContactsApi* | [**GetAlliancesAllianceIdContacts**](docs/ContactsApi.md#getalliancesallianceidcontacts) | **Get** /v1/alliances/{alliance_id}/contacts/ | Get alliance contacts
*ContactsApi* | [**GetCharactersCharacterIdContacts**](docs/ContactsApi.md#getcharacterscharacteridcontacts) | **Get** /v1/characters/{character_id}/contacts/ | Get contacts
*ContactsApi* | [**GetCharactersCharacterIdContactsLabels**](docs/ContactsApi.md#getcharacterscharacteridcontactslabels) | **Get** /v1/characters/{character_id}/contacts/labels/ | Get contact labels
*ContactsApi* | [**GetCorporationsCorporationIdContacts**](docs/ContactsApi.md#getcorporationscorporationidcontacts) | **Get** /v1/corporations/{corporation_id}/contacts/ | Get corporation contacts
*ContactsApi* | [**PostCharactersCharacterIdContacts**](docs/ContactsApi.md#postcharacterscharacteridcontacts) | **Post** /v1/characters/{character_id}/contacts/ | Add contacts
*ContactsApi* | [**PutCharactersCharacterIdContacts**](docs/ContactsApi.md#putcharacterscharacteridcontacts) | **Put** /v1/characters/{character_id}/contacts/ | Edit contacts
*ContractsApi* | [**GetCharactersCharacterIdContracts**](docs/ContractsApi.md#getcharacterscharacteridcontracts) | **Get** /v1/characters/{character_id}/contracts/ | Get contracts
*ContractsApi* | [**GetCharactersCharacterIdContractsContractIdBids**](docs/ContractsApi.md#getcharacterscharacteridcontractscontractidbids) | **Get** /v1/characters/{character_id}/contracts/{contract_id}/bids/ | Get contract bids
*ContractsApi* | [**GetCharactersCharacterIdContractsContractIdItems**](docs/ContractsApi.md#getcharacterscharacteridcontractscontractiditems) | **Get** /v1/characters/{character_id}/contracts/{contract_id}/items/ | Get contract items
*ContractsApi* | [**GetCorporationsCorporationIdContracts**](docs/ContractsApi.md#getcorporationscorporationidcontracts) | **Get** /v1/corporations/{corporation_id}/contracts/ | Get coporation contracts
*ContractsApi* | [**GetCorporationsCorporationIdContractsContractIdBids**](docs/ContractsApi.md#getcorporationscorporationidcontractscontractidbids) | **Get** /v1/corporations/{corporation_id}/contracts/{contract_id}/bids/ | Get corporation contract bids
*ContractsApi* | [**GetCorporationsCorporationIdContractsContractIdItems**](docs/ContractsApi.md#getcorporationscorporationidcontractscontractiditems) | **Get** /v1/corporations/{corporation_id}/contracts/{contract_id}/items/ | Get corporation contract items
*CorporationApi* | [**GetCorporationsCorporationId**](docs/CorporationApi.md#getcorporationscorporationid) | **Get** /v4/corporations/{corporation_id}/ | Get corporation information
*CorporationApi* | [**GetCorporationsCorporationIdAlliancehistory**](docs/CorporationApi.md#getcorporationscorporationidalliancehistory) | **Get** /v2/corporations/{corporation_id}/alliancehistory/ | Get alliance history
*CorporationApi* | [**GetCorporationsCorporationIdBlueprints**](docs/CorporationApi.md#getcorporationscorporationidblueprints) | **Get** /v1/corporations/{corporation_id}/blueprints/ | Get corporation blueprints
*CorporationApi* | [**GetCorporationsCorporationIdContainersLogs**](docs/CorporationApi.md#getcorporationscorporationidcontainerslogs) | **Get** /v1/corporations/{corporation_id}/containers/logs/ | Get all corporation ALSC logs
*CorporationApi* | [**GetCorporationsCorporationIdDivisions**](docs/CorporationApi.md#getcorporationscorporationiddivisions) | **Get** /v1/corporations/{corporation_id}/divisions/ | Get corporation divisions
*CorporationApi* | [**GetCorporationsCorporationIdFacilities**](docs/CorporationApi.md#getcorporationscorporationidfacilities) | **Get** /v1/corporations/{corporation_id}/facilities/ | Get corporation facilities
*CorporationApi* | [**GetCorporationsCorporationIdIcons**](docs/CorporationApi.md#getcorporationscorporationidicons) | **Get** /v1/corporations/{corporation_id}/icons/ | Get corporation icon
*CorporationApi* | [**GetCorporationsCorporationIdMedals**](docs/CorporationApi.md#getcorporationscorporationidmedals) | **Get** /v1/corporations/{corporation_id}/medals/ | Get corporation medals
*CorporationApi* | [**GetCorporationsCorporationIdMedalsIssued**](docs/CorporationApi.md#getcorporationscorporationidmedalsissued) | **Get** /v1/corporations/{corporation_id}/medals/issued/ | Get corporation issued medals
*CorporationApi* | [**GetCorporationsCorporationIdMembers**](docs/CorporationApi.md#getcorporationscorporationidmembers) | **Get** /v3/corporations/{corporation_id}/members/ | Get corporation members
*CorporationApi* | [**GetCorporationsCorporationIdMembersLimit**](docs/CorporationApi.md#getcorporationscorporationidmemberslimit) | **Get** /v1/corporations/{corporation_id}/members/limit/ | Get corporation member limit
*CorporationApi* | [**GetCorporationsCorporationIdMembersTitles**](docs/CorporationApi.md#getcorporationscorporationidmemberstitles) | **Get** /v1/corporations/{corporation_id}/members/titles/ | Get corporation&#39;s members&#39; titles
*CorporationApi* | [**GetCorporationsCorporationIdMembertracking**](docs/CorporationApi.md#getcorporationscorporationidmembertracking) | **Get** /v1/corporations/{corporation_id}/membertracking/ | Track corporation members
*CorporationApi* | [**GetCorporationsCorporationIdOutposts**](docs/CorporationApi.md#getcorporationscorporationidoutposts) | **Get** /v1/corporations/{corporation_id}/outposts/ | Get corporation outposts
*CorporationApi* | [**GetCorporationsCorporationIdOutpostsOutpostId**](docs/CorporationApi.md#getcorporationscorporationidoutpostsoutpostid) | **Get** /v1/corporations/{corporation_id}/outposts/{outpost_id}/ | Get corporation outpost details
*CorporationApi* | [**GetCorporationsCorporationIdRoles**](docs/CorporationApi.md#getcorporationscorporationidroles) | **Get** /v1/corporations/{corporation_id}/roles/ | Get corporation member roles
*CorporationApi* | [**GetCorporationsCorporationIdRolesHistory**](docs/CorporationApi.md#getcorporationscorporationidroleshistory) | **Get** /v1/corporations/{corporation_id}/roles/history/ | Get corporation member roles history
*CorporationApi* | [**GetCorporationsCorporationIdShareholders**](docs/CorporationApi.md#getcorporationscorporationidshareholders) | **Get** /v1/corporations/{corporation_id}/shareholders/ | Get corporation members
*CorporationApi* | [**GetCorporationsCorporationIdStandings**](docs/CorporationApi.md#getcorporationscorporationidstandings) | **Get** /v1/corporations/{corporation_id}/standings/ | Get corporation standings
*CorporationApi* | [**GetCorporationsCorporationIdStarbases**](docs/CorporationApi.md#getcorporationscorporationidstarbases) | **Get** /v1/corporations/{corporation_id}/starbases/ | Get corporation starbases (POSes)
*CorporationApi* | [**GetCorporationsCorporationIdStarbasesStarbaseId**](docs/CorporationApi.md#getcorporationscorporationidstarbasesstarbaseid) | **Get** /v1/corporations/{corporation_id}/starbases/{starbase_id}/ | Get starbase (POS) detail
*CorporationApi* | [**GetCorporationsCorporationIdStructures**](docs/CorporationApi.md#getcorporationscorporationidstructures) | **Get** /v1/corporations/{corporation_id}/structures/ | Get corporation structures
*CorporationApi* | [**GetCorporationsCorporationIdTitles**](docs/CorporationApi.md#getcorporationscorporationidtitles) | **Get** /v1/corporations/{corporation_id}/titles/ | Get corporation titles
*CorporationApi* | [**GetCorporationsNames**](docs/CorporationApi.md#getcorporationsnames) | **Get** /v2/corporations/names/ | Get corporation names
*CorporationApi* | [**GetCorporationsNpccorps**](docs/CorporationApi.md#getcorporationsnpccorps) | **Get** /v1/corporations/npccorps/ | Get npc corporations
*CorporationApi* | [**PutCorporationsCorporationIdStructuresStructureId**](docs/CorporationApi.md#putcorporationscorporationidstructuresstructureid) | **Put** /v1/corporations/{corporation_id}/structures/{structure_id}/ | Update structure vulnerability schedule
*DogmaApi* | [**GetDogmaAttributes**](docs/DogmaApi.md#getdogmaattributes) | **Get** /v1/dogma/attributes/ | Get attributes
*DogmaApi* | [**GetDogmaAttributesAttributeId**](docs/DogmaApi.md#getdogmaattributesattributeid) | **Get** /v1/dogma/attributes/{attribute_id}/ | Get attribute information
*DogmaApi* | [**GetDogmaEffects**](docs/DogmaApi.md#getdogmaeffects) | **Get** /v1/dogma/effects/ | Get effects
*DogmaApi* | [**GetDogmaEffectsEffectId**](docs/DogmaApi.md#getdogmaeffectseffectid) | **Get** /v2/dogma/effects/{effect_id}/ | Get effect information
*FactionWarfareApi* | [**GetCharactersCharacterIdFwStats**](docs/FactionWarfareApi.md#getcharacterscharacteridfwstats) | **Get** /v1/characters/{character_id}/fw/stats/ | Overview of a character involved in faction warfare
*FactionWarfareApi* | [**GetCorporationsCorporationIdFwStats**](docs/FactionWarfareApi.md#getcorporationscorporationidfwstats) | **Get** /v1/corporations/{corporation_id}/fw/stats/ | Overview of a corporation involved in faction warfare
*FactionWarfareApi* | [**GetFwLeaderboards**](docs/FactionWarfareApi.md#getfwleaderboards) | **Get** /v1/fw/leaderboards/ | List of the top factions in faction warfare
*FactionWarfareApi* | [**GetFwLeaderboardsCharacters**](docs/FactionWarfareApi.md#getfwleaderboardscharacters) | **Get** /v1/fw/leaderboards/characters/ | List of the top pilots in faction warfare
*FactionWarfareApi* | [**GetFwLeaderboardsCorporations**](docs/FactionWarfareApi.md#getfwleaderboardscorporations) | **Get** /v1/fw/leaderboards/corporations/ | List of the top corporations in faction warfare
*FactionWarfareApi* | [**GetFwStats**](docs/FactionWarfareApi.md#getfwstats) | **Get** /v1/fw/stats/ | An overview of statistics about factions involved in faction warfare
*FactionWarfareApi* | [**GetFwSystems**](docs/FactionWarfareApi.md#getfwsystems) | **Get** /v1/fw/systems/ | Ownership of faction warfare systems
*FactionWarfareApi* | [**GetFwWars**](docs/FactionWarfareApi.md#getfwwars) | **Get** /v1/fw/wars/ | Data about which NPC factions are at war
*FittingsApi* | [**DeleteCharactersCharacterIdFittingsFittingId**](docs/FittingsApi.md#deletecharacterscharacteridfittingsfittingid) | **Delete** /v1/characters/{character_id}/fittings/{fitting_id}/ | Delete fitting
*FittingsApi* | [**GetCharactersCharacterIdFittings**](docs/FittingsApi.md#getcharacterscharacteridfittings) | **Get** /v1/characters/{character_id}/fittings/ | Get fittings
*FittingsApi* | [**PostCharactersCharacterIdFittings**](docs/FittingsApi.md#postcharacterscharacteridfittings) | **Post** /v1/characters/{character_id}/fittings/ | Create fitting
*FleetsApi* | [**DeleteFleetsFleetIdMembersMemberId**](docs/FleetsApi.md#deletefleetsfleetidmembersmemberid) | **Delete** /v1/fleets/{fleet_id}/members/{member_id}/ | Kick fleet member
*FleetsApi* | [**DeleteFleetsFleetIdSquadsSquadId**](docs/FleetsApi.md#deletefleetsfleetidsquadssquadid) | **Delete** /v1/fleets/{fleet_id}/squads/{squad_id}/ | Delete fleet squad
*FleetsApi* | [**DeleteFleetsFleetIdWingsWingId**](docs/FleetsApi.md#deletefleetsfleetidwingswingid) | **Delete** /v1/fleets/{fleet_id}/wings/{wing_id}/ | Delete fleet wing
*FleetsApi* | [**GetCharactersCharacterIdFleet**](docs/FleetsApi.md#getcharacterscharacteridfleet) | **Get** /v1/characters/{character_id}/fleet/ | Get character fleet info
*FleetsApi* | [**GetFleetsFleetId**](docs/FleetsApi.md#getfleetsfleetid) | **Get** /v1/fleets/{fleet_id}/ | Get fleet information
*FleetsApi* | [**GetFleetsFleetIdMembers**](docs/FleetsApi.md#getfleetsfleetidmembers) | **Get** /v1/fleets/{fleet_id}/members/ | Get fleet members
*FleetsApi* | [**GetFleetsFleetIdWings**](docs/FleetsApi.md#getfleetsfleetidwings) | **Get** /v1/fleets/{fleet_id}/wings/ | Get fleet wings
*FleetsApi* | [**PostFleetsFleetIdMembers**](docs/FleetsApi.md#postfleetsfleetidmembers) | **Post** /v1/fleets/{fleet_id}/members/ | Create fleet invitation
*FleetsApi* | [**PostFleetsFleetIdWings**](docs/FleetsApi.md#postfleetsfleetidwings) | **Post** /v1/fleets/{fleet_id}/wings/ | Create fleet wing
*FleetsApi* | [**PostFleetsFleetIdWingsWingIdSquads**](docs/FleetsApi.md#postfleetsfleetidwingswingidsquads) | **Post** /v1/fleets/{fleet_id}/wings/{wing_id}/squads/ | Create fleet squad
*FleetsApi* | [**PutFleetsFleetId**](docs/FleetsApi.md#putfleetsfleetid) | **Put** /v1/fleets/{fleet_id}/ | Update fleet
*FleetsApi* | [**PutFleetsFleetIdMembersMemberId**](docs/FleetsApi.md#putfleetsfleetidmembersmemberid) | **Put** /v1/fleets/{fleet_id}/members/{member_id}/ | Move fleet member
*FleetsApi* | [**PutFleetsFleetIdSquadsSquadId**](docs/FleetsApi.md#putfleetsfleetidsquadssquadid) | **Put** /v1/fleets/{fleet_id}/squads/{squad_id}/ | Rename fleet squad
*FleetsApi* | [**PutFleetsFleetIdWingsWingId**](docs/FleetsApi.md#putfleetsfleetidwingswingid) | **Put** /v1/fleets/{fleet_id}/wings/{wing_id}/ | Rename fleet wing
*IncursionsApi* | [**GetIncursions**](docs/IncursionsApi.md#getincursions) | **Get** /v1/incursions/ | List incursions
*IndustryApi* | [**GetCharactersCharacterIdIndustryJobs**](docs/IndustryApi.md#getcharacterscharacteridindustryjobs) | **Get** /v1/characters/{character_id}/industry/jobs/ | List character industry jobs
*IndustryApi* | [**GetCharactersCharacterIdMining**](docs/IndustryApi.md#getcharacterscharacteridmining) | **Get** /v1/characters/{character_id}/mining/ | Character mining ledger
*IndustryApi* | [**GetCorporationCorporationIdMiningExtractions**](docs/IndustryApi.md#getcorporationcorporationidminingextractions) | **Get** /v1/corporation/{corporation_id}/mining/extractions/ | Moon extraction timers
*IndustryApi* | [**GetCorporationCorporationIdMiningObservers**](docs/IndustryApi.md#getcorporationcorporationidminingobservers) | **Get** /v1/corporation/{corporation_id}/mining/observers/ | Corporation mining observers
*IndustryApi* | [**GetCorporationCorporationIdMiningObserversObserverId**](docs/IndustryApi.md#getcorporationcorporationidminingobserversobserverid) | **Get** /v1/corporation/{corporation_id}/mining/observers/{observer_id}/ | Observed corporation mining
*IndustryApi* | [**GetCorporationsCorporationIdIndustryJobs**](docs/IndustryApi.md#getcorporationscorporationidindustryjobs) | **Get** /v1/corporations/{corporation_id}/industry/jobs/ | List corporation industry jobs
*IndustryApi* | [**GetIndustryFacilities**](docs/IndustryApi.md#getindustryfacilities) | **Get** /v1/industry/facilities/ | List industry facilities
*IndustryApi* | [**GetIndustrySystems**](docs/IndustryApi.md#getindustrysystems) | **Get** /v1/industry/systems/ | List solar system cost indices
*InsuranceApi* | [**GetInsurancePrices**](docs/InsuranceApi.md#getinsuranceprices) | **Get** /v1/insurance/prices/ | List insurance levels
*KillmailsApi* | [**GetCharactersCharacterIdKillmailsRecent**](docs/KillmailsApi.md#getcharacterscharacteridkillmailsrecent) | **Get** /v1/characters/{character_id}/killmails/recent/ | Get character kills and losses
*KillmailsApi* | [**GetCorporationsCorporationIdKillmailsRecent**](docs/KillmailsApi.md#getcorporationscorporationidkillmailsrecent) | **Get** /v1/corporations/{corporation_id}/killmails/recent/ | Get corporation kills and losses
*KillmailsApi* | [**GetKillmailsKillmailIdKillmailHash**](docs/KillmailsApi.md#getkillmailskillmailidkillmailhash) | **Get** /v1/killmails/{killmail_id}/{killmail_hash}/ | Get a single killmail
*LocationApi* | [**GetCharactersCharacterIdLocation**](docs/LocationApi.md#getcharacterscharacteridlocation) | **Get** /v1/characters/{character_id}/location/ | Get character location
*LocationApi* | [**GetCharactersCharacterIdOnline**](docs/LocationApi.md#getcharacterscharacteridonline) | **Get** /v2/characters/{character_id}/online/ | Get character online
*LocationApi* | [**GetCharactersCharacterIdShip**](docs/LocationApi.md#getcharacterscharacteridship) | **Get** /v1/characters/{character_id}/ship/ | Get current ship
*LoyaltyApi* | [**GetCharactersCharacterIdLoyaltyPoints**](docs/LoyaltyApi.md#getcharacterscharacteridloyaltypoints) | **Get** /v1/characters/{character_id}/loyalty/points/ | Get loyalty points
*LoyaltyApi* | [**GetLoyaltyStoresCorporationIdOffers**](docs/LoyaltyApi.md#getloyaltystorescorporationidoffers) | **Get** /v1/loyalty/stores/{corporation_id}/offers/ | List loyalty store offers
*MailApi* | [**DeleteCharactersCharacterIdMailLabelsLabelId**](docs/MailApi.md#deletecharacterscharacteridmaillabelslabelid) | **Delete** /v1/characters/{character_id}/mail/labels/{label_id}/ | Delete a mail label
*MailApi* | [**DeleteCharactersCharacterIdMailMailId**](docs/MailApi.md#deletecharacterscharacteridmailmailid) | **Delete** /v1/characters/{character_id}/mail/{mail_id}/ | Delete a mail
*MailApi* | [**GetCharactersCharacterIdMail**](docs/MailApi.md#getcharacterscharacteridmail) | **Get** /v1/characters/{character_id}/mail/ | Return mail headers
*MailApi* | [**GetCharactersCharacterIdMailLabels**](docs/MailApi.md#getcharacterscharacteridmaillabels) | **Get** /v3/characters/{character_id}/mail/labels/ | Get mail labels and unread counts
*MailApi* | [**GetCharactersCharacterIdMailLists**](docs/MailApi.md#getcharacterscharacteridmaillists) | **Get** /v1/characters/{character_id}/mail/lists/ | Return mailing list subscriptions
*MailApi* | [**GetCharactersCharacterIdMailMailId**](docs/MailApi.md#getcharacterscharacteridmailmailid) | **Get** /v1/characters/{character_id}/mail/{mail_id}/ | Return a mail
*MailApi* | [**PostCharactersCharacterIdMail**](docs/MailApi.md#postcharacterscharacteridmail) | **Post** /v1/characters/{character_id}/mail/ | Send a new mail
*MailApi* | [**PostCharactersCharacterIdMailLabels**](docs/MailApi.md#postcharacterscharacteridmaillabels) | **Post** /v2/characters/{character_id}/mail/labels/ | Create a mail label
*MailApi* | [**PutCharactersCharacterIdMailMailId**](docs/MailApi.md#putcharacterscharacteridmailmailid) | **Put** /v1/characters/{character_id}/mail/{mail_id}/ | Update metadata about a mail
*MarketApi* | [**GetCharactersCharacterIdOrders**](docs/MarketApi.md#getcharacterscharacteridorders) | **Get** /v1/characters/{character_id}/orders/ | List orders from a character
*MarketApi* | [**GetCorporationsCorporationIdOrders**](docs/MarketApi.md#getcorporationscorporationidorders) | **Get** /v1/corporations/{corporation_id}/orders/ | List orders from a corporation
*MarketApi* | [**GetMarketsGroups**](docs/MarketApi.md#getmarketsgroups) | **Get** /v1/markets/groups/ | Get item groups
*MarketApi* | [**GetMarketsGroupsMarketGroupId**](docs/MarketApi.md#getmarketsgroupsmarketgroupid) | **Get** /v1/markets/groups/{market_group_id}/ | Get item group information
*MarketApi* | [**GetMarketsPrices**](docs/MarketApi.md#getmarketsprices) | **Get** /v1/markets/prices/ | List market prices
*MarketApi* | [**GetMarketsRegionIdHistory**](docs/MarketApi.md#getmarketsregionidhistory) | **Get** /v1/markets/{region_id}/history/ | List historical market statistics in a region
*MarketApi* | [**GetMarketsRegionIdOrders**](docs/MarketApi.md#getmarketsregionidorders) | **Get** /v1/markets/{region_id}/orders/ | List orders in a region
*MarketApi* | [**GetMarketsRegionIdTypes**](docs/MarketApi.md#getmarketsregionidtypes) | **Get** /v1/markets/{region_id}/types/ | List type IDs relevant to a market
*MarketApi* | [**GetMarketsStructuresStructureId**](docs/MarketApi.md#getmarketsstructuresstructureid) | **Get** /v1/markets/structures/{structure_id}/ | List orders in a structure
*OpportunitiesApi* | [**GetCharactersCharacterIdOpportunities**](docs/OpportunitiesApi.md#getcharacterscharacteridopportunities) | **Get** /v1/characters/{character_id}/opportunities/ | Get a character&#39;s completed tasks
*OpportunitiesApi* | [**GetOpportunitiesGroups**](docs/OpportunitiesApi.md#getopportunitiesgroups) | **Get** /v1/opportunities/groups/ | Get opportunities groups
*OpportunitiesApi* | [**GetOpportunitiesGroupsGroupId**](docs/OpportunitiesApi.md#getopportunitiesgroupsgroupid) | **Get** /v1/opportunities/groups/{group_id}/ | Get opportunities group
*OpportunitiesApi* | [**GetOpportunitiesTasks**](docs/OpportunitiesApi.md#getopportunitiestasks) | **Get** /v1/opportunities/tasks/ | Get opportunities tasks
*OpportunitiesApi* | [**GetOpportunitiesTasksTaskId**](docs/OpportunitiesApi.md#getopportunitiestaskstaskid) | **Get** /v1/opportunities/tasks/{task_id}/ | Get opportunities task
*PlanetaryInteractionApi* | [**GetCharactersCharacterIdPlanets**](docs/PlanetaryInteractionApi.md#getcharacterscharacteridplanets) | **Get** /v1/characters/{character_id}/planets/ | Get colonies
*PlanetaryInteractionApi* | [**GetCharactersCharacterIdPlanetsPlanetId**](docs/PlanetaryInteractionApi.md#getcharacterscharacteridplanetsplanetid) | **Get** /v3/characters/{character_id}/planets/{planet_id}/ | Get colony layout
*PlanetaryInteractionApi* | [**GetCorporationsCorporationIdCustomsOffices**](docs/PlanetaryInteractionApi.md#getcorporationscorporationidcustomsoffices) | **Get** /v1/corporations/{corporation_id}/customs_offices/ | List corporation customs offices
*PlanetaryInteractionApi* | [**GetUniverseSchematicsSchematicId**](docs/PlanetaryInteractionApi.md#getuniverseschematicsschematicid) | **Get** /v1/universe/schematics/{schematic_id}/ | Get schematic information
*RoutesApi* | [**GetRouteOriginDestination**](docs/RoutesApi.md#getrouteorigindestination) | **Get** /v1/route/{origin}/{destination}/ | Get route
*SearchApi* | [**GetCharactersCharacterIdSearch**](docs/SearchApi.md#getcharacterscharacteridsearch) | **Get** /v3/characters/{character_id}/search/ | Search on a string
*SearchApi* | [**GetSearch**](docs/SearchApi.md#getsearch) | **Get** /v2/search/ | Search on a string
*SkillsApi* | [**GetCharactersCharacterIdAttributes**](docs/SkillsApi.md#getcharacterscharacteridattributes) | **Get** /v1/characters/{character_id}/attributes/ | Get character attributes
*SkillsApi* | [**GetCharactersCharacterIdSkillqueue**](docs/SkillsApi.md#getcharacterscharacteridskillqueue) | **Get** /v2/characters/{character_id}/skillqueue/ | Get character&#39;s skill queue
*SkillsApi* | [**GetCharactersCharacterIdSkills**](docs/SkillsApi.md#getcharacterscharacteridskills) | **Get** /v4/characters/{character_id}/skills/ | Get character skills
*SovereigntyApi* | [**GetSovereigntyCampaigns**](docs/SovereigntyApi.md#getsovereigntycampaigns) | **Get** /v1/sovereignty/campaigns/ | List sovereignty campaigns
*SovereigntyApi* | [**GetSovereigntyMap**](docs/SovereigntyApi.md#getsovereigntymap) | **Get** /v1/sovereignty/map/ | List sovereignty of systems
*SovereigntyApi* | [**GetSovereigntyStructures**](docs/SovereigntyApi.md#getsovereigntystructures) | **Get** /v1/sovereignty/structures/ | List sovereignty structures
*StatusApi* | [**GetStatus**](docs/StatusApi.md#getstatus) | **Get** /v1/status/ | Retrieve the uptime and player counts
*UniverseApi* | [**GetUniverseBloodlines**](docs/UniverseApi.md#getuniversebloodlines) | **Get** /v1/universe/bloodlines/ | Get bloodlines
*UniverseApi* | [**GetUniverseCategories**](docs/UniverseApi.md#getuniversecategories) | **Get** /v1/universe/categories/ | Get item categories
*UniverseApi* | [**GetUniverseCategoriesCategoryId**](docs/UniverseApi.md#getuniversecategoriescategoryid) | **Get** /v1/universe/categories/{category_id}/ | Get item category information
*UniverseApi* | [**GetUniverseConstellations**](docs/UniverseApi.md#getuniverseconstellations) | **Get** /v1/universe/constellations/ | Get constellations
*UniverseApi* | [**GetUniverseConstellationsConstellationId**](docs/UniverseApi.md#getuniverseconstellationsconstellationid) | **Get** /v1/universe/constellations/{constellation_id}/ | Get constellation information
*UniverseApi* | [**GetUniverseFactions**](docs/UniverseApi.md#getuniversefactions) | **Get** /v2/universe/factions/ | Get factions
*UniverseApi* | [**GetUniverseGraphics**](docs/UniverseApi.md#getuniversegraphics) | **Get** /v1/universe/graphics/ | Get graphics
*UniverseApi* | [**GetUniverseGraphicsGraphicId**](docs/UniverseApi.md#getuniversegraphicsgraphicid) | **Get** /v1/universe/graphics/{graphic_id}/ | Get graphic information
*UniverseApi* | [**GetUniverseGroups**](docs/UniverseApi.md#getuniversegroups) | **Get** /v1/universe/groups/ | Get item groups
*UniverseApi* | [**GetUniverseGroupsGroupId**](docs/UniverseApi.md#getuniversegroupsgroupid) | **Get** /v1/universe/groups/{group_id}/ | Get item group information
*UniverseApi* | [**GetUniverseMoonsMoonId**](docs/UniverseApi.md#getuniversemoonsmoonid) | **Get** /v1/universe/moons/{moon_id}/ | Get moon information
*UniverseApi* | [**GetUniversePlanetsPlanetId**](docs/UniverseApi.md#getuniverseplanetsplanetid) | **Get** /v1/universe/planets/{planet_id}/ | Get planet information
*UniverseApi* | [**GetUniverseRaces**](docs/UniverseApi.md#getuniverseraces) | **Get** /v1/universe/races/ | Get character races
*UniverseApi* | [**GetUniverseRegions**](docs/UniverseApi.md#getuniverseregions) | **Get** /v1/universe/regions/ | Get regions
*UniverseApi* | [**GetUniverseRegionsRegionId**](docs/UniverseApi.md#getuniverseregionsregionid) | **Get** /v1/universe/regions/{region_id}/ | Get region information
*UniverseApi* | [**GetUniverseStargatesStargateId**](docs/UniverseApi.md#getuniversestargatesstargateid) | **Get** /v1/universe/stargates/{stargate_id}/ | Get stargate information
*UniverseApi* | [**GetUniverseStarsStarId**](docs/UniverseApi.md#getuniversestarsstarid) | **Get** /v1/universe/stars/{star_id}/ | Get star information
*UniverseApi* | [**GetUniverseStationsStationId**](docs/UniverseApi.md#getuniversestationsstationid) | **Get** /v2/universe/stations/{station_id}/ | Get station information
*UniverseApi* | [**GetUniverseStructures**](docs/UniverseApi.md#getuniversestructures) | **Get** /v1/universe/structures/ | List all public structures
*UniverseApi* | [**GetUniverseStructuresStructureId**](docs/UniverseApi.md#getuniversestructuresstructureid) | **Get** /v1/universe/structures/{structure_id}/ | Get structure information
*UniverseApi* | [**GetUniverseSystemJumps**](docs/UniverseApi.md#getuniversesystemjumps) | **Get** /v1/universe/system_jumps/ | Get system jumps
*UniverseApi* | [**GetUniverseSystemKills**](docs/UniverseApi.md#getuniversesystemkills) | **Get** /v2/universe/system_kills/ | Get system kills
*UniverseApi* | [**GetUniverseSystems**](docs/UniverseApi.md#getuniversesystems) | **Get** /v1/universe/systems/ | Get solar systems
*UniverseApi* | [**GetUniverseSystemsSystemId**](docs/UniverseApi.md#getuniversesystemssystemid) | **Get** /v3/universe/systems/{system_id}/ | Get solar system information
*UniverseApi* | [**GetUniverseTypes**](docs/UniverseApi.md#getuniversetypes) | **Get** /v1/universe/types/ | Get types
*UniverseApi* | [**GetUniverseTypesTypeId**](docs/UniverseApi.md#getuniversetypestypeid) | **Get** /v3/universe/types/{type_id}/ | Get type information
*UniverseApi* | [**PostUniverseIds**](docs/UniverseApi.md#postuniverseids) | **Post** /v1/universe/ids/ | Bulk names to IDs
*UniverseApi* | [**PostUniverseNames**](docs/UniverseApi.md#postuniversenames) | **Post** /v2/universe/names/ | Get names and categories for a set of ID&#39;s
*UserInterfaceApi* | [**PostUiAutopilotWaypoint**](docs/UserInterfaceApi.md#postuiautopilotwaypoint) | **Post** /v2/ui/autopilot/waypoint/ | Set Autopilot Waypoint
*UserInterfaceApi* | [**PostUiOpenwindowContract**](docs/UserInterfaceApi.md#postuiopenwindowcontract) | **Post** /v1/ui/openwindow/contract/ | Open Contract Window
*UserInterfaceApi* | [**PostUiOpenwindowInformation**](docs/UserInterfaceApi.md#postuiopenwindowinformation) | **Post** /v1/ui/openwindow/information/ | Open Information Window
*UserInterfaceApi* | [**PostUiOpenwindowMarketdetails**](docs/UserInterfaceApi.md#postuiopenwindowmarketdetails) | **Post** /v1/ui/openwindow/marketdetails/ | Open Market Details
*UserInterfaceApi* | [**PostUiOpenwindowNewmail**](docs/UserInterfaceApi.md#postuiopenwindownewmail) | **Post** /v1/ui/openwindow/newmail/ | Open New Mail Window
*WalletApi* | [**GetCharactersCharacterIdWallet**](docs/WalletApi.md#getcharacterscharacteridwallet) | **Get** /v1/characters/{character_id}/wallet/ | Get a character&#39;s wallet balance
*WalletApi* | [**GetCharactersCharacterIdWalletJournal**](docs/WalletApi.md#getcharacterscharacteridwalletjournal) | **Get** /v3/characters/{character_id}/wallet/journal/ | Get character wallet journal
*WalletApi* | [**GetCharactersCharacterIdWalletTransactions**](docs/WalletApi.md#getcharacterscharacteridwallettransactions) | **Get** /v1/characters/{character_id}/wallet/transactions/ | Get wallet transactions
*WalletApi* | [**GetCorporationsCorporationIdWallets**](docs/WalletApi.md#getcorporationscorporationidwallets) | **Get** /v1/corporations/{corporation_id}/wallets/ | Returns a corporation&#39;s wallet balance
*WalletApi* | [**GetCorporationsCorporationIdWalletsDivisionJournal**](docs/WalletApi.md#getcorporationscorporationidwalletsdivisionjournal) | **Get** /v2/corporations/{corporation_id}/wallets/{division}/journal/ | Get corporation wallet journal
*WalletApi* | [**GetCorporationsCorporationIdWalletsDivisionTransactions**](docs/WalletApi.md#getcorporationscorporationidwalletsdivisiontransactions) | **Get** /v1/corporations/{corporation_id}/wallets/{division}/transactions/ | Get corporation wallet transactions
*WarsApi* | [**GetWars**](docs/WarsApi.md#getwars) | **Get** /v1/wars/ | List wars
*WarsApi* | [**GetWarsWarId**](docs/WarsApi.md#getwarswarid) | **Get** /v1/wars/{war_id}/ | Get war information
*WarsApi* | [**GetWarsWarIdKillmails**](docs/WarsApi.md#getwarswaridkillmails) | **Get** /v1/wars/{war_id}/killmails/ | List kills for a war


## Documentation For Models

 - [DeleteCharactersCharacterIdMailLabelsLabelIdUnprocessableEntity](docs/DeleteCharactersCharacterIdMailLabelsLabelIdUnprocessableEntity.md)
 - [DeleteFleetsFleetIdMembersMemberIdNotFound](docs/DeleteFleetsFleetIdMembersMemberIdNotFound.md)
 - [DeleteFleetsFleetIdSquadsSquadIdNotFound](docs/DeleteFleetsFleetIdSquadsSquadIdNotFound.md)
 - [DeleteFleetsFleetIdWingsWingIdNotFound](docs/DeleteFleetsFleetIdWingsWingIdNotFound.md)
 - [Forbidden](docs/Forbidden.md)
 - [GetAlliancesAllianceIdContacts200Ok](docs/GetAlliancesAllianceIdContacts200Ok.md)
 - [GetAlliancesAllianceIdIconsNotFound](docs/GetAlliancesAllianceIdIconsNotFound.md)
 - [GetAlliancesAllianceIdIconsOk](docs/GetAlliancesAllianceIdIconsOk.md)
 - [GetAlliancesAllianceIdNotFound](docs/GetAlliancesAllianceIdNotFound.md)
 - [GetAlliancesAllianceIdOk](docs/GetAlliancesAllianceIdOk.md)
 - [GetAlliancesNames200Ok](docs/GetAlliancesNames200Ok.md)
 - [GetCharactersCharacterIdAgentsResearch200Ok](docs/GetCharactersCharacterIdAgentsResearch200Ok.md)
 - [GetCharactersCharacterIdAssets200Ok](docs/GetCharactersCharacterIdAssets200Ok.md)
 - [GetCharactersCharacterIdAttributesOk](docs/GetCharactersCharacterIdAttributesOk.md)
 - [GetCharactersCharacterIdBlueprints200Ok](docs/GetCharactersCharacterIdBlueprints200Ok.md)
 - [GetCharactersCharacterIdBookmarks200Ok](docs/GetCharactersCharacterIdBookmarks200Ok.md)
 - [GetCharactersCharacterIdBookmarksCoordinates](docs/GetCharactersCharacterIdBookmarksCoordinates.md)
 - [GetCharactersCharacterIdBookmarksFolders200Ok](docs/GetCharactersCharacterIdBookmarksFolders200Ok.md)
 - [GetCharactersCharacterIdBookmarksItem](docs/GetCharactersCharacterIdBookmarksItem.md)
 - [GetCharactersCharacterIdCalendar200Ok](docs/GetCharactersCharacterIdCalendar200Ok.md)
 - [GetCharactersCharacterIdCalendarEventIdAttendees200Ok](docs/GetCharactersCharacterIdCalendarEventIdAttendees200Ok.md)
 - [GetCharactersCharacterIdCalendarEventIdOk](docs/GetCharactersCharacterIdCalendarEventIdOk.md)
 - [GetCharactersCharacterIdChatChannels200Ok](docs/GetCharactersCharacterIdChatChannels200Ok.md)
 - [GetCharactersCharacterIdChatChannelsAllowed](docs/GetCharactersCharacterIdChatChannelsAllowed.md)
 - [GetCharactersCharacterIdChatChannelsBlocked](docs/GetCharactersCharacterIdChatChannelsBlocked.md)
 - [GetCharactersCharacterIdChatChannelsMuted](docs/GetCharactersCharacterIdChatChannelsMuted.md)
 - [GetCharactersCharacterIdChatChannelsOperator](docs/GetCharactersCharacterIdChatChannelsOperator.md)
 - [GetCharactersCharacterIdClonesHomeLocation](docs/GetCharactersCharacterIdClonesHomeLocation.md)
 - [GetCharactersCharacterIdClonesJumpClone](docs/GetCharactersCharacterIdClonesJumpClone.md)
 - [GetCharactersCharacterIdClonesOk](docs/GetCharactersCharacterIdClonesOk.md)
 - [GetCharactersCharacterIdContacts200Ok](docs/GetCharactersCharacterIdContacts200Ok.md)
 - [GetCharactersCharacterIdContactsLabels200Ok](docs/GetCharactersCharacterIdContactsLabels200Ok.md)
 - [GetCharactersCharacterIdContracts200Ok](docs/GetCharactersCharacterIdContracts200Ok.md)
 - [GetCharactersCharacterIdContractsContractIdBids200Ok](docs/GetCharactersCharacterIdContractsContractIdBids200Ok.md)
 - [GetCharactersCharacterIdContractsContractIdItems200Ok](docs/GetCharactersCharacterIdContractsContractIdItems200Ok.md)
 - [GetCharactersCharacterIdCorporationhistory200Ok](docs/GetCharactersCharacterIdCorporationhistory200Ok.md)
 - [GetCharactersCharacterIdFatigueOk](docs/GetCharactersCharacterIdFatigueOk.md)
 - [GetCharactersCharacterIdFittings200Ok](docs/GetCharactersCharacterIdFittings200Ok.md)
 - [GetCharactersCharacterIdFittingsItem](docs/GetCharactersCharacterIdFittingsItem.md)
 - [GetCharactersCharacterIdFleetNotFound](docs/GetCharactersCharacterIdFleetNotFound.md)
 - [GetCharactersCharacterIdFleetOk](docs/GetCharactersCharacterIdFleetOk.md)
 - [GetCharactersCharacterIdFwStatsKills](docs/GetCharactersCharacterIdFwStatsKills.md)
 - [GetCharactersCharacterIdFwStatsOk](docs/GetCharactersCharacterIdFwStatsOk.md)
 - [GetCharactersCharacterIdFwStatsVictoryPoints](docs/GetCharactersCharacterIdFwStatsVictoryPoints.md)
 - [GetCharactersCharacterIdIndustryJobs200Ok](docs/GetCharactersCharacterIdIndustryJobs200Ok.md)
 - [GetCharactersCharacterIdKillmailsRecent200Ok](docs/GetCharactersCharacterIdKillmailsRecent200Ok.md)
 - [GetCharactersCharacterIdLocationOk](docs/GetCharactersCharacterIdLocationOk.md)
 - [GetCharactersCharacterIdLoyaltyPoints200Ok](docs/GetCharactersCharacterIdLoyaltyPoints200Ok.md)
 - [GetCharactersCharacterIdMail200Ok](docs/GetCharactersCharacterIdMail200Ok.md)
 - [GetCharactersCharacterIdMailLabelsLabel](docs/GetCharactersCharacterIdMailLabelsLabel.md)
 - [GetCharactersCharacterIdMailLabelsOk](docs/GetCharactersCharacterIdMailLabelsOk.md)
 - [GetCharactersCharacterIdMailLists200Ok](docs/GetCharactersCharacterIdMailLists200Ok.md)
 - [GetCharactersCharacterIdMailMailIdNotFound](docs/GetCharactersCharacterIdMailMailIdNotFound.md)
 - [GetCharactersCharacterIdMailMailIdOk](docs/GetCharactersCharacterIdMailMailIdOk.md)
 - [GetCharactersCharacterIdMailMailIdRecipient](docs/GetCharactersCharacterIdMailMailIdRecipient.md)
 - [GetCharactersCharacterIdMailRecipient](docs/GetCharactersCharacterIdMailRecipient.md)
 - [GetCharactersCharacterIdMedals200Ok](docs/GetCharactersCharacterIdMedals200Ok.md)
 - [GetCharactersCharacterIdMedalsGraphic](docs/GetCharactersCharacterIdMedalsGraphic.md)
 - [GetCharactersCharacterIdMining200Ok](docs/GetCharactersCharacterIdMining200Ok.md)
 - [GetCharactersCharacterIdNotFound](docs/GetCharactersCharacterIdNotFound.md)
 - [GetCharactersCharacterIdNotifications200Ok](docs/GetCharactersCharacterIdNotifications200Ok.md)
 - [GetCharactersCharacterIdNotificationsContacts200Ok](docs/GetCharactersCharacterIdNotificationsContacts200Ok.md)
 - [GetCharactersCharacterIdOk](docs/GetCharactersCharacterIdOk.md)
 - [GetCharactersCharacterIdOnlineOk](docs/GetCharactersCharacterIdOnlineOk.md)
 - [GetCharactersCharacterIdOpportunities200Ok](docs/GetCharactersCharacterIdOpportunities200Ok.md)
 - [GetCharactersCharacterIdOrders200Ok](docs/GetCharactersCharacterIdOrders200Ok.md)
 - [GetCharactersCharacterIdPlanets200Ok](docs/GetCharactersCharacterIdPlanets200Ok.md)
 - [GetCharactersCharacterIdPlanetsPlanetIdContent](docs/GetCharactersCharacterIdPlanetsPlanetIdContent.md)
 - [GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails](docs/GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails.md)
 - [GetCharactersCharacterIdPlanetsPlanetIdFactoryDetails](docs/GetCharactersCharacterIdPlanetsPlanetIdFactoryDetails.md)
 - [GetCharactersCharacterIdPlanetsPlanetIdHead](docs/GetCharactersCharacterIdPlanetsPlanetIdHead.md)
 - [GetCharactersCharacterIdPlanetsPlanetIdLink](docs/GetCharactersCharacterIdPlanetsPlanetIdLink.md)
 - [GetCharactersCharacterIdPlanetsPlanetIdNotFound](docs/GetCharactersCharacterIdPlanetsPlanetIdNotFound.md)
 - [GetCharactersCharacterIdPlanetsPlanetIdOk](docs/GetCharactersCharacterIdPlanetsPlanetIdOk.md)
 - [GetCharactersCharacterIdPlanetsPlanetIdPin](docs/GetCharactersCharacterIdPlanetsPlanetIdPin.md)
 - [GetCharactersCharacterIdPlanetsPlanetIdRoute](docs/GetCharactersCharacterIdPlanetsPlanetIdRoute.md)
 - [GetCharactersCharacterIdPortraitNotFound](docs/GetCharactersCharacterIdPortraitNotFound.md)
 - [GetCharactersCharacterIdPortraitOk](docs/GetCharactersCharacterIdPortraitOk.md)
 - [GetCharactersCharacterIdRolesOk](docs/GetCharactersCharacterIdRolesOk.md)
 - [GetCharactersCharacterIdSearchOk](docs/GetCharactersCharacterIdSearchOk.md)
 - [GetCharactersCharacterIdShipOk](docs/GetCharactersCharacterIdShipOk.md)
 - [GetCharactersCharacterIdSkillqueue200Ok](docs/GetCharactersCharacterIdSkillqueue200Ok.md)
 - [GetCharactersCharacterIdSkillsOk](docs/GetCharactersCharacterIdSkillsOk.md)
 - [GetCharactersCharacterIdSkillsSkill](docs/GetCharactersCharacterIdSkillsSkill.md)
 - [GetCharactersCharacterIdStandings200Ok](docs/GetCharactersCharacterIdStandings200Ok.md)
 - [GetCharactersCharacterIdStats200Ok](docs/GetCharactersCharacterIdStats200Ok.md)
 - [GetCharactersCharacterIdTitles200Ok](docs/GetCharactersCharacterIdTitles200Ok.md)
 - [GetCharactersCharacterIdWalletJournal200Ok](docs/GetCharactersCharacterIdWalletJournal200Ok.md)
 - [GetCharactersCharacterIdWalletJournalExtraInfo](docs/GetCharactersCharacterIdWalletJournalExtraInfo.md)
 - [GetCharactersCharacterIdWalletTransactions200Ok](docs/GetCharactersCharacterIdWalletTransactions200Ok.md)
 - [GetCharactersNames200Ok](docs/GetCharactersNames200Ok.md)
 - [GetCorporationCorporationIdMiningExtractions200Ok](docs/GetCorporationCorporationIdMiningExtractions200Ok.md)
 - [GetCorporationCorporationIdMiningObservers200Ok](docs/GetCorporationCorporationIdMiningObservers200Ok.md)
 - [GetCorporationCorporationIdMiningObserversObserverId200Ok](docs/GetCorporationCorporationIdMiningObserversObserverId200Ok.md)
 - [GetCorporationsCorporationIdAlliancehistory200Ok](docs/GetCorporationsCorporationIdAlliancehistory200Ok.md)
 - [GetCorporationsCorporationIdAssets200Ok](docs/GetCorporationsCorporationIdAssets200Ok.md)
 - [GetCorporationsCorporationIdBlueprints200Ok](docs/GetCorporationsCorporationIdBlueprints200Ok.md)
 - [GetCorporationsCorporationIdBookmarks200Ok](docs/GetCorporationsCorporationIdBookmarks200Ok.md)
 - [GetCorporationsCorporationIdBookmarksCoordinates](docs/GetCorporationsCorporationIdBookmarksCoordinates.md)
 - [GetCorporationsCorporationIdBookmarksFolders200Ok](docs/GetCorporationsCorporationIdBookmarksFolders200Ok.md)
 - [GetCorporationsCorporationIdBookmarksItem](docs/GetCorporationsCorporationIdBookmarksItem.md)
 - [GetCorporationsCorporationIdContacts200Ok](docs/GetCorporationsCorporationIdContacts200Ok.md)
 - [GetCorporationsCorporationIdContainersLogs200Ok](docs/GetCorporationsCorporationIdContainersLogs200Ok.md)
 - [GetCorporationsCorporationIdContracts200Ok](docs/GetCorporationsCorporationIdContracts200Ok.md)
 - [GetCorporationsCorporationIdContractsContractIdBids200Ok](docs/GetCorporationsCorporationIdContractsContractIdBids200Ok.md)
 - [GetCorporationsCorporationIdContractsContractIdItems200Ok](docs/GetCorporationsCorporationIdContractsContractIdItems200Ok.md)
 - [GetCorporationsCorporationIdCustomsOffices200Ok](docs/GetCorporationsCorporationIdCustomsOffices200Ok.md)
 - [GetCorporationsCorporationIdDivisionsHangar](docs/GetCorporationsCorporationIdDivisionsHangar.md)
 - [GetCorporationsCorporationIdDivisionsOk](docs/GetCorporationsCorporationIdDivisionsOk.md)
 - [GetCorporationsCorporationIdDivisionsWallet](docs/GetCorporationsCorporationIdDivisionsWallet.md)
 - [GetCorporationsCorporationIdFacilities200Ok](docs/GetCorporationsCorporationIdFacilities200Ok.md)
 - [GetCorporationsCorporationIdFwStatsKills](docs/GetCorporationsCorporationIdFwStatsKills.md)
 - [GetCorporationsCorporationIdFwStatsOk](docs/GetCorporationsCorporationIdFwStatsOk.md)
 - [GetCorporationsCorporationIdFwStatsVictoryPoints](docs/GetCorporationsCorporationIdFwStatsVictoryPoints.md)
 - [GetCorporationsCorporationIdIconsNotFound](docs/GetCorporationsCorporationIdIconsNotFound.md)
 - [GetCorporationsCorporationIdIconsOk](docs/GetCorporationsCorporationIdIconsOk.md)
 - [GetCorporationsCorporationIdIndustryJobs200Ok](docs/GetCorporationsCorporationIdIndustryJobs200Ok.md)
 - [GetCorporationsCorporationIdKillmailsRecent200Ok](docs/GetCorporationsCorporationIdKillmailsRecent200Ok.md)
 - [GetCorporationsCorporationIdMedals200Ok](docs/GetCorporationsCorporationIdMedals200Ok.md)
 - [GetCorporationsCorporationIdMedalsIssued200Ok](docs/GetCorporationsCorporationIdMedalsIssued200Ok.md)
 - [GetCorporationsCorporationIdMembersTitles200Ok](docs/GetCorporationsCorporationIdMembersTitles200Ok.md)
 - [GetCorporationsCorporationIdMembertracking200Ok](docs/GetCorporationsCorporationIdMembertracking200Ok.md)
 - [GetCorporationsCorporationIdNotFound](docs/GetCorporationsCorporationIdNotFound.md)
 - [GetCorporationsCorporationIdOk](docs/GetCorporationsCorporationIdOk.md)
 - [GetCorporationsCorporationIdOrders200Ok](docs/GetCorporationsCorporationIdOrders200Ok.md)
 - [GetCorporationsCorporationIdOutpostsOutpostIdCoordinates](docs/GetCorporationsCorporationIdOutpostsOutpostIdCoordinates.md)
 - [GetCorporationsCorporationIdOutpostsOutpostIdOk](docs/GetCorporationsCorporationIdOutpostsOutpostIdOk.md)
 - [GetCorporationsCorporationIdOutpostsOutpostIdService](docs/GetCorporationsCorporationIdOutpostsOutpostIdService.md)
 - [GetCorporationsCorporationIdRoles200Ok](docs/GetCorporationsCorporationIdRoles200Ok.md)
 - [GetCorporationsCorporationIdRolesHistory200Ok](docs/GetCorporationsCorporationIdRolesHistory200Ok.md)
 - [GetCorporationsCorporationIdShareholders200Ok](docs/GetCorporationsCorporationIdShareholders200Ok.md)
 - [GetCorporationsCorporationIdStandings200Ok](docs/GetCorporationsCorporationIdStandings200Ok.md)
 - [GetCorporationsCorporationIdStarbases200Ok](docs/GetCorporationsCorporationIdStarbases200Ok.md)
 - [GetCorporationsCorporationIdStarbasesStarbaseIdFuel](docs/GetCorporationsCorporationIdStarbasesStarbaseIdFuel.md)
 - [GetCorporationsCorporationIdStarbasesStarbaseIdOk](docs/GetCorporationsCorporationIdStarbasesStarbaseIdOk.md)
 - [GetCorporationsCorporationIdStructures200Ok](docs/GetCorporationsCorporationIdStructures200Ok.md)
 - [GetCorporationsCorporationIdStructuresCurrentVul](docs/GetCorporationsCorporationIdStructuresCurrentVul.md)
 - [GetCorporationsCorporationIdStructuresNextVul](docs/GetCorporationsCorporationIdStructuresNextVul.md)
 - [GetCorporationsCorporationIdStructuresService](docs/GetCorporationsCorporationIdStructuresService.md)
 - [GetCorporationsCorporationIdTitles200Ok](docs/GetCorporationsCorporationIdTitles200Ok.md)
 - [GetCorporationsCorporationIdWallets200Ok](docs/GetCorporationsCorporationIdWallets200Ok.md)
 - [GetCorporationsCorporationIdWalletsDivisionJournal200Ok](docs/GetCorporationsCorporationIdWalletsDivisionJournal200Ok.md)
 - [GetCorporationsCorporationIdWalletsDivisionJournalExtraInfo](docs/GetCorporationsCorporationIdWalletsDivisionJournalExtraInfo.md)
 - [GetCorporationsCorporationIdWalletsDivisionTransactions200Ok](docs/GetCorporationsCorporationIdWalletsDivisionTransactions200Ok.md)
 - [GetCorporationsNames200Ok](docs/GetCorporationsNames200Ok.md)
 - [GetDogmaAttributesAttributeIdNotFound](docs/GetDogmaAttributesAttributeIdNotFound.md)
 - [GetDogmaAttributesAttributeIdOk](docs/GetDogmaAttributesAttributeIdOk.md)
 - [GetDogmaEffectsEffectIdModifier](docs/GetDogmaEffectsEffectIdModifier.md)
 - [GetDogmaEffectsEffectIdNotFound](docs/GetDogmaEffectsEffectIdNotFound.md)
 - [GetDogmaEffectsEffectIdOk](docs/GetDogmaEffectsEffectIdOk.md)
 - [GetFleetsFleetIdMembers200Ok](docs/GetFleetsFleetIdMembers200Ok.md)
 - [GetFleetsFleetIdMembersNotFound](docs/GetFleetsFleetIdMembersNotFound.md)
 - [GetFleetsFleetIdNotFound](docs/GetFleetsFleetIdNotFound.md)
 - [GetFleetsFleetIdOk](docs/GetFleetsFleetIdOk.md)
 - [GetFleetsFleetIdWings200Ok](docs/GetFleetsFleetIdWings200Ok.md)
 - [GetFleetsFleetIdWingsNotFound](docs/GetFleetsFleetIdWingsNotFound.md)
 - [GetFleetsFleetIdWingsSquad](docs/GetFleetsFleetIdWingsSquad.md)
 - [GetFwLeaderboardsActiveTotal](docs/GetFwLeaderboardsActiveTotal.md)
 - [GetFwLeaderboardsActiveTotal1](docs/GetFwLeaderboardsActiveTotal1.md)
 - [GetFwLeaderboardsCharactersActiveTotal](docs/GetFwLeaderboardsCharactersActiveTotal.md)
 - [GetFwLeaderboardsCharactersActiveTotal1](docs/GetFwLeaderboardsCharactersActiveTotal1.md)
 - [GetFwLeaderboardsCharactersKills](docs/GetFwLeaderboardsCharactersKills.md)
 - [GetFwLeaderboardsCharactersLastWeek](docs/GetFwLeaderboardsCharactersLastWeek.md)
 - [GetFwLeaderboardsCharactersLastWeek1](docs/GetFwLeaderboardsCharactersLastWeek1.md)
 - [GetFwLeaderboardsCharactersOk](docs/GetFwLeaderboardsCharactersOk.md)
 - [GetFwLeaderboardsCharactersVictoryPoints](docs/GetFwLeaderboardsCharactersVictoryPoints.md)
 - [GetFwLeaderboardsCharactersYesterday](docs/GetFwLeaderboardsCharactersYesterday.md)
 - [GetFwLeaderboardsCharactersYesterday1](docs/GetFwLeaderboardsCharactersYesterday1.md)
 - [GetFwLeaderboardsCorporationsActiveTotal](docs/GetFwLeaderboardsCorporationsActiveTotal.md)
 - [GetFwLeaderboardsCorporationsActiveTotal1](docs/GetFwLeaderboardsCorporationsActiveTotal1.md)
 - [GetFwLeaderboardsCorporationsKills](docs/GetFwLeaderboardsCorporationsKills.md)
 - [GetFwLeaderboardsCorporationsLastWeek](docs/GetFwLeaderboardsCorporationsLastWeek.md)
 - [GetFwLeaderboardsCorporationsLastWeek1](docs/GetFwLeaderboardsCorporationsLastWeek1.md)
 - [GetFwLeaderboardsCorporationsOk](docs/GetFwLeaderboardsCorporationsOk.md)
 - [GetFwLeaderboardsCorporationsVictoryPoints](docs/GetFwLeaderboardsCorporationsVictoryPoints.md)
 - [GetFwLeaderboardsCorporationsYesterday](docs/GetFwLeaderboardsCorporationsYesterday.md)
 - [GetFwLeaderboardsCorporationsYesterday1](docs/GetFwLeaderboardsCorporationsYesterday1.md)
 - [GetFwLeaderboardsKills](docs/GetFwLeaderboardsKills.md)
 - [GetFwLeaderboardsLastWeek](docs/GetFwLeaderboardsLastWeek.md)
 - [GetFwLeaderboardsLastWeek1](docs/GetFwLeaderboardsLastWeek1.md)
 - [GetFwLeaderboardsOk](docs/GetFwLeaderboardsOk.md)
 - [GetFwLeaderboardsVictoryPoints](docs/GetFwLeaderboardsVictoryPoints.md)
 - [GetFwLeaderboardsYesterday](docs/GetFwLeaderboardsYesterday.md)
 - [GetFwLeaderboardsYesterday1](docs/GetFwLeaderboardsYesterday1.md)
 - [GetFwStats200Ok](docs/GetFwStats200Ok.md)
 - [GetFwStatsKills](docs/GetFwStatsKills.md)
 - [GetFwStatsVictoryPoints](docs/GetFwStatsVictoryPoints.md)
 - [GetFwSystems200Ok](docs/GetFwSystems200Ok.md)
 - [GetFwWars200Ok](docs/GetFwWars200Ok.md)
 - [GetIncursions200Ok](docs/GetIncursions200Ok.md)
 - [GetIndustryFacilities200Ok](docs/GetIndustryFacilities200Ok.md)
 - [GetIndustrySystems200Ok](docs/GetIndustrySystems200Ok.md)
 - [GetIndustrySystemsCostIndice](docs/GetIndustrySystemsCostIndice.md)
 - [GetInsurancePrices200Ok](docs/GetInsurancePrices200Ok.md)
 - [GetInsurancePricesLevel](docs/GetInsurancePricesLevel.md)
 - [GetKillmailsKillmailIdKillmailHashAttacker](docs/GetKillmailsKillmailIdKillmailHashAttacker.md)
 - [GetKillmailsKillmailIdKillmailHashItem](docs/GetKillmailsKillmailIdKillmailHashItem.md)
 - [GetKillmailsKillmailIdKillmailHashItem1](docs/GetKillmailsKillmailIdKillmailHashItem1.md)
 - [GetKillmailsKillmailIdKillmailHashOk](docs/GetKillmailsKillmailIdKillmailHashOk.md)
 - [GetKillmailsKillmailIdKillmailHashPosition](docs/GetKillmailsKillmailIdKillmailHashPosition.md)
 - [GetKillmailsKillmailIdKillmailHashUnprocessableEntity](docs/GetKillmailsKillmailIdKillmailHashUnprocessableEntity.md)
 - [GetKillmailsKillmailIdKillmailHashVictim](docs/GetKillmailsKillmailIdKillmailHashVictim.md)
 - [GetLoyaltyStoresCorporationIdOffers200Ok](docs/GetLoyaltyStoresCorporationIdOffers200Ok.md)
 - [GetLoyaltyStoresCorporationIdOffersRequiredItem](docs/GetLoyaltyStoresCorporationIdOffersRequiredItem.md)
 - [GetMarketsGroupsMarketGroupIdNotFound](docs/GetMarketsGroupsMarketGroupIdNotFound.md)
 - [GetMarketsGroupsMarketGroupIdOk](docs/GetMarketsGroupsMarketGroupIdOk.md)
 - [GetMarketsPrices200Ok](docs/GetMarketsPrices200Ok.md)
 - [GetMarketsRegionIdHistory200Ok](docs/GetMarketsRegionIdHistory200Ok.md)
 - [GetMarketsRegionIdHistoryUnprocessableEntity](docs/GetMarketsRegionIdHistoryUnprocessableEntity.md)
 - [GetMarketsRegionIdOrders200Ok](docs/GetMarketsRegionIdOrders200Ok.md)
 - [GetMarketsRegionIdOrdersUnprocessableEntity](docs/GetMarketsRegionIdOrdersUnprocessableEntity.md)
 - [GetMarketsStructuresStructureId200Ok](docs/GetMarketsStructuresStructureId200Ok.md)
 - [GetOpportunitiesGroupsGroupIdOk](docs/GetOpportunitiesGroupsGroupIdOk.md)
 - [GetOpportunitiesTasksTaskIdOk](docs/GetOpportunitiesTasksTaskIdOk.md)
 - [GetRouteOriginDestinationNotFound](docs/GetRouteOriginDestinationNotFound.md)
 - [GetSearchOk](docs/GetSearchOk.md)
 - [GetSovereigntyCampaigns200Ok](docs/GetSovereigntyCampaigns200Ok.md)
 - [GetSovereigntyCampaignsParticipant](docs/GetSovereigntyCampaignsParticipant.md)
 - [GetSovereigntyMap200Ok](docs/GetSovereigntyMap200Ok.md)
 - [GetSovereigntyStructures200Ok](docs/GetSovereigntyStructures200Ok.md)
 - [GetStatusOk](docs/GetStatusOk.md)
 - [GetUniverseBloodlines200Ok](docs/GetUniverseBloodlines200Ok.md)
 - [GetUniverseCategoriesCategoryIdNotFound](docs/GetUniverseCategoriesCategoryIdNotFound.md)
 - [GetUniverseCategoriesCategoryIdOk](docs/GetUniverseCategoriesCategoryIdOk.md)
 - [GetUniverseConstellationsConstellationIdNotFound](docs/GetUniverseConstellationsConstellationIdNotFound.md)
 - [GetUniverseConstellationsConstellationIdOk](docs/GetUniverseConstellationsConstellationIdOk.md)
 - [GetUniverseConstellationsConstellationIdPosition](docs/GetUniverseConstellationsConstellationIdPosition.md)
 - [GetUniverseFactions200Ok](docs/GetUniverseFactions200Ok.md)
 - [GetUniverseGraphicsGraphicIdNotFound](docs/GetUniverseGraphicsGraphicIdNotFound.md)
 - [GetUniverseGraphicsGraphicIdOk](docs/GetUniverseGraphicsGraphicIdOk.md)
 - [GetUniverseGroupsGroupIdNotFound](docs/GetUniverseGroupsGroupIdNotFound.md)
 - [GetUniverseGroupsGroupIdOk](docs/GetUniverseGroupsGroupIdOk.md)
 - [GetUniverseMoonsMoonIdNotFound](docs/GetUniverseMoonsMoonIdNotFound.md)
 - [GetUniverseMoonsMoonIdOk](docs/GetUniverseMoonsMoonIdOk.md)
 - [GetUniverseMoonsMoonIdPosition](docs/GetUniverseMoonsMoonIdPosition.md)
 - [GetUniversePlanetsPlanetIdNotFound](docs/GetUniversePlanetsPlanetIdNotFound.md)
 - [GetUniversePlanetsPlanetIdOk](docs/GetUniversePlanetsPlanetIdOk.md)
 - [GetUniversePlanetsPlanetIdPosition](docs/GetUniversePlanetsPlanetIdPosition.md)
 - [GetUniverseRaces200Ok](docs/GetUniverseRaces200Ok.md)
 - [GetUniverseRegionsRegionIdNotFound](docs/GetUniverseRegionsRegionIdNotFound.md)
 - [GetUniverseRegionsRegionIdOk](docs/GetUniverseRegionsRegionIdOk.md)
 - [GetUniverseSchematicsSchematicIdNotFound](docs/GetUniverseSchematicsSchematicIdNotFound.md)
 - [GetUniverseSchematicsSchematicIdOk](docs/GetUniverseSchematicsSchematicIdOk.md)
 - [GetUniverseStargatesStargateIdDestination](docs/GetUniverseStargatesStargateIdDestination.md)
 - [GetUniverseStargatesStargateIdNotFound](docs/GetUniverseStargatesStargateIdNotFound.md)
 - [GetUniverseStargatesStargateIdOk](docs/GetUniverseStargatesStargateIdOk.md)
 - [GetUniverseStargatesStargateIdPosition](docs/GetUniverseStargatesStargateIdPosition.md)
 - [GetUniverseStarsStarIdOk](docs/GetUniverseStarsStarIdOk.md)
 - [GetUniverseStationsStationIdNotFound](docs/GetUniverseStationsStationIdNotFound.md)
 - [GetUniverseStationsStationIdOk](docs/GetUniverseStationsStationIdOk.md)
 - [GetUniverseStationsStationIdPosition](docs/GetUniverseStationsStationIdPosition.md)
 - [GetUniverseStructuresStructureIdNotFound](docs/GetUniverseStructuresStructureIdNotFound.md)
 - [GetUniverseStructuresStructureIdOk](docs/GetUniverseStructuresStructureIdOk.md)
 - [GetUniverseStructuresStructureIdPosition](docs/GetUniverseStructuresStructureIdPosition.md)
 - [GetUniverseSystemJumps200Ok](docs/GetUniverseSystemJumps200Ok.md)
 - [GetUniverseSystemKills200Ok](docs/GetUniverseSystemKills200Ok.md)
 - [GetUniverseSystemsSystemIdNotFound](docs/GetUniverseSystemsSystemIdNotFound.md)
 - [GetUniverseSystemsSystemIdOk](docs/GetUniverseSystemsSystemIdOk.md)
 - [GetUniverseSystemsSystemIdPlanet](docs/GetUniverseSystemsSystemIdPlanet.md)
 - [GetUniverseSystemsSystemIdPosition](docs/GetUniverseSystemsSystemIdPosition.md)
 - [GetUniverseTypesTypeIdDogmaAttribute](docs/GetUniverseTypesTypeIdDogmaAttribute.md)
 - [GetUniverseTypesTypeIdDogmaEffect](docs/GetUniverseTypesTypeIdDogmaEffect.md)
 - [GetUniverseTypesTypeIdNotFound](docs/GetUniverseTypesTypeIdNotFound.md)
 - [GetUniverseTypesTypeIdOk](docs/GetUniverseTypesTypeIdOk.md)
 - [GetWarsWarIdAggressor](docs/GetWarsWarIdAggressor.md)
 - [GetWarsWarIdAlly](docs/GetWarsWarIdAlly.md)
 - [GetWarsWarIdDefender](docs/GetWarsWarIdDefender.md)
 - [GetWarsWarIdKillmails200Ok](docs/GetWarsWarIdKillmails200Ok.md)
 - [GetWarsWarIdKillmailsUnprocessableEntity](docs/GetWarsWarIdKillmailsUnprocessableEntity.md)
 - [GetWarsWarIdOk](docs/GetWarsWarIdOk.md)
 - [GetWarsWarIdUnprocessableEntity](docs/GetWarsWarIdUnprocessableEntity.md)
 - [InternalServerError](docs/InternalServerError.md)
 - [PostCharactersAffiliation200Ok](docs/PostCharactersAffiliation200Ok.md)
 - [PostCharactersAffiliationNotFound](docs/PostCharactersAffiliationNotFound.md)
 - [PostCharactersCharacterIdAssetsLocations200Ok](docs/PostCharactersCharacterIdAssetsLocations200Ok.md)
 - [PostCharactersCharacterIdAssetsLocationsPosition](docs/PostCharactersCharacterIdAssetsLocationsPosition.md)
 - [PostCharactersCharacterIdAssetsNames200Ok](docs/PostCharactersCharacterIdAssetsNames200Ok.md)
 - [PostCharactersCharacterIdFittingsCreated](docs/PostCharactersCharacterIdFittingsCreated.md)
 - [PostCharactersCharacterIdFittingsFitting](docs/PostCharactersCharacterIdFittingsFitting.md)
 - [PostCharactersCharacterIdFittingsItem](docs/PostCharactersCharacterIdFittingsItem.md)
 - [PostCharactersCharacterIdMailBadRequest](docs/PostCharactersCharacterIdMailBadRequest.md)
 - [PostCharactersCharacterIdMailLabelsLabel](docs/PostCharactersCharacterIdMailLabelsLabel.md)
 - [PostCharactersCharacterIdMailMail](docs/PostCharactersCharacterIdMailMail.md)
 - [PostCharactersCharacterIdMailRecipient](docs/PostCharactersCharacterIdMailRecipient.md)
 - [PostCorporationsCorporationIdAssetsLocations200Ok](docs/PostCorporationsCorporationIdAssetsLocations200Ok.md)
 - [PostCorporationsCorporationIdAssetsLocationsPosition](docs/PostCorporationsCorporationIdAssetsLocationsPosition.md)
 - [PostCorporationsCorporationIdAssetsNames200Ok](docs/PostCorporationsCorporationIdAssetsNames200Ok.md)
 - [PostFleetsFleetIdMembersInvitation](docs/PostFleetsFleetIdMembersInvitation.md)
 - [PostFleetsFleetIdMembersNotFound](docs/PostFleetsFleetIdMembersNotFound.md)
 - [PostFleetsFleetIdMembersUnprocessableEntity](docs/PostFleetsFleetIdMembersUnprocessableEntity.md)
 - [PostFleetsFleetIdWingsCreated](docs/PostFleetsFleetIdWingsCreated.md)
 - [PostFleetsFleetIdWingsNotFound](docs/PostFleetsFleetIdWingsNotFound.md)
 - [PostFleetsFleetIdWingsWingIdSquadsCreated](docs/PostFleetsFleetIdWingsWingIdSquadsCreated.md)
 - [PostFleetsFleetIdWingsWingIdSquadsNotFound](docs/PostFleetsFleetIdWingsWingIdSquadsNotFound.md)
 - [PostUiOpenwindowNewmailNewMail](docs/PostUiOpenwindowNewmailNewMail.md)
 - [PostUiOpenwindowNewmailUnprocessableEntity](docs/PostUiOpenwindowNewmailUnprocessableEntity.md)
 - [PostUniverseIdsAgent](docs/PostUniverseIdsAgent.md)
 - [PostUniverseIdsAlliance](docs/PostUniverseIdsAlliance.md)
 - [PostUniverseIdsCharacter](docs/PostUniverseIdsCharacter.md)
 - [PostUniverseIdsConstellation](docs/PostUniverseIdsConstellation.md)
 - [PostUniverseIdsCorporation](docs/PostUniverseIdsCorporation.md)
 - [PostUniverseIdsFaction](docs/PostUniverseIdsFaction.md)
 - [PostUniverseIdsInventoryType](docs/PostUniverseIdsInventoryType.md)
 - [PostUniverseIdsOk](docs/PostUniverseIdsOk.md)
 - [PostUniverseIdsRegion](docs/PostUniverseIdsRegion.md)
 - [PostUniverseIdsServiceUnavailable](docs/PostUniverseIdsServiceUnavailable.md)
 - [PostUniverseIdsStation](docs/PostUniverseIdsStation.md)
 - [PostUniverseIdsSystem](docs/PostUniverseIdsSystem.md)
 - [PostUniverseNames200Ok](docs/PostUniverseNames200Ok.md)
 - [PostUniverseNamesNotFound](docs/PostUniverseNamesNotFound.md)
 - [PutCharactersCharacterIdCalendarEventIdResponse](docs/PutCharactersCharacterIdCalendarEventIdResponse.md)
 - [PutCharactersCharacterIdMailMailIdBadRequest](docs/PutCharactersCharacterIdMailMailIdBadRequest.md)
 - [PutCharactersCharacterIdMailMailIdContents](docs/PutCharactersCharacterIdMailMailIdContents.md)
 - [PutCorporationsCorporationIdStructuresStructureIdNewSchedule](docs/PutCorporationsCorporationIdStructuresStructureIdNewSchedule.md)
 - [PutFleetsFleetIdBadRequest](docs/PutFleetsFleetIdBadRequest.md)
 - [PutFleetsFleetIdMembersMemberIdMovement](docs/PutFleetsFleetIdMembersMemberIdMovement.md)
 - [PutFleetsFleetIdMembersMemberIdNotFound](docs/PutFleetsFleetIdMembersMemberIdNotFound.md)
 - [PutFleetsFleetIdMembersMemberIdUnprocessableEntity](docs/PutFleetsFleetIdMembersMemberIdUnprocessableEntity.md)
 - [PutFleetsFleetIdNewSettings](docs/PutFleetsFleetIdNewSettings.md)
 - [PutFleetsFleetIdNotFound](docs/PutFleetsFleetIdNotFound.md)
 - [PutFleetsFleetIdSquadsSquadIdNaming](docs/PutFleetsFleetIdSquadsSquadIdNaming.md)
 - [PutFleetsFleetIdSquadsSquadIdNotFound](docs/PutFleetsFleetIdSquadsSquadIdNotFound.md)
 - [PutFleetsFleetIdWingsWingIdNaming](docs/PutFleetsFleetIdWingsWingIdNaming.md)
 - [PutFleetsFleetIdWingsWingIdNotFound](docs/PutFleetsFleetIdWingsWingIdNotFound.md)


## Documentation For Authorization


## evesso

- **Type**: OAuth
- **Flow**: implicit
- **Authorization URL**: https://login.eveonline.com/oauth/authorize
- **Scopes**: 
 - **esi-alliances.read_contacts.v1**: EVE SSO scope esi-alliances.read_contacts.v1
 - **esi-assets.read_assets.v1**: EVE SSO scope esi-assets.read_assets.v1
 - **esi-assets.read_corporation_assets.v1**: EVE SSO scope esi-assets.read_corporation_assets.v1
 - **esi-bookmarks.read_character_bookmarks.v1**: EVE SSO scope esi-bookmarks.read_character_bookmarks.v1
 - **esi-bookmarks.read_corporation_bookmarks.v1**: EVE SSO scope esi-bookmarks.read_corporation_bookmarks.v1
 - **esi-calendar.read_calendar_events.v1**: EVE SSO scope esi-calendar.read_calendar_events.v1
 - **esi-calendar.respond_calendar_events.v1**: EVE SSO scope esi-calendar.respond_calendar_events.v1
 - **esi-characters.read_agents_research.v1**: EVE SSO scope esi-characters.read_agents_research.v1
 - **esi-characters.read_blueprints.v1**: EVE SSO scope esi-characters.read_blueprints.v1
 - **esi-characters.read_chat_channels.v1**: EVE SSO scope esi-characters.read_chat_channels.v1
 - **esi-characters.read_contacts.v1**: EVE SSO scope esi-characters.read_contacts.v1
 - **esi-characters.read_corporation_roles.v1**: EVE SSO scope esi-characters.read_corporation_roles.v1
 - **esi-characters.read_fatigue.v1**: EVE SSO scope esi-characters.read_fatigue.v1
 - **esi-characters.read_fw_stats.v1**: EVE SSO scope esi-characters.read_fw_stats.v1
 - **esi-characters.read_loyalty.v1**: EVE SSO scope esi-characters.read_loyalty.v1
 - **esi-characters.read_medals.v1**: EVE SSO scope esi-characters.read_medals.v1
 - **esi-characters.read_notifications.v1**: EVE SSO scope esi-characters.read_notifications.v1
 - **esi-characters.read_opportunities.v1**: EVE SSO scope esi-characters.read_opportunities.v1
 - **esi-characters.read_standings.v1**: EVE SSO scope esi-characters.read_standings.v1
 - **esi-characters.read_titles.v1**: EVE SSO scope esi-characters.read_titles.v1
 - **esi-characters.write_contacts.v1**: EVE SSO scope esi-characters.write_contacts.v1
 - **esi-characterstats.read.v1**: EVE SSO scope esi-characterstats.read.v1
 - **esi-clones.read_clones.v1**: EVE SSO scope esi-clones.read_clones.v1
 - **esi-clones.read_implants.v1**: EVE SSO scope esi-clones.read_implants.v1
 - **esi-contracts.read_character_contracts.v1**: EVE SSO scope esi-contracts.read_character_contracts.v1
 - **esi-contracts.read_corporation_contracts.v1**: EVE SSO scope esi-contracts.read_corporation_contracts.v1
 - **esi-corporations.read_blueprints.v1**: EVE SSO scope esi-corporations.read_blueprints.v1
 - **esi-corporations.read_contacts.v1**: EVE SSO scope esi-corporations.read_contacts.v1
 - **esi-corporations.read_container_logs.v1**: EVE SSO scope esi-corporations.read_container_logs.v1
 - **esi-corporations.read_corporation_membership.v1**: EVE SSO scope esi-corporations.read_corporation_membership.v1
 - **esi-corporations.read_divisions.v1**: EVE SSO scope esi-corporations.read_divisions.v1
 - **esi-corporations.read_facilities.v1**: EVE SSO scope esi-corporations.read_facilities.v1
 - **esi-corporations.read_fw_stats.v1**: EVE SSO scope esi-corporations.read_fw_stats.v1
 - **esi-corporations.read_medals.v1**: EVE SSO scope esi-corporations.read_medals.v1
 - **esi-corporations.read_outposts.v1**: EVE SSO scope esi-corporations.read_outposts.v1
 - **esi-corporations.read_standings.v1**: EVE SSO scope esi-corporations.read_standings.v1
 - **esi-corporations.read_starbases.v1**: EVE SSO scope esi-corporations.read_starbases.v1
 - **esi-corporations.read_structures.v1**: EVE SSO scope esi-corporations.read_structures.v1
 - **esi-corporations.read_titles.v1**: EVE SSO scope esi-corporations.read_titles.v1
 - **esi-corporations.track_members.v1**: EVE SSO scope esi-corporations.track_members.v1
 - **esi-corporations.write_structures.v1**: EVE SSO scope esi-corporations.write_structures.v1
 - **esi-fittings.read_fittings.v1**: EVE SSO scope esi-fittings.read_fittings.v1
 - **esi-fittings.write_fittings.v1**: EVE SSO scope esi-fittings.write_fittings.v1
 - **esi-fleets.read_fleet.v1**: EVE SSO scope esi-fleets.read_fleet.v1
 - **esi-fleets.write_fleet.v1**: EVE SSO scope esi-fleets.write_fleet.v1
 - **esi-industry.read_character_jobs.v1**: EVE SSO scope esi-industry.read_character_jobs.v1
 - **esi-industry.read_character_mining.v1**: EVE SSO scope esi-industry.read_character_mining.v1
 - **esi-industry.read_corporation_jobs.v1**: EVE SSO scope esi-industry.read_corporation_jobs.v1
 - **esi-industry.read_corporation_mining.v1**: EVE SSO scope esi-industry.read_corporation_mining.v1
 - **esi-killmails.read_corporation_killmails.v1**: EVE SSO scope esi-killmails.read_corporation_killmails.v1
 - **esi-killmails.read_killmails.v1**: EVE SSO scope esi-killmails.read_killmails.v1
 - **esi-location.read_location.v1**: EVE SSO scope esi-location.read_location.v1
 - **esi-location.read_online.v1**: EVE SSO scope esi-location.read_online.v1
 - **esi-location.read_ship_type.v1**: EVE SSO scope esi-location.read_ship_type.v1
 - **esi-mail.organize_mail.v1**: EVE SSO scope esi-mail.organize_mail.v1
 - **esi-mail.read_mail.v1**: EVE SSO scope esi-mail.read_mail.v1
 - **esi-mail.send_mail.v1**: EVE SSO scope esi-mail.send_mail.v1
 - **esi-markets.read_character_orders.v1**: EVE SSO scope esi-markets.read_character_orders.v1
 - **esi-markets.read_corporation_orders.v1**: EVE SSO scope esi-markets.read_corporation_orders.v1
 - **esi-markets.structure_markets.v1**: EVE SSO scope esi-markets.structure_markets.v1
 - **esi-planets.manage_planets.v1**: EVE SSO scope esi-planets.manage_planets.v1
 - **esi-planets.read_customs_offices.v1**: EVE SSO scope esi-planets.read_customs_offices.v1
 - **esi-search.search_structures.v1**: EVE SSO scope esi-search.search_structures.v1
 - **esi-skills.read_skillqueue.v1**: EVE SSO scope esi-skills.read_skillqueue.v1
 - **esi-skills.read_skills.v1**: EVE SSO scope esi-skills.read_skills.v1
 - **esi-ui.open_window.v1**: EVE SSO scope esi-ui.open_window.v1
 - **esi-ui.write_waypoint.v1**: EVE SSO scope esi-ui.write_waypoint.v1
 - **esi-universe.read_structures.v1**: EVE SSO scope esi-universe.read_structures.v1
 - **esi-wallet.read_character_wallet.v1**: EVE SSO scope esi-wallet.read_character_wallet.v1
 - **esi-wallet.read_corporation_wallets.v1**: EVE SSO scope esi-wallet.read_corporation_wallets.v1

Example:
```
  auth, err = oauth2conf.TokenSource(http.Client, token) // Access and Refresh token structure

  client, err := esi.NewClient(nil, "esi client http://mysite.com")
  ctx := context.WithValue(context.TODO(), esi.ContextOAuth2, auth)
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


