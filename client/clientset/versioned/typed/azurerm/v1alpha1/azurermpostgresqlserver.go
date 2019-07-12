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

// AzurermPostgresqlServersGetter has a method to return a AzurermPostgresqlServerInterface.
// A group's client should implement this interface.
type AzurermPostgresqlServersGetter interface {
	AzurermPostgresqlServers() AzurermPostgresqlServerInterface
}

// AzurermPostgresqlServerInterface has methods to work with AzurermPostgresqlServer resources.
type AzurermPostgresqlServerInterface interface {
	Create(*v1alpha1.AzurermPostgresqlServer) (*v1alpha1.AzurermPostgresqlServer, error)
	Update(*v1alpha1.AzurermPostgresqlServer) (*v1alpha1.AzurermPostgresqlServer, error)
	UpdateStatus(*v1alpha1.AzurermPostgresqlServer) (*v1alpha1.AzurermPostgresqlServer, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermPostgresqlServer, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermPostgresqlServerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermPostgresqlServer, err error)
	AzurermPostgresqlServerExpansion
}

// azurermPostgresqlServers implements AzurermPostgresqlServerInterface
type azurermPostgresqlServers struct {
	client rest.Interface
}

// newAzurermPostgresqlServers returns a AzurermPostgresqlServers
func newAzurermPostgresqlServers(c *AzurermV1alpha1Client) *azurermPostgresqlServers {
	return &azurermPostgresqlServers{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermPostgresqlServer, and returns the corresponding azurermPostgresqlServer object, and an error if there is any.
func (c *azurermPostgresqlServers) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermPostgresqlServer, err error) {
	result = &v1alpha1.AzurermPostgresqlServer{}
	err = c.client.Get().
		Resource("azurermpostgresqlservers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermPostgresqlServers that match those selectors.
func (c *azurermPostgresqlServers) List(opts v1.ListOptions) (result *v1alpha1.AzurermPostgresqlServerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermPostgresqlServerList{}
	err = c.client.Get().
		Resource("azurermpostgresqlservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermPostgresqlServers.
func (c *azurermPostgresqlServers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermpostgresqlservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermPostgresqlServer and creates it.  Returns the server's representation of the azurermPostgresqlServer, and an error, if there is any.
func (c *azurermPostgresqlServers) Create(azurermPostgresqlServer *v1alpha1.AzurermPostgresqlServer) (result *v1alpha1.AzurermPostgresqlServer, err error) {
	result = &v1alpha1.AzurermPostgresqlServer{}
	err = c.client.Post().
		Resource("azurermpostgresqlservers").
		Body(azurermPostgresqlServer).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermPostgresqlServer and updates it. Returns the server's representation of the azurermPostgresqlServer, and an error, if there is any.
func (c *azurermPostgresqlServers) Update(azurermPostgresqlServer *v1alpha1.AzurermPostgresqlServer) (result *v1alpha1.AzurermPostgresqlServer, err error) {
	result = &v1alpha1.AzurermPostgresqlServer{}
	err = c.client.Put().
		Resource("azurermpostgresqlservers").
		Name(azurermPostgresqlServer.Name).
		Body(azurermPostgresqlServer).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermPostgresqlServers) UpdateStatus(azurermPostgresqlServer *v1alpha1.AzurermPostgresqlServer) (result *v1alpha1.AzurermPostgresqlServer, err error) {
	result = &v1alpha1.AzurermPostgresqlServer{}
	err = c.client.Put().
		Resource("azurermpostgresqlservers").
		Name(azurermPostgresqlServer.Name).
		SubResource("status").
		Body(azurermPostgresqlServer).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermPostgresqlServer and deletes it. Returns an error if one occurs.
func (c *azurermPostgresqlServers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermpostgresqlservers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermPostgresqlServers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermpostgresqlservers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermPostgresqlServer.
func (c *azurermPostgresqlServers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermPostgresqlServer, err error) {
	result = &v1alpha1.AzurermPostgresqlServer{}
	err = c.client.Patch(pt).
		Resource("azurermpostgresqlservers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}