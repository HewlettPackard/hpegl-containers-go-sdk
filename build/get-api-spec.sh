#!/bin/bash
# (C) Copyright 2019-2020 Hewlett Packard Enterprise Development LP

OUTPUT=
API_SPEC=

while getopts ":a:o:" opt; do
  case ${opt} in
    a ) # process option a
      API_SPEC=$OPTARG
      ;;
    o ) # process option o
      OUTPUT=$OPTARG
      ;;
    \? ) echo "Usage: cmd -a -o"
      ;;
  esac
done

function validate_input {
  if [[ -z "$GITHUB_TOKEN" ]]; then
    echo "Set your GITHUB_TOKEN env var"
    exit 1
  fi

  if [[ -z "$API_SPEC" ]]; then
    echo "you must specify provide -a with the locations of the api spec to retrieve"
    exit 1
  else
    echo "Using api spec $API_SPEC"
  fi

  if [[ -z "$OUTPUT" ]]; then
    echo "you must specify provide -o with the output location"
    exit 1
  else
    echo "Using api spec $API_SPEC"
  fi
}

function get_spec {
  #Delete to ensure old isn't proccessed accidentally if failure here
  if [ -e "$OUTPUT" ]; then
      rm "$OUTPUT"
  fi

  curl --header "Authorization: token $GITHUB_TOKEN" \
       --output "$OUTPUT" \
       "$API_SPEC"
}

validate_input
get_spec

