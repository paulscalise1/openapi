/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.512 V17.11.0; 5G System; Session Management Policy Control Service.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains parameters determining how flows associated with a PCC Rule are treated (e.g. blocked, redirected, etc).
type TrafficControlData struct {
	// Univocally identifies the traffic control policy data within a PDU session.
	TcId            string                `json:"tcId" yaml:"tcId" bson:"tcId,omitempty"`
	FlowStatus      FlowStatus            `json:"flowStatus,omitempty" yaml:"flowStatus" bson:"flowStatus,omitempty"`
	RedirectInfo    *RedirectInformation  `json:"redirectInfo,omitempty" yaml:"redirectInfo" bson:"redirectInfo,omitempty"`
	AddRedirectInfo []RedirectInformation `json:"addRedirectInfo,omitempty" yaml:"addRedirectInfo" bson:"addRedirectInfo,omitempty"`
	// Indicates whether applicat'on's start or stop notification is to be muted.
	MuteNotif bool `json:"muteNotif,omitempty" yaml:"muteNotif" bson:"muteNotif,omitempty"`
	// Reference to a pre-configured traffic steering policy for downlink traffic at the SMF.
	TrafficSteeringPolIdDl string `json:"trafficSteeringPolIdDl,omitempty" yaml:"trafficSteeringPolIdDl" bson:"trafficSteeringPolIdDl,omitempty"`
	// Reference to a pre-configured traffic steering policy for uplink traffic at the SMF.
	TrafficSteeringPolIdUl string `json:"trafficSteeringPolIdUl,omitempty" yaml:"trafficSteeringPolIdUl" bson:"trafficSteeringPolIdUl,omitempty"`
	// A list of location which the traffic shall be routed to for the AF request
	RouteToLocs []*RouteToLocation `json:"routeToLocs,omitempty" yaml:"routeToLocs" bson:"routeToLocs,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI 'nullable: true' property.
	MaxAllowedUpLat int32 `json:"maxAllowedUpLat,omitempty" yaml:"maxAllowedUpLat" bson:"maxAllowedUpLat,omitempty"`
	// Contains EAS IP replacement information.
	EasIpReplaceInfos []EasIpReplacementInfo `json:"easIpReplaceInfos,omitempty" yaml:"easIpReplaceInfos" bson:"easIpReplaceInfos,omitempty"`
	TraffCorreInd     bool                   `json:"traffCorreInd,omitempty" yaml:"traffCorreInd" bson:"traffCorreInd,omitempty"`
	// Indicates whether simultaneous connectivity should be temporarily maintained for the source and target PSA.
	SimConnInd bool `json:"simConnInd,omitempty" yaml:"simConnInd" bson:"simConnInd,omitempty"`
	// indicating a time in seconds.
	SimConnTerm    int32                                    `json:"simConnTerm,omitempty" yaml:"simConnTerm" bson:"simConnTerm,omitempty"`
	UpPathChgEvent *UpPathChgEvent                          `json:"upPathChgEvent,omitempty" yaml:"upPathChgEvent" bson:"upPathChgEvent,omitempty"`
	SteerFun       SteeringFunctionality                    `json:"steerFun,omitempty" yaml:"steerFun" bson:"steerFun,omitempty"`
	SteerModeDl    *SteeringMode                            `json:"steerModeDl,omitempty" yaml:"steerModeDl" bson:"steerModeDl,omitempty"`
	SteerModeUl    *SteeringMode                            `json:"steerModeUl,omitempty" yaml:"steerModeUl" bson:"steerModeUl,omitempty"`
	MulAccCtrl     PcfSmPolicyControlMulticastAccessControl `json:"mulAccCtrl,omitempty" yaml:"mulAccCtrl" bson:"mulAccCtrl,omitempty"`
}
