# \WalletApi

All URIs are relative to *https://esi.tech.ccp.is/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdWallets**](WalletApi.md#GetCharactersCharacterIdWallets) | **Get** /characters/{character_id}/wallets/ | List wallets and balances
[**GetCharactersCharacterIdWalletsJournal**](WalletApi.md#GetCharactersCharacterIdWalletsJournal) | **Get** /characters/{character_id}/wallets/journal/ | Get character wallet journal


# **GetCharactersCharacterIdWallets**
> []GetCharactersCharacterIdWallets200Ok GetCharactersCharacterIdWallets(ctx, characterId, optional)
List wallets and balances

List your wallets and their balances. Characters typically have only one wallet, with wallet_id 1000 being the master wallet.  ---  Alternate route: `/legacy/characters/{character_id}/wallets/`  Alternate route: `/latest/characters/{character_id}/wallets/`  Alternate route: `/dev/characters/{character_id}/wallets/`   ---  This route is cached for up to 120 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **characterId** | **int32**| An EVE character ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **token** | **string**| Access token to use, if preferred over a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCharactersCharacterIdWallets200Ok**](get_characters_character_id_wallets_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdWalletsJournal**
> []GetCharactersCharacterIdWalletsJournal200Ok GetCharactersCharacterIdWalletsJournal(ctx, characterId, optional)
Get character wallet journal

Retrieve character wallet journal  ---  Alternate route: `/legacy/characters/{character_id}/wallets/journal/`  Alternate route: `/latest/characters/{character_id}/wallets/journal/`  Alternate route: `/dev/characters/{character_id}/wallets/journal/`   ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **characterId** | **int32**| An EVE character ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **fromId** | **int64**| Only show journal entries happened before the transaction referenced by this id | 
 **token** | **string**| Access token to use, if preferred over a header | 
 **userAgent** | **string**| Client identifier, takes precedence over headers | 
 **xUserAgent** | **string**| Client identifier, takes precedence over User-Agent | 

### Return type

[**[]GetCharactersCharacterIdWalletsJournal200Ok**](get_characters_character_id_wallets_journal_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

