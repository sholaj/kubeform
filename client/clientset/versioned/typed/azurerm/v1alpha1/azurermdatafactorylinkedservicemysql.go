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

// AzurermDataFactoryLinkedServiceMysqlsGetter has a method to return a AzurermDataFactoryLinkedServiceMysqlInterface.
// A group's client should implement this interface.
type AzurermDataFactoryLinkedServiceMysqlsGetter interface {
	AzurermDataFactoryLinkedServiceMysqls() AzurermDataFactoryLinkedServiceMysqlInterface
}

// AzurermDataFactoryLinkedServiceMysqlInterface has methods to work with AzurermDataFactoryLinkedServiceMysql resources.
type AzurermDataFactoryLinkedServiceMysqlInterface interface {
	Create(*v1alpha1.AzurermDataFactoryLinkedServiceMysql) (*v1alpha1.AzurermDataFactoryLinkedServiceMysql, error)
	Update(*v1alpha1.AzurermDataFactoryLinkedServiceMysql) (*v1alpha1.AzurermDataFactoryLinkedServiceMysql, error)
	UpdateStatus(*v1alpha1.AzurermDataFactoryLinkedServiceMysql) (*v1alpha1.AzurermDataFactoryLinkedServiceMysql, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermDataFactoryLinkedServiceMysql, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermDataFactoryLinkedServiceMysqlList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDataFactoryLinkedServiceMysql, err error)
	AzurermDataFactoryLinkedServiceMysqlExpansion
}

// azurermDataFactoryLinkedServiceMysqls implements AzurermDataFactoryLinkedServiceMysqlInterface
type azurermDataFactoryLinkedServiceMysqls struct {
	client rest.Interface
}

// newAzurermDataFactoryLinkedServiceMysqls returns a AzurermDataFactoryLinkedServiceMysqls
func newAzurermDataFactoryLinkedServiceMysqls(c *AzurermV1alpha1Client) *azurermDataFactoryLinkedServiceMysqls {
	return &azurermDataFactoryLinkedServiceMysqls{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermDataFactoryLinkedServiceMysql, and returns the corresponding azurermDataFactoryLinkedServiceMysql object, and an error if there is any.
func (c *azurermDataFactoryLinkedServiceMysqls) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermDataFactoryLinkedServiceMysql, err error) {
	result = &v1alpha1.AzurermDataFactoryLinkedServiceMysql{}
	err = c.client.Get().
		Resource("azurermdatafactorylinkedservicemysqls").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermDataFactoryLinkedServiceMysqls that match those selectors.
func (c *azurermDataFactoryLinkedServiceMysqls) List(opts v1.ListOptions) (result *v1alpha1.AzurermDataFactoryLinkedServiceMysqlList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermDataFactoryLinkedServiceMysqlList{}
	err = c.client.Get().
		Resource("azurermdatafactorylinkedservicemysqls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermDataFactoryLinkedServiceMysqls.
func (c *azurermDataFactoryLinkedServiceMysqls) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermdatafactorylinkedservicemysqls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermDataFactoryLinkedServiceMysql and creates it.  Returns the server's representation of the azurermDataFactoryLinkedServiceMysql, and an error, if there is any.
func (c *azurermDataFactoryLinkedServiceMysqls) Create(azurermDataFactoryLinkedServiceMysql *v1alpha1.AzurermDataFactoryLinkedServiceMysql) (result *v1alpha1.AzurermDataFactoryLinkedServiceMysql, err error) {
	result = &v1alpha1.AzurermDataFactoryLinkedServiceMysql{}
	err = c.client.Post().
		Resource("azurermdatafactorylinkedservicemysqls").
		Body(azurermDataFactoryLinkedServiceMysql).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermDataFactoryLinkedServiceMysql and updates it. Returns the server's representation of the azurermDataFactoryLinkedServiceMysql, and an error, if there is any.
func (c *azurermDataFactoryLinkedServiceMysqls) Update(azurermDataFactoryLinkedServiceMysql *v1alpha1.AzurermDataFactoryLinkedServiceMysql) (result *v1alpha1.AzurermDataFactoryLinkedServiceMysql, err error) {
	result = &v1alpha1.AzurermDataFactoryLinkedServiceMysql{}
	err = c.client.Put().
		Resource("azurermdatafactorylinkedservicemysqls").
		Name(azurermDataFactoryLinkedServiceMysql.Name).
		Body(azurermDataFactoryLinkedServiceMysql).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermDataFactoryLinkedServiceMysqls) UpdateStatus(azurermDataFactoryLinkedServiceMysql *v1alpha1.AzurermDataFactoryLinkedServiceMysql) (result *v1alpha1.AzurermDataFactoryLinkedServiceMysql, err error) {
	result = &v1alpha1.AzurermDataFactoryLinkedServiceMysql{}
	err = c.client.Put().
		Resource("azurermdatafactorylinkedservicemysqls").
		Name(azurermDataFactoryLinkedServiceMysql.Name).
		SubResource("status").
		Body(azurermDataFactoryLinkedServiceMysql).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermDataFactoryLinkedServiceMysql and deletes it. Returns an error if one occurs.
func (c *azurermDataFactoryLinkedServiceMysqls) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermdatafactorylinkedservicemysqls").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermDataFactoryLinkedServiceMysqls) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermdatafactorylinkedservicemysqls").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermDataFactoryLinkedServiceMysql.
func (c *azurermDataFactoryLinkedServiceMysqls) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDataFactoryLinkedServiceMysql, err error) {
	result = &v1alpha1.AzurermDataFactoryLinkedServiceMysql{}
	err = c.client.Patch(pt).
		Resource("azurermdatafactorylinkedservicemysqls").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}