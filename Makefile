OPERATOR_SDK ?= operator-sdk

# Generate Kubernetes scaffholding code for types
generate: $(wildcard pkg/apis/*/*/*_types.go)
	${OPERATOR_SDK} generate k8s
