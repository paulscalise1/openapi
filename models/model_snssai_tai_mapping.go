/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.9.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// List of restricted or unrestricted S-NSSAIs per TAI(s)
type SnssaiTaiMapping struct {
	ReportingArea       *TargetArea       `json:"reportingArea" yaml:"reportingArea" bson:"reportingArea,omitempty"`
	AccessTypeList      []AccessType      `json:"accessTypeList,omitempty" yaml:"accessTypeList" bson:"accessTypeList,omitempty"`
	SupportedSnssaiList []SupportedSnssai `json:"supportedSnssaiList,omitempty" yaml:"supportedSnssaiList" bson:"supportedSnssaiList,omitempty"`
}
