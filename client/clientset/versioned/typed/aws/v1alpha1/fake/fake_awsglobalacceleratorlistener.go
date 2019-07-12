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

// FakeAwsGlobalacceleratorListeners implements AwsGlobalacceleratorListenerInterface
type FakeAwsGlobalacceleratorListeners struct {
	Fake *FakeAwsV1alpha1
}

var awsglobalacceleratorlistenersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsglobalacceleratorlisteners"}

var awsglobalacceleratorlistenersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsGlobalacceleratorListener"}

// Get takes name of the awsGlobalacceleratorListener, and returns the corresponding awsGlobalacceleratorListener object, and an error if there is any.
func (c *FakeAwsGlobalacceleratorListeners) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsGlobalacceleratorListener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsglobalacceleratorlistenersResource, name), &v1alpha1.AwsGlobalacceleratorListener{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGlobalacceleratorListener), err
}

// List takes label and field selectors, and returns the list of AwsGlobalacceleratorListeners that match those selectors.
func (c *FakeAwsGlobalacceleratorListeners) List(opts v1.ListOptions) (result *v1alpha1.AwsGlobalacceleratorListenerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsglobalacceleratorlistenersResource, awsglobalacceleratorlistenersKind, opts), &v1alpha1.AwsGlobalacceleratorListenerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsGlobalacceleratorListenerList{ListMeta: obj.(*v1alpha1.AwsGlobalacceleratorListenerList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsGlobalacceleratorListenerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsGlobalacceleratorListeners.
func (c *FakeAwsGlobalacceleratorListeners) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsglobalacceleratorlistenersResource, opts))
}

// Create takes the representation of a awsGlobalacceleratorListener and creates it.  Returns the server's representation of the awsGlobalacceleratorListener, and an error, if there is any.
func (c *FakeAwsGlobalacceleratorListeners) Create(awsGlobalacceleratorListener *v1alpha1.AwsGlobalacceleratorListener) (result *v1alpha1.AwsGlobalacceleratorListener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsglobalacceleratorlistenersResource, awsGlobalacceleratorListener), &v1alpha1.AwsGlobalacceleratorListener{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGlobalacceleratorListener), err
}

// Update takes the representation of a awsGlobalacceleratorListener and updates it. Returns the server's representation of the awsGlobalacceleratorListener, and an error, if there is any.
func (c *FakeAwsGlobalacceleratorListeners) Update(awsGlobalacceleratorListener *v1alpha1.AwsGlobalacceleratorListener) (result *v1alpha1.AwsGlobalacceleratorListener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsglobalacceleratorlistenersResource, awsGlobalacceleratorListener), &v1alpha1.AwsGlobalacceleratorListener{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGlobalacceleratorListener), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsGlobalacceleratorListeners) UpdateStatus(awsGlobalacceleratorListener *v1alpha1.AwsGlobalacceleratorListener) (*v1alpha1.AwsGlobalacceleratorListener, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsglobalacceleratorlistenersResource, "status", awsGlobalacceleratorListener), &v1alpha1.AwsGlobalacceleratorListener{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGlobalacceleratorListener), err
}

// Delete takes name of the awsGlobalacceleratorListener and deletes it. Returns an error if one occurs.
func (c *FakeAwsGlobalacceleratorListeners) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsglobalacceleratorlistenersResource, name), &v1alpha1.AwsGlobalacceleratorListener{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsGlobalacceleratorListeners) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsglobalacceleratorlistenersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsGlobalacceleratorListenerList{})
	return err
}

// Patch applies the patch and returns the patched awsGlobalacceleratorListener.
func (c *FakeAwsGlobalacceleratorListeners) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsGlobalacceleratorListener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsglobalacceleratorlistenersResource, name, pt, data, subresources...), &v1alpha1.AwsGlobalacceleratorListener{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGlobalacceleratorListener), err
}