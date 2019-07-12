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

// FakeAwsElastictranscoderPresets implements AwsElastictranscoderPresetInterface
type FakeAwsElastictranscoderPresets struct {
	Fake *FakeAwsV1alpha1
}

var awselastictranscoderpresetsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awselastictranscoderpresets"}

var awselastictranscoderpresetsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsElastictranscoderPreset"}

// Get takes name of the awsElastictranscoderPreset, and returns the corresponding awsElastictranscoderPreset object, and an error if there is any.
func (c *FakeAwsElastictranscoderPresets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsElastictranscoderPreset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awselastictranscoderpresetsResource, name), &v1alpha1.AwsElastictranscoderPreset{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElastictranscoderPreset), err
}

// List takes label and field selectors, and returns the list of AwsElastictranscoderPresets that match those selectors.
func (c *FakeAwsElastictranscoderPresets) List(opts v1.ListOptions) (result *v1alpha1.AwsElastictranscoderPresetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awselastictranscoderpresetsResource, awselastictranscoderpresetsKind, opts), &v1alpha1.AwsElastictranscoderPresetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsElastictranscoderPresetList{ListMeta: obj.(*v1alpha1.AwsElastictranscoderPresetList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsElastictranscoderPresetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsElastictranscoderPresets.
func (c *FakeAwsElastictranscoderPresets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awselastictranscoderpresetsResource, opts))
}

// Create takes the representation of a awsElastictranscoderPreset and creates it.  Returns the server's representation of the awsElastictranscoderPreset, and an error, if there is any.
func (c *FakeAwsElastictranscoderPresets) Create(awsElastictranscoderPreset *v1alpha1.AwsElastictranscoderPreset) (result *v1alpha1.AwsElastictranscoderPreset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awselastictranscoderpresetsResource, awsElastictranscoderPreset), &v1alpha1.AwsElastictranscoderPreset{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElastictranscoderPreset), err
}

// Update takes the representation of a awsElastictranscoderPreset and updates it. Returns the server's representation of the awsElastictranscoderPreset, and an error, if there is any.
func (c *FakeAwsElastictranscoderPresets) Update(awsElastictranscoderPreset *v1alpha1.AwsElastictranscoderPreset) (result *v1alpha1.AwsElastictranscoderPreset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awselastictranscoderpresetsResource, awsElastictranscoderPreset), &v1alpha1.AwsElastictranscoderPreset{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElastictranscoderPreset), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsElastictranscoderPresets) UpdateStatus(awsElastictranscoderPreset *v1alpha1.AwsElastictranscoderPreset) (*v1alpha1.AwsElastictranscoderPreset, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awselastictranscoderpresetsResource, "status", awsElastictranscoderPreset), &v1alpha1.AwsElastictranscoderPreset{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElastictranscoderPreset), err
}

// Delete takes name of the awsElastictranscoderPreset and deletes it. Returns an error if one occurs.
func (c *FakeAwsElastictranscoderPresets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awselastictranscoderpresetsResource, name), &v1alpha1.AwsElastictranscoderPreset{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsElastictranscoderPresets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awselastictranscoderpresetsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsElastictranscoderPresetList{})
	return err
}

// Patch applies the patch and returns the patched awsElastictranscoderPreset.
func (c *FakeAwsElastictranscoderPresets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsElastictranscoderPreset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awselastictranscoderpresetsResource, name, pt, data, subresources...), &v1alpha1.AwsElastictranscoderPreset{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElastictranscoderPreset), err
}