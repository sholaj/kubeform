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

// FakeAwsRamResourceAssociations implements AwsRamResourceAssociationInterface
type FakeAwsRamResourceAssociations struct {
	Fake *FakeAwsV1alpha1
}

var awsramresourceassociationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsramresourceassociations"}

var awsramresourceassociationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsRamResourceAssociation"}

// Get takes name of the awsRamResourceAssociation, and returns the corresponding awsRamResourceAssociation object, and an error if there is any.
func (c *FakeAwsRamResourceAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsRamResourceAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsramresourceassociationsResource, name), &v1alpha1.AwsRamResourceAssociation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRamResourceAssociation), err
}

// List takes label and field selectors, and returns the list of AwsRamResourceAssociations that match those selectors.
func (c *FakeAwsRamResourceAssociations) List(opts v1.ListOptions) (result *v1alpha1.AwsRamResourceAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsramresourceassociationsResource, awsramresourceassociationsKind, opts), &v1alpha1.AwsRamResourceAssociationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsRamResourceAssociationList{ListMeta: obj.(*v1alpha1.AwsRamResourceAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsRamResourceAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsRamResourceAssociations.
func (c *FakeAwsRamResourceAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsramresourceassociationsResource, opts))
}

// Create takes the representation of a awsRamResourceAssociation and creates it.  Returns the server's representation of the awsRamResourceAssociation, and an error, if there is any.
func (c *FakeAwsRamResourceAssociations) Create(awsRamResourceAssociation *v1alpha1.AwsRamResourceAssociation) (result *v1alpha1.AwsRamResourceAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsramresourceassociationsResource, awsRamResourceAssociation), &v1alpha1.AwsRamResourceAssociation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRamResourceAssociation), err
}

// Update takes the representation of a awsRamResourceAssociation and updates it. Returns the server's representation of the awsRamResourceAssociation, and an error, if there is any.
func (c *FakeAwsRamResourceAssociations) Update(awsRamResourceAssociation *v1alpha1.AwsRamResourceAssociation) (result *v1alpha1.AwsRamResourceAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsramresourceassociationsResource, awsRamResourceAssociation), &v1alpha1.AwsRamResourceAssociation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRamResourceAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsRamResourceAssociations) UpdateStatus(awsRamResourceAssociation *v1alpha1.AwsRamResourceAssociation) (*v1alpha1.AwsRamResourceAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsramresourceassociationsResource, "status", awsRamResourceAssociation), &v1alpha1.AwsRamResourceAssociation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRamResourceAssociation), err
}

// Delete takes name of the awsRamResourceAssociation and deletes it. Returns an error if one occurs.
func (c *FakeAwsRamResourceAssociations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsramresourceassociationsResource, name), &v1alpha1.AwsRamResourceAssociation{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsRamResourceAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsramresourceassociationsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsRamResourceAssociationList{})
	return err
}

// Patch applies the patch and returns the patched awsRamResourceAssociation.
func (c *FakeAwsRamResourceAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsRamResourceAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsramresourceassociationsResource, name, pt, data, subresources...), &v1alpha1.AwsRamResourceAssociation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRamResourceAssociation), err
}