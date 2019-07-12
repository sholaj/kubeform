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

// GoogleResourceManagerLiensGetter has a method to return a GoogleResourceManagerLienInterface.
// A group's client should implement this interface.
type GoogleResourceManagerLiensGetter interface {
	GoogleResourceManagerLiens() GoogleResourceManagerLienInterface
}

// GoogleResourceManagerLienInterface has methods to work with GoogleResourceManagerLien resources.
type GoogleResourceManagerLienInterface interface {
	Create(*v1alpha1.GoogleResourceManagerLien) (*v1alpha1.GoogleResourceManagerLien, error)
	Update(*v1alpha1.GoogleResourceManagerLien) (*v1alpha1.GoogleResourceManagerLien, error)
	UpdateStatus(*v1alpha1.GoogleResourceManagerLien) (*v1alpha1.GoogleResourceManagerLien, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleResourceManagerLien, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleResourceManagerLienList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleResourceManagerLien, err error)
	GoogleResourceManagerLienExpansion
}

// googleResourceManagerLiens implements GoogleResourceManagerLienInterface
type googleResourceManagerLiens struct {
	client rest.Interface
}

// newGoogleResourceManagerLiens returns a GoogleResourceManagerLiens
func newGoogleResourceManagerLiens(c *GoogleV1alpha1Client) *googleResourceManagerLiens {
	return &googleResourceManagerLiens{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleResourceManagerLien, and returns the corresponding googleResourceManagerLien object, and an error if there is any.
func (c *googleResourceManagerLiens) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleResourceManagerLien, err error) {
	result = &v1alpha1.GoogleResourceManagerLien{}
	err = c.client.Get().
		Resource("googleresourcemanagerliens").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleResourceManagerLiens that match those selectors.
func (c *googleResourceManagerLiens) List(opts v1.ListOptions) (result *v1alpha1.GoogleResourceManagerLienList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleResourceManagerLienList{}
	err = c.client.Get().
		Resource("googleresourcemanagerliens").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleResourceManagerLiens.
func (c *googleResourceManagerLiens) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googleresourcemanagerliens").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleResourceManagerLien and creates it.  Returns the server's representation of the googleResourceManagerLien, and an error, if there is any.
func (c *googleResourceManagerLiens) Create(googleResourceManagerLien *v1alpha1.GoogleResourceManagerLien) (result *v1alpha1.GoogleResourceManagerLien, err error) {
	result = &v1alpha1.GoogleResourceManagerLien{}
	err = c.client.Post().
		Resource("googleresourcemanagerliens").
		Body(googleResourceManagerLien).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleResourceManagerLien and updates it. Returns the server's representation of the googleResourceManagerLien, and an error, if there is any.
func (c *googleResourceManagerLiens) Update(googleResourceManagerLien *v1alpha1.GoogleResourceManagerLien) (result *v1alpha1.GoogleResourceManagerLien, err error) {
	result = &v1alpha1.GoogleResourceManagerLien{}
	err = c.client.Put().
		Resource("googleresourcemanagerliens").
		Name(googleResourceManagerLien.Name).
		Body(googleResourceManagerLien).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleResourceManagerLiens) UpdateStatus(googleResourceManagerLien *v1alpha1.GoogleResourceManagerLien) (result *v1alpha1.GoogleResourceManagerLien, err error) {
	result = &v1alpha1.GoogleResourceManagerLien{}
	err = c.client.Put().
		Resource("googleresourcemanagerliens").
		Name(googleResourceManagerLien.Name).
		SubResource("status").
		Body(googleResourceManagerLien).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleResourceManagerLien and deletes it. Returns an error if one occurs.
func (c *googleResourceManagerLiens) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googleresourcemanagerliens").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleResourceManagerLiens) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googleresourcemanagerliens").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleResourceManagerLien.
func (c *googleResourceManagerLiens) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleResourceManagerLien, err error) {
	result = &v1alpha1.GoogleResourceManagerLien{}
	err = c.client.Patch(pt).
		Resource("googleresourcemanagerliens").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}