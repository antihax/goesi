# GetCorporationsCorporationIdStarbases200Ok

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoonId** | **int32** | The moon this starbase (POS) is anchored on, unanchored POSes do not have this information | [optional] [default to null]
**OnlinedSince** | [**time.Time**](time.Time.md) | When the POS onlined, for starbases (POSes) in online state | [optional] [default to null]
**ReinforcedUntil** | [**time.Time**](time.Time.md) | When the POS will be out of reinforcement, for starbases (POSes) in reinforced state | [optional] [default to null]
**StarbaseId** | **int64** | Unique ID for this starbase (POS) | [default to null]
**State** | **string** | state string | [optional] [default to null]
**SystemId** | **int32** | The solar system this starbase (POS) is in, unanchored POSes have this information | [default to null]
**TypeId** | **int32** | Starbase (POS) type | [default to null]
**UnanchorAt** | [**time.Time**](time.Time.md) | When the POS started unanchoring, for starbases (POSes) in unanchoring state | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


