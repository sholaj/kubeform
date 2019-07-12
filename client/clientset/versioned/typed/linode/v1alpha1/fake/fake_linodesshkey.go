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
	v1alpha1 "kubeform.dev/kubeform/apis/linode/v1alpha1"
)

// FakeLinodeSshkeys implements LinodeSshkeyInterface
type FakeLinodeSshkeys struct {
	Fake *FakeLinodeV1alpha1
}

var linodesshkeysResource = schema.GroupVersionResource{Group: "linode.kubeform.com", Version: "v1alpha1", Resource: "linodesshkeys"}

var linodesshkeysKind = schema.GroupVersionKind{Group: "linode.kubeform.com", Version: "v1alpha1", Kind: "LinodeSshkey"}

// Get takes name of the linodeSshkey, and returns the corresponding linodeSshkey object, and an error if there is any.
func (c *FakeLinodeSshkeys) Get(name string, options v1.GetOptions) (result *v1alpha1.LinodeSshkey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(linodesshkeysResource, name), &v1alpha1.LinodeSshkey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinodeSshkey), err
}

// List takes label and field selectors, and returns the list of LinodeSshkeys that match those selectors.
func (c *FakeLinodeSshkeys) List(opts v1.ListOptions) (result *v1alpha1.LinodeSshkeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(linodesshkeysResource, linodesshkeysKind, opts), &v1alpha1.LinodeSshkeyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LinodeSshkeyList{ListMeta: obj.(*v1alpha1.LinodeSshkeyList).ListMeta}
	for _, item := range obj.(*v1alpha1.LinodeSshkeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested linodeSshkeys.
func (c *FakeLinodeSshkeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(linodesshkeysResource, opts))
}

// Create takes the representation of a linodeSshkey and creates it.  Returns the server's representation of the linodeSshkey, and an error, if there is any.
func (c *FakeLinodeSshkeys) Create(linodeSshkey *v1alpha1.LinodeSshkey) (result *v1alpha1.LinodeSshkey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(linodesshkeysResource, linodeSshkey), &v1alpha1.LinodeSshkey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinodeSshkey), err
}

// Update takes the representation of a linodeSshkey and updates it. Returns the server's representation of the linodeSshkey, and an error, if there is any.
func (c *FakeLinodeSshkeys) Update(linodeSshkey *v1alpha1.LinodeSshkey) (result *v1alpha1.LinodeSshkey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(linodesshkeysResource, linodeSshkey), &v1alpha1.LinodeSshkey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinodeSshkey), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLinodeSshkeys) UpdateStatus(linodeSshkey *v1alpha1.LinodeSshkey) (*v1alpha1.LinodeSshkey, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(linodesshkeysResource, "status", linodeSshkey), &v1alpha1.LinodeSshkey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinodeSshkey), err
}

// Delete takes name of the linodeSshkey and deletes it. Returns an error if one occurs.
func (c *FakeLinodeSshkeys) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(linodesshkeysResource, name), &v1alpha1.LinodeSshkey{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLinodeSshkeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(linodesshkeysResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LinodeSshkeyList{})
	return err
}

// Patch applies the patch and returns the patched linodeSshkey.
func (c *FakeLinodeSshkeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LinodeSshkey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(linodesshkeysResource, name, pt, data, subresources...), &v1alpha1.LinodeSshkey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinodeSshkey), err
}