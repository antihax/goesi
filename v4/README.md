# Go API client for goesiv4

## Documentation for API Endpoints

All URIs are relative to *https://esi.tech.ccp.is/v4*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CharacterApi* | [**GetCharactersCharacterId**](docs/CharacterApi.md#getcharacterscharacterid) | **Get** /characters/{character_id}/ | Get character&#39;s public information


## Documentation For Models

 - [GetCharactersCharacterIdInternalServerError](docs/GetCharactersCharacterIdInternalServerError.md)
 - [GetCharactersCharacterIdNotFound](docs/GetCharactersCharacterIdNotFound.md)
 - [GetCharactersCharacterIdOk](docs/GetCharactersCharacterIdOk.md)


## Documentation For Authorization

 All endpoints do not require authorization.


## Credits
https://github.com/go-resty/resty (MIT license) Copyright Â© 2015-2016 Jeevanandam M (jeeva@myjeeva.com)
 - Uses modified setBody and detectContentType

https://github.com/gregjones/httpcache (MIT license) Copyright Â© 2012 Greg Jones (greg.jones@gmail.com)
  - Uses parseCacheControl and CacheExpires as a helper function


