#!/bin/bash
# (C) Copyright 2019-2020 Hewlett Packard Enterprise Development LP

TAG=v0.3.1
SOURCE_API_SPEC="https://raw.githubusercontent.com/hpe-hcss/caas/$TAG/api/openapi-spec/caas-public-api.yaml"
PACKAGE=mcaasapi

BUILD_DIR=./build/latest
LATEST_API_SPEC=$BUILD_DIR/api.yaml

function clean() {
  rm -rf "$BUILD_DIR"/*
  [[ -d "$BUILD_DIR" ]] || mkdir "$BUILD_DIR"
}

clean

set -e

./build/get-api-spec.sh -a $SOURCE_API_SPEC -o $LATEST_API_SPEC

./build/gen-client.sh -a $LATEST_API_SPEC -o $BUILD_DIR -p $PACKAGE -t $TAG

clean
