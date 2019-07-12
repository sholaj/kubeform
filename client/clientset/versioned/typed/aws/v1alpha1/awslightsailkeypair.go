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

// AwsLightsailKeyPairsGetter has a method to return a AwsLightsailKeyPairInterface.
// A group's client should implement this interface.
type AwsLightsailKeyPairsGetter interface {
	AwsLightsailKeyPairs() AwsLightsailKeyPairInterface
}

// AwsLightsailKeyPairInterface has methods to work with AwsLightsailKeyPair resources.
type AwsLightsailKeyPairInterface interface {
	Create(*v1alpha1.AwsLightsailKeyPair) (*v1alpha1.AwsLightsailKeyPair, error)
	Update(*v1alpha1.AwsLightsailKeyPair) (*v1alpha1.AwsLightsailKeyPair, error)
	UpdateStatus(*v1alpha1.AwsLightsailKeyPair) (*v1alpha1.AwsLightsailKeyPair, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsLightsailKeyPair, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsLightsailKeyPairList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsLightsailKeyPair, err error)
	AwsLightsailKeyPairExpansion
}

// awsLightsailKeyPairs implements AwsLightsailKeyPairInterface
type awsLightsailKeyPairs struct {
	client rest.Interface
}

// newAwsLightsailKeyPairs returns a AwsLightsailKeyPairs
func newAwsLightsailKeyPairs(c *AwsV1alpha1Client) *awsLightsailKeyPairs {
	return &awsLightsailKeyPairs{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsLightsailKeyPair, and returns the corresponding awsLightsailKeyPair object, and an error if there is any.
func (c *awsLightsailKeyPairs) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsLightsailKeyPair, err error) {
	result = &v1alpha1.AwsLightsailKeyPair{}
	err = c.client.Get().
		Resource("awslightsailkeypairs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsLightsailKeyPairs that match those selectors.
func (c *awsLightsailKeyPairs) List(opts v1.ListOptions) (result *v1alpha1.AwsLightsailKeyPairList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsLightsailKeyPairList{}
	err = c.client.Get().
		Resource("awslightsailkeypairs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsLightsailKeyPairs.
func (c *awsLightsailKeyPairs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awslightsailkeypairs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsLightsailKeyPair and creates it.  Returns the server's representation of the awsLightsailKeyPair, and an error, if there is any.
func (c *awsLightsailKeyPairs) Create(awsLightsailKeyPair *v1alpha1.AwsLightsailKeyPair) (result *v1alpha1.AwsLightsailKeyPair, err error) {
	result = &v1alpha1.AwsLightsailKeyPair{}
	err = c.client.Post().
		Resource("awslightsailkeypairs").
		Body(awsLightsailKeyPair).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsLightsailKeyPair and updates it. Returns the server's representation of the awsLightsailKeyPair, and an error, if there is any.
func (c *awsLightsailKeyPairs) Update(awsLightsailKeyPair *v1alpha1.AwsLightsailKeyPair) (result *v1alpha1.AwsLightsailKeyPair, err error) {
	result = &v1alpha1.AwsLightsailKeyPair{}
	err = c.client.Put().
		Resource("awslightsailkeypairs").
		Name(awsLightsailKeyPair.Name).
		Body(awsLightsailKeyPair).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsLightsailKeyPairs) UpdateStatus(awsLightsailKeyPair *v1alpha1.AwsLightsailKeyPair) (result *v1alpha1.AwsLightsailKeyPair, err error) {
	result = &v1alpha1.AwsLightsailKeyPair{}
	err = c.client.Put().
		Resource("awslightsailkeypairs").
		Name(awsLightsailKeyPair.Name).
		SubResource("status").
		Body(awsLightsailKeyPair).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsLightsailKeyPair and deletes it. Returns an error if one occurs.
func (c *awsLightsailKeyPairs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awslightsailkeypairs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsLightsailKeyPairs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awslightsailkeypairs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsLightsailKeyPair.
func (c *awsLightsailKeyPairs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsLightsailKeyPair, err error) {
	result = &v1alpha1.AwsLightsailKeyPair{}
	err = c.client.Patch(pt).
		Resource("awslightsailkeypairs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}