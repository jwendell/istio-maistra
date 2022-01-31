#!/bin/bash

set -e

: "${ISTIO_OUT_LINUX?Missing variable}"

WD=$(dirname "$0")
WD=$(cd "$WD"; pwd)

DIR=$(mktemp -d)
trap 'rm -rf "${DIR}"' EXIT

cp "${WD}/Dockerfile.pilot" "${DIR}/Dockerfile"
cp "${ISTIO_OUT_LINUX}/pilot-discovery" "${DIR}"

CONTAINER_CLI="${CONTAINER_CLI:-docker}"

# HUB="${HUB:-quay.io/maistra-dev}"
# TAG="${TAG:-${MAISTRA_VERSION}-daily-$(date '+%Y%m%d')}"

HUB="quay.io/maistra-dev"
TAG="${MAISTRA_VERSION}-daily-$(date '+%Y%m%d')"

cd "${DIR}"
${CONTAINER_CLI} build --build-arg VERSION="${MAISTRA_VERSION}" --build-arg ISTIO_VERSION="${ISTIO_VERSION}" -t "${HUB}/pilot-ubi8:${TAG}" .
