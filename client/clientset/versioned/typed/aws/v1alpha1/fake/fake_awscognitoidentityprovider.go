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

// FakeAwsCognitoIdentityProviders implements AwsCognitoIdentityProviderInterface
type FakeAwsCognitoIdentityProviders struct {
	Fake *FakeAwsV1alpha1
}

var awscognitoidentityprovidersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awscognitoidentityproviders"}

var awscognitoidentityprovidersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsCognitoIdentityProvider"}

// Get takes name of the awsCognitoIdentityProvider, and returns the corresponding awsCognitoIdentityProvider object, and an error if there is any.
func (c *FakeAwsCognitoIdentityProviders) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsCognitoIdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awscognitoidentityprovidersResource, name), &v1alpha1.AwsCognitoIdentityProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCognitoIdentityProvider), err
}

// List takes label and field selectors, and returns the list of AwsCognitoIdentityProviders that match those selectors.
func (c *FakeAwsCognitoIdentityProviders) List(opts v1.ListOptions) (result *v1alpha1.AwsCognitoIdentityProviderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awscognitoidentityprovidersResource, awscognitoidentityprovidersKind, opts), &v1alpha1.AwsCognitoIdentityProviderList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsCognitoIdentityProviderList{ListMeta: obj.(*v1alpha1.AwsCognitoIdentityProviderList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsCognitoIdentityProviderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsCognitoIdentityProviders.
func (c *FakeAwsCognitoIdentityProviders) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awscognitoidentityprovidersResource, opts))
}

// Create takes the representation of a awsCognitoIdentityProvider and creates it.  Returns the server's representation of the awsCognitoIdentityProvider, and an error, if there is any.
func (c *FakeAwsCognitoIdentityProviders) Create(awsCognitoIdentityProvider *v1alpha1.AwsCognitoIdentityProvider) (result *v1alpha1.AwsCognitoIdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awscognitoidentityprovidersResource, awsCognitoIdentityProvider), &v1alpha1.AwsCognitoIdentityProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCognitoIdentityProvider), err
}

// Update takes the representation of a awsCognitoIdentityProvider and updates it. Returns the server's representation of the awsCognitoIdentityProvider, and an error, if there is any.
func (c *FakeAwsCognitoIdentityProviders) Update(awsCognitoIdentityProvider *v1alpha1.AwsCognitoIdentityProvider) (result *v1alpha1.AwsCognitoIdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awscognitoidentityprovidersResource, awsCognitoIdentityProvider), &v1alpha1.AwsCognitoIdentityProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCognitoIdentityProvider), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsCognitoIdentityProviders) UpdateStatus(awsCognitoIdentityProvider *v1alpha1.AwsCognitoIdentityProvider) (*v1alpha1.AwsCognitoIdentityProvider, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awscognitoidentityprovidersResource, "status", awsCognitoIdentityProvider), &v1alpha1.AwsCognitoIdentityProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCognitoIdentityProvider), err
}

// Delete takes name of the awsCognitoIdentityProvider and deletes it. Returns an error if one occurs.
func (c *FakeAwsCognitoIdentityProviders) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awscognitoidentityprovidersResource, name), &v1alpha1.AwsCognitoIdentityProvider{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsCognitoIdentityProviders) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awscognitoidentityprovidersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsCognitoIdentityProviderList{})
	return err
}

// Patch applies the patch and returns the patched awsCognitoIdentityProvider.
func (c *FakeAwsCognitoIdentityProviders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCognitoIdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awscognitoidentityprovidersResource, name, pt, data, subresources...), &v1alpha1.AwsCognitoIdentityProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCognitoIdentityProvider), err
}