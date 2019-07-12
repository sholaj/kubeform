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

// FakeAwsIotCertificates implements AwsIotCertificateInterface
type FakeAwsIotCertificates struct {
	Fake *FakeAwsV1alpha1
}

var awsiotcertificatesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsiotcertificates"}

var awsiotcertificatesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsIotCertificate"}

// Get takes name of the awsIotCertificate, and returns the corresponding awsIotCertificate object, and an error if there is any.
func (c *FakeAwsIotCertificates) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsIotCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsiotcertificatesResource, name), &v1alpha1.AwsIotCertificate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIotCertificate), err
}

// List takes label and field selectors, and returns the list of AwsIotCertificates that match those selectors.
func (c *FakeAwsIotCertificates) List(opts v1.ListOptions) (result *v1alpha1.AwsIotCertificateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsiotcertificatesResource, awsiotcertificatesKind, opts), &v1alpha1.AwsIotCertificateList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsIotCertificateList{ListMeta: obj.(*v1alpha1.AwsIotCertificateList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsIotCertificateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsIotCertificates.
func (c *FakeAwsIotCertificates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsiotcertificatesResource, opts))
}

// Create takes the representation of a awsIotCertificate and creates it.  Returns the server's representation of the awsIotCertificate, and an error, if there is any.
func (c *FakeAwsIotCertificates) Create(awsIotCertificate *v1alpha1.AwsIotCertificate) (result *v1alpha1.AwsIotCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsiotcertificatesResource, awsIotCertificate), &v1alpha1.AwsIotCertificate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIotCertificate), err
}

// Update takes the representation of a awsIotCertificate and updates it. Returns the server's representation of the awsIotCertificate, and an error, if there is any.
func (c *FakeAwsIotCertificates) Update(awsIotCertificate *v1alpha1.AwsIotCertificate) (result *v1alpha1.AwsIotCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsiotcertificatesResource, awsIotCertificate), &v1alpha1.AwsIotCertificate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIotCertificate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsIotCertificates) UpdateStatus(awsIotCertificate *v1alpha1.AwsIotCertificate) (*v1alpha1.AwsIotCertificate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsiotcertificatesResource, "status", awsIotCertificate), &v1alpha1.AwsIotCertificate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIotCertificate), err
}

// Delete takes name of the awsIotCertificate and deletes it. Returns an error if one occurs.
func (c *FakeAwsIotCertificates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsiotcertificatesResource, name), &v1alpha1.AwsIotCertificate{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsIotCertificates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsiotcertificatesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsIotCertificateList{})
	return err
}

// Patch applies the patch and returns the patched awsIotCertificate.
func (c *FakeAwsIotCertificates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsIotCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsiotcertificatesResource, name, pt, data, subresources...), &v1alpha1.AwsIotCertificate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIotCertificate), err
}