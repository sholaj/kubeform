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

// AwsKmsGrantsGetter has a method to return a AwsKmsGrantInterface.
// A group's client should implement this interface.
type AwsKmsGrantsGetter interface {
	AwsKmsGrants() AwsKmsGrantInterface
}

// AwsKmsGrantInterface has methods to work with AwsKmsGrant resources.
type AwsKmsGrantInterface interface {
	Create(*v1alpha1.AwsKmsGrant) (*v1alpha1.AwsKmsGrant, error)
	Update(*v1alpha1.AwsKmsGrant) (*v1alpha1.AwsKmsGrant, error)
	UpdateStatus(*v1alpha1.AwsKmsGrant) (*v1alpha1.AwsKmsGrant, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsKmsGrant, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsKmsGrantList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsKmsGrant, err error)
	AwsKmsGrantExpansion
}

// awsKmsGrants implements AwsKmsGrantInterface
type awsKmsGrants struct {
	client rest.Interface
}

// newAwsKmsGrants returns a AwsKmsGrants
func newAwsKmsGrants(c *AwsV1alpha1Client) *awsKmsGrants {
	return &awsKmsGrants{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsKmsGrant, and returns the corresponding awsKmsGrant object, and an error if there is any.
func (c *awsKmsGrants) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsKmsGrant, err error) {
	result = &v1alpha1.AwsKmsGrant{}
	err = c.client.Get().
		Resource("awskmsgrants").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsKmsGrants that match those selectors.
func (c *awsKmsGrants) List(opts v1.ListOptions) (result *v1alpha1.AwsKmsGrantList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsKmsGrantList{}
	err = c.client.Get().
		Resource("awskmsgrants").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsKmsGrants.
func (c *awsKmsGrants) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awskmsgrants").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsKmsGrant and creates it.  Returns the server's representation of the awsKmsGrant, and an error, if there is any.
func (c *awsKmsGrants) Create(awsKmsGrant *v1alpha1.AwsKmsGrant) (result *v1alpha1.AwsKmsGrant, err error) {
	result = &v1alpha1.AwsKmsGrant{}
	err = c.client.Post().
		Resource("awskmsgrants").
		Body(awsKmsGrant).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsKmsGrant and updates it. Returns the server's representation of the awsKmsGrant, and an error, if there is any.
func (c *awsKmsGrants) Update(awsKmsGrant *v1alpha1.AwsKmsGrant) (result *v1alpha1.AwsKmsGrant, err error) {
	result = &v1alpha1.AwsKmsGrant{}
	err = c.client.Put().
		Resource("awskmsgrants").
		Name(awsKmsGrant.Name).
		Body(awsKmsGrant).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsKmsGrants) UpdateStatus(awsKmsGrant *v1alpha1.AwsKmsGrant) (result *v1alpha1.AwsKmsGrant, err error) {
	result = &v1alpha1.AwsKmsGrant{}
	err = c.client.Put().
		Resource("awskmsgrants").
		Name(awsKmsGrant.Name).
		SubResource("status").
		Body(awsKmsGrant).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsKmsGrant and deletes it. Returns an error if one occurs.
func (c *awsKmsGrants) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awskmsgrants").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsKmsGrants) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awskmsgrants").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsKmsGrant.
func (c *awsKmsGrants) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsKmsGrant, err error) {
	result = &v1alpha1.AwsKmsGrant{}
	err = c.client.Patch(pt).
		Resource("awskmsgrants").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}