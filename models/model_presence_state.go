/*
 * Nchf_OfflineOnlyCharging
 *
 * OfflineOnlyCharging Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 32.291  V17.0.0: Telecommunication management; Charging management;  5G system, charging service; Stage 3.
 * Url: http://www.3gpp.org/ftp/Specs/archive/32_series/32.291/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type PresenceState string

// List of PresenceState
const (
	PresenceState_IN_AREA     PresenceState = "IN_AREA"
	PresenceState_OUT_OF_AREA PresenceState = "OUT_OF_AREA"
	PresenceState_UNKNOWN     PresenceState = "UNKNOWN"
	PresenceState_INACTIVE    PresenceState = "INACTIVE"
)
