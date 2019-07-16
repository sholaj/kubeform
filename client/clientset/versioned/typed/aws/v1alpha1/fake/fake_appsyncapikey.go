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

// FakeAppsyncApiKeys implements AppsyncApiKeyInterface
type FakeAppsyncApiKeys struct {
	Fake *FakeAwsV1alpha1
}

var appsyncapikeysResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "appsyncapikeys"}

var appsyncapikeysKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AppsyncApiKey"}

// Get takes name of the appsyncApiKey, and returns the corresponding appsyncApiKey object, and an error if there is any.
func (c *FakeAppsyncApiKeys) Get(name string, options v1.GetOptions) (result *v1alpha1.AppsyncApiKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(appsyncapikeysResource, name), &v1alpha1.AppsyncApiKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppsyncApiKey), err
}

// List takes label and field selectors, and returns the list of AppsyncApiKeys that match those selectors.
func (c *FakeAppsyncApiKeys) List(opts v1.ListOptions) (result *v1alpha1.AppsyncApiKeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(appsyncapikeysResource, appsyncapikeysKind, opts), &v1alpha1.AppsyncApiKeyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AppsyncApiKeyList{ListMeta: obj.(*v1alpha1.AppsyncApiKeyList).ListMeta}
	for _, item := range obj.(*v1alpha1.AppsyncApiKeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested appsyncApiKeys.
func (c *FakeAppsyncApiKeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(appsyncapikeysResource, opts))
}

// Create takes the representation of a appsyncApiKey and creates it.  Returns the server's representation of the appsyncApiKey, and an error, if there is any.
func (c *FakeAppsyncApiKeys) Create(appsyncApiKey *v1alpha1.AppsyncApiKey) (result *v1alpha1.AppsyncApiKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(appsyncapikeysResource, appsyncApiKey), &v1alpha1.AppsyncApiKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppsyncApiKey), err
}

// Update takes the representation of a appsyncApiKey and updates it. Returns the server's representation of the appsyncApiKey, and an error, if there is any.
func (c *FakeAppsyncApiKeys) Update(appsyncApiKey *v1alpha1.AppsyncApiKey) (result *v1alpha1.AppsyncApiKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(appsyncapikeysResource, appsyncApiKey), &v1alpha1.AppsyncApiKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppsyncApiKey), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAppsyncApiKeys) UpdateStatus(appsyncApiKey *v1alpha1.AppsyncApiKey) (*v1alpha1.AppsyncApiKey, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(appsyncapikeysResource, "status", appsyncApiKey), &v1alpha1.AppsyncApiKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppsyncApiKey), err
}

// Delete takes name of the appsyncApiKey and deletes it. Returns an error if one occurs.
func (c *FakeAppsyncApiKeys) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(appsyncapikeysResource, name), &v1alpha1.AppsyncApiKey{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAppsyncApiKeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(appsyncapikeysResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AppsyncApiKeyList{})
	return err
}

// Patch applies the patch and returns the patched appsyncApiKey.
func (c *FakeAppsyncApiKeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppsyncApiKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(appsyncapikeysResource, name, pt, data, subresources...), &v1alpha1.AppsyncApiKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppsyncApiKey), err
}