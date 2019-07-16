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

// AutoscalingGroupsGetter has a method to return a AutoscalingGroupInterface.
// A group's client should implement this interface.
type AutoscalingGroupsGetter interface {
	AutoscalingGroups() AutoscalingGroupInterface
}

// AutoscalingGroupInterface has methods to work with AutoscalingGroup resources.
type AutoscalingGroupInterface interface {
	Create(*v1alpha1.AutoscalingGroup) (*v1alpha1.AutoscalingGroup, error)
	Update(*v1alpha1.AutoscalingGroup) (*v1alpha1.AutoscalingGroup, error)
	UpdateStatus(*v1alpha1.AutoscalingGroup) (*v1alpha1.AutoscalingGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AutoscalingGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.AutoscalingGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AutoscalingGroup, err error)
	AutoscalingGroupExpansion
}

// autoscalingGroups implements AutoscalingGroupInterface
type autoscalingGroups struct {
	client rest.Interface
}

// newAutoscalingGroups returns a AutoscalingGroups
func newAutoscalingGroups(c *AwsV1alpha1Client) *autoscalingGroups {
	return &autoscalingGroups{
		client: c.RESTClient(),
	}
}

// Get takes name of the autoscalingGroup, and returns the corresponding autoscalingGroup object, and an error if there is any.
func (c *autoscalingGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AutoscalingGroup, err error) {
	result = &v1alpha1.AutoscalingGroup{}
	err = c.client.Get().
		Resource("autoscalinggroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AutoscalingGroups that match those selectors.
func (c *autoscalingGroups) List(opts v1.ListOptions) (result *v1alpha1.AutoscalingGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AutoscalingGroupList{}
	err = c.client.Get().
		Resource("autoscalinggroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested autoscalingGroups.
func (c *autoscalingGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("autoscalinggroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a autoscalingGroup and creates it.  Returns the server's representation of the autoscalingGroup, and an error, if there is any.
func (c *autoscalingGroups) Create(autoscalingGroup *v1alpha1.AutoscalingGroup) (result *v1alpha1.AutoscalingGroup, err error) {
	result = &v1alpha1.AutoscalingGroup{}
	err = c.client.Post().
		Resource("autoscalinggroups").
		Body(autoscalingGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a autoscalingGroup and updates it. Returns the server's representation of the autoscalingGroup, and an error, if there is any.
func (c *autoscalingGroups) Update(autoscalingGroup *v1alpha1.AutoscalingGroup) (result *v1alpha1.AutoscalingGroup, err error) {
	result = &v1alpha1.AutoscalingGroup{}
	err = c.client.Put().
		Resource("autoscalinggroups").
		Name(autoscalingGroup.Name).
		Body(autoscalingGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *autoscalingGroups) UpdateStatus(autoscalingGroup *v1alpha1.AutoscalingGroup) (result *v1alpha1.AutoscalingGroup, err error) {
	result = &v1alpha1.AutoscalingGroup{}
	err = c.client.Put().
		Resource("autoscalinggroups").
		Name(autoscalingGroup.Name).
		SubResource("status").
		Body(autoscalingGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the autoscalingGroup and deletes it. Returns an error if one occurs.
func (c *autoscalingGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("autoscalinggroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *autoscalingGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("autoscalinggroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched autoscalingGroup.
func (c *autoscalingGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AutoscalingGroup, err error) {
	result = &v1alpha1.AutoscalingGroup{}
	err = c.client.Patch(pt).
		Resource("autoscalinggroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}