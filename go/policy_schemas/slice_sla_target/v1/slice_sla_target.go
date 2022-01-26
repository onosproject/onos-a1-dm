// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

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
