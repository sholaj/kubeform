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

// AwsSesTemplatesGetter has a method to return a AwsSesTemplateInterface.
// A group's client should implement this interface.
type AwsSesTemplatesGetter interface {
	AwsSesTemplates() AwsSesTemplateInterface
}

// AwsSesTemplateInterface has methods to work with AwsSesTemplate resources.
type AwsSesTemplateInterface interface {
	Create(*v1alpha1.AwsSesTemplate) (*v1alpha1.AwsSesTemplate, error)
	Update(*v1alpha1.AwsSesTemplate) (*v1alpha1.AwsSesTemplate, error)
	UpdateStatus(*v1alpha1.AwsSesTemplate) (*v1alpha1.AwsSesTemplate, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsSesTemplate, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsSesTemplateList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSesTemplate, err error)
	AwsSesTemplateExpansion
}

// awsSesTemplates implements AwsSesTemplateInterface
type awsSesTemplates struct {
	client rest.Interface
}

// newAwsSesTemplates returns a AwsSesTemplates
func newAwsSesTemplates(c *AwsV1alpha1Client) *awsSesTemplates {
	return &awsSesTemplates{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsSesTemplate, and returns the corresponding awsSesTemplate object, and an error if there is any.
func (c *awsSesTemplates) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSesTemplate, err error) {
	result = &v1alpha1.AwsSesTemplate{}
	err = c.client.Get().
		Resource("awssestemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsSesTemplates that match those selectors.
func (c *awsSesTemplates) List(opts v1.ListOptions) (result *v1alpha1.AwsSesTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsSesTemplateList{}
	err = c.client.Get().
		Resource("awssestemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsSesTemplates.
func (c *awsSesTemplates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awssestemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsSesTemplate and creates it.  Returns the server's representation of the awsSesTemplate, and an error, if there is any.
func (c *awsSesTemplates) Create(awsSesTemplate *v1alpha1.AwsSesTemplate) (result *v1alpha1.AwsSesTemplate, err error) {
	result = &v1alpha1.AwsSesTemplate{}
	err = c.client.Post().
		Resource("awssestemplates").
		Body(awsSesTemplate).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsSesTemplate and updates it. Returns the server's representation of the awsSesTemplate, and an error, if there is any.
func (c *awsSesTemplates) Update(awsSesTemplate *v1alpha1.AwsSesTemplate) (result *v1alpha1.AwsSesTemplate, err error) {
	result = &v1alpha1.AwsSesTemplate{}
	err = c.client.Put().
		Resource("awssestemplates").
		Name(awsSesTemplate.Name).
		Body(awsSesTemplate).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsSesTemplates) UpdateStatus(awsSesTemplate *v1alpha1.AwsSesTemplate) (result *v1alpha1.AwsSesTemplate, err error) {
	result = &v1alpha1.AwsSesTemplate{}
	err = c.client.Put().
		Resource("awssestemplates").
		Name(awsSesTemplate.Name).
		SubResource("status").
		Body(awsSesTemplate).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsSesTemplate and deletes it. Returns an error if one occurs.
func (c *awsSesTemplates) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awssestemplates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsSesTemplates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awssestemplates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsSesTemplate.
func (c *awsSesTemplates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSesTemplate, err error) {
	result = &v1alpha1.AwsSesTemplate{}
	err = c.client.Patch(pt).
		Resource("awssestemplates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}