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

// FakeAzurermExpressRouteCircuits implements AzurermExpressRouteCircuitInterface
type FakeAzurermExpressRouteCircuits struct {
	Fake *FakeAzurermV1alpha1
}

var azurermexpressroutecircuitsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermexpressroutecircuits"}

var azurermexpressroutecircuitsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermExpressRouteCircuit"}

// Get takes name of the azurermExpressRouteCircuit, and returns the corresponding azurermExpressRouteCircuit object, and an error if there is any.
func (c *FakeAzurermExpressRouteCircuits) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermExpressRouteCircuit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermexpressroutecircuitsResource, name), &v1alpha1.AzurermExpressRouteCircuit{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermExpressRouteCircuit), err
}

// List takes label and field selectors, and returns the list of AzurermExpressRouteCircuits that match those selectors.
func (c *FakeAzurermExpressRouteCircuits) List(opts v1.ListOptions) (result *v1alpha1.AzurermExpressRouteCircuitList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermexpressroutecircuitsResource, azurermexpressroutecircuitsKind, opts), &v1alpha1.AzurermExpressRouteCircuitList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermExpressRouteCircuitList{ListMeta: obj.(*v1alpha1.AzurermExpressRouteCircuitList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermExpressRouteCircuitList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermExpressRouteCircuits.
func (c *FakeAzurermExpressRouteCircuits) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermexpressroutecircuitsResource, opts))
}

// Create takes the representation of a azurermExpressRouteCircuit and creates it.  Returns the server's representation of the azurermExpressRouteCircuit, and an error, if there is any.
func (c *FakeAzurermExpressRouteCircuits) Create(azurermExpressRouteCircuit *v1alpha1.AzurermExpressRouteCircuit) (result *v1alpha1.AzurermExpressRouteCircuit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermexpressroutecircuitsResource, azurermExpressRouteCircuit), &v1alpha1.AzurermExpressRouteCircuit{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermExpressRouteCircuit), err
}

// Update takes the representation of a azurermExpressRouteCircuit and updates it. Returns the server's representation of the azurermExpressRouteCircuit, and an error, if there is any.
func (c *FakeAzurermExpressRouteCircuits) Update(azurermExpressRouteCircuit *v1alpha1.AzurermExpressRouteCircuit) (result *v1alpha1.AzurermExpressRouteCircuit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermexpressroutecircuitsResource, azurermExpressRouteCircuit), &v1alpha1.AzurermExpressRouteCircuit{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermExpressRouteCircuit), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermExpressRouteCircuits) UpdateStatus(azurermExpressRouteCircuit *v1alpha1.AzurermExpressRouteCircuit) (*v1alpha1.AzurermExpressRouteCircuit, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermexpressroutecircuitsResource, "status", azurermExpressRouteCircuit), &v1alpha1.AzurermExpressRouteCircuit{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermExpressRouteCircuit), err
}

// Delete takes name of the azurermExpressRouteCircuit and deletes it. Returns an error if one occurs.
func (c *FakeAzurermExpressRouteCircuits) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermexpressroutecircuitsResource, name), &v1alpha1.AzurermExpressRouteCircuit{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermExpressRouteCircuits) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermexpressroutecircuitsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermExpressRouteCircuitList{})
	return err
}

// Patch applies the patch and returns the patched azurermExpressRouteCircuit.
func (c *FakeAzurermExpressRouteCircuits) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermExpressRouteCircuit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermexpressroutecircuitsResource, name, pt, data, subresources...), &v1alpha1.AzurermExpressRouteCircuit{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermExpressRouteCircuit), err
}