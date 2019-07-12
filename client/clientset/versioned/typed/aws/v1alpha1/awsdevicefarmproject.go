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

// AwsDevicefarmProjectsGetter has a method to return a AwsDevicefarmProjectInterface.
// A group's client should implement this interface.
type AwsDevicefarmProjectsGetter interface {
	AwsDevicefarmProjects() AwsDevicefarmProjectInterface
}

// AwsDevicefarmProjectInterface has methods to work with AwsDevicefarmProject resources.
type AwsDevicefarmProjectInterface interface {
	Create(*v1alpha1.AwsDevicefarmProject) (*v1alpha1.AwsDevicefarmProject, error)
	Update(*v1alpha1.AwsDevicefarmProject) (*v1alpha1.AwsDevicefarmProject, error)
	UpdateStatus(*v1alpha1.AwsDevicefarmProject) (*v1alpha1.AwsDevicefarmProject, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsDevicefarmProject, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsDevicefarmProjectList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDevicefarmProject, err error)
	AwsDevicefarmProjectExpansion
}

// awsDevicefarmProjects implements AwsDevicefarmProjectInterface
type awsDevicefarmProjects struct {
	client rest.Interface
}

// newAwsDevicefarmProjects returns a AwsDevicefarmProjects
func newAwsDevicefarmProjects(c *AwsV1alpha1Client) *awsDevicefarmProjects {
	return &awsDevicefarmProjects{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsDevicefarmProject, and returns the corresponding awsDevicefarmProject object, and an error if there is any.
func (c *awsDevicefarmProjects) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDevicefarmProject, err error) {
	result = &v1alpha1.AwsDevicefarmProject{}
	err = c.client.Get().
		Resource("awsdevicefarmprojects").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsDevicefarmProjects that match those selectors.
func (c *awsDevicefarmProjects) List(opts v1.ListOptions) (result *v1alpha1.AwsDevicefarmProjectList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsDevicefarmProjectList{}
	err = c.client.Get().
		Resource("awsdevicefarmprojects").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsDevicefarmProjects.
func (c *awsDevicefarmProjects) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsdevicefarmprojects").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsDevicefarmProject and creates it.  Returns the server's representation of the awsDevicefarmProject, and an error, if there is any.
func (c *awsDevicefarmProjects) Create(awsDevicefarmProject *v1alpha1.AwsDevicefarmProject) (result *v1alpha1.AwsDevicefarmProject, err error) {
	result = &v1alpha1.AwsDevicefarmProject{}
	err = c.client.Post().
		Resource("awsdevicefarmprojects").
		Body(awsDevicefarmProject).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsDevicefarmProject and updates it. Returns the server's representation of the awsDevicefarmProject, and an error, if there is any.
func (c *awsDevicefarmProjects) Update(awsDevicefarmProject *v1alpha1.AwsDevicefarmProject) (result *v1alpha1.AwsDevicefarmProject, err error) {
	result = &v1alpha1.AwsDevicefarmProject{}
	err = c.client.Put().
		Resource("awsdevicefarmprojects").
		Name(awsDevicefarmProject.Name).
		Body(awsDevicefarmProject).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsDevicefarmProjects) UpdateStatus(awsDevicefarmProject *v1alpha1.AwsDevicefarmProject) (result *v1alpha1.AwsDevicefarmProject, err error) {
	result = &v1alpha1.AwsDevicefarmProject{}
	err = c.client.Put().
		Resource("awsdevicefarmprojects").
		Name(awsDevicefarmProject.Name).
		SubResource("status").
		Body(awsDevicefarmProject).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsDevicefarmProject and deletes it. Returns an error if one occurs.
func (c *awsDevicefarmProjects) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsdevicefarmprojects").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsDevicefarmProjects) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsdevicefarmprojects").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsDevicefarmProject.
func (c *awsDevicefarmProjects) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDevicefarmProject, err error) {
	result = &v1alpha1.AwsDevicefarmProject{}
	err = c.client.Patch(pt).
		Resource("awsdevicefarmprojects").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}