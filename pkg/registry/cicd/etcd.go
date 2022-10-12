package cicd

import (
	"cicd-apiserver/pkg/apis/cicd"
	"cicd-apiserver/pkg/registry"

	gRegistry "k8s.io/apimachinery/pkg/registry/generic/registry"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
)

func NewREST(scheme *runtime.Scheme, optsGetter generic.RESTOptionsGetter) (*registry.REST, error) {
	strategy := NewStrategy(scheme)

	store := gRegistry.Store{
		NewFunc:                  func() runtime.Object { return &cicd.JenkinsService{} },
		NewListFunc:              func() runtime.Object { return &cicd.JenkinsServiceList{} },
		PredicateFunc:            MatchJenkinsService,
		DefaultQualifiedResource: cicd.Resource("jenkinsservice"),

		CreateStrategy: strategy,
		UpdateStrategy: strategy,
		DeleteStrategy: strategy,

		TableConvertor: rest.NewDefaultTableConvertor(cicd.Resource("jenkinsservice")),
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}
	return &registry.REST{store}, nil
}
