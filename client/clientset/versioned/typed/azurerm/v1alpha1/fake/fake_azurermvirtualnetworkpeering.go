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

// FakeAzurermVirtualNetworkPeerings implements AzurermVirtualNetworkPeeringInterface
type FakeAzurermVirtualNetworkPeerings struct {
	Fake *FakeAzurermV1alpha1
}

var azurermvirtualnetworkpeeringsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermvirtualnetworkpeerings"}

var azurermvirtualnetworkpeeringsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermVirtualNetworkPeering"}

// Get takes name of the azurermVirtualNetworkPeering, and returns the corresponding azurermVirtualNetworkPeering object, and an error if there is any.
func (c *FakeAzurermVirtualNetworkPeerings) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermVirtualNetworkPeering, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermvirtualnetworkpeeringsResource, name), &v1alpha1.AzurermVirtualNetworkPeering{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermVirtualNetworkPeering), err
}

// List takes label and field selectors, and returns the list of AzurermVirtualNetworkPeerings that match those selectors.
func (c *FakeAzurermVirtualNetworkPeerings) List(opts v1.ListOptions) (result *v1alpha1.AzurermVirtualNetworkPeeringList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermvirtualnetworkpeeringsResource, azurermvirtualnetworkpeeringsKind, opts), &v1alpha1.AzurermVirtualNetworkPeeringList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermVirtualNetworkPeeringList{ListMeta: obj.(*v1alpha1.AzurermVirtualNetworkPeeringList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermVirtualNetworkPeeringList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermVirtualNetworkPeerings.
func (c *FakeAzurermVirtualNetworkPeerings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermvirtualnetworkpeeringsResource, opts))
}

// Create takes the representation of a azurermVirtualNetworkPeering and creates it.  Returns the server's representation of the azurermVirtualNetworkPeering, and an error, if there is any.
func (c *FakeAzurermVirtualNetworkPeerings) Create(azurermVirtualNetworkPeering *v1alpha1.AzurermVirtualNetworkPeering) (result *v1alpha1.AzurermVirtualNetworkPeering, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermvirtualnetworkpeeringsResource, azurermVirtualNetworkPeering), &v1alpha1.AzurermVirtualNetworkPeering{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermVirtualNetworkPeering), err
}

// Update takes the representation of a azurermVirtualNetworkPeering and updates it. Returns the server's representation of the azurermVirtualNetworkPeering, and an error, if there is any.
func (c *FakeAzurermVirtualNetworkPeerings) Update(azurermVirtualNetworkPeering *v1alpha1.AzurermVirtualNetworkPeering) (result *v1alpha1.AzurermVirtualNetworkPeering, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermvirtualnetworkpeeringsResource, azurermVirtualNetworkPeering), &v1alpha1.AzurermVirtualNetworkPeering{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermVirtualNetworkPeering), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermVirtualNetworkPeerings) UpdateStatus(azurermVirtualNetworkPeering *v1alpha1.AzurermVirtualNetworkPeering) (*v1alpha1.AzurermVirtualNetworkPeering, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermvirtualnetworkpeeringsResource, "status", azurermVirtualNetworkPeering), &v1alpha1.AzurermVirtualNetworkPeering{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermVirtualNetworkPeering), err
}

// Delete takes name of the azurermVirtualNetworkPeering and deletes it. Returns an error if one occurs.
func (c *FakeAzurermVirtualNetworkPeerings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermvirtualnetworkpeeringsResource, name), &v1alpha1.AzurermVirtualNetworkPeering{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermVirtualNetworkPeerings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermvirtualnetworkpeeringsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermVirtualNetworkPeeringList{})
	return err
}

// Patch applies the patch and returns the patched azurermVirtualNetworkPeering.
func (c *FakeAzurermVirtualNetworkPeerings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermVirtualNetworkPeering, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermvirtualnetworkpeeringsResource, name, pt, data, subresources...), &v1alpha1.AzurermVirtualNetworkPeering{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermVirtualNetworkPeering), err
}