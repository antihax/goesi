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

/* A list of PutCharactersCharacterIdMailMailIdContents. */
//easyjson:json
type PutCharactersCharacterIdMailMailIdContentsList []PutCharactersCharacterIdMailMailIdContents

/* contents object */
//easyjson:json
type PutCharactersCharacterIdMailMailIdContents struct {
	Labels []int32 `json:"labels,omitempty"` /* Labels to assign to the mail. Pre-existing labels are unassigned. */
	Read   bool    `json:"read,omitempty"`   /* Whether the mail is flagged as read */
}
