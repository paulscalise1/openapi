/*
 * Namf_Communication
 *
 * AMF Communication Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.518 V17.12.0; 5G System; Access and Mobility Management Services
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.518/
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// indicates the DNN and S-NSSAI combination of a PDU session.
type PduSessionInfo struct {
	Snssai *Snssai `json:"snssai" yaml:"snssai" bson:"snssai,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn string `json:"dnn" yaml:"dnn" bson:"dnn,omitempty"`
}
