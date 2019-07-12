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

// AwsOpsworksJavaAppLayersGetter has a method to return a AwsOpsworksJavaAppLayerInterface.
// A group's client should implement this interface.
type AwsOpsworksJavaAppLayersGetter interface {
	AwsOpsworksJavaAppLayers() AwsOpsworksJavaAppLayerInterface
}

// AwsOpsworksJavaAppLayerInterface has methods to work with AwsOpsworksJavaAppLayer resources.
type AwsOpsworksJavaAppLayerInterface interface {
	Create(*v1alpha1.AwsOpsworksJavaAppLayer) (*v1alpha1.AwsOpsworksJavaAppLayer, error)
	Update(*v1alpha1.AwsOpsworksJavaAppLayer) (*v1alpha1.AwsOpsworksJavaAppLayer, error)
	UpdateStatus(*v1alpha1.AwsOpsworksJavaAppLayer) (*v1alpha1.AwsOpsworksJavaAppLayer, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsOpsworksJavaAppLayer, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsOpsworksJavaAppLayerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsOpsworksJavaAppLayer, err error)
	AwsOpsworksJavaAppLayerExpansion
}

// awsOpsworksJavaAppLayers implements AwsOpsworksJavaAppLayerInterface
type awsOpsworksJavaAppLayers struct {
	client rest.Interface
}

// newAwsOpsworksJavaAppLayers returns a AwsOpsworksJavaAppLayers
func newAwsOpsworksJavaAppLayers(c *AwsV1alpha1Client) *awsOpsworksJavaAppLayers {
	return &awsOpsworksJavaAppLayers{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsOpsworksJavaAppLayer, and returns the corresponding awsOpsworksJavaAppLayer object, and an error if there is any.
func (c *awsOpsworksJavaAppLayers) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsOpsworksJavaAppLayer, err error) {
	result = &v1alpha1.AwsOpsworksJavaAppLayer{}
	err = c.client.Get().
		Resource("awsopsworksjavaapplayers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsOpsworksJavaAppLayers that match those selectors.
func (c *awsOpsworksJavaAppLayers) List(opts v1.ListOptions) (result *v1alpha1.AwsOpsworksJavaAppLayerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsOpsworksJavaAppLayerList{}
	err = c.client.Get().
		Resource("awsopsworksjavaapplayers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsOpsworksJavaAppLayers.
func (c *awsOpsworksJavaAppLayers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsopsworksjavaapplayers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsOpsworksJavaAppLayer and creates it.  Returns the server's representation of the awsOpsworksJavaAppLayer, and an error, if there is any.
func (c *awsOpsworksJavaAppLayers) Create(awsOpsworksJavaAppLayer *v1alpha1.AwsOpsworksJavaAppLayer) (result *v1alpha1.AwsOpsworksJavaAppLayer, err error) {
	result = &v1alpha1.AwsOpsworksJavaAppLayer{}
	err = c.client.Post().
		Resource("awsopsworksjavaapplayers").
		Body(awsOpsworksJavaAppLayer).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsOpsworksJavaAppLayer and updates it. Returns the server's representation of the awsOpsworksJavaAppLayer, and an error, if there is any.
func (c *awsOpsworksJavaAppLayers) Update(awsOpsworksJavaAppLayer *v1alpha1.AwsOpsworksJavaAppLayer) (result *v1alpha1.AwsOpsworksJavaAppLayer, err error) {
	result = &v1alpha1.AwsOpsworksJavaAppLayer{}
	err = c.client.Put().
		Resource("awsopsworksjavaapplayers").
		Name(awsOpsworksJavaAppLayer.Name).
		Body(awsOpsworksJavaAppLayer).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsOpsworksJavaAppLayers) UpdateStatus(awsOpsworksJavaAppLayer *v1alpha1.AwsOpsworksJavaAppLayer) (result *v1alpha1.AwsOpsworksJavaAppLayer, err error) {
	result = &v1alpha1.AwsOpsworksJavaAppLayer{}
	err = c.client.Put().
		Resource("awsopsworksjavaapplayers").
		Name(awsOpsworksJavaAppLayer.Name).
		SubResource("status").
		Body(awsOpsworksJavaAppLayer).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsOpsworksJavaAppLayer and deletes it. Returns an error if one occurs.
func (c *awsOpsworksJavaAppLayers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsopsworksjavaapplayers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsOpsworksJavaAppLayers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsopsworksjavaapplayers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsOpsworksJavaAppLayer.
func (c *awsOpsworksJavaAppLayers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsOpsworksJavaAppLayer, err error) {
	result = &v1alpha1.AwsOpsworksJavaAppLayer{}
	err = c.client.Patch(pt).
		Resource("awsopsworksjavaapplayers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}