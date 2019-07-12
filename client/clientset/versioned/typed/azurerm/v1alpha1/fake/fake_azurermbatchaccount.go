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

// FakeAzurermBatchAccounts implements AzurermBatchAccountInterface
type FakeAzurermBatchAccounts struct {
	Fake *FakeAzurermV1alpha1
}

var azurermbatchaccountsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermbatchaccounts"}

var azurermbatchaccountsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermBatchAccount"}

// Get takes name of the azurermBatchAccount, and returns the corresponding azurermBatchAccount object, and an error if there is any.
func (c *FakeAzurermBatchAccounts) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermBatchAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermbatchaccountsResource, name), &v1alpha1.AzurermBatchAccount{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermBatchAccount), err
}

// List takes label and field selectors, and returns the list of AzurermBatchAccounts that match those selectors.
func (c *FakeAzurermBatchAccounts) List(opts v1.ListOptions) (result *v1alpha1.AzurermBatchAccountList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermbatchaccountsResource, azurermbatchaccountsKind, opts), &v1alpha1.AzurermBatchAccountList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermBatchAccountList{ListMeta: obj.(*v1alpha1.AzurermBatchAccountList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermBatchAccountList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermBatchAccounts.
func (c *FakeAzurermBatchAccounts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermbatchaccountsResource, opts))
}

// Create takes the representation of a azurermBatchAccount and creates it.  Returns the server's representation of the azurermBatchAccount, and an error, if there is any.
func (c *FakeAzurermBatchAccounts) Create(azurermBatchAccount *v1alpha1.AzurermBatchAccount) (result *v1alpha1.AzurermBatchAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermbatchaccountsResource, azurermBatchAccount), &v1alpha1.AzurermBatchAccount{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermBatchAccount), err
}

// Update takes the representation of a azurermBatchAccount and updates it. Returns the server's representation of the azurermBatchAccount, and an error, if there is any.
func (c *FakeAzurermBatchAccounts) Update(azurermBatchAccount *v1alpha1.AzurermBatchAccount) (result *v1alpha1.AzurermBatchAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermbatchaccountsResource, azurermBatchAccount), &v1alpha1.AzurermBatchAccount{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermBatchAccount), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermBatchAccounts) UpdateStatus(azurermBatchAccount *v1alpha1.AzurermBatchAccount) (*v1alpha1.AzurermBatchAccount, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermbatchaccountsResource, "status", azurermBatchAccount), &v1alpha1.AzurermBatchAccount{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermBatchAccount), err
}

// Delete takes name of the azurermBatchAccount and deletes it. Returns an error if one occurs.
func (c *FakeAzurermBatchAccounts) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermbatchaccountsResource, name), &v1alpha1.AzurermBatchAccount{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermBatchAccounts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermbatchaccountsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermBatchAccountList{})
	return err
}

// Patch applies the patch and returns the patched azurermBatchAccount.
func (c *FakeAzurermBatchAccounts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermBatchAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermbatchaccountsResource, name, pt, data, subresources...), &v1alpha1.AzurermBatchAccount{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermBatchAccount), err
}