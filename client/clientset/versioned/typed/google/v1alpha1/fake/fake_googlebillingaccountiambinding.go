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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeGoogleBillingAccountIamBindings implements GoogleBillingAccountIamBindingInterface
type FakeGoogleBillingAccountIamBindings struct {
	Fake *FakeGoogleV1alpha1
}

var googlebillingaccountiambindingsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlebillingaccountiambindings"}

var googlebillingaccountiambindingsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleBillingAccountIamBinding"}

// Get takes name of the googleBillingAccountIamBinding, and returns the corresponding googleBillingAccountIamBinding object, and an error if there is any.
func (c *FakeGoogleBillingAccountIamBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleBillingAccountIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlebillingaccountiambindingsResource, name), &v1alpha1.GoogleBillingAccountIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleBillingAccountIamBinding), err
}

// List takes label and field selectors, and returns the list of GoogleBillingAccountIamBindings that match those selectors.
func (c *FakeGoogleBillingAccountIamBindings) List(opts v1.ListOptions) (result *v1alpha1.GoogleBillingAccountIamBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlebillingaccountiambindingsResource, googlebillingaccountiambindingsKind, opts), &v1alpha1.GoogleBillingAccountIamBindingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleBillingAccountIamBindingList{ListMeta: obj.(*v1alpha1.GoogleBillingAccountIamBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleBillingAccountIamBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleBillingAccountIamBindings.
func (c *FakeGoogleBillingAccountIamBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlebillingaccountiambindingsResource, opts))
}

// Create takes the representation of a googleBillingAccountIamBinding and creates it.  Returns the server's representation of the googleBillingAccountIamBinding, and an error, if there is any.
func (c *FakeGoogleBillingAccountIamBindings) Create(googleBillingAccountIamBinding *v1alpha1.GoogleBillingAccountIamBinding) (result *v1alpha1.GoogleBillingAccountIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlebillingaccountiambindingsResource, googleBillingAccountIamBinding), &v1alpha1.GoogleBillingAccountIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleBillingAccountIamBinding), err
}

// Update takes the representation of a googleBillingAccountIamBinding and updates it. Returns the server's representation of the googleBillingAccountIamBinding, and an error, if there is any.
func (c *FakeGoogleBillingAccountIamBindings) Update(googleBillingAccountIamBinding *v1alpha1.GoogleBillingAccountIamBinding) (result *v1alpha1.GoogleBillingAccountIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlebillingaccountiambindingsResource, googleBillingAccountIamBinding), &v1alpha1.GoogleBillingAccountIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleBillingAccountIamBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleBillingAccountIamBindings) UpdateStatus(googleBillingAccountIamBinding *v1alpha1.GoogleBillingAccountIamBinding) (*v1alpha1.GoogleBillingAccountIamBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlebillingaccountiambindingsResource, "status", googleBillingAccountIamBinding), &v1alpha1.GoogleBillingAccountIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleBillingAccountIamBinding), err
}

// Delete takes name of the googleBillingAccountIamBinding and deletes it. Returns an error if one occurs.
func (c *FakeGoogleBillingAccountIamBindings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlebillingaccountiambindingsResource, name), &v1alpha1.GoogleBillingAccountIamBinding{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleBillingAccountIamBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlebillingaccountiambindingsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleBillingAccountIamBindingList{})
	return err
}

// Patch applies the patch and returns the patched googleBillingAccountIamBinding.
func (c *FakeGoogleBillingAccountIamBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleBillingAccountIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlebillingaccountiambindingsResource, name, pt, data, subresources...), &v1alpha1.GoogleBillingAccountIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleBillingAccountIamBinding), err
}