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

// AwsTransferUsersGetter has a method to return a AwsTransferUserInterface.
// A group's client should implement this interface.
type AwsTransferUsersGetter interface {
	AwsTransferUsers() AwsTransferUserInterface
}

// AwsTransferUserInterface has methods to work with AwsTransferUser resources.
type AwsTransferUserInterface interface {
	Create(*v1alpha1.AwsTransferUser) (*v1alpha1.AwsTransferUser, error)
	Update(*v1alpha1.AwsTransferUser) (*v1alpha1.AwsTransferUser, error)
	UpdateStatus(*v1alpha1.AwsTransferUser) (*v1alpha1.AwsTransferUser, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsTransferUser, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsTransferUserList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsTransferUser, err error)
	AwsTransferUserExpansion
}

// awsTransferUsers implements AwsTransferUserInterface
type awsTransferUsers struct {
	client rest.Interface
}

// newAwsTransferUsers returns a AwsTransferUsers
func newAwsTransferUsers(c *AwsV1alpha1Client) *awsTransferUsers {
	return &awsTransferUsers{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsTransferUser, and returns the corresponding awsTransferUser object, and an error if there is any.
func (c *awsTransferUsers) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsTransferUser, err error) {
	result = &v1alpha1.AwsTransferUser{}
	err = c.client.Get().
		Resource("awstransferusers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsTransferUsers that match those selectors.
func (c *awsTransferUsers) List(opts v1.ListOptions) (result *v1alpha1.AwsTransferUserList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsTransferUserList{}
	err = c.client.Get().
		Resource("awstransferusers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsTransferUsers.
func (c *awsTransferUsers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awstransferusers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsTransferUser and creates it.  Returns the server's representation of the awsTransferUser, and an error, if there is any.
func (c *awsTransferUsers) Create(awsTransferUser *v1alpha1.AwsTransferUser) (result *v1alpha1.AwsTransferUser, err error) {
	result = &v1alpha1.AwsTransferUser{}
	err = c.client.Post().
		Resource("awstransferusers").
		Body(awsTransferUser).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsTransferUser and updates it. Returns the server's representation of the awsTransferUser, and an error, if there is any.
func (c *awsTransferUsers) Update(awsTransferUser *v1alpha1.AwsTransferUser) (result *v1alpha1.AwsTransferUser, err error) {
	result = &v1alpha1.AwsTransferUser{}
	err = c.client.Put().
		Resource("awstransferusers").
		Name(awsTransferUser.Name).
		Body(awsTransferUser).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsTransferUsers) UpdateStatus(awsTransferUser *v1alpha1.AwsTransferUser) (result *v1alpha1.AwsTransferUser, err error) {
	result = &v1alpha1.AwsTransferUser{}
	err = c.client.Put().
		Resource("awstransferusers").
		Name(awsTransferUser.Name).
		SubResource("status").
		Body(awsTransferUser).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsTransferUser and deletes it. Returns an error if one occurs.
func (c *awsTransferUsers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awstransferusers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsTransferUsers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awstransferusers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsTransferUser.
func (c *awsTransferUsers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsTransferUser, err error) {
	result = &v1alpha1.AwsTransferUser{}
	err = c.client.Patch(pt).
		Resource("awstransferusers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}