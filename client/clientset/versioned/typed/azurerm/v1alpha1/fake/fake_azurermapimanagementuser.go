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

// FakeAzurermApiManagementUsers implements AzurermApiManagementUserInterface
type FakeAzurermApiManagementUsers struct {
	Fake *FakeAzurermV1alpha1
}

var azurermapimanagementusersResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermapimanagementusers"}

var azurermapimanagementusersKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermApiManagementUser"}

// Get takes name of the azurermApiManagementUser, and returns the corresponding azurermApiManagementUser object, and an error if there is any.
func (c *FakeAzurermApiManagementUsers) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermApiManagementUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermapimanagementusersResource, name), &v1alpha1.AzurermApiManagementUser{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApiManagementUser), err
}

// List takes label and field selectors, and returns the list of AzurermApiManagementUsers that match those selectors.
func (c *FakeAzurermApiManagementUsers) List(opts v1.ListOptions) (result *v1alpha1.AzurermApiManagementUserList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermapimanagementusersResource, azurermapimanagementusersKind, opts), &v1alpha1.AzurermApiManagementUserList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermApiManagementUserList{ListMeta: obj.(*v1alpha1.AzurermApiManagementUserList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermApiManagementUserList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermApiManagementUsers.
func (c *FakeAzurermApiManagementUsers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermapimanagementusersResource, opts))
}

// Create takes the representation of a azurermApiManagementUser and creates it.  Returns the server's representation of the azurermApiManagementUser, and an error, if there is any.
func (c *FakeAzurermApiManagementUsers) Create(azurermApiManagementUser *v1alpha1.AzurermApiManagementUser) (result *v1alpha1.AzurermApiManagementUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermapimanagementusersResource, azurermApiManagementUser), &v1alpha1.AzurermApiManagementUser{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApiManagementUser), err
}

// Update takes the representation of a azurermApiManagementUser and updates it. Returns the server's representation of the azurermApiManagementUser, and an error, if there is any.
func (c *FakeAzurermApiManagementUsers) Update(azurermApiManagementUser *v1alpha1.AzurermApiManagementUser) (result *v1alpha1.AzurermApiManagementUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermapimanagementusersResource, azurermApiManagementUser), &v1alpha1.AzurermApiManagementUser{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApiManagementUser), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermApiManagementUsers) UpdateStatus(azurermApiManagementUser *v1alpha1.AzurermApiManagementUser) (*v1alpha1.AzurermApiManagementUser, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermapimanagementusersResource, "status", azurermApiManagementUser), &v1alpha1.AzurermApiManagementUser{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApiManagementUser), err
}

// Delete takes name of the azurermApiManagementUser and deletes it. Returns an error if one occurs.
func (c *FakeAzurermApiManagementUsers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermapimanagementusersResource, name), &v1alpha1.AzurermApiManagementUser{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermApiManagementUsers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermapimanagementusersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermApiManagementUserList{})
	return err
}

// Patch applies the patch and returns the patched azurermApiManagementUser.
func (c *FakeAzurermApiManagementUsers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermApiManagementUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermapimanagementusersResource, name, pt, data, subresources...), &v1alpha1.AzurermApiManagementUser{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApiManagementUser), err
}