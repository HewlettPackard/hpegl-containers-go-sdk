#!/bin/bash
# (C) Copyright 2019-2020 Hewlett Packard Enterprise Development LP

TAG=$1
SOURCE_API_SPEC=$2
PACKAGE=mcaasapi

BUILD_DIR=build/latest
LATEST_API_SPEC=$BUILD_DIR/api.yaml

function clean() {
  rm -rf ${PWD}/"$BUILD_DIR"/*
  [[ -d ${PWD}/"$BUILD_DIR" ]] || mkdir ${PWD}/"$BUILD_DIR"
}

clean

set -e

# Run outside the docker container
./build/get-api-spec.sh -a $SOURCE_API_SPEC -o ${PWD}/$LATEST_API_SPEC

# Run inside the docker container
./build/gen-client.sh -a $LATEST_API_SPEC -o $BUILD_DIR -p $PACKAGE -t $TAG

clean
