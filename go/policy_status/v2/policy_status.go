// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    aPI, err := UnmarshalAPI(bytes)
//    bytes, err = aPI.Marshal()

package policystatusv2

import "encoding/json"

func UnmarshalAPI(data []byte) (API, error) {
	var r API
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *API) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// O-RAN standard policy status
type API struct {
	EnforceReason *EnforceReason `json:"enforceReason,omitempty"`
	EnforceStatus EnforceStatus  `json:"enforceStatus"`          
}

type EnforceReason string
const (
	OtherReason EnforceReason = "OTHER_REASON"
	ScopeNotApplicable EnforceReason = "SCOPE_NOT_APPLICABLE"
	StatementNotApplicable EnforceReason = "STATEMENT_NOT_APPLICABLE"
)

type EnforceStatus string
const (
	Enforced EnforceStatus = "ENFORCED"
	NotEnforced EnforceStatus = "NOT_ENFORCED"
)

var RawSchema = `{ "$schema": "http://json-schema.org/draft-07/schema#", "description": "O-RAN standard policy status", "type": "object", "properties": { "enforceStatus": { "type": "string", "enum": [ "ENFORCED", "NOT_ENFORCED" ] }, "enforceReason": { "type": "string", "enum": [ "SCOPE_NOT_APPLICABLE", "STATEMENT_NOT_APPLICABLE", "OTHER_REASON" ] } }, "additionalProperties": false, "required": ["enforceStatus"] }`
