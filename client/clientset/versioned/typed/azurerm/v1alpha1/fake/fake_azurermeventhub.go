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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeAzurermEventhubs implements AzurermEventhubInterface
type FakeAzurermEventhubs struct {
	Fake *FakeAzurermV1alpha1
}

var azurermeventhubsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermeventhubs"}

var azurermeventhubsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermEventhub"}

// Get takes name of the azurermEventhub, and returns the corresponding azurermEventhub object, and an error if there is any.
func (c *FakeAzurermEventhubs) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermEventhub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermeventhubsResource, name), &v1alpha1.AzurermEventhub{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermEventhub), err
}

// List takes label and field selectors, and returns the list of AzurermEventhubs that match those selectors.
func (c *FakeAzurermEventhubs) List(opts v1.ListOptions) (result *v1alpha1.AzurermEventhubList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermeventhubsResource, azurermeventhubsKind, opts), &v1alpha1.AzurermEventhubList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermEventhubList{ListMeta: obj.(*v1alpha1.AzurermEventhubList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermEventhubList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermEventhubs.
func (c *FakeAzurermEventhubs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermeventhubsResource, opts))
}

// Create takes the representation of a azurermEventhub and creates it.  Returns the server's representation of the azurermEventhub, and an error, if there is any.
func (c *FakeAzurermEventhubs) Create(azurermEventhub *v1alpha1.AzurermEventhub) (result *v1alpha1.AzurermEventhub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermeventhubsResource, azurermEventhub), &v1alpha1.AzurermEventhub{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermEventhub), err
}

// Update takes the representation of a azurermEventhub and updates it. Returns the server's representation of the azurermEventhub, and an error, if there is any.
func (c *FakeAzurermEventhubs) Update(azurermEventhub *v1alpha1.AzurermEventhub) (result *v1alpha1.AzurermEventhub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermeventhubsResource, azurermEventhub), &v1alpha1.AzurermEventhub{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermEventhub), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermEventhubs) UpdateStatus(azurermEventhub *v1alpha1.AzurermEventhub) (*v1alpha1.AzurermEventhub, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermeventhubsResource, "status", azurermEventhub), &v1alpha1.AzurermEventhub{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermEventhub), err
}

// Delete takes name of the azurermEventhub and deletes it. Returns an error if one occurs.
func (c *FakeAzurermEventhubs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermeventhubsResource, name), &v1alpha1.AzurermEventhub{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermEventhubs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermeventhubsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermEventhubList{})
	return err
}

// Patch applies the patch and returns the patched azurermEventhub.
func (c *FakeAzurermEventhubs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermEventhub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermeventhubsResource, name, pt, data, subresources...), &v1alpha1.AzurermEventhub{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermEventhub), err
}