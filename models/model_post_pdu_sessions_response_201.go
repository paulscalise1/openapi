/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.502 V17.11.0; 5G System; Session Management Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.502/
 *
 * API version: 1.2.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type PostPduSessionsResponse201 struct {
	JsonData               *PduSessionCreatedData `json:"jsonData,omitempty" yaml:"jsonData" bson:"jsonData,omitempty"`
	BinaryDataN1SmInfoToUe []byte                 `json:"binaryDataN1SmInfoToUe,omitempty" yaml:"binaryDataN1SmInfoToUe" bson:"binaryDataN1SmInfoToUe,omitempty"`
}
