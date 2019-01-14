# GetCorporationsCorporationIdStructures200Ok

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorporationId** | **int32** | ID of the corporation that owns the structure | [default to null]
**FuelExpires** | [**time.Time**](time.Time.md) | Date on which the structure will run out of fuel | [optional] [default to null]
**NextReinforceApply** | [**time.Time**](time.Time.md) | The date and time when the structure&#39;s newly requested reinforcement times (e.g. next_reinforce_hour and next_reinforce_day) will take effect | [optional] [default to null]
**NextReinforceHour** | **int32** | The requested change to reinforce_hour that will take effect at the time shown by next_reinforce_apply | [optional] [default to null]
**NextReinforceWeekday** | **int32** | The requested change to reinforce_weekday that will take effect at the time shown by next_reinforce_apply | [optional] [default to null]
**ProfileId** | **int32** | The id of the ACL profile for this citadel | [default to null]
**ReinforceHour** | **int32** | The hour of day that determines the four hour window when the structure will randomly exit its reinforcement periods and become vulnerable to attack against its armor and/or hull. The structure will become vulnerable at a random time that is +/- 2 hours centered on the value of this property | [default to null]
**ReinforceWeekday** | **int32** | The day of the week when the structure exits its final reinforcement period and becomes vulnerable to attack against its hull. Monday is 0 and Sunday is 6 | [optional] [default to null]
**Services** | [**[]GetCorporationsCorporationIdStructuresService**](get_corporations_corporation_id_structures_service.md) | Contains a list of service upgrades, and their state | [optional] [default to null]
**State** | **string** | state string | [default to null]
**StateTimerEnd** | [**time.Time**](time.Time.md) | Date at which the structure will move to it&#39;s next state | [optional] [default to null]
**StateTimerStart** | [**time.Time**](time.Time.md) | Date at which the structure entered it&#39;s current state | [optional] [default to null]
**StructureId** | **int64** | The Item ID of the structure | [default to null]
**SystemId** | **int32** | The solar system the structure is in | [default to null]
**TypeId** | **int32** | The type id of the structure | [default to null]
**UnanchorsAt** | [**time.Time**](time.Time.md) | Date at which the structure will unanchor | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


