.PHONY: deploy

OPERATOR_SDK ?= operator-sdk
KUBECTL ?= oc
CONTAINER ?= podman
IMAGE_BUILDER ?= buildah

CONTAINER_VERSION ?= latest
CONTAINER_REPO ?= quay.io/kscout/operator:${CONTAINER_VERSION}

# Deploy CRD and operator deployment to cluster
deploy:
	${KUBECTL} apply -f deploy/crds/kscout_v1_kscout_crd.yaml

	${OPERATOR_SDK} build --image-builder ${IMAGE_BUILDER} ${CONTAINER_REPO}
	${CONTAINER} push ${CONTAINER_REPO}

	sed 's|REPLACE_IMAGE|${CONTAINER_REPO}|g' deploy/operator.in.yaml > deploy/operator.yaml

	${KUBECTL} apply \
		-f deploy/service_account.yaml \
		-f deploy/role.yaml \
		-f deploy/role_binding.yaml \
		-f deploy/operator.yaml

# Generate Kubernetes scaffholding code for types
generate: $(wildcard pkg/apis/*/*/*_types.go)
	${OPERATOR_SDK} generate k8s
