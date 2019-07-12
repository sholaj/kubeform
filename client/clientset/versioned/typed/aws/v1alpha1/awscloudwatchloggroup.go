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

// AwsCloudwatchLogGroupsGetter has a method to return a AwsCloudwatchLogGroupInterface.
// A group's client should implement this interface.
type AwsCloudwatchLogGroupsGetter interface {
	AwsCloudwatchLogGroups() AwsCloudwatchLogGroupInterface
}

// AwsCloudwatchLogGroupInterface has methods to work with AwsCloudwatchLogGroup resources.
type AwsCloudwatchLogGroupInterface interface {
	Create(*v1alpha1.AwsCloudwatchLogGroup) (*v1alpha1.AwsCloudwatchLogGroup, error)
	Update(*v1alpha1.AwsCloudwatchLogGroup) (*v1alpha1.AwsCloudwatchLogGroup, error)
	UpdateStatus(*v1alpha1.AwsCloudwatchLogGroup) (*v1alpha1.AwsCloudwatchLogGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsCloudwatchLogGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsCloudwatchLogGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCloudwatchLogGroup, err error)
	AwsCloudwatchLogGroupExpansion
}

// awsCloudwatchLogGroups implements AwsCloudwatchLogGroupInterface
type awsCloudwatchLogGroups struct {
	client rest.Interface
}

// newAwsCloudwatchLogGroups returns a AwsCloudwatchLogGroups
func newAwsCloudwatchLogGroups(c *AwsV1alpha1Client) *awsCloudwatchLogGroups {
	return &awsCloudwatchLogGroups{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsCloudwatchLogGroup, and returns the corresponding awsCloudwatchLogGroup object, and an error if there is any.
func (c *awsCloudwatchLogGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsCloudwatchLogGroup, err error) {
	result = &v1alpha1.AwsCloudwatchLogGroup{}
	err = c.client.Get().
		Resource("awscloudwatchloggroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsCloudwatchLogGroups that match those selectors.
func (c *awsCloudwatchLogGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsCloudwatchLogGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsCloudwatchLogGroupList{}
	err = c.client.Get().
		Resource("awscloudwatchloggroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsCloudwatchLogGroups.
func (c *awsCloudwatchLogGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awscloudwatchloggroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsCloudwatchLogGroup and creates it.  Returns the server's representation of the awsCloudwatchLogGroup, and an error, if there is any.
func (c *awsCloudwatchLogGroups) Create(awsCloudwatchLogGroup *v1alpha1.AwsCloudwatchLogGroup) (result *v1alpha1.AwsCloudwatchLogGroup, err error) {
	result = &v1alpha1.AwsCloudwatchLogGroup{}
	err = c.client.Post().
		Resource("awscloudwatchloggroups").
		Body(awsCloudwatchLogGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsCloudwatchLogGroup and updates it. Returns the server's representation of the awsCloudwatchLogGroup, and an error, if there is any.
func (c *awsCloudwatchLogGroups) Update(awsCloudwatchLogGroup *v1alpha1.AwsCloudwatchLogGroup) (result *v1alpha1.AwsCloudwatchLogGroup, err error) {
	result = &v1alpha1.AwsCloudwatchLogGroup{}
	err = c.client.Put().
		Resource("awscloudwatchloggroups").
		Name(awsCloudwatchLogGroup.Name).
		Body(awsCloudwatchLogGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsCloudwatchLogGroups) UpdateStatus(awsCloudwatchLogGroup *v1alpha1.AwsCloudwatchLogGroup) (result *v1alpha1.AwsCloudwatchLogGroup, err error) {
	result = &v1alpha1.AwsCloudwatchLogGroup{}
	err = c.client.Put().
		Resource("awscloudwatchloggroups").
		Name(awsCloudwatchLogGroup.Name).
		SubResource("status").
		Body(awsCloudwatchLogGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsCloudwatchLogGroup and deletes it. Returns an error if one occurs.
func (c *awsCloudwatchLogGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awscloudwatchloggroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsCloudwatchLogGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awscloudwatchloggroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsCloudwatchLogGroup.
func (c *awsCloudwatchLogGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCloudwatchLogGroup, err error) {
	result = &v1alpha1.AwsCloudwatchLogGroup{}
	err = c.client.Patch(pt).
		Resource("awscloudwatchloggroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}