/*
 * 3gpp-traffic-influence
 *
 * API for AF traffic influence   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.522 V17.7.0; 5G System; Network Exposure Function Northbound APIs.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.522/
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Ellipsoid with uncertainty
type UncertaintyEllipsoid struct {
	// Indicates value of uncertainty.
	SemiMajor float32 `json:"semiMajor" yaml:"semiMajor" bson:"semiMajor,omitempty"`
	// Indicates value of uncertainty.
	SemiMinor float32 `json:"semiMinor" yaml:"semiMinor" bson:"semiMinor,omitempty"`
	// Indicates value of uncertainty.
	Vertical float32 `json:"vertical" yaml:"vertical" bson:"vertical,omitempty"`
	// Indicates value of orientation angle.
	OrientationMajor int32 `json:"orientationMajor" yaml:"orientationMajor" bson:"orientationMajor,omitempty"`
}
