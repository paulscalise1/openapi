/*
 * 3gpp-pfd-management
 *
 * API for PFD management. © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models_nef

type FailureCode string

// List of FailureCode
const (
	FailureCode_MALFUNCTION         FailureCode = "MALFUNCTION"
	FailureCode_RESOURCE_LIMITATION FailureCode = "RESOURCE_LIMITATION"
	FailureCode_SHORT_DELAY         FailureCode = "SHORT_DELAY"
	FailureCode_APP_ID_DUPLICATED   FailureCode = "APP_ID_DUPLICATED"
	FailureCode_PARTIAL_FAILURE     FailureCode = "PARTIAL_FAILURE"
	FailureCode_OTHER_REASON        FailureCode = "OTHER_REASON"
)
