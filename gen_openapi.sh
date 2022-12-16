#!/bin/sh

# This script generates the OpenAPI spec for the API server.
# It is intended to be run from the root of the repository.

docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli help
# Generate the OpenAPI spec.
docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i /local/api/openapi-spec/swagger.yaml \
    --git-host github.com \
    --git-user-id sptuan \
    --git-repo-id stargazer/modules/generated/openapi \
    -g go-gin-server \
    --package-name "service" \
    -o /local/modules/generated/openapi