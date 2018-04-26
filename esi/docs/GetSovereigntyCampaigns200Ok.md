# GetSovereigntyCampaigns200Ok

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttackersScore** | **float32** | Score for all attacking parties, only present in Defense Events.  | [optional] [default to null]
**CampaignId** | **int32** | Unique ID for this campaign. | [default to null]
**ConstellationId** | **int32** | The constellation in which the campaign will take place.  | [default to null]
**DefenderId** | **int32** | Defending alliance, only present in Defense Events  | [optional] [default to null]
**DefenderScore** | **float32** | Score for the defending alliance, only present in Defense Events.  | [optional] [default to null]
**EventType** | **string** | Type of event this campaign is for. tcu_defense, ihub_defense and station_defense are referred to as \&quot;Defense Events\&quot;, station_freeport as \&quot;Freeport Events\&quot;.  | [default to null]
**Participants** | [**[]GetSovereigntyCampaignsParticipant**](get_sovereignty_campaigns_participant.md) | Alliance participating and their respective scores, only present in Freeport Events.  | [optional] [default to null]
**SolarSystemId** | **int32** | The solar system the structure is located in.  | [default to null]
**StartTime** | [**time.Time**](time.Time.md) | Time the event is scheduled to start.  | [default to null]
**StructureId** | **int64** | The structure item ID that is related to this campaign.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


