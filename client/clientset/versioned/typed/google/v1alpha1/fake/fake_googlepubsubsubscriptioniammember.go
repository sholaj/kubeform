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

// FakeGooglePubsubSubscriptionIamMembers implements GooglePubsubSubscriptionIamMemberInterface
type FakeGooglePubsubSubscriptionIamMembers struct {
	Fake *FakeGoogleV1alpha1
}

var googlepubsubsubscriptioniammembersResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlepubsubsubscriptioniammembers"}

var googlepubsubsubscriptioniammembersKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GooglePubsubSubscriptionIamMember"}

// Get takes name of the googlePubsubSubscriptionIamMember, and returns the corresponding googlePubsubSubscriptionIamMember object, and an error if there is any.
func (c *FakeGooglePubsubSubscriptionIamMembers) Get(name string, options v1.GetOptions) (result *v1alpha1.GooglePubsubSubscriptionIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlepubsubsubscriptioniammembersResource, name), &v1alpha1.GooglePubsubSubscriptionIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GooglePubsubSubscriptionIamMember), err
}

// List takes label and field selectors, and returns the list of GooglePubsubSubscriptionIamMembers that match those selectors.
func (c *FakeGooglePubsubSubscriptionIamMembers) List(opts v1.ListOptions) (result *v1alpha1.GooglePubsubSubscriptionIamMemberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlepubsubsubscriptioniammembersResource, googlepubsubsubscriptioniammembersKind, opts), &v1alpha1.GooglePubsubSubscriptionIamMemberList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GooglePubsubSubscriptionIamMemberList{ListMeta: obj.(*v1alpha1.GooglePubsubSubscriptionIamMemberList).ListMeta}
	for _, item := range obj.(*v1alpha1.GooglePubsubSubscriptionIamMemberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googlePubsubSubscriptionIamMembers.
func (c *FakeGooglePubsubSubscriptionIamMembers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlepubsubsubscriptioniammembersResource, opts))
}

// Create takes the representation of a googlePubsubSubscriptionIamMember and creates it.  Returns the server's representation of the googlePubsubSubscriptionIamMember, and an error, if there is any.
func (c *FakeGooglePubsubSubscriptionIamMembers) Create(googlePubsubSubscriptionIamMember *v1alpha1.GooglePubsubSubscriptionIamMember) (result *v1alpha1.GooglePubsubSubscriptionIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlepubsubsubscriptioniammembersResource, googlePubsubSubscriptionIamMember), &v1alpha1.GooglePubsubSubscriptionIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GooglePubsubSubscriptionIamMember), err
}

// Update takes the representation of a googlePubsubSubscriptionIamMember and updates it. Returns the server's representation of the googlePubsubSubscriptionIamMember, and an error, if there is any.
func (c *FakeGooglePubsubSubscriptionIamMembers) Update(googlePubsubSubscriptionIamMember *v1alpha1.GooglePubsubSubscriptionIamMember) (result *v1alpha1.GooglePubsubSubscriptionIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlepubsubsubscriptioniammembersResource, googlePubsubSubscriptionIamMember), &v1alpha1.GooglePubsubSubscriptionIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GooglePubsubSubscriptionIamMember), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGooglePubsubSubscriptionIamMembers) UpdateStatus(googlePubsubSubscriptionIamMember *v1alpha1.GooglePubsubSubscriptionIamMember) (*v1alpha1.GooglePubsubSubscriptionIamMember, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlepubsubsubscriptioniammembersResource, "status", googlePubsubSubscriptionIamMember), &v1alpha1.GooglePubsubSubscriptionIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GooglePubsubSubscriptionIamMember), err
}

// Delete takes name of the googlePubsubSubscriptionIamMember and deletes it. Returns an error if one occurs.
func (c *FakeGooglePubsubSubscriptionIamMembers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlepubsubsubscriptioniammembersResource, name), &v1alpha1.GooglePubsubSubscriptionIamMember{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGooglePubsubSubscriptionIamMembers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlepubsubsubscriptioniammembersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GooglePubsubSubscriptionIamMemberList{})
	return err
}

// Patch applies the patch and returns the patched googlePubsubSubscriptionIamMember.
func (c *FakeGooglePubsubSubscriptionIamMembers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GooglePubsubSubscriptionIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlepubsubsubscriptioniammembersResource, name, pt, data, subresources...), &v1alpha1.GooglePubsubSubscriptionIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GooglePubsubSubscriptionIamMember), err
}