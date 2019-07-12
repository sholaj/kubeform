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

// AwsEgressOnlyInternetGatewaysGetter has a method to return a AwsEgressOnlyInternetGatewayInterface.
// A group's client should implement this interface.
type AwsEgressOnlyInternetGatewaysGetter interface {
	AwsEgressOnlyInternetGateways() AwsEgressOnlyInternetGatewayInterface
}

// AwsEgressOnlyInternetGatewayInterface has methods to work with AwsEgressOnlyInternetGateway resources.
type AwsEgressOnlyInternetGatewayInterface interface {
	Create(*v1alpha1.AwsEgressOnlyInternetGateway) (*v1alpha1.AwsEgressOnlyInternetGateway, error)
	Update(*v1alpha1.AwsEgressOnlyInternetGateway) (*v1alpha1.AwsEgressOnlyInternetGateway, error)
	UpdateStatus(*v1alpha1.AwsEgressOnlyInternetGateway) (*v1alpha1.AwsEgressOnlyInternetGateway, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsEgressOnlyInternetGateway, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsEgressOnlyInternetGatewayList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsEgressOnlyInternetGateway, err error)
	AwsEgressOnlyInternetGatewayExpansion
}

// awsEgressOnlyInternetGateways implements AwsEgressOnlyInternetGatewayInterface
type awsEgressOnlyInternetGateways struct {
	client rest.Interface
}

// newAwsEgressOnlyInternetGateways returns a AwsEgressOnlyInternetGateways
func newAwsEgressOnlyInternetGateways(c *AwsV1alpha1Client) *awsEgressOnlyInternetGateways {
	return &awsEgressOnlyInternetGateways{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsEgressOnlyInternetGateway, and returns the corresponding awsEgressOnlyInternetGateway object, and an error if there is any.
func (c *awsEgressOnlyInternetGateways) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsEgressOnlyInternetGateway, err error) {
	result = &v1alpha1.AwsEgressOnlyInternetGateway{}
	err = c.client.Get().
		Resource("awsegressonlyinternetgateways").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsEgressOnlyInternetGateways that match those selectors.
func (c *awsEgressOnlyInternetGateways) List(opts v1.ListOptions) (result *v1alpha1.AwsEgressOnlyInternetGatewayList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsEgressOnlyInternetGatewayList{}
	err = c.client.Get().
		Resource("awsegressonlyinternetgateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsEgressOnlyInternetGateways.
func (c *awsEgressOnlyInternetGateways) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsegressonlyinternetgateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsEgressOnlyInternetGateway and creates it.  Returns the server's representation of the awsEgressOnlyInternetGateway, and an error, if there is any.
func (c *awsEgressOnlyInternetGateways) Create(awsEgressOnlyInternetGateway *v1alpha1.AwsEgressOnlyInternetGateway) (result *v1alpha1.AwsEgressOnlyInternetGateway, err error) {
	result = &v1alpha1.AwsEgressOnlyInternetGateway{}
	err = c.client.Post().
		Resource("awsegressonlyinternetgateways").
		Body(awsEgressOnlyInternetGateway).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsEgressOnlyInternetGateway and updates it. Returns the server's representation of the awsEgressOnlyInternetGateway, and an error, if there is any.
func (c *awsEgressOnlyInternetGateways) Update(awsEgressOnlyInternetGateway *v1alpha1.AwsEgressOnlyInternetGateway) (result *v1alpha1.AwsEgressOnlyInternetGateway, err error) {
	result = &v1alpha1.AwsEgressOnlyInternetGateway{}
	err = c.client.Put().
		Resource("awsegressonlyinternetgateways").
		Name(awsEgressOnlyInternetGateway.Name).
		Body(awsEgressOnlyInternetGateway).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsEgressOnlyInternetGateways) UpdateStatus(awsEgressOnlyInternetGateway *v1alpha1.AwsEgressOnlyInternetGateway) (result *v1alpha1.AwsEgressOnlyInternetGateway, err error) {
	result = &v1alpha1.AwsEgressOnlyInternetGateway{}
	err = c.client.Put().
		Resource("awsegressonlyinternetgateways").
		Name(awsEgressOnlyInternetGateway.Name).
		SubResource("status").
		Body(awsEgressOnlyInternetGateway).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsEgressOnlyInternetGateway and deletes it. Returns an error if one occurs.
func (c *awsEgressOnlyInternetGateways) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsegressonlyinternetgateways").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsEgressOnlyInternetGateways) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsegressonlyinternetgateways").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsEgressOnlyInternetGateway.
func (c *awsEgressOnlyInternetGateways) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsEgressOnlyInternetGateway, err error) {
	result = &v1alpha1.AwsEgressOnlyInternetGateway{}
	err = c.client.Patch(pt).
		Resource("awsegressonlyinternetgateways").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}