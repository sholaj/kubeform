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

// AzurermApiManagementAuthorizationServersGetter has a method to return a AzurermApiManagementAuthorizationServerInterface.
// A group's client should implement this interface.
type AzurermApiManagementAuthorizationServersGetter interface {
	AzurermApiManagementAuthorizationServers() AzurermApiManagementAuthorizationServerInterface
}

// AzurermApiManagementAuthorizationServerInterface has methods to work with AzurermApiManagementAuthorizationServer resources.
type AzurermApiManagementAuthorizationServerInterface interface {
	Create(*v1alpha1.AzurermApiManagementAuthorizationServer) (*v1alpha1.AzurermApiManagementAuthorizationServer, error)
	Update(*v1alpha1.AzurermApiManagementAuthorizationServer) (*v1alpha1.AzurermApiManagementAuthorizationServer, error)
	UpdateStatus(*v1alpha1.AzurermApiManagementAuthorizationServer) (*v1alpha1.AzurermApiManagementAuthorizationServer, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermApiManagementAuthorizationServer, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermApiManagementAuthorizationServerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermApiManagementAuthorizationServer, err error)
	AzurermApiManagementAuthorizationServerExpansion
}

// azurermApiManagementAuthorizationServers implements AzurermApiManagementAuthorizationServerInterface
type azurermApiManagementAuthorizationServers struct {
	client rest.Interface
}

// newAzurermApiManagementAuthorizationServers returns a AzurermApiManagementAuthorizationServers
func newAzurermApiManagementAuthorizationServers(c *AzurermV1alpha1Client) *azurermApiManagementAuthorizationServers {
	return &azurermApiManagementAuthorizationServers{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermApiManagementAuthorizationServer, and returns the corresponding azurermApiManagementAuthorizationServer object, and an error if there is any.
func (c *azurermApiManagementAuthorizationServers) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermApiManagementAuthorizationServer, err error) {
	result = &v1alpha1.AzurermApiManagementAuthorizationServer{}
	err = c.client.Get().
		Resource("azurermapimanagementauthorizationservers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermApiManagementAuthorizationServers that match those selectors.
func (c *azurermApiManagementAuthorizationServers) List(opts v1.ListOptions) (result *v1alpha1.AzurermApiManagementAuthorizationServerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermApiManagementAuthorizationServerList{}
	err = c.client.Get().
		Resource("azurermapimanagementauthorizationservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermApiManagementAuthorizationServers.
func (c *azurermApiManagementAuthorizationServers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermapimanagementauthorizationservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermApiManagementAuthorizationServer and creates it.  Returns the server's representation of the azurermApiManagementAuthorizationServer, and an error, if there is any.
func (c *azurermApiManagementAuthorizationServers) Create(azurermApiManagementAuthorizationServer *v1alpha1.AzurermApiManagementAuthorizationServer) (result *v1alpha1.AzurermApiManagementAuthorizationServer, err error) {
	result = &v1alpha1.AzurermApiManagementAuthorizationServer{}
	err = c.client.Post().
		Resource("azurermapimanagementauthorizationservers").
		Body(azurermApiManagementAuthorizationServer).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermApiManagementAuthorizationServer and updates it. Returns the server's representation of the azurermApiManagementAuthorizationServer, and an error, if there is any.
func (c *azurermApiManagementAuthorizationServers) Update(azurermApiManagementAuthorizationServer *v1alpha1.AzurermApiManagementAuthorizationServer) (result *v1alpha1.AzurermApiManagementAuthorizationServer, err error) {
	result = &v1alpha1.AzurermApiManagementAuthorizationServer{}
	err = c.client.Put().
		Resource("azurermapimanagementauthorizationservers").
		Name(azurermApiManagementAuthorizationServer.Name).
		Body(azurermApiManagementAuthorizationServer).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermApiManagementAuthorizationServers) UpdateStatus(azurermApiManagementAuthorizationServer *v1alpha1.AzurermApiManagementAuthorizationServer) (result *v1alpha1.AzurermApiManagementAuthorizationServer, err error) {
	result = &v1alpha1.AzurermApiManagementAuthorizationServer{}
	err = c.client.Put().
		Resource("azurermapimanagementauthorizationservers").
		Name(azurermApiManagementAuthorizationServer.Name).
		SubResource("status").
		Body(azurermApiManagementAuthorizationServer).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermApiManagementAuthorizationServer and deletes it. Returns an error if one occurs.
func (c *azurermApiManagementAuthorizationServers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermapimanagementauthorizationservers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermApiManagementAuthorizationServers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermapimanagementauthorizationservers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermApiManagementAuthorizationServer.
func (c *azurermApiManagementAuthorizationServers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermApiManagementAuthorizationServer, err error) {
	result = &v1alpha1.AzurermApiManagementAuthorizationServer{}
	err = c.client.Patch(pt).
		Resource("azurermapimanagementauthorizationservers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}