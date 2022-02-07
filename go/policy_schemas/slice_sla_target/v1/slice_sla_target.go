// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    aPI, err := UnmarshalAPI(bytes)
//    bytes, err = aPI.Marshal()

package sliceslatargetv1

import "encoding/json"

func UnmarshalAPI(data []byte) (API, error) {
	var r API
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *API) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// O-RAN standard slice SLA policy
type API struct {
	Scope              Scope              `json:"scope"`                      
	SliceSlaObjectives SliceSlaObjectives `json:"sliceSlaObjectives"`         
	SliceSlaResources  *SliceSlaResources `json:"sliceSlaResources,omitempty"`
}

type Scope struct {
	SliceID SliceID `json:"sliceId"`
}

type SliceID struct {
	PlmnID PlmnID  `json:"plmnId"`      
	SD     *string `json:"sd,omitempty"`
	Sst    int64   `json:"sst"`         
}

type PlmnID struct {
	Mcc string `json:"mcc"`
	Mnc string `json:"mnc"`
}

type SliceSlaObjectives struct {
	GuaDLThptPerSlice      *float64 `json:"guaDlThptPerSlice,omitempty"`     
	GuaULThptPerSlice      *float64 `json:"guaUlThptPerSlice,omitempty"`     
	MaxDLThptPerSlice      *float64 `json:"maxDlThptPerSlice,omitempty"`     
	MaxDLThptPerUe         *float64 `json:"maxDlThptPerUe,omitempty"`        
	MaxNumberOfPduSessions *float64 `json:"maxNumberOfPduSessions,omitempty"`
	MaxNumberOfUes         *float64 `json:"maxNumberOfUes,omitempty"`        
	MaxULThptPerSlice      *float64 `json:"maxUlThptPerSlice,omitempty"`     
	MaxULThptPerUe         *float64 `json:"maxUlThptPerUe,omitempty"`        
}

type SliceSlaResources struct {
	CellIDList []CellID `json:"cellIdList,omitempty"`
	TaIList    []TaI    `json:"taIList,omitempty"`   
}

type CellID struct {
	CID    CID    `json:"cId"`   
	PlmnID PlmnID `json:"plmnId"`
}

type CID struct {
	NcI *int64 `json:"ncI,omitempty"`
	EcI *int64 `json:"ecI,omitempty"`
}

type TaI struct {
	PlmnID PlmnID `json:"plmnId"`
	Tac    string `json:"tac"`   
}

var RawSchema = `{ "$schema": "http://json-schema.org/draft-07/schema#", "description": "O-RAN standard slice SLA policy", "type": "object", "properties": { "scope": { "type": "object", "properties": { "sliceId": {"$ref": "#/definitions/SliceId"} }, "additionalProperties": false, "required": ["sliceId"] }, "sliceSlaObjectives": { "type": "object", "properties": { "maxNumberOfUes": {"type": "number"}, "maxNumberOfPduSessions": {"type": "number"}, "guaDlThptPerSlice": {"type": "number"}, "maxDlThptPerSlice": {"type": "number"}, "maxDlThptPerUe": {"type": "number"}, "guaUlThptPerSlice": {"type": "number"}, "maxUlThptPerSlice": {"type": "number"}, "maxUlThptPerUe": {"type": "number"} }, "minProperties": 1, "additionalProperties": false }, "sliceSlaResources": { "type": "object", "properties": { "cellIdList": {"$ref": "#/definitions/CellIdList"}, "taIList": {"$ref": "#/definitions/TaIList"} }, "additionalProperties": false, "oneOf": [ {"required": ["cellIdList"]}, {"required": ["taIList"]} ] } }, "additionalProperties": false, "required": ["scope", "sliceSlaObjectives"], "definitions": { "SliceId": { "type": "object", "properties": { "sst": { "type": "integer", "minimum": 0, "maximum": 255 }, "sd": { "type": "string", "pattern": "^[A-Fa-f0-9]{6}$" }, "plmnId": {"$ref": "#/definitions/PlmnId"} }, "additionalProperties": false, "required": ["sst","plmnId"] }, "CellId": { "type": "object", "properties": { "plmnId": {"$ref": "#/definitions/PlmnId"}, "cId": {"$ref": "#/definitions/CId"} }, "additionalProperties": false, "required": ["plmnId", "cId"] }, "CId": { "oneOf": [ { "type":"object", "properties": { "ncI": {"$ref": "#/definitions/NcI"} }, "additionalProperties": false, "required": ["ncI"] }, { "type": "object", "properties": { "ecI": {"$ref": "#/definitions/EcI"} }, "additionalProperties": false, "required": ["ecI"] } ] }, "NcI": { "type": "integer", "minimum": 0, "maximum": 68719476735 }, "EcI": { "type": "integer", "minimum": 0, "maximum": 268435455 }, "PlmnId": { "type": "object", "properties": { "mcc": { "type": "string", "pattern": "^[0-9]{3}$" }, "mnc": { "type": "string", "pattern": "^[0-9]{2,3}$" } }, "additionalProperties": false, "required": ["mcc", "mnc"] }, "TaI": { "type": "object", "properties": { "plmnId": {"$ref": "#/definitions/PlmnId"}, "tac": { "type": "string", "pattern": "^[A-Fa-f0-9]{6}$" } }, "additionalProperties": false, "required": ["plmnId","tac"] }, "CellIdList": { "type":"array", "items": { "$ref": "#/definitions/CellId" } }, "TaIList": { "type":"array", "items": { "$ref": "#/definitions/TaI" } } } }`

var PolicyTypeID = "ORAN_SliceSLATarget_1.0.0"
