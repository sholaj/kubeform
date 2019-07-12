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

// FakeAwsEc2TransitGatewayRouteTablePropagations implements AwsEc2TransitGatewayRouteTablePropagationInterface
type FakeAwsEc2TransitGatewayRouteTablePropagations struct {
	Fake *FakeAwsV1alpha1
}

var awsec2transitgatewayroutetablepropagationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsec2transitgatewayroutetablepropagations"}

var awsec2transitgatewayroutetablepropagationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsEc2TransitGatewayRouteTablePropagation"}

// Get takes name of the awsEc2TransitGatewayRouteTablePropagation, and returns the corresponding awsEc2TransitGatewayRouteTablePropagation object, and an error if there is any.
func (c *FakeAwsEc2TransitGatewayRouteTablePropagations) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsEc2TransitGatewayRouteTablePropagation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsec2transitgatewayroutetablepropagationsResource, name), &v1alpha1.AwsEc2TransitGatewayRouteTablePropagation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEc2TransitGatewayRouteTablePropagation), err
}

// List takes label and field selectors, and returns the list of AwsEc2TransitGatewayRouteTablePropagations that match those selectors.
func (c *FakeAwsEc2TransitGatewayRouteTablePropagations) List(opts v1.ListOptions) (result *v1alpha1.AwsEc2TransitGatewayRouteTablePropagationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsec2transitgatewayroutetablepropagationsResource, awsec2transitgatewayroutetablepropagationsKind, opts), &v1alpha1.AwsEc2TransitGatewayRouteTablePropagationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsEc2TransitGatewayRouteTablePropagationList{ListMeta: obj.(*v1alpha1.AwsEc2TransitGatewayRouteTablePropagationList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsEc2TransitGatewayRouteTablePropagationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsEc2TransitGatewayRouteTablePropagations.
func (c *FakeAwsEc2TransitGatewayRouteTablePropagations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsec2transitgatewayroutetablepropagationsResource, opts))
}

// Create takes the representation of a awsEc2TransitGatewayRouteTablePropagation and creates it.  Returns the server's representation of the awsEc2TransitGatewayRouteTablePropagation, and an error, if there is any.
func (c *FakeAwsEc2TransitGatewayRouteTablePropagations) Create(awsEc2TransitGatewayRouteTablePropagation *v1alpha1.AwsEc2TransitGatewayRouteTablePropagation) (result *v1alpha1.AwsEc2TransitGatewayRouteTablePropagation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsec2transitgatewayroutetablepropagationsResource, awsEc2TransitGatewayRouteTablePropagation), &v1alpha1.AwsEc2TransitGatewayRouteTablePropagation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEc2TransitGatewayRouteTablePropagation), err
}

// Update takes the representation of a awsEc2TransitGatewayRouteTablePropagation and updates it. Returns the server's representation of the awsEc2TransitGatewayRouteTablePropagation, and an error, if there is any.
func (c *FakeAwsEc2TransitGatewayRouteTablePropagations) Update(awsEc2TransitGatewayRouteTablePropagation *v1alpha1.AwsEc2TransitGatewayRouteTablePropagation) (result *v1alpha1.AwsEc2TransitGatewayRouteTablePropagation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsec2transitgatewayroutetablepropagationsResource, awsEc2TransitGatewayRouteTablePropagation), &v1alpha1.AwsEc2TransitGatewayRouteTablePropagation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEc2TransitGatewayRouteTablePropagation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsEc2TransitGatewayRouteTablePropagations) UpdateStatus(awsEc2TransitGatewayRouteTablePropagation *v1alpha1.AwsEc2TransitGatewayRouteTablePropagation) (*v1alpha1.AwsEc2TransitGatewayRouteTablePropagation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsec2transitgatewayroutetablepropagationsResource, "status", awsEc2TransitGatewayRouteTablePropagation), &v1alpha1.AwsEc2TransitGatewayRouteTablePropagation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEc2TransitGatewayRouteTablePropagation), err
}

// Delete takes name of the awsEc2TransitGatewayRouteTablePropagation and deletes it. Returns an error if one occurs.
func (c *FakeAwsEc2TransitGatewayRouteTablePropagations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsec2transitgatewayroutetablepropagationsResource, name), &v1alpha1.AwsEc2TransitGatewayRouteTablePropagation{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsEc2TransitGatewayRouteTablePropagations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsec2transitgatewayroutetablepropagationsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsEc2TransitGatewayRouteTablePropagationList{})
	return err
}

// Patch applies the patch and returns the patched awsEc2TransitGatewayRouteTablePropagation.
func (c *FakeAwsEc2TransitGatewayRouteTablePropagations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsEc2TransitGatewayRouteTablePropagation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsec2transitgatewayroutetablepropagationsResource, name, pt, data, subresources...), &v1alpha1.AwsEc2TransitGatewayRouteTablePropagation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEc2TransitGatewayRouteTablePropagation), err
}