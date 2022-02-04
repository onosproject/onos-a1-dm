// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    aPI, err := UnmarshalAPI(bytes)
//    bytes, err = aPI.Marshal()

package qoeandtspv2

import "encoding/json"

func UnmarshalAPI(data []byte) (API, error) {
	var r API
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *API) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// O-RAN standard QoE Target and Traffic Steering Preference policy
type API struct {
	QoeObjectives QoeObjectives `json:"qoeObjectives"`
	Scope         Scope         `json:"scope"`        
	TSPResources  []TSPResource `json:"tspResources"` 
}

type QoeObjectives struct {
	InitialBuffering *float64 `json:"initialBuffering,omitempty"`
	QoeScore         *float64 `json:"qoeScore,omitempty"`        
	ReBuffFreq       *float64 `json:"reBuffFreq,omitempty"`      
	StallRatio       *float64 `json:"stallRatio,omitempty"`      
}

type Scope struct {
	CellID  *CellID  `json:"cellId,omitempty"` 
	QosID   *QosID   `json:"qosId,omitempty"`  
	SliceID *SliceID `json:"sliceId,omitempty"`
	UeID    *string  `json:"ueId,omitempty"`   
}

type CellID struct {
	CID    CID    `json:"cId"`   
	PlmnID PlmnID `json:"plmnId"`
}

type CID struct {
	NcI *int64 `json:"ncI,omitempty"`
	EcI *int64 `json:"ecI,omitempty"`
}

type PlmnID struct {
	Mcc string `json:"mcc"`
	Mnc string `json:"mnc"`
}

type QosID struct {
	The5QI *int64 `json:"5qI,omitempty"`
	QcI    *int64 `json:"qcI,omitempty"`
}

type SliceID struct {
	PlmnID PlmnID  `json:"plmnId"`      
	SD     *string `json:"sd,omitempty"`
	Sst    int64   `json:"sst"`         
}

type TSPResource struct {
	CellIDList []CellID       `json:"cellIdList"`       
	Preference PreferenceType `json:"preference"`       
	Primary    *bool          `json:"primary,omitempty"`
}

type PreferenceType string
const (
	Avoid PreferenceType = "AVOID"
	Forbid PreferenceType = "FORBID"
	Prefer PreferenceType = "PREFER"
	Shall PreferenceType = "SHALL"
)

var RawSchema = `{ "$schema": "http://json-schema.org/draft-07/schema#", "description": "O-RAN standard QoE Target and Traffic Steering Preference policy", "type": "object", "properties": { "scope": { "anyOf": [ { "type": "object", "properties": { "ueId": {"$ref": "#/definitions/UeId"}, "sliceId": {"$ref": "#/definitions/SliceId"}, "qosId": {"$ref": "#/definitions/QosId"}, "cellId": {"$ref": "#/definitions/CellId"} }, "additionalProperties": false, "required": ["ueId", "sliceId"] }, { "type": "object", "properties": { "ueId": {"$ref": "#/definitions/UeId"}, "qosId": {"$ref": "#/definitions/QosId"}, "cellId": {"$ref": "#/definitions/CellId"} }, "additionalProperties": false, "required": ["ueId", "qosId"] }, { "type": "object", "properties": { "sliceId": {"$ref": "#/definitions/SliceId"}, "qosId": {"$ref": "#/definitions/QosId"}, "cellId": {"$ref": "#/definitions/CellId"} }, "additionalProperties": false, "required": ["sliceId"] } ] }, "qoeObjectives": { "type": "object", "properties": { "qoeScore": {"type": "number"}, "initialBuffering": {"type": "number"}, "reBuffFreq": {"type": "number"}, "stallRatio": {"type": "number"} }, "minProperties": 1, "additionalProperties": false }, "tspResources": { "type": "array", "items": { "$ref": "#/definitions/TspResource" }, "minItems": 1 } }, "additionalProperties": false, "required": ["scope", "qoeObjectives", "tspResources"], "definitions": { "UeId": { "type": "string", "pattern": "^[A-Fa-f0-9]{16}$" }, "SliceId": { "type": "object", "properties": { "sst": { "type": "integer", "minimum": 0, "maximum": 255 }, "sd": { "type": "string", "pattern": "^[A-Fa-f0-9]{6}$" }, "plmnId": {"$ref": "#/definitions/PlmnId"} }, "additionalProperties": false, "required": ["sst","plmnId"] }, "QosId": { "oneOf": [ { "type":"object", "properties": { "5qI": { "type": "integer", "minimum": 1, "maximum": 256 } }, "additionalProperties": false, "required": ["5qI"] }, { "type": "object", "properties": { "qcI": { "type": "integer", "minimum": 1, "maximum": 256 } }, "additionalProperties": false, "required": ["qcI"] } ] }, "CellId": { "type": "object", "properties": { "plmnId": {"$ref": "#/definitions/PlmnId"}, "cId": {"$ref": "#/definitions/CId"} }, "additionalProperties": false, "required": ["plmnId", "cId"] }, "CId": { "oneOf": [ { "type":"object", "properties": { "ncI": {"$ref": "#/definitions/NcI"} }, "additionalProperties": false, "required": ["ncI"] }, { "type": "object", "properties": { "ecI": {"$ref": "#/definitions/EcI"} }, "additionalProperties": false, "required": ["ecI"] } ] }, "NcI": { "type": "integer", "minimum": 0, "maximum": 68719476735 }, "EcI": { "type": "integer", "minimum": 0, "maximum": 268435455 }, "PlmnId": { "type": "object", "properties": { "mcc": { "type": "string", "pattern": "^[0-9]{3}$" }, "mnc": { "type": "string", "pattern": "^[0-9]{2,3}$" } }, "additionalProperties": false, "required": ["mcc", "mnc"] }, "PreferenceType": { "type": "string", "enum": [ "SHALL", "PREFER", "AVOID", "FORBID" ] }, "CellIdList": { "type":"array", "items": { "$ref": "#/definitions/CellId" } }, "TspResource": { "type": "object", "properties": { "cellIdList": {"$ref": "#/definitions/CellIdList"}, "preference": {"$ref": "#/definitions/PreferenceType"}, "primary": {"type": "boolean"} }, "required": ["cellIdList", "preference"], "additionalProperties": false } } }`

var PolicyTypeID = "ORAN_QoEandTSP_2.0.0"
