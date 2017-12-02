# GetCorporationsCorporationIdStructures200Ok

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StructureId** | **int64** | The Item ID of the structure | [default to null]
**TypeId** | **int32** | The type id of the structure | [default to null]
**CorporationId** | **int32** | ID of the corporation that owns the structure | [default to null]
**SystemId** | **int32** | The solar system the structure is in | [default to null]
**ProfileId** | **int32** | The id of the ACL profile for this citadel | [default to null]
**CurrentVul** | [**[]GetCorporationsCorporationIdStructuresCurrentVul**](get_corporations_corporation_id_structures_current_vul.md) | This week&#39;s vulnerability windows, Monday is day 0 | [default to null]
**NextVul** | [**[]GetCorporationsCorporationIdStructuresNextVul**](get_corporations_corporation_id_structures_next_vul.md) | Next week&#39;s vulnerability windows, Monday is day 0 | [default to null]
**FuelExpires** | [**time.Time**](time.Time.md) | Date on which the structure will run out of fuel | [optional] [default to null]
**Services** | [**[]GetCorporationsCorporationIdStructuresService**](get_corporations_corporation_id_structures_service.md) | Contains a list of service upgrades, and their state | [optional] [default to null]
**StateTimerStart** | [**time.Time**](time.Time.md) | Date at which the structure entered it&#39;s current state | [optional] [default to null]
**StateTimerEnd** | [**time.Time**](time.Time.md) | Date at which the structure will move to it&#39;s next state | [optional] [default to null]
**UnanchorsAt** | [**time.Time**](time.Time.md) | Date at which the structure will unanchor | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


