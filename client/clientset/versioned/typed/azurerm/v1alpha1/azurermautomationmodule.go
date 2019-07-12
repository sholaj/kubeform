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

// AzurermAutomationModulesGetter has a method to return a AzurermAutomationModuleInterface.
// A group's client should implement this interface.
type AzurermAutomationModulesGetter interface {
	AzurermAutomationModules() AzurermAutomationModuleInterface
}

// AzurermAutomationModuleInterface has methods to work with AzurermAutomationModule resources.
type AzurermAutomationModuleInterface interface {
	Create(*v1alpha1.AzurermAutomationModule) (*v1alpha1.AzurermAutomationModule, error)
	Update(*v1alpha1.AzurermAutomationModule) (*v1alpha1.AzurermAutomationModule, error)
	UpdateStatus(*v1alpha1.AzurermAutomationModule) (*v1alpha1.AzurermAutomationModule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermAutomationModule, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermAutomationModuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermAutomationModule, err error)
	AzurermAutomationModuleExpansion
}

// azurermAutomationModules implements AzurermAutomationModuleInterface
type azurermAutomationModules struct {
	client rest.Interface
}

// newAzurermAutomationModules returns a AzurermAutomationModules
func newAzurermAutomationModules(c *AzurermV1alpha1Client) *azurermAutomationModules {
	return &azurermAutomationModules{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermAutomationModule, and returns the corresponding azurermAutomationModule object, and an error if there is any.
func (c *azurermAutomationModules) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermAutomationModule, err error) {
	result = &v1alpha1.AzurermAutomationModule{}
	err = c.client.Get().
		Resource("azurermautomationmodules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermAutomationModules that match those selectors.
func (c *azurermAutomationModules) List(opts v1.ListOptions) (result *v1alpha1.AzurermAutomationModuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermAutomationModuleList{}
	err = c.client.Get().
		Resource("azurermautomationmodules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermAutomationModules.
func (c *azurermAutomationModules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermautomationmodules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermAutomationModule and creates it.  Returns the server's representation of the azurermAutomationModule, and an error, if there is any.
func (c *azurermAutomationModules) Create(azurermAutomationModule *v1alpha1.AzurermAutomationModule) (result *v1alpha1.AzurermAutomationModule, err error) {
	result = &v1alpha1.AzurermAutomationModule{}
	err = c.client.Post().
		Resource("azurermautomationmodules").
		Body(azurermAutomationModule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermAutomationModule and updates it. Returns the server's representation of the azurermAutomationModule, and an error, if there is any.
func (c *azurermAutomationModules) Update(azurermAutomationModule *v1alpha1.AzurermAutomationModule) (result *v1alpha1.AzurermAutomationModule, err error) {
	result = &v1alpha1.AzurermAutomationModule{}
	err = c.client.Put().
		Resource("azurermautomationmodules").
		Name(azurermAutomationModule.Name).
		Body(azurermAutomationModule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermAutomationModules) UpdateStatus(azurermAutomationModule *v1alpha1.AzurermAutomationModule) (result *v1alpha1.AzurermAutomationModule, err error) {
	result = &v1alpha1.AzurermAutomationModule{}
	err = c.client.Put().
		Resource("azurermautomationmodules").
		Name(azurermAutomationModule.Name).
		SubResource("status").
		Body(azurermAutomationModule).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermAutomationModule and deletes it. Returns an error if one occurs.
func (c *azurermAutomationModules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermautomationmodules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermAutomationModules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermautomationmodules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermAutomationModule.
func (c *azurermAutomationModules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermAutomationModule, err error) {
	result = &v1alpha1.AzurermAutomationModule{}
	err = c.client.Patch(pt).
		Resource("azurermautomationmodules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}