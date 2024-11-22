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

import (
	"time"
)

/* A list of GetCorporationsCorporationIdIndustryJobs200Ok. */
//easyjson:json
type GetCorporationsCorporationIdIndustryJobs200OkList []GetCorporationsCorporationIdIndustryJobs200Ok

/* 200 ok object */
//easyjson:json
type GetCorporationsCorporationIdIndustryJobs200Ok struct {
	ActivityId           int32     `json:"activity_id,omitempty"`            /* Job activity ID */
	BlueprintId          int64     `json:"blueprint_id,omitempty"`           /* blueprint_id integer */
	BlueprintLocationId  int64     `json:"blueprint_location_id,omitempty"`  /* Location ID of the location from which the blueprint was installed. Normally a station ID, but can also be an asset (e.g. container) or corporation facility */
	BlueprintTypeId      int32     `json:"blueprint_type_id,omitempty"`      /* blueprint_type_id integer */
	CompletedCharacterId int32     `json:"completed_character_id,omitempty"` /* ID of the character which completed this job */
	CompletedDate        time.Time `json:"completed_date,omitempty"`         /* Date and time when this job was completed */
	Cost                 float64   `json:"cost,omitempty"`                   /* The sume of job installation fee and industry facility tax */
	Duration             int32     `json:"duration,omitempty"`               /* Job duration in seconds */
	EndDate              time.Time `json:"end_date,omitempty"`               /* Date and time when this job finished */
	FacilityId           int64     `json:"facility_id,omitempty"`            /* ID of the facility where this job is running */
	InstallerId          int32     `json:"installer_id,omitempty"`           /* ID of the character which installed this job */
	JobId                int32     `json:"job_id,omitempty"`                 /* Unique job ID */
	LicensedRuns         int32     `json:"licensed_runs,omitempty"`          /* Number of runs blueprint is licensed for */
	LocationId           int64     `json:"location_id,omitempty"`            /* ID of the location for the industry facility */
	OutputLocationId     int64     `json:"output_location_id,omitempty"`     /* Location ID of the location to which the output of the job will be delivered. Normally a station ID, but can also be a corporation facility */
	PauseDate            time.Time `json:"pause_date,omitempty"`             /* Date and time when this job was paused (i.e. time when the facility where this job was installed went offline) */
	Probability          float32   `json:"probability,omitempty"`            /* Chance of success for invention */
	ProductTypeId        int32     `json:"product_type_id,omitempty"`        /* Type ID of product (manufactured, copied or invented) */
	Runs                 int32     `json:"runs,omitempty"`                   /* Number of runs for a manufacturing job, or number of copies to make for a blueprint copy */
	StartDate            time.Time `json:"start_date,omitempty"`             /* Date and time when this job started */
	Status               string    `json:"status,omitempty"`                 /* status string */
	SuccessfulRuns       int32     `json:"successful_runs,omitempty"`        /* Number of successful runs for this job. Equal to runs unless this is an invention job */
}
