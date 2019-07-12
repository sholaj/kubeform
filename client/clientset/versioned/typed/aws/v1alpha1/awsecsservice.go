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

// AwsEcsServicesGetter has a method to return a AwsEcsServiceInterface.
// A group's client should implement this interface.
type AwsEcsServicesGetter interface {
	AwsEcsServices() AwsEcsServiceInterface
}

// AwsEcsServiceInterface has methods to work with AwsEcsService resources.
type AwsEcsServiceInterface interface {
	Create(*v1alpha1.AwsEcsService) (*v1alpha1.AwsEcsService, error)
	Update(*v1alpha1.AwsEcsService) (*v1alpha1.AwsEcsService, error)
	UpdateStatus(*v1alpha1.AwsEcsService) (*v1alpha1.AwsEcsService, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsEcsService, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsEcsServiceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsEcsService, err error)
	AwsEcsServiceExpansion
}

// awsEcsServices implements AwsEcsServiceInterface
type awsEcsServices struct {
	client rest.Interface
}

// newAwsEcsServices returns a AwsEcsServices
func newAwsEcsServices(c *AwsV1alpha1Client) *awsEcsServices {
	return &awsEcsServices{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsEcsService, and returns the corresponding awsEcsService object, and an error if there is any.
func (c *awsEcsServices) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsEcsService, err error) {
	result = &v1alpha1.AwsEcsService{}
	err = c.client.Get().
		Resource("awsecsservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsEcsServices that match those selectors.
func (c *awsEcsServices) List(opts v1.ListOptions) (result *v1alpha1.AwsEcsServiceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsEcsServiceList{}
	err = c.client.Get().
		Resource("awsecsservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsEcsServices.
func (c *awsEcsServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsecsservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsEcsService and creates it.  Returns the server's representation of the awsEcsService, and an error, if there is any.
func (c *awsEcsServices) Create(awsEcsService *v1alpha1.AwsEcsService) (result *v1alpha1.AwsEcsService, err error) {
	result = &v1alpha1.AwsEcsService{}
	err = c.client.Post().
		Resource("awsecsservices").
		Body(awsEcsService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsEcsService and updates it. Returns the server's representation of the awsEcsService, and an error, if there is any.
func (c *awsEcsServices) Update(awsEcsService *v1alpha1.AwsEcsService) (result *v1alpha1.AwsEcsService, err error) {
	result = &v1alpha1.AwsEcsService{}
	err = c.client.Put().
		Resource("awsecsservices").
		Name(awsEcsService.Name).
		Body(awsEcsService).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsEcsServices) UpdateStatus(awsEcsService *v1alpha1.AwsEcsService) (result *v1alpha1.AwsEcsService, err error) {
	result = &v1alpha1.AwsEcsService{}
	err = c.client.Put().
		Resource("awsecsservices").
		Name(awsEcsService.Name).
		SubResource("status").
		Body(awsEcsService).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsEcsService and deletes it. Returns an error if one occurs.
func (c *awsEcsServices) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsecsservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsEcsServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsecsservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsEcsService.
func (c *awsEcsServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsEcsService, err error) {
	result = &v1alpha1.AwsEcsService{}
	err = c.client.Patch(pt).
		Resource("awsecsservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}