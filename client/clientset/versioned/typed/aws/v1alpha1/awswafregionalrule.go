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

// AwsWafregionalRulesGetter has a method to return a AwsWafregionalRuleInterface.
// A group's client should implement this interface.
type AwsWafregionalRulesGetter interface {
	AwsWafregionalRules() AwsWafregionalRuleInterface
}

// AwsWafregionalRuleInterface has methods to work with AwsWafregionalRule resources.
type AwsWafregionalRuleInterface interface {
	Create(*v1alpha1.AwsWafregionalRule) (*v1alpha1.AwsWafregionalRule, error)
	Update(*v1alpha1.AwsWafregionalRule) (*v1alpha1.AwsWafregionalRule, error)
	UpdateStatus(*v1alpha1.AwsWafregionalRule) (*v1alpha1.AwsWafregionalRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsWafregionalRule, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsWafregionalRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafregionalRule, err error)
	AwsWafregionalRuleExpansion
}

// awsWafregionalRules implements AwsWafregionalRuleInterface
type awsWafregionalRules struct {
	client rest.Interface
}

// newAwsWafregionalRules returns a AwsWafregionalRules
func newAwsWafregionalRules(c *AwsV1alpha1Client) *awsWafregionalRules {
	return &awsWafregionalRules{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsWafregionalRule, and returns the corresponding awsWafregionalRule object, and an error if there is any.
func (c *awsWafregionalRules) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsWafregionalRule, err error) {
	result = &v1alpha1.AwsWafregionalRule{}
	err = c.client.Get().
		Resource("awswafregionalrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsWafregionalRules that match those selectors.
func (c *awsWafregionalRules) List(opts v1.ListOptions) (result *v1alpha1.AwsWafregionalRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsWafregionalRuleList{}
	err = c.client.Get().
		Resource("awswafregionalrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsWafregionalRules.
func (c *awsWafregionalRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awswafregionalrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsWafregionalRule and creates it.  Returns the server's representation of the awsWafregionalRule, and an error, if there is any.
func (c *awsWafregionalRules) Create(awsWafregionalRule *v1alpha1.AwsWafregionalRule) (result *v1alpha1.AwsWafregionalRule, err error) {
	result = &v1alpha1.AwsWafregionalRule{}
	err = c.client.Post().
		Resource("awswafregionalrules").
		Body(awsWafregionalRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsWafregionalRule and updates it. Returns the server's representation of the awsWafregionalRule, and an error, if there is any.
func (c *awsWafregionalRules) Update(awsWafregionalRule *v1alpha1.AwsWafregionalRule) (result *v1alpha1.AwsWafregionalRule, err error) {
	result = &v1alpha1.AwsWafregionalRule{}
	err = c.client.Put().
		Resource("awswafregionalrules").
		Name(awsWafregionalRule.Name).
		Body(awsWafregionalRule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsWafregionalRules) UpdateStatus(awsWafregionalRule *v1alpha1.AwsWafregionalRule) (result *v1alpha1.AwsWafregionalRule, err error) {
	result = &v1alpha1.AwsWafregionalRule{}
	err = c.client.Put().
		Resource("awswafregionalrules").
		Name(awsWafregionalRule.Name).
		SubResource("status").
		Body(awsWafregionalRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsWafregionalRule and deletes it. Returns an error if one occurs.
func (c *awsWafregionalRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awswafregionalrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsWafregionalRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awswafregionalrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsWafregionalRule.
func (c *awsWafregionalRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafregionalRule, err error) {
	result = &v1alpha1.AwsWafregionalRule{}
	err = c.client.Patch(pt).
		Resource("awswafregionalrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}