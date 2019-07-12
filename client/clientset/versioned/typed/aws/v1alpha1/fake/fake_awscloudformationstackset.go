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

// FakeAwsCloudformationStackSets implements AwsCloudformationStackSetInterface
type FakeAwsCloudformationStackSets struct {
	Fake *FakeAwsV1alpha1
}

var awscloudformationstacksetsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awscloudformationstacksets"}

var awscloudformationstacksetsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsCloudformationStackSet"}

// Get takes name of the awsCloudformationStackSet, and returns the corresponding awsCloudformationStackSet object, and an error if there is any.
func (c *FakeAwsCloudformationStackSets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsCloudformationStackSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awscloudformationstacksetsResource, name), &v1alpha1.AwsCloudformationStackSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCloudformationStackSet), err
}

// List takes label and field selectors, and returns the list of AwsCloudformationStackSets that match those selectors.
func (c *FakeAwsCloudformationStackSets) List(opts v1.ListOptions) (result *v1alpha1.AwsCloudformationStackSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awscloudformationstacksetsResource, awscloudformationstacksetsKind, opts), &v1alpha1.AwsCloudformationStackSetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsCloudformationStackSetList{ListMeta: obj.(*v1alpha1.AwsCloudformationStackSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsCloudformationStackSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsCloudformationStackSets.
func (c *FakeAwsCloudformationStackSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awscloudformationstacksetsResource, opts))
}

// Create takes the representation of a awsCloudformationStackSet and creates it.  Returns the server's representation of the awsCloudformationStackSet, and an error, if there is any.
func (c *FakeAwsCloudformationStackSets) Create(awsCloudformationStackSet *v1alpha1.AwsCloudformationStackSet) (result *v1alpha1.AwsCloudformationStackSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awscloudformationstacksetsResource, awsCloudformationStackSet), &v1alpha1.AwsCloudformationStackSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCloudformationStackSet), err
}

// Update takes the representation of a awsCloudformationStackSet and updates it. Returns the server's representation of the awsCloudformationStackSet, and an error, if there is any.
func (c *FakeAwsCloudformationStackSets) Update(awsCloudformationStackSet *v1alpha1.AwsCloudformationStackSet) (result *v1alpha1.AwsCloudformationStackSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awscloudformationstacksetsResource, awsCloudformationStackSet), &v1alpha1.AwsCloudformationStackSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCloudformationStackSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsCloudformationStackSets) UpdateStatus(awsCloudformationStackSet *v1alpha1.AwsCloudformationStackSet) (*v1alpha1.AwsCloudformationStackSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awscloudformationstacksetsResource, "status", awsCloudformationStackSet), &v1alpha1.AwsCloudformationStackSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCloudformationStackSet), err
}

// Delete takes name of the awsCloudformationStackSet and deletes it. Returns an error if one occurs.
func (c *FakeAwsCloudformationStackSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awscloudformationstacksetsResource, name), &v1alpha1.AwsCloudformationStackSet{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsCloudformationStackSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awscloudformationstacksetsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsCloudformationStackSetList{})
	return err
}

// Patch applies the patch and returns the patched awsCloudformationStackSet.
func (c *FakeAwsCloudformationStackSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCloudformationStackSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awscloudformationstacksetsResource, name, pt, data, subresources...), &v1alpha1.AwsCloudformationStackSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCloudformationStackSet), err
}