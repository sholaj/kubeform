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

// FakeAwsIamGroupPolicyAttachments implements AwsIamGroupPolicyAttachmentInterface
type FakeAwsIamGroupPolicyAttachments struct {
	Fake *FakeAwsV1alpha1
}

var awsiamgrouppolicyattachmentsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsiamgrouppolicyattachments"}

var awsiamgrouppolicyattachmentsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsIamGroupPolicyAttachment"}

// Get takes name of the awsIamGroupPolicyAttachment, and returns the corresponding awsIamGroupPolicyAttachment object, and an error if there is any.
func (c *FakeAwsIamGroupPolicyAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsIamGroupPolicyAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsiamgrouppolicyattachmentsResource, name), &v1alpha1.AwsIamGroupPolicyAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamGroupPolicyAttachment), err
}

// List takes label and field selectors, and returns the list of AwsIamGroupPolicyAttachments that match those selectors.
func (c *FakeAwsIamGroupPolicyAttachments) List(opts v1.ListOptions) (result *v1alpha1.AwsIamGroupPolicyAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsiamgrouppolicyattachmentsResource, awsiamgrouppolicyattachmentsKind, opts), &v1alpha1.AwsIamGroupPolicyAttachmentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsIamGroupPolicyAttachmentList{ListMeta: obj.(*v1alpha1.AwsIamGroupPolicyAttachmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsIamGroupPolicyAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsIamGroupPolicyAttachments.
func (c *FakeAwsIamGroupPolicyAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsiamgrouppolicyattachmentsResource, opts))
}

// Create takes the representation of a awsIamGroupPolicyAttachment and creates it.  Returns the server's representation of the awsIamGroupPolicyAttachment, and an error, if there is any.
func (c *FakeAwsIamGroupPolicyAttachments) Create(awsIamGroupPolicyAttachment *v1alpha1.AwsIamGroupPolicyAttachment) (result *v1alpha1.AwsIamGroupPolicyAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsiamgrouppolicyattachmentsResource, awsIamGroupPolicyAttachment), &v1alpha1.AwsIamGroupPolicyAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamGroupPolicyAttachment), err
}

// Update takes the representation of a awsIamGroupPolicyAttachment and updates it. Returns the server's representation of the awsIamGroupPolicyAttachment, and an error, if there is any.
func (c *FakeAwsIamGroupPolicyAttachments) Update(awsIamGroupPolicyAttachment *v1alpha1.AwsIamGroupPolicyAttachment) (result *v1alpha1.AwsIamGroupPolicyAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsiamgrouppolicyattachmentsResource, awsIamGroupPolicyAttachment), &v1alpha1.AwsIamGroupPolicyAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamGroupPolicyAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsIamGroupPolicyAttachments) UpdateStatus(awsIamGroupPolicyAttachment *v1alpha1.AwsIamGroupPolicyAttachment) (*v1alpha1.AwsIamGroupPolicyAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsiamgrouppolicyattachmentsResource, "status", awsIamGroupPolicyAttachment), &v1alpha1.AwsIamGroupPolicyAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamGroupPolicyAttachment), err
}

// Delete takes name of the awsIamGroupPolicyAttachment and deletes it. Returns an error if one occurs.
func (c *FakeAwsIamGroupPolicyAttachments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsiamgrouppolicyattachmentsResource, name), &v1alpha1.AwsIamGroupPolicyAttachment{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsIamGroupPolicyAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsiamgrouppolicyattachmentsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsIamGroupPolicyAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched awsIamGroupPolicyAttachment.
func (c *FakeAwsIamGroupPolicyAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsIamGroupPolicyAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsiamgrouppolicyattachmentsResource, name, pt, data, subresources...), &v1alpha1.AwsIamGroupPolicyAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamGroupPolicyAttachment), err
}