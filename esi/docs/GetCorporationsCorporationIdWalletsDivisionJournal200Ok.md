# GetCorporationsCorporationIdWalletsDivisionJournal200Ok

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float64** | The amount of ISK given or taken from the wallet as a result of the given transaction. Positive when ISK is deposited into the wallet and negative when ISK is withdrawn | [optional] [default to null]
**Balance** | **float64** | Wallet balance after transaction occurred | [optional] [default to null]
**ContextId** | **int64** | An ID that gives extra context to the particular transaction. Because of legacy reasons the context is completely different per ref_type and means different things. It is also possible to not have a context_id | [optional] [default to null]
**ContextIdType** | **string** | The type of the given context_id if present | [optional] [default to null]
**Date** | [**time.Time**](time.Time.md) | Date and time of transaction | [default to null]
**Description** | **string** | The reason for the transaction, mirrors what is seen in the client | [default to null]
**FirstPartyId** | **int32** | The id of the first party involved in the transaction. This attribute has no consistency and is different or non existant for particular ref_types. The description attribute will help make sense of what this attribute means. For more info about the given ID it can be dropped into the /universe/names/ ESI route to determine its type and name | [optional] [default to null]
**Id** | **int64** | Unique journal reference ID | [default to null]
**Reason** | **string** | The user stated reason for the transaction. Only applies to some ref_types | [optional] [default to null]
**RefType** | **string** | \&quot;The transaction type for the given. transaction. Different transaction types will populate different attributes. Note: If you have an existing XML API application that is using ref_types, you will need to know which string ESI ref_type maps to which integer. You can look at the following file to see string-&gt;int mappings: https://github.com/ccpgames/eve-glue/blob/master/eve_glue/wallet_journal_ref.py\&quot; | [default to null]
**SecondPartyId** | **int32** | The id of the second party involved in the transaction. This attribute has no consistency and is different or non existant for particular ref_types. The description attribute will help make sense of what this attribute means. For more info about the given ID it can be dropped into the /universe/names/ ESI route to determine its type and name | [optional] [default to null]
**Tax** | **float64** | Tax amount received. Only applies to tax related transactions | [optional] [default to null]
**TaxReceiverId** | **int32** | The corporation ID receiving any tax paid. Only applies to tax related transactions | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


