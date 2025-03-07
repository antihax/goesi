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

/* A list of PostFleetsFleetIdMembersInvitation. */
//easyjson:json
type PostFleetsFleetIdMembersInvitationList []PostFleetsFleetIdMembersInvitation

/* invitation object */
//easyjson:json
type PostFleetsFleetIdMembersInvitation struct {
	CharacterId int32  `json:"character_id,omitempty"` /* The character you want to invite */
	Role        string `json:"role,omitempty"`         /* If a character is invited with the `fleet_commander` role, neither `wing_id` or `squad_id` should be specified. If a character is invited with the `wing_commander` role, only `wing_id` should be specified. If a character is invited with the `squad_commander` role, both `wing_id` and `squad_id` should be specified. If a character is invited with the `squad_member` role, `wing_id` and `squad_id` should either both be specified or not specified at all. If they aren’t specified, the invited character will join any squad with available positions. */
	SquadId     int64  `json:"squad_id,omitempty"`     /* squad_id integer */
	WingId      int64  `json:"wing_id,omitempty"`      /* wing_id integer */
}
