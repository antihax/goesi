# Go API client for goesiv1

## Documentation for API Endpoints

All URIs are relative to *https://esi.tech.ccp.is/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AllianceApi* | [**GetAlliances**](docs/AllianceApi.md#getalliances) | **Get** /alliances/ | List all alliances
*AllianceApi* | [**GetAlliancesAllianceId**](docs/AllianceApi.md#getalliancesallianceid) | **Get** /alliances/{alliance_id}/ | Get alliance information
*AllianceApi* | [**GetAlliancesAllianceIdCorporations**](docs/AllianceApi.md#getalliancesallianceidcorporations) | **Get** /alliances/{alliance_id}/corporations/ | List alliance&#39;s corporations
*AllianceApi* | [**GetAlliancesAllianceIdIcons**](docs/AllianceApi.md#getalliancesallianceidicons) | **Get** /alliances/{alliance_id}/icons/ | Get alliance icon
*AllianceApi* | [**GetAlliancesNames**](docs/AllianceApi.md#getalliancesnames) | **Get** /alliances/names/ | Get alliance names
*AssetsApi* | [**GetCharactersCharacterIdAssets**](docs/AssetsApi.md#getcharacterscharacteridassets) | **Get** /characters/{character_id}/assets/ | Get character assets
*BookmarksApi* | [**GetCharactersCharacterIdBookmarks**](docs/BookmarksApi.md#getcharacterscharacteridbookmarks) | **Get** /characters/{character_id}/bookmarks/ | List bookmarks
*BookmarksApi* | [**GetCharactersCharacterIdBookmarksFolders**](docs/BookmarksApi.md#getcharacterscharacteridbookmarksfolders) | **Get** /characters/{character_id}/bookmarks/folders/ | List bookmark folders
*CalendarApi* | [**GetCharactersCharacterIdCalendar**](docs/CalendarApi.md#getcharacterscharacteridcalendar) | **Get** /characters/{character_id}/calendar/ | List calendar event summaries
*CharacterApi* | [**GetCharactersCharacterIdCorporationhistory**](docs/CharacterApi.md#getcharacterscharacteridcorporationhistory) | **Get** /characters/{character_id}/corporationhistory/ | Get corporation history
*CharacterApi* | [**GetCharactersCharacterIdPortrait**](docs/CharacterApi.md#getcharacterscharacteridportrait) | **Get** /characters/{character_id}/portrait/ | Get character portraits
*CharacterApi* | [**GetCharactersNames**](docs/CharacterApi.md#getcharactersnames) | **Get** /characters/names/ | Get character names
*ClonesApi* | [**GetCharactersCharacterIdClones**](docs/ClonesApi.md#getcharacterscharacteridclones) | **Get** /characters/{character_id}/clones/ | Get clones
*ContactsApi* | [**DeleteCharactersCharacterIdContacts**](docs/ContactsApi.md#deletecharacterscharacteridcontacts) | **Delete** /characters/{character_id}/contacts/ | Delete contacts
*ContactsApi* | [**GetCharactersCharacterIdContacts**](docs/ContactsApi.md#getcharacterscharacteridcontacts) | **Get** /characters/{character_id}/contacts/ | Get contacts
*ContactsApi* | [**GetCharactersCharacterIdContactsLabels**](docs/ContactsApi.md#getcharacterscharacteridcontactslabels) | **Get** /characters/{character_id}/contacts/labels/ | Get contact labels
*ContactsApi* | [**PostCharactersCharacterIdContacts**](docs/ContactsApi.md#postcharacterscharacteridcontacts) | **Post** /characters/{character_id}/contacts/ | Add contacts
*ContactsApi* | [**PutCharactersCharacterIdContacts**](docs/ContactsApi.md#putcharacterscharacteridcontacts) | **Put** /characters/{character_id}/contacts/ | Edit contacts
*CorporationApi* | [**GetCorporationsCorporationIdAlliancehistory**](docs/CorporationApi.md#getcorporationscorporationidalliancehistory) | **Get** /corporations/{corporation_id}/alliancehistory/ | Get alliance history
*CorporationApi* | [**GetCorporationsCorporationIdIcons**](docs/CorporationApi.md#getcorporationscorporationidicons) | **Get** /corporations/{corporation_id}/icons/ | Get corporation icon
*CorporationApi* | [**GetCorporationsCorporationIdRoles**](docs/CorporationApi.md#getcorporationscorporationidroles) | **Get** /corporations/{corporation_id}/roles/ | Get corporation member roles
*CorporationApi* | [**GetCorporationsCorporationIdStructures**](docs/CorporationApi.md#getcorporationscorporationidstructures) | **Get** /corporations/{corporation_id}/structures/ | Get corporation structures
*CorporationApi* | [**GetCorporationsNames**](docs/CorporationApi.md#getcorporationsnames) | **Get** /corporations/names/ | Get corporation names
*CorporationApi* | [**GetCorporationsNpccorps**](docs/CorporationApi.md#getcorporationsnpccorps) | **Get** /corporations/npccorps/ | Get npc corporations
*CorporationApi* | [**PutCorporationsCorporationIdStructuresStructureId**](docs/CorporationApi.md#putcorporationscorporationidstructuresstructureid) | **Put** /corporations/{corporation_id}/structures/{structure_id}/ | Update structure vulnerability schedule
*DogmaApi* | [**GetDogmaAttributes**](docs/DogmaApi.md#getdogmaattributes) | **Get** /dogma/attributes/ | Get attributes
*DogmaApi* | [**GetDogmaAttributesAttributeId**](docs/DogmaApi.md#getdogmaattributesattributeid) | **Get** /dogma/attributes/{attribute_id}/ | Get attribute information
*DogmaApi* | [**GetDogmaEffects**](docs/DogmaApi.md#getdogmaeffects) | **Get** /dogma/effects/ | Get effects
*DogmaApi* | [**GetDogmaEffectsEffectId**](docs/DogmaApi.md#getdogmaeffectseffectid) | **Get** /dogma/effects/{effect_id}/ | Get effect information
*FittingsApi* | [**DeleteCharactersCharacterIdFittingsFittingId**](docs/FittingsApi.md#deletecharacterscharacteridfittingsfittingid) | **Delete** /characters/{character_id}/fittings/{fitting_id}/ | Delete fitting
*FittingsApi* | [**GetCharactersCharacterIdFittings**](docs/FittingsApi.md#getcharacterscharacteridfittings) | **Get** /characters/{character_id}/fittings/ | Get fittings
*FittingsApi* | [**PostCharactersCharacterIdFittings**](docs/FittingsApi.md#postcharacterscharacteridfittings) | **Post** /characters/{character_id}/fittings/ | Create fitting
*FleetsApi* | [**DeleteFleetsFleetIdMembersMemberId**](docs/FleetsApi.md#deletefleetsfleetidmembersmemberid) | **Delete** /fleets/{fleet_id}/members/{member_id}/ | Kick fleet member
*FleetsApi* | [**DeleteFleetsFleetIdSquadsSquadId**](docs/FleetsApi.md#deletefleetsfleetidsquadssquadid) | **Delete** /fleets/{fleet_id}/squads/{squad_id}/ | Delete fleet squad
*FleetsApi* | [**DeleteFleetsFleetIdWingsWingId**](docs/FleetsApi.md#deletefleetsfleetidwingswingid) | **Delete** /fleets/{fleet_id}/wings/{wing_id}/ | Delete fleet wing
*FleetsApi* | [**GetFleetsFleetId**](docs/FleetsApi.md#getfleetsfleetid) | **Get** /fleets/{fleet_id}/ | Get fleet information
*FleetsApi* | [**GetFleetsFleetIdMembers**](docs/FleetsApi.md#getfleetsfleetidmembers) | **Get** /fleets/{fleet_id}/members/ | Get fleet members
*FleetsApi* | [**GetFleetsFleetIdWings**](docs/FleetsApi.md#getfleetsfleetidwings) | **Get** /fleets/{fleet_id}/wings/ | Get fleet wings
*FleetsApi* | [**PostFleetsFleetIdMembers**](docs/FleetsApi.md#postfleetsfleetidmembers) | **Post** /fleets/{fleet_id}/members/ | Create fleet invitation
*FleetsApi* | [**PostFleetsFleetIdWings**](docs/FleetsApi.md#postfleetsfleetidwings) | **Post** /fleets/{fleet_id}/wings/ | Create fleet wing
*FleetsApi* | [**PostFleetsFleetIdWingsWingIdSquads**](docs/FleetsApi.md#postfleetsfleetidwingswingidsquads) | **Post** /fleets/{fleet_id}/wings/{wing_id}/squads/ | Create fleet squad
*FleetsApi* | [**PutFleetsFleetId**](docs/FleetsApi.md#putfleetsfleetid) | **Put** /fleets/{fleet_id}/ | Update fleet
*FleetsApi* | [**PutFleetsFleetIdMembersMemberId**](docs/FleetsApi.md#putfleetsfleetidmembersmemberid) | **Put** /fleets/{fleet_id}/members/{member_id}/ | Move fleet member
*FleetsApi* | [**PutFleetsFleetIdSquadsSquadId**](docs/FleetsApi.md#putfleetsfleetidsquadssquadid) | **Put** /fleets/{fleet_id}/squads/{squad_id}/ | Rename fleet squad
*FleetsApi* | [**PutFleetsFleetIdWingsWingId**](docs/FleetsApi.md#putfleetsfleetidwingswingid) | **Put** /fleets/{fleet_id}/wings/{wing_id}/ | Rename fleet wing
*IncursionsApi* | [**GetIncursions**](docs/IncursionsApi.md#getincursions) | **Get** /incursions/ | List incursions
*IndustryApi* | [**GetIndustryFacilities**](docs/IndustryApi.md#getindustryfacilities) | **Get** /industry/facilities/ | List industry facilities
*IndustryApi* | [**GetIndustrySystems**](docs/IndustryApi.md#getindustrysystems) | **Get** /industry/systems/ | List solar system cost indices
*InsuranceApi* | [**GetInsurancePrices**](docs/InsuranceApi.md#getinsuranceprices) | **Get** /insurance/prices/ | List insurance levels
*KillmailsApi* | [**GetCharactersCharacterIdKillmailsRecent**](docs/KillmailsApi.md#getcharacterscharacteridkillmailsrecent) | **Get** /characters/{character_id}/killmails/recent/ | List kills and losses
*KillmailsApi* | [**GetKillmailsKillmailIdKillmailHash**](docs/KillmailsApi.md#getkillmailskillmailidkillmailhash) | **Get** /killmails/{killmail_id}/{killmail_hash}/ | Get a single killmail
*LocationApi* | [**GetCharactersCharacterIdLocation**](docs/LocationApi.md#getcharacterscharacteridlocation) | **Get** /characters/{character_id}/location/ | Get character location
*LocationApi* | [**GetCharactersCharacterIdShip**](docs/LocationApi.md#getcharacterscharacteridship) | **Get** /characters/{character_id}/ship/ | Get current ship
*LoyaltyApi* | [**GetCharactersCharacterIdLoyaltyPoints**](docs/LoyaltyApi.md#getcharacterscharacteridloyaltypoints) | **Get** /characters/{character_id}/loyalty/points/ | Get loyalty points
*LoyaltyApi* | [**GetLoyaltyStoresCorporationIdOffers**](docs/LoyaltyApi.md#getloyaltystorescorporationidoffers) | **Get** /loyalty/stores/{corporation_id}/offers/ | List loyalty store offers
*MailApi* | [**DeleteCharactersCharacterIdMailLabelsLabelId**](docs/MailApi.md#deletecharacterscharacteridmaillabelslabelid) | **Delete** /characters/{character_id}/mail/labels/{label_id}/ | Delete a mail label
*MailApi* | [**DeleteCharactersCharacterIdMailMailId**](docs/MailApi.md#deletecharacterscharacteridmailmailid) | **Delete** /characters/{character_id}/mail/{mail_id}/ | Delete a mail
*MailApi* | [**GetCharactersCharacterIdMail**](docs/MailApi.md#getcharacterscharacteridmail) | **Get** /characters/{character_id}/mail/ | Return mail headers
*MailApi* | [**GetCharactersCharacterIdMailLists**](docs/MailApi.md#getcharacterscharacteridmaillists) | **Get** /characters/{character_id}/mail/lists/ | Return mailing list subscriptions
*MailApi* | [**GetCharactersCharacterIdMailMailId**](docs/MailApi.md#getcharacterscharacteridmailmailid) | **Get** /characters/{character_id}/mail/{mail_id}/ | Return a mail
*MailApi* | [**GetCharactersCharacterIdMailUnread**](docs/MailApi.md#getcharacterscharacteridmailunread) | **Get** /characters/{character_id}/mail/unread/ | Return the number of unread mails
*MailApi* | [**PostCharactersCharacterIdMail**](docs/MailApi.md#postcharacterscharacteridmail) | **Post** /characters/{character_id}/mail/ | Send a new mail
*MailApi* | [**PutCharactersCharacterIdMailMailId**](docs/MailApi.md#putcharacterscharacteridmailmailid) | **Put** /characters/{character_id}/mail/{mail_id}/ | Update metadata about a mail
*MarketApi* | [**GetMarketsGroups**](docs/MarketApi.md#getmarketsgroups) | **Get** /markets/groups/ | Get item groups
*MarketApi* | [**GetMarketsGroupsMarketGroupId**](docs/MarketApi.md#getmarketsgroupsmarketgroupid) | **Get** /markets/groups/{market_group_id}/ | Get item group information
*MarketApi* | [**GetMarketsPrices**](docs/MarketApi.md#getmarketsprices) | **Get** /markets/prices/ | List market prices
*MarketApi* | [**GetMarketsRegionIdHistory**](docs/MarketApi.md#getmarketsregionidhistory) | **Get** /markets/{region_id}/history/ | List historical market statistics in a region
*MarketApi* | [**GetMarketsRegionIdOrders**](docs/MarketApi.md#getmarketsregionidorders) | **Get** /markets/{region_id}/orders/ | List orders in a region
*MarketApi* | [**GetMarketsStructuresStructureId**](docs/MarketApi.md#getmarketsstructuresstructureid) | **Get** /markets/structures/{structure_id}/ | List orders in a structure
*PlanetaryInteractionApi* | [**GetCharactersCharacterIdPlanets**](docs/PlanetaryInteractionApi.md#getcharacterscharacteridplanets) | **Get** /characters/{character_id}/planets/ | Get colonies
*PlanetaryInteractionApi* | [**GetCharactersCharacterIdPlanetsPlanetId**](docs/PlanetaryInteractionApi.md#getcharacterscharacteridplanetsplanetid) | **Get** /characters/{character_id}/planets/{planet_id}/ | Get colony layout
*PlanetaryInteractionApi* | [**GetUniverseSchematicsSchematicId**](docs/PlanetaryInteractionApi.md#getuniverseschematicsschematicid) | **Get** /universe/schematics/{schematic_id}/ | Get schematic information
*SearchApi* | [**GetCharactersCharacterIdSearch**](docs/SearchApi.md#getcharacterscharacteridsearch) | **Get** /characters/{character_id}/search/ | Search on a string
*SearchApi* | [**GetSearch**](docs/SearchApi.md#getsearch) | **Get** /search/ | Search on a string
*SovereigntyApi* | [**GetSovereigntyCampaigns**](docs/SovereigntyApi.md#getsovereigntycampaigns) | **Get** /sovereignty/campaigns/ | List sovereignty campaigns
*SovereigntyApi* | [**GetSovereigntyStructures**](docs/SovereigntyApi.md#getsovereigntystructures) | **Get** /sovereignty/structures/ | List sovereignty structures
*UniverseApi* | [**GetUniverseBloodlines**](docs/UniverseApi.md#getuniversebloodlines) | **Get** /universe/bloodlines/ | Get bloodlines
*UniverseApi* | [**GetUniverseCategories**](docs/UniverseApi.md#getuniversecategories) | **Get** /universe/categories/ | Get item categories
*UniverseApi* | [**GetUniverseCategoriesCategoryId**](docs/UniverseApi.md#getuniversecategoriescategoryid) | **Get** /universe/categories/{category_id}/ | Get item category information
*UniverseApi* | [**GetUniverseConstellations**](docs/UniverseApi.md#getuniverseconstellations) | **Get** /universe/constellations/ | Get constellations
*UniverseApi* | [**GetUniverseConstellationsConstellationId**](docs/UniverseApi.md#getuniverseconstellationsconstellationid) | **Get** /universe/constellations/{constellation_id}/ | Get constellation information
*UniverseApi* | [**GetUniverseFactions**](docs/UniverseApi.md#getuniversefactions) | **Get** /universe/factions/ | Get factions
*UniverseApi* | [**GetUniverseGraphics**](docs/UniverseApi.md#getuniversegraphics) | **Get** /universe/graphics/ | Get graphics
*UniverseApi* | [**GetUniverseGraphicsGraphicId**](docs/UniverseApi.md#getuniversegraphicsgraphicid) | **Get** /universe/graphics/{graphic_id}/ | Get graphic information
*UniverseApi* | [**GetUniverseGroups**](docs/UniverseApi.md#getuniversegroups) | **Get** /universe/groups/ | Get item groups
*UniverseApi* | [**GetUniverseGroupsGroupId**](docs/UniverseApi.md#getuniversegroupsgroupid) | **Get** /universe/groups/{group_id}/ | Get item group information
*UniverseApi* | [**GetUniverseMoonsMoonId**](docs/UniverseApi.md#getuniversemoonsmoonid) | **Get** /universe/moons/{moon_id}/ | Get moon information
*UniverseApi* | [**GetUniversePlanetsPlanetId**](docs/UniverseApi.md#getuniverseplanetsplanetid) | **Get** /universe/planets/{planet_id}/ | Get planet information
*UniverseApi* | [**GetUniverseRaces**](docs/UniverseApi.md#getuniverseraces) | **Get** /universe/races/ | Get character races
*UniverseApi* | [**GetUniverseRegions**](docs/UniverseApi.md#getuniverseregions) | **Get** /universe/regions/ | Get regions
*UniverseApi* | [**GetUniverseRegionsRegionId**](docs/UniverseApi.md#getuniverseregionsregionid) | **Get** /universe/regions/{region_id}/ | Get region information
*UniverseApi* | [**GetUniverseStargatesStargateId**](docs/UniverseApi.md#getuniversestargatesstargateid) | **Get** /universe/stargates/{stargate_id}/ | Get stargate information
*UniverseApi* | [**GetUniverseStationsStationId**](docs/UniverseApi.md#getuniversestationsstationid) | **Get** /universe/stations/{station_id}/ | Get station information
*UniverseApi* | [**GetUniverseStructures**](docs/UniverseApi.md#getuniversestructures) | **Get** /universe/structures/ | List all public structures
*UniverseApi* | [**GetUniverseStructuresStructureId**](docs/UniverseApi.md#getuniversestructuresstructureid) | **Get** /universe/structures/{structure_id}/ | Get structure information
*UniverseApi* | [**GetUniverseSystems**](docs/UniverseApi.md#getuniversesystems) | **Get** /universe/systems/ | Get solar systems
*UniverseApi* | [**GetUniverseSystemsSystemId**](docs/UniverseApi.md#getuniversesystemssystemid) | **Get** /universe/systems/{system_id}/ | Get solar system information
*UniverseApi* | [**GetUniverseTypes**](docs/UniverseApi.md#getuniversetypes) | **Get** /universe/types/ | Get types
*UniverseApi* | [**GetUniverseTypesTypeId**](docs/UniverseApi.md#getuniversetypestypeid) | **Get** /universe/types/{type_id}/ | Get type information
*UniverseApi* | [**PostUniverseNames**](docs/UniverseApi.md#postuniversenames) | **Post** /universe/names/ | Get names and categories for a set of ID&#39;s
*UserInterfaceApi* | [**PostUiAutopilotWaypoint**](docs/UserInterfaceApi.md#postuiautopilotwaypoint) | **Post** /ui/autopilot/waypoint/ | Set Autopilot Waypoint
*UserInterfaceApi* | [**PostUiOpenwindowContract**](docs/UserInterfaceApi.md#postuiopenwindowcontract) | **Post** /ui/openwindow/contract/ | Open Contract Window
*UserInterfaceApi* | [**PostUiOpenwindowInformation**](docs/UserInterfaceApi.md#postuiopenwindowinformation) | **Post** /ui/openwindow/information/ | Open Information Window
*UserInterfaceApi* | [**PostUiOpenwindowMarketdetails**](docs/UserInterfaceApi.md#postuiopenwindowmarketdetails) | **Post** /ui/openwindow/marketdetails/ | Open Market Details
*UserInterfaceApi* | [**PostUiOpenwindowNewmail**](docs/UserInterfaceApi.md#postuiopenwindownewmail) | **Post** /ui/openwindow/newmail/ | Open New Mail Window
*WalletApi* | [**GetCharactersCharacterIdWallets**](docs/WalletApi.md#getcharacterscharacteridwallets) | **Get** /characters/{character_id}/wallets/ | List wallets and balances
*WarsApi* | [**GetWars**](docs/WarsApi.md#getwars) | **Get** /wars/ | List wars
*WarsApi* | [**GetWarsWarId**](docs/WarsApi.md#getwarswarid) | **Get** /wars/{war_id}/ | Get war information
*WarsApi* | [**GetWarsWarIdKillmails**](docs/WarsApi.md#getwarswaridkillmails) | **Get** /wars/{war_id}/killmails/ | List kills for a war


## Documentation For Models

 - [DeleteCharactersCharacterIdContactsForbidden](docs/DeleteCharactersCharacterIdContactsForbidden.md)
 - [DeleteCharactersCharacterIdContactsInternalServerError](docs/DeleteCharactersCharacterIdContactsInternalServerError.md)
 - [DeleteCharactersCharacterIdFittingsFittingIdForbidden](docs/DeleteCharactersCharacterIdFittingsFittingIdForbidden.md)
 - [DeleteCharactersCharacterIdFittingsFittingIdInternalServerError](docs/DeleteCharactersCharacterIdFittingsFittingIdInternalServerError.md)
 - [DeleteCharactersCharacterIdMailLabelsLabelIdForbidden](docs/DeleteCharactersCharacterIdMailLabelsLabelIdForbidden.md)
 - [DeleteCharactersCharacterIdMailLabelsLabelIdInternalServerError](docs/DeleteCharactersCharacterIdMailLabelsLabelIdInternalServerError.md)
 - [DeleteCharactersCharacterIdMailLabelsLabelIdUnprocessableEntity](docs/DeleteCharactersCharacterIdMailLabelsLabelIdUnprocessableEntity.md)
 - [DeleteCharactersCharacterIdMailMailIdForbidden](docs/DeleteCharactersCharacterIdMailMailIdForbidden.md)
 - [DeleteCharactersCharacterIdMailMailIdInternalServerError](docs/DeleteCharactersCharacterIdMailMailIdInternalServerError.md)
 - [DeleteFleetsFleetIdMembersMemberIdForbidden](docs/DeleteFleetsFleetIdMembersMemberIdForbidden.md)
 - [DeleteFleetsFleetIdMembersMemberIdInternalServerError](docs/DeleteFleetsFleetIdMembersMemberIdInternalServerError.md)
 - [DeleteFleetsFleetIdMembersMemberIdNotFound](docs/DeleteFleetsFleetIdMembersMemberIdNotFound.md)
 - [DeleteFleetsFleetIdSquadsSquadIdForbidden](docs/DeleteFleetsFleetIdSquadsSquadIdForbidden.md)
 - [DeleteFleetsFleetIdSquadsSquadIdInternalServerError](docs/DeleteFleetsFleetIdSquadsSquadIdInternalServerError.md)
 - [DeleteFleetsFleetIdSquadsSquadIdNotFound](docs/DeleteFleetsFleetIdSquadsSquadIdNotFound.md)
 - [DeleteFleetsFleetIdWingsWingIdForbidden](docs/DeleteFleetsFleetIdWingsWingIdForbidden.md)
 - [DeleteFleetsFleetIdWingsWingIdInternalServerError](docs/DeleteFleetsFleetIdWingsWingIdInternalServerError.md)
 - [DeleteFleetsFleetIdWingsWingIdNotFound](docs/DeleteFleetsFleetIdWingsWingIdNotFound.md)
 - [GetAlliancesAllianceIdCorporationsInternalServerError](docs/GetAlliancesAllianceIdCorporationsInternalServerError.md)
 - [GetAlliancesAllianceIdIconsInternalServerError](docs/GetAlliancesAllianceIdIconsInternalServerError.md)
 - [GetAlliancesAllianceIdIconsNotFound](docs/GetAlliancesAllianceIdIconsNotFound.md)
 - [GetAlliancesAllianceIdIconsOk](docs/GetAlliancesAllianceIdIconsOk.md)
 - [GetAlliancesAllianceIdInternalServerError](docs/GetAlliancesAllianceIdInternalServerError.md)
 - [GetAlliancesAllianceIdNotFound](docs/GetAlliancesAllianceIdNotFound.md)
 - [GetAlliancesAllianceIdOk](docs/GetAlliancesAllianceIdOk.md)
 - [GetAlliancesInternalServerError](docs/GetAlliancesInternalServerError.md)
 - [GetAlliancesNames200Ok](docs/GetAlliancesNames200Ok.md)
 - [GetAlliancesNamesInternalServerError](docs/GetAlliancesNamesInternalServerError.md)
 - [GetCharactersCharacterIdAssets200Ok](docs/GetCharactersCharacterIdAssets200Ok.md)
 - [GetCharactersCharacterIdAssetsForbidden](docs/GetCharactersCharacterIdAssetsForbidden.md)
 - [GetCharactersCharacterIdAssetsInternalServerError](docs/GetCharactersCharacterIdAssetsInternalServerError.md)
 - [GetCharactersCharacterIdBookmarks200Ok](docs/GetCharactersCharacterIdBookmarks200Ok.md)
 - [GetCharactersCharacterIdBookmarksCoordinates](docs/GetCharactersCharacterIdBookmarksCoordinates.md)
 - [GetCharactersCharacterIdBookmarksFolders200Ok](docs/GetCharactersCharacterIdBookmarksFolders200Ok.md)
 - [GetCharactersCharacterIdBookmarksFoldersForbidden](docs/GetCharactersCharacterIdBookmarksFoldersForbidden.md)
 - [GetCharactersCharacterIdBookmarksFoldersInternalServerError](docs/GetCharactersCharacterIdBookmarksFoldersInternalServerError.md)
 - [GetCharactersCharacterIdBookmarksForbidden](docs/GetCharactersCharacterIdBookmarksForbidden.md)
 - [GetCharactersCharacterIdBookmarksInternalServerError](docs/GetCharactersCharacterIdBookmarksInternalServerError.md)
 - [GetCharactersCharacterIdBookmarksItem](docs/GetCharactersCharacterIdBookmarksItem.md)
 - [GetCharactersCharacterIdBookmarksTarget](docs/GetCharactersCharacterIdBookmarksTarget.md)
 - [GetCharactersCharacterIdCalendar200Ok](docs/GetCharactersCharacterIdCalendar200Ok.md)
 - [GetCharactersCharacterIdCalendarForbidden](docs/GetCharactersCharacterIdCalendarForbidden.md)
 - [GetCharactersCharacterIdCalendarInternalServerError](docs/GetCharactersCharacterIdCalendarInternalServerError.md)
 - [GetCharactersCharacterIdClones200Ok](docs/GetCharactersCharacterIdClones200Ok.md)
 - [GetCharactersCharacterIdClonesForbidden](docs/GetCharactersCharacterIdClonesForbidden.md)
 - [GetCharactersCharacterIdClonesInternalServerError](docs/GetCharactersCharacterIdClonesInternalServerError.md)
 - [GetCharactersCharacterIdContacts200Ok](docs/GetCharactersCharacterIdContacts200Ok.md)
 - [GetCharactersCharacterIdContactsForbidden](docs/GetCharactersCharacterIdContactsForbidden.md)
 - [GetCharactersCharacterIdContactsInternalServerError](docs/GetCharactersCharacterIdContactsInternalServerError.md)
 - [GetCharactersCharacterIdContactsLabels200Ok](docs/GetCharactersCharacterIdContactsLabels200Ok.md)
 - [GetCharactersCharacterIdContactsLabelsForbidden](docs/GetCharactersCharacterIdContactsLabelsForbidden.md)
 - [GetCharactersCharacterIdContactsLabelsInternalServerError](docs/GetCharactersCharacterIdContactsLabelsInternalServerError.md)
 - [GetCharactersCharacterIdCorporationhistory200Ok](docs/GetCharactersCharacterIdCorporationhistory200Ok.md)
 - [GetCharactersCharacterIdCorporationhistoryInternalServerError](docs/GetCharactersCharacterIdCorporationhistoryInternalServerError.md)
 - [GetCharactersCharacterIdFittings200Ok](docs/GetCharactersCharacterIdFittings200Ok.md)
 - [GetCharactersCharacterIdFittingsForbidden](docs/GetCharactersCharacterIdFittingsForbidden.md)
 - [GetCharactersCharacterIdFittingsInternalServerError](docs/GetCharactersCharacterIdFittingsInternalServerError.md)
 - [GetCharactersCharacterIdFittingsItem](docs/GetCharactersCharacterIdFittingsItem.md)
 - [GetCharactersCharacterIdKillmailsRecent200Ok](docs/GetCharactersCharacterIdKillmailsRecent200Ok.md)
 - [GetCharactersCharacterIdKillmailsRecentForbidden](docs/GetCharactersCharacterIdKillmailsRecentForbidden.md)
 - [GetCharactersCharacterIdKillmailsRecentInternalServerError](docs/GetCharactersCharacterIdKillmailsRecentInternalServerError.md)
 - [GetCharactersCharacterIdLocationForbidden](docs/GetCharactersCharacterIdLocationForbidden.md)
 - [GetCharactersCharacterIdLocationInternalServerError](docs/GetCharactersCharacterIdLocationInternalServerError.md)
 - [GetCharactersCharacterIdLocationOk](docs/GetCharactersCharacterIdLocationOk.md)
 - [GetCharactersCharacterIdLoyaltyPoints200Ok](docs/GetCharactersCharacterIdLoyaltyPoints200Ok.md)
 - [GetCharactersCharacterIdLoyaltyPointsForbidden](docs/GetCharactersCharacterIdLoyaltyPointsForbidden.md)
 - [GetCharactersCharacterIdLoyaltyPointsInternalServerError](docs/GetCharactersCharacterIdLoyaltyPointsInternalServerError.md)
 - [GetCharactersCharacterIdMail200Ok](docs/GetCharactersCharacterIdMail200Ok.md)
 - [GetCharactersCharacterIdMailForbidden](docs/GetCharactersCharacterIdMailForbidden.md)
 - [GetCharactersCharacterIdMailInternalServerError](docs/GetCharactersCharacterIdMailInternalServerError.md)
 - [GetCharactersCharacterIdMailLists200Ok](docs/GetCharactersCharacterIdMailLists200Ok.md)
 - [GetCharactersCharacterIdMailListsForbidden](docs/GetCharactersCharacterIdMailListsForbidden.md)
 - [GetCharactersCharacterIdMailListsInternalServerError](docs/GetCharactersCharacterIdMailListsInternalServerError.md)
 - [GetCharactersCharacterIdMailMailIdForbidden](docs/GetCharactersCharacterIdMailMailIdForbidden.md)
 - [GetCharactersCharacterIdMailMailIdInternalServerError](docs/GetCharactersCharacterIdMailMailIdInternalServerError.md)
 - [GetCharactersCharacterIdMailMailIdNotFound](docs/GetCharactersCharacterIdMailMailIdNotFound.md)
 - [GetCharactersCharacterIdMailMailIdOk](docs/GetCharactersCharacterIdMailMailIdOk.md)
 - [GetCharactersCharacterIdMailMailIdRecipient](docs/GetCharactersCharacterIdMailMailIdRecipient.md)
 - [GetCharactersCharacterIdMailRecipient](docs/GetCharactersCharacterIdMailRecipient.md)
 - [GetCharactersCharacterIdMailUnreadForbidden](docs/GetCharactersCharacterIdMailUnreadForbidden.md)
 - [GetCharactersCharacterIdMailUnreadInternalServerError](docs/GetCharactersCharacterIdMailUnreadInternalServerError.md)
 - [GetCharactersCharacterIdPlanets200Ok](docs/GetCharactersCharacterIdPlanets200Ok.md)
 - [GetCharactersCharacterIdPlanetsForbidden](docs/GetCharactersCharacterIdPlanetsForbidden.md)
 - [GetCharactersCharacterIdPlanetsInternalServerError](docs/GetCharactersCharacterIdPlanetsInternalServerError.md)
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
 - [GetCharactersCharacterIdShipForbidden](docs/GetCharactersCharacterIdShipForbidden.md)
 - [GetCharactersCharacterIdShipInternalServerError](docs/GetCharactersCharacterIdShipInternalServerError.md)
 - [GetCharactersCharacterIdShipOk](docs/GetCharactersCharacterIdShipOk.md)
 - [GetCharactersCharacterIdWallets200Ok](docs/GetCharactersCharacterIdWallets200Ok.md)
 - [GetCharactersCharacterIdWalletsForbidden](docs/GetCharactersCharacterIdWalletsForbidden.md)
 - [GetCharactersCharacterIdWalletsInternalServerError](docs/GetCharactersCharacterIdWalletsInternalServerError.md)
 - [GetCharactersNames200Ok](docs/GetCharactersNames200Ok.md)
 - [GetCharactersNamesInternalServerError](docs/GetCharactersNamesInternalServerError.md)
 - [GetCorporationsCorporationIdAlliancehistory200Ok](docs/GetCorporationsCorporationIdAlliancehistory200Ok.md)
 - [GetCorporationsCorporationIdAlliancehistoryAlliance](docs/GetCorporationsCorporationIdAlliancehistoryAlliance.md)
 - [GetCorporationsCorporationIdAlliancehistoryInternalServerError](docs/GetCorporationsCorporationIdAlliancehistoryInternalServerError.md)
 - [GetCorporationsCorporationIdIconsInternalServerError](docs/GetCorporationsCorporationIdIconsInternalServerError.md)
 - [GetCorporationsCorporationIdIconsNotFound](docs/GetCorporationsCorporationIdIconsNotFound.md)
 - [GetCorporationsCorporationIdIconsOk](docs/GetCorporationsCorporationIdIconsOk.md)
 - [GetCorporationsCorporationIdRoles200Ok](docs/GetCorporationsCorporationIdRoles200Ok.md)
 - [GetCorporationsCorporationIdRolesForbidden](docs/GetCorporationsCorporationIdRolesForbidden.md)
 - [GetCorporationsCorporationIdRolesInternalServerError](docs/GetCorporationsCorporationIdRolesInternalServerError.md)
 - [GetCorporationsCorporationIdStructures200Ok](docs/GetCorporationsCorporationIdStructures200Ok.md)
 - [GetCorporationsCorporationIdStructuresCurrentVul](docs/GetCorporationsCorporationIdStructuresCurrentVul.md)
 - [GetCorporationsCorporationIdStructuresForbidden](docs/GetCorporationsCorporationIdStructuresForbidden.md)
 - [GetCorporationsCorporationIdStructuresInternalServerError](docs/GetCorporationsCorporationIdStructuresInternalServerError.md)
 - [GetCorporationsCorporationIdStructuresNextVul](docs/GetCorporationsCorporationIdStructuresNextVul.md)
 - [GetCorporationsCorporationIdStructuresService](docs/GetCorporationsCorporationIdStructuresService.md)
 - [GetCorporationsNames200Ok](docs/GetCorporationsNames200Ok.md)
 - [GetCorporationsNamesInternalServerError](docs/GetCorporationsNamesInternalServerError.md)
 - [GetCorporationsNpccorpsInternalServerError](docs/GetCorporationsNpccorpsInternalServerError.md)
 - [GetDogmaAttributesAttributeIdInternalServerError](docs/GetDogmaAttributesAttributeIdInternalServerError.md)
 - [GetDogmaAttributesAttributeIdNotFound](docs/GetDogmaAttributesAttributeIdNotFound.md)
 - [GetDogmaAttributesAttributeIdOk](docs/GetDogmaAttributesAttributeIdOk.md)
 - [GetDogmaAttributesInternalServerError](docs/GetDogmaAttributesInternalServerError.md)
 - [GetDogmaEffectsEffectIdInternalServerError](docs/GetDogmaEffectsEffectIdInternalServerError.md)
 - [GetDogmaEffectsEffectIdModifier](docs/GetDogmaEffectsEffectIdModifier.md)
 - [GetDogmaEffectsEffectIdNotFound](docs/GetDogmaEffectsEffectIdNotFound.md)
 - [GetDogmaEffectsEffectIdOk](docs/GetDogmaEffectsEffectIdOk.md)
 - [GetDogmaEffectsInternalServerError](docs/GetDogmaEffectsInternalServerError.md)
 - [GetFleetsFleetIdForbidden](docs/GetFleetsFleetIdForbidden.md)
 - [GetFleetsFleetIdInternalServerError](docs/GetFleetsFleetIdInternalServerError.md)
 - [GetFleetsFleetIdMembers200Ok](docs/GetFleetsFleetIdMembers200Ok.md)
 - [GetFleetsFleetIdMembersForbidden](docs/GetFleetsFleetIdMembersForbidden.md)
 - [GetFleetsFleetIdMembersInternalServerError](docs/GetFleetsFleetIdMembersInternalServerError.md)
 - [GetFleetsFleetIdMembersNotFound](docs/GetFleetsFleetIdMembersNotFound.md)
 - [GetFleetsFleetIdNotFound](docs/GetFleetsFleetIdNotFound.md)
 - [GetFleetsFleetIdOk](docs/GetFleetsFleetIdOk.md)
 - [GetFleetsFleetIdWings200Ok](docs/GetFleetsFleetIdWings200Ok.md)
 - [GetFleetsFleetIdWingsForbidden](docs/GetFleetsFleetIdWingsForbidden.md)
 - [GetFleetsFleetIdWingsInternalServerError](docs/GetFleetsFleetIdWingsInternalServerError.md)
 - [GetFleetsFleetIdWingsNotFound](docs/GetFleetsFleetIdWingsNotFound.md)
 - [GetFleetsFleetIdWingsSquad](docs/GetFleetsFleetIdWingsSquad.md)
 - [GetIncursions200Ok](docs/GetIncursions200Ok.md)
 - [GetIncursionsInternalServerError](docs/GetIncursionsInternalServerError.md)
 - [GetIndustryFacilities200Ok](docs/GetIndustryFacilities200Ok.md)
 - [GetIndustryFacilitiesInternalServerError](docs/GetIndustryFacilitiesInternalServerError.md)
 - [GetIndustrySystems200Ok](docs/GetIndustrySystems200Ok.md)
 - [GetIndustrySystemsCostIndice](docs/GetIndustrySystemsCostIndice.md)
 - [GetIndustrySystemsInternalServerError](docs/GetIndustrySystemsInternalServerError.md)
 - [GetInsurancePrices200Ok](docs/GetInsurancePrices200Ok.md)
 - [GetInsurancePricesInternalServerError](docs/GetInsurancePricesInternalServerError.md)
 - [GetInsurancePricesLevel](docs/GetInsurancePricesLevel.md)
 - [GetKillmailsKillmailIdKillmailHashAttacker](docs/GetKillmailsKillmailIdKillmailHashAttacker.md)
 - [GetKillmailsKillmailIdKillmailHashInternalServerError](docs/GetKillmailsKillmailIdKillmailHashInternalServerError.md)
 - [GetKillmailsKillmailIdKillmailHashItem](docs/GetKillmailsKillmailIdKillmailHashItem.md)
 - [GetKillmailsKillmailIdKillmailHashItem1](docs/GetKillmailsKillmailIdKillmailHashItem1.md)
 - [GetKillmailsKillmailIdKillmailHashOk](docs/GetKillmailsKillmailIdKillmailHashOk.md)
 - [GetKillmailsKillmailIdKillmailHashPosition](docs/GetKillmailsKillmailIdKillmailHashPosition.md)
 - [GetKillmailsKillmailIdKillmailHashUnprocessableEntity](docs/GetKillmailsKillmailIdKillmailHashUnprocessableEntity.md)
 - [GetKillmailsKillmailIdKillmailHashVictim](docs/GetKillmailsKillmailIdKillmailHashVictim.md)
 - [GetLoyaltyStoresCorporationIdOffers200Ok](docs/GetLoyaltyStoresCorporationIdOffers200Ok.md)
 - [GetLoyaltyStoresCorporationIdOffersInternalServerError](docs/GetLoyaltyStoresCorporationIdOffersInternalServerError.md)
 - [GetLoyaltyStoresCorporationIdOffersRequiredItem](docs/GetLoyaltyStoresCorporationIdOffersRequiredItem.md)
 - [GetMarketsGroupsInternalServerError](docs/GetMarketsGroupsInternalServerError.md)
 - [GetMarketsGroupsMarketGroupIdInternalServerError](docs/GetMarketsGroupsMarketGroupIdInternalServerError.md)
 - [GetMarketsGroupsMarketGroupIdNotFound](docs/GetMarketsGroupsMarketGroupIdNotFound.md)
 - [GetMarketsGroupsMarketGroupIdOk](docs/GetMarketsGroupsMarketGroupIdOk.md)
 - [GetMarketsPrices200Ok](docs/GetMarketsPrices200Ok.md)
 - [GetMarketsPricesInternalServerError](docs/GetMarketsPricesInternalServerError.md)
 - [GetMarketsRegionIdHistory200Ok](docs/GetMarketsRegionIdHistory200Ok.md)
 - [GetMarketsRegionIdHistoryInternalServerError](docs/GetMarketsRegionIdHistoryInternalServerError.md)
 - [GetMarketsRegionIdHistoryUnprocessableEntity](docs/GetMarketsRegionIdHistoryUnprocessableEntity.md)
 - [GetMarketsRegionIdOrders200Ok](docs/GetMarketsRegionIdOrders200Ok.md)
 - [GetMarketsRegionIdOrdersInternalServerError](docs/GetMarketsRegionIdOrdersInternalServerError.md)
 - [GetMarketsRegionIdOrdersUnprocessableEntity](docs/GetMarketsRegionIdOrdersUnprocessableEntity.md)
 - [GetMarketsStructuresStructureId200Ok](docs/GetMarketsStructuresStructureId200Ok.md)
 - [GetMarketsStructuresStructureIdForbidden](docs/GetMarketsStructuresStructureIdForbidden.md)
 - [GetMarketsStructuresStructureIdInternalServerError](docs/GetMarketsStructuresStructureIdInternalServerError.md)
 - [GetSearchInternalServerError](docs/GetSearchInternalServerError.md)
 - [GetSearchOk](docs/GetSearchOk.md)
 - [GetSovereigntyCampaigns200Ok](docs/GetSovereigntyCampaigns200Ok.md)
 - [GetSovereigntyCampaignsInternalServerError](docs/GetSovereigntyCampaignsInternalServerError.md)
 - [GetSovereigntyCampaignsParticipant](docs/GetSovereigntyCampaignsParticipant.md)
 - [GetSovereigntyStructures200Ok](docs/GetSovereigntyStructures200Ok.md)
 - [GetSovereigntyStructuresInternalServerError](docs/GetSovereigntyStructuresInternalServerError.md)
 - [GetUniverseBloodlines200Ok](docs/GetUniverseBloodlines200Ok.md)
 - [GetUniverseBloodlinesInternalServerError](docs/GetUniverseBloodlinesInternalServerError.md)
 - [GetUniverseCategoriesCategoryIdInternalServerError](docs/GetUniverseCategoriesCategoryIdInternalServerError.md)
 - [GetUniverseCategoriesCategoryIdNotFound](docs/GetUniverseCategoriesCategoryIdNotFound.md)
 - [GetUniverseCategoriesCategoryIdOk](docs/GetUniverseCategoriesCategoryIdOk.md)
 - [GetUniverseCategoriesInternalServerError](docs/GetUniverseCategoriesInternalServerError.md)
 - [GetUniverseConstellationsConstellationIdInternalServerError](docs/GetUniverseConstellationsConstellationIdInternalServerError.md)
 - [GetUniverseConstellationsConstellationIdNotFound](docs/GetUniverseConstellationsConstellationIdNotFound.md)
 - [GetUniverseConstellationsConstellationIdOk](docs/GetUniverseConstellationsConstellationIdOk.md)
 - [GetUniverseConstellationsConstellationIdPosition](docs/GetUniverseConstellationsConstellationIdPosition.md)
 - [GetUniverseConstellationsInternalServerError](docs/GetUniverseConstellationsInternalServerError.md)
 - [GetUniverseFactions200Ok](docs/GetUniverseFactions200Ok.md)
 - [GetUniverseFactionsInternalServerError](docs/GetUniverseFactionsInternalServerError.md)
 - [GetUniverseGraphicsGraphicIdInternalServerError](docs/GetUniverseGraphicsGraphicIdInternalServerError.md)
 - [GetUniverseGraphicsGraphicIdNotFound](docs/GetUniverseGraphicsGraphicIdNotFound.md)
 - [GetUniverseGraphicsGraphicIdOk](docs/GetUniverseGraphicsGraphicIdOk.md)
 - [GetUniverseGraphicsInternalServerError](docs/GetUniverseGraphicsInternalServerError.md)
 - [GetUniverseGroupsGroupIdInternalServerError](docs/GetUniverseGroupsGroupIdInternalServerError.md)
 - [GetUniverseGroupsGroupIdNotFound](docs/GetUniverseGroupsGroupIdNotFound.md)
 - [GetUniverseGroupsGroupIdOk](docs/GetUniverseGroupsGroupIdOk.md)
 - [GetUniverseGroupsInternalServerError](docs/GetUniverseGroupsInternalServerError.md)
 - [GetUniverseMoonsMoonIdInternalServerError](docs/GetUniverseMoonsMoonIdInternalServerError.md)
 - [GetUniverseMoonsMoonIdNotFound](docs/GetUniverseMoonsMoonIdNotFound.md)
 - [GetUniverseMoonsMoonIdOk](docs/GetUniverseMoonsMoonIdOk.md)
 - [GetUniverseMoonsMoonIdPosition](docs/GetUniverseMoonsMoonIdPosition.md)
 - [GetUniversePlanetsPlanetIdInternalServerError](docs/GetUniversePlanetsPlanetIdInternalServerError.md)
 - [GetUniversePlanetsPlanetIdNotFound](docs/GetUniversePlanetsPlanetIdNotFound.md)
 - [GetUniversePlanetsPlanetIdOk](docs/GetUniversePlanetsPlanetIdOk.md)
 - [GetUniversePlanetsPlanetIdPosition](docs/GetUniversePlanetsPlanetIdPosition.md)
 - [GetUniverseRaces200Ok](docs/GetUniverseRaces200Ok.md)
 - [GetUniverseRacesInternalServerError](docs/GetUniverseRacesInternalServerError.md)
 - [GetUniverseRegionsInternalServerError](docs/GetUniverseRegionsInternalServerError.md)
 - [GetUniverseRegionsRegionIdInternalServerError](docs/GetUniverseRegionsRegionIdInternalServerError.md)
 - [GetUniverseRegionsRegionIdNotFound](docs/GetUniverseRegionsRegionIdNotFound.md)
 - [GetUniverseRegionsRegionIdOk](docs/GetUniverseRegionsRegionIdOk.md)
 - [GetUniverseSchematicsSchematicIdInternalServerError](docs/GetUniverseSchematicsSchematicIdInternalServerError.md)
 - [GetUniverseSchematicsSchematicIdNotFound](docs/GetUniverseSchematicsSchematicIdNotFound.md)
 - [GetUniverseSchematicsSchematicIdOk](docs/GetUniverseSchematicsSchematicIdOk.md)
 - [GetUniverseStargatesStargateIdDestination](docs/GetUniverseStargatesStargateIdDestination.md)
 - [GetUniverseStargatesStargateIdInternalServerError](docs/GetUniverseStargatesStargateIdInternalServerError.md)
 - [GetUniverseStargatesStargateIdNotFound](docs/GetUniverseStargatesStargateIdNotFound.md)
 - [GetUniverseStargatesStargateIdOk](docs/GetUniverseStargatesStargateIdOk.md)
 - [GetUniverseStargatesStargateIdPosition](docs/GetUniverseStargatesStargateIdPosition.md)
 - [GetUniverseStationsStationIdInternalServerError](docs/GetUniverseStationsStationIdInternalServerError.md)
 - [GetUniverseStationsStationIdOk](docs/GetUniverseStationsStationIdOk.md)
 - [GetUniverseStructuresInternalServerError](docs/GetUniverseStructuresInternalServerError.md)
 - [GetUniverseStructuresStructureIdForbidden](docs/GetUniverseStructuresStructureIdForbidden.md)
 - [GetUniverseStructuresStructureIdInternalServerError](docs/GetUniverseStructuresStructureIdInternalServerError.md)
 - [GetUniverseStructuresStructureIdNotFound](docs/GetUniverseStructuresStructureIdNotFound.md)
 - [GetUniverseStructuresStructureIdOk](docs/GetUniverseStructuresStructureIdOk.md)
 - [GetUniverseStructuresStructureIdPosition](docs/GetUniverseStructuresStructureIdPosition.md)
 - [GetUniverseSystemsInternalServerError](docs/GetUniverseSystemsInternalServerError.md)
 - [GetUniverseSystemsSystemIdInternalServerError](docs/GetUniverseSystemsSystemIdInternalServerError.md)
 - [GetUniverseSystemsSystemIdNotFound](docs/GetUniverseSystemsSystemIdNotFound.md)
 - [GetUniverseSystemsSystemIdOk](docs/GetUniverseSystemsSystemIdOk.md)
 - [GetUniverseTypesInternalServerError](docs/GetUniverseTypesInternalServerError.md)
 - [GetUniverseTypesTypeIdInternalServerError](docs/GetUniverseTypesTypeIdInternalServerError.md)
 - [GetUniverseTypesTypeIdNotFound](docs/GetUniverseTypesTypeIdNotFound.md)
 - [GetUniverseTypesTypeIdOk](docs/GetUniverseTypesTypeIdOk.md)
 - [GetWarsInternalServerError](docs/GetWarsInternalServerError.md)
 - [GetWarsWarIdAggressor](docs/GetWarsWarIdAggressor.md)
 - [GetWarsWarIdAlly](docs/GetWarsWarIdAlly.md)
 - [GetWarsWarIdDefender](docs/GetWarsWarIdDefender.md)
 - [GetWarsWarIdInternalServerError](docs/GetWarsWarIdInternalServerError.md)
 - [GetWarsWarIdKillmails200Ok](docs/GetWarsWarIdKillmails200Ok.md)
 - [GetWarsWarIdKillmailsInternalServerError](docs/GetWarsWarIdKillmailsInternalServerError.md)
 - [GetWarsWarIdKillmailsUnprocessableEntity](docs/GetWarsWarIdKillmailsUnprocessableEntity.md)
 - [GetWarsWarIdOk](docs/GetWarsWarIdOk.md)
 - [GetWarsWarIdUnprocessableEntity](docs/GetWarsWarIdUnprocessableEntity.md)
 - [PostCharactersCharacterIdContactsForbidden](docs/PostCharactersCharacterIdContactsForbidden.md)
 - [PostCharactersCharacterIdContactsInternalServerError](docs/PostCharactersCharacterIdContactsInternalServerError.md)
 - [PostCharactersCharacterIdFittingsCreated](docs/PostCharactersCharacterIdFittingsCreated.md)
 - [PostCharactersCharacterIdFittingsFitting](docs/PostCharactersCharacterIdFittingsFitting.md)
 - [PostCharactersCharacterIdFittingsForbidden](docs/PostCharactersCharacterIdFittingsForbidden.md)
 - [PostCharactersCharacterIdFittingsInternalServerError](docs/PostCharactersCharacterIdFittingsInternalServerError.md)
 - [PostCharactersCharacterIdFittingsItem](docs/PostCharactersCharacterIdFittingsItem.md)
 - [PostCharactersCharacterIdMailBadRequest](docs/PostCharactersCharacterIdMailBadRequest.md)
 - [PostCharactersCharacterIdMailForbidden](docs/PostCharactersCharacterIdMailForbidden.md)
 - [PostCharactersCharacterIdMailInternalServerError](docs/PostCharactersCharacterIdMailInternalServerError.md)
 - [PostCharactersCharacterIdMailMail](docs/PostCharactersCharacterIdMailMail.md)
 - [PostCharactersCharacterIdMailRecipient](docs/PostCharactersCharacterIdMailRecipient.md)
 - [PostFleetsFleetIdMembersForbidden](docs/PostFleetsFleetIdMembersForbidden.md)
 - [PostFleetsFleetIdMembersInternalServerError](docs/PostFleetsFleetIdMembersInternalServerError.md)
 - [PostFleetsFleetIdMembersInvitation](docs/PostFleetsFleetIdMembersInvitation.md)
 - [PostFleetsFleetIdMembersNotFound](docs/PostFleetsFleetIdMembersNotFound.md)
 - [PostFleetsFleetIdMembersUnprocessableEntity](docs/PostFleetsFleetIdMembersUnprocessableEntity.md)
 - [PostFleetsFleetIdWingsCreated](docs/PostFleetsFleetIdWingsCreated.md)
 - [PostFleetsFleetIdWingsForbidden](docs/PostFleetsFleetIdWingsForbidden.md)
 - [PostFleetsFleetIdWingsInternalServerError](docs/PostFleetsFleetIdWingsInternalServerError.md)
 - [PostFleetsFleetIdWingsNotFound](docs/PostFleetsFleetIdWingsNotFound.md)
 - [PostFleetsFleetIdWingsWingIdSquadsCreated](docs/PostFleetsFleetIdWingsWingIdSquadsCreated.md)
 - [PostFleetsFleetIdWingsWingIdSquadsForbidden](docs/PostFleetsFleetIdWingsWingIdSquadsForbidden.md)
 - [PostFleetsFleetIdWingsWingIdSquadsInternalServerError](docs/PostFleetsFleetIdWingsWingIdSquadsInternalServerError.md)
 - [PostFleetsFleetIdWingsWingIdSquadsNotFound](docs/PostFleetsFleetIdWingsWingIdSquadsNotFound.md)
 - [PostUiAutopilotWaypointForbidden](docs/PostUiAutopilotWaypointForbidden.md)
 - [PostUiAutopilotWaypointInternalServerError](docs/PostUiAutopilotWaypointInternalServerError.md)
 - [PostUiOpenwindowContractForbidden](docs/PostUiOpenwindowContractForbidden.md)
 - [PostUiOpenwindowContractInternalServerError](docs/PostUiOpenwindowContractInternalServerError.md)
 - [PostUiOpenwindowInformationForbidden](docs/PostUiOpenwindowInformationForbidden.md)
 - [PostUiOpenwindowInformationInternalServerError](docs/PostUiOpenwindowInformationInternalServerError.md)
 - [PostUiOpenwindowMarketdetailsForbidden](docs/PostUiOpenwindowMarketdetailsForbidden.md)
 - [PostUiOpenwindowMarketdetailsInternalServerError](docs/PostUiOpenwindowMarketdetailsInternalServerError.md)
 - [PostUiOpenwindowNewmailForbidden](docs/PostUiOpenwindowNewmailForbidden.md)
 - [PostUiOpenwindowNewmailInternalServerError](docs/PostUiOpenwindowNewmailInternalServerError.md)
 - [PostUiOpenwindowNewmailNewMail](docs/PostUiOpenwindowNewmailNewMail.md)
 - [PostUiOpenwindowNewmailUnprocessableEntity](docs/PostUiOpenwindowNewmailUnprocessableEntity.md)
 - [PostUniverseNames200Ok](docs/PostUniverseNames200Ok.md)
 - [PostUniverseNamesIds](docs/PostUniverseNamesIds.md)
 - [PostUniverseNamesInternalServerError](docs/PostUniverseNamesInternalServerError.md)
 - [PostUniverseNamesNotFound](docs/PostUniverseNamesNotFound.md)
 - [PutCharactersCharacterIdContactsForbidden](docs/PutCharactersCharacterIdContactsForbidden.md)
 - [PutCharactersCharacterIdContactsInternalServerError](docs/PutCharactersCharacterIdContactsInternalServerError.md)
 - [PutCharactersCharacterIdMailMailIdBadRequest](docs/PutCharactersCharacterIdMailMailIdBadRequest.md)
 - [PutCharactersCharacterIdMailMailIdContents](docs/PutCharactersCharacterIdMailMailIdContents.md)
 - [PutCharactersCharacterIdMailMailIdForbidden](docs/PutCharactersCharacterIdMailMailIdForbidden.md)
 - [PutCharactersCharacterIdMailMailIdInternalServerError](docs/PutCharactersCharacterIdMailMailIdInternalServerError.md)
 - [PutCorporationsCorporationIdStructuresStructureIdForbidden](docs/PutCorporationsCorporationIdStructuresStructureIdForbidden.md)
 - [PutCorporationsCorporationIdStructuresStructureIdInternalServerError](docs/PutCorporationsCorporationIdStructuresStructureIdInternalServerError.md)
 - [PutCorporationsCorporationIdStructuresStructureIdNewSchedule](docs/PutCorporationsCorporationIdStructuresStructureIdNewSchedule.md)
 - [PutFleetsFleetIdBadRequest](docs/PutFleetsFleetIdBadRequest.md)
 - [PutFleetsFleetIdForbidden](docs/PutFleetsFleetIdForbidden.md)
 - [PutFleetsFleetIdInternalServerError](docs/PutFleetsFleetIdInternalServerError.md)
 - [PutFleetsFleetIdMembersMemberIdForbidden](docs/PutFleetsFleetIdMembersMemberIdForbidden.md)
 - [PutFleetsFleetIdMembersMemberIdInternalServerError](docs/PutFleetsFleetIdMembersMemberIdInternalServerError.md)
 - [PutFleetsFleetIdMembersMemberIdMovement](docs/PutFleetsFleetIdMembersMemberIdMovement.md)
 - [PutFleetsFleetIdMembersMemberIdNotFound](docs/PutFleetsFleetIdMembersMemberIdNotFound.md)
 - [PutFleetsFleetIdMembersMemberIdUnprocessableEntity](docs/PutFleetsFleetIdMembersMemberIdUnprocessableEntity.md)
 - [PutFleetsFleetIdNewSettings](docs/PutFleetsFleetIdNewSettings.md)
 - [PutFleetsFleetIdNotFound](docs/PutFleetsFleetIdNotFound.md)
 - [PutFleetsFleetIdSquadsSquadIdForbidden](docs/PutFleetsFleetIdSquadsSquadIdForbidden.md)
 - [PutFleetsFleetIdSquadsSquadIdInternalServerError](docs/PutFleetsFleetIdSquadsSquadIdInternalServerError.md)
 - [PutFleetsFleetIdSquadsSquadIdNaming](docs/PutFleetsFleetIdSquadsSquadIdNaming.md)
 - [PutFleetsFleetIdSquadsSquadIdNotFound](docs/PutFleetsFleetIdSquadsSquadIdNotFound.md)
 - [PutFleetsFleetIdWingsWingIdForbidden](docs/PutFleetsFleetIdWingsWingIdForbidden.md)
 - [PutFleetsFleetIdWingsWingIdInternalServerError](docs/PutFleetsFleetIdWingsWingIdInternalServerError.md)
 - [PutFleetsFleetIdWingsWingIdNaming](docs/PutFleetsFleetIdWingsWingIdNaming.md)
 - [PutFleetsFleetIdWingsWingIdNotFound](docs/PutFleetsFleetIdWingsWingIdNotFound.md)


## Documentation For Authorization


## evesso

- **Type**: OAuth
- **Flow**: implicit
- **Authorization URL**: https://login.eveonline.com/oauth/authorize
- **Scopes**: 
 - **esi-assets.read_assets.v1**: EVE SSO scope esi-assets.read_assets.v1
 - **esi-bookmarks.read_character_bookmarks.v1**: EVE SSO scope esi-bookmarks.read_character_bookmarks.v1
 - **esi-calendar.read_calendar_events.v1**: EVE SSO scope esi-calendar.read_calendar_events.v1
 - **esi-characters.read_contacts.v1**: EVE SSO scope esi-characters.read_contacts.v1
 - **esi-characters.read_loyalty.v1**: EVE SSO scope esi-characters.read_loyalty.v1
 - **esi-characters.write_contacts.v1**: EVE SSO scope esi-characters.write_contacts.v1
 - **esi-clones.read_clones.v1**: EVE SSO scope esi-clones.read_clones.v1
 - **esi-corporations.read_corporation_membership.v1**: EVE SSO scope esi-corporations.read_corporation_membership.v1
 - **esi-corporations.read_structures.v1**: EVE SSO scope esi-corporations.read_structures.v1
 - **esi-corporations.write_structures.v1**: EVE SSO scope esi-corporations.write_structures.v1
 - **esi-fittings.read_fittings.v1**: EVE SSO scope esi-fittings.read_fittings.v1
 - **esi-fittings.write_fittings.v1**: EVE SSO scope esi-fittings.write_fittings.v1
 - **esi-fleets.read_fleet.v1**: EVE SSO scope esi-fleets.read_fleet.v1
 - **esi-fleets.write_fleet.v1**: EVE SSO scope esi-fleets.write_fleet.v1
 - **esi-killmails.read_killmails.v1**: EVE SSO scope esi-killmails.read_killmails.v1
 - **esi-location.read_location.v1**: EVE SSO scope esi-location.read_location.v1
 - **esi-location.read_ship_type.v1**: EVE SSO scope esi-location.read_ship_type.v1
 - **esi-mail.organize_mail.v1**: EVE SSO scope esi-mail.organize_mail.v1
 - **esi-mail.read_mail.v1**: EVE SSO scope esi-mail.read_mail.v1
 - **esi-mail.send_mail.v1**: EVE SSO scope esi-mail.send_mail.v1
 - **esi-markets.structure_markets.v1**: EVE SSO scope esi-markets.structure_markets.v1
 - **esi-planets.manage_planets.v1**: EVE SSO scope esi-planets.manage_planets.v1
 - **esi-search.search_structures.v1**: EVE SSO scope esi-search.search_structures.v1
 - **esi-ui.open_window.v1**: EVE SSO scope esi-ui.open_window.v1
 - **esi-ui.write_waypoint.v1**: EVE SSO scope esi-ui.write_waypoint.v1
 - **esi-universe.read_structures.v1**: EVE SSO scope esi-universe.read_structures.v1
 - **esi-wallet.read_character_wallet.v1**: EVE SSO scope esi-wallet.read_character_wallet.v1

Example:
```
  auth, err = oauth2conf.TokenSource(http.Client, token) // Access and Refresh token structure

  client, err := goesiv1.NewClient(nil, "goesiv1 client http://mysite.com")
  ctx := context.WithValue(context.TODO(), goesiv1.ContextOAuth2, auth)
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


