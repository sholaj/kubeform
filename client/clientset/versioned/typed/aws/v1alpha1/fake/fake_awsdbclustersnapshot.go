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

// FakeAwsDbClusterSnapshots implements AwsDbClusterSnapshotInterface
type FakeAwsDbClusterSnapshots struct {
	Fake *FakeAwsV1alpha1
}

var awsdbclustersnapshotsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsdbclustersnapshots"}

var awsdbclustersnapshotsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsDbClusterSnapshot"}

// Get takes name of the awsDbClusterSnapshot, and returns the corresponding awsDbClusterSnapshot object, and an error if there is any.
func (c *FakeAwsDbClusterSnapshots) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDbClusterSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsdbclustersnapshotsResource, name), &v1alpha1.AwsDbClusterSnapshot{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDbClusterSnapshot), err
}

// List takes label and field selectors, and returns the list of AwsDbClusterSnapshots that match those selectors.
func (c *FakeAwsDbClusterSnapshots) List(opts v1.ListOptions) (result *v1alpha1.AwsDbClusterSnapshotList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsdbclustersnapshotsResource, awsdbclustersnapshotsKind, opts), &v1alpha1.AwsDbClusterSnapshotList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsDbClusterSnapshotList{ListMeta: obj.(*v1alpha1.AwsDbClusterSnapshotList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsDbClusterSnapshotList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsDbClusterSnapshots.
func (c *FakeAwsDbClusterSnapshots) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsdbclustersnapshotsResource, opts))
}

// Create takes the representation of a awsDbClusterSnapshot and creates it.  Returns the server's representation of the awsDbClusterSnapshot, and an error, if there is any.
func (c *FakeAwsDbClusterSnapshots) Create(awsDbClusterSnapshot *v1alpha1.AwsDbClusterSnapshot) (result *v1alpha1.AwsDbClusterSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsdbclustersnapshotsResource, awsDbClusterSnapshot), &v1alpha1.AwsDbClusterSnapshot{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDbClusterSnapshot), err
}

// Update takes the representation of a awsDbClusterSnapshot and updates it. Returns the server's representation of the awsDbClusterSnapshot, and an error, if there is any.
func (c *FakeAwsDbClusterSnapshots) Update(awsDbClusterSnapshot *v1alpha1.AwsDbClusterSnapshot) (result *v1alpha1.AwsDbClusterSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsdbclustersnapshotsResource, awsDbClusterSnapshot), &v1alpha1.AwsDbClusterSnapshot{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDbClusterSnapshot), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsDbClusterSnapshots) UpdateStatus(awsDbClusterSnapshot *v1alpha1.AwsDbClusterSnapshot) (*v1alpha1.AwsDbClusterSnapshot, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsdbclustersnapshotsResource, "status", awsDbClusterSnapshot), &v1alpha1.AwsDbClusterSnapshot{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDbClusterSnapshot), err
}

// Delete takes name of the awsDbClusterSnapshot and deletes it. Returns an error if one occurs.
func (c *FakeAwsDbClusterSnapshots) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsdbclustersnapshotsResource, name), &v1alpha1.AwsDbClusterSnapshot{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsDbClusterSnapshots) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsdbclustersnapshotsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsDbClusterSnapshotList{})
	return err
}

// Patch applies the patch and returns the patched awsDbClusterSnapshot.
func (c *FakeAwsDbClusterSnapshots) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDbClusterSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsdbclustersnapshotsResource, name, pt, data, subresources...), &v1alpha1.AwsDbClusterSnapshot{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDbClusterSnapshot), err
}