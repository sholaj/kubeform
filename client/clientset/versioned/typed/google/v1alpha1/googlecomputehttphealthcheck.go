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

// GoogleComputeHttpHealthChecksGetter has a method to return a GoogleComputeHttpHealthCheckInterface.
// A group's client should implement this interface.
type GoogleComputeHttpHealthChecksGetter interface {
	GoogleComputeHttpHealthChecks() GoogleComputeHttpHealthCheckInterface
}

// GoogleComputeHttpHealthCheckInterface has methods to work with GoogleComputeHttpHealthCheck resources.
type GoogleComputeHttpHealthCheckInterface interface {
	Create(*v1alpha1.GoogleComputeHttpHealthCheck) (*v1alpha1.GoogleComputeHttpHealthCheck, error)
	Update(*v1alpha1.GoogleComputeHttpHealthCheck) (*v1alpha1.GoogleComputeHttpHealthCheck, error)
	UpdateStatus(*v1alpha1.GoogleComputeHttpHealthCheck) (*v1alpha1.GoogleComputeHttpHealthCheck, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleComputeHttpHealthCheck, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleComputeHttpHealthCheckList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComputeHttpHealthCheck, err error)
	GoogleComputeHttpHealthCheckExpansion
}

// googleComputeHttpHealthChecks implements GoogleComputeHttpHealthCheckInterface
type googleComputeHttpHealthChecks struct {
	client rest.Interface
}

// newGoogleComputeHttpHealthChecks returns a GoogleComputeHttpHealthChecks
func newGoogleComputeHttpHealthChecks(c *GoogleV1alpha1Client) *googleComputeHttpHealthChecks {
	return &googleComputeHttpHealthChecks{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleComputeHttpHealthCheck, and returns the corresponding googleComputeHttpHealthCheck object, and an error if there is any.
func (c *googleComputeHttpHealthChecks) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleComputeHttpHealthCheck, err error) {
	result = &v1alpha1.GoogleComputeHttpHealthCheck{}
	err = c.client.Get().
		Resource("googlecomputehttphealthchecks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleComputeHttpHealthChecks that match those selectors.
func (c *googleComputeHttpHealthChecks) List(opts v1.ListOptions) (result *v1alpha1.GoogleComputeHttpHealthCheckList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleComputeHttpHealthCheckList{}
	err = c.client.Get().
		Resource("googlecomputehttphealthchecks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleComputeHttpHealthChecks.
func (c *googleComputeHttpHealthChecks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googlecomputehttphealthchecks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleComputeHttpHealthCheck and creates it.  Returns the server's representation of the googleComputeHttpHealthCheck, and an error, if there is any.
func (c *googleComputeHttpHealthChecks) Create(googleComputeHttpHealthCheck *v1alpha1.GoogleComputeHttpHealthCheck) (result *v1alpha1.GoogleComputeHttpHealthCheck, err error) {
	result = &v1alpha1.GoogleComputeHttpHealthCheck{}
	err = c.client.Post().
		Resource("googlecomputehttphealthchecks").
		Body(googleComputeHttpHealthCheck).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleComputeHttpHealthCheck and updates it. Returns the server's representation of the googleComputeHttpHealthCheck, and an error, if there is any.
func (c *googleComputeHttpHealthChecks) Update(googleComputeHttpHealthCheck *v1alpha1.GoogleComputeHttpHealthCheck) (result *v1alpha1.GoogleComputeHttpHealthCheck, err error) {
	result = &v1alpha1.GoogleComputeHttpHealthCheck{}
	err = c.client.Put().
		Resource("googlecomputehttphealthchecks").
		Name(googleComputeHttpHealthCheck.Name).
		Body(googleComputeHttpHealthCheck).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleComputeHttpHealthChecks) UpdateStatus(googleComputeHttpHealthCheck *v1alpha1.GoogleComputeHttpHealthCheck) (result *v1alpha1.GoogleComputeHttpHealthCheck, err error) {
	result = &v1alpha1.GoogleComputeHttpHealthCheck{}
	err = c.client.Put().
		Resource("googlecomputehttphealthchecks").
		Name(googleComputeHttpHealthCheck.Name).
		SubResource("status").
		Body(googleComputeHttpHealthCheck).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleComputeHttpHealthCheck and deletes it. Returns an error if one occurs.
func (c *googleComputeHttpHealthChecks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googlecomputehttphealthchecks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleComputeHttpHealthChecks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googlecomputehttphealthchecks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleComputeHttpHealthCheck.
func (c *googleComputeHttpHealthChecks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComputeHttpHealthCheck, err error) {
	result = &v1alpha1.GoogleComputeHttpHealthCheck{}
	err = c.client.Patch(pt).
		Resource("googlecomputehttphealthchecks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}