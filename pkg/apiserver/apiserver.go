package apiserver

import (
	"cicd-apiserver/pkg/apis/cicd"
	cicdregistry "cicd-apiserver/pkg/registry"
	jsstorage "cicd-apiserver/pkg/registry/cicd/jenkinsservice"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"

	"cicd-apiserver/pkg/apis/cicd/install"
)

var (
	Scheme = runtime.NewScheme()
	Codecs = serializer.NewCodecFactory(Scheme)
)

// 如下方法需要更新至相应phase，开始漏掉了
func init() {
	install.Install(Scheme)
	metav1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
	unversioned := schema.GroupVersion{Group: "", Version: "v1"}
	Scheme.AddUnversionedTypes(
		unversioned,
		&metav1.Status{},
		&metav1.APIVersions{},
		&metav1.APIGroupList{},
		&metav1.APIGroup{},
		&metav1.APIResourceList{},
	)
}

//如下环节制作Server的Config
type Config struct {
	GenericConfig *genericapiserver.RecommendedConfig
	// ExtraConfig   ExtraConfig // 如果有自己需要的config的话，可以扩展field
}

type completedConfig struct {
	GenericConfig genericapiserver.CompletedConfig
}

//完善后的config
type CompletedConfig struct {
	*completedConfig
}

type CicdServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}

//完善初始的config
func (cfg *Config) Complete() CompletedConfig {
	cconfig := completedConfig{
		cfg.GenericConfig.Complete(),
	}
	cconfig.GenericConfig.Version = &version.Info{
		Major: "1",
		Minor: "0",
	}
	return CompletedConfig{&cconfig}
}

//有了这个方法，完善后的config就可以制作server的instance了
func (ccfg completedConfig) NewServer() (*CicdServer, error) {
	genericServer, err := ccfg.GenericConfig.New("cicd-apiserver", genericapiserver.NewEmptyDelegate())
	if err != nil {
		return nil, err
	}

	server := &CicdServer{
		GenericAPIServer: genericServer,
	}

	//重点是把我们各个版本的api object都注入到server中去，开始
	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(
		cicd.GroupName,
		Scheme,
		metav1.ParameterCodec,
		Codecs,
	)
	v1alphastorage := map[string]rest.Storage{}
	v1alphastorage["jenkinsservices"] = cicdregistry.RESTWithErrorHandler(jsstorage.NewREST(Scheme, ccfg.GenericConfig.RESTOptionsGetter))
	apiGroupInfo.VersionedResourcesStorageMap["v1alpha"] = v1alphastorage

	if err := server.GenericAPIServer.InstallAPIGroup(&apiGroupInfo); err != nil {
		return nil, err
	}

	return server, nil
}
