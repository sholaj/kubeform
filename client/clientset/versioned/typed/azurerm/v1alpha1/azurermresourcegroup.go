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

// AzurermResourceGroupsGetter has a method to return a AzurermResourceGroupInterface.
// A group's client should implement this interface.
type AzurermResourceGroupsGetter interface {
	AzurermResourceGroups() AzurermResourceGroupInterface
}

// AzurermResourceGroupInterface has methods to work with AzurermResourceGroup resources.
type AzurermResourceGroupInterface interface {
	Create(*v1alpha1.AzurermResourceGroup) (*v1alpha1.AzurermResourceGroup, error)
	Update(*v1alpha1.AzurermResourceGroup) (*v1alpha1.AzurermResourceGroup, error)
	UpdateStatus(*v1alpha1.AzurermResourceGroup) (*v1alpha1.AzurermResourceGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermResourceGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermResourceGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermResourceGroup, err error)
	AzurermResourceGroupExpansion
}

// azurermResourceGroups implements AzurermResourceGroupInterface
type azurermResourceGroups struct {
	client rest.Interface
}

// newAzurermResourceGroups returns a AzurermResourceGroups
func newAzurermResourceGroups(c *AzurermV1alpha1Client) *azurermResourceGroups {
	return &azurermResourceGroups{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermResourceGroup, and returns the corresponding azurermResourceGroup object, and an error if there is any.
func (c *azurermResourceGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermResourceGroup, err error) {
	result = &v1alpha1.AzurermResourceGroup{}
	err = c.client.Get().
		Resource("azurermresourcegroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermResourceGroups that match those selectors.
func (c *azurermResourceGroups) List(opts v1.ListOptions) (result *v1alpha1.AzurermResourceGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermResourceGroupList{}
	err = c.client.Get().
		Resource("azurermresourcegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermResourceGroups.
func (c *azurermResourceGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermresourcegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermResourceGroup and creates it.  Returns the server's representation of the azurermResourceGroup, and an error, if there is any.
func (c *azurermResourceGroups) Create(azurermResourceGroup *v1alpha1.AzurermResourceGroup) (result *v1alpha1.AzurermResourceGroup, err error) {
	result = &v1alpha1.AzurermResourceGroup{}
	err = c.client.Post().
		Resource("azurermresourcegroups").
		Body(azurermResourceGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermResourceGroup and updates it. Returns the server's representation of the azurermResourceGroup, and an error, if there is any.
func (c *azurermResourceGroups) Update(azurermResourceGroup *v1alpha1.AzurermResourceGroup) (result *v1alpha1.AzurermResourceGroup, err error) {
	result = &v1alpha1.AzurermResourceGroup{}
	err = c.client.Put().
		Resource("azurermresourcegroups").
		Name(azurermResourceGroup.Name).
		Body(azurermResourceGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermResourceGroups) UpdateStatus(azurermResourceGroup *v1alpha1.AzurermResourceGroup) (result *v1alpha1.AzurermResourceGroup, err error) {
	result = &v1alpha1.AzurermResourceGroup{}
	err = c.client.Put().
		Resource("azurermresourcegroups").
		Name(azurermResourceGroup.Name).
		SubResource("status").
		Body(azurermResourceGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermResourceGroup and deletes it. Returns an error if one occurs.
func (c *azurermResourceGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermresourcegroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermResourceGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermresourcegroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermResourceGroup.
func (c *azurermResourceGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermResourceGroup, err error) {
	result = &v1alpha1.AzurermResourceGroup{}
	err = c.client.Patch(pt).
		Resource("azurermresourcegroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}