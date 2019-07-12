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

// FakeAwsAlbs implements AwsAlbInterface
type FakeAwsAlbs struct {
	Fake *FakeAwsV1alpha1
}

var awsalbsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsalbs"}

var awsalbsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsAlb"}

// Get takes name of the awsAlb, and returns the corresponding awsAlb object, and an error if there is any.
func (c *FakeAwsAlbs) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsAlb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsalbsResource, name), &v1alpha1.AwsAlb{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAlb), err
}

// List takes label and field selectors, and returns the list of AwsAlbs that match those selectors.
func (c *FakeAwsAlbs) List(opts v1.ListOptions) (result *v1alpha1.AwsAlbList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsalbsResource, awsalbsKind, opts), &v1alpha1.AwsAlbList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsAlbList{ListMeta: obj.(*v1alpha1.AwsAlbList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsAlbList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsAlbs.
func (c *FakeAwsAlbs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsalbsResource, opts))
}

// Create takes the representation of a awsAlb and creates it.  Returns the server's representation of the awsAlb, and an error, if there is any.
func (c *FakeAwsAlbs) Create(awsAlb *v1alpha1.AwsAlb) (result *v1alpha1.AwsAlb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsalbsResource, awsAlb), &v1alpha1.AwsAlb{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAlb), err
}

// Update takes the representation of a awsAlb and updates it. Returns the server's representation of the awsAlb, and an error, if there is any.
func (c *FakeAwsAlbs) Update(awsAlb *v1alpha1.AwsAlb) (result *v1alpha1.AwsAlb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsalbsResource, awsAlb), &v1alpha1.AwsAlb{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAlb), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsAlbs) UpdateStatus(awsAlb *v1alpha1.AwsAlb) (*v1alpha1.AwsAlb, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsalbsResource, "status", awsAlb), &v1alpha1.AwsAlb{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAlb), err
}

// Delete takes name of the awsAlb and deletes it. Returns an error if one occurs.
func (c *FakeAwsAlbs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsalbsResource, name), &v1alpha1.AwsAlb{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsAlbs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsalbsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsAlbList{})
	return err
}

// Patch applies the patch and returns the patched awsAlb.
func (c *FakeAwsAlbs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAlb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsalbsResource, name, pt, data, subresources...), &v1alpha1.AwsAlb{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAlb), err
}