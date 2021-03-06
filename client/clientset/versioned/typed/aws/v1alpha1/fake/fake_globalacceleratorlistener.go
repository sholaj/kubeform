/*
Copyright The Kubeform Authors.

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
	"context"

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGlobalacceleratorListeners implements GlobalacceleratorListenerInterface
type FakeGlobalacceleratorListeners struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var globalacceleratorlistenersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "globalacceleratorlisteners"}

var globalacceleratorlistenersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "GlobalacceleratorListener"}

// Get takes name of the globalacceleratorListener, and returns the corresponding globalacceleratorListener object, and an error if there is any.
func (c *FakeGlobalacceleratorListeners) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GlobalacceleratorListener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(globalacceleratorlistenersResource, c.ns, name), &v1alpha1.GlobalacceleratorListener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlobalacceleratorListener), err
}

// List takes label and field selectors, and returns the list of GlobalacceleratorListeners that match those selectors.
func (c *FakeGlobalacceleratorListeners) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GlobalacceleratorListenerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(globalacceleratorlistenersResource, globalacceleratorlistenersKind, c.ns, opts), &v1alpha1.GlobalacceleratorListenerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GlobalacceleratorListenerList{ListMeta: obj.(*v1alpha1.GlobalacceleratorListenerList).ListMeta}
	for _, item := range obj.(*v1alpha1.GlobalacceleratorListenerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested globalacceleratorListeners.
func (c *FakeGlobalacceleratorListeners) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(globalacceleratorlistenersResource, c.ns, opts))

}

// Create takes the representation of a globalacceleratorListener and creates it.  Returns the server's representation of the globalacceleratorListener, and an error, if there is any.
func (c *FakeGlobalacceleratorListeners) Create(ctx context.Context, globalacceleratorListener *v1alpha1.GlobalacceleratorListener, opts v1.CreateOptions) (result *v1alpha1.GlobalacceleratorListener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(globalacceleratorlistenersResource, c.ns, globalacceleratorListener), &v1alpha1.GlobalacceleratorListener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlobalacceleratorListener), err
}

// Update takes the representation of a globalacceleratorListener and updates it. Returns the server's representation of the globalacceleratorListener, and an error, if there is any.
func (c *FakeGlobalacceleratorListeners) Update(ctx context.Context, globalacceleratorListener *v1alpha1.GlobalacceleratorListener, opts v1.UpdateOptions) (result *v1alpha1.GlobalacceleratorListener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(globalacceleratorlistenersResource, c.ns, globalacceleratorListener), &v1alpha1.GlobalacceleratorListener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlobalacceleratorListener), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGlobalacceleratorListeners) UpdateStatus(ctx context.Context, globalacceleratorListener *v1alpha1.GlobalacceleratorListener, opts v1.UpdateOptions) (*v1alpha1.GlobalacceleratorListener, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(globalacceleratorlistenersResource, "status", c.ns, globalacceleratorListener), &v1alpha1.GlobalacceleratorListener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlobalacceleratorListener), err
}

// Delete takes name of the globalacceleratorListener and deletes it. Returns an error if one occurs.
func (c *FakeGlobalacceleratorListeners) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(globalacceleratorlistenersResource, c.ns, name), &v1alpha1.GlobalacceleratorListener{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGlobalacceleratorListeners) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(globalacceleratorlistenersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.GlobalacceleratorListenerList{})
	return err
}

// Patch applies the patch and returns the patched globalacceleratorListener.
func (c *FakeGlobalacceleratorListeners) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GlobalacceleratorListener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(globalacceleratorlistenersResource, c.ns, name, pt, data, subresources...), &v1alpha1.GlobalacceleratorListener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlobalacceleratorListener), err
}
