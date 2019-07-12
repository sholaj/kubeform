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

// AzurermContainerRegistriesGetter has a method to return a AzurermContainerRegistryInterface.
// A group's client should implement this interface.
type AzurermContainerRegistriesGetter interface {
	AzurermContainerRegistries() AzurermContainerRegistryInterface
}

// AzurermContainerRegistryInterface has methods to work with AzurermContainerRegistry resources.
type AzurermContainerRegistryInterface interface {
	Create(*v1alpha1.AzurermContainerRegistry) (*v1alpha1.AzurermContainerRegistry, error)
	Update(*v1alpha1.AzurermContainerRegistry) (*v1alpha1.AzurermContainerRegistry, error)
	UpdateStatus(*v1alpha1.AzurermContainerRegistry) (*v1alpha1.AzurermContainerRegistry, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermContainerRegistry, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermContainerRegistryList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermContainerRegistry, err error)
	AzurermContainerRegistryExpansion
}

// azurermContainerRegistries implements AzurermContainerRegistryInterface
type azurermContainerRegistries struct {
	client rest.Interface
}

// newAzurermContainerRegistries returns a AzurermContainerRegistries
func newAzurermContainerRegistries(c *AzurermV1alpha1Client) *azurermContainerRegistries {
	return &azurermContainerRegistries{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermContainerRegistry, and returns the corresponding azurermContainerRegistry object, and an error if there is any.
func (c *azurermContainerRegistries) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermContainerRegistry, err error) {
	result = &v1alpha1.AzurermContainerRegistry{}
	err = c.client.Get().
		Resource("azurermcontainerregistries").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermContainerRegistries that match those selectors.
func (c *azurermContainerRegistries) List(opts v1.ListOptions) (result *v1alpha1.AzurermContainerRegistryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermContainerRegistryList{}
	err = c.client.Get().
		Resource("azurermcontainerregistries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermContainerRegistries.
func (c *azurermContainerRegistries) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermcontainerregistries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermContainerRegistry and creates it.  Returns the server's representation of the azurermContainerRegistry, and an error, if there is any.
func (c *azurermContainerRegistries) Create(azurermContainerRegistry *v1alpha1.AzurermContainerRegistry) (result *v1alpha1.AzurermContainerRegistry, err error) {
	result = &v1alpha1.AzurermContainerRegistry{}
	err = c.client.Post().
		Resource("azurermcontainerregistries").
		Body(azurermContainerRegistry).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermContainerRegistry and updates it. Returns the server's representation of the azurermContainerRegistry, and an error, if there is any.
func (c *azurermContainerRegistries) Update(azurermContainerRegistry *v1alpha1.AzurermContainerRegistry) (result *v1alpha1.AzurermContainerRegistry, err error) {
	result = &v1alpha1.AzurermContainerRegistry{}
	err = c.client.Put().
		Resource("azurermcontainerregistries").
		Name(azurermContainerRegistry.Name).
		Body(azurermContainerRegistry).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermContainerRegistries) UpdateStatus(azurermContainerRegistry *v1alpha1.AzurermContainerRegistry) (result *v1alpha1.AzurermContainerRegistry, err error) {
	result = &v1alpha1.AzurermContainerRegistry{}
	err = c.client.Put().
		Resource("azurermcontainerregistries").
		Name(azurermContainerRegistry.Name).
		SubResource("status").
		Body(azurermContainerRegistry).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermContainerRegistry and deletes it. Returns an error if one occurs.
func (c *azurermContainerRegistries) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermcontainerregistries").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermContainerRegistries) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermcontainerregistries").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermContainerRegistry.
func (c *azurermContainerRegistries) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermContainerRegistry, err error) {
	result = &v1alpha1.AzurermContainerRegistry{}
	err = c.client.Patch(pt).
		Resource("azurermcontainerregistries").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}