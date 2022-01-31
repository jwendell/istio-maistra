export MAISTRA_VERSION ?= 2.2.0

.PHONY: vendor
vendor:
	@echo "updating vendor"
	@go mod vendor
	@echo "done updating vendor"

gen: vendor

# STANDARD_BINARIES += ./mec/cmd/mec

.PHONY: maistra-images maistra-image-pilot maistra-image-proxyv2
maistra-images: maistra-image-pilot maistra-image-proxy maistra-image-cni

maistra-image-pilot: ISTIO_VERSION=${VERSION} 
maistra-image-pilot: VERSION=${MAISTRA_VERSION}
maistra-image-pilot: build
	@maistra/build-image.sh pilot

maistra-image-proxy: ISTIO_VERSION=${VERSION} VERSION=${MAISTRA_VERSION}
maistra-image-proxy: build
	@maistra/build-image.sh proxy

maistra-image-cni: ISTIO_VERSION=${VERSION} VERSION=${MAISTRA_VERSION}
maistra-image-cni: build
	@maistra/build-image.sh cni
