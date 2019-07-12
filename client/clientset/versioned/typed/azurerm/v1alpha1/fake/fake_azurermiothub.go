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

// FakeAzurermIothubs implements AzurermIothubInterface
type FakeAzurermIothubs struct {
	Fake *FakeAzurermV1alpha1
}

var azurermiothubsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermiothubs"}

var azurermiothubsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermIothub"}

// Get takes name of the azurermIothub, and returns the corresponding azurermIothub object, and an error if there is any.
func (c *FakeAzurermIothubs) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermIothub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermiothubsResource, name), &v1alpha1.AzurermIothub{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermIothub), err
}

// List takes label and field selectors, and returns the list of AzurermIothubs that match those selectors.
func (c *FakeAzurermIothubs) List(opts v1.ListOptions) (result *v1alpha1.AzurermIothubList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermiothubsResource, azurermiothubsKind, opts), &v1alpha1.AzurermIothubList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermIothubList{ListMeta: obj.(*v1alpha1.AzurermIothubList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermIothubList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermIothubs.
func (c *FakeAzurermIothubs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermiothubsResource, opts))
}

// Create takes the representation of a azurermIothub and creates it.  Returns the server's representation of the azurermIothub, and an error, if there is any.
func (c *FakeAzurermIothubs) Create(azurermIothub *v1alpha1.AzurermIothub) (result *v1alpha1.AzurermIothub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermiothubsResource, azurermIothub), &v1alpha1.AzurermIothub{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermIothub), err
}

// Update takes the representation of a azurermIothub and updates it. Returns the server's representation of the azurermIothub, and an error, if there is any.
func (c *FakeAzurermIothubs) Update(azurermIothub *v1alpha1.AzurermIothub) (result *v1alpha1.AzurermIothub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermiothubsResource, azurermIothub), &v1alpha1.AzurermIothub{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermIothub), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermIothubs) UpdateStatus(azurermIothub *v1alpha1.AzurermIothub) (*v1alpha1.AzurermIothub, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermiothubsResource, "status", azurermIothub), &v1alpha1.AzurermIothub{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermIothub), err
}

// Delete takes name of the azurermIothub and deletes it. Returns an error if one occurs.
func (c *FakeAzurermIothubs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermiothubsResource, name), &v1alpha1.AzurermIothub{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermIothubs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermiothubsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermIothubList{})
	return err
}

// Patch applies the patch and returns the patched azurermIothub.
func (c *FakeAzurermIothubs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermIothub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermiothubsResource, name, pt, data, subresources...), &v1alpha1.AzurermIothub{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermIothub), err
}