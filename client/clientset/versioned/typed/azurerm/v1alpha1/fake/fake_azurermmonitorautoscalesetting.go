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

// FakeAzurermMonitorAutoscaleSettings implements AzurermMonitorAutoscaleSettingInterface
type FakeAzurermMonitorAutoscaleSettings struct {
	Fake *FakeAzurermV1alpha1
}

var azurermmonitorautoscalesettingsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermmonitorautoscalesettings"}

var azurermmonitorautoscalesettingsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermMonitorAutoscaleSetting"}

// Get takes name of the azurermMonitorAutoscaleSetting, and returns the corresponding azurermMonitorAutoscaleSetting object, and an error if there is any.
func (c *FakeAzurermMonitorAutoscaleSettings) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermMonitorAutoscaleSetting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermmonitorautoscalesettingsResource, name), &v1alpha1.AzurermMonitorAutoscaleSetting{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMonitorAutoscaleSetting), err
}

// List takes label and field selectors, and returns the list of AzurermMonitorAutoscaleSettings that match those selectors.
func (c *FakeAzurermMonitorAutoscaleSettings) List(opts v1.ListOptions) (result *v1alpha1.AzurermMonitorAutoscaleSettingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermmonitorautoscalesettingsResource, azurermmonitorautoscalesettingsKind, opts), &v1alpha1.AzurermMonitorAutoscaleSettingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermMonitorAutoscaleSettingList{ListMeta: obj.(*v1alpha1.AzurermMonitorAutoscaleSettingList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermMonitorAutoscaleSettingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermMonitorAutoscaleSettings.
func (c *FakeAzurermMonitorAutoscaleSettings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermmonitorautoscalesettingsResource, opts))
}

// Create takes the representation of a azurermMonitorAutoscaleSetting and creates it.  Returns the server's representation of the azurermMonitorAutoscaleSetting, and an error, if there is any.
func (c *FakeAzurermMonitorAutoscaleSettings) Create(azurermMonitorAutoscaleSetting *v1alpha1.AzurermMonitorAutoscaleSetting) (result *v1alpha1.AzurermMonitorAutoscaleSetting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermmonitorautoscalesettingsResource, azurermMonitorAutoscaleSetting), &v1alpha1.AzurermMonitorAutoscaleSetting{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMonitorAutoscaleSetting), err
}

// Update takes the representation of a azurermMonitorAutoscaleSetting and updates it. Returns the server's representation of the azurermMonitorAutoscaleSetting, and an error, if there is any.
func (c *FakeAzurermMonitorAutoscaleSettings) Update(azurermMonitorAutoscaleSetting *v1alpha1.AzurermMonitorAutoscaleSetting) (result *v1alpha1.AzurermMonitorAutoscaleSetting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermmonitorautoscalesettingsResource, azurermMonitorAutoscaleSetting), &v1alpha1.AzurermMonitorAutoscaleSetting{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMonitorAutoscaleSetting), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermMonitorAutoscaleSettings) UpdateStatus(azurermMonitorAutoscaleSetting *v1alpha1.AzurermMonitorAutoscaleSetting) (*v1alpha1.AzurermMonitorAutoscaleSetting, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermmonitorautoscalesettingsResource, "status", azurermMonitorAutoscaleSetting), &v1alpha1.AzurermMonitorAutoscaleSetting{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMonitorAutoscaleSetting), err
}

// Delete takes name of the azurermMonitorAutoscaleSetting and deletes it. Returns an error if one occurs.
func (c *FakeAzurermMonitorAutoscaleSettings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermmonitorautoscalesettingsResource, name), &v1alpha1.AzurermMonitorAutoscaleSetting{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermMonitorAutoscaleSettings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermmonitorautoscalesettingsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermMonitorAutoscaleSettingList{})
	return err
}

// Patch applies the patch and returns the patched azurermMonitorAutoscaleSetting.
func (c *FakeAzurermMonitorAutoscaleSettings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermMonitorAutoscaleSetting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermmonitorautoscalesettingsResource, name, pt, data, subresources...), &v1alpha1.AzurermMonitorAutoscaleSetting{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMonitorAutoscaleSetting), err
}