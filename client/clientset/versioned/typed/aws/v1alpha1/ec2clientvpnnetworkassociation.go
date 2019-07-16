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

// Ec2ClientVpnNetworkAssociationsGetter has a method to return a Ec2ClientVpnNetworkAssociationInterface.
// A group's client should implement this interface.
type Ec2ClientVpnNetworkAssociationsGetter interface {
	Ec2ClientVpnNetworkAssociations() Ec2ClientVpnNetworkAssociationInterface
}

// Ec2ClientVpnNetworkAssociationInterface has methods to work with Ec2ClientVpnNetworkAssociation resources.
type Ec2ClientVpnNetworkAssociationInterface interface {
	Create(*v1alpha1.Ec2ClientVpnNetworkAssociation) (*v1alpha1.Ec2ClientVpnNetworkAssociation, error)
	Update(*v1alpha1.Ec2ClientVpnNetworkAssociation) (*v1alpha1.Ec2ClientVpnNetworkAssociation, error)
	UpdateStatus(*v1alpha1.Ec2ClientVpnNetworkAssociation) (*v1alpha1.Ec2ClientVpnNetworkAssociation, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Ec2ClientVpnNetworkAssociation, error)
	List(opts v1.ListOptions) (*v1alpha1.Ec2ClientVpnNetworkAssociationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Ec2ClientVpnNetworkAssociation, err error)
	Ec2ClientVpnNetworkAssociationExpansion
}

// ec2ClientVpnNetworkAssociations implements Ec2ClientVpnNetworkAssociationInterface
type ec2ClientVpnNetworkAssociations struct {
	client rest.Interface
}

// newEc2ClientVpnNetworkAssociations returns a Ec2ClientVpnNetworkAssociations
func newEc2ClientVpnNetworkAssociations(c *AwsV1alpha1Client) *ec2ClientVpnNetworkAssociations {
	return &ec2ClientVpnNetworkAssociations{
		client: c.RESTClient(),
	}
}

// Get takes name of the ec2ClientVpnNetworkAssociation, and returns the corresponding ec2ClientVpnNetworkAssociation object, and an error if there is any.
func (c *ec2ClientVpnNetworkAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.Ec2ClientVpnNetworkAssociation, err error) {
	result = &v1alpha1.Ec2ClientVpnNetworkAssociation{}
	err = c.client.Get().
		Resource("ec2clientvpnnetworkassociations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Ec2ClientVpnNetworkAssociations that match those selectors.
func (c *ec2ClientVpnNetworkAssociations) List(opts v1.ListOptions) (result *v1alpha1.Ec2ClientVpnNetworkAssociationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.Ec2ClientVpnNetworkAssociationList{}
	err = c.client.Get().
		Resource("ec2clientvpnnetworkassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ec2ClientVpnNetworkAssociations.
func (c *ec2ClientVpnNetworkAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("ec2clientvpnnetworkassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a ec2ClientVpnNetworkAssociation and creates it.  Returns the server's representation of the ec2ClientVpnNetworkAssociation, and an error, if there is any.
func (c *ec2ClientVpnNetworkAssociations) Create(ec2ClientVpnNetworkAssociation *v1alpha1.Ec2ClientVpnNetworkAssociation) (result *v1alpha1.Ec2ClientVpnNetworkAssociation, err error) {
	result = &v1alpha1.Ec2ClientVpnNetworkAssociation{}
	err = c.client.Post().
		Resource("ec2clientvpnnetworkassociations").
		Body(ec2ClientVpnNetworkAssociation).
		Do().
		Into(result)
	return
}

// Update takes the representation of a ec2ClientVpnNetworkAssociation and updates it. Returns the server's representation of the ec2ClientVpnNetworkAssociation, and an error, if there is any.
func (c *ec2ClientVpnNetworkAssociations) Update(ec2ClientVpnNetworkAssociation *v1alpha1.Ec2ClientVpnNetworkAssociation) (result *v1alpha1.Ec2ClientVpnNetworkAssociation, err error) {
	result = &v1alpha1.Ec2ClientVpnNetworkAssociation{}
	err = c.client.Put().
		Resource("ec2clientvpnnetworkassociations").
		Name(ec2ClientVpnNetworkAssociation.Name).
		Body(ec2ClientVpnNetworkAssociation).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *ec2ClientVpnNetworkAssociations) UpdateStatus(ec2ClientVpnNetworkAssociation *v1alpha1.Ec2ClientVpnNetworkAssociation) (result *v1alpha1.Ec2ClientVpnNetworkAssociation, err error) {
	result = &v1alpha1.Ec2ClientVpnNetworkAssociation{}
	err = c.client.Put().
		Resource("ec2clientvpnnetworkassociations").
		Name(ec2ClientVpnNetworkAssociation.Name).
		SubResource("status").
		Body(ec2ClientVpnNetworkAssociation).
		Do().
		Into(result)
	return
}

// Delete takes name of the ec2ClientVpnNetworkAssociation and deletes it. Returns an error if one occurs.
func (c *ec2ClientVpnNetworkAssociations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("ec2clientvpnnetworkassociations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ec2ClientVpnNetworkAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("ec2clientvpnnetworkassociations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched ec2ClientVpnNetworkAssociation.
func (c *ec2ClientVpnNetworkAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Ec2ClientVpnNetworkAssociation, err error) {
	result = &v1alpha1.Ec2ClientVpnNetworkAssociation{}
	err = c.client.Patch(pt).
		Resource("ec2clientvpnnetworkassociations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}