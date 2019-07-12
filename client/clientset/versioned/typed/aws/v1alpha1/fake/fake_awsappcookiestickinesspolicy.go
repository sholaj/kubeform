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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeAwsAppCookieStickinessPolicies implements AwsAppCookieStickinessPolicyInterface
type FakeAwsAppCookieStickinessPolicies struct {
	Fake *FakeAwsV1alpha1
}

var awsappcookiestickinesspoliciesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsappcookiestickinesspolicies"}

var awsappcookiestickinesspoliciesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsAppCookieStickinessPolicy"}

// Get takes name of the awsAppCookieStickinessPolicy, and returns the corresponding awsAppCookieStickinessPolicy object, and an error if there is any.
func (c *FakeAwsAppCookieStickinessPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsAppCookieStickinessPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsappcookiestickinesspoliciesResource, name), &v1alpha1.AwsAppCookieStickinessPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAppCookieStickinessPolicy), err
}

// List takes label and field selectors, and returns the list of AwsAppCookieStickinessPolicies that match those selectors.
func (c *FakeAwsAppCookieStickinessPolicies) List(opts v1.ListOptions) (result *v1alpha1.AwsAppCookieStickinessPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsappcookiestickinesspoliciesResource, awsappcookiestickinesspoliciesKind, opts), &v1alpha1.AwsAppCookieStickinessPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsAppCookieStickinessPolicyList{ListMeta: obj.(*v1alpha1.AwsAppCookieStickinessPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsAppCookieStickinessPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsAppCookieStickinessPolicies.
func (c *FakeAwsAppCookieStickinessPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsappcookiestickinesspoliciesResource, opts))
}

// Create takes the representation of a awsAppCookieStickinessPolicy and creates it.  Returns the server's representation of the awsAppCookieStickinessPolicy, and an error, if there is any.
func (c *FakeAwsAppCookieStickinessPolicies) Create(awsAppCookieStickinessPolicy *v1alpha1.AwsAppCookieStickinessPolicy) (result *v1alpha1.AwsAppCookieStickinessPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsappcookiestickinesspoliciesResource, awsAppCookieStickinessPolicy), &v1alpha1.AwsAppCookieStickinessPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAppCookieStickinessPolicy), err
}

// Update takes the representation of a awsAppCookieStickinessPolicy and updates it. Returns the server's representation of the awsAppCookieStickinessPolicy, and an error, if there is any.
func (c *FakeAwsAppCookieStickinessPolicies) Update(awsAppCookieStickinessPolicy *v1alpha1.AwsAppCookieStickinessPolicy) (result *v1alpha1.AwsAppCookieStickinessPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsappcookiestickinesspoliciesResource, awsAppCookieStickinessPolicy), &v1alpha1.AwsAppCookieStickinessPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAppCookieStickinessPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsAppCookieStickinessPolicies) UpdateStatus(awsAppCookieStickinessPolicy *v1alpha1.AwsAppCookieStickinessPolicy) (*v1alpha1.AwsAppCookieStickinessPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsappcookiestickinesspoliciesResource, "status", awsAppCookieStickinessPolicy), &v1alpha1.AwsAppCookieStickinessPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAppCookieStickinessPolicy), err
}

// Delete takes name of the awsAppCookieStickinessPolicy and deletes it. Returns an error if one occurs.
func (c *FakeAwsAppCookieStickinessPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsappcookiestickinesspoliciesResource, name), &v1alpha1.AwsAppCookieStickinessPolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsAppCookieStickinessPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsappcookiestickinesspoliciesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsAppCookieStickinessPolicyList{})
	return err
}

// Patch applies the patch and returns the patched awsAppCookieStickinessPolicy.
func (c *FakeAwsAppCookieStickinessPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAppCookieStickinessPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsappcookiestickinesspoliciesResource, name, pt, data, subresources...), &v1alpha1.AwsAppCookieStickinessPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAppCookieStickinessPolicy), err
}