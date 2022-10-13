package apiserver

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	genericapiserver "k8s.io/apiserver/pkg/server"
)

var (
	Scheme = runtime.NewScheme()
	Codecs = serializer.NewCodecFactory(Scheme)
)

type Config struct {
	GenericConfig *genericapiserver.RecommendedConfig
	// ExtraConfig   ExtraConfig // 如果有自己需要的config的话，可以扩展field
}

type completedConfig struct{
	GenericConfig *genericapiserver.completedConfig
}

type CompletedConfig struct{
	*completedConfig
}

type CicdServer struct{
	GenericAPIServer *genericapiserver.GenericAPIServer
}

func (cfg *)