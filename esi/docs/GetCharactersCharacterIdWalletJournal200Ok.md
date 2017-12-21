# GetCharactersCharacterIdWalletJournal200Ok

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | [**time.Time**](time.Time.md) | Date and time of transaction | [default to null]
**RefId** | **int64** | Unique journal reference ID | [default to null]
**RefType** | **string** | Transaction type, different type of transaction will populate different fields in &#x60;extra_info&#x60; Note: If you have an existing XML API application that is using ref_types, you will need to know which string ESI ref_type maps to which integer. You can use the following gist to see string-&gt;int mappings: https://gist.github.com/ccp-zoetrope/c03db66d90c2148724c06171bc52e0ec | [default to null]
**FirstPartyId** | **int32** | first_party_id integer | [optional] [default to null]
**FirstPartyType** | **string** | first_party_type string | [optional] [default to null]
**SecondPartyId** | **int32** | second_party_id integer | [optional] [default to null]
**SecondPartyType** | **string** | second_party_type string | [optional] [default to null]
**Amount** | **float64** | Transaction amount. Positive when value transferred to the first party. Negative otherwise | [optional] [default to null]
**Balance** | **float64** | Wallet balance after transaction occurred | [optional] [default to null]
**Reason** | **string** | reason string | [optional] [default to null]
**TaxReceiverId** | **int32** | the corporation ID receiving any tax paid | [optional] [default to null]
**Tax** | **float64** | Tax amount received for tax related transactions | [optional] [default to null]
**ExtraInfo** | [**GetCharactersCharacterIdWalletJournalExtraInfo**](get_characters_character_id_wallet_journal_extra_info.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


