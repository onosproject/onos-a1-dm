#!/bin/bash

# SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

export IMAGE_VER=latest

pushd ../docker/quicktype
docker build . -t onosproject/quicktype:${IMAGE_VER}
popd

# Commons for A1-P

cat ../template/license_header.txt > ../../go/policy_status/v2/policy_status.go
SCHEMA=$(cat ../../api/policy_management/policy_status.json)
docker run -v $(pwd)/../../api/policy_management/policy_status.json:/api.json onosproject/quicktype:${IMAGE_VER} /api.json -l go -s schema --package policystatusv2>> ../../go/policy_status/v2/policy_status.go
echo >> ../../go/policy_status/v2/policy_status.go
echo var RawSchema = \`$SCHEMA\` >> ../../go/policy_status/v2/policy_status.go

cat ../template/license_header.txt > ../../go/policy_scope_id/v2/policy_scope_identifier.go
SCHEMA=$(cat ../../api/policy_management/policy_scope_identifier.json)
docker run -v $(pwd)/../../api/policy_management/policy_scope_identifier.json:/api.json onosproject/quicktype:${IMAGE_VER} /api.json -l go -s schema --package policyscopeidv2 >> ../../go/policy_scope_id/v2/policy_scope_identifier.go
echo >> ../../go/policy_scope_id/v2/policy_scope_identifier.go
echo var RawSchema = \`$SCHEMA\` >> ../../go/policy_scope_id/v2/policy_scope_identifier.go

# PolicyTypes

cat ../template/license_header.txt > ../../go/policy_schemas/qoe_and_tsp/v2/qoe_and_tsp.go
SCHEMA=$(cat ../../api/policy_management/policy_schema_qoe_and_tsp.json)
docker run -v $(pwd)/../../api/policy_management/policy_schema_qoe_and_tsp.json:/api.json onosproject/quicktype:${IMAGE_VER} /api.json -l go -s schema --package qoeandtspv2 >> ../../go/policy_schemas/qoe_and_tsp/v2/qoe_and_tsp.go
echo >> ../../go/policy_schemas/qoe_and_tsp/v2/qoe_and_tsp.go
echo var RawSchema = \`$SCHEMA\` >> ../../go/policy_schemas/qoe_and_tsp/v2/qoe_and_tsp.go
echo >> ../../go/policy_schemas/qoe_and_tsp/v2/qoe_and_tsp.go
echo var PolicyTypeID = \"ORAN_QoEandTSP_2.0.0\" >> ../../go/policy_schemas/qoe_and_tsp/v2/qoe_and_tsp.go

cat ../template/license_header.txt > ../../go/policy_schemas/qoe_target/v2/qoe_target.go
SCHEMA=$(cat ../../api/policy_management/policy_schema_qoe_target.json)
docker run -v $(pwd)/../../api/policy_management/policy_schema_qoe_target.json:/api.json onosproject/quicktype:${IMAGE_VER} /api.json -l go -s schema --package qoetargetv2 >> ../../go/policy_schemas/qoe_target/v2/qoe_target.go
echo >> ../../go/policy_schemas/qoe_target/v2/qoe_target.go
echo var RawSchema = \`$SCHEMA\` >> ../../go/policy_schemas/qoe_target/v2/qoe_target.go
echo >> ../../go/policy_schemas/qoe_target/v2/qoe_target.go
echo var PolicyTypeID = \"ORAN_QoETarget_2.0.0\" >> ../../go/policy_schemas/qoe_target/v2/qoe_target.go

cat ../template/license_header.txt > ../../go/policy_schemas/qos_and_tsp/v2/qos_and_tsp.go
SCHEMA=$(cat ../../api/policy_management/policy_schema_qos_and_tsp.json)
docker run -v $(pwd)/../../api/policy_management/policy_schema_qos_and_tsp.json:/api.json onosproject/quicktype:${IMAGE_VER} /api.json -l go -s schema --package qosandtspv2 >> ../../go/policy_schemas/qos_and_tsp/v2/qos_and_tsp.go
echo >> ../../go/policy_schemas/qos_and_tsp/v2/qos_and_tsp.go
echo var RawSchema = \`$SCHEMA\` >> ../../go/policy_schemas/qos_and_tsp/v2/qos_and_tsp.go
echo >> ../../go/policy_schemas/qos_and_tsp/v2/qos_and_tsp.go
echo var PolicyTypeID = \"ORAN_QoSandTSP_2.0.0\" >> ../../go/policy_schemas/qos_and_tsp/v2/qos_and_tsp.go

cat ../template/license_header.txt > ../../go/policy_schemas/qos_target/v2/qos_target.go
SCHEMA=$(cat ../../api/policy_management/policy_schema_qos_target.json)
docker run -v $(pwd)/../../api/policy_management/policy_schema_qos_target.json:/api.json onosproject/quicktype:${IMAGE_VER} /api.json -l go -s schema --package qostargetv2 >> ../../go/policy_schemas/qos_target/v2/qos_target.go
echo >> ../../go/policy_schemas/qos_target/v2/qos_target.go
echo var RawSchema = \`$SCHEMA\` >> ../../go/policy_schemas/qos_target/v2/qos_target.go
echo >> ../../go/policy_schemas/qos_target/v2/qos_target.go
echo var PolicyTypeID = \"ORAN_QoSTarget_2.0.0\" >> ../../go/policy_schemas/qos_target/v2/qos_target.go

cat ../template/license_header.txt > ../../go/policy_schemas/slice_sla_target/v1/slice_sla_target.go
SCHEMA=$(cat ../../api/policy_management/policy_schema_slice_sla_target.json)
docker run -v $(pwd)/../../api/policy_management/policy_schema_slice_sla_target.json:/api.json onosproject/quicktype:${IMAGE_VER} /api.json -l go -s schema --package sliceslatargetv1 >> ../../go/policy_schemas/slice_sla_target/v1/slice_sla_target.go
echo >> ../../go/policy_schemas/slice_sla_target/v1/slice_sla_target.go
echo var RawSchema = \`$SCHEMA\` >> ../../go/policy_schemas/slice_sla_target/v1/slice_sla_target.go
echo >> ../../go/policy_schemas/slice_sla_target/v1/slice_sla_target.go
echo var PolicyTypeID = \"ORAN_SliceSLATarget_1.0.0\" >> ../../go/policy_schemas/slice_sla_target/v1/slice_sla_target.go

cat ../template/license_header.txt > ../../go/policy_schemas/traffic_steering_preference/v2/traffic_steering_preference.go
SCHEMA=$(cat ../../api/policy_management/policy_schema_traffic_steering_preference.json)
docker run -v $(pwd)/../../api/policy_management/policy_schema_traffic_steering_preference.json:/api.json onosproject/quicktype:${IMAGE_VER} /api.json -l go -s schema --package tspv2 >> ../../go/policy_schemas/traffic_steering_preference/v2/traffic_steering_preference.go
echo >> ../../go/policy_schemas/traffic_steering_preference/v2/traffic_steering_preference.go
echo var RawSchema = \`$SCHEMA\` >> ../../go/policy_schemas/traffic_steering_preference/v2/traffic_steering_preference.go
echo >> ../../go/policy_schemas/traffic_steering_preference/v2/traffic_steering_preference.go
echo var PolicyTypeID = \"ORAN_TrafficSteeringPreference_2.0.0\" >> ../../go/policy_schemas/traffic_steering_preference/v2/traffic_steering_preference.go

cat ../template/license_header.txt > ../../go/policy_schemas/ue_level_target/v1/ue_level_target.go
SCHEMA=$(cat ../../api/policy_management/policy_schema_ue_level_target.json)
docker run -v $(pwd)/../../api/policy_management/policy_schema_ue_level_target.json:/api.json onosproject/quicktype:${IMAGE_VER} /api.json -l go -s schema --package ueleveltargetv1 >> ../../go/policy_schemas/ue_level_target/v1/ue_level_target.go
echo >> ../../go/policy_schemas/ue_level_target/v1/ue_level_target.go
echo var RawSchema = \`$SCHEMA\` >> ../../go/policy_schemas/ue_level_target/v1/ue_level_target.go
echo >> ../../go/policy_schemas/ue_level_target/v1/ue_level_target.go
echo var PolicyTypeID = \"ORAN_UELevelTarget_1.0.0\" >> ../../go/policy_schemas/ue_level_target/v1/ue_level_target.go
