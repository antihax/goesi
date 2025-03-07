/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 1.30
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package esi

/* A list of GetCharactersCharacterIdContractsContractIdItems200Ok. */
//easyjson:json
type GetCharactersCharacterIdContractsContractIdItems200OkList []GetCharactersCharacterIdContractsContractIdItems200Ok

/* 200 ok object */
//easyjson:json
type GetCharactersCharacterIdContractsContractIdItems200Ok struct {
	IsIncluded  bool  `json:"is_included,omitempty"`  /* true if the contract issuer has submitted this item with the contract, false if the isser is asking for this item in the contract */
	IsSingleton bool  `json:"is_singleton,omitempty"` /* is_singleton boolean */
	Quantity    int32 `json:"quantity,omitempty"`     /* Number of items in the stack */
	RawQuantity int32 `json:"raw_quantity,omitempty"` /* -1 indicates that the item is a singleton (non-stackable). If the item happens to be a Blueprint, -1 is an Original and -2 is a Blueprint Copy */
	RecordId    int64 `json:"record_id,omitempty"`    /* Unique ID for the item */
	TypeId      int32 `json:"type_id,omitempty"`      /* Type ID for item */
}
