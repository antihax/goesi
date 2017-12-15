# GetCharactersCharacterIdOrders200Ok

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **int64** | Unique order ID | [default to null]
**TypeId** | **int32** | The type ID of the item transacted in this order | [default to null]
**RegionId** | **int32** | ID of the region where order was placed | [default to null]
**LocationId** | **int64** | ID of the location where order was placed | [default to null]
**Range_** | **string** | Valid order range, numbers are ranges in jumps | [default to null]
**IsBuyOrder** | **bool** | True for a bid (buy) order. False for an offer (sell) order | [default to null]
**Price** | **float64** | Cost per unit for this order | [default to null]
**VolumeTotal** | **int32** | Quantity of items required or offered at time order was placed | [default to null]
**VolumeRemain** | **int32** | Quantity of items still required or offered | [default to null]
**Issued** | [**time.Time**](time.Time.md) | Date and time when this order was issued | [default to null]
**State** | **string** | Current order state | [default to null]
**MinVolume** | **int32** | For bids (buy orders), the minimum quantity that will be accepted in a matching offer (sell order) | [default to null]
**AccountId** | **int32** | Wallet division for the buyer or seller of this order. Always 1000 for characters. Currently 1000 through 1006 for corporations | [default to null]
**Duration** | **int32** | Numer of days for which order is valid (starting from the issued date). An order expires at time issued + duration | [default to null]
**IsCorp** | **bool** | is_corp boolean | [default to null]
**Escrow** | **float64** | For buy orders, the amount of ISK in escrow | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


