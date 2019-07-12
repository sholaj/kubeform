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

// FakeAzurermKubernetesClusters implements AzurermKubernetesClusterInterface
type FakeAzurermKubernetesClusters struct {
	Fake *FakeAzurermV1alpha1
}

var azurermkubernetesclustersResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermkubernetesclusters"}

var azurermkubernetesclustersKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermKubernetesCluster"}

// Get takes name of the azurermKubernetesCluster, and returns the corresponding azurermKubernetesCluster object, and an error if there is any.
func (c *FakeAzurermKubernetesClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermKubernetesCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermkubernetesclustersResource, name), &v1alpha1.AzurermKubernetesCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermKubernetesCluster), err
}

// List takes label and field selectors, and returns the list of AzurermKubernetesClusters that match those selectors.
func (c *FakeAzurermKubernetesClusters) List(opts v1.ListOptions) (result *v1alpha1.AzurermKubernetesClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermkubernetesclustersResource, azurermkubernetesclustersKind, opts), &v1alpha1.AzurermKubernetesClusterList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermKubernetesClusterList{ListMeta: obj.(*v1alpha1.AzurermKubernetesClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermKubernetesClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermKubernetesClusters.
func (c *FakeAzurermKubernetesClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermkubernetesclustersResource, opts))
}

// Create takes the representation of a azurermKubernetesCluster and creates it.  Returns the server's representation of the azurermKubernetesCluster, and an error, if there is any.
func (c *FakeAzurermKubernetesClusters) Create(azurermKubernetesCluster *v1alpha1.AzurermKubernetesCluster) (result *v1alpha1.AzurermKubernetesCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermkubernetesclustersResource, azurermKubernetesCluster), &v1alpha1.AzurermKubernetesCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermKubernetesCluster), err
}

// Update takes the representation of a azurermKubernetesCluster and updates it. Returns the server's representation of the azurermKubernetesCluster, and an error, if there is any.
func (c *FakeAzurermKubernetesClusters) Update(azurermKubernetesCluster *v1alpha1.AzurermKubernetesCluster) (result *v1alpha1.AzurermKubernetesCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermkubernetesclustersResource, azurermKubernetesCluster), &v1alpha1.AzurermKubernetesCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermKubernetesCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermKubernetesClusters) UpdateStatus(azurermKubernetesCluster *v1alpha1.AzurermKubernetesCluster) (*v1alpha1.AzurermKubernetesCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermkubernetesclustersResource, "status", azurermKubernetesCluster), &v1alpha1.AzurermKubernetesCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermKubernetesCluster), err
}

// Delete takes name of the azurermKubernetesCluster and deletes it. Returns an error if one occurs.
func (c *FakeAzurermKubernetesClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermkubernetesclustersResource, name), &v1alpha1.AzurermKubernetesCluster{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermKubernetesClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermkubernetesclustersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermKubernetesClusterList{})
	return err
}

// Patch applies the patch and returns the patched azurermKubernetesCluster.
func (c *FakeAzurermKubernetesClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermKubernetesCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermkubernetesclustersResource, name, pt, data, subresources...), &v1alpha1.AzurermKubernetesCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermKubernetesCluster), err
}