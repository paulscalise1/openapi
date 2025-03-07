/*
 * 3gpp-traffic-influence
 *
 * API for AF traffic influence © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models_nef

import (
	"github.com/paulscalise1/openapi/models"
)

type TrafficInfluSubPatch struct {
	// Identifies whether an application can be relocated once a location of the application has been selected.
	AppReloInd bool `json:"appReloInd,omitempty" bson:"appReloInd"`

	// Identifies IP packet filters.
	TrafficFilters []models.FlowInfo `json:"trafficFilters,omitempty" bson:"trafficFilters"`

	// Identifies Ethernet packet filters.
	EthTrafficFilters []models.EthFlowDescription `json:"ethTrafficFilters,omitempty" bson:"ethTrafficFilters"`

	// Identifies the N6 traffic routing requirement.
	TrafficRoutes []models.RouteToLocation `json:"trafficRoutes,omitempty" bson:"trafficRoutes"`

	TfcCorrInd bool `json:"tfcCorrInd,omitempty" bson:"tfcCorrInd"`

	TempValidities []models.TemporalValidity `json:"tempValidities,omitempty" bson:"tempValidities"`

	// Identifies a geographic zone that the AF request applies only to the traffic of UE(s) located in this specific zone.
	ValidGeoZoneIds []string `json:"validGeoZoneIds,omitempty" bson:"validGeoZoneIds"`

	AfAckInd bool `json:"afAckInd,omitempty" bson:"afAckInd"`

	AddrPreserInd bool `json:"addrPreserInd,omitempty" bson:"addrPreserInd"`
}
