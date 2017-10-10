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

/* A list of GetUniverseTypesTypeIdDogmaEffect. */
//easyjson:json
type GetUniverseTypesTypeIdDogmaEffectList []GetUniverseTypesTypeIdDogmaEffect

/* dogma_effect object */
//easyjson:json
type GetUniverseTypesTypeIdDogmaEffect struct {
	EffectId  int32 `json:"effect_id,omitempty"`  /* effect_id integer */
	IsDefault bool  `json:"is_default,omitempty"` /* is_default boolean */
}
