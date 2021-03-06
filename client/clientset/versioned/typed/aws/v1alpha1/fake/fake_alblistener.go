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

// FakeAlbListeners implements AlbListenerInterface
type FakeAlbListeners struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var alblistenersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "alblisteners"}

var alblistenersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AlbListener"}

// Get takes name of the albListener, and returns the corresponding albListener object, and an error if there is any.
func (c *FakeAlbListeners) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AlbListener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(alblistenersResource, c.ns, name), &v1alpha1.AlbListener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlbListener), err
}

// List takes label and field selectors, and returns the list of AlbListeners that match those selectors.
func (c *FakeAlbListeners) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AlbListenerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(alblistenersResource, alblistenersKind, c.ns, opts), &v1alpha1.AlbListenerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AlbListenerList{ListMeta: obj.(*v1alpha1.AlbListenerList).ListMeta}
	for _, item := range obj.(*v1alpha1.AlbListenerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested albListeners.
func (c *FakeAlbListeners) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(alblistenersResource, c.ns, opts))

}

// Create takes the representation of a albListener and creates it.  Returns the server's representation of the albListener, and an error, if there is any.
func (c *FakeAlbListeners) Create(ctx context.Context, albListener *v1alpha1.AlbListener, opts v1.CreateOptions) (result *v1alpha1.AlbListener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(alblistenersResource, c.ns, albListener), &v1alpha1.AlbListener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlbListener), err
}

// Update takes the representation of a albListener and updates it. Returns the server's representation of the albListener, and an error, if there is any.
func (c *FakeAlbListeners) Update(ctx context.Context, albListener *v1alpha1.AlbListener, opts v1.UpdateOptions) (result *v1alpha1.AlbListener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(alblistenersResource, c.ns, albListener), &v1alpha1.AlbListener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlbListener), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAlbListeners) UpdateStatus(ctx context.Context, albListener *v1alpha1.AlbListener, opts v1.UpdateOptions) (*v1alpha1.AlbListener, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(alblistenersResource, "status", c.ns, albListener), &v1alpha1.AlbListener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlbListener), err
}

// Delete takes name of the albListener and deletes it. Returns an error if one occurs.
func (c *FakeAlbListeners) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(alblistenersResource, c.ns, name), &v1alpha1.AlbListener{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAlbListeners) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(alblistenersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AlbListenerList{})
	return err
}

// Patch applies the patch and returns the patched albListener.
func (c *FakeAlbListeners) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AlbListener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(alblistenersResource, c.ns, name, pt, data, subresources...), &v1alpha1.AlbListener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlbListener), err
}
