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

// FakeAwsDocdbClusters implements AwsDocdbClusterInterface
type FakeAwsDocdbClusters struct {
	Fake *FakeAwsV1alpha1
}

var awsdocdbclustersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsdocdbclusters"}

var awsdocdbclustersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsDocdbCluster"}

// Get takes name of the awsDocdbCluster, and returns the corresponding awsDocdbCluster object, and an error if there is any.
func (c *FakeAwsDocdbClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDocdbCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsdocdbclustersResource, name), &v1alpha1.AwsDocdbCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDocdbCluster), err
}

// List takes label and field selectors, and returns the list of AwsDocdbClusters that match those selectors.
func (c *FakeAwsDocdbClusters) List(opts v1.ListOptions) (result *v1alpha1.AwsDocdbClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsdocdbclustersResource, awsdocdbclustersKind, opts), &v1alpha1.AwsDocdbClusterList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsDocdbClusterList{ListMeta: obj.(*v1alpha1.AwsDocdbClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsDocdbClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsDocdbClusters.
func (c *FakeAwsDocdbClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsdocdbclustersResource, opts))
}

// Create takes the representation of a awsDocdbCluster and creates it.  Returns the server's representation of the awsDocdbCluster, and an error, if there is any.
func (c *FakeAwsDocdbClusters) Create(awsDocdbCluster *v1alpha1.AwsDocdbCluster) (result *v1alpha1.AwsDocdbCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsdocdbclustersResource, awsDocdbCluster), &v1alpha1.AwsDocdbCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDocdbCluster), err
}

// Update takes the representation of a awsDocdbCluster and updates it. Returns the server's representation of the awsDocdbCluster, and an error, if there is any.
func (c *FakeAwsDocdbClusters) Update(awsDocdbCluster *v1alpha1.AwsDocdbCluster) (result *v1alpha1.AwsDocdbCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsdocdbclustersResource, awsDocdbCluster), &v1alpha1.AwsDocdbCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDocdbCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsDocdbClusters) UpdateStatus(awsDocdbCluster *v1alpha1.AwsDocdbCluster) (*v1alpha1.AwsDocdbCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsdocdbclustersResource, "status", awsDocdbCluster), &v1alpha1.AwsDocdbCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDocdbCluster), err
}

// Delete takes name of the awsDocdbCluster and deletes it. Returns an error if one occurs.
func (c *FakeAwsDocdbClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsdocdbclustersResource, name), &v1alpha1.AwsDocdbCluster{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsDocdbClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsdocdbclustersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsDocdbClusterList{})
	return err
}

// Patch applies the patch and returns the patched awsDocdbCluster.
func (c *FakeAwsDocdbClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDocdbCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsdocdbclustersResource, name, pt, data, subresources...), &v1alpha1.AwsDocdbCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDocdbCluster), err
}