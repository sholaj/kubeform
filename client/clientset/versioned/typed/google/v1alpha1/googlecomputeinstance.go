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

// GoogleComputeInstancesGetter has a method to return a GoogleComputeInstanceInterface.
// A group's client should implement this interface.
type GoogleComputeInstancesGetter interface {
	GoogleComputeInstances() GoogleComputeInstanceInterface
}

// GoogleComputeInstanceInterface has methods to work with GoogleComputeInstance resources.
type GoogleComputeInstanceInterface interface {
	Create(*v1alpha1.GoogleComputeInstance) (*v1alpha1.GoogleComputeInstance, error)
	Update(*v1alpha1.GoogleComputeInstance) (*v1alpha1.GoogleComputeInstance, error)
	UpdateStatus(*v1alpha1.GoogleComputeInstance) (*v1alpha1.GoogleComputeInstance, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleComputeInstance, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleComputeInstanceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComputeInstance, err error)
	GoogleComputeInstanceExpansion
}

// googleComputeInstances implements GoogleComputeInstanceInterface
type googleComputeInstances struct {
	client rest.Interface
}

// newGoogleComputeInstances returns a GoogleComputeInstances
func newGoogleComputeInstances(c *GoogleV1alpha1Client) *googleComputeInstances {
	return &googleComputeInstances{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleComputeInstance, and returns the corresponding googleComputeInstance object, and an error if there is any.
func (c *googleComputeInstances) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleComputeInstance, err error) {
	result = &v1alpha1.GoogleComputeInstance{}
	err = c.client.Get().
		Resource("googlecomputeinstances").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleComputeInstances that match those selectors.
func (c *googleComputeInstances) List(opts v1.ListOptions) (result *v1alpha1.GoogleComputeInstanceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleComputeInstanceList{}
	err = c.client.Get().
		Resource("googlecomputeinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleComputeInstances.
func (c *googleComputeInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googlecomputeinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleComputeInstance and creates it.  Returns the server's representation of the googleComputeInstance, and an error, if there is any.
func (c *googleComputeInstances) Create(googleComputeInstance *v1alpha1.GoogleComputeInstance) (result *v1alpha1.GoogleComputeInstance, err error) {
	result = &v1alpha1.GoogleComputeInstance{}
	err = c.client.Post().
		Resource("googlecomputeinstances").
		Body(googleComputeInstance).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleComputeInstance and updates it. Returns the server's representation of the googleComputeInstance, and an error, if there is any.
func (c *googleComputeInstances) Update(googleComputeInstance *v1alpha1.GoogleComputeInstance) (result *v1alpha1.GoogleComputeInstance, err error) {
	result = &v1alpha1.GoogleComputeInstance{}
	err = c.client.Put().
		Resource("googlecomputeinstances").
		Name(googleComputeInstance.Name).
		Body(googleComputeInstance).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleComputeInstances) UpdateStatus(googleComputeInstance *v1alpha1.GoogleComputeInstance) (result *v1alpha1.GoogleComputeInstance, err error) {
	result = &v1alpha1.GoogleComputeInstance{}
	err = c.client.Put().
		Resource("googlecomputeinstances").
		Name(googleComputeInstance.Name).
		SubResource("status").
		Body(googleComputeInstance).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleComputeInstance and deletes it. Returns an error if one occurs.
func (c *googleComputeInstances) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googlecomputeinstances").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleComputeInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googlecomputeinstances").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleComputeInstance.
func (c *googleComputeInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComputeInstance, err error) {
	result = &v1alpha1.GoogleComputeInstance{}
	err = c.client.Patch(pt).
		Resource("googlecomputeinstances").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}