module github.com/kscout/operator

go 1.12

// Pinned to kubernetes-1.13.1
replace (
	k8s.io/api => k8s.io/api v0.0.0-20181213150558-05914d821849
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20181213153335-0fe22c71c476
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20181127025237-2b1284ed4c93
	k8s.io/client-go => k8s.io/client-go v0.0.0-20181213151034-8d9ed539ba31
)

require (
	github.com/NYTimes/gziphandler v1.0.1 // indirect
	github.com/go-openapi/spec v0.19.2
	github.com/gobuffalo/flect v0.1.5 // indirect
	github.com/google/go-cmp v0.3.1 // indirect
	github.com/google/go-containerregistry v0.0.0-20190729175742-ef12d49c8daf // indirect
	github.com/operator-framework/operator-sdk v0.9.0
	github.com/spf13/pflag v1.0.3
	k8s.io/api v0.0.0-20190612125737-db0771252981
	k8s.io/apimachinery v0.0.0-20190612125636-6a5db36e93ad
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/code-generator v0.0.0-20190803082810-c4ef572adb98
	k8s.io/gengo v0.0.0-20190327210449-e17681d19d3a
	k8s.io/kube-openapi v0.0.0-20190722073852-5e22f3d471e6
	knative.dev/pkg v0.0.0-20190807140856-4707aad818fe // indirect
	knative.dev/serving v0.8.0
	sigs.k8s.io/controller-runtime v0.1.12
	sigs.k8s.io/controller-tools v0.1.12
)
