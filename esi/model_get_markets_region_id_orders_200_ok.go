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

/* A list of GetMarketsRegionIdOrders200Ok. */
//easyjson:json
type GetMarketsRegionIdOrders200OkList []GetMarketsRegionIdOrders200Ok

/* 200 ok object */
//easyjson:json
type GetMarketsRegionIdOrders200Ok struct {
	Duration     int32     `json:"duration,omitempty"`      /* duration integer */
	IsBuyOrder   bool      `json:"is_buy_order,omitempty"`  /* is_buy_order boolean */
	Issued       time.Time `json:"issued,omitempty"`        /* issued string */
	LocationId   int64     `json:"location_id,omitempty"`   /* location_id integer */
	MinVolume    int32     `json:"min_volume,omitempty"`    /* min_volume integer */
	OrderId      int64     `json:"order_id,omitempty"`      /* order_id integer */
	Price        float64   `json:"price,omitempty"`         /* price number */
	Range_       string    `json:"range,omitempty"`         /* range string */
	SystemId     int32     `json:"system_id,omitempty"`     /* The solar system this order was placed */
	TypeId       int32     `json:"type_id,omitempty"`       /* type_id integer */
	VolumeRemain int32     `json:"volume_remain,omitempty"` /* volume_remain integer */
	VolumeTotal  int32     `json:"volume_total,omitempty"`  /* volume_total integer */
}
