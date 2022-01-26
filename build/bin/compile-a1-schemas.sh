#!/bin/bash

export IMAGE_VER=latest

pushd ../docker/quicktype
docker build . -t onosproject/quicktype:${IMAGE_VER}
popd

cat ../template/license_header.txt > ../../go/policy_status/v2/policy_status.go
docker run -v $(pwd)/../../api/policy_management/policy_status.json:/api.json onosproject/quicktype:${IMAGE_VER} /api.json -l go -s schema --package policystatusv2>> ../../go/policy_status/v2/policy_status.go

cat ../template/license_header.txt > ../../go/policy_scope_id/v2/policy_scope_identifier.go
docker run -v $(pwd)/../../api/policy_management/policy_scope_identifier.json:/api.json onosproject/quicktype:${IMAGE_VER} /api.json -l go -s schema --package policyscopeidv2 >> ../../go/policy_scope_id/v2/policy_scope_identifier.go

cat ../template/license_header.txt > ../../go/policy_schemas/qoe_and_tsp/v2/qoe_and_tsp.go
docker run -v $(pwd)/../../api/policy_management/policy_schema_qoe_and_tsp.json:/api.json onosproject/quicktype:${IMAGE_VER} /api.json -l go -s schema --package qoeandtspv2 >> ../../go/policy_schemas/qoe_and_tsp/v2/qoe_and_tsp.go

cat ../template/license_header.txt > ../../go/policy_schemas/qoe_target/v2/qoe_target.go
docker run -v $(pwd)/../../api/policy_management/policy_schema_qoe_target.json:/api.json onosproject/quicktype:${IMAGE_VER} /api.json -l go -s schema --package qoetargetv2 >> ../../go/policy_schemas/qoe_target/v2/qoe_target.go

cat ../template/license_header.txt > ../../go/policy_schemas/qos_and_tsp/v2/qos_and_tsp.go
docker run -v $(pwd)/../../api/policy_management/policy_schema_qos_and_tsp.json:/api.json onosproject/quicktype:${IMAGE_VER} /api.json -l go -s schema --package qosandtspv2 >> ../../go/policy_schemas/qos_and_tsp/v2/qos_and_tsp.go

cat ../template/license_header.txt > ../../go/policy_schemas/qos_target/v2/qos_target.go
docker run -v $(pwd)/../../api/policy_management/policy_schema_qos_target.json:/api.json onosproject/quicktype:${IMAGE_VER} /api.json -l go -s schema --package qostargetv2 >> ../../go/policy_schemas/qos_target/v2/qos_target.go

cat ../template/license_header.txt > ../../go/policy_schemas/slice_sla_target/v1/slice_sla_target.go
docker run -v $(pwd)/../../api/policy_management/policy_schema_slice_sla_target.json:/api.json onosproject/quicktype:${IMAGE_VER} /api.json -l go -s schema --package sliceslatargetv1 >> ../../go/policy_schemas/slice_sla_target/v1/slice_sla_target.go

cat ../template/license_header.txt > ../../go/policy_schemas/traffic_steering_preference/v2/traffic_steering_preference.go
docker run -v $(pwd)/../../api/policy_management/policy_schema_traffic_steering_preference.json:/api.json onosproject/quicktype:${IMAGE_VER} /api.json -l go -s schema --package tspv2 >> ../../go/policy_schemas/traffic_steering_preference/v2/traffic_steering_preference.go

cat ../template/license_header.txt > ../../go/policy_schemas/ue_level_target/v1/ue_level_target.go
docker run -v $(pwd)/../../api/policy_management/policy_schema_ue_level_target.json:/api.json onosproject/quicktype:${IMAGE_VER} /api.json -l go -s schema --package ueleveltargetv1 >> ../../go/policy_schemas/ue_level_target/v1/ue_level_target.go