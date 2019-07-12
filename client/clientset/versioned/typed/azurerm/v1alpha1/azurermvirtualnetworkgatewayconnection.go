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

// AzurermVirtualNetworkGatewayConnectionsGetter has a method to return a AzurermVirtualNetworkGatewayConnectionInterface.
// A group's client should implement this interface.
type AzurermVirtualNetworkGatewayConnectionsGetter interface {
	AzurermVirtualNetworkGatewayConnections() AzurermVirtualNetworkGatewayConnectionInterface
}

// AzurermVirtualNetworkGatewayConnectionInterface has methods to work with AzurermVirtualNetworkGatewayConnection resources.
type AzurermVirtualNetworkGatewayConnectionInterface interface {
	Create(*v1alpha1.AzurermVirtualNetworkGatewayConnection) (*v1alpha1.AzurermVirtualNetworkGatewayConnection, error)
	Update(*v1alpha1.AzurermVirtualNetworkGatewayConnection) (*v1alpha1.AzurermVirtualNetworkGatewayConnection, error)
	UpdateStatus(*v1alpha1.AzurermVirtualNetworkGatewayConnection) (*v1alpha1.AzurermVirtualNetworkGatewayConnection, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermVirtualNetworkGatewayConnection, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermVirtualNetworkGatewayConnectionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermVirtualNetworkGatewayConnection, err error)
	AzurermVirtualNetworkGatewayConnectionExpansion
}

// azurermVirtualNetworkGatewayConnections implements AzurermVirtualNetworkGatewayConnectionInterface
type azurermVirtualNetworkGatewayConnections struct {
	client rest.Interface
}

// newAzurermVirtualNetworkGatewayConnections returns a AzurermVirtualNetworkGatewayConnections
func newAzurermVirtualNetworkGatewayConnections(c *AzurermV1alpha1Client) *azurermVirtualNetworkGatewayConnections {
	return &azurermVirtualNetworkGatewayConnections{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermVirtualNetworkGatewayConnection, and returns the corresponding azurermVirtualNetworkGatewayConnection object, and an error if there is any.
func (c *azurermVirtualNetworkGatewayConnections) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermVirtualNetworkGatewayConnection, err error) {
	result = &v1alpha1.AzurermVirtualNetworkGatewayConnection{}
	err = c.client.Get().
		Resource("azurermvirtualnetworkgatewayconnections").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermVirtualNetworkGatewayConnections that match those selectors.
func (c *azurermVirtualNetworkGatewayConnections) List(opts v1.ListOptions) (result *v1alpha1.AzurermVirtualNetworkGatewayConnectionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermVirtualNetworkGatewayConnectionList{}
	err = c.client.Get().
		Resource("azurermvirtualnetworkgatewayconnections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermVirtualNetworkGatewayConnections.
func (c *azurermVirtualNetworkGatewayConnections) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermvirtualnetworkgatewayconnections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermVirtualNetworkGatewayConnection and creates it.  Returns the server's representation of the azurermVirtualNetworkGatewayConnection, and an error, if there is any.
func (c *azurermVirtualNetworkGatewayConnections) Create(azurermVirtualNetworkGatewayConnection *v1alpha1.AzurermVirtualNetworkGatewayConnection) (result *v1alpha1.AzurermVirtualNetworkGatewayConnection, err error) {
	result = &v1alpha1.AzurermVirtualNetworkGatewayConnection{}
	err = c.client.Post().
		Resource("azurermvirtualnetworkgatewayconnections").
		Body(azurermVirtualNetworkGatewayConnection).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermVirtualNetworkGatewayConnection and updates it. Returns the server's representation of the azurermVirtualNetworkGatewayConnection, and an error, if there is any.
func (c *azurermVirtualNetworkGatewayConnections) Update(azurermVirtualNetworkGatewayConnection *v1alpha1.AzurermVirtualNetworkGatewayConnection) (result *v1alpha1.AzurermVirtualNetworkGatewayConnection, err error) {
	result = &v1alpha1.AzurermVirtualNetworkGatewayConnection{}
	err = c.client.Put().
		Resource("azurermvirtualnetworkgatewayconnections").
		Name(azurermVirtualNetworkGatewayConnection.Name).
		Body(azurermVirtualNetworkGatewayConnection).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermVirtualNetworkGatewayConnections) UpdateStatus(azurermVirtualNetworkGatewayConnection *v1alpha1.AzurermVirtualNetworkGatewayConnection) (result *v1alpha1.AzurermVirtualNetworkGatewayConnection, err error) {
	result = &v1alpha1.AzurermVirtualNetworkGatewayConnection{}
	err = c.client.Put().
		Resource("azurermvirtualnetworkgatewayconnections").
		Name(azurermVirtualNetworkGatewayConnection.Name).
		SubResource("status").
		Body(azurermVirtualNetworkGatewayConnection).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermVirtualNetworkGatewayConnection and deletes it. Returns an error if one occurs.
func (c *azurermVirtualNetworkGatewayConnections) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermvirtualnetworkgatewayconnections").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermVirtualNetworkGatewayConnections) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermvirtualnetworkgatewayconnections").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermVirtualNetworkGatewayConnection.
func (c *azurermVirtualNetworkGatewayConnections) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermVirtualNetworkGatewayConnection, err error) {
	result = &v1alpha1.AzurermVirtualNetworkGatewayConnection{}
	err = c.client.Patch(pt).
		Resource("azurermvirtualnetworkgatewayconnections").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}