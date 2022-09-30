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


// Code generated by C:\Users\I042102\go\bin\client-gen.exe. DO NOT EDIT.

package internalversion

import (
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	cicd "cicd-apiserver/pkg/apis/cicd"
	cicdapiserver\pkg\generated\clientset\internalversion\scheme "cicd-apiserver\pkg\generated\clientset\internalversion\scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


// JenkinsServiceListsGetter has a method to return a JenkinsServiceListInterface.
// A group's client should implement this interface.
type JenkinsServiceListsGetter interface {
	JenkinsServiceLists(namespace string) JenkinsServiceListInterface
}

// JenkinsServiceListInterface has methods to work with JenkinsServiceList resources.
type JenkinsServiceListInterface interface {
Create(ctx context.Context, jenkinsServiceList *cicd.JenkinsServiceList, opts v1.CreateOptions) (*cicd.JenkinsServiceList, error)
Update(ctx context.Context, jenkinsServiceList *cicd.JenkinsServiceList, opts v1.UpdateOptions) (*cicd.JenkinsServiceList, error)
Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
Get(ctx context.Context, name string, opts v1.GetOptions) (*cicd.JenkinsServiceList, error)
List(ctx context.Context, opts v1.ListOptions) (*cicd.JenkinsServiceListList, error)
Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *cicd.JenkinsServiceList, err error)
	JenkinsServiceListExpansion
}

// jenkinsServiceLists implements JenkinsServiceListInterface
type jenkinsServiceLists struct {
	client rest.Interface
	ns     string
}

// newJenkinsServiceLists returns a JenkinsServiceLists
func newJenkinsServiceLists(c *AutobusiClient, namespace string) *jenkinsServiceLists {
	return &jenkinsServiceLists{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the jenkinsServiceList, and returns the corresponding jenkinsServiceList object, and an error if there is any.
func (c *jenkinsServiceLists) Get(ctx context.Context, name string, options v1.GetOptions) (result *cicd.JenkinsServiceList, err error) {
	result = &cicd.JenkinsServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("jenkinsservicelists").
		Name(name).
		VersionedParams(&options, cicdapiserver\pkg\generated\clientset\internalversion\scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of JenkinsServiceLists that match those selectors.
func (c *jenkinsServiceLists) List(ctx context.Context, opts v1.ListOptions) (result *cicd.JenkinsServiceListList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil{
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &cicd.JenkinsServiceListList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("jenkinsservicelists").
		VersionedParams(&opts, cicdapiserver\pkg\generated\clientset\internalversion\scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested jenkinsServiceLists.
func (c *jenkinsServiceLists) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil{
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("jenkinsservicelists").
		VersionedParams(&opts, cicdapiserver\pkg\generated\clientset\internalversion\scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a jenkinsServiceList and creates it.  Returns the server's representation of the jenkinsServiceList, and an error, if there is any.
func (c *jenkinsServiceLists) Create(ctx context.Context, jenkinsServiceList *cicd.JenkinsServiceList, opts v1.CreateOptions) (result *cicd.JenkinsServiceList, err error) {
	result = &cicd.JenkinsServiceList{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("jenkinsservicelists").
		VersionedParams(&opts, cicdapiserver\pkg\generated\clientset\internalversion\scheme.ParameterCodec).
		Body(jenkinsServiceList).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a jenkinsServiceList and updates it. Returns the server's representation of the jenkinsServiceList, and an error, if there is any.
func (c *jenkinsServiceLists) Update(ctx context.Context, jenkinsServiceList *cicd.JenkinsServiceList, opts v1.UpdateOptions) (result *cicd.JenkinsServiceList, err error) {
	result = &cicd.JenkinsServiceList{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("jenkinsservicelists").
		Name(jenkinsServiceList.Name).
		VersionedParams(&opts, cicdapiserver\pkg\generated\clientset\internalversion\scheme.ParameterCodec).
		Body(jenkinsServiceList).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the jenkinsServiceList and deletes it. Returns an error if one occurs.
func (c *jenkinsServiceLists) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("jenkinsservicelists").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *jenkinsServiceLists) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil{
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("jenkinsservicelists").
		VersionedParams(&listOpts, cicdapiserver\pkg\generated\clientset\internalversion\scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched jenkinsServiceList.
func (c *jenkinsServiceLists) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *cicd.JenkinsServiceList, err error) {
	result = &cicd.JenkinsServiceList{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("jenkinsservicelists").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, cicdapiserver\pkg\generated\clientset\internalversion\scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
