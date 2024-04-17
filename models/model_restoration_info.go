/*
 * Nhss_imsUECM
 *
 * Nhss UE Context Management Service for IMS. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * Source file: 3GPP TS 29.562 HSS Services for interworking with IMS, version 16.5.0
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.562/
 *
 * API version: 1.0.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models



// The information relevant to a specific registration required for an S-CSCF to handle the requests for a user
type RestorationInfo struct {
	Path string `json:"path" yaml:"path" bson:"path"`
	Contact string `json:"contact" yaml:"contact" bson:"contact"`
	InitialCSeqSequenceNumber int32 `json:"initialCSeqSequenceNumber,omitempty" yaml:"initialCSeqSequenceNumber" bson:"initialCSeqSequenceNumber"`
	CallIdSipHeader string `json:"callIdSipHeader,omitempty" yaml:"callIdSipHeader" bson:"callIdSipHeader"`
	UesubscriptionInfo *UeSubscriptionInfo `json:"uesubscriptionInfo,omitempty" yaml:"uesubscriptionInfo" bson:"uesubscriptionInfo"`
	PcscfSubscriptionInfo *PcscfSubscriptionInfo `json:"pcscfSubscriptionInfo,omitempty" yaml:"pcscfSubscriptionInfo" bson:"pcscfSubscriptionInfo"`
	// A map (list of key-value pairs where subscriptionId serves as key) of ImsSdmSubscription
	ImsSdmSubscriptions map[string]*ImsSdmSubscription `json:"imsSdmSubscriptions,omitempty" yaml:"imsSdmSubscriptions" bson:"imsSdmSubscriptions"`
}
