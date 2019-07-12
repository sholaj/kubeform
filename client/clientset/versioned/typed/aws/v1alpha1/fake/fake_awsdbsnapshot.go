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

// FakeAwsDbSnapshots implements AwsDbSnapshotInterface
type FakeAwsDbSnapshots struct {
	Fake *FakeAwsV1alpha1
}

var awsdbsnapshotsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsdbsnapshots"}

var awsdbsnapshotsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsDbSnapshot"}

// Get takes name of the awsDbSnapshot, and returns the corresponding awsDbSnapshot object, and an error if there is any.
func (c *FakeAwsDbSnapshots) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDbSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsdbsnapshotsResource, name), &v1alpha1.AwsDbSnapshot{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDbSnapshot), err
}

// List takes label and field selectors, and returns the list of AwsDbSnapshots that match those selectors.
func (c *FakeAwsDbSnapshots) List(opts v1.ListOptions) (result *v1alpha1.AwsDbSnapshotList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsdbsnapshotsResource, awsdbsnapshotsKind, opts), &v1alpha1.AwsDbSnapshotList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsDbSnapshotList{ListMeta: obj.(*v1alpha1.AwsDbSnapshotList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsDbSnapshotList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsDbSnapshots.
func (c *FakeAwsDbSnapshots) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsdbsnapshotsResource, opts))
}

// Create takes the representation of a awsDbSnapshot and creates it.  Returns the server's representation of the awsDbSnapshot, and an error, if there is any.
func (c *FakeAwsDbSnapshots) Create(awsDbSnapshot *v1alpha1.AwsDbSnapshot) (result *v1alpha1.AwsDbSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsdbsnapshotsResource, awsDbSnapshot), &v1alpha1.AwsDbSnapshot{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDbSnapshot), err
}

// Update takes the representation of a awsDbSnapshot and updates it. Returns the server's representation of the awsDbSnapshot, and an error, if there is any.
func (c *FakeAwsDbSnapshots) Update(awsDbSnapshot *v1alpha1.AwsDbSnapshot) (result *v1alpha1.AwsDbSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsdbsnapshotsResource, awsDbSnapshot), &v1alpha1.AwsDbSnapshot{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDbSnapshot), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsDbSnapshots) UpdateStatus(awsDbSnapshot *v1alpha1.AwsDbSnapshot) (*v1alpha1.AwsDbSnapshot, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsdbsnapshotsResource, "status", awsDbSnapshot), &v1alpha1.AwsDbSnapshot{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDbSnapshot), err
}

// Delete takes name of the awsDbSnapshot and deletes it. Returns an error if one occurs.
func (c *FakeAwsDbSnapshots) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsdbsnapshotsResource, name), &v1alpha1.AwsDbSnapshot{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsDbSnapshots) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsdbsnapshotsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsDbSnapshotList{})
	return err
}

// Patch applies the patch and returns the patched awsDbSnapshot.
func (c *FakeAwsDbSnapshots) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDbSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsdbsnapshotsResource, name, pt, data, subresources...), &v1alpha1.AwsDbSnapshot{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDbSnapshot), err
}