/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 1.26
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

/* A list of GetMarketsPrices200Ok. */
//easyjson:json
type GetMarketsPrices200OkList []GetMarketsPrices200Ok

/* 200 ok object */
//easyjson:json
type GetMarketsPrices200Ok struct {
	AdjustedPrice float64 `json:"adjusted_price,omitempty"` /* adjusted_price number */
	AveragePrice  float64 `json:"average_price,omitempty"`  /* average_price number */
	TypeId        int32   `json:"type_id,omitempty"`        /* type_id integer */
}
