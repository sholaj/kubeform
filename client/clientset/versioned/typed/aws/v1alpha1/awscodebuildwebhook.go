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

// AwsCodebuildWebhooksGetter has a method to return a AwsCodebuildWebhookInterface.
// A group's client should implement this interface.
type AwsCodebuildWebhooksGetter interface {
	AwsCodebuildWebhooks() AwsCodebuildWebhookInterface
}

// AwsCodebuildWebhookInterface has methods to work with AwsCodebuildWebhook resources.
type AwsCodebuildWebhookInterface interface {
	Create(*v1alpha1.AwsCodebuildWebhook) (*v1alpha1.AwsCodebuildWebhook, error)
	Update(*v1alpha1.AwsCodebuildWebhook) (*v1alpha1.AwsCodebuildWebhook, error)
	UpdateStatus(*v1alpha1.AwsCodebuildWebhook) (*v1alpha1.AwsCodebuildWebhook, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsCodebuildWebhook, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsCodebuildWebhookList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCodebuildWebhook, err error)
	AwsCodebuildWebhookExpansion
}

// awsCodebuildWebhooks implements AwsCodebuildWebhookInterface
type awsCodebuildWebhooks struct {
	client rest.Interface
}

// newAwsCodebuildWebhooks returns a AwsCodebuildWebhooks
func newAwsCodebuildWebhooks(c *AwsV1alpha1Client) *awsCodebuildWebhooks {
	return &awsCodebuildWebhooks{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsCodebuildWebhook, and returns the corresponding awsCodebuildWebhook object, and an error if there is any.
func (c *awsCodebuildWebhooks) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsCodebuildWebhook, err error) {
	result = &v1alpha1.AwsCodebuildWebhook{}
	err = c.client.Get().
		Resource("awscodebuildwebhooks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsCodebuildWebhooks that match those selectors.
func (c *awsCodebuildWebhooks) List(opts v1.ListOptions) (result *v1alpha1.AwsCodebuildWebhookList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsCodebuildWebhookList{}
	err = c.client.Get().
		Resource("awscodebuildwebhooks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsCodebuildWebhooks.
func (c *awsCodebuildWebhooks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awscodebuildwebhooks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsCodebuildWebhook and creates it.  Returns the server's representation of the awsCodebuildWebhook, and an error, if there is any.
func (c *awsCodebuildWebhooks) Create(awsCodebuildWebhook *v1alpha1.AwsCodebuildWebhook) (result *v1alpha1.AwsCodebuildWebhook, err error) {
	result = &v1alpha1.AwsCodebuildWebhook{}
	err = c.client.Post().
		Resource("awscodebuildwebhooks").
		Body(awsCodebuildWebhook).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsCodebuildWebhook and updates it. Returns the server's representation of the awsCodebuildWebhook, and an error, if there is any.
func (c *awsCodebuildWebhooks) Update(awsCodebuildWebhook *v1alpha1.AwsCodebuildWebhook) (result *v1alpha1.AwsCodebuildWebhook, err error) {
	result = &v1alpha1.AwsCodebuildWebhook{}
	err = c.client.Put().
		Resource("awscodebuildwebhooks").
		Name(awsCodebuildWebhook.Name).
		Body(awsCodebuildWebhook).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsCodebuildWebhooks) UpdateStatus(awsCodebuildWebhook *v1alpha1.AwsCodebuildWebhook) (result *v1alpha1.AwsCodebuildWebhook, err error) {
	result = &v1alpha1.AwsCodebuildWebhook{}
	err = c.client.Put().
		Resource("awscodebuildwebhooks").
		Name(awsCodebuildWebhook.Name).
		SubResource("status").
		Body(awsCodebuildWebhook).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsCodebuildWebhook and deletes it. Returns an error if one occurs.
func (c *awsCodebuildWebhooks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awscodebuildwebhooks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsCodebuildWebhooks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awscodebuildwebhooks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsCodebuildWebhook.
func (c *awsCodebuildWebhooks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCodebuildWebhook, err error) {
	result = &v1alpha1.AwsCodebuildWebhook{}
	err = c.client.Patch(pt).
		Resource("awscodebuildwebhooks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}