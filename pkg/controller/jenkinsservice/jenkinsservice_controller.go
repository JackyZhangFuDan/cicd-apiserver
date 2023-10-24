package jenkinsservice

import (
	"cicd-apiserver/pkg/apis/cicd/v1alpha"

	informer "cicd-apiserver/pkg/generated/informers/externalversions/cicd/v1alpha"
	lister "cicd-apiserver/pkg/generated/listers/cicd/v1alpha"
	"context"
	"fmt"
	"time"

	apps "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/apimachinery/pkg/util/wait"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
)

type Controller struct {
	jenkinsServiceSynced cache.InformerSynced
	coreAPIClient        clientset.Interface
	jenkinsServiceLister lister.JenkinsServiceLister
	queue                workqueue.RateLimitingInterface

	syncHandler func(ctx context.Context, key string) error
}

func NewJenkinsServiceController(jsInformer informer.JenkinsServiceInformer, coreAPIClient clientset.Interface) *Controller {
	c := &Controller{
		coreAPIClient:        coreAPIClient,
		jenkinsServiceLister: jsInformer.Lister(),
		jenkinsServiceSynced: jsInformer.Informer().HasSynced,

		queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "jenkinsservice"),
	}
	c.syncHandler = c.sync

	jsInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		// 只跟踪JS的创建
		AddFunc: func(obj interface{}) {
			klog.Info("New Jenkins Service Object is found")
			cast := obj.(*v1alpha.JenkinsService)
			key, err := cache.MetaNamespaceKeyFunc(cast)
			if err != nil {
				klog.ErrorS(err, "Failed when extracting key for Jenkins Service Object")
				return
			}
			c.queue.Add(key)
		},
	})
	return c
}

// 启动控制器，直到 stopCh关闭
func (c *Controller) Run(ctx context.Context, workers int) {
	defer utilruntime.HandleCrash()
	defer c.queue.ShutDown()

	klog.Info("Starting Jenkins Service Controller")
	defer klog.Info("Shutting down Jenkins Service Controller")

	// 等待本地JS的cache同步完毕
	klog.Info("Waiting for caches to sync for Jenkins Service controller")
	if !cache.WaitForCacheSync(ctx.Done(), c.jenkinsServiceSynced) {
		utilruntime.HandleError(fmt.Errorf("Unable to sync caches for Jenkins Service controller"))
		return
	}
	klog.Infof("Caches are synced for Jenkins Service controller")

	// 启动指定数量的协程，每个都跑控制循环
	// wait.Until的作用是如果runWorker处理失败退出，那么再次启动它
	for i := 0; i < workers; i++ {
		go wait.UntilWithContext(ctx, c.runWorker, time.Second)
	}

	<-ctx.Done()
}

func (c *Controller) runWorker(ctx context.Context) {
	for c.processNextWorkItem(ctx) {
	}
}

func (c *Controller) processNextWorkItem(ctx context.Context) bool {
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	//告诉queue，这个key的事情你已经处理完毕了
	defer c.queue.Done(key)

	err := c.syncHandler(ctx, key.(string))
	if err == nil {
		//处理完毕没出错，告诉queue清理这个key相关的追踪记录，例如失败计数等
		c.queue.Forget(key)
		klog.Infof("Finish processing key %s", key)
		return true
	}
	//上报错误，这个调用允许插拔错误处理逻辑，例如集群监控
	utilruntime.HandleError(fmt.Errorf("%v failed with: %v", key, err))
	//把失败的key放入队列，但在最后，从而赢得一段冷静事件，马上处理的话可能还会失败
	c.queue.AddRateLimited(key)
	return true
}

// 一个控制循环的处理逻辑
func (c *Controller) sync(ctx context.Context, key string) (err error) {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		klog.ErrorS(err, "Failed to split meta namespace cache key", "cacheKey", name)
	}

	js, err := c.jenkinsServiceLister.JenkinsServices(namespace).Get(name)
	if errors.IsNotFound(err) {
		klog.Infof("Jenkins Service %s has been deleted", key)
		return nil
	}
	if err != nil {
		return err
	}

	var replicas int32
	replicas = int32(js.Spec.InstanceAmount)
	selector := map[string]string{}
	selector["type"] = "jenkinsservice"
	selector["jsname"] = name

	// 创建一个Deployment
	d := apps.Deployment{
		TypeMeta: metav1.TypeMeta{APIVersion: "apps/v1", Kind: "Deployment"},
		ObjectMeta: metav1.ObjectMeta{
			UID:         uuid.NewUUID(),
			Name:        name,
			Namespace:   namespace,
			Annotations: make(map[string]string),
		},
		Spec: apps.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{MatchLabels: selector},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: selector,
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:            "jenkinsforjs",
							Image:           "nginx:latest",
							ImagePullPolicy: "IfNotPresent",
						},
					},
				},
			},
		},
	}
	_, err = c.coreAPIClient.AppsV1().Deployments(namespace).Get(ctx, d.Name, metav1.GetOptions{})
	if err == nil {
		klog.Infof("Deployment for Jenkins Service %s already exists", key)
		return nil
	}
	_, err = c.coreAPIClient.AppsV1().Deployments(namespace).Create(ctx, &d, metav1.CreateOptions{})
	if err != nil {
		klog.ErrorS(err, "Failed when creating deployment for Jenkins Service")
		return err
	}

	klog.Infof("Successfully create deploymnet for Jenkins Service %s", key)
	return nil
}
