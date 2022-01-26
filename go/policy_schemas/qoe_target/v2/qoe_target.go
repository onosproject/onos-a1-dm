// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    aPI, err := UnmarshalAPI(bytes)
//    bytes, err = aPI.Marshal()

package qoetargetv2

import "encoding/json"

func UnmarshalAPI(data []byte) (API, error) {
	var r API
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *API) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// O-RAN standard QoE Target policy
type API struct {
	QoeObjectives QoeObjectives `json:"qoeObjectives"`
	Scope         Scope         `json:"scope"`        
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
