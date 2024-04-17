/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.504 V17.12.0; 5G System; Unified Data Repository Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.504/
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type LcsPrivacy struct {
	AfInstanceId string `json:"afInstanceId,omitempty" yaml:"afInstanceId" bson:"afInstanceId,omitempty"`
	ReferenceId  int32  `json:"referenceId,omitempty" yaml:"referenceId" bson:"referenceId,omitempty"`
	Lpi          *Lpi   `json:"lpi,omitempty" yaml:"lpi" bson:"lpi,omitempty"`
	// String uniquely identifying MTC provider information.
	MtcProviderInformation string `json:"mtcProviderInformation,omitempty" yaml:"mtcProviderInformation" bson:"mtcProviderInformation,omitempty"`
}
