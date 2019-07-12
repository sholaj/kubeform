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

// FakeAwsKinesisFirehoseDeliveryStreams implements AwsKinesisFirehoseDeliveryStreamInterface
type FakeAwsKinesisFirehoseDeliveryStreams struct {
	Fake *FakeAwsV1alpha1
}

var awskinesisfirehosedeliverystreamsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awskinesisfirehosedeliverystreams"}

var awskinesisfirehosedeliverystreamsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsKinesisFirehoseDeliveryStream"}

// Get takes name of the awsKinesisFirehoseDeliveryStream, and returns the corresponding awsKinesisFirehoseDeliveryStream object, and an error if there is any.
func (c *FakeAwsKinesisFirehoseDeliveryStreams) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsKinesisFirehoseDeliveryStream, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awskinesisfirehosedeliverystreamsResource, name), &v1alpha1.AwsKinesisFirehoseDeliveryStream{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsKinesisFirehoseDeliveryStream), err
}

// List takes label and field selectors, and returns the list of AwsKinesisFirehoseDeliveryStreams that match those selectors.
func (c *FakeAwsKinesisFirehoseDeliveryStreams) List(opts v1.ListOptions) (result *v1alpha1.AwsKinesisFirehoseDeliveryStreamList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awskinesisfirehosedeliverystreamsResource, awskinesisfirehosedeliverystreamsKind, opts), &v1alpha1.AwsKinesisFirehoseDeliveryStreamList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsKinesisFirehoseDeliveryStreamList{ListMeta: obj.(*v1alpha1.AwsKinesisFirehoseDeliveryStreamList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsKinesisFirehoseDeliveryStreamList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsKinesisFirehoseDeliveryStreams.
func (c *FakeAwsKinesisFirehoseDeliveryStreams) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awskinesisfirehosedeliverystreamsResource, opts))
}

// Create takes the representation of a awsKinesisFirehoseDeliveryStream and creates it.  Returns the server's representation of the awsKinesisFirehoseDeliveryStream, and an error, if there is any.
func (c *FakeAwsKinesisFirehoseDeliveryStreams) Create(awsKinesisFirehoseDeliveryStream *v1alpha1.AwsKinesisFirehoseDeliveryStream) (result *v1alpha1.AwsKinesisFirehoseDeliveryStream, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awskinesisfirehosedeliverystreamsResource, awsKinesisFirehoseDeliveryStream), &v1alpha1.AwsKinesisFirehoseDeliveryStream{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsKinesisFirehoseDeliveryStream), err
}

// Update takes the representation of a awsKinesisFirehoseDeliveryStream and updates it. Returns the server's representation of the awsKinesisFirehoseDeliveryStream, and an error, if there is any.
func (c *FakeAwsKinesisFirehoseDeliveryStreams) Update(awsKinesisFirehoseDeliveryStream *v1alpha1.AwsKinesisFirehoseDeliveryStream) (result *v1alpha1.AwsKinesisFirehoseDeliveryStream, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awskinesisfirehosedeliverystreamsResource, awsKinesisFirehoseDeliveryStream), &v1alpha1.AwsKinesisFirehoseDeliveryStream{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsKinesisFirehoseDeliveryStream), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsKinesisFirehoseDeliveryStreams) UpdateStatus(awsKinesisFirehoseDeliveryStream *v1alpha1.AwsKinesisFirehoseDeliveryStream) (*v1alpha1.AwsKinesisFirehoseDeliveryStream, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awskinesisfirehosedeliverystreamsResource, "status", awsKinesisFirehoseDeliveryStream), &v1alpha1.AwsKinesisFirehoseDeliveryStream{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsKinesisFirehoseDeliveryStream), err
}

// Delete takes name of the awsKinesisFirehoseDeliveryStream and deletes it. Returns an error if one occurs.
func (c *FakeAwsKinesisFirehoseDeliveryStreams) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awskinesisfirehosedeliverystreamsResource, name), &v1alpha1.AwsKinesisFirehoseDeliveryStream{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsKinesisFirehoseDeliveryStreams) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awskinesisfirehosedeliverystreamsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsKinesisFirehoseDeliveryStreamList{})
	return err
}

// Patch applies the patch and returns the patched awsKinesisFirehoseDeliveryStream.
func (c *FakeAwsKinesisFirehoseDeliveryStreams) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsKinesisFirehoseDeliveryStream, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awskinesisfirehosedeliverystreamsResource, name, pt, data, subresources...), &v1alpha1.AwsKinesisFirehoseDeliveryStream{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsKinesisFirehoseDeliveryStream), err
}