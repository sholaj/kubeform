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

// AwsCognitoIdentityProvidersGetter has a method to return a AwsCognitoIdentityProviderInterface.
// A group's client should implement this interface.
type AwsCognitoIdentityProvidersGetter interface {
	AwsCognitoIdentityProviders() AwsCognitoIdentityProviderInterface
}

// AwsCognitoIdentityProviderInterface has methods to work with AwsCognitoIdentityProvider resources.
type AwsCognitoIdentityProviderInterface interface {
	Create(*v1alpha1.AwsCognitoIdentityProvider) (*v1alpha1.AwsCognitoIdentityProvider, error)
	Update(*v1alpha1.AwsCognitoIdentityProvider) (*v1alpha1.AwsCognitoIdentityProvider, error)
	UpdateStatus(*v1alpha1.AwsCognitoIdentityProvider) (*v1alpha1.AwsCognitoIdentityProvider, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsCognitoIdentityProvider, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsCognitoIdentityProviderList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCognitoIdentityProvider, err error)
	AwsCognitoIdentityProviderExpansion
}

// awsCognitoIdentityProviders implements AwsCognitoIdentityProviderInterface
type awsCognitoIdentityProviders struct {
	client rest.Interface
}

// newAwsCognitoIdentityProviders returns a AwsCognitoIdentityProviders
func newAwsCognitoIdentityProviders(c *AwsV1alpha1Client) *awsCognitoIdentityProviders {
	return &awsCognitoIdentityProviders{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsCognitoIdentityProvider, and returns the corresponding awsCognitoIdentityProvider object, and an error if there is any.
func (c *awsCognitoIdentityProviders) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsCognitoIdentityProvider, err error) {
	result = &v1alpha1.AwsCognitoIdentityProvider{}
	err = c.client.Get().
		Resource("awscognitoidentityproviders").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsCognitoIdentityProviders that match those selectors.
func (c *awsCognitoIdentityProviders) List(opts v1.ListOptions) (result *v1alpha1.AwsCognitoIdentityProviderList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsCognitoIdentityProviderList{}
	err = c.client.Get().
		Resource("awscognitoidentityproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsCognitoIdentityProviders.
func (c *awsCognitoIdentityProviders) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awscognitoidentityproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsCognitoIdentityProvider and creates it.  Returns the server's representation of the awsCognitoIdentityProvider, and an error, if there is any.
func (c *awsCognitoIdentityProviders) Create(awsCognitoIdentityProvider *v1alpha1.AwsCognitoIdentityProvider) (result *v1alpha1.AwsCognitoIdentityProvider, err error) {
	result = &v1alpha1.AwsCognitoIdentityProvider{}
	err = c.client.Post().
		Resource("awscognitoidentityproviders").
		Body(awsCognitoIdentityProvider).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsCognitoIdentityProvider and updates it. Returns the server's representation of the awsCognitoIdentityProvider, and an error, if there is any.
func (c *awsCognitoIdentityProviders) Update(awsCognitoIdentityProvider *v1alpha1.AwsCognitoIdentityProvider) (result *v1alpha1.AwsCognitoIdentityProvider, err error) {
	result = &v1alpha1.AwsCognitoIdentityProvider{}
	err = c.client.Put().
		Resource("awscognitoidentityproviders").
		Name(awsCognitoIdentityProvider.Name).
		Body(awsCognitoIdentityProvider).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsCognitoIdentityProviders) UpdateStatus(awsCognitoIdentityProvider *v1alpha1.AwsCognitoIdentityProvider) (result *v1alpha1.AwsCognitoIdentityProvider, err error) {
	result = &v1alpha1.AwsCognitoIdentityProvider{}
	err = c.client.Put().
		Resource("awscognitoidentityproviders").
		Name(awsCognitoIdentityProvider.Name).
		SubResource("status").
		Body(awsCognitoIdentityProvider).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsCognitoIdentityProvider and deletes it. Returns an error if one occurs.
func (c *awsCognitoIdentityProviders) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awscognitoidentityproviders").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsCognitoIdentityProviders) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awscognitoidentityproviders").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsCognitoIdentityProvider.
func (c *awsCognitoIdentityProviders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCognitoIdentityProvider, err error) {
	result = &v1alpha1.AwsCognitoIdentityProvider{}
	err = c.client.Patch(pt).
		Resource("awscognitoidentityproviders").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}