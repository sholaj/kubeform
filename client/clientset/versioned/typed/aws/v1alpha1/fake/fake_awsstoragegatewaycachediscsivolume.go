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

// FakeAwsStoragegatewayCachedIscsiVolumes implements AwsStoragegatewayCachedIscsiVolumeInterface
type FakeAwsStoragegatewayCachedIscsiVolumes struct {
	Fake *FakeAwsV1alpha1
}

var awsstoragegatewaycachediscsivolumesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsstoragegatewaycachediscsivolumes"}

var awsstoragegatewaycachediscsivolumesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsStoragegatewayCachedIscsiVolume"}

// Get takes name of the awsStoragegatewayCachedIscsiVolume, and returns the corresponding awsStoragegatewayCachedIscsiVolume object, and an error if there is any.
func (c *FakeAwsStoragegatewayCachedIscsiVolumes) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsStoragegatewayCachedIscsiVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsstoragegatewaycachediscsivolumesResource, name), &v1alpha1.AwsStoragegatewayCachedIscsiVolume{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsStoragegatewayCachedIscsiVolume), err
}

// List takes label and field selectors, and returns the list of AwsStoragegatewayCachedIscsiVolumes that match those selectors.
func (c *FakeAwsStoragegatewayCachedIscsiVolumes) List(opts v1.ListOptions) (result *v1alpha1.AwsStoragegatewayCachedIscsiVolumeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsstoragegatewaycachediscsivolumesResource, awsstoragegatewaycachediscsivolumesKind, opts), &v1alpha1.AwsStoragegatewayCachedIscsiVolumeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsStoragegatewayCachedIscsiVolumeList{ListMeta: obj.(*v1alpha1.AwsStoragegatewayCachedIscsiVolumeList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsStoragegatewayCachedIscsiVolumeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsStoragegatewayCachedIscsiVolumes.
func (c *FakeAwsStoragegatewayCachedIscsiVolumes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsstoragegatewaycachediscsivolumesResource, opts))
}

// Create takes the representation of a awsStoragegatewayCachedIscsiVolume and creates it.  Returns the server's representation of the awsStoragegatewayCachedIscsiVolume, and an error, if there is any.
func (c *FakeAwsStoragegatewayCachedIscsiVolumes) Create(awsStoragegatewayCachedIscsiVolume *v1alpha1.AwsStoragegatewayCachedIscsiVolume) (result *v1alpha1.AwsStoragegatewayCachedIscsiVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsstoragegatewaycachediscsivolumesResource, awsStoragegatewayCachedIscsiVolume), &v1alpha1.AwsStoragegatewayCachedIscsiVolume{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsStoragegatewayCachedIscsiVolume), err
}

// Update takes the representation of a awsStoragegatewayCachedIscsiVolume and updates it. Returns the server's representation of the awsStoragegatewayCachedIscsiVolume, and an error, if there is any.
func (c *FakeAwsStoragegatewayCachedIscsiVolumes) Update(awsStoragegatewayCachedIscsiVolume *v1alpha1.AwsStoragegatewayCachedIscsiVolume) (result *v1alpha1.AwsStoragegatewayCachedIscsiVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsstoragegatewaycachediscsivolumesResource, awsStoragegatewayCachedIscsiVolume), &v1alpha1.AwsStoragegatewayCachedIscsiVolume{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsStoragegatewayCachedIscsiVolume), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsStoragegatewayCachedIscsiVolumes) UpdateStatus(awsStoragegatewayCachedIscsiVolume *v1alpha1.AwsStoragegatewayCachedIscsiVolume) (*v1alpha1.AwsStoragegatewayCachedIscsiVolume, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsstoragegatewaycachediscsivolumesResource, "status", awsStoragegatewayCachedIscsiVolume), &v1alpha1.AwsStoragegatewayCachedIscsiVolume{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsStoragegatewayCachedIscsiVolume), err
}

// Delete takes name of the awsStoragegatewayCachedIscsiVolume and deletes it. Returns an error if one occurs.
func (c *FakeAwsStoragegatewayCachedIscsiVolumes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsstoragegatewaycachediscsivolumesResource, name), &v1alpha1.AwsStoragegatewayCachedIscsiVolume{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsStoragegatewayCachedIscsiVolumes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsstoragegatewaycachediscsivolumesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsStoragegatewayCachedIscsiVolumeList{})
	return err
}

// Patch applies the patch and returns the patched awsStoragegatewayCachedIscsiVolume.
func (c *FakeAwsStoragegatewayCachedIscsiVolumes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsStoragegatewayCachedIscsiVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsstoragegatewaycachediscsivolumesResource, name, pt, data, subresources...), &v1alpha1.AwsStoragegatewayCachedIscsiVolume{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsStoragegatewayCachedIscsiVolume), err
}