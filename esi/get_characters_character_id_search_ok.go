/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.6.2
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

/* A list of GetCharactersCharacterIdSearchOk. */
//easyjson:json
type GetCharactersCharacterIdSearchOkList []GetCharactersCharacterIdSearchOk

/* 200 ok object */
//easyjson:json
type GetCharactersCharacterIdSearchOk struct {
	Agent         []int32 `json:"agent,omitempty"`         /* agent array */
	Alliance      []int32 `json:"alliance,omitempty"`      /* alliance array */
	Character     []int32 `json:"character,omitempty"`     /* character array */
	Constellation []int32 `json:"constellation,omitempty"` /* constellation array */
	Corporation   []int32 `json:"corporation,omitempty"`   /* corporation array */
	Faction       []int32 `json:"faction,omitempty"`       /* faction array */
	Inventorytype []int32 `json:"inventorytype,omitempty"` /* inventorytype array */
	Region        []int32 `json:"region,omitempty"`        /* region array */
	Solarsystem   []int32 `json:"solarsystem,omitempty"`   /* solarsystem array */
	Station       []int32 `json:"station,omitempty"`       /* station array */
	Structure     []int64 `json:"structure,omitempty"`     /* structure array */
	Wormhole      []int32 `json:"wormhole,omitempty"`      /* wormhole array */
}
