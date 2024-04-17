/*
 * Npcf_MBSPolicyControl API
 *
 * MBS Policy Control Service   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.537 V17.3.0; 5G System; Multicast/Broadcast Policy Control Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.537/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Represent MBS Media Information.
type MbsMediaInfo struct {
	MbsMedType MediaType `json:"mbsMedType,omitempty" yaml:"mbsMedType" bson:"mbsMedType,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MaxReqMbsBwDl string `json:"maxReqMbsBwDl,omitempty" yaml:"maxReqMbsBwDl" bson:"maxReqMbsBwDl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MinReqMbsBwDl string   `json:"minReqMbsBwDl,omitempty" yaml:"minReqMbsBwDl" bson:"minReqMbsBwDl,omitempty"`
	Codecs        []string `json:"codecs,omitempty" yaml:"codecs" bson:"codecs,omitempty"`
}
