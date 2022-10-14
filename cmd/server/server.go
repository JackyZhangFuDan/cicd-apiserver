package server

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apiserver/pkg/admission"
	genericserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"

	myadmission "cicd-apiserver/pkg/admission"
	"cicd-apiserver/pkg/admission/plugin"
	"cicd-apiserver/pkg/apis/cicd/v1alpha"
	"cicd-apiserver/pkg/apiserver"
	clientset "cicd-apiserver/pkg/generated/clientset/internalversion"
	informers "cicd-apiserver/pkg/generated/informers/internalversion"
)

const etcdPathPrefix = "/registry/cicd-apiserver.autobusi.com"

// 以下环节制作option，option是server启动参数，由它进一步制作config，再然后由config制造server
type ServerOptions struct {
	RecommendedOptions    *genericoptions.RecommendedOptions
	SharedInformerFactory informers.SharedInformerFactory
}

func NewServerOptions() *ServerOptions {
	o := &ServerOptions{
		RecommendedOptions: genericoptions.NewRecommendedOptions(
			etcdPathPrefix,
			apiserver.Codecs.LegacyCodec(v1alpha.SchemeGroupVersion),
		),
	}
	return o
}
func (o *ServerOptions) Validate() error {
	errors := []error{}
	errors = append(errors, o.RecommendedOptions.Validate()...)
	//把errors数组合并成单独error
	return utilerrors.NewAggregate(errors)
}
func (o *ServerOptions) Complete() error {
	//把我们的admission plugin加进去
	plugin.Register(o.RecommendedOptions.Admission.Plugins)
	//plugin 的相互顺序很重要，最好不要破坏已有顺序，直接加在尾部
	o.RecommendedOptions.Admission.RecommendedPluginOrder =
		append(o.RecommendedOptions.Admission.RecommendedPluginOrder, "JenkinsService")
	return nil
}
func (o *ServerOptions) Config() (*apiserver.Config, error) {
	//申请系统分派证书
	if err := o.RecommendedOptions.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost", nil, []net.IP{net.ParseIP("127.0.0.1")}); err != nil {
		return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
	}
	//把我们admission用的informer做出来，放入admission初始化器。informer的初始化在post start hook中进行
	o.RecommendedOptions.ExtraAdmissionInitializers = func(cfg *genericserver.RecommendedConfig) ([]admission.PluginInitializer, error) {
		client, err := clientset.NewForConfig(cfg.LoopbackClientConfig)
		if err != nil {
			return nil, err
		}
		informerFactory := informers.NewSharedInformerFactory(client, cfg.LoopbackClientConfig.Timeout)
		o.SharedInformerFactory = informerFactory
		return []admission.PluginInitializer{myadmission.NewCicdInformerPluginInitializer(informerFactory)}, nil
	}
	//做config，用于返回
	standardConfig := genericserver.NewRecommendedConfig(apiserver.Codecs)
	if err := o.RecommendedOptions.ApplyTo(standardConfig); err != nil {
		return nil, err
	}
	myConfig := &apiserver.Config{
		GenericConfig: standardConfig,
	}
	return myConfig, nil
}

// 把server做出来并跑起来
func (o *ServerOptions) Run(stopCh <-chan struct{}) error {
	c, err := o.Config()
	if err != nil {
		return err
	}

	s, err := c.Complete().NewServer()
	if err != nil {
		return err
	}

	s.GenericAPIServer.AddPostStartHook("start-cicd-apiserver-informers", func(context genericserver.PostStartHookContext) error {
		c.GenericConfig.SharedInformerFactory.Start(context.StopCh)
		o.SharedInformerFactory.Start(context.StopCh)
		return nil
	})
	return s.GenericAPIServer.PrepareRun().Run(stopCh)
}

//以下环节制作cobra command，它可以启动server
func NewCommandStartServer(defaultOptions *ServerOptions, stopCh <-chan struct{}) *cobra.Command {
	options := *defaultOptions
	cmd := &cobra.Command{
		Short: "my cicd api server",
		Long:  "my api server for demostration purpose",
		RunE: func(c *cobra.Command, args []string) error {
			if err := options.Complete(); err != nil {
				return err
			}
			if err := options.Validate(); err != nil {
				return err
			}
			if err := options.Run(stopCh); err != nil {
				return err
			}
			return nil
		},
	}
	flags := cmd.Flags()
	options.RecommendedOptions.AddFlags(flags)
	return cmd
}
