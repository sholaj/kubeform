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

// FakeAzurermDataFactoryLinkedServiceSqlServers implements AzurermDataFactoryLinkedServiceSqlServerInterface
type FakeAzurermDataFactoryLinkedServiceSqlServers struct {
	Fake *FakeAzurermV1alpha1
}

var azurermdatafactorylinkedservicesqlserversResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermdatafactorylinkedservicesqlservers"}

var azurermdatafactorylinkedservicesqlserversKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermDataFactoryLinkedServiceSqlServer"}

// Get takes name of the azurermDataFactoryLinkedServiceSqlServer, and returns the corresponding azurermDataFactoryLinkedServiceSqlServer object, and an error if there is any.
func (c *FakeAzurermDataFactoryLinkedServiceSqlServers) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermDataFactoryLinkedServiceSqlServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermdatafactorylinkedservicesqlserversResource, name), &v1alpha1.AzurermDataFactoryLinkedServiceSqlServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDataFactoryLinkedServiceSqlServer), err
}

// List takes label and field selectors, and returns the list of AzurermDataFactoryLinkedServiceSqlServers that match those selectors.
func (c *FakeAzurermDataFactoryLinkedServiceSqlServers) List(opts v1.ListOptions) (result *v1alpha1.AzurermDataFactoryLinkedServiceSqlServerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermdatafactorylinkedservicesqlserversResource, azurermdatafactorylinkedservicesqlserversKind, opts), &v1alpha1.AzurermDataFactoryLinkedServiceSqlServerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermDataFactoryLinkedServiceSqlServerList{ListMeta: obj.(*v1alpha1.AzurermDataFactoryLinkedServiceSqlServerList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermDataFactoryLinkedServiceSqlServerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermDataFactoryLinkedServiceSqlServers.
func (c *FakeAzurermDataFactoryLinkedServiceSqlServers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermdatafactorylinkedservicesqlserversResource, opts))
}

// Create takes the representation of a azurermDataFactoryLinkedServiceSqlServer and creates it.  Returns the server's representation of the azurermDataFactoryLinkedServiceSqlServer, and an error, if there is any.
func (c *FakeAzurermDataFactoryLinkedServiceSqlServers) Create(azurermDataFactoryLinkedServiceSqlServer *v1alpha1.AzurermDataFactoryLinkedServiceSqlServer) (result *v1alpha1.AzurermDataFactoryLinkedServiceSqlServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermdatafactorylinkedservicesqlserversResource, azurermDataFactoryLinkedServiceSqlServer), &v1alpha1.AzurermDataFactoryLinkedServiceSqlServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDataFactoryLinkedServiceSqlServer), err
}

// Update takes the representation of a azurermDataFactoryLinkedServiceSqlServer and updates it. Returns the server's representation of the azurermDataFactoryLinkedServiceSqlServer, and an error, if there is any.
func (c *FakeAzurermDataFactoryLinkedServiceSqlServers) Update(azurermDataFactoryLinkedServiceSqlServer *v1alpha1.AzurermDataFactoryLinkedServiceSqlServer) (result *v1alpha1.AzurermDataFactoryLinkedServiceSqlServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermdatafactorylinkedservicesqlserversResource, azurermDataFactoryLinkedServiceSqlServer), &v1alpha1.AzurermDataFactoryLinkedServiceSqlServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDataFactoryLinkedServiceSqlServer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermDataFactoryLinkedServiceSqlServers) UpdateStatus(azurermDataFactoryLinkedServiceSqlServer *v1alpha1.AzurermDataFactoryLinkedServiceSqlServer) (*v1alpha1.AzurermDataFactoryLinkedServiceSqlServer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermdatafactorylinkedservicesqlserversResource, "status", azurermDataFactoryLinkedServiceSqlServer), &v1alpha1.AzurermDataFactoryLinkedServiceSqlServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDataFactoryLinkedServiceSqlServer), err
}

// Delete takes name of the azurermDataFactoryLinkedServiceSqlServer and deletes it. Returns an error if one occurs.
func (c *FakeAzurermDataFactoryLinkedServiceSqlServers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermdatafactorylinkedservicesqlserversResource, name), &v1alpha1.AzurermDataFactoryLinkedServiceSqlServer{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermDataFactoryLinkedServiceSqlServers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermdatafactorylinkedservicesqlserversResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermDataFactoryLinkedServiceSqlServerList{})
	return err
}

// Patch applies the patch and returns the patched azurermDataFactoryLinkedServiceSqlServer.
func (c *FakeAzurermDataFactoryLinkedServiceSqlServers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDataFactoryLinkedServiceSqlServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermdatafactorylinkedservicesqlserversResource, name, pt, data, subresources...), &v1alpha1.AzurermDataFactoryLinkedServiceSqlServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDataFactoryLinkedServiceSqlServer), err
}