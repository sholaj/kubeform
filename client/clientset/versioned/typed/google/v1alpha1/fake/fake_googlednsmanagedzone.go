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

// FakeGoogleDnsManagedZones implements GoogleDnsManagedZoneInterface
type FakeGoogleDnsManagedZones struct {
	Fake *FakeGoogleV1alpha1
}

var googlednsmanagedzonesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlednsmanagedzones"}

var googlednsmanagedzonesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleDnsManagedZone"}

// Get takes name of the googleDnsManagedZone, and returns the corresponding googleDnsManagedZone object, and an error if there is any.
func (c *FakeGoogleDnsManagedZones) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleDnsManagedZone, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlednsmanagedzonesResource, name), &v1alpha1.GoogleDnsManagedZone{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleDnsManagedZone), err
}

// List takes label and field selectors, and returns the list of GoogleDnsManagedZones that match those selectors.
func (c *FakeGoogleDnsManagedZones) List(opts v1.ListOptions) (result *v1alpha1.GoogleDnsManagedZoneList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlednsmanagedzonesResource, googlednsmanagedzonesKind, opts), &v1alpha1.GoogleDnsManagedZoneList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleDnsManagedZoneList{ListMeta: obj.(*v1alpha1.GoogleDnsManagedZoneList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleDnsManagedZoneList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleDnsManagedZones.
func (c *FakeGoogleDnsManagedZones) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlednsmanagedzonesResource, opts))
}

// Create takes the representation of a googleDnsManagedZone and creates it.  Returns the server's representation of the googleDnsManagedZone, and an error, if there is any.
func (c *FakeGoogleDnsManagedZones) Create(googleDnsManagedZone *v1alpha1.GoogleDnsManagedZone) (result *v1alpha1.GoogleDnsManagedZone, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlednsmanagedzonesResource, googleDnsManagedZone), &v1alpha1.GoogleDnsManagedZone{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleDnsManagedZone), err
}

// Update takes the representation of a googleDnsManagedZone and updates it. Returns the server's representation of the googleDnsManagedZone, and an error, if there is any.
func (c *FakeGoogleDnsManagedZones) Update(googleDnsManagedZone *v1alpha1.GoogleDnsManagedZone) (result *v1alpha1.GoogleDnsManagedZone, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlednsmanagedzonesResource, googleDnsManagedZone), &v1alpha1.GoogleDnsManagedZone{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleDnsManagedZone), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleDnsManagedZones) UpdateStatus(googleDnsManagedZone *v1alpha1.GoogleDnsManagedZone) (*v1alpha1.GoogleDnsManagedZone, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlednsmanagedzonesResource, "status", googleDnsManagedZone), &v1alpha1.GoogleDnsManagedZone{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleDnsManagedZone), err
}

// Delete takes name of the googleDnsManagedZone and deletes it. Returns an error if one occurs.
func (c *FakeGoogleDnsManagedZones) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlednsmanagedzonesResource, name), &v1alpha1.GoogleDnsManagedZone{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleDnsManagedZones) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlednsmanagedzonesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleDnsManagedZoneList{})
	return err
}

// Patch applies the patch and returns the patched googleDnsManagedZone.
func (c *FakeGoogleDnsManagedZones) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleDnsManagedZone, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlednsmanagedzonesResource, name, pt, data, subresources...), &v1alpha1.GoogleDnsManagedZone{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleDnsManagedZone), err
}