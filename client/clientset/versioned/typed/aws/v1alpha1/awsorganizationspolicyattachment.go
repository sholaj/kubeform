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

// AwsOrganizationsPolicyAttachmentsGetter has a method to return a AwsOrganizationsPolicyAttachmentInterface.
// A group's client should implement this interface.
type AwsOrganizationsPolicyAttachmentsGetter interface {
	AwsOrganizationsPolicyAttachments() AwsOrganizationsPolicyAttachmentInterface
}

// AwsOrganizationsPolicyAttachmentInterface has methods to work with AwsOrganizationsPolicyAttachment resources.
type AwsOrganizationsPolicyAttachmentInterface interface {
	Create(*v1alpha1.AwsOrganizationsPolicyAttachment) (*v1alpha1.AwsOrganizationsPolicyAttachment, error)
	Update(*v1alpha1.AwsOrganizationsPolicyAttachment) (*v1alpha1.AwsOrganizationsPolicyAttachment, error)
	UpdateStatus(*v1alpha1.AwsOrganizationsPolicyAttachment) (*v1alpha1.AwsOrganizationsPolicyAttachment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsOrganizationsPolicyAttachment, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsOrganizationsPolicyAttachmentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsOrganizationsPolicyAttachment, err error)
	AwsOrganizationsPolicyAttachmentExpansion
}

// awsOrganizationsPolicyAttachments implements AwsOrganizationsPolicyAttachmentInterface
type awsOrganizationsPolicyAttachments struct {
	client rest.Interface
}

// newAwsOrganizationsPolicyAttachments returns a AwsOrganizationsPolicyAttachments
func newAwsOrganizationsPolicyAttachments(c *AwsV1alpha1Client) *awsOrganizationsPolicyAttachments {
	return &awsOrganizationsPolicyAttachments{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsOrganizationsPolicyAttachment, and returns the corresponding awsOrganizationsPolicyAttachment object, and an error if there is any.
func (c *awsOrganizationsPolicyAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsOrganizationsPolicyAttachment, err error) {
	result = &v1alpha1.AwsOrganizationsPolicyAttachment{}
	err = c.client.Get().
		Resource("awsorganizationspolicyattachments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsOrganizationsPolicyAttachments that match those selectors.
func (c *awsOrganizationsPolicyAttachments) List(opts v1.ListOptions) (result *v1alpha1.AwsOrganizationsPolicyAttachmentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsOrganizationsPolicyAttachmentList{}
	err = c.client.Get().
		Resource("awsorganizationspolicyattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsOrganizationsPolicyAttachments.
func (c *awsOrganizationsPolicyAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsorganizationspolicyattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsOrganizationsPolicyAttachment and creates it.  Returns the server's representation of the awsOrganizationsPolicyAttachment, and an error, if there is any.
func (c *awsOrganizationsPolicyAttachments) Create(awsOrganizationsPolicyAttachment *v1alpha1.AwsOrganizationsPolicyAttachment) (result *v1alpha1.AwsOrganizationsPolicyAttachment, err error) {
	result = &v1alpha1.AwsOrganizationsPolicyAttachment{}
	err = c.client.Post().
		Resource("awsorganizationspolicyattachments").
		Body(awsOrganizationsPolicyAttachment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsOrganizationsPolicyAttachment and updates it. Returns the server's representation of the awsOrganizationsPolicyAttachment, and an error, if there is any.
func (c *awsOrganizationsPolicyAttachments) Update(awsOrganizationsPolicyAttachment *v1alpha1.AwsOrganizationsPolicyAttachment) (result *v1alpha1.AwsOrganizationsPolicyAttachment, err error) {
	result = &v1alpha1.AwsOrganizationsPolicyAttachment{}
	err = c.client.Put().
		Resource("awsorganizationspolicyattachments").
		Name(awsOrganizationsPolicyAttachment.Name).
		Body(awsOrganizationsPolicyAttachment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsOrganizationsPolicyAttachments) UpdateStatus(awsOrganizationsPolicyAttachment *v1alpha1.AwsOrganizationsPolicyAttachment) (result *v1alpha1.AwsOrganizationsPolicyAttachment, err error) {
	result = &v1alpha1.AwsOrganizationsPolicyAttachment{}
	err = c.client.Put().
		Resource("awsorganizationspolicyattachments").
		Name(awsOrganizationsPolicyAttachment.Name).
		SubResource("status").
		Body(awsOrganizationsPolicyAttachment).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsOrganizationsPolicyAttachment and deletes it. Returns an error if one occurs.
func (c *awsOrganizationsPolicyAttachments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsorganizationspolicyattachments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsOrganizationsPolicyAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsorganizationspolicyattachments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsOrganizationsPolicyAttachment.
func (c *awsOrganizationsPolicyAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsOrganizationsPolicyAttachment, err error) {
	result = &v1alpha1.AwsOrganizationsPolicyAttachment{}
	err = c.client.Patch(pt).
		Resource("awsorganizationspolicyattachments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}