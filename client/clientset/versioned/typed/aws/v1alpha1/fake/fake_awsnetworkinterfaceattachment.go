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

// FakeAwsNetworkInterfaceAttachments implements AwsNetworkInterfaceAttachmentInterface
type FakeAwsNetworkInterfaceAttachments struct {
	Fake *FakeAwsV1alpha1
}

var awsnetworkinterfaceattachmentsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsnetworkinterfaceattachments"}

var awsnetworkinterfaceattachmentsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsNetworkInterfaceAttachment"}

// Get takes name of the awsNetworkInterfaceAttachment, and returns the corresponding awsNetworkInterfaceAttachment object, and an error if there is any.
func (c *FakeAwsNetworkInterfaceAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsNetworkInterfaceAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsnetworkinterfaceattachmentsResource, name), &v1alpha1.AwsNetworkInterfaceAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsNetworkInterfaceAttachment), err
}

// List takes label and field selectors, and returns the list of AwsNetworkInterfaceAttachments that match those selectors.
func (c *FakeAwsNetworkInterfaceAttachments) List(opts v1.ListOptions) (result *v1alpha1.AwsNetworkInterfaceAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsnetworkinterfaceattachmentsResource, awsnetworkinterfaceattachmentsKind, opts), &v1alpha1.AwsNetworkInterfaceAttachmentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsNetworkInterfaceAttachmentList{ListMeta: obj.(*v1alpha1.AwsNetworkInterfaceAttachmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsNetworkInterfaceAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsNetworkInterfaceAttachments.
func (c *FakeAwsNetworkInterfaceAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsnetworkinterfaceattachmentsResource, opts))
}

// Create takes the representation of a awsNetworkInterfaceAttachment and creates it.  Returns the server's representation of the awsNetworkInterfaceAttachment, and an error, if there is any.
func (c *FakeAwsNetworkInterfaceAttachments) Create(awsNetworkInterfaceAttachment *v1alpha1.AwsNetworkInterfaceAttachment) (result *v1alpha1.AwsNetworkInterfaceAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsnetworkinterfaceattachmentsResource, awsNetworkInterfaceAttachment), &v1alpha1.AwsNetworkInterfaceAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsNetworkInterfaceAttachment), err
}

// Update takes the representation of a awsNetworkInterfaceAttachment and updates it. Returns the server's representation of the awsNetworkInterfaceAttachment, and an error, if there is any.
func (c *FakeAwsNetworkInterfaceAttachments) Update(awsNetworkInterfaceAttachment *v1alpha1.AwsNetworkInterfaceAttachment) (result *v1alpha1.AwsNetworkInterfaceAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsnetworkinterfaceattachmentsResource, awsNetworkInterfaceAttachment), &v1alpha1.AwsNetworkInterfaceAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsNetworkInterfaceAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsNetworkInterfaceAttachments) UpdateStatus(awsNetworkInterfaceAttachment *v1alpha1.AwsNetworkInterfaceAttachment) (*v1alpha1.AwsNetworkInterfaceAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsnetworkinterfaceattachmentsResource, "status", awsNetworkInterfaceAttachment), &v1alpha1.AwsNetworkInterfaceAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsNetworkInterfaceAttachment), err
}

// Delete takes name of the awsNetworkInterfaceAttachment and deletes it. Returns an error if one occurs.
func (c *FakeAwsNetworkInterfaceAttachments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsnetworkinterfaceattachmentsResource, name), &v1alpha1.AwsNetworkInterfaceAttachment{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsNetworkInterfaceAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsnetworkinterfaceattachmentsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsNetworkInterfaceAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched awsNetworkInterfaceAttachment.
func (c *FakeAwsNetworkInterfaceAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsNetworkInterfaceAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsnetworkinterfaceattachmentsResource, name, pt, data, subresources...), &v1alpha1.AwsNetworkInterfaceAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsNetworkInterfaceAttachment), err
}