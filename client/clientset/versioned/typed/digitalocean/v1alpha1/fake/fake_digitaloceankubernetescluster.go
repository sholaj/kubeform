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
	v1alpha1 "kubeform.dev/kubeform/apis/digitalocean/v1alpha1"
)

// FakeDigitaloceanKubernetesClusters implements DigitaloceanKubernetesClusterInterface
type FakeDigitaloceanKubernetesClusters struct {
	Fake *FakeDigitaloceanV1alpha1
}

var digitaloceankubernetesclustersResource = schema.GroupVersionResource{Group: "digitalocean.kubeform.com", Version: "v1alpha1", Resource: "digitaloceankubernetesclusters"}

var digitaloceankubernetesclustersKind = schema.GroupVersionKind{Group: "digitalocean.kubeform.com", Version: "v1alpha1", Kind: "DigitaloceanKubernetesCluster"}

// Get takes name of the digitaloceanKubernetesCluster, and returns the corresponding digitaloceanKubernetesCluster object, and an error if there is any.
func (c *FakeDigitaloceanKubernetesClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.DigitaloceanKubernetesCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(digitaloceankubernetesclustersResource, name), &v1alpha1.DigitaloceanKubernetesCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanKubernetesCluster), err
}

// List takes label and field selectors, and returns the list of DigitaloceanKubernetesClusters that match those selectors.
func (c *FakeDigitaloceanKubernetesClusters) List(opts v1.ListOptions) (result *v1alpha1.DigitaloceanKubernetesClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(digitaloceankubernetesclustersResource, digitaloceankubernetesclustersKind, opts), &v1alpha1.DigitaloceanKubernetesClusterList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DigitaloceanKubernetesClusterList{ListMeta: obj.(*v1alpha1.DigitaloceanKubernetesClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.DigitaloceanKubernetesClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested digitaloceanKubernetesClusters.
func (c *FakeDigitaloceanKubernetesClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(digitaloceankubernetesclustersResource, opts))
}

// Create takes the representation of a digitaloceanKubernetesCluster and creates it.  Returns the server's representation of the digitaloceanKubernetesCluster, and an error, if there is any.
func (c *FakeDigitaloceanKubernetesClusters) Create(digitaloceanKubernetesCluster *v1alpha1.DigitaloceanKubernetesCluster) (result *v1alpha1.DigitaloceanKubernetesCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(digitaloceankubernetesclustersResource, digitaloceanKubernetesCluster), &v1alpha1.DigitaloceanKubernetesCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanKubernetesCluster), err
}

// Update takes the representation of a digitaloceanKubernetesCluster and updates it. Returns the server's representation of the digitaloceanKubernetesCluster, and an error, if there is any.
func (c *FakeDigitaloceanKubernetesClusters) Update(digitaloceanKubernetesCluster *v1alpha1.DigitaloceanKubernetesCluster) (result *v1alpha1.DigitaloceanKubernetesCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(digitaloceankubernetesclustersResource, digitaloceanKubernetesCluster), &v1alpha1.DigitaloceanKubernetesCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanKubernetesCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDigitaloceanKubernetesClusters) UpdateStatus(digitaloceanKubernetesCluster *v1alpha1.DigitaloceanKubernetesCluster) (*v1alpha1.DigitaloceanKubernetesCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(digitaloceankubernetesclustersResource, "status", digitaloceanKubernetesCluster), &v1alpha1.DigitaloceanKubernetesCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanKubernetesCluster), err
}

// Delete takes name of the digitaloceanKubernetesCluster and deletes it. Returns an error if one occurs.
func (c *FakeDigitaloceanKubernetesClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(digitaloceankubernetesclustersResource, name), &v1alpha1.DigitaloceanKubernetesCluster{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDigitaloceanKubernetesClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(digitaloceankubernetesclustersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DigitaloceanKubernetesClusterList{})
	return err
}

// Patch applies the patch and returns the patched digitaloceanKubernetesCluster.
func (c *FakeDigitaloceanKubernetesClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DigitaloceanKubernetesCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(digitaloceankubernetesclustersResource, name, pt, data, subresources...), &v1alpha1.DigitaloceanKubernetesCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanKubernetesCluster), err
}