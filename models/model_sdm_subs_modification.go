/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.503 Unified Data Management Services, version 17.13.0
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.503/
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type SdmSubsModification struct {
	// string with format 'date-time' as defined in OpenAPI.
	Expires               *time.Time `json:"expires,omitempty" yaml:"expires" bson:"expires,omitempty"`
	MonitoredResourceUris []string   `json:"monitoredResourceUris,omitempty" yaml:"monitoredResourceUris" bson:"monitoredResourceUris,omitempty"`
}
