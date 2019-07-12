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

// FakeGoogleComputeProjectMetadataItems implements GoogleComputeProjectMetadataItemInterface
type FakeGoogleComputeProjectMetadataItems struct {
	Fake *FakeGoogleV1alpha1
}

var googlecomputeprojectmetadataitemsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlecomputeprojectmetadataitems"}

var googlecomputeprojectmetadataitemsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleComputeProjectMetadataItem"}

// Get takes name of the googleComputeProjectMetadataItem, and returns the corresponding googleComputeProjectMetadataItem object, and an error if there is any.
func (c *FakeGoogleComputeProjectMetadataItems) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleComputeProjectMetadataItem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlecomputeprojectmetadataitemsResource, name), &v1alpha1.GoogleComputeProjectMetadataItem{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeProjectMetadataItem), err
}

// List takes label and field selectors, and returns the list of GoogleComputeProjectMetadataItems that match those selectors.
func (c *FakeGoogleComputeProjectMetadataItems) List(opts v1.ListOptions) (result *v1alpha1.GoogleComputeProjectMetadataItemList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlecomputeprojectmetadataitemsResource, googlecomputeprojectmetadataitemsKind, opts), &v1alpha1.GoogleComputeProjectMetadataItemList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleComputeProjectMetadataItemList{ListMeta: obj.(*v1alpha1.GoogleComputeProjectMetadataItemList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleComputeProjectMetadataItemList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleComputeProjectMetadataItems.
func (c *FakeGoogleComputeProjectMetadataItems) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlecomputeprojectmetadataitemsResource, opts))
}

// Create takes the representation of a googleComputeProjectMetadataItem and creates it.  Returns the server's representation of the googleComputeProjectMetadataItem, and an error, if there is any.
func (c *FakeGoogleComputeProjectMetadataItems) Create(googleComputeProjectMetadataItem *v1alpha1.GoogleComputeProjectMetadataItem) (result *v1alpha1.GoogleComputeProjectMetadataItem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlecomputeprojectmetadataitemsResource, googleComputeProjectMetadataItem), &v1alpha1.GoogleComputeProjectMetadataItem{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeProjectMetadataItem), err
}

// Update takes the representation of a googleComputeProjectMetadataItem and updates it. Returns the server's representation of the googleComputeProjectMetadataItem, and an error, if there is any.
func (c *FakeGoogleComputeProjectMetadataItems) Update(googleComputeProjectMetadataItem *v1alpha1.GoogleComputeProjectMetadataItem) (result *v1alpha1.GoogleComputeProjectMetadataItem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlecomputeprojectmetadataitemsResource, googleComputeProjectMetadataItem), &v1alpha1.GoogleComputeProjectMetadataItem{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeProjectMetadataItem), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleComputeProjectMetadataItems) UpdateStatus(googleComputeProjectMetadataItem *v1alpha1.GoogleComputeProjectMetadataItem) (*v1alpha1.GoogleComputeProjectMetadataItem, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlecomputeprojectmetadataitemsResource, "status", googleComputeProjectMetadataItem), &v1alpha1.GoogleComputeProjectMetadataItem{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeProjectMetadataItem), err
}

// Delete takes name of the googleComputeProjectMetadataItem and deletes it. Returns an error if one occurs.
func (c *FakeGoogleComputeProjectMetadataItems) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlecomputeprojectmetadataitemsResource, name), &v1alpha1.GoogleComputeProjectMetadataItem{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleComputeProjectMetadataItems) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlecomputeprojectmetadataitemsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleComputeProjectMetadataItemList{})
	return err
}

// Patch applies the patch and returns the patched googleComputeProjectMetadataItem.
func (c *FakeGoogleComputeProjectMetadataItems) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComputeProjectMetadataItem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlecomputeprojectmetadataitemsResource, name, pt, data, subresources...), &v1alpha1.GoogleComputeProjectMetadataItem{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeProjectMetadataItem), err
}