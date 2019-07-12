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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeAwsDaxClusters implements AwsDaxClusterInterface
type FakeAwsDaxClusters struct {
	Fake *FakeAwsV1alpha1
}

var awsdaxclustersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsdaxclusters"}

var awsdaxclustersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsDaxCluster"}

// Get takes name of the awsDaxCluster, and returns the corresponding awsDaxCluster object, and an error if there is any.
func (c *FakeAwsDaxClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDaxCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsdaxclustersResource, name), &v1alpha1.AwsDaxCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDaxCluster), err
}

// List takes label and field selectors, and returns the list of AwsDaxClusters that match those selectors.
func (c *FakeAwsDaxClusters) List(opts v1.ListOptions) (result *v1alpha1.AwsDaxClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsdaxclustersResource, awsdaxclustersKind, opts), &v1alpha1.AwsDaxClusterList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsDaxClusterList{ListMeta: obj.(*v1alpha1.AwsDaxClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsDaxClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsDaxClusters.
func (c *FakeAwsDaxClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsdaxclustersResource, opts))
}

// Create takes the representation of a awsDaxCluster and creates it.  Returns the server's representation of the awsDaxCluster, and an error, if there is any.
func (c *FakeAwsDaxClusters) Create(awsDaxCluster *v1alpha1.AwsDaxCluster) (result *v1alpha1.AwsDaxCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsdaxclustersResource, awsDaxCluster), &v1alpha1.AwsDaxCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDaxCluster), err
}

// Update takes the representation of a awsDaxCluster and updates it. Returns the server's representation of the awsDaxCluster, and an error, if there is any.
func (c *FakeAwsDaxClusters) Update(awsDaxCluster *v1alpha1.AwsDaxCluster) (result *v1alpha1.AwsDaxCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsdaxclustersResource, awsDaxCluster), &v1alpha1.AwsDaxCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDaxCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsDaxClusters) UpdateStatus(awsDaxCluster *v1alpha1.AwsDaxCluster) (*v1alpha1.AwsDaxCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsdaxclustersResource, "status", awsDaxCluster), &v1alpha1.AwsDaxCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDaxCluster), err
}

// Delete takes name of the awsDaxCluster and deletes it. Returns an error if one occurs.
func (c *FakeAwsDaxClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsdaxclustersResource, name), &v1alpha1.AwsDaxCluster{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsDaxClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsdaxclustersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsDaxClusterList{})
	return err
}

// Patch applies the patch and returns the patched awsDaxCluster.
func (c *FakeAwsDaxClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDaxCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsdaxclustersResource, name, pt, data, subresources...), &v1alpha1.AwsDaxCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDaxCluster), err
}