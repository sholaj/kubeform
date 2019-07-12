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

// FakeAwsVpcEndpointServiceAllowedPrincipals implements AwsVpcEndpointServiceAllowedPrincipalInterface
type FakeAwsVpcEndpointServiceAllowedPrincipals struct {
	Fake *FakeAwsV1alpha1
}

var awsvpcendpointserviceallowedprincipalsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsvpcendpointserviceallowedprincipals"}

var awsvpcendpointserviceallowedprincipalsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsVpcEndpointServiceAllowedPrincipal"}

// Get takes name of the awsVpcEndpointServiceAllowedPrincipal, and returns the corresponding awsVpcEndpointServiceAllowedPrincipal object, and an error if there is any.
func (c *FakeAwsVpcEndpointServiceAllowedPrincipals) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsVpcEndpointServiceAllowedPrincipal, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsvpcendpointserviceallowedprincipalsResource, name), &v1alpha1.AwsVpcEndpointServiceAllowedPrincipal{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpcEndpointServiceAllowedPrincipal), err
}

// List takes label and field selectors, and returns the list of AwsVpcEndpointServiceAllowedPrincipals that match those selectors.
func (c *FakeAwsVpcEndpointServiceAllowedPrincipals) List(opts v1.ListOptions) (result *v1alpha1.AwsVpcEndpointServiceAllowedPrincipalList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsvpcendpointserviceallowedprincipalsResource, awsvpcendpointserviceallowedprincipalsKind, opts), &v1alpha1.AwsVpcEndpointServiceAllowedPrincipalList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsVpcEndpointServiceAllowedPrincipalList{ListMeta: obj.(*v1alpha1.AwsVpcEndpointServiceAllowedPrincipalList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsVpcEndpointServiceAllowedPrincipalList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsVpcEndpointServiceAllowedPrincipals.
func (c *FakeAwsVpcEndpointServiceAllowedPrincipals) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsvpcendpointserviceallowedprincipalsResource, opts))
}

// Create takes the representation of a awsVpcEndpointServiceAllowedPrincipal and creates it.  Returns the server's representation of the awsVpcEndpointServiceAllowedPrincipal, and an error, if there is any.
func (c *FakeAwsVpcEndpointServiceAllowedPrincipals) Create(awsVpcEndpointServiceAllowedPrincipal *v1alpha1.AwsVpcEndpointServiceAllowedPrincipal) (result *v1alpha1.AwsVpcEndpointServiceAllowedPrincipal, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsvpcendpointserviceallowedprincipalsResource, awsVpcEndpointServiceAllowedPrincipal), &v1alpha1.AwsVpcEndpointServiceAllowedPrincipal{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpcEndpointServiceAllowedPrincipal), err
}

// Update takes the representation of a awsVpcEndpointServiceAllowedPrincipal and updates it. Returns the server's representation of the awsVpcEndpointServiceAllowedPrincipal, and an error, if there is any.
func (c *FakeAwsVpcEndpointServiceAllowedPrincipals) Update(awsVpcEndpointServiceAllowedPrincipal *v1alpha1.AwsVpcEndpointServiceAllowedPrincipal) (result *v1alpha1.AwsVpcEndpointServiceAllowedPrincipal, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsvpcendpointserviceallowedprincipalsResource, awsVpcEndpointServiceAllowedPrincipal), &v1alpha1.AwsVpcEndpointServiceAllowedPrincipal{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpcEndpointServiceAllowedPrincipal), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsVpcEndpointServiceAllowedPrincipals) UpdateStatus(awsVpcEndpointServiceAllowedPrincipal *v1alpha1.AwsVpcEndpointServiceAllowedPrincipal) (*v1alpha1.AwsVpcEndpointServiceAllowedPrincipal, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsvpcendpointserviceallowedprincipalsResource, "status", awsVpcEndpointServiceAllowedPrincipal), &v1alpha1.AwsVpcEndpointServiceAllowedPrincipal{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpcEndpointServiceAllowedPrincipal), err
}

// Delete takes name of the awsVpcEndpointServiceAllowedPrincipal and deletes it. Returns an error if one occurs.
func (c *FakeAwsVpcEndpointServiceAllowedPrincipals) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsvpcendpointserviceallowedprincipalsResource, name), &v1alpha1.AwsVpcEndpointServiceAllowedPrincipal{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsVpcEndpointServiceAllowedPrincipals) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsvpcendpointserviceallowedprincipalsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsVpcEndpointServiceAllowedPrincipalList{})
	return err
}

// Patch applies the patch and returns the patched awsVpcEndpointServiceAllowedPrincipal.
func (c *FakeAwsVpcEndpointServiceAllowedPrincipals) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsVpcEndpointServiceAllowedPrincipal, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsvpcendpointserviceallowedprincipalsResource, name, pt, data, subresources...), &v1alpha1.AwsVpcEndpointServiceAllowedPrincipal{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpcEndpointServiceAllowedPrincipal), err
}