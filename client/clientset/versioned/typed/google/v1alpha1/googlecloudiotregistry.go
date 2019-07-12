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

// GoogleCloudiotRegistriesGetter has a method to return a GoogleCloudiotRegistryInterface.
// A group's client should implement this interface.
type GoogleCloudiotRegistriesGetter interface {
	GoogleCloudiotRegistries() GoogleCloudiotRegistryInterface
}

// GoogleCloudiotRegistryInterface has methods to work with GoogleCloudiotRegistry resources.
type GoogleCloudiotRegistryInterface interface {
	Create(*v1alpha1.GoogleCloudiotRegistry) (*v1alpha1.GoogleCloudiotRegistry, error)
	Update(*v1alpha1.GoogleCloudiotRegistry) (*v1alpha1.GoogleCloudiotRegistry, error)
	UpdateStatus(*v1alpha1.GoogleCloudiotRegistry) (*v1alpha1.GoogleCloudiotRegistry, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleCloudiotRegistry, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleCloudiotRegistryList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleCloudiotRegistry, err error)
	GoogleCloudiotRegistryExpansion
}

// googleCloudiotRegistries implements GoogleCloudiotRegistryInterface
type googleCloudiotRegistries struct {
	client rest.Interface
}

// newGoogleCloudiotRegistries returns a GoogleCloudiotRegistries
func newGoogleCloudiotRegistries(c *GoogleV1alpha1Client) *googleCloudiotRegistries {
	return &googleCloudiotRegistries{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleCloudiotRegistry, and returns the corresponding googleCloudiotRegistry object, and an error if there is any.
func (c *googleCloudiotRegistries) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleCloudiotRegistry, err error) {
	result = &v1alpha1.GoogleCloudiotRegistry{}
	err = c.client.Get().
		Resource("googlecloudiotregistries").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleCloudiotRegistries that match those selectors.
func (c *googleCloudiotRegistries) List(opts v1.ListOptions) (result *v1alpha1.GoogleCloudiotRegistryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleCloudiotRegistryList{}
	err = c.client.Get().
		Resource("googlecloudiotregistries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleCloudiotRegistries.
func (c *googleCloudiotRegistries) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googlecloudiotregistries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleCloudiotRegistry and creates it.  Returns the server's representation of the googleCloudiotRegistry, and an error, if there is any.
func (c *googleCloudiotRegistries) Create(googleCloudiotRegistry *v1alpha1.GoogleCloudiotRegistry) (result *v1alpha1.GoogleCloudiotRegistry, err error) {
	result = &v1alpha1.GoogleCloudiotRegistry{}
	err = c.client.Post().
		Resource("googlecloudiotregistries").
		Body(googleCloudiotRegistry).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleCloudiotRegistry and updates it. Returns the server's representation of the googleCloudiotRegistry, and an error, if there is any.
func (c *googleCloudiotRegistries) Update(googleCloudiotRegistry *v1alpha1.GoogleCloudiotRegistry) (result *v1alpha1.GoogleCloudiotRegistry, err error) {
	result = &v1alpha1.GoogleCloudiotRegistry{}
	err = c.client.Put().
		Resource("googlecloudiotregistries").
		Name(googleCloudiotRegistry.Name).
		Body(googleCloudiotRegistry).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleCloudiotRegistries) UpdateStatus(googleCloudiotRegistry *v1alpha1.GoogleCloudiotRegistry) (result *v1alpha1.GoogleCloudiotRegistry, err error) {
	result = &v1alpha1.GoogleCloudiotRegistry{}
	err = c.client.Put().
		Resource("googlecloudiotregistries").
		Name(googleCloudiotRegistry.Name).
		SubResource("status").
		Body(googleCloudiotRegistry).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleCloudiotRegistry and deletes it. Returns an error if one occurs.
func (c *googleCloudiotRegistries) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googlecloudiotregistries").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleCloudiotRegistries) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googlecloudiotregistries").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleCloudiotRegistry.
func (c *googleCloudiotRegistries) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleCloudiotRegistry, err error) {
	result = &v1alpha1.GoogleCloudiotRegistry{}
	err = c.client.Patch(pt).
		Resource("googlecloudiotregistries").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}