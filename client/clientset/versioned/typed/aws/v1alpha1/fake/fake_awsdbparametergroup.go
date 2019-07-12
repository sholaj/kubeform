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

// FakeAwsDbParameterGroups implements AwsDbParameterGroupInterface
type FakeAwsDbParameterGroups struct {
	Fake *FakeAwsV1alpha1
}

var awsdbparametergroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsdbparametergroups"}

var awsdbparametergroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsDbParameterGroup"}

// Get takes name of the awsDbParameterGroup, and returns the corresponding awsDbParameterGroup object, and an error if there is any.
func (c *FakeAwsDbParameterGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDbParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsdbparametergroupsResource, name), &v1alpha1.AwsDbParameterGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDbParameterGroup), err
}

// List takes label and field selectors, and returns the list of AwsDbParameterGroups that match those selectors.
func (c *FakeAwsDbParameterGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsDbParameterGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsdbparametergroupsResource, awsdbparametergroupsKind, opts), &v1alpha1.AwsDbParameterGroupList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsDbParameterGroupList{ListMeta: obj.(*v1alpha1.AwsDbParameterGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsDbParameterGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsDbParameterGroups.
func (c *FakeAwsDbParameterGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsdbparametergroupsResource, opts))
}

// Create takes the representation of a awsDbParameterGroup and creates it.  Returns the server's representation of the awsDbParameterGroup, and an error, if there is any.
func (c *FakeAwsDbParameterGroups) Create(awsDbParameterGroup *v1alpha1.AwsDbParameterGroup) (result *v1alpha1.AwsDbParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsdbparametergroupsResource, awsDbParameterGroup), &v1alpha1.AwsDbParameterGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDbParameterGroup), err
}

// Update takes the representation of a awsDbParameterGroup and updates it. Returns the server's representation of the awsDbParameterGroup, and an error, if there is any.
func (c *FakeAwsDbParameterGroups) Update(awsDbParameterGroup *v1alpha1.AwsDbParameterGroup) (result *v1alpha1.AwsDbParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsdbparametergroupsResource, awsDbParameterGroup), &v1alpha1.AwsDbParameterGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDbParameterGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsDbParameterGroups) UpdateStatus(awsDbParameterGroup *v1alpha1.AwsDbParameterGroup) (*v1alpha1.AwsDbParameterGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsdbparametergroupsResource, "status", awsDbParameterGroup), &v1alpha1.AwsDbParameterGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDbParameterGroup), err
}

// Delete takes name of the awsDbParameterGroup and deletes it. Returns an error if one occurs.
func (c *FakeAwsDbParameterGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsdbparametergroupsResource, name), &v1alpha1.AwsDbParameterGroup{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsDbParameterGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsdbparametergroupsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsDbParameterGroupList{})
	return err
}

// Patch applies the patch and returns the patched awsDbParameterGroup.
func (c *FakeAwsDbParameterGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDbParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsdbparametergroupsResource, name, pt, data, subresources...), &v1alpha1.AwsDbParameterGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDbParameterGroup), err
}