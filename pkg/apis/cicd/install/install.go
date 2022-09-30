package install

import (
	"cicd-apiserver/pkg/apis/cicd"
	"cicd-apiserver/pkg/apis/cicd/v1alpha"

	"k8s.io/apimachinery/pkg/runtime"
	util "k8s.io/apimachinery/pkg/util/runtime"
)

// 当我们有了scheme实例时，就可以来调这个install来把这个api server支持的object信息注册进来了
func Install(scheme *runtime.Scheme) {
	util.Must(cicd.AddToScheme(scheme))
	util.Must(v1alpha.AddToScheme(scheme))
	util.Must(scheme.SetVersionPriority(v1alpha.SchemeGroupVersion, cicd.SchemeGroupVersion))
}
