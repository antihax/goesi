/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.6.dev5
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

package goesiv2

/* 200 ok object */
type GetCorporationsCorporationIdOk struct {
	AllianceId      int32  `json:"alliance_id,omitempty"`      /* id of alliance that corporation is a member of, if any */
	CeoId           int32  `json:"ceo_id,omitempty"`           /* ceo_id integer */
	CorporationName string `json:"corporation_name,omitempty"` /* the full name of the corporation */
	MemberCount     int32  `json:"member_count,omitempty"`     /* member_count integer */
	Ticker          string `json:"ticker,omitempty"`           /* the short name of the corporation */
}
