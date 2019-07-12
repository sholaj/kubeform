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
	v1alpha1 "kubeform.dev/kubeform/apis/digitalocean/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// DigitaloceanFloatingIpsGetter has a method to return a DigitaloceanFloatingIpInterface.
// A group's client should implement this interface.
type DigitaloceanFloatingIpsGetter interface {
	DigitaloceanFloatingIps() DigitaloceanFloatingIpInterface
}

// DigitaloceanFloatingIpInterface has methods to work with DigitaloceanFloatingIp resources.
type DigitaloceanFloatingIpInterface interface {
	Create(*v1alpha1.DigitaloceanFloatingIp) (*v1alpha1.DigitaloceanFloatingIp, error)
	Update(*v1alpha1.DigitaloceanFloatingIp) (*v1alpha1.DigitaloceanFloatingIp, error)
	UpdateStatus(*v1alpha1.DigitaloceanFloatingIp) (*v1alpha1.DigitaloceanFloatingIp, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DigitaloceanFloatingIp, error)
	List(opts v1.ListOptions) (*v1alpha1.DigitaloceanFloatingIpList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DigitaloceanFloatingIp, err error)
	DigitaloceanFloatingIpExpansion
}

// digitaloceanFloatingIps implements DigitaloceanFloatingIpInterface
type digitaloceanFloatingIps struct {
	client rest.Interface
}

// newDigitaloceanFloatingIps returns a DigitaloceanFloatingIps
func newDigitaloceanFloatingIps(c *DigitaloceanV1alpha1Client) *digitaloceanFloatingIps {
	return &digitaloceanFloatingIps{
		client: c.RESTClient(),
	}
}

// Get takes name of the digitaloceanFloatingIp, and returns the corresponding digitaloceanFloatingIp object, and an error if there is any.
func (c *digitaloceanFloatingIps) Get(name string, options v1.GetOptions) (result *v1alpha1.DigitaloceanFloatingIp, err error) {
	result = &v1alpha1.DigitaloceanFloatingIp{}
	err = c.client.Get().
		Resource("digitaloceanfloatingips").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DigitaloceanFloatingIps that match those selectors.
func (c *digitaloceanFloatingIps) List(opts v1.ListOptions) (result *v1alpha1.DigitaloceanFloatingIpList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DigitaloceanFloatingIpList{}
	err = c.client.Get().
		Resource("digitaloceanfloatingips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested digitaloceanFloatingIps.
func (c *digitaloceanFloatingIps) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("digitaloceanfloatingips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a digitaloceanFloatingIp and creates it.  Returns the server's representation of the digitaloceanFloatingIp, and an error, if there is any.
func (c *digitaloceanFloatingIps) Create(digitaloceanFloatingIp *v1alpha1.DigitaloceanFloatingIp) (result *v1alpha1.DigitaloceanFloatingIp, err error) {
	result = &v1alpha1.DigitaloceanFloatingIp{}
	err = c.client.Post().
		Resource("digitaloceanfloatingips").
		Body(digitaloceanFloatingIp).
		Do().
		Into(result)
	return
}

// Update takes the representation of a digitaloceanFloatingIp and updates it. Returns the server's representation of the digitaloceanFloatingIp, and an error, if there is any.
func (c *digitaloceanFloatingIps) Update(digitaloceanFloatingIp *v1alpha1.DigitaloceanFloatingIp) (result *v1alpha1.DigitaloceanFloatingIp, err error) {
	result = &v1alpha1.DigitaloceanFloatingIp{}
	err = c.client.Put().
		Resource("digitaloceanfloatingips").
		Name(digitaloceanFloatingIp.Name).
		Body(digitaloceanFloatingIp).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *digitaloceanFloatingIps) UpdateStatus(digitaloceanFloatingIp *v1alpha1.DigitaloceanFloatingIp) (result *v1alpha1.DigitaloceanFloatingIp, err error) {
	result = &v1alpha1.DigitaloceanFloatingIp{}
	err = c.client.Put().
		Resource("digitaloceanfloatingips").
		Name(digitaloceanFloatingIp.Name).
		SubResource("status").
		Body(digitaloceanFloatingIp).
		Do().
		Into(result)
	return
}

// Delete takes name of the digitaloceanFloatingIp and deletes it. Returns an error if one occurs.
func (c *digitaloceanFloatingIps) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("digitaloceanfloatingips").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *digitaloceanFloatingIps) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("digitaloceanfloatingips").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched digitaloceanFloatingIp.
func (c *digitaloceanFloatingIps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DigitaloceanFloatingIp, err error) {
	result = &v1alpha1.DigitaloceanFloatingIp{}
	err = c.client.Patch(pt).
		Resource("digitaloceanfloatingips").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}