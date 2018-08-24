# GetContractsPublicItemsContractId200Ok

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsBlueprintCopy** | **bool** | is_blueprint_copy boolean | [optional] [default to null]
**IsIncluded** | **bool** | true if the contract issuer has submitted this item with the contract, false if the isser is asking for this item in the contract | [default to null]
**ItemId** | **int64** | Unique ID for the item being sold. Not present if item is being requested by contract rather than sold with contract | [optional] [default to null]
**MaterialEfficiency** | **int32** | Material Efficiency Level of the blueprint | [optional] [default to null]
**Quantity** | **int32** | Number of items in the stack | [default to null]
**RecordId** | **int64** | Unique ID for the item, used by the contract system | [default to null]
**Runs** | **int32** | Number of runs remaining if the blueprint is a copy, -1 if it is an original | [optional] [default to null]
**TimeEfficiency** | **int32** | Time Efficiency Level of the blueprint | [optional] [default to null]
**TypeId** | **int32** | Type ID for item | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


