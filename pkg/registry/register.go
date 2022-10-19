package registry

import (
	"fmt"

	gRegistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
)

// 直接把Store内嵌，相当于“改名”为REST，是由于store这个结构体也是最终响应Restful请求的实体
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
