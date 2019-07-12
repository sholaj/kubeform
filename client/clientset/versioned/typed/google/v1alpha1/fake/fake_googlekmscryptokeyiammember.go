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

// FakeGoogleKmsCryptoKeyIamMembers implements GoogleKmsCryptoKeyIamMemberInterface
type FakeGoogleKmsCryptoKeyIamMembers struct {
	Fake *FakeGoogleV1alpha1
}

var googlekmscryptokeyiammembersResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlekmscryptokeyiammembers"}

var googlekmscryptokeyiammembersKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleKmsCryptoKeyIamMember"}

// Get takes name of the googleKmsCryptoKeyIamMember, and returns the corresponding googleKmsCryptoKeyIamMember object, and an error if there is any.
func (c *FakeGoogleKmsCryptoKeyIamMembers) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleKmsCryptoKeyIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlekmscryptokeyiammembersResource, name), &v1alpha1.GoogleKmsCryptoKeyIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleKmsCryptoKeyIamMember), err
}

// List takes label and field selectors, and returns the list of GoogleKmsCryptoKeyIamMembers that match those selectors.
func (c *FakeGoogleKmsCryptoKeyIamMembers) List(opts v1.ListOptions) (result *v1alpha1.GoogleKmsCryptoKeyIamMemberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlekmscryptokeyiammembersResource, googlekmscryptokeyiammembersKind, opts), &v1alpha1.GoogleKmsCryptoKeyIamMemberList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleKmsCryptoKeyIamMemberList{ListMeta: obj.(*v1alpha1.GoogleKmsCryptoKeyIamMemberList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleKmsCryptoKeyIamMemberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleKmsCryptoKeyIamMembers.
func (c *FakeGoogleKmsCryptoKeyIamMembers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlekmscryptokeyiammembersResource, opts))
}

// Create takes the representation of a googleKmsCryptoKeyIamMember and creates it.  Returns the server's representation of the googleKmsCryptoKeyIamMember, and an error, if there is any.
func (c *FakeGoogleKmsCryptoKeyIamMembers) Create(googleKmsCryptoKeyIamMember *v1alpha1.GoogleKmsCryptoKeyIamMember) (result *v1alpha1.GoogleKmsCryptoKeyIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlekmscryptokeyiammembersResource, googleKmsCryptoKeyIamMember), &v1alpha1.GoogleKmsCryptoKeyIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleKmsCryptoKeyIamMember), err
}

// Update takes the representation of a googleKmsCryptoKeyIamMember and updates it. Returns the server's representation of the googleKmsCryptoKeyIamMember, and an error, if there is any.
func (c *FakeGoogleKmsCryptoKeyIamMembers) Update(googleKmsCryptoKeyIamMember *v1alpha1.GoogleKmsCryptoKeyIamMember) (result *v1alpha1.GoogleKmsCryptoKeyIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlekmscryptokeyiammembersResource, googleKmsCryptoKeyIamMember), &v1alpha1.GoogleKmsCryptoKeyIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleKmsCryptoKeyIamMember), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleKmsCryptoKeyIamMembers) UpdateStatus(googleKmsCryptoKeyIamMember *v1alpha1.GoogleKmsCryptoKeyIamMember) (*v1alpha1.GoogleKmsCryptoKeyIamMember, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlekmscryptokeyiammembersResource, "status", googleKmsCryptoKeyIamMember), &v1alpha1.GoogleKmsCryptoKeyIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleKmsCryptoKeyIamMember), err
}

// Delete takes name of the googleKmsCryptoKeyIamMember and deletes it. Returns an error if one occurs.
func (c *FakeGoogleKmsCryptoKeyIamMembers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlekmscryptokeyiammembersResource, name), &v1alpha1.GoogleKmsCryptoKeyIamMember{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleKmsCryptoKeyIamMembers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlekmscryptokeyiammembersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleKmsCryptoKeyIamMemberList{})
	return err
}

// Patch applies the patch and returns the patched googleKmsCryptoKeyIamMember.
func (c *FakeGoogleKmsCryptoKeyIamMembers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleKmsCryptoKeyIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlekmscryptokeyiammembersResource, name, pt, data, subresources...), &v1alpha1.GoogleKmsCryptoKeyIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleKmsCryptoKeyIamMember), err
}