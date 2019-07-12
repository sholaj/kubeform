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

// GoogleStorageBucketsGetter has a method to return a GoogleStorageBucketInterface.
// A group's client should implement this interface.
type GoogleStorageBucketsGetter interface {
	GoogleStorageBuckets() GoogleStorageBucketInterface
}

// GoogleStorageBucketInterface has methods to work with GoogleStorageBucket resources.
type GoogleStorageBucketInterface interface {
	Create(*v1alpha1.GoogleStorageBucket) (*v1alpha1.GoogleStorageBucket, error)
	Update(*v1alpha1.GoogleStorageBucket) (*v1alpha1.GoogleStorageBucket, error)
	UpdateStatus(*v1alpha1.GoogleStorageBucket) (*v1alpha1.GoogleStorageBucket, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleStorageBucket, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleStorageBucketList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleStorageBucket, err error)
	GoogleStorageBucketExpansion
}

// googleStorageBuckets implements GoogleStorageBucketInterface
type googleStorageBuckets struct {
	client rest.Interface
}

// newGoogleStorageBuckets returns a GoogleStorageBuckets
func newGoogleStorageBuckets(c *GoogleV1alpha1Client) *googleStorageBuckets {
	return &googleStorageBuckets{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleStorageBucket, and returns the corresponding googleStorageBucket object, and an error if there is any.
func (c *googleStorageBuckets) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleStorageBucket, err error) {
	result = &v1alpha1.GoogleStorageBucket{}
	err = c.client.Get().
		Resource("googlestoragebuckets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleStorageBuckets that match those selectors.
func (c *googleStorageBuckets) List(opts v1.ListOptions) (result *v1alpha1.GoogleStorageBucketList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleStorageBucketList{}
	err = c.client.Get().
		Resource("googlestoragebuckets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleStorageBuckets.
func (c *googleStorageBuckets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googlestoragebuckets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleStorageBucket and creates it.  Returns the server's representation of the googleStorageBucket, and an error, if there is any.
func (c *googleStorageBuckets) Create(googleStorageBucket *v1alpha1.GoogleStorageBucket) (result *v1alpha1.GoogleStorageBucket, err error) {
	result = &v1alpha1.GoogleStorageBucket{}
	err = c.client.Post().
		Resource("googlestoragebuckets").
		Body(googleStorageBucket).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleStorageBucket and updates it. Returns the server's representation of the googleStorageBucket, and an error, if there is any.
func (c *googleStorageBuckets) Update(googleStorageBucket *v1alpha1.GoogleStorageBucket) (result *v1alpha1.GoogleStorageBucket, err error) {
	result = &v1alpha1.GoogleStorageBucket{}
	err = c.client.Put().
		Resource("googlestoragebuckets").
		Name(googleStorageBucket.Name).
		Body(googleStorageBucket).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleStorageBuckets) UpdateStatus(googleStorageBucket *v1alpha1.GoogleStorageBucket) (result *v1alpha1.GoogleStorageBucket, err error) {
	result = &v1alpha1.GoogleStorageBucket{}
	err = c.client.Put().
		Resource("googlestoragebuckets").
		Name(googleStorageBucket.Name).
		SubResource("status").
		Body(googleStorageBucket).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleStorageBucket and deletes it. Returns an error if one occurs.
func (c *googleStorageBuckets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googlestoragebuckets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleStorageBuckets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googlestoragebuckets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleStorageBucket.
func (c *googleStorageBuckets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleStorageBucket, err error) {
	result = &v1alpha1.GoogleStorageBucket{}
	err = c.client.Patch(pt).
		Resource("googlestoragebuckets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}