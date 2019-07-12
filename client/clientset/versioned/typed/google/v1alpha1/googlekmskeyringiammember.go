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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// GoogleKmsKeyRingIamMembersGetter has a method to return a GoogleKmsKeyRingIamMemberInterface.
// A group's client should implement this interface.
type GoogleKmsKeyRingIamMembersGetter interface {
	GoogleKmsKeyRingIamMembers() GoogleKmsKeyRingIamMemberInterface
}

// GoogleKmsKeyRingIamMemberInterface has methods to work with GoogleKmsKeyRingIamMember resources.
type GoogleKmsKeyRingIamMemberInterface interface {
	Create(*v1alpha1.GoogleKmsKeyRingIamMember) (*v1alpha1.GoogleKmsKeyRingIamMember, error)
	Update(*v1alpha1.GoogleKmsKeyRingIamMember) (*v1alpha1.GoogleKmsKeyRingIamMember, error)
	UpdateStatus(*v1alpha1.GoogleKmsKeyRingIamMember) (*v1alpha1.GoogleKmsKeyRingIamMember, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleKmsKeyRingIamMember, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleKmsKeyRingIamMemberList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleKmsKeyRingIamMember, err error)
	GoogleKmsKeyRingIamMemberExpansion
}

// googleKmsKeyRingIamMembers implements GoogleKmsKeyRingIamMemberInterface
type googleKmsKeyRingIamMembers struct {
	client rest.Interface
}

// newGoogleKmsKeyRingIamMembers returns a GoogleKmsKeyRingIamMembers
func newGoogleKmsKeyRingIamMembers(c *GoogleV1alpha1Client) *googleKmsKeyRingIamMembers {
	return &googleKmsKeyRingIamMembers{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleKmsKeyRingIamMember, and returns the corresponding googleKmsKeyRingIamMember object, and an error if there is any.
func (c *googleKmsKeyRingIamMembers) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleKmsKeyRingIamMember, err error) {
	result = &v1alpha1.GoogleKmsKeyRingIamMember{}
	err = c.client.Get().
		Resource("googlekmskeyringiammembers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleKmsKeyRingIamMembers that match those selectors.
func (c *googleKmsKeyRingIamMembers) List(opts v1.ListOptions) (result *v1alpha1.GoogleKmsKeyRingIamMemberList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleKmsKeyRingIamMemberList{}
	err = c.client.Get().
		Resource("googlekmskeyringiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleKmsKeyRingIamMembers.
func (c *googleKmsKeyRingIamMembers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googlekmskeyringiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleKmsKeyRingIamMember and creates it.  Returns the server's representation of the googleKmsKeyRingIamMember, and an error, if there is any.
func (c *googleKmsKeyRingIamMembers) Create(googleKmsKeyRingIamMember *v1alpha1.GoogleKmsKeyRingIamMember) (result *v1alpha1.GoogleKmsKeyRingIamMember, err error) {
	result = &v1alpha1.GoogleKmsKeyRingIamMember{}
	err = c.client.Post().
		Resource("googlekmskeyringiammembers").
		Body(googleKmsKeyRingIamMember).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleKmsKeyRingIamMember and updates it. Returns the server's representation of the googleKmsKeyRingIamMember, and an error, if there is any.
func (c *googleKmsKeyRingIamMembers) Update(googleKmsKeyRingIamMember *v1alpha1.GoogleKmsKeyRingIamMember) (result *v1alpha1.GoogleKmsKeyRingIamMember, err error) {
	result = &v1alpha1.GoogleKmsKeyRingIamMember{}
	err = c.client.Put().
		Resource("googlekmskeyringiammembers").
		Name(googleKmsKeyRingIamMember.Name).
		Body(googleKmsKeyRingIamMember).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleKmsKeyRingIamMembers) UpdateStatus(googleKmsKeyRingIamMember *v1alpha1.GoogleKmsKeyRingIamMember) (result *v1alpha1.GoogleKmsKeyRingIamMember, err error) {
	result = &v1alpha1.GoogleKmsKeyRingIamMember{}
	err = c.client.Put().
		Resource("googlekmskeyringiammembers").
		Name(googleKmsKeyRingIamMember.Name).
		SubResource("status").
		Body(googleKmsKeyRingIamMember).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleKmsKeyRingIamMember and deletes it. Returns an error if one occurs.
func (c *googleKmsKeyRingIamMembers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googlekmskeyringiammembers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleKmsKeyRingIamMembers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googlekmskeyringiammembers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleKmsKeyRingIamMember.
func (c *googleKmsKeyRingIamMembers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleKmsKeyRingIamMember, err error) {
	result = &v1alpha1.GoogleKmsKeyRingIamMember{}
	err = c.client.Patch(pt).
		Resource("googlekmskeyringiammembers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}