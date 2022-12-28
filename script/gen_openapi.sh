#!/bin/sh

# This script generates the OpenAPI spec for the API server.
# TODO: Not powerful yet. You may need to manually move and merge the generated spec.
cd ..
docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli help
# Generate the OpenAPI spec.
docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i /local/api/openapi-spec/swagger.yaml \
    --git-host github.com \
    --git-user-id sptuan \
    --git-repo-id stargazer/internal/generated/openapi \
    -g go-gin-server \
    --package-name "service" \
    -o /local/internal/generated/openapi