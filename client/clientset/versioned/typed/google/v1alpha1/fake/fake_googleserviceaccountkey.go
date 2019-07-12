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

// FakeGoogleServiceAccountKeys implements GoogleServiceAccountKeyInterface
type FakeGoogleServiceAccountKeys struct {
	Fake *FakeGoogleV1alpha1
}

var googleserviceaccountkeysResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googleserviceaccountkeys"}

var googleserviceaccountkeysKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleServiceAccountKey"}

// Get takes name of the googleServiceAccountKey, and returns the corresponding googleServiceAccountKey object, and an error if there is any.
func (c *FakeGoogleServiceAccountKeys) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleServiceAccountKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googleserviceaccountkeysResource, name), &v1alpha1.GoogleServiceAccountKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleServiceAccountKey), err
}

// List takes label and field selectors, and returns the list of GoogleServiceAccountKeys that match those selectors.
func (c *FakeGoogleServiceAccountKeys) List(opts v1.ListOptions) (result *v1alpha1.GoogleServiceAccountKeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googleserviceaccountkeysResource, googleserviceaccountkeysKind, opts), &v1alpha1.GoogleServiceAccountKeyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleServiceAccountKeyList{ListMeta: obj.(*v1alpha1.GoogleServiceAccountKeyList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleServiceAccountKeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleServiceAccountKeys.
func (c *FakeGoogleServiceAccountKeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googleserviceaccountkeysResource, opts))
}

// Create takes the representation of a googleServiceAccountKey and creates it.  Returns the server's representation of the googleServiceAccountKey, and an error, if there is any.
func (c *FakeGoogleServiceAccountKeys) Create(googleServiceAccountKey *v1alpha1.GoogleServiceAccountKey) (result *v1alpha1.GoogleServiceAccountKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googleserviceaccountkeysResource, googleServiceAccountKey), &v1alpha1.GoogleServiceAccountKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleServiceAccountKey), err
}

// Update takes the representation of a googleServiceAccountKey and updates it. Returns the server's representation of the googleServiceAccountKey, and an error, if there is any.
func (c *FakeGoogleServiceAccountKeys) Update(googleServiceAccountKey *v1alpha1.GoogleServiceAccountKey) (result *v1alpha1.GoogleServiceAccountKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googleserviceaccountkeysResource, googleServiceAccountKey), &v1alpha1.GoogleServiceAccountKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleServiceAccountKey), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleServiceAccountKeys) UpdateStatus(googleServiceAccountKey *v1alpha1.GoogleServiceAccountKey) (*v1alpha1.GoogleServiceAccountKey, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googleserviceaccountkeysResource, "status", googleServiceAccountKey), &v1alpha1.GoogleServiceAccountKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleServiceAccountKey), err
}

// Delete takes name of the googleServiceAccountKey and deletes it. Returns an error if one occurs.
func (c *FakeGoogleServiceAccountKeys) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googleserviceaccountkeysResource, name), &v1alpha1.GoogleServiceAccountKey{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleServiceAccountKeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googleserviceaccountkeysResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleServiceAccountKeyList{})
	return err
}

// Patch applies the patch and returns the patched googleServiceAccountKey.
func (c *FakeGoogleServiceAccountKeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleServiceAccountKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googleserviceaccountkeysResource, name, pt, data, subresources...), &v1alpha1.GoogleServiceAccountKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleServiceAccountKey), err
}