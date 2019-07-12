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

// AwsLbSslNegotiationPoliciesGetter has a method to return a AwsLbSslNegotiationPolicyInterface.
// A group's client should implement this interface.
type AwsLbSslNegotiationPoliciesGetter interface {
	AwsLbSslNegotiationPolicies() AwsLbSslNegotiationPolicyInterface
}

// AwsLbSslNegotiationPolicyInterface has methods to work with AwsLbSslNegotiationPolicy resources.
type AwsLbSslNegotiationPolicyInterface interface {
	Create(*v1alpha1.AwsLbSslNegotiationPolicy) (*v1alpha1.AwsLbSslNegotiationPolicy, error)
	Update(*v1alpha1.AwsLbSslNegotiationPolicy) (*v1alpha1.AwsLbSslNegotiationPolicy, error)
	UpdateStatus(*v1alpha1.AwsLbSslNegotiationPolicy) (*v1alpha1.AwsLbSslNegotiationPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsLbSslNegotiationPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsLbSslNegotiationPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsLbSslNegotiationPolicy, err error)
	AwsLbSslNegotiationPolicyExpansion
}

// awsLbSslNegotiationPolicies implements AwsLbSslNegotiationPolicyInterface
type awsLbSslNegotiationPolicies struct {
	client rest.Interface
}

// newAwsLbSslNegotiationPolicies returns a AwsLbSslNegotiationPolicies
func newAwsLbSslNegotiationPolicies(c *AwsV1alpha1Client) *awsLbSslNegotiationPolicies {
	return &awsLbSslNegotiationPolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsLbSslNegotiationPolicy, and returns the corresponding awsLbSslNegotiationPolicy object, and an error if there is any.
func (c *awsLbSslNegotiationPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsLbSslNegotiationPolicy, err error) {
	result = &v1alpha1.AwsLbSslNegotiationPolicy{}
	err = c.client.Get().
		Resource("awslbsslnegotiationpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsLbSslNegotiationPolicies that match those selectors.
func (c *awsLbSslNegotiationPolicies) List(opts v1.ListOptions) (result *v1alpha1.AwsLbSslNegotiationPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsLbSslNegotiationPolicyList{}
	err = c.client.Get().
		Resource("awslbsslnegotiationpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsLbSslNegotiationPolicies.
func (c *awsLbSslNegotiationPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awslbsslnegotiationpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsLbSslNegotiationPolicy and creates it.  Returns the server's representation of the awsLbSslNegotiationPolicy, and an error, if there is any.
func (c *awsLbSslNegotiationPolicies) Create(awsLbSslNegotiationPolicy *v1alpha1.AwsLbSslNegotiationPolicy) (result *v1alpha1.AwsLbSslNegotiationPolicy, err error) {
	result = &v1alpha1.AwsLbSslNegotiationPolicy{}
	err = c.client.Post().
		Resource("awslbsslnegotiationpolicies").
		Body(awsLbSslNegotiationPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsLbSslNegotiationPolicy and updates it. Returns the server's representation of the awsLbSslNegotiationPolicy, and an error, if there is any.
func (c *awsLbSslNegotiationPolicies) Update(awsLbSslNegotiationPolicy *v1alpha1.AwsLbSslNegotiationPolicy) (result *v1alpha1.AwsLbSslNegotiationPolicy, err error) {
	result = &v1alpha1.AwsLbSslNegotiationPolicy{}
	err = c.client.Put().
		Resource("awslbsslnegotiationpolicies").
		Name(awsLbSslNegotiationPolicy.Name).
		Body(awsLbSslNegotiationPolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsLbSslNegotiationPolicies) UpdateStatus(awsLbSslNegotiationPolicy *v1alpha1.AwsLbSslNegotiationPolicy) (result *v1alpha1.AwsLbSslNegotiationPolicy, err error) {
	result = &v1alpha1.AwsLbSslNegotiationPolicy{}
	err = c.client.Put().
		Resource("awslbsslnegotiationpolicies").
		Name(awsLbSslNegotiationPolicy.Name).
		SubResource("status").
		Body(awsLbSslNegotiationPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsLbSslNegotiationPolicy and deletes it. Returns an error if one occurs.
func (c *awsLbSslNegotiationPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awslbsslnegotiationpolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsLbSslNegotiationPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awslbsslnegotiationpolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsLbSslNegotiationPolicy.
func (c *awsLbSslNegotiationPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsLbSslNegotiationPolicy, err error) {
	result = &v1alpha1.AwsLbSslNegotiationPolicy{}
	err = c.client.Patch(pt).
		Resource("awslbsslnegotiationpolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}