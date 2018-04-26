# GetCorporationsCorporationIdStarbasesStarbaseIdOk

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowAllianceMembers** | **bool** | allow_alliance_members boolean | [default to null]
**AllowCorporationMembers** | **bool** | allow_corporation_members boolean | [default to null]
**Anchor** | **string** | Who can anchor starbase (POS) and its structures | [default to null]
**AttackIfAtWar** | **bool** | attack_if_at_war boolean | [default to null]
**AttackIfOtherSecurityStatusDropping** | **bool** | attack_if_other_security_status_dropping boolean | [default to null]
**AttackSecurityStatusThreshold** | **float32** | Starbase (POS) will attack if target&#39;s security standing is lower than this value | [optional] [default to null]
**AttackStandingThreshold** | **float32** | Starbase (POS) will attack if target&#39;s standing is lower than this value | [optional] [default to null]
**FuelBayTake** | **string** | Who can take fuel blocks out of the starbase (POS)&#39;s fuel bay | [default to null]
**FuelBayView** | **string** | Who can view the starbase (POS)&#39;s fule bay. Characters either need to have required role or belong to the starbase (POS) owner&#39;s corporation or alliance, as described by the enum, all other access settings follows the same scheme | [default to null]
**Fuels** | [**[]GetCorporationsCorporationIdStarbasesStarbaseIdFuel**](get_corporations_corporation_id_starbases_starbase_id_fuel.md) | Fuel blocks and other things that will be consumed when operating a starbase (POS) | [optional] [default to null]
**Offline** | **string** | Who can offline starbase (POS) and its structures | [default to null]
**Online** | **string** | Who can online starbase (POS) and its structures | [default to null]
**Unanchor** | **string** | Who can unanchor starbase (POS) and its structures | [default to null]
**UseAllianceStandings** | **bool** | True if the starbase (POS) is using alliance standings, otherwise using corporation&#39;s | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


