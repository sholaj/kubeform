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

// AzurermLogicAppTriggerRecurrencesGetter has a method to return a AzurermLogicAppTriggerRecurrenceInterface.
// A group's client should implement this interface.
type AzurermLogicAppTriggerRecurrencesGetter interface {
	AzurermLogicAppTriggerRecurrences() AzurermLogicAppTriggerRecurrenceInterface
}

// AzurermLogicAppTriggerRecurrenceInterface has methods to work with AzurermLogicAppTriggerRecurrence resources.
type AzurermLogicAppTriggerRecurrenceInterface interface {
	Create(*v1alpha1.AzurermLogicAppTriggerRecurrence) (*v1alpha1.AzurermLogicAppTriggerRecurrence, error)
	Update(*v1alpha1.AzurermLogicAppTriggerRecurrence) (*v1alpha1.AzurermLogicAppTriggerRecurrence, error)
	UpdateStatus(*v1alpha1.AzurermLogicAppTriggerRecurrence) (*v1alpha1.AzurermLogicAppTriggerRecurrence, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermLogicAppTriggerRecurrence, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermLogicAppTriggerRecurrenceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermLogicAppTriggerRecurrence, err error)
	AzurermLogicAppTriggerRecurrenceExpansion
}

// azurermLogicAppTriggerRecurrences implements AzurermLogicAppTriggerRecurrenceInterface
type azurermLogicAppTriggerRecurrences struct {
	client rest.Interface
}

// newAzurermLogicAppTriggerRecurrences returns a AzurermLogicAppTriggerRecurrences
func newAzurermLogicAppTriggerRecurrences(c *AzurermV1alpha1Client) *azurermLogicAppTriggerRecurrences {
	return &azurermLogicAppTriggerRecurrences{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermLogicAppTriggerRecurrence, and returns the corresponding azurermLogicAppTriggerRecurrence object, and an error if there is any.
func (c *azurermLogicAppTriggerRecurrences) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermLogicAppTriggerRecurrence, err error) {
	result = &v1alpha1.AzurermLogicAppTriggerRecurrence{}
	err = c.client.Get().
		Resource("azurermlogicapptriggerrecurrences").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermLogicAppTriggerRecurrences that match those selectors.
func (c *azurermLogicAppTriggerRecurrences) List(opts v1.ListOptions) (result *v1alpha1.AzurermLogicAppTriggerRecurrenceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermLogicAppTriggerRecurrenceList{}
	err = c.client.Get().
		Resource("azurermlogicapptriggerrecurrences").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermLogicAppTriggerRecurrences.
func (c *azurermLogicAppTriggerRecurrences) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermlogicapptriggerrecurrences").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermLogicAppTriggerRecurrence and creates it.  Returns the server's representation of the azurermLogicAppTriggerRecurrence, and an error, if there is any.
func (c *azurermLogicAppTriggerRecurrences) Create(azurermLogicAppTriggerRecurrence *v1alpha1.AzurermLogicAppTriggerRecurrence) (result *v1alpha1.AzurermLogicAppTriggerRecurrence, err error) {
	result = &v1alpha1.AzurermLogicAppTriggerRecurrence{}
	err = c.client.Post().
		Resource("azurermlogicapptriggerrecurrences").
		Body(azurermLogicAppTriggerRecurrence).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermLogicAppTriggerRecurrence and updates it. Returns the server's representation of the azurermLogicAppTriggerRecurrence, and an error, if there is any.
func (c *azurermLogicAppTriggerRecurrences) Update(azurermLogicAppTriggerRecurrence *v1alpha1.AzurermLogicAppTriggerRecurrence) (result *v1alpha1.AzurermLogicAppTriggerRecurrence, err error) {
	result = &v1alpha1.AzurermLogicAppTriggerRecurrence{}
	err = c.client.Put().
		Resource("azurermlogicapptriggerrecurrences").
		Name(azurermLogicAppTriggerRecurrence.Name).
		Body(azurermLogicAppTriggerRecurrence).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermLogicAppTriggerRecurrences) UpdateStatus(azurermLogicAppTriggerRecurrence *v1alpha1.AzurermLogicAppTriggerRecurrence) (result *v1alpha1.AzurermLogicAppTriggerRecurrence, err error) {
	result = &v1alpha1.AzurermLogicAppTriggerRecurrence{}
	err = c.client.Put().
		Resource("azurermlogicapptriggerrecurrences").
		Name(azurermLogicAppTriggerRecurrence.Name).
		SubResource("status").
		Body(azurermLogicAppTriggerRecurrence).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermLogicAppTriggerRecurrence and deletes it. Returns an error if one occurs.
func (c *azurermLogicAppTriggerRecurrences) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermlogicapptriggerrecurrences").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermLogicAppTriggerRecurrences) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermlogicapptriggerrecurrences").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermLogicAppTriggerRecurrence.
func (c *azurermLogicAppTriggerRecurrences) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermLogicAppTriggerRecurrence, err error) {
	result = &v1alpha1.AzurermLogicAppTriggerRecurrence{}
	err = c.client.Patch(pt).
		Resource("azurermlogicapptriggerrecurrences").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}