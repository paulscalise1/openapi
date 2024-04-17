/*
 * Nnef_SMService
 *
 * Nnef SMService Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.541 V17.4.0; 5G System; Session Management Services for Non-IP Data Delivery (NIDD).
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.541/
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Information within response message invoking MtForwardSm service operation, for delivering MT SMS Delivery Report.
type SmsDeliveryData struct {
	SmsPayload *RefToBinaryData `json:"smsPayload" yaml:"smsPayload" bson:"smsPayload,omitempty"`
}
