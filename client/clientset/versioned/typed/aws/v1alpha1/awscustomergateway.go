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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// AwsCustomerGatewaysGetter has a method to return a AwsCustomerGatewayInterface.
// A group's client should implement this interface.
type AwsCustomerGatewaysGetter interface {
	AwsCustomerGateways() AwsCustomerGatewayInterface
}

// AwsCustomerGatewayInterface has methods to work with AwsCustomerGateway resources.
type AwsCustomerGatewayInterface interface {
	Create(*v1alpha1.AwsCustomerGateway) (*v1alpha1.AwsCustomerGateway, error)
	Update(*v1alpha1.AwsCustomerGateway) (*v1alpha1.AwsCustomerGateway, error)
	UpdateStatus(*v1alpha1.AwsCustomerGateway) (*v1alpha1.AwsCustomerGateway, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsCustomerGateway, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsCustomerGatewayList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCustomerGateway, err error)
	AwsCustomerGatewayExpansion
}

// awsCustomerGateways implements AwsCustomerGatewayInterface
type awsCustomerGateways struct {
	client rest.Interface
}

// newAwsCustomerGateways returns a AwsCustomerGateways
func newAwsCustomerGateways(c *AwsV1alpha1Client) *awsCustomerGateways {
	return &awsCustomerGateways{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsCustomerGateway, and returns the corresponding awsCustomerGateway object, and an error if there is any.
func (c *awsCustomerGateways) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsCustomerGateway, err error) {
	result = &v1alpha1.AwsCustomerGateway{}
	err = c.client.Get().
		Resource("awscustomergateways").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsCustomerGateways that match those selectors.
func (c *awsCustomerGateways) List(opts v1.ListOptions) (result *v1alpha1.AwsCustomerGatewayList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsCustomerGatewayList{}
	err = c.client.Get().
		Resource("awscustomergateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsCustomerGateways.
func (c *awsCustomerGateways) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awscustomergateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsCustomerGateway and creates it.  Returns the server's representation of the awsCustomerGateway, and an error, if there is any.
func (c *awsCustomerGateways) Create(awsCustomerGateway *v1alpha1.AwsCustomerGateway) (result *v1alpha1.AwsCustomerGateway, err error) {
	result = &v1alpha1.AwsCustomerGateway{}
	err = c.client.Post().
		Resource("awscustomergateways").
		Body(awsCustomerGateway).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsCustomerGateway and updates it. Returns the server's representation of the awsCustomerGateway, and an error, if there is any.
func (c *awsCustomerGateways) Update(awsCustomerGateway *v1alpha1.AwsCustomerGateway) (result *v1alpha1.AwsCustomerGateway, err error) {
	result = &v1alpha1.AwsCustomerGateway{}
	err = c.client.Put().
		Resource("awscustomergateways").
		Name(awsCustomerGateway.Name).
		Body(awsCustomerGateway).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsCustomerGateways) UpdateStatus(awsCustomerGateway *v1alpha1.AwsCustomerGateway) (result *v1alpha1.AwsCustomerGateway, err error) {
	result = &v1alpha1.AwsCustomerGateway{}
	err = c.client.Put().
		Resource("awscustomergateways").
		Name(awsCustomerGateway.Name).
		SubResource("status").
		Body(awsCustomerGateway).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsCustomerGateway and deletes it. Returns an error if one occurs.
func (c *awsCustomerGateways) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awscustomergateways").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsCustomerGateways) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awscustomergateways").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsCustomerGateway.
func (c *awsCustomerGateways) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCustomerGateway, err error) {
	result = &v1alpha1.AwsCustomerGateway{}
	err = c.client.Patch(pt).
		Resource("awscustomergateways").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}