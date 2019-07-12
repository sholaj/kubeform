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

// FakeAzurermServiceFabricClusters implements AzurermServiceFabricClusterInterface
type FakeAzurermServiceFabricClusters struct {
	Fake *FakeAzurermV1alpha1
}

var azurermservicefabricclustersResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermservicefabricclusters"}

var azurermservicefabricclustersKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermServiceFabricCluster"}

// Get takes name of the azurermServiceFabricCluster, and returns the corresponding azurermServiceFabricCluster object, and an error if there is any.
func (c *FakeAzurermServiceFabricClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermServiceFabricCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermservicefabricclustersResource, name), &v1alpha1.AzurermServiceFabricCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermServiceFabricCluster), err
}

// List takes label and field selectors, and returns the list of AzurermServiceFabricClusters that match those selectors.
func (c *FakeAzurermServiceFabricClusters) List(opts v1.ListOptions) (result *v1alpha1.AzurermServiceFabricClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermservicefabricclustersResource, azurermservicefabricclustersKind, opts), &v1alpha1.AzurermServiceFabricClusterList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermServiceFabricClusterList{ListMeta: obj.(*v1alpha1.AzurermServiceFabricClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermServiceFabricClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermServiceFabricClusters.
func (c *FakeAzurermServiceFabricClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermservicefabricclustersResource, opts))
}

// Create takes the representation of a azurermServiceFabricCluster and creates it.  Returns the server's representation of the azurermServiceFabricCluster, and an error, if there is any.
func (c *FakeAzurermServiceFabricClusters) Create(azurermServiceFabricCluster *v1alpha1.AzurermServiceFabricCluster) (result *v1alpha1.AzurermServiceFabricCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermservicefabricclustersResource, azurermServiceFabricCluster), &v1alpha1.AzurermServiceFabricCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermServiceFabricCluster), err
}

// Update takes the representation of a azurermServiceFabricCluster and updates it. Returns the server's representation of the azurermServiceFabricCluster, and an error, if there is any.
func (c *FakeAzurermServiceFabricClusters) Update(azurermServiceFabricCluster *v1alpha1.AzurermServiceFabricCluster) (result *v1alpha1.AzurermServiceFabricCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermservicefabricclustersResource, azurermServiceFabricCluster), &v1alpha1.AzurermServiceFabricCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermServiceFabricCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermServiceFabricClusters) UpdateStatus(azurermServiceFabricCluster *v1alpha1.AzurermServiceFabricCluster) (*v1alpha1.AzurermServiceFabricCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermservicefabricclustersResource, "status", azurermServiceFabricCluster), &v1alpha1.AzurermServiceFabricCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermServiceFabricCluster), err
}

// Delete takes name of the azurermServiceFabricCluster and deletes it. Returns an error if one occurs.
func (c *FakeAzurermServiceFabricClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermservicefabricclustersResource, name), &v1alpha1.AzurermServiceFabricCluster{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermServiceFabricClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermservicefabricclustersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermServiceFabricClusterList{})
	return err
}

// Patch applies the patch and returns the patched azurermServiceFabricCluster.
func (c *FakeAzurermServiceFabricClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermServiceFabricCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermservicefabricclustersResource, name, pt, data, subresources...), &v1alpha1.AzurermServiceFabricCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermServiceFabricCluster), err
}