#!/usr/bin/env bash

GENERATOR_IMAGE=openapitools/openapi-generator-cli:v6.0.0
SWAGGER_CLI_IMAGE=stepancons/swagger-cli:v1

GENERATED="/resources"
OPENAPI_SPEC="/docs/openapi/v1"
COMPILED_SPEC_PATH="${OPENAPI_SPEC}/_build/openapi.yaml"

function generate_openapi {
  docker run --rm -v "${PWD}:/local" ${SWAGGER_CLI_IMAGE} bundle "/local${OPENAPI_SPEC}/openapi.yaml" \
      --outfile "/local${COMPILED_SPEC_PATH}" --type yaml
  # validate resulting spec
  docker run --rm -v "${PWD}:/local" ${SWAGGER_CLI_IMAGE} validate "/local${COMPILED_SPEC_PATH}"
  docker run --rm -v "${PWD}:/local" ${GENERATOR_IMAGE} generate \
      -i "/local${COMPILED_SPEC_PATH}" \
      -g go \
      -o "/local${GENERATED}/" \
      --additional-properties packageName="resources" \
      --global-property=models
}

function cleanup_pre {
  find ".${GENERATED}" -type f ! -name "*.keep.go" -exec rm -rf {} \;
  mkdir -p ".${GENERATED}"
}

cleanup_pre
generate_openapi