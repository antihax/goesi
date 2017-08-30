# GetCorporationsCorporationIdWalletsDivisionJournal200Ok

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Transaction amount. Positive when value transferred to the first party. Negative otherwise | [optional] [default to null]
**Balance** | **float32** | Wallet balance after transaction occurred | [optional] [default to null]
**Date** | [**time.Time**](time.Time.md) | Date and time of transaction | [default to null]
**ExtraInfo** | [**GetCorporationsCorporationIdWalletsDivisionJournalExtraInfo**](get_corporations_corporation_id_wallets_division_journal_extra_info.md) |  | [optional] [default to null]
**FirstPartyId** | **int32** | first_party_id integer | [optional] [default to null]
**FirstPartyType** | **string** | first_party_type string | [optional] [default to null]
**Reason** | **string** | reason string | [optional] [default to null]
**RefId** | **int64** | Unique journal reference ID | [default to null]
**RefType** | **string** | Transaction type, different type of transaction will populate different fields in &#x60;extra_info&#x60; | [default to null]
**SecondPartyId** | **int32** | second_party_id integer | [optional] [default to null]
**SecondPartyType** | **string** | second_party_type string | [optional] [default to null]
**Tax** | **float32** | Tax amount received for tax related transactions | [optional] [default to null]
**TaxRecieverId** | **int32** | the corporation ID receiving any tax paid | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


