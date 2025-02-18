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

/* A list of GetCorporationsCorporationIdStructures200Ok. */
//easyjson:json
type GetCorporationsCorporationIdStructures200OkList []GetCorporationsCorporationIdStructures200Ok

/* 200 ok object */
//easyjson:json
type GetCorporationsCorporationIdStructures200Ok struct {
	CorporationId      int32                                           `json:"corporation_id,omitempty"`       /* ID of the corporation that owns the structure */
	FuelExpires        time.Time                                       `json:"fuel_expires,omitempty"`         /* Date on which the structure will run out of fuel */
	Name               string                                          `json:"name,omitempty"`                 /* The structure name */
	NextReinforceApply time.Time                                       `json:"next_reinforce_apply,omitempty"` /* The date and time when the structure's newly requested reinforcement times (e.g. next_reinforce_hour and next_reinforce_day) will take effect */
	NextReinforceHour  int32                                           `json:"next_reinforce_hour,omitempty"`  /* The requested change to reinforce_hour that will take effect at the time shown by next_reinforce_apply */
	ProfileId          int32                                           `json:"profile_id,omitempty"`           /* The id of the ACL profile for this citadel */
	ReinforceHour      int32                                           `json:"reinforce_hour,omitempty"`       /* The hour of day that determines the four hour window when the structure will randomly exit its reinforcement periods and become vulnerable to attack against its armor and/or hull. The structure will become vulnerable at a random time that is +/- 2 hours centered on the value of this property */
	Services           []GetCorporationsCorporationIdStructuresService `json:"services,omitempty"`             /* Contains a list of service upgrades, and their state */
	State              string                                          `json:"state,omitempty"`                /* state string */
	StateTimerEnd      time.Time                                       `json:"state_timer_end,omitempty"`      /* Date at which the structure will move to it's next state */
	StateTimerStart    time.Time                                       `json:"state_timer_start,omitempty"`    /* Date at which the structure entered it's current state */
	StructureId        int64                                           `json:"structure_id,omitempty"`         /* The Item ID of the structure */
	SystemId           int32                                           `json:"system_id,omitempty"`            /* The solar system the structure is in */
	TypeId             int32                                           `json:"type_id,omitempty"`              /* The type id of the structure */
	UnanchorsAt        time.Time                                       `json:"unanchors_at,omitempty"`         /* Date at which the structure will unanchor */
}
