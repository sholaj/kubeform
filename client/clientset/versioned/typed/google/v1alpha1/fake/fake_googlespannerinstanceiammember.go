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

// FakeGoogleSpannerInstanceIamMembers implements GoogleSpannerInstanceIamMemberInterface
type FakeGoogleSpannerInstanceIamMembers struct {
	Fake *FakeGoogleV1alpha1
}

var googlespannerinstanceiammembersResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlespannerinstanceiammembers"}

var googlespannerinstanceiammembersKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleSpannerInstanceIamMember"}

// Get takes name of the googleSpannerInstanceIamMember, and returns the corresponding googleSpannerInstanceIamMember object, and an error if there is any.
func (c *FakeGoogleSpannerInstanceIamMembers) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleSpannerInstanceIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlespannerinstanceiammembersResource, name), &v1alpha1.GoogleSpannerInstanceIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleSpannerInstanceIamMember), err
}

// List takes label and field selectors, and returns the list of GoogleSpannerInstanceIamMembers that match those selectors.
func (c *FakeGoogleSpannerInstanceIamMembers) List(opts v1.ListOptions) (result *v1alpha1.GoogleSpannerInstanceIamMemberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlespannerinstanceiammembersResource, googlespannerinstanceiammembersKind, opts), &v1alpha1.GoogleSpannerInstanceIamMemberList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleSpannerInstanceIamMemberList{ListMeta: obj.(*v1alpha1.GoogleSpannerInstanceIamMemberList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleSpannerInstanceIamMemberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleSpannerInstanceIamMembers.
func (c *FakeGoogleSpannerInstanceIamMembers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlespannerinstanceiammembersResource, opts))
}

// Create takes the representation of a googleSpannerInstanceIamMember and creates it.  Returns the server's representation of the googleSpannerInstanceIamMember, and an error, if there is any.
func (c *FakeGoogleSpannerInstanceIamMembers) Create(googleSpannerInstanceIamMember *v1alpha1.GoogleSpannerInstanceIamMember) (result *v1alpha1.GoogleSpannerInstanceIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlespannerinstanceiammembersResource, googleSpannerInstanceIamMember), &v1alpha1.GoogleSpannerInstanceIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleSpannerInstanceIamMember), err
}

// Update takes the representation of a googleSpannerInstanceIamMember and updates it. Returns the server's representation of the googleSpannerInstanceIamMember, and an error, if there is any.
func (c *FakeGoogleSpannerInstanceIamMembers) Update(googleSpannerInstanceIamMember *v1alpha1.GoogleSpannerInstanceIamMember) (result *v1alpha1.GoogleSpannerInstanceIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlespannerinstanceiammembersResource, googleSpannerInstanceIamMember), &v1alpha1.GoogleSpannerInstanceIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleSpannerInstanceIamMember), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleSpannerInstanceIamMembers) UpdateStatus(googleSpannerInstanceIamMember *v1alpha1.GoogleSpannerInstanceIamMember) (*v1alpha1.GoogleSpannerInstanceIamMember, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlespannerinstanceiammembersResource, "status", googleSpannerInstanceIamMember), &v1alpha1.GoogleSpannerInstanceIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleSpannerInstanceIamMember), err
}

// Delete takes name of the googleSpannerInstanceIamMember and deletes it. Returns an error if one occurs.
func (c *FakeGoogleSpannerInstanceIamMembers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlespannerinstanceiammembersResource, name), &v1alpha1.GoogleSpannerInstanceIamMember{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleSpannerInstanceIamMembers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlespannerinstanceiammembersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleSpannerInstanceIamMemberList{})
	return err
}

// Patch applies the patch and returns the patched googleSpannerInstanceIamMember.
func (c *FakeGoogleSpannerInstanceIamMembers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleSpannerInstanceIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlespannerinstanceiammembersResource, name, pt, data, subresources...), &v1alpha1.GoogleSpannerInstanceIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleSpannerInstanceIamMember), err
}