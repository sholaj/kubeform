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

// FakeGoogleStorageBucketIamBindings implements GoogleStorageBucketIamBindingInterface
type FakeGoogleStorageBucketIamBindings struct {
	Fake *FakeGoogleV1alpha1
}

var googlestoragebucketiambindingsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlestoragebucketiambindings"}

var googlestoragebucketiambindingsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleStorageBucketIamBinding"}

// Get takes name of the googleStorageBucketIamBinding, and returns the corresponding googleStorageBucketIamBinding object, and an error if there is any.
func (c *FakeGoogleStorageBucketIamBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleStorageBucketIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlestoragebucketiambindingsResource, name), &v1alpha1.GoogleStorageBucketIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleStorageBucketIamBinding), err
}

// List takes label and field selectors, and returns the list of GoogleStorageBucketIamBindings that match those selectors.
func (c *FakeGoogleStorageBucketIamBindings) List(opts v1.ListOptions) (result *v1alpha1.GoogleStorageBucketIamBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlestoragebucketiambindingsResource, googlestoragebucketiambindingsKind, opts), &v1alpha1.GoogleStorageBucketIamBindingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleStorageBucketIamBindingList{ListMeta: obj.(*v1alpha1.GoogleStorageBucketIamBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleStorageBucketIamBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleStorageBucketIamBindings.
func (c *FakeGoogleStorageBucketIamBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlestoragebucketiambindingsResource, opts))
}

// Create takes the representation of a googleStorageBucketIamBinding and creates it.  Returns the server's representation of the googleStorageBucketIamBinding, and an error, if there is any.
func (c *FakeGoogleStorageBucketIamBindings) Create(googleStorageBucketIamBinding *v1alpha1.GoogleStorageBucketIamBinding) (result *v1alpha1.GoogleStorageBucketIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlestoragebucketiambindingsResource, googleStorageBucketIamBinding), &v1alpha1.GoogleStorageBucketIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleStorageBucketIamBinding), err
}

// Update takes the representation of a googleStorageBucketIamBinding and updates it. Returns the server's representation of the googleStorageBucketIamBinding, and an error, if there is any.
func (c *FakeGoogleStorageBucketIamBindings) Update(googleStorageBucketIamBinding *v1alpha1.GoogleStorageBucketIamBinding) (result *v1alpha1.GoogleStorageBucketIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlestoragebucketiambindingsResource, googleStorageBucketIamBinding), &v1alpha1.GoogleStorageBucketIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleStorageBucketIamBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleStorageBucketIamBindings) UpdateStatus(googleStorageBucketIamBinding *v1alpha1.GoogleStorageBucketIamBinding) (*v1alpha1.GoogleStorageBucketIamBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlestoragebucketiambindingsResource, "status", googleStorageBucketIamBinding), &v1alpha1.GoogleStorageBucketIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleStorageBucketIamBinding), err
}

// Delete takes name of the googleStorageBucketIamBinding and deletes it. Returns an error if one occurs.
func (c *FakeGoogleStorageBucketIamBindings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlestoragebucketiambindingsResource, name), &v1alpha1.GoogleStorageBucketIamBinding{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleStorageBucketIamBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlestoragebucketiambindingsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleStorageBucketIamBindingList{})
	return err
}

// Patch applies the patch and returns the patched googleStorageBucketIamBinding.
func (c *FakeGoogleStorageBucketIamBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleStorageBucketIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlestoragebucketiambindingsResource, name, pt, data, subresources...), &v1alpha1.GoogleStorageBucketIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleStorageBucketIamBinding), err
}