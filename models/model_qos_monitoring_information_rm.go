/*
 * 3gpp-as-session-with-qos
 *
 * API for setting us an AS session with required QoS.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.122 V17.9.0 T8 reference point for Northbound APIs
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.122/
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Represents the same as the QosMonitoringInformation data type but with the nullable:true property.
type QosMonitoringInformationRm struct {
	ReqQosMonParams []RequestedQosMonitoringParameter `json:"reqQosMonParams,omitempty" yaml:"reqQosMonParams" bson:"reqQosMonParams,omitempty"`
	RepFreqs        []ReportingFrequency              `json:"repFreqs,omitempty" yaml:"repFreqs" bson:"repFreqs,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI 'nullable: true' property.
	RepThreshDl int32 `json:"repThreshDl,omitempty" yaml:"repThreshDl" bson:"repThreshDl,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI 'nullable: true' property.
	RepThreshUl int32 `json:"repThreshUl,omitempty" yaml:"repThreshUl" bson:"repThreshUl,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI 'nullable: true' property.
	RepThreshRp int32 `json:"repThreshRp,omitempty" yaml:"repThreshRp" bson:"repThreshRp,omitempty"`
	// Unsigned integer identifying a period of time in units of seconds with \"nullable=true\" property.
	WaitTime int32 `json:"waitTime,omitempty" yaml:"waitTime" bson:"waitTime,omitempty"`
	// Unsigned integer identifying a period of time in units of seconds with \"nullable=true\" property.
	RepPeriod int32 `json:"repPeriod,omitempty" yaml:"repPeriod" bson:"repPeriod,omitempty"`
}
