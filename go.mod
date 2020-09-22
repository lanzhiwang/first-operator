module github.com/imooc-com/imooc-operator

go 1.13

require (
	github.com/operator-framework/operator-sdk v0.17.0
	sigs.k8s.io/controller-runtime v0.5.2
)

replace (
	k8s.io/client-go => k8s.io/client-go v0.17.4 // Required by prometheus-operator
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.3.2+incompatible // Required by OLM
)
