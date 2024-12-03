/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 1.28
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

/* A list of GetWarsWarIdDefender. */
//easyjson:json
type GetWarsWarIdDefenderList []GetWarsWarIdDefender

/* The defending corporation or alliance that declared this war, only contains either corporation_id or alliance_id */
//easyjson:json
type GetWarsWarIdDefender struct {
	AllianceId    int32   `json:"alliance_id,omitempty"`    /* Alliance ID if and only if the defender is an alliance */
	CorporationId int32   `json:"corporation_id,omitempty"` /* Corporation ID if and only if the defender is a corporation */
	IskDestroyed  float32 `json:"isk_destroyed,omitempty"`  /* ISK value of ships the defender has killed */
	ShipsKilled   int32   `json:"ships_killed,omitempty"`   /* The number of ships the defender has killed */
}
