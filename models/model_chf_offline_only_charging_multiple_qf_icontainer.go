/*
 * Nchf_OfflineOnlyCharging
 *
 * OfflineOnlyCharging Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 32.291  V17.0.0: Telecommunication management; Charging management;  5G system, charging service; Stage 3.
 * Url: http://www.3gpp.org/ftp/Specs/archive/32_series/32.291/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type ChfOfflineOnlyChargingMultipleQfIcontainer struct {
	Triggers []ChfOfflineOnlyChargingTrigger `json:"triggers,omitempty" yaml:"triggers" bson:"triggers,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	TriggerTimestamp *time.Time `json:"triggerTimestamp,omitempty" yaml:"triggerTimestamp" bson:"triggerTimestamp,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	Time int32 `json:"time,omitempty" yaml:"time" bson:"time,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.
	TotalVolume int32 `json:"totalVolume,omitempty" yaml:"totalVolume" bson:"totalVolume,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.
	UplinkVolume            int32                                          `json:"uplinkVolume,omitempty" yaml:"uplinkVolume" bson:"uplinkVolume,omitempty"`
	LocalSequenceNumber     int32                                          `json:"localSequenceNumber" yaml:"localSequenceNumber" bson:"localSequenceNumber,omitempty"`
	QFIContainerInformation *ChfOfflineOnlyChargingQfiContainerInformation `json:"qFIContainerInformation,omitempty" yaml:"qFIContainerInformation" bson:"qFIContainerInformation,omitempty"`
}
