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

// FakePinpointSmsChannels implements PinpointSmsChannelInterface
type FakePinpointSmsChannels struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var pinpointsmschannelsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "pinpointsmschannels"}

var pinpointsmschannelsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "PinpointSmsChannel"}

// Get takes name of the pinpointSmsChannel, and returns the corresponding pinpointSmsChannel object, and an error if there is any.
func (c *FakePinpointSmsChannels) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PinpointSmsChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pinpointsmschannelsResource, c.ns, name), &v1alpha1.PinpointSmsChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointSmsChannel), err
}

// List takes label and field selectors, and returns the list of PinpointSmsChannels that match those selectors.
func (c *FakePinpointSmsChannels) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PinpointSmsChannelList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pinpointsmschannelsResource, pinpointsmschannelsKind, c.ns, opts), &v1alpha1.PinpointSmsChannelList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PinpointSmsChannelList{ListMeta: obj.(*v1alpha1.PinpointSmsChannelList).ListMeta}
	for _, item := range obj.(*v1alpha1.PinpointSmsChannelList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pinpointSmsChannels.
func (c *FakePinpointSmsChannels) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pinpointsmschannelsResource, c.ns, opts))

}

// Create takes the representation of a pinpointSmsChannel and creates it.  Returns the server's representation of the pinpointSmsChannel, and an error, if there is any.
func (c *FakePinpointSmsChannels) Create(ctx context.Context, pinpointSmsChannel *v1alpha1.PinpointSmsChannel, opts v1.CreateOptions) (result *v1alpha1.PinpointSmsChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pinpointsmschannelsResource, c.ns, pinpointSmsChannel), &v1alpha1.PinpointSmsChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointSmsChannel), err
}

// Update takes the representation of a pinpointSmsChannel and updates it. Returns the server's representation of the pinpointSmsChannel, and an error, if there is any.
func (c *FakePinpointSmsChannels) Update(ctx context.Context, pinpointSmsChannel *v1alpha1.PinpointSmsChannel, opts v1.UpdateOptions) (result *v1alpha1.PinpointSmsChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pinpointsmschannelsResource, c.ns, pinpointSmsChannel), &v1alpha1.PinpointSmsChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointSmsChannel), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePinpointSmsChannels) UpdateStatus(ctx context.Context, pinpointSmsChannel *v1alpha1.PinpointSmsChannel, opts v1.UpdateOptions) (*v1alpha1.PinpointSmsChannel, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(pinpointsmschannelsResource, "status", c.ns, pinpointSmsChannel), &v1alpha1.PinpointSmsChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointSmsChannel), err
}

// Delete takes name of the pinpointSmsChannel and deletes it. Returns an error if one occurs.
func (c *FakePinpointSmsChannels) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pinpointsmschannelsResource, c.ns, name), &v1alpha1.PinpointSmsChannel{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePinpointSmsChannels) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pinpointsmschannelsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PinpointSmsChannelList{})
	return err
}

// Patch applies the patch and returns the patched pinpointSmsChannel.
func (c *FakePinpointSmsChannels) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PinpointSmsChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pinpointsmschannelsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PinpointSmsChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointSmsChannel), err
}
