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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// AzurermPolicyAssignmentsGetter has a method to return a AzurermPolicyAssignmentInterface.
// A group's client should implement this interface.
type AzurermPolicyAssignmentsGetter interface {
	AzurermPolicyAssignments() AzurermPolicyAssignmentInterface
}

// AzurermPolicyAssignmentInterface has methods to work with AzurermPolicyAssignment resources.
type AzurermPolicyAssignmentInterface interface {
	Create(*v1alpha1.AzurermPolicyAssignment) (*v1alpha1.AzurermPolicyAssignment, error)
	Update(*v1alpha1.AzurermPolicyAssignment) (*v1alpha1.AzurermPolicyAssignment, error)
	UpdateStatus(*v1alpha1.AzurermPolicyAssignment) (*v1alpha1.AzurermPolicyAssignment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermPolicyAssignment, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermPolicyAssignmentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermPolicyAssignment, err error)
	AzurermPolicyAssignmentExpansion
}

// azurermPolicyAssignments implements AzurermPolicyAssignmentInterface
type azurermPolicyAssignments struct {
	client rest.Interface
}

// newAzurermPolicyAssignments returns a AzurermPolicyAssignments
func newAzurermPolicyAssignments(c *AzurermV1alpha1Client) *azurermPolicyAssignments {
	return &azurermPolicyAssignments{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermPolicyAssignment, and returns the corresponding azurermPolicyAssignment object, and an error if there is any.
func (c *azurermPolicyAssignments) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermPolicyAssignment, err error) {
	result = &v1alpha1.AzurermPolicyAssignment{}
	err = c.client.Get().
		Resource("azurermpolicyassignments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermPolicyAssignments that match those selectors.
func (c *azurermPolicyAssignments) List(opts v1.ListOptions) (result *v1alpha1.AzurermPolicyAssignmentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermPolicyAssignmentList{}
	err = c.client.Get().
		Resource("azurermpolicyassignments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermPolicyAssignments.
func (c *azurermPolicyAssignments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermpolicyassignments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermPolicyAssignment and creates it.  Returns the server's representation of the azurermPolicyAssignment, and an error, if there is any.
func (c *azurermPolicyAssignments) Create(azurermPolicyAssignment *v1alpha1.AzurermPolicyAssignment) (result *v1alpha1.AzurermPolicyAssignment, err error) {
	result = &v1alpha1.AzurermPolicyAssignment{}
	err = c.client.Post().
		Resource("azurermpolicyassignments").
		Body(azurermPolicyAssignment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermPolicyAssignment and updates it. Returns the server's representation of the azurermPolicyAssignment, and an error, if there is any.
func (c *azurermPolicyAssignments) Update(azurermPolicyAssignment *v1alpha1.AzurermPolicyAssignment) (result *v1alpha1.AzurermPolicyAssignment, err error) {
	result = &v1alpha1.AzurermPolicyAssignment{}
	err = c.client.Put().
		Resource("azurermpolicyassignments").
		Name(azurermPolicyAssignment.Name).
		Body(azurermPolicyAssignment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermPolicyAssignments) UpdateStatus(azurermPolicyAssignment *v1alpha1.AzurermPolicyAssignment) (result *v1alpha1.AzurermPolicyAssignment, err error) {
	result = &v1alpha1.AzurermPolicyAssignment{}
	err = c.client.Put().
		Resource("azurermpolicyassignments").
		Name(azurermPolicyAssignment.Name).
		SubResource("status").
		Body(azurermPolicyAssignment).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermPolicyAssignment and deletes it. Returns an error if one occurs.
func (c *azurermPolicyAssignments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermpolicyassignments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermPolicyAssignments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermpolicyassignments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermPolicyAssignment.
func (c *azurermPolicyAssignments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermPolicyAssignment, err error) {
	result = &v1alpha1.AzurermPolicyAssignment{}
	err = c.client.Patch(pt).
		Resource("azurermpolicyassignments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}