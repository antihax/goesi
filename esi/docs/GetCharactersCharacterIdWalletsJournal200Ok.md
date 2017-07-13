# GetCharactersCharacterIdWalletsJournal200Ok

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Transaction amount. Positive when value transferred to the first party. Negative otherwise | [optional] [default to null]
**ArgumentName** | **string** | argument_name string | [optional] [default to null]
**ArgumentValue** | **int32** | argument_value integer | [optional] [default to null]
**Balance** | **float32** | Wallet balance after transaction occurred | [optional] [default to null]
**Date** | [**time.Time**](time.Time.md) | Date and time of transaction | [default to null]
**FirstPartyId** | **int32** | first_party_id integer | [optional] [default to null]
**FirstPartyType** | **string** | first_party_type string | [optional] [default to null]
**Reason** | **string** | reason string | [optional] [default to null]
**RefId** | **int64** | Unique journal reference ID | [default to null]
**RefTypeId** | **int32** | Transaction type | [default to null]
**SecondPartyId** | **int32** | second_party_id integer | [optional] [default to null]
**SecondPartyType** | **string** | second_party_type string | [optional] [default to null]
**TaxAmount** | **float32** | Tax amount received for tax related transactions | [optional] [default to null]
**TaxRecieverId** | **int32** | For tax related transactions, gives the corporation ID of the entity receiving the tax | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


