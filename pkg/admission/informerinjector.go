package admission

import (
	informers "cicd-apiserver/pkg/generated/informers/internalversion"

	"k8s.io/apiserver/pkg/admission"
)

// 需要admission plugin去实现这个接口，从而保证可以接收informerfactory；
type WantsCicdInformerFactory interface {
	SetInformerFactory(informers.SharedInformerFactory)
}

type cicdInformerPluginInitializer struct {
	informers informers.SharedInformerFactory
}

func (i cicdInformerPluginInitializer) Initialize(plugin admission.Interface) {
	if wants, ok := plugin.(WantsCicdInformerFactory); ok { //如果目标plugin通过实现接口，声明需要cicd informer，那么我们就给它
		wants.SetInformerFactory(i.informers)
	}
}

// server启动时在config阶段被调用，从而把informer交给plugin
func NewCicdInformerPluginInitializer(informers informers.SharedInformerFactory) cicdInformerPluginInitializer {
	return cicdInformerPluginInitializer{
		informers: informers,
	}
}
