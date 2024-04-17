/*
 * Npcf_AMPolicyControl
 *
 * Access and Mobility Policy Control Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.507 V17.9.0; 5G System; Access and Mobility Policy Control Service.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.507/
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Information which the NF service consumer provides when requesting the creation of a policy association. The serviveName property corresponds to the serviceName in the main body of the specification.
type PcfAmPolicyControlPolicyAssociationRequest struct {
	// String providing an URI formatted according to RFC 3986.
	NotificationUri string `json:"notificationUri" yaml:"notificationUri" bson:"notificationUri,omitempty"`
	// Alternate or backup IPv4 Address(es) where to send Notifications.
	AltNotifIpv4Addrs []string `json:"altNotifIpv4Addrs,omitempty" yaml:"altNotifIpv4Addrs" bson:"altNotifIpv4Addrs,omitempty"`
	// Alternate or backup IPv6 Address(es) where to send Notifications.
	AltNotifIpv6Addrs []string `json:"altNotifIpv6Addrs,omitempty" yaml:"altNotifIpv6Addrs" bson:"altNotifIpv6Addrs,omitempty"`
	// Alternate or backup FQDN(s) where to send Notifications.
	AltNotifFqdns []string `json:"altNotifFqdns,omitempty" yaml:"altNotifFqdns" bson:"altNotifFqdns,omitempty"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501.
	Supi string `json:"supi" yaml:"supi" bson:"supi,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi        string       `json:"gpsi,omitempty" yaml:"gpsi" bson:"gpsi,omitempty"`
	AccessType  AccessType   `json:"accessType,omitempty" yaml:"accessType" bson:"accessType,omitempty"`
	AccessTypes []AccessType `json:"accessTypes,omitempty" yaml:"accessTypes" bson:"accessTypes,omitempty"`
	// String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.
	Pei     string        `json:"pei,omitempty" yaml:"pei" bson:"pei,omitempty"`
	UserLoc *UserLocation `json:"userLoc,omitempty" yaml:"userLoc" bson:"userLoc,omitempty"`
	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.
	TimeZone      string                          `json:"timeZone,omitempty" yaml:"timeZone" bson:"timeZone,omitempty"`
	ServingPlmn   *PlmnIdNid                      `json:"servingPlmn,omitempty" yaml:"servingPlmn" bson:"servingPlmn,omitempty"`
	RatType       RatType                         `json:"ratType,omitempty" yaml:"ratType" bson:"ratType,omitempty"`
	RatTypes      []RatType                       `json:"ratTypes,omitempty" yaml:"ratTypes" bson:"ratTypes,omitempty"`
	GroupIds      []string                        `json:"groupIds,omitempty" yaml:"groupIds" bson:"groupIds,omitempty"`
	ServAreaRes   *ServiceAreaRestriction         `json:"servAreaRes,omitempty" yaml:"servAreaRes" bson:"servAreaRes,omitempty"`
	WlServAreaRes *WirelineServiceAreaRestriction `json:"wlServAreaRes,omitempty" yaml:"wlServAreaRes" bson:"wlServAreaRes,omitempty"`
	// Unsigned integer representing the \"Subscriber Profile ID for RAT/Frequency Priority\"  as specified in 3GPP TS 36.413.
	Rfsp   int32 `json:"rfsp,omitempty" yaml:"rfsp" bson:"rfsp,omitempty"`
	UeAmbr *Ambr `json:"ueAmbr,omitempty" yaml:"ueAmbr" bson:"ueAmbr,omitempty"`
	// The subscribed UE Slice-MBR for each subscribed S-NSSAI of the home PLMN mapping  to a S-NSSAI of the serving PLMN Shall be provided when available.
	UeSliceMbrs []*UeSliceMbr `json:"ueSliceMbrs,omitempty" yaml:"ueSliceMbrs" bson:"ueSliceMbrs,omitempty"`
	// array of allowed S-NSSAIs for the 3GPP access.
	AllowedSnssais []Snssai `json:"allowedSnssais,omitempty" yaml:"allowedSnssais" bson:"allowedSnssais,omitempty"`
	// array of target S-NSSAIs.
	TargetSnssais []Snssai `json:"targetSnssais,omitempty" yaml:"targetSnssais" bson:"targetSnssais,omitempty"`
	// mapping of each S-NSSAI of the Allowed NSSAI to the corresponding S-NSSAI of the HPLMN.
	MappingSnssais []MappingOfSnssai `json:"mappingSnssais,omitempty" yaml:"mappingSnssais" bson:"mappingSnssais,omitempty"`
	// array of allowed S-NSSAIs for the Non-3GPP access.
	N3gAllowedSnssais []Snssai    `json:"n3gAllowedSnssais,omitempty" yaml:"n3gAllowedSnssais" bson:"n3gAllowedSnssais,omitempty"`
	Guami             *Guami      `json:"guami,omitempty" yaml:"guami" bson:"guami,omitempty"`
	ServiveName       ServiceName `json:"serviveName,omitempty" yaml:"serviveName" bson:"serviveName,omitempty"`
	TraceReq          *TraceData  `json:"traceReq,omitempty" yaml:"traceReq" bson:"traceReq,omitempty"`
	NwdafDatas        []NwdafData `json:"nwdafDatas,omitempty" yaml:"nwdafDatas" bson:"nwdafDatas,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SuppFeat string `json:"suppFeat" yaml:"suppFeat" bson:"suppFeat,omitempty"`
}
