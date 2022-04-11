#!/bin/bash
# (C) Copyright 2019-2020 Hewlett Packard Enterprise Development LP
# Pretty hacky.  Skipping all the best practices like clean

API_SPEC=
TAG=
PACKAGE=
OUTPUT_DIR=
REPO=
PKG_DIR=
DOC_DIR=./docs

while getopts ":a:o:p:t:" opt; do
  case ${opt} in
    a ) # process option a
      API_SPEC=$OPTARG
      ;;
    o ) # process option o
      OUTPUT_DIR="$OPTARG"/generated
      ;;
    p ) # process option p
      PACKAGE="$OPTARG"
      PKG_DIR=./pkg/$PACKAGE
      ;;
    t ) # process option t
      TAG="$OPTARG"
      ;;
    \? ) echo "Usage: cmd -a -o"
      ;;
  esac
done

function validate_input {
  if [[ -z "$API_SPEC" ]]; then
    echo "you must specify provide -a with the locations of the api spec to retrieve"
    exit 1
  else
    echo "Using api spec $API_SPEC"
  fi

  if [[ -z "$OUTPUT_DIR" ]]; then
    echo "you must specify provide -o with the output location"
    exit 1
  else
    echo "Using api spec $API_SPEC"
  fi

  if [[ -z "$PACKAGE" ]]; then
    echo "you must specify provide -p with the target package"
    exit 1
  else
    echo "Using api spec $API_SPEC"
  fi

  if [[ -z "$TAG" ]]; then
    echo "you must specify provide -p with the target package"
    exit 1
  else
    echo "Using api spec $API_SPEC"
  fi
}

function generate {
  docker pull swaggerapi/swagger-generator-v3

  docker run --rm -v "${PWD}"/:/local swaggerapi/swagger-codegen-cli-v3 generate -i /local/"$API_SPEC" -l go -o /local/"$OUTPUT_DIR"  --additional-properties=packageName="$PACKAGE",packageVersion="$TAG",hidegenerationTimestamp=false

  rm -rf "${PKG_DIR}"
  mkdir "${PKG_DIR}"
  cp "$OUTPUT_DIR"/*.go $PKG_DIR

  rm -rf ${DOC_DIR}/*
  cp "$OUTPUT_DIR"/docs/* $DOC_DIR

  # Add custom usage example - Note could try to change to customer mustache template
  README="$OUTPUT_DIR"/README.md
  REPO=$(git config --get remote.origin.url)
  REPO="${REPO/git@github.com:/github.com/}"
  REPO="${REPO/.git/}"
  PACKAGE_PATH=${REPO}/pkg/${PACKAGE}
  EXAMPLE_PATH=$(dirname $(realpath -s $0))/example.md
  sed -i '' 's|Put the package under your project folder and add the following in import:||g' "$README"
  sed -i '' "/## Installation/r $EXAMPLE_PATH" "$README"
  sed -i '' "s|{{PACKAGE_PATH}}|${PACKAGE_PATH}|g" "$README"
  sed -i '' "s|{{PACKAGE}}|${PACKAGE}|g" "$README"

  tmp="$README".tmp
  awk '/## Documentation For Authorization/ {exit} {print}' "$README" > "$tmp" && mv "$tmp" $README

  cp "$README" ./
}

validate_input
generate

