package plugin

import (
	"cicd-apiserver/pkg/apis/cicd"
	"context"
	"fmt"
	"io"

	informers "cicd-apiserver/pkg/generated/informers/internalversion"
	listers "cicd-apiserver/pkg/generated/listers/cicd/internalversion"

	"k8s.io/apiserver/pkg/admission"

	"k8s.io/apimachinery/pkg/labels"
)

// plugin必须实现admission.Interface接口，而内嵌的admission.Handler结构体就实现了
type JenkinsServicePlugin struct {
	*admission.Handler
	jsLister listers.JenkinsServiceLister
}

// 把这个plugin注册进api server体系中的方法，会在server启动的代码中调用
// plugin 参数是server的plugins集合，需要把我们的放进去
func Register(plugin *admission.Plugins) {
	plugin.Register("JenkinsService", func(config io.Reader) (admission.Interface, error) {
		return New()
	})
}

// 创建plugin结构体实例的方法
func New() (*JenkinsServicePlugin, error) {
	return &JenkinsServicePlugin{
		Handler: admission.NewHandler(admission.Create),
	}, nil
}

// 有了validate方法就实现了admission.ValidationInterface，从而在validating阶段被调用
func (jsp *JenkinsServicePlugin) Validate(ctx context.Context, a admission.Attributes, _ admission.ObjectInterfaces) error {
	if a.GetKind().GroupKind() != cicd.Kind("JenkinsService") { //所有object的valid都会进来，所以我们要验一下是不是该关心的
		return nil
	}

	if !jsp.WaitForReady() { // 例如informer还没有把远程信息sync到本地
		return admission.NewForbidden(a, fmt.Errorf("the plugin isn't ready for handling request"))
	}

	// 下面就可以进行我们期望的校验了
	// 区别于registry部分strategy中的valid strategy，此处的校验更多是多实体之间的关联正确性，而不是单个jenkins service内容的正确
	// 例如，我们规定整个系统中只能存在10 个JenkinsService对象，多了不行，就可以在这里做检查
	existedJenkinsServices, err := jsp.jsLister.List(labels.Everything())
	if err != nil {
		return admission.NewForbidden(a, fmt.Errorf("the plugin encounter internal error during retrieve jenkins service objects from api server"))
	}
	if len(existedJenkinsServices) >= 10 {
		return admission.NewForbidden(a, fmt.Errorf("too many service instances exist, %d exist but max is 10", len(existedJenkinsServices)))
	}

	return nil

}

// 有了这个方法，plugin就实现了WantsCicdInformerFactory接口，可以获取到cicd informer了
func (jsp *JenkinsServicePlugin) SetInformerFactory(f informers.SharedInformerFactory) {
	jsp.jsLister = f.Autobusi().InternalVersion().JenkinsServices().Lister()
	jsp.SetReadyFunc(f.Autobusi().InternalVersion().JenkinsServices().Informer().HasSynced)
}
