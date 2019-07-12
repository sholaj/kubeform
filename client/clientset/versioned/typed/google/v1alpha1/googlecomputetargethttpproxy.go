/*
Copyright 2019 The Kubeform Authors.

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

package v1alpha1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// GoogleComputeTargetHttpProxiesGetter has a method to return a GoogleComputeTargetHttpProxyInterface.
// A group's client should implement this interface.
type GoogleComputeTargetHttpProxiesGetter interface {
	GoogleComputeTargetHttpProxies() GoogleComputeTargetHttpProxyInterface
}

// GoogleComputeTargetHttpProxyInterface has methods to work with GoogleComputeTargetHttpProxy resources.
type GoogleComputeTargetHttpProxyInterface interface {
	Create(*v1alpha1.GoogleComputeTargetHttpProxy) (*v1alpha1.GoogleComputeTargetHttpProxy, error)
	Update(*v1alpha1.GoogleComputeTargetHttpProxy) (*v1alpha1.GoogleComputeTargetHttpProxy, error)
	UpdateStatus(*v1alpha1.GoogleComputeTargetHttpProxy) (*v1alpha1.GoogleComputeTargetHttpProxy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleComputeTargetHttpProxy, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleComputeTargetHttpProxyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComputeTargetHttpProxy, err error)
	GoogleComputeTargetHttpProxyExpansion
}

// googleComputeTargetHttpProxies implements GoogleComputeTargetHttpProxyInterface
type googleComputeTargetHttpProxies struct {
	client rest.Interface
}

// newGoogleComputeTargetHttpProxies returns a GoogleComputeTargetHttpProxies
func newGoogleComputeTargetHttpProxies(c *GoogleV1alpha1Client) *googleComputeTargetHttpProxies {
	return &googleComputeTargetHttpProxies{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleComputeTargetHttpProxy, and returns the corresponding googleComputeTargetHttpProxy object, and an error if there is any.
func (c *googleComputeTargetHttpProxies) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleComputeTargetHttpProxy, err error) {
	result = &v1alpha1.GoogleComputeTargetHttpProxy{}
	err = c.client.Get().
		Resource("googlecomputetargethttpproxies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleComputeTargetHttpProxies that match those selectors.
func (c *googleComputeTargetHttpProxies) List(opts v1.ListOptions) (result *v1alpha1.GoogleComputeTargetHttpProxyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleComputeTargetHttpProxyList{}
	err = c.client.Get().
		Resource("googlecomputetargethttpproxies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleComputeTargetHttpProxies.
func (c *googleComputeTargetHttpProxies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googlecomputetargethttpproxies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleComputeTargetHttpProxy and creates it.  Returns the server's representation of the googleComputeTargetHttpProxy, and an error, if there is any.
func (c *googleComputeTargetHttpProxies) Create(googleComputeTargetHttpProxy *v1alpha1.GoogleComputeTargetHttpProxy) (result *v1alpha1.GoogleComputeTargetHttpProxy, err error) {
	result = &v1alpha1.GoogleComputeTargetHttpProxy{}
	err = c.client.Post().
		Resource("googlecomputetargethttpproxies").
		Body(googleComputeTargetHttpProxy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleComputeTargetHttpProxy and updates it. Returns the server's representation of the googleComputeTargetHttpProxy, and an error, if there is any.
func (c *googleComputeTargetHttpProxies) Update(googleComputeTargetHttpProxy *v1alpha1.GoogleComputeTargetHttpProxy) (result *v1alpha1.GoogleComputeTargetHttpProxy, err error) {
	result = &v1alpha1.GoogleComputeTargetHttpProxy{}
	err = c.client.Put().
		Resource("googlecomputetargethttpproxies").
		Name(googleComputeTargetHttpProxy.Name).
		Body(googleComputeTargetHttpProxy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleComputeTargetHttpProxies) UpdateStatus(googleComputeTargetHttpProxy *v1alpha1.GoogleComputeTargetHttpProxy) (result *v1alpha1.GoogleComputeTargetHttpProxy, err error) {
	result = &v1alpha1.GoogleComputeTargetHttpProxy{}
	err = c.client.Put().
		Resource("googlecomputetargethttpproxies").
		Name(googleComputeTargetHttpProxy.Name).
		SubResource("status").
		Body(googleComputeTargetHttpProxy).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleComputeTargetHttpProxy and deletes it. Returns an error if one occurs.
func (c *googleComputeTargetHttpProxies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googlecomputetargethttpproxies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleComputeTargetHttpProxies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googlecomputetargethttpproxies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleComputeTargetHttpProxy.
func (c *googleComputeTargetHttpProxies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComputeTargetHttpProxy, err error) {
	result = &v1alpha1.GoogleComputeTargetHttpProxy{}
	err = c.client.Patch(pt).
		Resource("googlecomputetargethttpproxies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}