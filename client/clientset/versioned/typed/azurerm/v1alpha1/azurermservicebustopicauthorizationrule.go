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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// AzurermServicebusTopicAuthorizationRulesGetter has a method to return a AzurermServicebusTopicAuthorizationRuleInterface.
// A group's client should implement this interface.
type AzurermServicebusTopicAuthorizationRulesGetter interface {
	AzurermServicebusTopicAuthorizationRules() AzurermServicebusTopicAuthorizationRuleInterface
}

// AzurermServicebusTopicAuthorizationRuleInterface has methods to work with AzurermServicebusTopicAuthorizationRule resources.
type AzurermServicebusTopicAuthorizationRuleInterface interface {
	Create(*v1alpha1.AzurermServicebusTopicAuthorizationRule) (*v1alpha1.AzurermServicebusTopicAuthorizationRule, error)
	Update(*v1alpha1.AzurermServicebusTopicAuthorizationRule) (*v1alpha1.AzurermServicebusTopicAuthorizationRule, error)
	UpdateStatus(*v1alpha1.AzurermServicebusTopicAuthorizationRule) (*v1alpha1.AzurermServicebusTopicAuthorizationRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermServicebusTopicAuthorizationRule, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermServicebusTopicAuthorizationRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermServicebusTopicAuthorizationRule, err error)
	AzurermServicebusTopicAuthorizationRuleExpansion
}

// azurermServicebusTopicAuthorizationRules implements AzurermServicebusTopicAuthorizationRuleInterface
type azurermServicebusTopicAuthorizationRules struct {
	client rest.Interface
}

// newAzurermServicebusTopicAuthorizationRules returns a AzurermServicebusTopicAuthorizationRules
func newAzurermServicebusTopicAuthorizationRules(c *AzurermV1alpha1Client) *azurermServicebusTopicAuthorizationRules {
	return &azurermServicebusTopicAuthorizationRules{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermServicebusTopicAuthorizationRule, and returns the corresponding azurermServicebusTopicAuthorizationRule object, and an error if there is any.
func (c *azurermServicebusTopicAuthorizationRules) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermServicebusTopicAuthorizationRule, err error) {
	result = &v1alpha1.AzurermServicebusTopicAuthorizationRule{}
	err = c.client.Get().
		Resource("azurermservicebustopicauthorizationrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermServicebusTopicAuthorizationRules that match those selectors.
func (c *azurermServicebusTopicAuthorizationRules) List(opts v1.ListOptions) (result *v1alpha1.AzurermServicebusTopicAuthorizationRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermServicebusTopicAuthorizationRuleList{}
	err = c.client.Get().
		Resource("azurermservicebustopicauthorizationrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermServicebusTopicAuthorizationRules.
func (c *azurermServicebusTopicAuthorizationRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermservicebustopicauthorizationrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermServicebusTopicAuthorizationRule and creates it.  Returns the server's representation of the azurermServicebusTopicAuthorizationRule, and an error, if there is any.
func (c *azurermServicebusTopicAuthorizationRules) Create(azurermServicebusTopicAuthorizationRule *v1alpha1.AzurermServicebusTopicAuthorizationRule) (result *v1alpha1.AzurermServicebusTopicAuthorizationRule, err error) {
	result = &v1alpha1.AzurermServicebusTopicAuthorizationRule{}
	err = c.client.Post().
		Resource("azurermservicebustopicauthorizationrules").
		Body(azurermServicebusTopicAuthorizationRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermServicebusTopicAuthorizationRule and updates it. Returns the server's representation of the azurermServicebusTopicAuthorizationRule, and an error, if there is any.
func (c *azurermServicebusTopicAuthorizationRules) Update(azurermServicebusTopicAuthorizationRule *v1alpha1.AzurermServicebusTopicAuthorizationRule) (result *v1alpha1.AzurermServicebusTopicAuthorizationRule, err error) {
	result = &v1alpha1.AzurermServicebusTopicAuthorizationRule{}
	err = c.client.Put().
		Resource("azurermservicebustopicauthorizationrules").
		Name(azurermServicebusTopicAuthorizationRule.Name).
		Body(azurermServicebusTopicAuthorizationRule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermServicebusTopicAuthorizationRules) UpdateStatus(azurermServicebusTopicAuthorizationRule *v1alpha1.AzurermServicebusTopicAuthorizationRule) (result *v1alpha1.AzurermServicebusTopicAuthorizationRule, err error) {
	result = &v1alpha1.AzurermServicebusTopicAuthorizationRule{}
	err = c.client.Put().
		Resource("azurermservicebustopicauthorizationrules").
		Name(azurermServicebusTopicAuthorizationRule.Name).
		SubResource("status").
		Body(azurermServicebusTopicAuthorizationRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermServicebusTopicAuthorizationRule and deletes it. Returns an error if one occurs.
func (c *azurermServicebusTopicAuthorizationRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermservicebustopicauthorizationrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermServicebusTopicAuthorizationRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermservicebustopicauthorizationrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermServicebusTopicAuthorizationRule.
func (c *azurermServicebusTopicAuthorizationRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermServicebusTopicAuthorizationRule, err error) {
	result = &v1alpha1.AzurermServicebusTopicAuthorizationRule{}
	err = c.client.Patch(pt).
		Resource("azurermservicebustopicauthorizationrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}