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

// FakeGoogleStorageDefaultObjectAccessControls implements GoogleStorageDefaultObjectAccessControlInterface
type FakeGoogleStorageDefaultObjectAccessControls struct {
	Fake *FakeGoogleV1alpha1
}

var googlestoragedefaultobjectaccesscontrolsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlestoragedefaultobjectaccesscontrols"}

var googlestoragedefaultobjectaccesscontrolsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleStorageDefaultObjectAccessControl"}

// Get takes name of the googleStorageDefaultObjectAccessControl, and returns the corresponding googleStorageDefaultObjectAccessControl object, and an error if there is any.
func (c *FakeGoogleStorageDefaultObjectAccessControls) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleStorageDefaultObjectAccessControl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlestoragedefaultobjectaccesscontrolsResource, name), &v1alpha1.GoogleStorageDefaultObjectAccessControl{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleStorageDefaultObjectAccessControl), err
}

// List takes label and field selectors, and returns the list of GoogleStorageDefaultObjectAccessControls that match those selectors.
func (c *FakeGoogleStorageDefaultObjectAccessControls) List(opts v1.ListOptions) (result *v1alpha1.GoogleStorageDefaultObjectAccessControlList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlestoragedefaultobjectaccesscontrolsResource, googlestoragedefaultobjectaccesscontrolsKind, opts), &v1alpha1.GoogleStorageDefaultObjectAccessControlList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleStorageDefaultObjectAccessControlList{ListMeta: obj.(*v1alpha1.GoogleStorageDefaultObjectAccessControlList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleStorageDefaultObjectAccessControlList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleStorageDefaultObjectAccessControls.
func (c *FakeGoogleStorageDefaultObjectAccessControls) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlestoragedefaultobjectaccesscontrolsResource, opts))
}

// Create takes the representation of a googleStorageDefaultObjectAccessControl and creates it.  Returns the server's representation of the googleStorageDefaultObjectAccessControl, and an error, if there is any.
func (c *FakeGoogleStorageDefaultObjectAccessControls) Create(googleStorageDefaultObjectAccessControl *v1alpha1.GoogleStorageDefaultObjectAccessControl) (result *v1alpha1.GoogleStorageDefaultObjectAccessControl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlestoragedefaultobjectaccesscontrolsResource, googleStorageDefaultObjectAccessControl), &v1alpha1.GoogleStorageDefaultObjectAccessControl{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleStorageDefaultObjectAccessControl), err
}

// Update takes the representation of a googleStorageDefaultObjectAccessControl and updates it. Returns the server's representation of the googleStorageDefaultObjectAccessControl, and an error, if there is any.
func (c *FakeGoogleStorageDefaultObjectAccessControls) Update(googleStorageDefaultObjectAccessControl *v1alpha1.GoogleStorageDefaultObjectAccessControl) (result *v1alpha1.GoogleStorageDefaultObjectAccessControl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlestoragedefaultobjectaccesscontrolsResource, googleStorageDefaultObjectAccessControl), &v1alpha1.GoogleStorageDefaultObjectAccessControl{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleStorageDefaultObjectAccessControl), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleStorageDefaultObjectAccessControls) UpdateStatus(googleStorageDefaultObjectAccessControl *v1alpha1.GoogleStorageDefaultObjectAccessControl) (*v1alpha1.GoogleStorageDefaultObjectAccessControl, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlestoragedefaultobjectaccesscontrolsResource, "status", googleStorageDefaultObjectAccessControl), &v1alpha1.GoogleStorageDefaultObjectAccessControl{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleStorageDefaultObjectAccessControl), err
}

// Delete takes name of the googleStorageDefaultObjectAccessControl and deletes it. Returns an error if one occurs.
func (c *FakeGoogleStorageDefaultObjectAccessControls) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlestoragedefaultobjectaccesscontrolsResource, name), &v1alpha1.GoogleStorageDefaultObjectAccessControl{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleStorageDefaultObjectAccessControls) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlestoragedefaultobjectaccesscontrolsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleStorageDefaultObjectAccessControlList{})
	return err
}

// Patch applies the patch and returns the patched googleStorageDefaultObjectAccessControl.
func (c *FakeGoogleStorageDefaultObjectAccessControls) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleStorageDefaultObjectAccessControl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlestoragedefaultobjectaccesscontrolsResource, name, pt, data, subresources...), &v1alpha1.GoogleStorageDefaultObjectAccessControl{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleStorageDefaultObjectAccessControl), err
}