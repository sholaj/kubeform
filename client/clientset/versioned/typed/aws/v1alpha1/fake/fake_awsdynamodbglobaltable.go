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

// FakeAwsDynamodbGlobalTables implements AwsDynamodbGlobalTableInterface
type FakeAwsDynamodbGlobalTables struct {
	Fake *FakeAwsV1alpha1
}

var awsdynamodbglobaltablesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsdynamodbglobaltables"}

var awsdynamodbglobaltablesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsDynamodbGlobalTable"}

// Get takes name of the awsDynamodbGlobalTable, and returns the corresponding awsDynamodbGlobalTable object, and an error if there is any.
func (c *FakeAwsDynamodbGlobalTables) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDynamodbGlobalTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsdynamodbglobaltablesResource, name), &v1alpha1.AwsDynamodbGlobalTable{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDynamodbGlobalTable), err
}

// List takes label and field selectors, and returns the list of AwsDynamodbGlobalTables that match those selectors.
func (c *FakeAwsDynamodbGlobalTables) List(opts v1.ListOptions) (result *v1alpha1.AwsDynamodbGlobalTableList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsdynamodbglobaltablesResource, awsdynamodbglobaltablesKind, opts), &v1alpha1.AwsDynamodbGlobalTableList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsDynamodbGlobalTableList{ListMeta: obj.(*v1alpha1.AwsDynamodbGlobalTableList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsDynamodbGlobalTableList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsDynamodbGlobalTables.
func (c *FakeAwsDynamodbGlobalTables) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsdynamodbglobaltablesResource, opts))
}

// Create takes the representation of a awsDynamodbGlobalTable and creates it.  Returns the server's representation of the awsDynamodbGlobalTable, and an error, if there is any.
func (c *FakeAwsDynamodbGlobalTables) Create(awsDynamodbGlobalTable *v1alpha1.AwsDynamodbGlobalTable) (result *v1alpha1.AwsDynamodbGlobalTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsdynamodbglobaltablesResource, awsDynamodbGlobalTable), &v1alpha1.AwsDynamodbGlobalTable{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDynamodbGlobalTable), err
}

// Update takes the representation of a awsDynamodbGlobalTable and updates it. Returns the server's representation of the awsDynamodbGlobalTable, and an error, if there is any.
func (c *FakeAwsDynamodbGlobalTables) Update(awsDynamodbGlobalTable *v1alpha1.AwsDynamodbGlobalTable) (result *v1alpha1.AwsDynamodbGlobalTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsdynamodbglobaltablesResource, awsDynamodbGlobalTable), &v1alpha1.AwsDynamodbGlobalTable{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDynamodbGlobalTable), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsDynamodbGlobalTables) UpdateStatus(awsDynamodbGlobalTable *v1alpha1.AwsDynamodbGlobalTable) (*v1alpha1.AwsDynamodbGlobalTable, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsdynamodbglobaltablesResource, "status", awsDynamodbGlobalTable), &v1alpha1.AwsDynamodbGlobalTable{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDynamodbGlobalTable), err
}

// Delete takes name of the awsDynamodbGlobalTable and deletes it. Returns an error if one occurs.
func (c *FakeAwsDynamodbGlobalTables) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsdynamodbglobaltablesResource, name), &v1alpha1.AwsDynamodbGlobalTable{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsDynamodbGlobalTables) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsdynamodbglobaltablesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsDynamodbGlobalTableList{})
	return err
}

// Patch applies the patch and returns the patched awsDynamodbGlobalTable.
func (c *FakeAwsDynamodbGlobalTables) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDynamodbGlobalTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsdynamodbglobaltablesResource, name, pt, data, subresources...), &v1alpha1.AwsDynamodbGlobalTable{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDynamodbGlobalTable), err
}