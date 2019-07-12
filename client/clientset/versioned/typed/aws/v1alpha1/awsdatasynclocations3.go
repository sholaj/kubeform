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

// AwsDatasyncLocationS3sGetter has a method to return a AwsDatasyncLocationS3Interface.
// A group's client should implement this interface.
type AwsDatasyncLocationS3sGetter interface {
	AwsDatasyncLocationS3s() AwsDatasyncLocationS3Interface
}

// AwsDatasyncLocationS3Interface has methods to work with AwsDatasyncLocationS3 resources.
type AwsDatasyncLocationS3Interface interface {
	Create(*v1alpha1.AwsDatasyncLocationS3) (*v1alpha1.AwsDatasyncLocationS3, error)
	Update(*v1alpha1.AwsDatasyncLocationS3) (*v1alpha1.AwsDatasyncLocationS3, error)
	UpdateStatus(*v1alpha1.AwsDatasyncLocationS3) (*v1alpha1.AwsDatasyncLocationS3, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsDatasyncLocationS3, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsDatasyncLocationS3List, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDatasyncLocationS3, err error)
	AwsDatasyncLocationS3Expansion
}

// awsDatasyncLocationS3s implements AwsDatasyncLocationS3Interface
type awsDatasyncLocationS3s struct {
	client rest.Interface
}

// newAwsDatasyncLocationS3s returns a AwsDatasyncLocationS3s
func newAwsDatasyncLocationS3s(c *AwsV1alpha1Client) *awsDatasyncLocationS3s {
	return &awsDatasyncLocationS3s{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsDatasyncLocationS3, and returns the corresponding awsDatasyncLocationS3 object, and an error if there is any.
func (c *awsDatasyncLocationS3s) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDatasyncLocationS3, err error) {
	result = &v1alpha1.AwsDatasyncLocationS3{}
	err = c.client.Get().
		Resource("awsdatasynclocations3s").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsDatasyncLocationS3s that match those selectors.
func (c *awsDatasyncLocationS3s) List(opts v1.ListOptions) (result *v1alpha1.AwsDatasyncLocationS3List, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsDatasyncLocationS3List{}
	err = c.client.Get().
		Resource("awsdatasynclocations3s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsDatasyncLocationS3s.
func (c *awsDatasyncLocationS3s) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsdatasynclocations3s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsDatasyncLocationS3 and creates it.  Returns the server's representation of the awsDatasyncLocationS3, and an error, if there is any.
func (c *awsDatasyncLocationS3s) Create(awsDatasyncLocationS3 *v1alpha1.AwsDatasyncLocationS3) (result *v1alpha1.AwsDatasyncLocationS3, err error) {
	result = &v1alpha1.AwsDatasyncLocationS3{}
	err = c.client.Post().
		Resource("awsdatasynclocations3s").
		Body(awsDatasyncLocationS3).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsDatasyncLocationS3 and updates it. Returns the server's representation of the awsDatasyncLocationS3, and an error, if there is any.
func (c *awsDatasyncLocationS3s) Update(awsDatasyncLocationS3 *v1alpha1.AwsDatasyncLocationS3) (result *v1alpha1.AwsDatasyncLocationS3, err error) {
	result = &v1alpha1.AwsDatasyncLocationS3{}
	err = c.client.Put().
		Resource("awsdatasynclocations3s").
		Name(awsDatasyncLocationS3.Name).
		Body(awsDatasyncLocationS3).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsDatasyncLocationS3s) UpdateStatus(awsDatasyncLocationS3 *v1alpha1.AwsDatasyncLocationS3) (result *v1alpha1.AwsDatasyncLocationS3, err error) {
	result = &v1alpha1.AwsDatasyncLocationS3{}
	err = c.client.Put().
		Resource("awsdatasynclocations3s").
		Name(awsDatasyncLocationS3.Name).
		SubResource("status").
		Body(awsDatasyncLocationS3).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsDatasyncLocationS3 and deletes it. Returns an error if one occurs.
func (c *awsDatasyncLocationS3s) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsdatasynclocations3s").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsDatasyncLocationS3s) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsdatasynclocations3s").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsDatasyncLocationS3.
func (c *awsDatasyncLocationS3s) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDatasyncLocationS3, err error) {
	result = &v1alpha1.AwsDatasyncLocationS3{}
	err = c.client.Patch(pt).
		Resource("awsdatasynclocations3s").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}