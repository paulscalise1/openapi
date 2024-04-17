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

// AMF event subscription extended with additional information received for the subscription
type ExtAmfEventSubscription struct {
	EventList []AmfEvent `json:"eventList" yaml:"eventList" bson:"eventList,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	EventNotifyUri      string `json:"eventNotifyUri" yaml:"eventNotifyUri" bson:"eventNotifyUri,omitempty"`
	NotifyCorrelationId string `json:"notifyCorrelationId" yaml:"notifyCorrelationId" bson:"notifyCorrelationId,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	NfId string `json:"nfId" yaml:"nfId" bson:"nfId,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	SubsChangeNotifyUri           string `json:"subsChangeNotifyUri,omitempty" yaml:"subsChangeNotifyUri" bson:"subsChangeNotifyUri,omitempty"`
	SubsChangeNotifyCorrelationId string `json:"subsChangeNotifyCorrelationId,omitempty" yaml:"subsChangeNotifyCorrelationId" bson:"subsChangeNotifyCorrelationId,omitempty"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501.
	Supi string `json:"supi,omitempty" yaml:"supi" bson:"supi,omitempty"`
	// String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.
	GroupId         string   `json:"groupId,omitempty" yaml:"groupId" bson:"groupId,omitempty"`
	ExcludeSupiList []string `json:"excludeSupiList,omitempty" yaml:"excludeSupiList" bson:"excludeSupiList,omitempty"`
	ExcludeGpsiList []string `json:"excludeGpsiList,omitempty" yaml:"excludeGpsiList" bson:"excludeGpsiList,omitempty"`
	IncludeSupiList []string `json:"includeSupiList,omitempty" yaml:"includeSupiList" bson:"includeSupiList,omitempty"`
	IncludeGpsiList []string `json:"includeGpsiList,omitempty" yaml:"includeGpsiList" bson:"includeGpsiList,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi string `json:"gpsi,omitempty" yaml:"gpsi" bson:"gpsi,omitempty"`
	// String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.
	Pei               string                `json:"pei,omitempty" yaml:"pei" bson:"pei,omitempty"`
	AnyUE             bool                  `json:"anyUE,omitempty" yaml:"anyUE" bson:"anyUE,omitempty"`
	Options           *AmfEventMode         `json:"options,omitempty" yaml:"options" bson:"options,omitempty"`
	SourceNfType      NrfNfManagementNfType `json:"sourceNfType,omitempty" yaml:"sourceNfType" bson:"sourceNfType,omitempty"`
	BindingInfo       []string              `json:"bindingInfo,omitempty" yaml:"bindingInfo" bson:"bindingInfo,omitempty"`
	SubscribingNfType NrfNfManagementNfType `json:"subscribingNfType,omitempty" yaml:"subscribingNfType" bson:"subscribingNfType,omitempty"`
	EventSyncInd      bool                  `json:"eventSyncInd,omitempty" yaml:"eventSyncInd" bson:"eventSyncInd,omitempty"`
	NfConsumerInfo    []string              `json:"nfConsumerInfo,omitempty" yaml:"nfConsumerInfo" bson:"nfConsumerInfo,omitempty"`
	// Map of subscribed Area of Interest (AoI) Event State in the old AMF. The JSON pointer to an AmfEventArea element in the areaList IE (or a PresenceInfo element in  presenceInfoList IE) of the AmfEvent data type shall be the key of the map.
	AoiStateList map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty" yaml:"aoiStateList" bson:"aoiStateList,omitempty"`
}
