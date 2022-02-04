// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package policyschemas

import (
	qoeandtspv2 "github.com/onosproject/onos-a1-dm/go/policy_schemas/qoe_and_tsp/v2"
	qoetargetv2 "github.com/onosproject/onos-a1-dm/go/policy_schemas/qoe_target/v2"
	qosandtspv2 "github.com/onosproject/onos-a1-dm/go/policy_schemas/qos_and_tsp/v2"
	qostargetv2 "github.com/onosproject/onos-a1-dm/go/policy_schemas/qos_target/v2"
	sliceslatargetv1 "github.com/onosproject/onos-a1-dm/go/policy_schemas/slice_sla_target/v1"
	tspv2 "github.com/onosproject/onos-a1-dm/go/policy_schemas/traffic_steering_preference/v2"
	ueleveltargetv1 "github.com/onosproject/onos-a1-dm/go/policy_schemas/ue_level_target/v1"
)

var PolicySchemas = map[string]string {
	qoeandtspv2.PolicyTypeID: qoeandtspv2.RawSchema,
	qoetargetv2.PolicyTypeID: qoetargetv2.RawSchema,
	qosandtspv2.PolicyTypeID: qosandtspv2.RawSchema,
	qostargetv2.PolicyTypeID: qostargetv2.RawSchema,
	sliceslatargetv1.PolicyTypeID: sliceslatargetv1.RawSchema,
	tspv2.PolicyTypeID: tspv2.RawSchema,
	ueleveltargetv1.PolicyTypeID: ueleveltargetv1.RawSchema,
}