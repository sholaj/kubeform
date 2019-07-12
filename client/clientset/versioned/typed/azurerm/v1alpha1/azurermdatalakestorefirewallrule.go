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

// AzurermDataLakeStoreFirewallRulesGetter has a method to return a AzurermDataLakeStoreFirewallRuleInterface.
// A group's client should implement this interface.
type AzurermDataLakeStoreFirewallRulesGetter interface {
	AzurermDataLakeStoreFirewallRules() AzurermDataLakeStoreFirewallRuleInterface
}

// AzurermDataLakeStoreFirewallRuleInterface has methods to work with AzurermDataLakeStoreFirewallRule resources.
type AzurermDataLakeStoreFirewallRuleInterface interface {
	Create(*v1alpha1.AzurermDataLakeStoreFirewallRule) (*v1alpha1.AzurermDataLakeStoreFirewallRule, error)
	Update(*v1alpha1.AzurermDataLakeStoreFirewallRule) (*v1alpha1.AzurermDataLakeStoreFirewallRule, error)
	UpdateStatus(*v1alpha1.AzurermDataLakeStoreFirewallRule) (*v1alpha1.AzurermDataLakeStoreFirewallRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermDataLakeStoreFirewallRule, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermDataLakeStoreFirewallRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDataLakeStoreFirewallRule, err error)
	AzurermDataLakeStoreFirewallRuleExpansion
}

// azurermDataLakeStoreFirewallRules implements AzurermDataLakeStoreFirewallRuleInterface
type azurermDataLakeStoreFirewallRules struct {
	client rest.Interface
}

// newAzurermDataLakeStoreFirewallRules returns a AzurermDataLakeStoreFirewallRules
func newAzurermDataLakeStoreFirewallRules(c *AzurermV1alpha1Client) *azurermDataLakeStoreFirewallRules {
	return &azurermDataLakeStoreFirewallRules{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermDataLakeStoreFirewallRule, and returns the corresponding azurermDataLakeStoreFirewallRule object, and an error if there is any.
func (c *azurermDataLakeStoreFirewallRules) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermDataLakeStoreFirewallRule, err error) {
	result = &v1alpha1.AzurermDataLakeStoreFirewallRule{}
	err = c.client.Get().
		Resource("azurermdatalakestorefirewallrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermDataLakeStoreFirewallRules that match those selectors.
func (c *azurermDataLakeStoreFirewallRules) List(opts v1.ListOptions) (result *v1alpha1.AzurermDataLakeStoreFirewallRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermDataLakeStoreFirewallRuleList{}
	err = c.client.Get().
		Resource("azurermdatalakestorefirewallrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermDataLakeStoreFirewallRules.
func (c *azurermDataLakeStoreFirewallRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermdatalakestorefirewallrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermDataLakeStoreFirewallRule and creates it.  Returns the server's representation of the azurermDataLakeStoreFirewallRule, and an error, if there is any.
func (c *azurermDataLakeStoreFirewallRules) Create(azurermDataLakeStoreFirewallRule *v1alpha1.AzurermDataLakeStoreFirewallRule) (result *v1alpha1.AzurermDataLakeStoreFirewallRule, err error) {
	result = &v1alpha1.AzurermDataLakeStoreFirewallRule{}
	err = c.client.Post().
		Resource("azurermdatalakestorefirewallrules").
		Body(azurermDataLakeStoreFirewallRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermDataLakeStoreFirewallRule and updates it. Returns the server's representation of the azurermDataLakeStoreFirewallRule, and an error, if there is any.
func (c *azurermDataLakeStoreFirewallRules) Update(azurermDataLakeStoreFirewallRule *v1alpha1.AzurermDataLakeStoreFirewallRule) (result *v1alpha1.AzurermDataLakeStoreFirewallRule, err error) {
	result = &v1alpha1.AzurermDataLakeStoreFirewallRule{}
	err = c.client.Put().
		Resource("azurermdatalakestorefirewallrules").
		Name(azurermDataLakeStoreFirewallRule.Name).
		Body(azurermDataLakeStoreFirewallRule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermDataLakeStoreFirewallRules) UpdateStatus(azurermDataLakeStoreFirewallRule *v1alpha1.AzurermDataLakeStoreFirewallRule) (result *v1alpha1.AzurermDataLakeStoreFirewallRule, err error) {
	result = &v1alpha1.AzurermDataLakeStoreFirewallRule{}
	err = c.client.Put().
		Resource("azurermdatalakestorefirewallrules").
		Name(azurermDataLakeStoreFirewallRule.Name).
		SubResource("status").
		Body(azurermDataLakeStoreFirewallRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermDataLakeStoreFirewallRule and deletes it. Returns an error if one occurs.
func (c *azurermDataLakeStoreFirewallRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermdatalakestorefirewallrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermDataLakeStoreFirewallRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermdatalakestorefirewallrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermDataLakeStoreFirewallRule.
func (c *azurermDataLakeStoreFirewallRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDataLakeStoreFirewallRule, err error) {
	result = &v1alpha1.AzurermDataLakeStoreFirewallRule{}
	err = c.client.Patch(pt).
		Resource("azurermdatalakestorefirewallrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}