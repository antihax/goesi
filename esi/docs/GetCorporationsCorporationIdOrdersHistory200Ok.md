# GetCorporationsCorporationIdOrdersHistory200Ok

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** | Number of days the order was valid for (starting from the issued date). An order expires at time issued + duration | [default to null]
**Escrow** | **float64** | For buy orders, the amount of ISK in escrow | [optional] [default to null]
**IsBuyOrder** | **bool** | True if the order is a bid (buy) order | [optional] [default to null]
**Issued** | [**time.Time**](time.Time.md) | Date and time when this order was issued | [default to null]
**IssuedBy** | **int32** | The character who issued this order | [optional] [default to null]
**LocationId** | **int64** | ID of the location where order was placed | [default to null]
**MinVolume** | **int32** | For buy orders, the minimum quantity that will be accepted in a matching sell order | [optional] [default to null]
**OrderId** | **int64** | Unique order ID | [default to null]
**Price** | **float64** | Cost per unit for this order | [default to null]
**Range_** | **string** | Valid order range, numbers are ranges in jumps | [default to null]
**RegionId** | **int32** | ID of the region where order was placed | [default to null]
**State** | **string** | Current order state | [default to null]
**TypeId** | **int32** | The type ID of the item transacted in this order | [default to null]
**VolumeRemain** | **int32** | Quantity of items still required or offered | [default to null]
**VolumeTotal** | **int32** | Quantity of items required or offered at time order was placed | [default to null]
**WalletDivision** | **int32** | The corporation wallet division used for this order | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


