/*
 * AUSF API
 *
 * AUSF UE Authentication Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.509 V17.9.0; 5G System; 3GPP TS Authentication Server services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.509
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains the result of the authentication
type ConfirmationDataResponse struct {
	AuthResult AusfUeAuthenticationAuthResult `json:"authResult" yaml:"authResult" bson:"authResult,omitempty"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501.
	Supi string `json:"supi,omitempty" yaml:"supi" bson:"supi,omitempty"`
	// Contains the Kseaf.
	Kseaf   string                 `json:"kseaf,omitempty" yaml:"kseaf" bson:"kseaf,omitempty"`
	PvsInfo []ServerAddressingInfo `json:"pvsInfo,omitempty" yaml:"pvsInfo" bson:"pvsInfo,omitempty"`
}
