# GetCorporationsCorporationIdCustomsOffices200Ok

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllianceTaxRate** | **float32** | Only present if alliance access is allowed | [optional] [default to null]
**AllowAccessWithStandings** | **bool** | standing_level and any standing related tax rate only present when this is true | [default to null]
**AllowAllianceAccess** | **bool** | allow_alliance_access boolean | [default to null]
**BadStandingTaxRate** | **float32** | bad_standing_tax_rate number | [optional] [default to null]
**CorporationTaxRate** | **float32** | corporation_tax_rate number | [optional] [default to null]
**ExcellentStandingTaxRate** | **float32** | Tax rate for entities with excellent level of standing, only present if this level is allowed, same for all other standing related tax rates | [optional] [default to null]
**GoodStandingTaxRate** | **float32** | good_standing_tax_rate number | [optional] [default to null]
**NeutralStandingTaxRate** | **float32** | neutral_standing_tax_rate number | [optional] [default to null]
**OfficeId** | **int64** | unique ID of this customs office | [default to null]
**ReinforceExitEnd** | **int32** | reinforce_exit_end integer | [default to null]
**ReinforceExitStart** | **int32** | Together with reinforce_exit_end, marks a 2-hour period where this customs office could exit reinforcement mode during the day after initial attack | [default to null]
**StandingLevel** | **string** | Access is allowed only for entities with this level of standing or better | [optional] [default to null]
**SystemId** | **int32** | ID of the solar system this customs office is located in | [default to null]
**TerribleStandingTaxRate** | **float32** | terrible_standing_tax_rate number | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


