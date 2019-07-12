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

// AwsWafByteMatchSetsGetter has a method to return a AwsWafByteMatchSetInterface.
// A group's client should implement this interface.
type AwsWafByteMatchSetsGetter interface {
	AwsWafByteMatchSets() AwsWafByteMatchSetInterface
}

// AwsWafByteMatchSetInterface has methods to work with AwsWafByteMatchSet resources.
type AwsWafByteMatchSetInterface interface {
	Create(*v1alpha1.AwsWafByteMatchSet) (*v1alpha1.AwsWafByteMatchSet, error)
	Update(*v1alpha1.AwsWafByteMatchSet) (*v1alpha1.AwsWafByteMatchSet, error)
	UpdateStatus(*v1alpha1.AwsWafByteMatchSet) (*v1alpha1.AwsWafByteMatchSet, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsWafByteMatchSet, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsWafByteMatchSetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafByteMatchSet, err error)
	AwsWafByteMatchSetExpansion
}

// awsWafByteMatchSets implements AwsWafByteMatchSetInterface
type awsWafByteMatchSets struct {
	client rest.Interface
}

// newAwsWafByteMatchSets returns a AwsWafByteMatchSets
func newAwsWafByteMatchSets(c *AwsV1alpha1Client) *awsWafByteMatchSets {
	return &awsWafByteMatchSets{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsWafByteMatchSet, and returns the corresponding awsWafByteMatchSet object, and an error if there is any.
func (c *awsWafByteMatchSets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsWafByteMatchSet, err error) {
	result = &v1alpha1.AwsWafByteMatchSet{}
	err = c.client.Get().
		Resource("awswafbytematchsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsWafByteMatchSets that match those selectors.
func (c *awsWafByteMatchSets) List(opts v1.ListOptions) (result *v1alpha1.AwsWafByteMatchSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsWafByteMatchSetList{}
	err = c.client.Get().
		Resource("awswafbytematchsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsWafByteMatchSets.
func (c *awsWafByteMatchSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awswafbytematchsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsWafByteMatchSet and creates it.  Returns the server's representation of the awsWafByteMatchSet, and an error, if there is any.
func (c *awsWafByteMatchSets) Create(awsWafByteMatchSet *v1alpha1.AwsWafByteMatchSet) (result *v1alpha1.AwsWafByteMatchSet, err error) {
	result = &v1alpha1.AwsWafByteMatchSet{}
	err = c.client.Post().
		Resource("awswafbytematchsets").
		Body(awsWafByteMatchSet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsWafByteMatchSet and updates it. Returns the server's representation of the awsWafByteMatchSet, and an error, if there is any.
func (c *awsWafByteMatchSets) Update(awsWafByteMatchSet *v1alpha1.AwsWafByteMatchSet) (result *v1alpha1.AwsWafByteMatchSet, err error) {
	result = &v1alpha1.AwsWafByteMatchSet{}
	err = c.client.Put().
		Resource("awswafbytematchsets").
		Name(awsWafByteMatchSet.Name).
		Body(awsWafByteMatchSet).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsWafByteMatchSets) UpdateStatus(awsWafByteMatchSet *v1alpha1.AwsWafByteMatchSet) (result *v1alpha1.AwsWafByteMatchSet, err error) {
	result = &v1alpha1.AwsWafByteMatchSet{}
	err = c.client.Put().
		Resource("awswafbytematchsets").
		Name(awsWafByteMatchSet.Name).
		SubResource("status").
		Body(awsWafByteMatchSet).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsWafByteMatchSet and deletes it. Returns an error if one occurs.
func (c *awsWafByteMatchSets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awswafbytematchsets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsWafByteMatchSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awswafbytematchsets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsWafByteMatchSet.
func (c *awsWafByteMatchSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafByteMatchSet, err error) {
	result = &v1alpha1.AwsWafByteMatchSet{}
	err = c.client.Patch(pt).
		Resource("awswafbytematchsets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}