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

// AwsWorklinkWebsiteCertificateAuthorityAssociationsGetter has a method to return a AwsWorklinkWebsiteCertificateAuthorityAssociationInterface.
// A group's client should implement this interface.
type AwsWorklinkWebsiteCertificateAuthorityAssociationsGetter interface {
	AwsWorklinkWebsiteCertificateAuthorityAssociations() AwsWorklinkWebsiteCertificateAuthorityAssociationInterface
}

// AwsWorklinkWebsiteCertificateAuthorityAssociationInterface has methods to work with AwsWorklinkWebsiteCertificateAuthorityAssociation resources.
type AwsWorklinkWebsiteCertificateAuthorityAssociationInterface interface {
	Create(*v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation) (*v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation, error)
	Update(*v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation) (*v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation, error)
	UpdateStatus(*v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation) (*v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation, err error)
	AwsWorklinkWebsiteCertificateAuthorityAssociationExpansion
}

// awsWorklinkWebsiteCertificateAuthorityAssociations implements AwsWorklinkWebsiteCertificateAuthorityAssociationInterface
type awsWorklinkWebsiteCertificateAuthorityAssociations struct {
	client rest.Interface
}

// newAwsWorklinkWebsiteCertificateAuthorityAssociations returns a AwsWorklinkWebsiteCertificateAuthorityAssociations
func newAwsWorklinkWebsiteCertificateAuthorityAssociations(c *AwsV1alpha1Client) *awsWorklinkWebsiteCertificateAuthorityAssociations {
	return &awsWorklinkWebsiteCertificateAuthorityAssociations{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsWorklinkWebsiteCertificateAuthorityAssociation, and returns the corresponding awsWorklinkWebsiteCertificateAuthorityAssociation object, and an error if there is any.
func (c *awsWorklinkWebsiteCertificateAuthorityAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation, err error) {
	result = &v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation{}
	err = c.client.Get().
		Resource("awsworklinkwebsitecertificateauthorityassociations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsWorklinkWebsiteCertificateAuthorityAssociations that match those selectors.
func (c *awsWorklinkWebsiteCertificateAuthorityAssociations) List(opts v1.ListOptions) (result *v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociationList{}
	err = c.client.Get().
		Resource("awsworklinkwebsitecertificateauthorityassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsWorklinkWebsiteCertificateAuthorityAssociations.
func (c *awsWorklinkWebsiteCertificateAuthorityAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsworklinkwebsitecertificateauthorityassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsWorklinkWebsiteCertificateAuthorityAssociation and creates it.  Returns the server's representation of the awsWorklinkWebsiteCertificateAuthorityAssociation, and an error, if there is any.
func (c *awsWorklinkWebsiteCertificateAuthorityAssociations) Create(awsWorklinkWebsiteCertificateAuthorityAssociation *v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation) (result *v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation, err error) {
	result = &v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation{}
	err = c.client.Post().
		Resource("awsworklinkwebsitecertificateauthorityassociations").
		Body(awsWorklinkWebsiteCertificateAuthorityAssociation).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsWorklinkWebsiteCertificateAuthorityAssociation and updates it. Returns the server's representation of the awsWorklinkWebsiteCertificateAuthorityAssociation, and an error, if there is any.
func (c *awsWorklinkWebsiteCertificateAuthorityAssociations) Update(awsWorklinkWebsiteCertificateAuthorityAssociation *v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation) (result *v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation, err error) {
	result = &v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation{}
	err = c.client.Put().
		Resource("awsworklinkwebsitecertificateauthorityassociations").
		Name(awsWorklinkWebsiteCertificateAuthorityAssociation.Name).
		Body(awsWorklinkWebsiteCertificateAuthorityAssociation).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsWorklinkWebsiteCertificateAuthorityAssociations) UpdateStatus(awsWorklinkWebsiteCertificateAuthorityAssociation *v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation) (result *v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation, err error) {
	result = &v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation{}
	err = c.client.Put().
		Resource("awsworklinkwebsitecertificateauthorityassociations").
		Name(awsWorklinkWebsiteCertificateAuthorityAssociation.Name).
		SubResource("status").
		Body(awsWorklinkWebsiteCertificateAuthorityAssociation).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsWorklinkWebsiteCertificateAuthorityAssociation and deletes it. Returns an error if one occurs.
func (c *awsWorklinkWebsiteCertificateAuthorityAssociations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsworklinkwebsitecertificateauthorityassociations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsWorklinkWebsiteCertificateAuthorityAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsworklinkwebsitecertificateauthorityassociations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsWorklinkWebsiteCertificateAuthorityAssociation.
func (c *awsWorklinkWebsiteCertificateAuthorityAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation, err error) {
	result = &v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation{}
	err = c.client.Patch(pt).
		Resource("awsworklinkwebsitecertificateauthorityassociations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}