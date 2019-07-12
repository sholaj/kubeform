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

// GoogleKmsKeyRingIamBindingsGetter has a method to return a GoogleKmsKeyRingIamBindingInterface.
// A group's client should implement this interface.
type GoogleKmsKeyRingIamBindingsGetter interface {
	GoogleKmsKeyRingIamBindings() GoogleKmsKeyRingIamBindingInterface
}

// GoogleKmsKeyRingIamBindingInterface has methods to work with GoogleKmsKeyRingIamBinding resources.
type GoogleKmsKeyRingIamBindingInterface interface {
	Create(*v1alpha1.GoogleKmsKeyRingIamBinding) (*v1alpha1.GoogleKmsKeyRingIamBinding, error)
	Update(*v1alpha1.GoogleKmsKeyRingIamBinding) (*v1alpha1.GoogleKmsKeyRingIamBinding, error)
	UpdateStatus(*v1alpha1.GoogleKmsKeyRingIamBinding) (*v1alpha1.GoogleKmsKeyRingIamBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleKmsKeyRingIamBinding, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleKmsKeyRingIamBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleKmsKeyRingIamBinding, err error)
	GoogleKmsKeyRingIamBindingExpansion
}

// googleKmsKeyRingIamBindings implements GoogleKmsKeyRingIamBindingInterface
type googleKmsKeyRingIamBindings struct {
	client rest.Interface
}

// newGoogleKmsKeyRingIamBindings returns a GoogleKmsKeyRingIamBindings
func newGoogleKmsKeyRingIamBindings(c *GoogleV1alpha1Client) *googleKmsKeyRingIamBindings {
	return &googleKmsKeyRingIamBindings{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleKmsKeyRingIamBinding, and returns the corresponding googleKmsKeyRingIamBinding object, and an error if there is any.
func (c *googleKmsKeyRingIamBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleKmsKeyRingIamBinding, err error) {
	result = &v1alpha1.GoogleKmsKeyRingIamBinding{}
	err = c.client.Get().
		Resource("googlekmskeyringiambindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleKmsKeyRingIamBindings that match those selectors.
func (c *googleKmsKeyRingIamBindings) List(opts v1.ListOptions) (result *v1alpha1.GoogleKmsKeyRingIamBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleKmsKeyRingIamBindingList{}
	err = c.client.Get().
		Resource("googlekmskeyringiambindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleKmsKeyRingIamBindings.
func (c *googleKmsKeyRingIamBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googlekmskeyringiambindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleKmsKeyRingIamBinding and creates it.  Returns the server's representation of the googleKmsKeyRingIamBinding, and an error, if there is any.
func (c *googleKmsKeyRingIamBindings) Create(googleKmsKeyRingIamBinding *v1alpha1.GoogleKmsKeyRingIamBinding) (result *v1alpha1.GoogleKmsKeyRingIamBinding, err error) {
	result = &v1alpha1.GoogleKmsKeyRingIamBinding{}
	err = c.client.Post().
		Resource("googlekmskeyringiambindings").
		Body(googleKmsKeyRingIamBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleKmsKeyRingIamBinding and updates it. Returns the server's representation of the googleKmsKeyRingIamBinding, and an error, if there is any.
func (c *googleKmsKeyRingIamBindings) Update(googleKmsKeyRingIamBinding *v1alpha1.GoogleKmsKeyRingIamBinding) (result *v1alpha1.GoogleKmsKeyRingIamBinding, err error) {
	result = &v1alpha1.GoogleKmsKeyRingIamBinding{}
	err = c.client.Put().
		Resource("googlekmskeyringiambindings").
		Name(googleKmsKeyRingIamBinding.Name).
		Body(googleKmsKeyRingIamBinding).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleKmsKeyRingIamBindings) UpdateStatus(googleKmsKeyRingIamBinding *v1alpha1.GoogleKmsKeyRingIamBinding) (result *v1alpha1.GoogleKmsKeyRingIamBinding, err error) {
	result = &v1alpha1.GoogleKmsKeyRingIamBinding{}
	err = c.client.Put().
		Resource("googlekmskeyringiambindings").
		Name(googleKmsKeyRingIamBinding.Name).
		SubResource("status").
		Body(googleKmsKeyRingIamBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleKmsKeyRingIamBinding and deletes it. Returns an error if one occurs.
func (c *googleKmsKeyRingIamBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googlekmskeyringiambindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleKmsKeyRingIamBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googlekmskeyringiambindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleKmsKeyRingIamBinding.
func (c *googleKmsKeyRingIamBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleKmsKeyRingIamBinding, err error) {
	result = &v1alpha1.GoogleKmsKeyRingIamBinding{}
	err = c.client.Patch(pt).
		Resource("googlekmskeyringiambindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}