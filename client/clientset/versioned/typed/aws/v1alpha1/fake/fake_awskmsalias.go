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

// FakeAwsKmsAliases implements AwsKmsAliasInterface
type FakeAwsKmsAliases struct {
	Fake *FakeAwsV1alpha1
}

var awskmsaliasesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awskmsaliases"}

var awskmsaliasesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsKmsAlias"}

// Get takes name of the awsKmsAlias, and returns the corresponding awsKmsAlias object, and an error if there is any.
func (c *FakeAwsKmsAliases) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsKmsAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awskmsaliasesResource, name), &v1alpha1.AwsKmsAlias{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsKmsAlias), err
}

// List takes label and field selectors, and returns the list of AwsKmsAliases that match those selectors.
func (c *FakeAwsKmsAliases) List(opts v1.ListOptions) (result *v1alpha1.AwsKmsAliasList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awskmsaliasesResource, awskmsaliasesKind, opts), &v1alpha1.AwsKmsAliasList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsKmsAliasList{ListMeta: obj.(*v1alpha1.AwsKmsAliasList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsKmsAliasList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsKmsAliases.
func (c *FakeAwsKmsAliases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awskmsaliasesResource, opts))
}

// Create takes the representation of a awsKmsAlias and creates it.  Returns the server's representation of the awsKmsAlias, and an error, if there is any.
func (c *FakeAwsKmsAliases) Create(awsKmsAlias *v1alpha1.AwsKmsAlias) (result *v1alpha1.AwsKmsAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awskmsaliasesResource, awsKmsAlias), &v1alpha1.AwsKmsAlias{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsKmsAlias), err
}

// Update takes the representation of a awsKmsAlias and updates it. Returns the server's representation of the awsKmsAlias, and an error, if there is any.
func (c *FakeAwsKmsAliases) Update(awsKmsAlias *v1alpha1.AwsKmsAlias) (result *v1alpha1.AwsKmsAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awskmsaliasesResource, awsKmsAlias), &v1alpha1.AwsKmsAlias{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsKmsAlias), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsKmsAliases) UpdateStatus(awsKmsAlias *v1alpha1.AwsKmsAlias) (*v1alpha1.AwsKmsAlias, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awskmsaliasesResource, "status", awsKmsAlias), &v1alpha1.AwsKmsAlias{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsKmsAlias), err
}

// Delete takes name of the awsKmsAlias and deletes it. Returns an error if one occurs.
func (c *FakeAwsKmsAliases) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awskmsaliasesResource, name), &v1alpha1.AwsKmsAlias{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsKmsAliases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awskmsaliasesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsKmsAliasList{})
	return err
}

// Patch applies the patch and returns the patched awsKmsAlias.
func (c *FakeAwsKmsAliases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsKmsAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awskmsaliasesResource, name, pt, data, subresources...), &v1alpha1.AwsKmsAlias{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsKmsAlias), err
}