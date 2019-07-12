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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// AzurermSubnetsGetter has a method to return a AzurermSubnetInterface.
// A group's client should implement this interface.
type AzurermSubnetsGetter interface {
	AzurermSubnets() AzurermSubnetInterface
}

// AzurermSubnetInterface has methods to work with AzurermSubnet resources.
type AzurermSubnetInterface interface {
	Create(*v1alpha1.AzurermSubnet) (*v1alpha1.AzurermSubnet, error)
	Update(*v1alpha1.AzurermSubnet) (*v1alpha1.AzurermSubnet, error)
	UpdateStatus(*v1alpha1.AzurermSubnet) (*v1alpha1.AzurermSubnet, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermSubnet, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermSubnetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermSubnet, err error)
	AzurermSubnetExpansion
}

// azurermSubnets implements AzurermSubnetInterface
type azurermSubnets struct {
	client rest.Interface
}

// newAzurermSubnets returns a AzurermSubnets
func newAzurermSubnets(c *AzurermV1alpha1Client) *azurermSubnets {
	return &azurermSubnets{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermSubnet, and returns the corresponding azurermSubnet object, and an error if there is any.
func (c *azurermSubnets) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermSubnet, err error) {
	result = &v1alpha1.AzurermSubnet{}
	err = c.client.Get().
		Resource("azurermsubnets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermSubnets that match those selectors.
func (c *azurermSubnets) List(opts v1.ListOptions) (result *v1alpha1.AzurermSubnetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermSubnetList{}
	err = c.client.Get().
		Resource("azurermsubnets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermSubnets.
func (c *azurermSubnets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermsubnets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermSubnet and creates it.  Returns the server's representation of the azurermSubnet, and an error, if there is any.
func (c *azurermSubnets) Create(azurermSubnet *v1alpha1.AzurermSubnet) (result *v1alpha1.AzurermSubnet, err error) {
	result = &v1alpha1.AzurermSubnet{}
	err = c.client.Post().
		Resource("azurermsubnets").
		Body(azurermSubnet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermSubnet and updates it. Returns the server's representation of the azurermSubnet, and an error, if there is any.
func (c *azurermSubnets) Update(azurermSubnet *v1alpha1.AzurermSubnet) (result *v1alpha1.AzurermSubnet, err error) {
	result = &v1alpha1.AzurermSubnet{}
	err = c.client.Put().
		Resource("azurermsubnets").
		Name(azurermSubnet.Name).
		Body(azurermSubnet).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermSubnets) UpdateStatus(azurermSubnet *v1alpha1.AzurermSubnet) (result *v1alpha1.AzurermSubnet, err error) {
	result = &v1alpha1.AzurermSubnet{}
	err = c.client.Put().
		Resource("azurermsubnets").
		Name(azurermSubnet.Name).
		SubResource("status").
		Body(azurermSubnet).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermSubnet and deletes it. Returns an error if one occurs.
func (c *azurermSubnets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermsubnets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermSubnets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermsubnets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermSubnet.
func (c *azurermSubnets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermSubnet, err error) {
	result = &v1alpha1.AzurermSubnet{}
	err = c.client.Patch(pt).
		Resource("azurermsubnets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}