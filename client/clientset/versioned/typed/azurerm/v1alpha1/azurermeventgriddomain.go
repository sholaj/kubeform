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

// AzurermEventgridDomainsGetter has a method to return a AzurermEventgridDomainInterface.
// A group's client should implement this interface.
type AzurermEventgridDomainsGetter interface {
	AzurermEventgridDomains() AzurermEventgridDomainInterface
}

// AzurermEventgridDomainInterface has methods to work with AzurermEventgridDomain resources.
type AzurermEventgridDomainInterface interface {
	Create(*v1alpha1.AzurermEventgridDomain) (*v1alpha1.AzurermEventgridDomain, error)
	Update(*v1alpha1.AzurermEventgridDomain) (*v1alpha1.AzurermEventgridDomain, error)
	UpdateStatus(*v1alpha1.AzurermEventgridDomain) (*v1alpha1.AzurermEventgridDomain, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermEventgridDomain, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermEventgridDomainList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermEventgridDomain, err error)
	AzurermEventgridDomainExpansion
}

// azurermEventgridDomains implements AzurermEventgridDomainInterface
type azurermEventgridDomains struct {
	client rest.Interface
}

// newAzurermEventgridDomains returns a AzurermEventgridDomains
func newAzurermEventgridDomains(c *AzurermV1alpha1Client) *azurermEventgridDomains {
	return &azurermEventgridDomains{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermEventgridDomain, and returns the corresponding azurermEventgridDomain object, and an error if there is any.
func (c *azurermEventgridDomains) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermEventgridDomain, err error) {
	result = &v1alpha1.AzurermEventgridDomain{}
	err = c.client.Get().
		Resource("azurermeventgriddomains").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermEventgridDomains that match those selectors.
func (c *azurermEventgridDomains) List(opts v1.ListOptions) (result *v1alpha1.AzurermEventgridDomainList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermEventgridDomainList{}
	err = c.client.Get().
		Resource("azurermeventgriddomains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermEventgridDomains.
func (c *azurermEventgridDomains) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermeventgriddomains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermEventgridDomain and creates it.  Returns the server's representation of the azurermEventgridDomain, and an error, if there is any.
func (c *azurermEventgridDomains) Create(azurermEventgridDomain *v1alpha1.AzurermEventgridDomain) (result *v1alpha1.AzurermEventgridDomain, err error) {
	result = &v1alpha1.AzurermEventgridDomain{}
	err = c.client.Post().
		Resource("azurermeventgriddomains").
		Body(azurermEventgridDomain).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermEventgridDomain and updates it. Returns the server's representation of the azurermEventgridDomain, and an error, if there is any.
func (c *azurermEventgridDomains) Update(azurermEventgridDomain *v1alpha1.AzurermEventgridDomain) (result *v1alpha1.AzurermEventgridDomain, err error) {
	result = &v1alpha1.AzurermEventgridDomain{}
	err = c.client.Put().
		Resource("azurermeventgriddomains").
		Name(azurermEventgridDomain.Name).
		Body(azurermEventgridDomain).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermEventgridDomains) UpdateStatus(azurermEventgridDomain *v1alpha1.AzurermEventgridDomain) (result *v1alpha1.AzurermEventgridDomain, err error) {
	result = &v1alpha1.AzurermEventgridDomain{}
	err = c.client.Put().
		Resource("azurermeventgriddomains").
		Name(azurermEventgridDomain.Name).
		SubResource("status").
		Body(azurermEventgridDomain).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermEventgridDomain and deletes it. Returns an error if one occurs.
func (c *azurermEventgridDomains) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermeventgriddomains").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermEventgridDomains) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermeventgriddomains").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermEventgridDomain.
func (c *azurermEventgridDomains) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermEventgridDomain, err error) {
	result = &v1alpha1.AzurermEventgridDomain{}
	err = c.client.Patch(pt).
		Resource("azurermeventgriddomains").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}