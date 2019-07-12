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

// AzurermEventhubNamespaceAuthorizationRulesGetter has a method to return a AzurermEventhubNamespaceAuthorizationRuleInterface.
// A group's client should implement this interface.
type AzurermEventhubNamespaceAuthorizationRulesGetter interface {
	AzurermEventhubNamespaceAuthorizationRules() AzurermEventhubNamespaceAuthorizationRuleInterface
}

// AzurermEventhubNamespaceAuthorizationRuleInterface has methods to work with AzurermEventhubNamespaceAuthorizationRule resources.
type AzurermEventhubNamespaceAuthorizationRuleInterface interface {
	Create(*v1alpha1.AzurermEventhubNamespaceAuthorizationRule) (*v1alpha1.AzurermEventhubNamespaceAuthorizationRule, error)
	Update(*v1alpha1.AzurermEventhubNamespaceAuthorizationRule) (*v1alpha1.AzurermEventhubNamespaceAuthorizationRule, error)
	UpdateStatus(*v1alpha1.AzurermEventhubNamespaceAuthorizationRule) (*v1alpha1.AzurermEventhubNamespaceAuthorizationRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermEventhubNamespaceAuthorizationRule, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermEventhubNamespaceAuthorizationRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermEventhubNamespaceAuthorizationRule, err error)
	AzurermEventhubNamespaceAuthorizationRuleExpansion
}

// azurermEventhubNamespaceAuthorizationRules implements AzurermEventhubNamespaceAuthorizationRuleInterface
type azurermEventhubNamespaceAuthorizationRules struct {
	client rest.Interface
}

// newAzurermEventhubNamespaceAuthorizationRules returns a AzurermEventhubNamespaceAuthorizationRules
func newAzurermEventhubNamespaceAuthorizationRules(c *AzurermV1alpha1Client) *azurermEventhubNamespaceAuthorizationRules {
	return &azurermEventhubNamespaceAuthorizationRules{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermEventhubNamespaceAuthorizationRule, and returns the corresponding azurermEventhubNamespaceAuthorizationRule object, and an error if there is any.
func (c *azurermEventhubNamespaceAuthorizationRules) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermEventhubNamespaceAuthorizationRule, err error) {
	result = &v1alpha1.AzurermEventhubNamespaceAuthorizationRule{}
	err = c.client.Get().
		Resource("azurermeventhubnamespaceauthorizationrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermEventhubNamespaceAuthorizationRules that match those selectors.
func (c *azurermEventhubNamespaceAuthorizationRules) List(opts v1.ListOptions) (result *v1alpha1.AzurermEventhubNamespaceAuthorizationRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermEventhubNamespaceAuthorizationRuleList{}
	err = c.client.Get().
		Resource("azurermeventhubnamespaceauthorizationrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermEventhubNamespaceAuthorizationRules.
func (c *azurermEventhubNamespaceAuthorizationRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermeventhubnamespaceauthorizationrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermEventhubNamespaceAuthorizationRule and creates it.  Returns the server's representation of the azurermEventhubNamespaceAuthorizationRule, and an error, if there is any.
func (c *azurermEventhubNamespaceAuthorizationRules) Create(azurermEventhubNamespaceAuthorizationRule *v1alpha1.AzurermEventhubNamespaceAuthorizationRule) (result *v1alpha1.AzurermEventhubNamespaceAuthorizationRule, err error) {
	result = &v1alpha1.AzurermEventhubNamespaceAuthorizationRule{}
	err = c.client.Post().
		Resource("azurermeventhubnamespaceauthorizationrules").
		Body(azurermEventhubNamespaceAuthorizationRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermEventhubNamespaceAuthorizationRule and updates it. Returns the server's representation of the azurermEventhubNamespaceAuthorizationRule, and an error, if there is any.
func (c *azurermEventhubNamespaceAuthorizationRules) Update(azurermEventhubNamespaceAuthorizationRule *v1alpha1.AzurermEventhubNamespaceAuthorizationRule) (result *v1alpha1.AzurermEventhubNamespaceAuthorizationRule, err error) {
	result = &v1alpha1.AzurermEventhubNamespaceAuthorizationRule{}
	err = c.client.Put().
		Resource("azurermeventhubnamespaceauthorizationrules").
		Name(azurermEventhubNamespaceAuthorizationRule.Name).
		Body(azurermEventhubNamespaceAuthorizationRule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermEventhubNamespaceAuthorizationRules) UpdateStatus(azurermEventhubNamespaceAuthorizationRule *v1alpha1.AzurermEventhubNamespaceAuthorizationRule) (result *v1alpha1.AzurermEventhubNamespaceAuthorizationRule, err error) {
	result = &v1alpha1.AzurermEventhubNamespaceAuthorizationRule{}
	err = c.client.Put().
		Resource("azurermeventhubnamespaceauthorizationrules").
		Name(azurermEventhubNamespaceAuthorizationRule.Name).
		SubResource("status").
		Body(azurermEventhubNamespaceAuthorizationRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermEventhubNamespaceAuthorizationRule and deletes it. Returns an error if one occurs.
func (c *azurermEventhubNamespaceAuthorizationRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermeventhubnamespaceauthorizationrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermEventhubNamespaceAuthorizationRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermeventhubnamespaceauthorizationrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermEventhubNamespaceAuthorizationRule.
func (c *azurermEventhubNamespaceAuthorizationRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermEventhubNamespaceAuthorizationRule, err error) {
	result = &v1alpha1.AzurermEventhubNamespaceAuthorizationRule{}
	err = c.client.Patch(pt).
		Resource("azurermeventhubnamespaceauthorizationrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}