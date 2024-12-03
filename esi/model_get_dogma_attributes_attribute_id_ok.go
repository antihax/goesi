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

/* A list of GetDogmaAttributesAttributeIdOk. */
//easyjson:json
type GetDogmaAttributesAttributeIdOkList []GetDogmaAttributesAttributeIdOk

/* 200 ok object */
//easyjson:json
type GetDogmaAttributesAttributeIdOk struct {
	AttributeId  int32   `json:"attribute_id,omitempty"`  /* attribute_id integer */
	DefaultValue float32 `json:"default_value,omitempty"` /* default_value number */
	Description  string  `json:"description,omitempty"`   /* description string */
	DisplayName  string  `json:"display_name,omitempty"`  /* display_name string */
	HighIsGood   bool    `json:"high_is_good,omitempty"`  /* high_is_good boolean */
	IconId       int32   `json:"icon_id,omitempty"`       /* icon_id integer */
	Name         string  `json:"name,omitempty"`          /* name string */
	Published    bool    `json:"published,omitempty"`     /* published boolean */
	Stackable    bool    `json:"stackable,omitempty"`     /* stackable boolean */
	UnitId       int32   `json:"unit_id,omitempty"`       /* unit_id integer */
}
