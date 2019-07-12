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

// FakeAzurermCdnEndpoints implements AzurermCdnEndpointInterface
type FakeAzurermCdnEndpoints struct {
	Fake *FakeAzurermV1alpha1
}

var azurermcdnendpointsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermcdnendpoints"}

var azurermcdnendpointsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermCdnEndpoint"}

// Get takes name of the azurermCdnEndpoint, and returns the corresponding azurermCdnEndpoint object, and an error if there is any.
func (c *FakeAzurermCdnEndpoints) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermCdnEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermcdnendpointsResource, name), &v1alpha1.AzurermCdnEndpoint{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermCdnEndpoint), err
}

// List takes label and field selectors, and returns the list of AzurermCdnEndpoints that match those selectors.
func (c *FakeAzurermCdnEndpoints) List(opts v1.ListOptions) (result *v1alpha1.AzurermCdnEndpointList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermcdnendpointsResource, azurermcdnendpointsKind, opts), &v1alpha1.AzurermCdnEndpointList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermCdnEndpointList{ListMeta: obj.(*v1alpha1.AzurermCdnEndpointList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermCdnEndpointList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermCdnEndpoints.
func (c *FakeAzurermCdnEndpoints) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermcdnendpointsResource, opts))
}

// Create takes the representation of a azurermCdnEndpoint and creates it.  Returns the server's representation of the azurermCdnEndpoint, and an error, if there is any.
func (c *FakeAzurermCdnEndpoints) Create(azurermCdnEndpoint *v1alpha1.AzurermCdnEndpoint) (result *v1alpha1.AzurermCdnEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermcdnendpointsResource, azurermCdnEndpoint), &v1alpha1.AzurermCdnEndpoint{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermCdnEndpoint), err
}

// Update takes the representation of a azurermCdnEndpoint and updates it. Returns the server's representation of the azurermCdnEndpoint, and an error, if there is any.
func (c *FakeAzurermCdnEndpoints) Update(azurermCdnEndpoint *v1alpha1.AzurermCdnEndpoint) (result *v1alpha1.AzurermCdnEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermcdnendpointsResource, azurermCdnEndpoint), &v1alpha1.AzurermCdnEndpoint{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermCdnEndpoint), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermCdnEndpoints) UpdateStatus(azurermCdnEndpoint *v1alpha1.AzurermCdnEndpoint) (*v1alpha1.AzurermCdnEndpoint, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermcdnendpointsResource, "status", azurermCdnEndpoint), &v1alpha1.AzurermCdnEndpoint{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermCdnEndpoint), err
}

// Delete takes name of the azurermCdnEndpoint and deletes it. Returns an error if one occurs.
func (c *FakeAzurermCdnEndpoints) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermcdnendpointsResource, name), &v1alpha1.AzurermCdnEndpoint{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermCdnEndpoints) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermcdnendpointsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermCdnEndpointList{})
	return err
}

// Patch applies the patch and returns the patched azurermCdnEndpoint.
func (c *FakeAzurermCdnEndpoints) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermCdnEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermcdnendpointsResource, name, pt, data, subresources...), &v1alpha1.AzurermCdnEndpoint{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermCdnEndpoint), err
}