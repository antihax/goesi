# GetCharactersCharacterIdIndustryJobs200Ok

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityId** | **int32** | Job activity ID | [default to null]
**BlueprintId** | **int64** | blueprint_id integer | [default to null]
**BlueprintLocationId** | **int64** | Location ID of the location from which the blueprint was installed. Normally a station ID, but can also be an asset (e.g. container) or corporation facility | [default to null]
**BlueprintTypeId** | **int32** | blueprint_type_id integer | [default to null]
**CompletedCharacterId** | **int32** | ID of the character which completed this job | [optional] [default to null]
**CompletedDate** | [**time.Time**](time.Time.md) | Date and time when this job was completed | [optional] [default to null]
**Cost** | **float64** | The sume of job installation fee and industry facility tax | [optional] [default to null]
**Duration** | **int32** | Job duration in seconds | [default to null]
**EndDate** | [**time.Time**](time.Time.md) | Date and time when this job finished | [default to null]
**FacilityId** | **int64** | ID of the facility where this job is running | [default to null]
**InstallerId** | **int32** | ID of the character which installed this job | [default to null]
**JobId** | **int32** | Unique job ID | [default to null]
**LicensedRuns** | **int32** | Number of runs blueprint is licensed for | [optional] [default to null]
**OutputLocationId** | **int64** | Location ID of the location to which the output of the job will be delivered. Normally a station ID, but can also be a corporation facility | [default to null]
**PauseDate** | [**time.Time**](time.Time.md) | Date and time when this job was paused (i.e. time when the facility where this job was installed went offline) | [optional] [default to null]
**Probability** | **float32** | Chance of success for invention | [optional] [default to null]
**ProductTypeId** | **int32** | Type ID of product (manufactured, copied or invented) | [optional] [default to null]
**Runs** | **int32** | Number of runs for a manufacturing job, or number of copies to make for a blueprint copy | [default to null]
**StartDate** | [**time.Time**](time.Time.md) | Date and time when this job started | [default to null]
**StationId** | **int64** | ID of the station where industry facility is located | [default to null]
**Status** | **string** | status string | [default to null]
**SuccessfulRuns** | **int32** | Number of successful runs for this job. Equal to runs unless this is an invention job | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


