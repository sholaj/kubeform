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

// GoogleComputeSharedVpcServiceProjectsGetter has a method to return a GoogleComputeSharedVpcServiceProjectInterface.
// A group's client should implement this interface.
type GoogleComputeSharedVpcServiceProjectsGetter interface {
	GoogleComputeSharedVpcServiceProjects() GoogleComputeSharedVpcServiceProjectInterface
}

// GoogleComputeSharedVpcServiceProjectInterface has methods to work with GoogleComputeSharedVpcServiceProject resources.
type GoogleComputeSharedVpcServiceProjectInterface interface {
	Create(*v1alpha1.GoogleComputeSharedVpcServiceProject) (*v1alpha1.GoogleComputeSharedVpcServiceProject, error)
	Update(*v1alpha1.GoogleComputeSharedVpcServiceProject) (*v1alpha1.GoogleComputeSharedVpcServiceProject, error)
	UpdateStatus(*v1alpha1.GoogleComputeSharedVpcServiceProject) (*v1alpha1.GoogleComputeSharedVpcServiceProject, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleComputeSharedVpcServiceProject, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleComputeSharedVpcServiceProjectList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComputeSharedVpcServiceProject, err error)
	GoogleComputeSharedVpcServiceProjectExpansion
}

// googleComputeSharedVpcServiceProjects implements GoogleComputeSharedVpcServiceProjectInterface
type googleComputeSharedVpcServiceProjects struct {
	client rest.Interface
}

// newGoogleComputeSharedVpcServiceProjects returns a GoogleComputeSharedVpcServiceProjects
func newGoogleComputeSharedVpcServiceProjects(c *GoogleV1alpha1Client) *googleComputeSharedVpcServiceProjects {
	return &googleComputeSharedVpcServiceProjects{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleComputeSharedVpcServiceProject, and returns the corresponding googleComputeSharedVpcServiceProject object, and an error if there is any.
func (c *googleComputeSharedVpcServiceProjects) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleComputeSharedVpcServiceProject, err error) {
	result = &v1alpha1.GoogleComputeSharedVpcServiceProject{}
	err = c.client.Get().
		Resource("googlecomputesharedvpcserviceprojects").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleComputeSharedVpcServiceProjects that match those selectors.
func (c *googleComputeSharedVpcServiceProjects) List(opts v1.ListOptions) (result *v1alpha1.GoogleComputeSharedVpcServiceProjectList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleComputeSharedVpcServiceProjectList{}
	err = c.client.Get().
		Resource("googlecomputesharedvpcserviceprojects").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleComputeSharedVpcServiceProjects.
func (c *googleComputeSharedVpcServiceProjects) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googlecomputesharedvpcserviceprojects").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleComputeSharedVpcServiceProject and creates it.  Returns the server's representation of the googleComputeSharedVpcServiceProject, and an error, if there is any.
func (c *googleComputeSharedVpcServiceProjects) Create(googleComputeSharedVpcServiceProject *v1alpha1.GoogleComputeSharedVpcServiceProject) (result *v1alpha1.GoogleComputeSharedVpcServiceProject, err error) {
	result = &v1alpha1.GoogleComputeSharedVpcServiceProject{}
	err = c.client.Post().
		Resource("googlecomputesharedvpcserviceprojects").
		Body(googleComputeSharedVpcServiceProject).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleComputeSharedVpcServiceProject and updates it. Returns the server's representation of the googleComputeSharedVpcServiceProject, and an error, if there is any.
func (c *googleComputeSharedVpcServiceProjects) Update(googleComputeSharedVpcServiceProject *v1alpha1.GoogleComputeSharedVpcServiceProject) (result *v1alpha1.GoogleComputeSharedVpcServiceProject, err error) {
	result = &v1alpha1.GoogleComputeSharedVpcServiceProject{}
	err = c.client.Put().
		Resource("googlecomputesharedvpcserviceprojects").
		Name(googleComputeSharedVpcServiceProject.Name).
		Body(googleComputeSharedVpcServiceProject).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleComputeSharedVpcServiceProjects) UpdateStatus(googleComputeSharedVpcServiceProject *v1alpha1.GoogleComputeSharedVpcServiceProject) (result *v1alpha1.GoogleComputeSharedVpcServiceProject, err error) {
	result = &v1alpha1.GoogleComputeSharedVpcServiceProject{}
	err = c.client.Put().
		Resource("googlecomputesharedvpcserviceprojects").
		Name(googleComputeSharedVpcServiceProject.Name).
		SubResource("status").
		Body(googleComputeSharedVpcServiceProject).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleComputeSharedVpcServiceProject and deletes it. Returns an error if one occurs.
func (c *googleComputeSharedVpcServiceProjects) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googlecomputesharedvpcserviceprojects").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleComputeSharedVpcServiceProjects) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googlecomputesharedvpcserviceprojects").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleComputeSharedVpcServiceProject.
func (c *googleComputeSharedVpcServiceProjects) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComputeSharedVpcServiceProject, err error) {
	result = &v1alpha1.GoogleComputeSharedVpcServiceProject{}
	err = c.client.Patch(pt).
		Resource("googlecomputesharedvpcserviceprojects").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}