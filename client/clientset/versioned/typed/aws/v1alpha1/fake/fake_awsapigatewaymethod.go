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

// FakeAwsApiGatewayMethods implements AwsApiGatewayMethodInterface
type FakeAwsApiGatewayMethods struct {
	Fake *FakeAwsV1alpha1
}

var awsapigatewaymethodsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsapigatewaymethods"}

var awsapigatewaymethodsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsApiGatewayMethod"}

// Get takes name of the awsApiGatewayMethod, and returns the corresponding awsApiGatewayMethod object, and an error if there is any.
func (c *FakeAwsApiGatewayMethods) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsApiGatewayMethod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsapigatewaymethodsResource, name), &v1alpha1.AwsApiGatewayMethod{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayMethod), err
}

// List takes label and field selectors, and returns the list of AwsApiGatewayMethods that match those selectors.
func (c *FakeAwsApiGatewayMethods) List(opts v1.ListOptions) (result *v1alpha1.AwsApiGatewayMethodList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsapigatewaymethodsResource, awsapigatewaymethodsKind, opts), &v1alpha1.AwsApiGatewayMethodList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsApiGatewayMethodList{ListMeta: obj.(*v1alpha1.AwsApiGatewayMethodList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsApiGatewayMethodList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsApiGatewayMethods.
func (c *FakeAwsApiGatewayMethods) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsapigatewaymethodsResource, opts))
}

// Create takes the representation of a awsApiGatewayMethod and creates it.  Returns the server's representation of the awsApiGatewayMethod, and an error, if there is any.
func (c *FakeAwsApiGatewayMethods) Create(awsApiGatewayMethod *v1alpha1.AwsApiGatewayMethod) (result *v1alpha1.AwsApiGatewayMethod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsapigatewaymethodsResource, awsApiGatewayMethod), &v1alpha1.AwsApiGatewayMethod{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayMethod), err
}

// Update takes the representation of a awsApiGatewayMethod and updates it. Returns the server's representation of the awsApiGatewayMethod, and an error, if there is any.
func (c *FakeAwsApiGatewayMethods) Update(awsApiGatewayMethod *v1alpha1.AwsApiGatewayMethod) (result *v1alpha1.AwsApiGatewayMethod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsapigatewaymethodsResource, awsApiGatewayMethod), &v1alpha1.AwsApiGatewayMethod{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayMethod), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsApiGatewayMethods) UpdateStatus(awsApiGatewayMethod *v1alpha1.AwsApiGatewayMethod) (*v1alpha1.AwsApiGatewayMethod, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsapigatewaymethodsResource, "status", awsApiGatewayMethod), &v1alpha1.AwsApiGatewayMethod{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayMethod), err
}

// Delete takes name of the awsApiGatewayMethod and deletes it. Returns an error if one occurs.
func (c *FakeAwsApiGatewayMethods) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsapigatewaymethodsResource, name), &v1alpha1.AwsApiGatewayMethod{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsApiGatewayMethods) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsapigatewaymethodsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsApiGatewayMethodList{})
	return err
}

// Patch applies the patch and returns the patched awsApiGatewayMethod.
func (c *FakeAwsApiGatewayMethods) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayMethod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsapigatewaymethodsResource, name, pt, data, subresources...), &v1alpha1.AwsApiGatewayMethod{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayMethod), err
}