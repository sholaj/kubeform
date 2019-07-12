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

// FakeAwsStoragegatewayCaches implements AwsStoragegatewayCacheInterface
type FakeAwsStoragegatewayCaches struct {
	Fake *FakeAwsV1alpha1
}

var awsstoragegatewaycachesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsstoragegatewaycaches"}

var awsstoragegatewaycachesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsStoragegatewayCache"}

// Get takes name of the awsStoragegatewayCache, and returns the corresponding awsStoragegatewayCache object, and an error if there is any.
func (c *FakeAwsStoragegatewayCaches) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsStoragegatewayCache, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsstoragegatewaycachesResource, name), &v1alpha1.AwsStoragegatewayCache{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsStoragegatewayCache), err
}

// List takes label and field selectors, and returns the list of AwsStoragegatewayCaches that match those selectors.
func (c *FakeAwsStoragegatewayCaches) List(opts v1.ListOptions) (result *v1alpha1.AwsStoragegatewayCacheList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsstoragegatewaycachesResource, awsstoragegatewaycachesKind, opts), &v1alpha1.AwsStoragegatewayCacheList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsStoragegatewayCacheList{ListMeta: obj.(*v1alpha1.AwsStoragegatewayCacheList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsStoragegatewayCacheList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsStoragegatewayCaches.
func (c *FakeAwsStoragegatewayCaches) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsstoragegatewaycachesResource, opts))
}

// Create takes the representation of a awsStoragegatewayCache and creates it.  Returns the server's representation of the awsStoragegatewayCache, and an error, if there is any.
func (c *FakeAwsStoragegatewayCaches) Create(awsStoragegatewayCache *v1alpha1.AwsStoragegatewayCache) (result *v1alpha1.AwsStoragegatewayCache, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsstoragegatewaycachesResource, awsStoragegatewayCache), &v1alpha1.AwsStoragegatewayCache{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsStoragegatewayCache), err
}

// Update takes the representation of a awsStoragegatewayCache and updates it. Returns the server's representation of the awsStoragegatewayCache, and an error, if there is any.
func (c *FakeAwsStoragegatewayCaches) Update(awsStoragegatewayCache *v1alpha1.AwsStoragegatewayCache) (result *v1alpha1.AwsStoragegatewayCache, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsstoragegatewaycachesResource, awsStoragegatewayCache), &v1alpha1.AwsStoragegatewayCache{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsStoragegatewayCache), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsStoragegatewayCaches) UpdateStatus(awsStoragegatewayCache *v1alpha1.AwsStoragegatewayCache) (*v1alpha1.AwsStoragegatewayCache, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsstoragegatewaycachesResource, "status", awsStoragegatewayCache), &v1alpha1.AwsStoragegatewayCache{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsStoragegatewayCache), err
}

// Delete takes name of the awsStoragegatewayCache and deletes it. Returns an error if one occurs.
func (c *FakeAwsStoragegatewayCaches) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsstoragegatewaycachesResource, name), &v1alpha1.AwsStoragegatewayCache{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsStoragegatewayCaches) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsstoragegatewaycachesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsStoragegatewayCacheList{})
	return err
}

// Patch applies the patch and returns the patched awsStoragegatewayCache.
func (c *FakeAwsStoragegatewayCaches) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsStoragegatewayCache, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsstoragegatewaycachesResource, name, pt, data, subresources...), &v1alpha1.AwsStoragegatewayCache{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsStoragegatewayCache), err
}