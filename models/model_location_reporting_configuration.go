/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.504 V16.9.0; 5G System; Unified Data Repository Services; Stage 3
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.504/
 *
 * API version: 2.1.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type LocationReportingConfiguration struct {
	CurrentLocation bool             `json:"currentLocation" yaml:"currentLocation" bson:"currentLocation,omitempty"`
	OneTime         bool             `json:"oneTime,omitempty" yaml:"oneTime" bson:"oneTime,omitempty"`
	Accuracy        LocationAccuracy `json:"accuracy,omitempty" yaml:"accuracy" bson:"accuracy,omitempty"`
	N3gppAccuracy   LocationAccuracy `json:"n3gppAccuracy,omitempty" yaml:"n3gppAccuracy" bson:"n3gppAccuracy,omitempty"`
}
