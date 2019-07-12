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

// FakeAwsLoadBalancerBackendServerPolicies implements AwsLoadBalancerBackendServerPolicyInterface
type FakeAwsLoadBalancerBackendServerPolicies struct {
	Fake *FakeAwsV1alpha1
}

var awsloadbalancerbackendserverpoliciesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsloadbalancerbackendserverpolicies"}

var awsloadbalancerbackendserverpoliciesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsLoadBalancerBackendServerPolicy"}

// Get takes name of the awsLoadBalancerBackendServerPolicy, and returns the corresponding awsLoadBalancerBackendServerPolicy object, and an error if there is any.
func (c *FakeAwsLoadBalancerBackendServerPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsLoadBalancerBackendServerPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsloadbalancerbackendserverpoliciesResource, name), &v1alpha1.AwsLoadBalancerBackendServerPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLoadBalancerBackendServerPolicy), err
}

// List takes label and field selectors, and returns the list of AwsLoadBalancerBackendServerPolicies that match those selectors.
func (c *FakeAwsLoadBalancerBackendServerPolicies) List(opts v1.ListOptions) (result *v1alpha1.AwsLoadBalancerBackendServerPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsloadbalancerbackendserverpoliciesResource, awsloadbalancerbackendserverpoliciesKind, opts), &v1alpha1.AwsLoadBalancerBackendServerPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsLoadBalancerBackendServerPolicyList{ListMeta: obj.(*v1alpha1.AwsLoadBalancerBackendServerPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsLoadBalancerBackendServerPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsLoadBalancerBackendServerPolicies.
func (c *FakeAwsLoadBalancerBackendServerPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsloadbalancerbackendserverpoliciesResource, opts))
}

// Create takes the representation of a awsLoadBalancerBackendServerPolicy and creates it.  Returns the server's representation of the awsLoadBalancerBackendServerPolicy, and an error, if there is any.
func (c *FakeAwsLoadBalancerBackendServerPolicies) Create(awsLoadBalancerBackendServerPolicy *v1alpha1.AwsLoadBalancerBackendServerPolicy) (result *v1alpha1.AwsLoadBalancerBackendServerPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsloadbalancerbackendserverpoliciesResource, awsLoadBalancerBackendServerPolicy), &v1alpha1.AwsLoadBalancerBackendServerPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLoadBalancerBackendServerPolicy), err
}

// Update takes the representation of a awsLoadBalancerBackendServerPolicy and updates it. Returns the server's representation of the awsLoadBalancerBackendServerPolicy, and an error, if there is any.
func (c *FakeAwsLoadBalancerBackendServerPolicies) Update(awsLoadBalancerBackendServerPolicy *v1alpha1.AwsLoadBalancerBackendServerPolicy) (result *v1alpha1.AwsLoadBalancerBackendServerPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsloadbalancerbackendserverpoliciesResource, awsLoadBalancerBackendServerPolicy), &v1alpha1.AwsLoadBalancerBackendServerPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLoadBalancerBackendServerPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsLoadBalancerBackendServerPolicies) UpdateStatus(awsLoadBalancerBackendServerPolicy *v1alpha1.AwsLoadBalancerBackendServerPolicy) (*v1alpha1.AwsLoadBalancerBackendServerPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsloadbalancerbackendserverpoliciesResource, "status", awsLoadBalancerBackendServerPolicy), &v1alpha1.AwsLoadBalancerBackendServerPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLoadBalancerBackendServerPolicy), err
}

// Delete takes name of the awsLoadBalancerBackendServerPolicy and deletes it. Returns an error if one occurs.
func (c *FakeAwsLoadBalancerBackendServerPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsloadbalancerbackendserverpoliciesResource, name), &v1alpha1.AwsLoadBalancerBackendServerPolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsLoadBalancerBackendServerPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsloadbalancerbackendserverpoliciesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsLoadBalancerBackendServerPolicyList{})
	return err
}

// Patch applies the patch and returns the patched awsLoadBalancerBackendServerPolicy.
func (c *FakeAwsLoadBalancerBackendServerPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsLoadBalancerBackendServerPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsloadbalancerbackendserverpoliciesResource, name, pt, data, subresources...), &v1alpha1.AwsLoadBalancerBackendServerPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLoadBalancerBackendServerPolicy), err
}