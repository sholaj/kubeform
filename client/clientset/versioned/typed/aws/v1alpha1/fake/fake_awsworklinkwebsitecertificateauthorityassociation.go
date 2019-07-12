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

// FakeAwsWorklinkWebsiteCertificateAuthorityAssociations implements AwsWorklinkWebsiteCertificateAuthorityAssociationInterface
type FakeAwsWorklinkWebsiteCertificateAuthorityAssociations struct {
	Fake *FakeAwsV1alpha1
}

var awsworklinkwebsitecertificateauthorityassociationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsworklinkwebsitecertificateauthorityassociations"}

var awsworklinkwebsitecertificateauthorityassociationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsWorklinkWebsiteCertificateAuthorityAssociation"}

// Get takes name of the awsWorklinkWebsiteCertificateAuthorityAssociation, and returns the corresponding awsWorklinkWebsiteCertificateAuthorityAssociation object, and an error if there is any.
func (c *FakeAwsWorklinkWebsiteCertificateAuthorityAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsworklinkwebsitecertificateauthorityassociationsResource, name), &v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation), err
}

// List takes label and field selectors, and returns the list of AwsWorklinkWebsiteCertificateAuthorityAssociations that match those selectors.
func (c *FakeAwsWorklinkWebsiteCertificateAuthorityAssociations) List(opts v1.ListOptions) (result *v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsworklinkwebsitecertificateauthorityassociationsResource, awsworklinkwebsitecertificateauthorityassociationsKind, opts), &v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociationList{ListMeta: obj.(*v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsWorklinkWebsiteCertificateAuthorityAssociations.
func (c *FakeAwsWorklinkWebsiteCertificateAuthorityAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsworklinkwebsitecertificateauthorityassociationsResource, opts))
}

// Create takes the representation of a awsWorklinkWebsiteCertificateAuthorityAssociation and creates it.  Returns the server's representation of the awsWorklinkWebsiteCertificateAuthorityAssociation, and an error, if there is any.
func (c *FakeAwsWorklinkWebsiteCertificateAuthorityAssociations) Create(awsWorklinkWebsiteCertificateAuthorityAssociation *v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation) (result *v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsworklinkwebsitecertificateauthorityassociationsResource, awsWorklinkWebsiteCertificateAuthorityAssociation), &v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation), err
}

// Update takes the representation of a awsWorklinkWebsiteCertificateAuthorityAssociation and updates it. Returns the server's representation of the awsWorklinkWebsiteCertificateAuthorityAssociation, and an error, if there is any.
func (c *FakeAwsWorklinkWebsiteCertificateAuthorityAssociations) Update(awsWorklinkWebsiteCertificateAuthorityAssociation *v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation) (result *v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsworklinkwebsitecertificateauthorityassociationsResource, awsWorklinkWebsiteCertificateAuthorityAssociation), &v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsWorklinkWebsiteCertificateAuthorityAssociations) UpdateStatus(awsWorklinkWebsiteCertificateAuthorityAssociation *v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation) (*v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsworklinkwebsitecertificateauthorityassociationsResource, "status", awsWorklinkWebsiteCertificateAuthorityAssociation), &v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation), err
}

// Delete takes name of the awsWorklinkWebsiteCertificateAuthorityAssociation and deletes it. Returns an error if one occurs.
func (c *FakeAwsWorklinkWebsiteCertificateAuthorityAssociations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsworklinkwebsitecertificateauthorityassociationsResource, name), &v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsWorklinkWebsiteCertificateAuthorityAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsworklinkwebsitecertificateauthorityassociationsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociationList{})
	return err
}

// Patch applies the patch and returns the patched awsWorklinkWebsiteCertificateAuthorityAssociation.
func (c *FakeAwsWorklinkWebsiteCertificateAuthorityAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsworklinkwebsitecertificateauthorityassociationsResource, name, pt, data, subresources...), &v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWorklinkWebsiteCertificateAuthorityAssociation), err
}