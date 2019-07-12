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

// AzurermDnsMxRecordsGetter has a method to return a AzurermDnsMxRecordInterface.
// A group's client should implement this interface.
type AzurermDnsMxRecordsGetter interface {
	AzurermDnsMxRecords() AzurermDnsMxRecordInterface
}

// AzurermDnsMxRecordInterface has methods to work with AzurermDnsMxRecord resources.
type AzurermDnsMxRecordInterface interface {
	Create(*v1alpha1.AzurermDnsMxRecord) (*v1alpha1.AzurermDnsMxRecord, error)
	Update(*v1alpha1.AzurermDnsMxRecord) (*v1alpha1.AzurermDnsMxRecord, error)
	UpdateStatus(*v1alpha1.AzurermDnsMxRecord) (*v1alpha1.AzurermDnsMxRecord, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermDnsMxRecord, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermDnsMxRecordList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDnsMxRecord, err error)
	AzurermDnsMxRecordExpansion
}

// azurermDnsMxRecords implements AzurermDnsMxRecordInterface
type azurermDnsMxRecords struct {
	client rest.Interface
}

// newAzurermDnsMxRecords returns a AzurermDnsMxRecords
func newAzurermDnsMxRecords(c *AzurermV1alpha1Client) *azurermDnsMxRecords {
	return &azurermDnsMxRecords{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermDnsMxRecord, and returns the corresponding azurermDnsMxRecord object, and an error if there is any.
func (c *azurermDnsMxRecords) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermDnsMxRecord, err error) {
	result = &v1alpha1.AzurermDnsMxRecord{}
	err = c.client.Get().
		Resource("azurermdnsmxrecords").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermDnsMxRecords that match those selectors.
func (c *azurermDnsMxRecords) List(opts v1.ListOptions) (result *v1alpha1.AzurermDnsMxRecordList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermDnsMxRecordList{}
	err = c.client.Get().
		Resource("azurermdnsmxrecords").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermDnsMxRecords.
func (c *azurermDnsMxRecords) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermdnsmxrecords").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermDnsMxRecord and creates it.  Returns the server's representation of the azurermDnsMxRecord, and an error, if there is any.
func (c *azurermDnsMxRecords) Create(azurermDnsMxRecord *v1alpha1.AzurermDnsMxRecord) (result *v1alpha1.AzurermDnsMxRecord, err error) {
	result = &v1alpha1.AzurermDnsMxRecord{}
	err = c.client.Post().
		Resource("azurermdnsmxrecords").
		Body(azurermDnsMxRecord).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermDnsMxRecord and updates it. Returns the server's representation of the azurermDnsMxRecord, and an error, if there is any.
func (c *azurermDnsMxRecords) Update(azurermDnsMxRecord *v1alpha1.AzurermDnsMxRecord) (result *v1alpha1.AzurermDnsMxRecord, err error) {
	result = &v1alpha1.AzurermDnsMxRecord{}
	err = c.client.Put().
		Resource("azurermdnsmxrecords").
		Name(azurermDnsMxRecord.Name).
		Body(azurermDnsMxRecord).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermDnsMxRecords) UpdateStatus(azurermDnsMxRecord *v1alpha1.AzurermDnsMxRecord) (result *v1alpha1.AzurermDnsMxRecord, err error) {
	result = &v1alpha1.AzurermDnsMxRecord{}
	err = c.client.Put().
		Resource("azurermdnsmxrecords").
		Name(azurermDnsMxRecord.Name).
		SubResource("status").
		Body(azurermDnsMxRecord).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermDnsMxRecord and deletes it. Returns an error if one occurs.
func (c *azurermDnsMxRecords) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermdnsmxrecords").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermDnsMxRecords) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermdnsmxrecords").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermDnsMxRecord.
func (c *azurermDnsMxRecords) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDnsMxRecord, err error) {
	result = &v1alpha1.AzurermDnsMxRecord{}
	err = c.client.Patch(pt).
		Resource("azurermdnsmxrecords").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}