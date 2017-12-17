# GetWarsWarIdOk

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID of the specified war | [default to null]
**Declared** | [**time.Time**](time.Time.md) | Time that the war was declared | [default to null]
**Started** | [**time.Time**](time.Time.md) | Time when the war started and both sides could shoot each other | [optional] [default to null]
**Retracted** | [**time.Time**](time.Time.md) | Time the war was retracted but both sides could still shoot each other | [optional] [default to null]
**Finished** | [**time.Time**](time.Time.md) | Time the war ended and shooting was no longer allowed | [optional] [default to null]
**Mutual** | **bool** | Was the war declared mutual by both parties | [default to null]
**OpenForAllies** | **bool** | Is the war currently open for allies or not | [default to null]
**Aggressor** | [**GetWarsWarIdAggressor**](get_wars_war_id_aggressor.md) |  | [default to null]
**Defender** | [**GetWarsWarIdDefender**](get_wars_war_id_defender.md) |  | [default to null]
**Allies** | [**[]GetWarsWarIdAlly**](get_wars_war_id_ally.md) | allied corporations or alliances, each object contains either corporation_id or alliance_id | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


