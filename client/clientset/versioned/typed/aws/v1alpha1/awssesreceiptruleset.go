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

// AwsSesReceiptRuleSetsGetter has a method to return a AwsSesReceiptRuleSetInterface.
// A group's client should implement this interface.
type AwsSesReceiptRuleSetsGetter interface {
	AwsSesReceiptRuleSets() AwsSesReceiptRuleSetInterface
}

// AwsSesReceiptRuleSetInterface has methods to work with AwsSesReceiptRuleSet resources.
type AwsSesReceiptRuleSetInterface interface {
	Create(*v1alpha1.AwsSesReceiptRuleSet) (*v1alpha1.AwsSesReceiptRuleSet, error)
	Update(*v1alpha1.AwsSesReceiptRuleSet) (*v1alpha1.AwsSesReceiptRuleSet, error)
	UpdateStatus(*v1alpha1.AwsSesReceiptRuleSet) (*v1alpha1.AwsSesReceiptRuleSet, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsSesReceiptRuleSet, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsSesReceiptRuleSetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSesReceiptRuleSet, err error)
	AwsSesReceiptRuleSetExpansion
}

// awsSesReceiptRuleSets implements AwsSesReceiptRuleSetInterface
type awsSesReceiptRuleSets struct {
	client rest.Interface
}

// newAwsSesReceiptRuleSets returns a AwsSesReceiptRuleSets
func newAwsSesReceiptRuleSets(c *AwsV1alpha1Client) *awsSesReceiptRuleSets {
	return &awsSesReceiptRuleSets{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsSesReceiptRuleSet, and returns the corresponding awsSesReceiptRuleSet object, and an error if there is any.
func (c *awsSesReceiptRuleSets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSesReceiptRuleSet, err error) {
	result = &v1alpha1.AwsSesReceiptRuleSet{}
	err = c.client.Get().
		Resource("awssesreceiptrulesets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsSesReceiptRuleSets that match those selectors.
func (c *awsSesReceiptRuleSets) List(opts v1.ListOptions) (result *v1alpha1.AwsSesReceiptRuleSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsSesReceiptRuleSetList{}
	err = c.client.Get().
		Resource("awssesreceiptrulesets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsSesReceiptRuleSets.
func (c *awsSesReceiptRuleSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awssesreceiptrulesets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsSesReceiptRuleSet and creates it.  Returns the server's representation of the awsSesReceiptRuleSet, and an error, if there is any.
func (c *awsSesReceiptRuleSets) Create(awsSesReceiptRuleSet *v1alpha1.AwsSesReceiptRuleSet) (result *v1alpha1.AwsSesReceiptRuleSet, err error) {
	result = &v1alpha1.AwsSesReceiptRuleSet{}
	err = c.client.Post().
		Resource("awssesreceiptrulesets").
		Body(awsSesReceiptRuleSet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsSesReceiptRuleSet and updates it. Returns the server's representation of the awsSesReceiptRuleSet, and an error, if there is any.
func (c *awsSesReceiptRuleSets) Update(awsSesReceiptRuleSet *v1alpha1.AwsSesReceiptRuleSet) (result *v1alpha1.AwsSesReceiptRuleSet, err error) {
	result = &v1alpha1.AwsSesReceiptRuleSet{}
	err = c.client.Put().
		Resource("awssesreceiptrulesets").
		Name(awsSesReceiptRuleSet.Name).
		Body(awsSesReceiptRuleSet).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsSesReceiptRuleSets) UpdateStatus(awsSesReceiptRuleSet *v1alpha1.AwsSesReceiptRuleSet) (result *v1alpha1.AwsSesReceiptRuleSet, err error) {
	result = &v1alpha1.AwsSesReceiptRuleSet{}
	err = c.client.Put().
		Resource("awssesreceiptrulesets").
		Name(awsSesReceiptRuleSet.Name).
		SubResource("status").
		Body(awsSesReceiptRuleSet).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsSesReceiptRuleSet and deletes it. Returns an error if one occurs.
func (c *awsSesReceiptRuleSets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awssesreceiptrulesets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsSesReceiptRuleSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awssesreceiptrulesets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsSesReceiptRuleSet.
func (c *awsSesReceiptRuleSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSesReceiptRuleSet, err error) {
	result = &v1alpha1.AwsSesReceiptRuleSet{}
	err = c.client.Patch(pt).
		Resource("awssesreceiptrulesets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}