# GetCharactersCharacterIdFwStatsOk

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FactionId** | **int32** | The faction the given character is enlisted to fight for. Will not be included if character is not enlisted in faction warfare | [optional] [default to null]
**EnlistedOn** | [**time.Time**](time.Time.md) | The enlistment date of the given character into faction warfare. Will not be included if character is not enlisted in faction warfare | [optional] [default to null]
**CurrentRank** | **int32** | The given character&#39;s current faction rank | [optional] [default to null]
**HighestRank** | **int32** | The given character&#39;s highest faction rank achieved | [optional] [default to null]
**Kills** | [**GetCharactersCharacterIdFwStatsKills**](get_characters_character_id_fw_stats_kills.md) |  | [default to null]
**VictoryPoints** | [**GetCharactersCharacterIdFwStatsVictoryPoints**](get_characters_character_id_fw_stats_victory_points.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


