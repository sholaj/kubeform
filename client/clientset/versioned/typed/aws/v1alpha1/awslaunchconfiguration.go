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

// AwsLaunchConfigurationsGetter has a method to return a AwsLaunchConfigurationInterface.
// A group's client should implement this interface.
type AwsLaunchConfigurationsGetter interface {
	AwsLaunchConfigurations() AwsLaunchConfigurationInterface
}

// AwsLaunchConfigurationInterface has methods to work with AwsLaunchConfiguration resources.
type AwsLaunchConfigurationInterface interface {
	Create(*v1alpha1.AwsLaunchConfiguration) (*v1alpha1.AwsLaunchConfiguration, error)
	Update(*v1alpha1.AwsLaunchConfiguration) (*v1alpha1.AwsLaunchConfiguration, error)
	UpdateStatus(*v1alpha1.AwsLaunchConfiguration) (*v1alpha1.AwsLaunchConfiguration, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsLaunchConfiguration, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsLaunchConfigurationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsLaunchConfiguration, err error)
	AwsLaunchConfigurationExpansion
}

// awsLaunchConfigurations implements AwsLaunchConfigurationInterface
type awsLaunchConfigurations struct {
	client rest.Interface
}

// newAwsLaunchConfigurations returns a AwsLaunchConfigurations
func newAwsLaunchConfigurations(c *AwsV1alpha1Client) *awsLaunchConfigurations {
	return &awsLaunchConfigurations{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsLaunchConfiguration, and returns the corresponding awsLaunchConfiguration object, and an error if there is any.
func (c *awsLaunchConfigurations) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsLaunchConfiguration, err error) {
	result = &v1alpha1.AwsLaunchConfiguration{}
	err = c.client.Get().
		Resource("awslaunchconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsLaunchConfigurations that match those selectors.
func (c *awsLaunchConfigurations) List(opts v1.ListOptions) (result *v1alpha1.AwsLaunchConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsLaunchConfigurationList{}
	err = c.client.Get().
		Resource("awslaunchconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsLaunchConfigurations.
func (c *awsLaunchConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awslaunchconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsLaunchConfiguration and creates it.  Returns the server's representation of the awsLaunchConfiguration, and an error, if there is any.
func (c *awsLaunchConfigurations) Create(awsLaunchConfiguration *v1alpha1.AwsLaunchConfiguration) (result *v1alpha1.AwsLaunchConfiguration, err error) {
	result = &v1alpha1.AwsLaunchConfiguration{}
	err = c.client.Post().
		Resource("awslaunchconfigurations").
		Body(awsLaunchConfiguration).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsLaunchConfiguration and updates it. Returns the server's representation of the awsLaunchConfiguration, and an error, if there is any.
func (c *awsLaunchConfigurations) Update(awsLaunchConfiguration *v1alpha1.AwsLaunchConfiguration) (result *v1alpha1.AwsLaunchConfiguration, err error) {
	result = &v1alpha1.AwsLaunchConfiguration{}
	err = c.client.Put().
		Resource("awslaunchconfigurations").
		Name(awsLaunchConfiguration.Name).
		Body(awsLaunchConfiguration).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsLaunchConfigurations) UpdateStatus(awsLaunchConfiguration *v1alpha1.AwsLaunchConfiguration) (result *v1alpha1.AwsLaunchConfiguration, err error) {
	result = &v1alpha1.AwsLaunchConfiguration{}
	err = c.client.Put().
		Resource("awslaunchconfigurations").
		Name(awsLaunchConfiguration.Name).
		SubResource("status").
		Body(awsLaunchConfiguration).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsLaunchConfiguration and deletes it. Returns an error if one occurs.
func (c *awsLaunchConfigurations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awslaunchconfigurations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsLaunchConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awslaunchconfigurations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsLaunchConfiguration.
func (c *awsLaunchConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsLaunchConfiguration, err error) {
	result = &v1alpha1.AwsLaunchConfiguration{}
	err = c.client.Patch(pt).
		Resource("awslaunchconfigurations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}