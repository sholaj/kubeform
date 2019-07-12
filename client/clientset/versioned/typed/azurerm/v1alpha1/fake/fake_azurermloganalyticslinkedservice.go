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

// FakeAzurermLogAnalyticsLinkedServices implements AzurermLogAnalyticsLinkedServiceInterface
type FakeAzurermLogAnalyticsLinkedServices struct {
	Fake *FakeAzurermV1alpha1
}

var azurermloganalyticslinkedservicesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermloganalyticslinkedservices"}

var azurermloganalyticslinkedservicesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermLogAnalyticsLinkedService"}

// Get takes name of the azurermLogAnalyticsLinkedService, and returns the corresponding azurermLogAnalyticsLinkedService object, and an error if there is any.
func (c *FakeAzurermLogAnalyticsLinkedServices) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermLogAnalyticsLinkedService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermloganalyticslinkedservicesResource, name), &v1alpha1.AzurermLogAnalyticsLinkedService{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermLogAnalyticsLinkedService), err
}

// List takes label and field selectors, and returns the list of AzurermLogAnalyticsLinkedServices that match those selectors.
func (c *FakeAzurermLogAnalyticsLinkedServices) List(opts v1.ListOptions) (result *v1alpha1.AzurermLogAnalyticsLinkedServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermloganalyticslinkedservicesResource, azurermloganalyticslinkedservicesKind, opts), &v1alpha1.AzurermLogAnalyticsLinkedServiceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermLogAnalyticsLinkedServiceList{ListMeta: obj.(*v1alpha1.AzurermLogAnalyticsLinkedServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermLogAnalyticsLinkedServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermLogAnalyticsLinkedServices.
func (c *FakeAzurermLogAnalyticsLinkedServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermloganalyticslinkedservicesResource, opts))
}

// Create takes the representation of a azurermLogAnalyticsLinkedService and creates it.  Returns the server's representation of the azurermLogAnalyticsLinkedService, and an error, if there is any.
func (c *FakeAzurermLogAnalyticsLinkedServices) Create(azurermLogAnalyticsLinkedService *v1alpha1.AzurermLogAnalyticsLinkedService) (result *v1alpha1.AzurermLogAnalyticsLinkedService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermloganalyticslinkedservicesResource, azurermLogAnalyticsLinkedService), &v1alpha1.AzurermLogAnalyticsLinkedService{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermLogAnalyticsLinkedService), err
}

// Update takes the representation of a azurermLogAnalyticsLinkedService and updates it. Returns the server's representation of the azurermLogAnalyticsLinkedService, and an error, if there is any.
func (c *FakeAzurermLogAnalyticsLinkedServices) Update(azurermLogAnalyticsLinkedService *v1alpha1.AzurermLogAnalyticsLinkedService) (result *v1alpha1.AzurermLogAnalyticsLinkedService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermloganalyticslinkedservicesResource, azurermLogAnalyticsLinkedService), &v1alpha1.AzurermLogAnalyticsLinkedService{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermLogAnalyticsLinkedService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermLogAnalyticsLinkedServices) UpdateStatus(azurermLogAnalyticsLinkedService *v1alpha1.AzurermLogAnalyticsLinkedService) (*v1alpha1.AzurermLogAnalyticsLinkedService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermloganalyticslinkedservicesResource, "status", azurermLogAnalyticsLinkedService), &v1alpha1.AzurermLogAnalyticsLinkedService{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermLogAnalyticsLinkedService), err
}

// Delete takes name of the azurermLogAnalyticsLinkedService and deletes it. Returns an error if one occurs.
func (c *FakeAzurermLogAnalyticsLinkedServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermloganalyticslinkedservicesResource, name), &v1alpha1.AzurermLogAnalyticsLinkedService{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermLogAnalyticsLinkedServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermloganalyticslinkedservicesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermLogAnalyticsLinkedServiceList{})
	return err
}

// Patch applies the patch and returns the patched azurermLogAnalyticsLinkedService.
func (c *FakeAzurermLogAnalyticsLinkedServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermLogAnalyticsLinkedService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermloganalyticslinkedservicesResource, name, pt, data, subresources...), &v1alpha1.AzurermLogAnalyticsLinkedService{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermLogAnalyticsLinkedService), err
}