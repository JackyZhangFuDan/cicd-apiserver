/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha

import (
	v1alpha "cicd-apiserver/pkg/apis/cicd/v1alpha"
	scheme "cicd-apiserver/pkg/generated/clientset/versioned/scheme"
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// JenkinsServicesGetter has a method to return a JenkinsServiceInterface.
// A group's client should implement this interface.
type JenkinsServicesGetter interface {
	JenkinsServices(namespace string) JenkinsServiceInterface
}

// JenkinsServiceInterface has methods to work with JenkinsService resources.
type JenkinsServiceInterface interface {
	Create(ctx context.Context, jenkinsService *v1alpha.JenkinsService, opts v1.CreateOptions) (*v1alpha.JenkinsService, error)
	Update(ctx context.Context, jenkinsService *v1alpha.JenkinsService, opts v1.UpdateOptions) (*v1alpha.JenkinsService, error)
	UpdateStatus(ctx context.Context, jenkinsService *v1alpha.JenkinsService, opts v1.UpdateOptions) (*v1alpha.JenkinsService, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha.JenkinsService, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha.JenkinsServiceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha.JenkinsService, err error)
	JenkinsServiceExpansion
}

// jenkinsServices implements JenkinsServiceInterface
type jenkinsServices struct {
	client rest.Interface
	ns     string
}

// newJenkinsServices returns a JenkinsServices
func newJenkinsServices(c *AutobusiV1alphaClient, namespace string) *jenkinsServices {
	return &jenkinsServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the jenkinsService, and returns the corresponding jenkinsService object, and an error if there is any.
func (c *jenkinsServices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha.JenkinsService, err error) {
	result = &v1alpha.JenkinsService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("jenkinsservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of JenkinsServices that match those selectors.
func (c *jenkinsServices) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha.JenkinsServiceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha.JenkinsServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("jenkinsservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested jenkinsServices.
func (c *jenkinsServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("jenkinsservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a jenkinsService and creates it.  Returns the server's representation of the jenkinsService, and an error, if there is any.
func (c *jenkinsServices) Create(ctx context.Context, jenkinsService *v1alpha.JenkinsService, opts v1.CreateOptions) (result *v1alpha.JenkinsService, err error) {
	result = &v1alpha.JenkinsService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("jenkinsservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(jenkinsService).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a jenkinsService and updates it. Returns the server's representation of the jenkinsService, and an error, if there is any.
func (c *jenkinsServices) Update(ctx context.Context, jenkinsService *v1alpha.JenkinsService, opts v1.UpdateOptions) (result *v1alpha.JenkinsService, err error) {
	result = &v1alpha.JenkinsService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("jenkinsservices").
		Name(jenkinsService.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(jenkinsService).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *jenkinsServices) UpdateStatus(ctx context.Context, jenkinsService *v1alpha.JenkinsService, opts v1.UpdateOptions) (result *v1alpha.JenkinsService, err error) {
	result = &v1alpha.JenkinsService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("jenkinsservices").
		Name(jenkinsService.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(jenkinsService).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the jenkinsService and deletes it. Returns an error if one occurs.
func (c *jenkinsServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("jenkinsservices").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *jenkinsServices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("jenkinsservices").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched jenkinsService.
func (c *jenkinsServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha.JenkinsService, err error) {
	result = &v1alpha.JenkinsService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("jenkinsservices").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
