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

// FakeAwsIamInstanceProfiles implements AwsIamInstanceProfileInterface
type FakeAwsIamInstanceProfiles struct {
	Fake *FakeAwsV1alpha1
}

var awsiaminstanceprofilesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsiaminstanceprofiles"}

var awsiaminstanceprofilesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsIamInstanceProfile"}

// Get takes name of the awsIamInstanceProfile, and returns the corresponding awsIamInstanceProfile object, and an error if there is any.
func (c *FakeAwsIamInstanceProfiles) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsIamInstanceProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsiaminstanceprofilesResource, name), &v1alpha1.AwsIamInstanceProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamInstanceProfile), err
}

// List takes label and field selectors, and returns the list of AwsIamInstanceProfiles that match those selectors.
func (c *FakeAwsIamInstanceProfiles) List(opts v1.ListOptions) (result *v1alpha1.AwsIamInstanceProfileList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsiaminstanceprofilesResource, awsiaminstanceprofilesKind, opts), &v1alpha1.AwsIamInstanceProfileList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsIamInstanceProfileList{ListMeta: obj.(*v1alpha1.AwsIamInstanceProfileList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsIamInstanceProfileList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsIamInstanceProfiles.
func (c *FakeAwsIamInstanceProfiles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsiaminstanceprofilesResource, opts))
}

// Create takes the representation of a awsIamInstanceProfile and creates it.  Returns the server's representation of the awsIamInstanceProfile, and an error, if there is any.
func (c *FakeAwsIamInstanceProfiles) Create(awsIamInstanceProfile *v1alpha1.AwsIamInstanceProfile) (result *v1alpha1.AwsIamInstanceProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsiaminstanceprofilesResource, awsIamInstanceProfile), &v1alpha1.AwsIamInstanceProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamInstanceProfile), err
}

// Update takes the representation of a awsIamInstanceProfile and updates it. Returns the server's representation of the awsIamInstanceProfile, and an error, if there is any.
func (c *FakeAwsIamInstanceProfiles) Update(awsIamInstanceProfile *v1alpha1.AwsIamInstanceProfile) (result *v1alpha1.AwsIamInstanceProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsiaminstanceprofilesResource, awsIamInstanceProfile), &v1alpha1.AwsIamInstanceProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamInstanceProfile), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsIamInstanceProfiles) UpdateStatus(awsIamInstanceProfile *v1alpha1.AwsIamInstanceProfile) (*v1alpha1.AwsIamInstanceProfile, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsiaminstanceprofilesResource, "status", awsIamInstanceProfile), &v1alpha1.AwsIamInstanceProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamInstanceProfile), err
}

// Delete takes name of the awsIamInstanceProfile and deletes it. Returns an error if one occurs.
func (c *FakeAwsIamInstanceProfiles) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsiaminstanceprofilesResource, name), &v1alpha1.AwsIamInstanceProfile{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsIamInstanceProfiles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsiaminstanceprofilesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsIamInstanceProfileList{})
	return err
}

// Patch applies the patch and returns the patched awsIamInstanceProfile.
func (c *FakeAwsIamInstanceProfiles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsIamInstanceProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsiaminstanceprofilesResource, name, pt, data, subresources...), &v1alpha1.AwsIamInstanceProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamInstanceProfile), err
}