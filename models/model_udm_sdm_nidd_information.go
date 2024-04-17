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

type UdmSdmNiddInformation struct {
	AfId string `json:"afId" yaml:"afId" bson:"afId,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi string `json:"gpsi,omitempty" yaml:"gpsi" bson:"gpsi,omitempty"`
	// String identifying External Group Identifier that identifies a group made up of one or more  subscriptions associated to a group of IMSIs, as specified in clause 19.7.3 of 3GPP TS 23.003.
	ExtGroupId string `json:"extGroupId,omitempty" yaml:"extGroupId" bson:"extGroupId,omitempty"`
}
