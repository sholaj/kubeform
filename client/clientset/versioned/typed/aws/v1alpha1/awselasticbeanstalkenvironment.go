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

// AwsElasticBeanstalkEnvironmentsGetter has a method to return a AwsElasticBeanstalkEnvironmentInterface.
// A group's client should implement this interface.
type AwsElasticBeanstalkEnvironmentsGetter interface {
	AwsElasticBeanstalkEnvironments() AwsElasticBeanstalkEnvironmentInterface
}

// AwsElasticBeanstalkEnvironmentInterface has methods to work with AwsElasticBeanstalkEnvironment resources.
type AwsElasticBeanstalkEnvironmentInterface interface {
	Create(*v1alpha1.AwsElasticBeanstalkEnvironment) (*v1alpha1.AwsElasticBeanstalkEnvironment, error)
	Update(*v1alpha1.AwsElasticBeanstalkEnvironment) (*v1alpha1.AwsElasticBeanstalkEnvironment, error)
	UpdateStatus(*v1alpha1.AwsElasticBeanstalkEnvironment) (*v1alpha1.AwsElasticBeanstalkEnvironment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsElasticBeanstalkEnvironment, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsElasticBeanstalkEnvironmentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsElasticBeanstalkEnvironment, err error)
	AwsElasticBeanstalkEnvironmentExpansion
}

// awsElasticBeanstalkEnvironments implements AwsElasticBeanstalkEnvironmentInterface
type awsElasticBeanstalkEnvironments struct {
	client rest.Interface
}

// newAwsElasticBeanstalkEnvironments returns a AwsElasticBeanstalkEnvironments
func newAwsElasticBeanstalkEnvironments(c *AwsV1alpha1Client) *awsElasticBeanstalkEnvironments {
	return &awsElasticBeanstalkEnvironments{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsElasticBeanstalkEnvironment, and returns the corresponding awsElasticBeanstalkEnvironment object, and an error if there is any.
func (c *awsElasticBeanstalkEnvironments) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsElasticBeanstalkEnvironment, err error) {
	result = &v1alpha1.AwsElasticBeanstalkEnvironment{}
	err = c.client.Get().
		Resource("awselasticbeanstalkenvironments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsElasticBeanstalkEnvironments that match those selectors.
func (c *awsElasticBeanstalkEnvironments) List(opts v1.ListOptions) (result *v1alpha1.AwsElasticBeanstalkEnvironmentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsElasticBeanstalkEnvironmentList{}
	err = c.client.Get().
		Resource("awselasticbeanstalkenvironments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsElasticBeanstalkEnvironments.
func (c *awsElasticBeanstalkEnvironments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awselasticbeanstalkenvironments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsElasticBeanstalkEnvironment and creates it.  Returns the server's representation of the awsElasticBeanstalkEnvironment, and an error, if there is any.
func (c *awsElasticBeanstalkEnvironments) Create(awsElasticBeanstalkEnvironment *v1alpha1.AwsElasticBeanstalkEnvironment) (result *v1alpha1.AwsElasticBeanstalkEnvironment, err error) {
	result = &v1alpha1.AwsElasticBeanstalkEnvironment{}
	err = c.client.Post().
		Resource("awselasticbeanstalkenvironments").
		Body(awsElasticBeanstalkEnvironment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsElasticBeanstalkEnvironment and updates it. Returns the server's representation of the awsElasticBeanstalkEnvironment, and an error, if there is any.
func (c *awsElasticBeanstalkEnvironments) Update(awsElasticBeanstalkEnvironment *v1alpha1.AwsElasticBeanstalkEnvironment) (result *v1alpha1.AwsElasticBeanstalkEnvironment, err error) {
	result = &v1alpha1.AwsElasticBeanstalkEnvironment{}
	err = c.client.Put().
		Resource("awselasticbeanstalkenvironments").
		Name(awsElasticBeanstalkEnvironment.Name).
		Body(awsElasticBeanstalkEnvironment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsElasticBeanstalkEnvironments) UpdateStatus(awsElasticBeanstalkEnvironment *v1alpha1.AwsElasticBeanstalkEnvironment) (result *v1alpha1.AwsElasticBeanstalkEnvironment, err error) {
	result = &v1alpha1.AwsElasticBeanstalkEnvironment{}
	err = c.client.Put().
		Resource("awselasticbeanstalkenvironments").
		Name(awsElasticBeanstalkEnvironment.Name).
		SubResource("status").
		Body(awsElasticBeanstalkEnvironment).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsElasticBeanstalkEnvironment and deletes it. Returns an error if one occurs.
func (c *awsElasticBeanstalkEnvironments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awselasticbeanstalkenvironments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsElasticBeanstalkEnvironments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awselasticbeanstalkenvironments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsElasticBeanstalkEnvironment.
func (c *awsElasticBeanstalkEnvironments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsElasticBeanstalkEnvironment, err error) {
	result = &v1alpha1.AwsElasticBeanstalkEnvironment{}
	err = c.client.Patch(pt).
		Resource("awselasticbeanstalkenvironments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}