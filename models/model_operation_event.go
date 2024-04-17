/*
 * Namf_MBSBroadcast
 *
 * AMF MBSBroadcast Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.518 V17.12.0; 5G System; Access and Mobility Management Services
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.518/
 *
 * API version: 1.0.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Operation Event for a Broadcast MBS Session.
type OperationEvent struct {
	OpEventType OpEventType `json:"opEventType" yaml:"opEventType" bson:"opEventType,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	AmfId                 string              `json:"amfId,omitempty" yaml:"amfId" bson:"amfId,omitempty"`
	NgranFailureEventList []NgranFailureEvent `json:"ngranFailureEventList,omitempty" yaml:"ngranFailureEventList" bson:"ngranFailureEventList,omitempty"`
}
