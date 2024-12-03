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

/* A list of GetUniverseFactions200Ok. */
//easyjson:json
type GetUniverseFactions200OkList []GetUniverseFactions200Ok

/* 200 ok object */
//easyjson:json
type GetUniverseFactions200Ok struct {
	CorporationId        int32   `json:"corporation_id,omitempty"`         /* corporation_id integer */
	Description          string  `json:"description,omitempty"`            /* description string */
	FactionId            int32   `json:"faction_id,omitempty"`             /* faction_id integer */
	IsUnique             bool    `json:"is_unique,omitempty"`              /* is_unique boolean */
	MilitiaCorporationId int32   `json:"militia_corporation_id,omitempty"` /* militia_corporation_id integer */
	Name                 string  `json:"name,omitempty"`                   /* name string */
	SizeFactor           float32 `json:"size_factor,omitempty"`            /* size_factor number */
	SolarSystemId        int32   `json:"solar_system_id,omitempty"`        /* solar_system_id integer */
	StationCount         int32   `json:"station_count,omitempty"`          /* station_count integer */
	StationSystemCount   int32   `json:"station_system_count,omitempty"`   /* station_system_count integer */
}
