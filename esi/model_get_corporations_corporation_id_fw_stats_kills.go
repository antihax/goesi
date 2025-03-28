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

/* A list of GetCorporationsCorporationIdFwStatsKills. */
//easyjson:json
type GetCorporationsCorporationIdFwStatsKillsList []GetCorporationsCorporationIdFwStatsKills

/* Summary of kills done by the given corporation against enemy factions */
//easyjson:json
type GetCorporationsCorporationIdFwStatsKills struct {
	LastWeek  int32 `json:"last_week,omitempty"` /* Last week's total number of kills by members of the given corporation against enemy factions */
	Total     int32 `json:"total,omitempty"`     /* Total number of kills by members of the given corporation against enemy factions since the corporation enlisted */
	Yesterday int32 `json:"yesterday,omitempty"` /* Yesterday's total number of kills by members of the given corporation against enemy factions */
}
