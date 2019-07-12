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

// GoogleSpannerDatabasesGetter has a method to return a GoogleSpannerDatabaseInterface.
// A group's client should implement this interface.
type GoogleSpannerDatabasesGetter interface {
	GoogleSpannerDatabases() GoogleSpannerDatabaseInterface
}

// GoogleSpannerDatabaseInterface has methods to work with GoogleSpannerDatabase resources.
type GoogleSpannerDatabaseInterface interface {
	Create(*v1alpha1.GoogleSpannerDatabase) (*v1alpha1.GoogleSpannerDatabase, error)
	Update(*v1alpha1.GoogleSpannerDatabase) (*v1alpha1.GoogleSpannerDatabase, error)
	UpdateStatus(*v1alpha1.GoogleSpannerDatabase) (*v1alpha1.GoogleSpannerDatabase, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleSpannerDatabase, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleSpannerDatabaseList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleSpannerDatabase, err error)
	GoogleSpannerDatabaseExpansion
}

// googleSpannerDatabases implements GoogleSpannerDatabaseInterface
type googleSpannerDatabases struct {
	client rest.Interface
}

// newGoogleSpannerDatabases returns a GoogleSpannerDatabases
func newGoogleSpannerDatabases(c *GoogleV1alpha1Client) *googleSpannerDatabases {
	return &googleSpannerDatabases{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleSpannerDatabase, and returns the corresponding googleSpannerDatabase object, and an error if there is any.
func (c *googleSpannerDatabases) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleSpannerDatabase, err error) {
	result = &v1alpha1.GoogleSpannerDatabase{}
	err = c.client.Get().
		Resource("googlespannerdatabases").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleSpannerDatabases that match those selectors.
func (c *googleSpannerDatabases) List(opts v1.ListOptions) (result *v1alpha1.GoogleSpannerDatabaseList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleSpannerDatabaseList{}
	err = c.client.Get().
		Resource("googlespannerdatabases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleSpannerDatabases.
func (c *googleSpannerDatabases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googlespannerdatabases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleSpannerDatabase and creates it.  Returns the server's representation of the googleSpannerDatabase, and an error, if there is any.
func (c *googleSpannerDatabases) Create(googleSpannerDatabase *v1alpha1.GoogleSpannerDatabase) (result *v1alpha1.GoogleSpannerDatabase, err error) {
	result = &v1alpha1.GoogleSpannerDatabase{}
	err = c.client.Post().
		Resource("googlespannerdatabases").
		Body(googleSpannerDatabase).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleSpannerDatabase and updates it. Returns the server's representation of the googleSpannerDatabase, and an error, if there is any.
func (c *googleSpannerDatabases) Update(googleSpannerDatabase *v1alpha1.GoogleSpannerDatabase) (result *v1alpha1.GoogleSpannerDatabase, err error) {
	result = &v1alpha1.GoogleSpannerDatabase{}
	err = c.client.Put().
		Resource("googlespannerdatabases").
		Name(googleSpannerDatabase.Name).
		Body(googleSpannerDatabase).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleSpannerDatabases) UpdateStatus(googleSpannerDatabase *v1alpha1.GoogleSpannerDatabase) (result *v1alpha1.GoogleSpannerDatabase, err error) {
	result = &v1alpha1.GoogleSpannerDatabase{}
	err = c.client.Put().
		Resource("googlespannerdatabases").
		Name(googleSpannerDatabase.Name).
		SubResource("status").
		Body(googleSpannerDatabase).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleSpannerDatabase and deletes it. Returns an error if one occurs.
func (c *googleSpannerDatabases) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googlespannerdatabases").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleSpannerDatabases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googlespannerdatabases").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleSpannerDatabase.
func (c *googleSpannerDatabases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleSpannerDatabase, err error) {
	result = &v1alpha1.GoogleSpannerDatabase{}
	err = c.client.Patch(pt).
		Resource("googlespannerdatabases").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}