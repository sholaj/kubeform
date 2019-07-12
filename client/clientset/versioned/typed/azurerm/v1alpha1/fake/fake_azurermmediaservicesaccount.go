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

// FakeAzurermMediaServicesAccounts implements AzurermMediaServicesAccountInterface
type FakeAzurermMediaServicesAccounts struct {
	Fake *FakeAzurermV1alpha1
}

var azurermmediaservicesaccountsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermmediaservicesaccounts"}

var azurermmediaservicesaccountsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermMediaServicesAccount"}

// Get takes name of the azurermMediaServicesAccount, and returns the corresponding azurermMediaServicesAccount object, and an error if there is any.
func (c *FakeAzurermMediaServicesAccounts) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermMediaServicesAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermmediaservicesaccountsResource, name), &v1alpha1.AzurermMediaServicesAccount{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMediaServicesAccount), err
}

// List takes label and field selectors, and returns the list of AzurermMediaServicesAccounts that match those selectors.
func (c *FakeAzurermMediaServicesAccounts) List(opts v1.ListOptions) (result *v1alpha1.AzurermMediaServicesAccountList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermmediaservicesaccountsResource, azurermmediaservicesaccountsKind, opts), &v1alpha1.AzurermMediaServicesAccountList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermMediaServicesAccountList{ListMeta: obj.(*v1alpha1.AzurermMediaServicesAccountList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermMediaServicesAccountList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermMediaServicesAccounts.
func (c *FakeAzurermMediaServicesAccounts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermmediaservicesaccountsResource, opts))
}

// Create takes the representation of a azurermMediaServicesAccount and creates it.  Returns the server's representation of the azurermMediaServicesAccount, and an error, if there is any.
func (c *FakeAzurermMediaServicesAccounts) Create(azurermMediaServicesAccount *v1alpha1.AzurermMediaServicesAccount) (result *v1alpha1.AzurermMediaServicesAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermmediaservicesaccountsResource, azurermMediaServicesAccount), &v1alpha1.AzurermMediaServicesAccount{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMediaServicesAccount), err
}

// Update takes the representation of a azurermMediaServicesAccount and updates it. Returns the server's representation of the azurermMediaServicesAccount, and an error, if there is any.
func (c *FakeAzurermMediaServicesAccounts) Update(azurermMediaServicesAccount *v1alpha1.AzurermMediaServicesAccount) (result *v1alpha1.AzurermMediaServicesAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermmediaservicesaccountsResource, azurermMediaServicesAccount), &v1alpha1.AzurermMediaServicesAccount{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMediaServicesAccount), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermMediaServicesAccounts) UpdateStatus(azurermMediaServicesAccount *v1alpha1.AzurermMediaServicesAccount) (*v1alpha1.AzurermMediaServicesAccount, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermmediaservicesaccountsResource, "status", azurermMediaServicesAccount), &v1alpha1.AzurermMediaServicesAccount{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMediaServicesAccount), err
}

// Delete takes name of the azurermMediaServicesAccount and deletes it. Returns an error if one occurs.
func (c *FakeAzurermMediaServicesAccounts) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermmediaservicesaccountsResource, name), &v1alpha1.AzurermMediaServicesAccount{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermMediaServicesAccounts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermmediaservicesaccountsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermMediaServicesAccountList{})
	return err
}

// Patch applies the patch and returns the patched azurermMediaServicesAccount.
func (c *FakeAzurermMediaServicesAccounts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermMediaServicesAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermmediaservicesaccountsResource, name, pt, data, subresources...), &v1alpha1.AzurermMediaServicesAccount{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMediaServicesAccount), err
}