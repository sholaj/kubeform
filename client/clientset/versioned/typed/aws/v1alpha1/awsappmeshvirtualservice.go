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

// AwsAppmeshVirtualServicesGetter has a method to return a AwsAppmeshVirtualServiceInterface.
// A group's client should implement this interface.
type AwsAppmeshVirtualServicesGetter interface {
	AwsAppmeshVirtualServices() AwsAppmeshVirtualServiceInterface
}

// AwsAppmeshVirtualServiceInterface has methods to work with AwsAppmeshVirtualService resources.
type AwsAppmeshVirtualServiceInterface interface {
	Create(*v1alpha1.AwsAppmeshVirtualService) (*v1alpha1.AwsAppmeshVirtualService, error)
	Update(*v1alpha1.AwsAppmeshVirtualService) (*v1alpha1.AwsAppmeshVirtualService, error)
	UpdateStatus(*v1alpha1.AwsAppmeshVirtualService) (*v1alpha1.AwsAppmeshVirtualService, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsAppmeshVirtualService, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsAppmeshVirtualServiceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAppmeshVirtualService, err error)
	AwsAppmeshVirtualServiceExpansion
}

// awsAppmeshVirtualServices implements AwsAppmeshVirtualServiceInterface
type awsAppmeshVirtualServices struct {
	client rest.Interface
}

// newAwsAppmeshVirtualServices returns a AwsAppmeshVirtualServices
func newAwsAppmeshVirtualServices(c *AwsV1alpha1Client) *awsAppmeshVirtualServices {
	return &awsAppmeshVirtualServices{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsAppmeshVirtualService, and returns the corresponding awsAppmeshVirtualService object, and an error if there is any.
func (c *awsAppmeshVirtualServices) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsAppmeshVirtualService, err error) {
	result = &v1alpha1.AwsAppmeshVirtualService{}
	err = c.client.Get().
		Resource("awsappmeshvirtualservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsAppmeshVirtualServices that match those selectors.
func (c *awsAppmeshVirtualServices) List(opts v1.ListOptions) (result *v1alpha1.AwsAppmeshVirtualServiceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsAppmeshVirtualServiceList{}
	err = c.client.Get().
		Resource("awsappmeshvirtualservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsAppmeshVirtualServices.
func (c *awsAppmeshVirtualServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsappmeshvirtualservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsAppmeshVirtualService and creates it.  Returns the server's representation of the awsAppmeshVirtualService, and an error, if there is any.
func (c *awsAppmeshVirtualServices) Create(awsAppmeshVirtualService *v1alpha1.AwsAppmeshVirtualService) (result *v1alpha1.AwsAppmeshVirtualService, err error) {
	result = &v1alpha1.AwsAppmeshVirtualService{}
	err = c.client.Post().
		Resource("awsappmeshvirtualservices").
		Body(awsAppmeshVirtualService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsAppmeshVirtualService and updates it. Returns the server's representation of the awsAppmeshVirtualService, and an error, if there is any.
func (c *awsAppmeshVirtualServices) Update(awsAppmeshVirtualService *v1alpha1.AwsAppmeshVirtualService) (result *v1alpha1.AwsAppmeshVirtualService, err error) {
	result = &v1alpha1.AwsAppmeshVirtualService{}
	err = c.client.Put().
		Resource("awsappmeshvirtualservices").
		Name(awsAppmeshVirtualService.Name).
		Body(awsAppmeshVirtualService).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsAppmeshVirtualServices) UpdateStatus(awsAppmeshVirtualService *v1alpha1.AwsAppmeshVirtualService) (result *v1alpha1.AwsAppmeshVirtualService, err error) {
	result = &v1alpha1.AwsAppmeshVirtualService{}
	err = c.client.Put().
		Resource("awsappmeshvirtualservices").
		Name(awsAppmeshVirtualService.Name).
		SubResource("status").
		Body(awsAppmeshVirtualService).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsAppmeshVirtualService and deletes it. Returns an error if one occurs.
func (c *awsAppmeshVirtualServices) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsappmeshvirtualservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsAppmeshVirtualServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsappmeshvirtualservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsAppmeshVirtualService.
func (c *awsAppmeshVirtualServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAppmeshVirtualService, err error) {
	result = &v1alpha1.AwsAppmeshVirtualService{}
	err = c.client.Patch(pt).
		Resource("awsappmeshvirtualservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}