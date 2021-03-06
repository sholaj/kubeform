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

// FakeVpcPeeringConnections implements VpcPeeringConnectionInterface
type FakeVpcPeeringConnections struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var vpcpeeringconnectionsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "vpcpeeringconnections"}

var vpcpeeringconnectionsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "VpcPeeringConnection"}

// Get takes name of the vpcPeeringConnection, and returns the corresponding vpcPeeringConnection object, and an error if there is any.
func (c *FakeVpcPeeringConnections) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VpcPeeringConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vpcpeeringconnectionsResource, c.ns, name), &v1alpha1.VpcPeeringConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcPeeringConnection), err
}

// List takes label and field selectors, and returns the list of VpcPeeringConnections that match those selectors.
func (c *FakeVpcPeeringConnections) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VpcPeeringConnectionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vpcpeeringconnectionsResource, vpcpeeringconnectionsKind, c.ns, opts), &v1alpha1.VpcPeeringConnectionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VpcPeeringConnectionList{ListMeta: obj.(*v1alpha1.VpcPeeringConnectionList).ListMeta}
	for _, item := range obj.(*v1alpha1.VpcPeeringConnectionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vpcPeeringConnections.
func (c *FakeVpcPeeringConnections) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vpcpeeringconnectionsResource, c.ns, opts))

}

// Create takes the representation of a vpcPeeringConnection and creates it.  Returns the server's representation of the vpcPeeringConnection, and an error, if there is any.
func (c *FakeVpcPeeringConnections) Create(ctx context.Context, vpcPeeringConnection *v1alpha1.VpcPeeringConnection, opts v1.CreateOptions) (result *v1alpha1.VpcPeeringConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vpcpeeringconnectionsResource, c.ns, vpcPeeringConnection), &v1alpha1.VpcPeeringConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcPeeringConnection), err
}

// Update takes the representation of a vpcPeeringConnection and updates it. Returns the server's representation of the vpcPeeringConnection, and an error, if there is any.
func (c *FakeVpcPeeringConnections) Update(ctx context.Context, vpcPeeringConnection *v1alpha1.VpcPeeringConnection, opts v1.UpdateOptions) (result *v1alpha1.VpcPeeringConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vpcpeeringconnectionsResource, c.ns, vpcPeeringConnection), &v1alpha1.VpcPeeringConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcPeeringConnection), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVpcPeeringConnections) UpdateStatus(ctx context.Context, vpcPeeringConnection *v1alpha1.VpcPeeringConnection, opts v1.UpdateOptions) (*v1alpha1.VpcPeeringConnection, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vpcpeeringconnectionsResource, "status", c.ns, vpcPeeringConnection), &v1alpha1.VpcPeeringConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcPeeringConnection), err
}

// Delete takes name of the vpcPeeringConnection and deletes it. Returns an error if one occurs.
func (c *FakeVpcPeeringConnections) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(vpcpeeringconnectionsResource, c.ns, name), &v1alpha1.VpcPeeringConnection{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVpcPeeringConnections) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vpcpeeringconnectionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VpcPeeringConnectionList{})
	return err
}

// Patch applies the patch and returns the patched vpcPeeringConnection.
func (c *FakeVpcPeeringConnections) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VpcPeeringConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vpcpeeringconnectionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.VpcPeeringConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcPeeringConnection), err
}
