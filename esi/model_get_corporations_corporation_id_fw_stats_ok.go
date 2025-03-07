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

import (
	"time"
)

/* A list of GetCorporationsCorporationIdFwStatsOk. */
//easyjson:json
type GetCorporationsCorporationIdFwStatsOkList []GetCorporationsCorporationIdFwStatsOk

/* 200 ok object */
//easyjson:json
type GetCorporationsCorporationIdFwStatsOk struct {
	EnlistedOn    time.Time                                        `json:"enlisted_on,omitempty"` /* The enlistment date of the given corporation into faction warfare. Will not be included if corporation is not enlisted in faction warfare */
	FactionId     int32                                            `json:"faction_id,omitempty"`  /* The faction the given corporation is enlisted to fight for. Will not be included if corporation is not enlisted in faction warfare */
	Kills         GetCorporationsCorporationIdFwStatsKills         `json:"kills,omitempty"`
	Pilots        int32                                            `json:"pilots,omitempty"` /* How many pilots the enlisted corporation has. Will not be included if corporation is not enlisted in faction warfare */
	VictoryPoints GetCorporationsCorporationIdFwStatsVictoryPoints `json:"victory_points,omitempty"`
}
