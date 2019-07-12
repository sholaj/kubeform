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

// FakeAzurermDataLakeAnalyticsAccounts implements AzurermDataLakeAnalyticsAccountInterface
type FakeAzurermDataLakeAnalyticsAccounts struct {
	Fake *FakeAzurermV1alpha1
}

var azurermdatalakeanalyticsaccountsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermdatalakeanalyticsaccounts"}

var azurermdatalakeanalyticsaccountsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermDataLakeAnalyticsAccount"}

// Get takes name of the azurermDataLakeAnalyticsAccount, and returns the corresponding azurermDataLakeAnalyticsAccount object, and an error if there is any.
func (c *FakeAzurermDataLakeAnalyticsAccounts) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermDataLakeAnalyticsAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermdatalakeanalyticsaccountsResource, name), &v1alpha1.AzurermDataLakeAnalyticsAccount{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDataLakeAnalyticsAccount), err
}

// List takes label and field selectors, and returns the list of AzurermDataLakeAnalyticsAccounts that match those selectors.
func (c *FakeAzurermDataLakeAnalyticsAccounts) List(opts v1.ListOptions) (result *v1alpha1.AzurermDataLakeAnalyticsAccountList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermdatalakeanalyticsaccountsResource, azurermdatalakeanalyticsaccountsKind, opts), &v1alpha1.AzurermDataLakeAnalyticsAccountList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermDataLakeAnalyticsAccountList{ListMeta: obj.(*v1alpha1.AzurermDataLakeAnalyticsAccountList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermDataLakeAnalyticsAccountList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermDataLakeAnalyticsAccounts.
func (c *FakeAzurermDataLakeAnalyticsAccounts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermdatalakeanalyticsaccountsResource, opts))
}

// Create takes the representation of a azurermDataLakeAnalyticsAccount and creates it.  Returns the server's representation of the azurermDataLakeAnalyticsAccount, and an error, if there is any.
func (c *FakeAzurermDataLakeAnalyticsAccounts) Create(azurermDataLakeAnalyticsAccount *v1alpha1.AzurermDataLakeAnalyticsAccount) (result *v1alpha1.AzurermDataLakeAnalyticsAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermdatalakeanalyticsaccountsResource, azurermDataLakeAnalyticsAccount), &v1alpha1.AzurermDataLakeAnalyticsAccount{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDataLakeAnalyticsAccount), err
}

// Update takes the representation of a azurermDataLakeAnalyticsAccount and updates it. Returns the server's representation of the azurermDataLakeAnalyticsAccount, and an error, if there is any.
func (c *FakeAzurermDataLakeAnalyticsAccounts) Update(azurermDataLakeAnalyticsAccount *v1alpha1.AzurermDataLakeAnalyticsAccount) (result *v1alpha1.AzurermDataLakeAnalyticsAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermdatalakeanalyticsaccountsResource, azurermDataLakeAnalyticsAccount), &v1alpha1.AzurermDataLakeAnalyticsAccount{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDataLakeAnalyticsAccount), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermDataLakeAnalyticsAccounts) UpdateStatus(azurermDataLakeAnalyticsAccount *v1alpha1.AzurermDataLakeAnalyticsAccount) (*v1alpha1.AzurermDataLakeAnalyticsAccount, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermdatalakeanalyticsaccountsResource, "status", azurermDataLakeAnalyticsAccount), &v1alpha1.AzurermDataLakeAnalyticsAccount{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDataLakeAnalyticsAccount), err
}

// Delete takes name of the azurermDataLakeAnalyticsAccount and deletes it. Returns an error if one occurs.
func (c *FakeAzurermDataLakeAnalyticsAccounts) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermdatalakeanalyticsaccountsResource, name), &v1alpha1.AzurermDataLakeAnalyticsAccount{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermDataLakeAnalyticsAccounts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermdatalakeanalyticsaccountsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermDataLakeAnalyticsAccountList{})
	return err
}

// Patch applies the patch and returns the patched azurermDataLakeAnalyticsAccount.
func (c *FakeAzurermDataLakeAnalyticsAccounts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDataLakeAnalyticsAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermdatalakeanalyticsaccountsResource, name, pt, data, subresources...), &v1alpha1.AzurermDataLakeAnalyticsAccount{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDataLakeAnalyticsAccount), err
}