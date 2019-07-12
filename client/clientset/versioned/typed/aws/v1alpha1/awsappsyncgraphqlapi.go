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

// AwsAppsyncGraphqlApisGetter has a method to return a AwsAppsyncGraphqlApiInterface.
// A group's client should implement this interface.
type AwsAppsyncGraphqlApisGetter interface {
	AwsAppsyncGraphqlApis() AwsAppsyncGraphqlApiInterface
}

// AwsAppsyncGraphqlApiInterface has methods to work with AwsAppsyncGraphqlApi resources.
type AwsAppsyncGraphqlApiInterface interface {
	Create(*v1alpha1.AwsAppsyncGraphqlApi) (*v1alpha1.AwsAppsyncGraphqlApi, error)
	Update(*v1alpha1.AwsAppsyncGraphqlApi) (*v1alpha1.AwsAppsyncGraphqlApi, error)
	UpdateStatus(*v1alpha1.AwsAppsyncGraphqlApi) (*v1alpha1.AwsAppsyncGraphqlApi, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsAppsyncGraphqlApi, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsAppsyncGraphqlApiList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAppsyncGraphqlApi, err error)
	AwsAppsyncGraphqlApiExpansion
}

// awsAppsyncGraphqlApis implements AwsAppsyncGraphqlApiInterface
type awsAppsyncGraphqlApis struct {
	client rest.Interface
}

// newAwsAppsyncGraphqlApis returns a AwsAppsyncGraphqlApis
func newAwsAppsyncGraphqlApis(c *AwsV1alpha1Client) *awsAppsyncGraphqlApis {
	return &awsAppsyncGraphqlApis{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsAppsyncGraphqlApi, and returns the corresponding awsAppsyncGraphqlApi object, and an error if there is any.
func (c *awsAppsyncGraphqlApis) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsAppsyncGraphqlApi, err error) {
	result = &v1alpha1.AwsAppsyncGraphqlApi{}
	err = c.client.Get().
		Resource("awsappsyncgraphqlapis").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsAppsyncGraphqlApis that match those selectors.
func (c *awsAppsyncGraphqlApis) List(opts v1.ListOptions) (result *v1alpha1.AwsAppsyncGraphqlApiList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsAppsyncGraphqlApiList{}
	err = c.client.Get().
		Resource("awsappsyncgraphqlapis").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsAppsyncGraphqlApis.
func (c *awsAppsyncGraphqlApis) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsappsyncgraphqlapis").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsAppsyncGraphqlApi and creates it.  Returns the server's representation of the awsAppsyncGraphqlApi, and an error, if there is any.
func (c *awsAppsyncGraphqlApis) Create(awsAppsyncGraphqlApi *v1alpha1.AwsAppsyncGraphqlApi) (result *v1alpha1.AwsAppsyncGraphqlApi, err error) {
	result = &v1alpha1.AwsAppsyncGraphqlApi{}
	err = c.client.Post().
		Resource("awsappsyncgraphqlapis").
		Body(awsAppsyncGraphqlApi).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsAppsyncGraphqlApi and updates it. Returns the server's representation of the awsAppsyncGraphqlApi, and an error, if there is any.
func (c *awsAppsyncGraphqlApis) Update(awsAppsyncGraphqlApi *v1alpha1.AwsAppsyncGraphqlApi) (result *v1alpha1.AwsAppsyncGraphqlApi, err error) {
	result = &v1alpha1.AwsAppsyncGraphqlApi{}
	err = c.client.Put().
		Resource("awsappsyncgraphqlapis").
		Name(awsAppsyncGraphqlApi.Name).
		Body(awsAppsyncGraphqlApi).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsAppsyncGraphqlApis) UpdateStatus(awsAppsyncGraphqlApi *v1alpha1.AwsAppsyncGraphqlApi) (result *v1alpha1.AwsAppsyncGraphqlApi, err error) {
	result = &v1alpha1.AwsAppsyncGraphqlApi{}
	err = c.client.Put().
		Resource("awsappsyncgraphqlapis").
		Name(awsAppsyncGraphqlApi.Name).
		SubResource("status").
		Body(awsAppsyncGraphqlApi).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsAppsyncGraphqlApi and deletes it. Returns an error if one occurs.
func (c *awsAppsyncGraphqlApis) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsappsyncgraphqlapis").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsAppsyncGraphqlApis) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsappsyncgraphqlapis").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsAppsyncGraphqlApi.
func (c *awsAppsyncGraphqlApis) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAppsyncGraphqlApi, err error) {
	result = &v1alpha1.AwsAppsyncGraphqlApi{}
	err = c.client.Patch(pt).
		Resource("awsappsyncgraphqlapis").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}