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

// FakeAwsInspectorResourceGroups implements AwsInspectorResourceGroupInterface
type FakeAwsInspectorResourceGroups struct {
	Fake *FakeAwsV1alpha1
}

var awsinspectorresourcegroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsinspectorresourcegroups"}

var awsinspectorresourcegroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsInspectorResourceGroup"}

// Get takes name of the awsInspectorResourceGroup, and returns the corresponding awsInspectorResourceGroup object, and an error if there is any.
func (c *FakeAwsInspectorResourceGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsInspectorResourceGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsinspectorresourcegroupsResource, name), &v1alpha1.AwsInspectorResourceGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsInspectorResourceGroup), err
}

// List takes label and field selectors, and returns the list of AwsInspectorResourceGroups that match those selectors.
func (c *FakeAwsInspectorResourceGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsInspectorResourceGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsinspectorresourcegroupsResource, awsinspectorresourcegroupsKind, opts), &v1alpha1.AwsInspectorResourceGroupList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsInspectorResourceGroupList{ListMeta: obj.(*v1alpha1.AwsInspectorResourceGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsInspectorResourceGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsInspectorResourceGroups.
func (c *FakeAwsInspectorResourceGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsinspectorresourcegroupsResource, opts))
}

// Create takes the representation of a awsInspectorResourceGroup and creates it.  Returns the server's representation of the awsInspectorResourceGroup, and an error, if there is any.
func (c *FakeAwsInspectorResourceGroups) Create(awsInspectorResourceGroup *v1alpha1.AwsInspectorResourceGroup) (result *v1alpha1.AwsInspectorResourceGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsinspectorresourcegroupsResource, awsInspectorResourceGroup), &v1alpha1.AwsInspectorResourceGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsInspectorResourceGroup), err
}

// Update takes the representation of a awsInspectorResourceGroup and updates it. Returns the server's representation of the awsInspectorResourceGroup, and an error, if there is any.
func (c *FakeAwsInspectorResourceGroups) Update(awsInspectorResourceGroup *v1alpha1.AwsInspectorResourceGroup) (result *v1alpha1.AwsInspectorResourceGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsinspectorresourcegroupsResource, awsInspectorResourceGroup), &v1alpha1.AwsInspectorResourceGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsInspectorResourceGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsInspectorResourceGroups) UpdateStatus(awsInspectorResourceGroup *v1alpha1.AwsInspectorResourceGroup) (*v1alpha1.AwsInspectorResourceGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsinspectorresourcegroupsResource, "status", awsInspectorResourceGroup), &v1alpha1.AwsInspectorResourceGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsInspectorResourceGroup), err
}

// Delete takes name of the awsInspectorResourceGroup and deletes it. Returns an error if one occurs.
func (c *FakeAwsInspectorResourceGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsinspectorresourcegroupsResource, name), &v1alpha1.AwsInspectorResourceGroup{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsInspectorResourceGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsinspectorresourcegroupsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsInspectorResourceGroupList{})
	return err
}

// Patch applies the patch and returns the patched awsInspectorResourceGroup.
func (c *FakeAwsInspectorResourceGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsInspectorResourceGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsinspectorresourcegroupsResource, name, pt, data, subresources...), &v1alpha1.AwsInspectorResourceGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsInspectorResourceGroup), err
}