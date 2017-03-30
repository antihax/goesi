/* 
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.2.dev16
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

package goesiv1

import (
	"time"
)

/* 200 ok object */
type GetCharactersCharacterIdMedals200Ok struct {

	/* corporation_id integer */
	CorporationId int32 `json:"corporation_id,omitempty"`

	/* date string */
	Date time.Time `json:"date,omitempty"`

	/* description string */
	Description string `json:"description,omitempty"`

	/* graphics array */
	Graphics []GetCharactersCharacterIdMedalsGraphic `json:"graphics,omitempty"`

	/* issuer_id integer */
	IssuerId int32 `json:"issuer_id,omitempty"`

	/* medal_id integer */
	MedalId int32 `json:"medal_id,omitempty"`

	/* reason string */
	Reason string `json:"reason,omitempty"`

	/* status string */
	Status string `json:"status,omitempty"`

	/* title string */
	Title string `json:"title,omitempty"`
}
