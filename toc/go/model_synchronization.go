/*
 * TOC API
 *
 * An API to create backup and restoration
 *
 * API version: 1.0.0
 * Contact: 147011991+gyk4j@users.noreply.github.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type Synchronization struct {

	Id int64 `json:"id,omitempty"`

	Time time.Time `json:"time,omitempty"`

	// Synchronization Status
	Status string `json:"status,omitempty"`
}
