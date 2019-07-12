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
	v1alpha1 "kubeform.dev/kubeform/apis/digitalocean/v1alpha1"
)

// FakeDigitaloceanFirewalls implements DigitaloceanFirewallInterface
type FakeDigitaloceanFirewalls struct {
	Fake *FakeDigitaloceanV1alpha1
}

var digitaloceanfirewallsResource = schema.GroupVersionResource{Group: "digitalocean.kubeform.com", Version: "v1alpha1", Resource: "digitaloceanfirewalls"}

var digitaloceanfirewallsKind = schema.GroupVersionKind{Group: "digitalocean.kubeform.com", Version: "v1alpha1", Kind: "DigitaloceanFirewall"}

// Get takes name of the digitaloceanFirewall, and returns the corresponding digitaloceanFirewall object, and an error if there is any.
func (c *FakeDigitaloceanFirewalls) Get(name string, options v1.GetOptions) (result *v1alpha1.DigitaloceanFirewall, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(digitaloceanfirewallsResource, name), &v1alpha1.DigitaloceanFirewall{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanFirewall), err
}

// List takes label and field selectors, and returns the list of DigitaloceanFirewalls that match those selectors.
func (c *FakeDigitaloceanFirewalls) List(opts v1.ListOptions) (result *v1alpha1.DigitaloceanFirewallList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(digitaloceanfirewallsResource, digitaloceanfirewallsKind, opts), &v1alpha1.DigitaloceanFirewallList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DigitaloceanFirewallList{ListMeta: obj.(*v1alpha1.DigitaloceanFirewallList).ListMeta}
	for _, item := range obj.(*v1alpha1.DigitaloceanFirewallList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested digitaloceanFirewalls.
func (c *FakeDigitaloceanFirewalls) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(digitaloceanfirewallsResource, opts))
}

// Create takes the representation of a digitaloceanFirewall and creates it.  Returns the server's representation of the digitaloceanFirewall, and an error, if there is any.
func (c *FakeDigitaloceanFirewalls) Create(digitaloceanFirewall *v1alpha1.DigitaloceanFirewall) (result *v1alpha1.DigitaloceanFirewall, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(digitaloceanfirewallsResource, digitaloceanFirewall), &v1alpha1.DigitaloceanFirewall{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanFirewall), err
}

// Update takes the representation of a digitaloceanFirewall and updates it. Returns the server's representation of the digitaloceanFirewall, and an error, if there is any.
func (c *FakeDigitaloceanFirewalls) Update(digitaloceanFirewall *v1alpha1.DigitaloceanFirewall) (result *v1alpha1.DigitaloceanFirewall, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(digitaloceanfirewallsResource, digitaloceanFirewall), &v1alpha1.DigitaloceanFirewall{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanFirewall), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDigitaloceanFirewalls) UpdateStatus(digitaloceanFirewall *v1alpha1.DigitaloceanFirewall) (*v1alpha1.DigitaloceanFirewall, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(digitaloceanfirewallsResource, "status", digitaloceanFirewall), &v1alpha1.DigitaloceanFirewall{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanFirewall), err
}

// Delete takes name of the digitaloceanFirewall and deletes it. Returns an error if one occurs.
func (c *FakeDigitaloceanFirewalls) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(digitaloceanfirewallsResource, name), &v1alpha1.DigitaloceanFirewall{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDigitaloceanFirewalls) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(digitaloceanfirewallsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DigitaloceanFirewallList{})
	return err
}

// Patch applies the patch and returns the patched digitaloceanFirewall.
func (c *FakeDigitaloceanFirewalls) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DigitaloceanFirewall, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(digitaloceanfirewallsResource, name, pt, data, subresources...), &v1alpha1.DigitaloceanFirewall{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanFirewall), err
}