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

// AzurermLogicAppTriggerCustomsGetter has a method to return a AzurermLogicAppTriggerCustomInterface.
// A group's client should implement this interface.
type AzurermLogicAppTriggerCustomsGetter interface {
	AzurermLogicAppTriggerCustoms() AzurermLogicAppTriggerCustomInterface
}

// AzurermLogicAppTriggerCustomInterface has methods to work with AzurermLogicAppTriggerCustom resources.
type AzurermLogicAppTriggerCustomInterface interface {
	Create(*v1alpha1.AzurermLogicAppTriggerCustom) (*v1alpha1.AzurermLogicAppTriggerCustom, error)
	Update(*v1alpha1.AzurermLogicAppTriggerCustom) (*v1alpha1.AzurermLogicAppTriggerCustom, error)
	UpdateStatus(*v1alpha1.AzurermLogicAppTriggerCustom) (*v1alpha1.AzurermLogicAppTriggerCustom, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermLogicAppTriggerCustom, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermLogicAppTriggerCustomList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermLogicAppTriggerCustom, err error)
	AzurermLogicAppTriggerCustomExpansion
}

// azurermLogicAppTriggerCustoms implements AzurermLogicAppTriggerCustomInterface
type azurermLogicAppTriggerCustoms struct {
	client rest.Interface
}

// newAzurermLogicAppTriggerCustoms returns a AzurermLogicAppTriggerCustoms
func newAzurermLogicAppTriggerCustoms(c *AzurermV1alpha1Client) *azurermLogicAppTriggerCustoms {
	return &azurermLogicAppTriggerCustoms{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermLogicAppTriggerCustom, and returns the corresponding azurermLogicAppTriggerCustom object, and an error if there is any.
func (c *azurermLogicAppTriggerCustoms) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermLogicAppTriggerCustom, err error) {
	result = &v1alpha1.AzurermLogicAppTriggerCustom{}
	err = c.client.Get().
		Resource("azurermlogicapptriggercustoms").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermLogicAppTriggerCustoms that match those selectors.
func (c *azurermLogicAppTriggerCustoms) List(opts v1.ListOptions) (result *v1alpha1.AzurermLogicAppTriggerCustomList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermLogicAppTriggerCustomList{}
	err = c.client.Get().
		Resource("azurermlogicapptriggercustoms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermLogicAppTriggerCustoms.
func (c *azurermLogicAppTriggerCustoms) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermlogicapptriggercustoms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermLogicAppTriggerCustom and creates it.  Returns the server's representation of the azurermLogicAppTriggerCustom, and an error, if there is any.
func (c *azurermLogicAppTriggerCustoms) Create(azurermLogicAppTriggerCustom *v1alpha1.AzurermLogicAppTriggerCustom) (result *v1alpha1.AzurermLogicAppTriggerCustom, err error) {
	result = &v1alpha1.AzurermLogicAppTriggerCustom{}
	err = c.client.Post().
		Resource("azurermlogicapptriggercustoms").
		Body(azurermLogicAppTriggerCustom).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermLogicAppTriggerCustom and updates it. Returns the server's representation of the azurermLogicAppTriggerCustom, and an error, if there is any.
func (c *azurermLogicAppTriggerCustoms) Update(azurermLogicAppTriggerCustom *v1alpha1.AzurermLogicAppTriggerCustom) (result *v1alpha1.AzurermLogicAppTriggerCustom, err error) {
	result = &v1alpha1.AzurermLogicAppTriggerCustom{}
	err = c.client.Put().
		Resource("azurermlogicapptriggercustoms").
		Name(azurermLogicAppTriggerCustom.Name).
		Body(azurermLogicAppTriggerCustom).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermLogicAppTriggerCustoms) UpdateStatus(azurermLogicAppTriggerCustom *v1alpha1.AzurermLogicAppTriggerCustom) (result *v1alpha1.AzurermLogicAppTriggerCustom, err error) {
	result = &v1alpha1.AzurermLogicAppTriggerCustom{}
	err = c.client.Put().
		Resource("azurermlogicapptriggercustoms").
		Name(azurermLogicAppTriggerCustom.Name).
		SubResource("status").
		Body(azurermLogicAppTriggerCustom).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermLogicAppTriggerCustom and deletes it. Returns an error if one occurs.
func (c *azurermLogicAppTriggerCustoms) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermlogicapptriggercustoms").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermLogicAppTriggerCustoms) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermlogicapptriggercustoms").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermLogicAppTriggerCustom.
func (c *azurermLogicAppTriggerCustoms) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermLogicAppTriggerCustom, err error) {
	result = &v1alpha1.AzurermLogicAppTriggerCustom{}
	err = c.client.Patch(pt).
		Resource("azurermlogicapptriggercustoms").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}