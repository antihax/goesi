# GetCharactersCharacterIdChatChannels200Ok

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | **int32** | Unique channel ID. Always negative for player-created channels. Permanent (CCP created) channels have a positive ID, but don&#39;t appear in the API | [default to null]
**Name** | **string** | Displayed name of channel | [default to null]
**OwnerId** | **int32** | owner_id integer | [default to null]
**ComparisonKey** | **string** | Normalized, unique string used to compare channel names | [default to null]
**HasPassword** | **bool** | If this is a password protected channel | [default to null]
**Motd** | **string** | Message of the day for this channel | [default to null]
**Allowed** | [**[]GetCharactersCharacterIdChatChannelsAllowed**](get_characters_character_id_chat_channels_allowed.md) | allowed array | [default to null]
**Operators** | [**[]GetCharactersCharacterIdChatChannelsOperator**](get_characters_character_id_chat_channels_operator.md) | operators array | [default to null]
**Blocked** | [**[]GetCharactersCharacterIdChatChannelsBlocked**](get_characters_character_id_chat_channels_blocked.md) | blocked array | [default to null]
**Muted** | [**[]GetCharactersCharacterIdChatChannelsMuted**](get_characters_character_id_chat_channels_muted.md) | muted array | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


