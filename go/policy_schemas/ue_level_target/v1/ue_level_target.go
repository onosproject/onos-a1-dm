// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    aPI, err := UnmarshalAPI(bytes)
//    bytes, err = aPI.Marshal()

package ueleveltargetv1

import "encoding/json"

func UnmarshalAPI(data []byte) (API, error) {
	var r API
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *API) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// O-RAN standard UE Level Target policy
type API struct {
	Scope             Scope             `json:"scope"`            
	UeLevelObjectives UeLevelObjectives `json:"ueLevelObjectives"`
}

type Scope struct {
	CellID  *CellID  `json:"cellId,omitempty"` 
	GroupID *GroupID `json:"groupId,omitempty"`
	SliceID *SliceID `json:"sliceId,omitempty"`
	UeID    string   `json:"ueId"`             
	QosID   *QosID   `json:"qosId,omitempty"`  
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

type GroupID struct {
	SPID      *int64 `json:"spId,omitempty"`     
	RfspIndex *int64 `json:"rfspIndex,omitempty"`
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

type UeLevelObjectives struct {
	DLPacketDelay           *float64         `json:"dlPacketDelay,omitempty"`          
	DLReliability           *ReliabilityType `json:"dlReliability,omitempty"`          
	DLRlcSduPacketLossRate  *float64         `json:"dlRlcSduPacketLossRate ,omitempty"`
	DLThroughput            *float64         `json:"dlThroughput,omitempty"`           
	ULPacketDelay           *float64         `json:"ulPacketDelay,omitempty"`          
	ULPdcpSduPacketLossRate *float64         `json:"ulPdcpSduPacketLossRate,omitempty"`
	ULReliability           *ReliabilityType `json:"ulReliability,omitempty"`          
	ULThroughput            *float64         `json:"ulThroughput,omitempty"`           
}

type ReliabilityType struct {
	PacketSize       float64 `json:"packetSize"`      
	SuccessProbility float64 `json:"successProbility"`
	UserPlaneLatency float64 `json:"userPlaneLatency"`
}

var RawSchema = `{ "$schema": "http://json-schema.org/draft-07/schema#", "description": "O-RAN standard UE Level Target policy", "type": "object", "properties": { "scope": { "anyOf": [ { "type": "object", "properties": { "ueId": {"$ref": "#/definitions/UeId"}, "groupId": {"$ref": "#/definitions/GroupId"}, "sliceId": {"$ref": "#/definitions/SliceId"}, "cellId": {"$ref": "#/definitions/CellId"} }, "additionalProperties": false, "required": ["ueId"] }, { "type": "object", "properties": { "ueId": {"$ref": "#/definitions/UeId"}, "groupId": {"$ref": "#/definitions/GroupId"}, "sliceId": {"$ref": "#/definitions/SliceId"}, "qosId": {"$ref": "#/definitions/QosId"}, "cellId": {"$ref": "#/definitions/CellId"} }, "additionalProperties": false, "required": ["ueId", "qosId"] }, { "type": "object", "properties": { "ueId": {"$ref": "#/definitions/UeId"}, "sliceId": {"$ref": "#/definitions/SliceId"}, "qosId": {"$ref": "#/definitions/QosId"}, "cellId": {"$ref": "#/definitions/CellId"} }, "additionalProperties": false, "required": ["ueId", "sliceId"] } ] }, "ueLevelObjectives": { "type": "object", "properties": { "ulThroughput": {"type": "number"}, "dlThroughput": {"type": "number"}, "ulPacketDelay": {"type": "number"}, "dlPacketDelay": {"type": "number"}, "ulPdcpSduPacketLossRate": {"type": "number"}, "dlRlcSduPacketLossRate ": {"type": "number"}, "dlReliability": {"$ref": "#/definitions/ReliabilityType"}, "ulReliability": {"$ref": "#/definitions/ReliabilityType"} }, "minProperties": 1, "additionalProperties": false } }, "additionalProperties": false, "required": ["scope", "ueLevelObjectives"], "definitions": { "UeId": { "type": "string", "pattern": "^[A-Fa-f0-9]{16}$" }, "GroupId": { "oneOf": [ { "type":"object", "properties": { "spId": { "type": "integer", "minimum": 1, "maximum": 256 } }, "additionalProperties": false, "required": ["spId"] }, { "type": "object", "properties": { "rfspIndex": { "type": "integer", "minimum": 1, "maximum": 256 } }, "additionalProperties": false, "required": ["rfspIndex"] } ] }, "SliceId": { "type": "object", "properties": { "sst": { "type": "integer", "minimum": 0, "maximum": 255 }, "sd": { "type": "string", "pattern": "^[A-Fa-f0-9]{6}$" }, "plmnId": {"$ref": "#/definitions/PlmnId"} }, "additionalProperties": false, "required": ["sst","plmnId"] }, "QosId": { "oneOf": [ { "type":"object", "properties": { "5qI": { "type": "integer", "minimum": 1, "maximum": 256 } }, "additionalProperties": false, "required": ["5qI"] }, { "type": "object", "properties": { "qcI": { "type": "integer", "minimum": 1, "maximum": 256 } }, "additionalProperties": false, "required": ["qcI"] } ] }, "CellId": { "type": "object", "properties": { "plmnId": {"$ref": "#/definitions/PlmnId"}, "cId": {"$ref": "#/definitions/CId"} }, "additionalProperties": false, "required": ["plmnId", "cId"] }, "CId": { "oneOf": [ { "type":"object", "properties": { "ncI": {"$ref": "#/definitions/NcI"} }, "additionalProperties": false, "required": ["ncI"] }, { "type": "object", "properties": { "ecI": {"$ref": "#/definitions/EcI"} }, "additionalProperties": false, "required": ["ecI"] } ] }, "NcI": { "type": "integer", "minimum": 0, "maximum": 68719476735 }, "EcI": { "type": "integer", "minimum": 0, "maximum": 268435455 }, "PlmnId": { "type": "object", "properties": { "mcc": { "type": "string", "pattern": "^[0-9]{3}$" }, "mnc": { "type": "string", "pattern": "^[0-9]{2,3}$" } }, "additionalProperties": false, "required": ["mcc", "mnc"] }, "ReliabilityType": { "type": "object", "properties": { "packetSize": {"type": "number"}, "userPlaneLatency": {"type": "number"}, "successProbility": {"type": "number"} }, "required": ["packetSize","userPlaneLatency","successProbility"] } } }`

var PolicyTypeID = "ORAN_UELevelTarget_1.0.0"
