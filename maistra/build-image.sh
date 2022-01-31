#!/bin/bash

set -e

: "${ISTIO_OUT_LINUX?Missing variable}"

WD=$(dirname "$0")
WD=$(cd "$WD"; pwd)

DIR=$(mktemp -d)
trap 'rm -rf "${DIR}"' EXIT

if [ "$1" == "pilot" ]; then
  cp "${WD}/Dockerfile.pilot" "${DIR}/Dockerfile"
  cp "${ISTIO_OUT_LINUX}/pilot-discovery" "${DIR}"
  cp "${ISTIO_ENVOY_BOOTSTRAP_CONFIG_DIR}/"*.json "${DIR}"
  NAME=pilot-ubi8

elif [ "$1" == "proxy" ]; then
  cp "${WD}/Dockerfile.proxyv2" "${DIR}/Dockerfile"
  cp "${ISTIO_OUT_LINUX}/pilot-agent" "${DIR}"
  cp "${ISTIO_OUT_LINUX}/envoy" "${DIR}"
  cp "${ISTIO_ENVOY_BOOTSTRAP_CONFIG_DIR}/"*.json "${DIR}"
  cp "${ISTIO_ENVOY_LINUX_RELEASE_DIR}/"*.wasm "${DIR}"
  NAME=proxyv2-ubi8

elif [ "$1" == "cni" ]; then
  cp "${WD}/Dockerfile.cni" "${DIR}/Dockerfile"
  cp "${ISTIO_OUT_LINUX}/istio-cni" "${DIR}"
  cp "${ISTIO_OUT_LINUX}/istio-cni-taint" "${DIR}"
  cp "${ISTIO_OUT_LINUX}/install-cni" "${DIR}"
  NAME=istio-cni-ubi8

else
  echo "Invalid image"
  exit 1
fi


CONTAINER_CLI="${CONTAINER_CLI:-docker}"

HUB="${HUB:-quay.io/maistra-dev}"
TAG="${TAG:-${MAISTRA_VERSION}-daily-$(date '+%Y%m%d')}"

# HUB="quay.io/maistra-dev"
# TAG="${MAISTRA_VERSION}-daily-$(date '+%Y%m%d')"

echo "Building ${HUB}/${NAME}:${TAG}"

cd "${DIR}"
${CONTAINER_CLI} build --build-arg VERSION="${MAISTRA_VERSION}" --build-arg ISTIO_VERSION="${ISTIO_VERSION}" -t "${HUB}/${NAME}:${TAG}" .

echo "Done building ${HUB}/${NAME}:${TAG}"
echo
