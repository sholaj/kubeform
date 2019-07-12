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

// AwsApiGatewayStagesGetter has a method to return a AwsApiGatewayStageInterface.
// A group's client should implement this interface.
type AwsApiGatewayStagesGetter interface {
	AwsApiGatewayStages() AwsApiGatewayStageInterface
}

// AwsApiGatewayStageInterface has methods to work with AwsApiGatewayStage resources.
type AwsApiGatewayStageInterface interface {
	Create(*v1alpha1.AwsApiGatewayStage) (*v1alpha1.AwsApiGatewayStage, error)
	Update(*v1alpha1.AwsApiGatewayStage) (*v1alpha1.AwsApiGatewayStage, error)
	UpdateStatus(*v1alpha1.AwsApiGatewayStage) (*v1alpha1.AwsApiGatewayStage, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsApiGatewayStage, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsApiGatewayStageList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayStage, err error)
	AwsApiGatewayStageExpansion
}

// awsApiGatewayStages implements AwsApiGatewayStageInterface
type awsApiGatewayStages struct {
	client rest.Interface
}

// newAwsApiGatewayStages returns a AwsApiGatewayStages
func newAwsApiGatewayStages(c *AwsV1alpha1Client) *awsApiGatewayStages {
	return &awsApiGatewayStages{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsApiGatewayStage, and returns the corresponding awsApiGatewayStage object, and an error if there is any.
func (c *awsApiGatewayStages) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsApiGatewayStage, err error) {
	result = &v1alpha1.AwsApiGatewayStage{}
	err = c.client.Get().
		Resource("awsapigatewaystages").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsApiGatewayStages that match those selectors.
func (c *awsApiGatewayStages) List(opts v1.ListOptions) (result *v1alpha1.AwsApiGatewayStageList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsApiGatewayStageList{}
	err = c.client.Get().
		Resource("awsapigatewaystages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsApiGatewayStages.
func (c *awsApiGatewayStages) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsapigatewaystages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsApiGatewayStage and creates it.  Returns the server's representation of the awsApiGatewayStage, and an error, if there is any.
func (c *awsApiGatewayStages) Create(awsApiGatewayStage *v1alpha1.AwsApiGatewayStage) (result *v1alpha1.AwsApiGatewayStage, err error) {
	result = &v1alpha1.AwsApiGatewayStage{}
	err = c.client.Post().
		Resource("awsapigatewaystages").
		Body(awsApiGatewayStage).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsApiGatewayStage and updates it. Returns the server's representation of the awsApiGatewayStage, and an error, if there is any.
func (c *awsApiGatewayStages) Update(awsApiGatewayStage *v1alpha1.AwsApiGatewayStage) (result *v1alpha1.AwsApiGatewayStage, err error) {
	result = &v1alpha1.AwsApiGatewayStage{}
	err = c.client.Put().
		Resource("awsapigatewaystages").
		Name(awsApiGatewayStage.Name).
		Body(awsApiGatewayStage).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsApiGatewayStages) UpdateStatus(awsApiGatewayStage *v1alpha1.AwsApiGatewayStage) (result *v1alpha1.AwsApiGatewayStage, err error) {
	result = &v1alpha1.AwsApiGatewayStage{}
	err = c.client.Put().
		Resource("awsapigatewaystages").
		Name(awsApiGatewayStage.Name).
		SubResource("status").
		Body(awsApiGatewayStage).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsApiGatewayStage and deletes it. Returns an error if one occurs.
func (c *awsApiGatewayStages) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsapigatewaystages").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsApiGatewayStages) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsapigatewaystages").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsApiGatewayStage.
func (c *awsApiGatewayStages) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayStage, err error) {
	result = &v1alpha1.AwsApiGatewayStage{}
	err = c.client.Patch(pt).
		Resource("awsapigatewaystages").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}