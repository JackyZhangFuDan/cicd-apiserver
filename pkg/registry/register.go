package registry

import (
	"fmt"

	gRegistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
)

type REST struct {
	*gRegistry.Store
}

func RESTWithErrorHandler(storage rest.StandardStorage, err error) rest.StandardStorage {
	if err != nil {
		err = fmt.Errorf("fail when creating rest storage for a resource due to %v, will die", err)
		panic(err)
	}
	return storage
}
