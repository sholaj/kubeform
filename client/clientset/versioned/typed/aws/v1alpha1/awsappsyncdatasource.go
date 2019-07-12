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

// AwsAppsyncDatasourcesGetter has a method to return a AwsAppsyncDatasourceInterface.
// A group's client should implement this interface.
type AwsAppsyncDatasourcesGetter interface {
	AwsAppsyncDatasources() AwsAppsyncDatasourceInterface
}

// AwsAppsyncDatasourceInterface has methods to work with AwsAppsyncDatasource resources.
type AwsAppsyncDatasourceInterface interface {
	Create(*v1alpha1.AwsAppsyncDatasource) (*v1alpha1.AwsAppsyncDatasource, error)
	Update(*v1alpha1.AwsAppsyncDatasource) (*v1alpha1.AwsAppsyncDatasource, error)
	UpdateStatus(*v1alpha1.AwsAppsyncDatasource) (*v1alpha1.AwsAppsyncDatasource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsAppsyncDatasource, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsAppsyncDatasourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAppsyncDatasource, err error)
	AwsAppsyncDatasourceExpansion
}

// awsAppsyncDatasources implements AwsAppsyncDatasourceInterface
type awsAppsyncDatasources struct {
	client rest.Interface
}

// newAwsAppsyncDatasources returns a AwsAppsyncDatasources
func newAwsAppsyncDatasources(c *AwsV1alpha1Client) *awsAppsyncDatasources {
	return &awsAppsyncDatasources{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsAppsyncDatasource, and returns the corresponding awsAppsyncDatasource object, and an error if there is any.
func (c *awsAppsyncDatasources) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsAppsyncDatasource, err error) {
	result = &v1alpha1.AwsAppsyncDatasource{}
	err = c.client.Get().
		Resource("awsappsyncdatasources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsAppsyncDatasources that match those selectors.
func (c *awsAppsyncDatasources) List(opts v1.ListOptions) (result *v1alpha1.AwsAppsyncDatasourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsAppsyncDatasourceList{}
	err = c.client.Get().
		Resource("awsappsyncdatasources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsAppsyncDatasources.
func (c *awsAppsyncDatasources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsappsyncdatasources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsAppsyncDatasource and creates it.  Returns the server's representation of the awsAppsyncDatasource, and an error, if there is any.
func (c *awsAppsyncDatasources) Create(awsAppsyncDatasource *v1alpha1.AwsAppsyncDatasource) (result *v1alpha1.AwsAppsyncDatasource, err error) {
	result = &v1alpha1.AwsAppsyncDatasource{}
	err = c.client.Post().
		Resource("awsappsyncdatasources").
		Body(awsAppsyncDatasource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsAppsyncDatasource and updates it. Returns the server's representation of the awsAppsyncDatasource, and an error, if there is any.
func (c *awsAppsyncDatasources) Update(awsAppsyncDatasource *v1alpha1.AwsAppsyncDatasource) (result *v1alpha1.AwsAppsyncDatasource, err error) {
	result = &v1alpha1.AwsAppsyncDatasource{}
	err = c.client.Put().
		Resource("awsappsyncdatasources").
		Name(awsAppsyncDatasource.Name).
		Body(awsAppsyncDatasource).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsAppsyncDatasources) UpdateStatus(awsAppsyncDatasource *v1alpha1.AwsAppsyncDatasource) (result *v1alpha1.AwsAppsyncDatasource, err error) {
	result = &v1alpha1.AwsAppsyncDatasource{}
	err = c.client.Put().
		Resource("awsappsyncdatasources").
		Name(awsAppsyncDatasource.Name).
		SubResource("status").
		Body(awsAppsyncDatasource).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsAppsyncDatasource and deletes it. Returns an error if one occurs.
func (c *awsAppsyncDatasources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsappsyncdatasources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsAppsyncDatasources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsappsyncdatasources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsAppsyncDatasource.
func (c *awsAppsyncDatasources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAppsyncDatasource, err error) {
	result = &v1alpha1.AwsAppsyncDatasource{}
	err = c.client.Patch(pt).
		Resource("awsappsyncdatasources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}