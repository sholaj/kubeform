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

	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBinaryAuthorizationAttestorIamPolicies implements BinaryAuthorizationAttestorIamPolicyInterface
type FakeBinaryAuthorizationAttestorIamPolicies struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var binaryauthorizationattestoriampoliciesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "binaryauthorizationattestoriampolicies"}

var binaryauthorizationattestoriampoliciesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "BinaryAuthorizationAttestorIamPolicy"}

// Get takes name of the binaryAuthorizationAttestorIamPolicy, and returns the corresponding binaryAuthorizationAttestorIamPolicy object, and an error if there is any.
func (c *FakeBinaryAuthorizationAttestorIamPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BinaryAuthorizationAttestorIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(binaryauthorizationattestoriampoliciesResource, c.ns, name), &v1alpha1.BinaryAuthorizationAttestorIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BinaryAuthorizationAttestorIamPolicy), err
}

// List takes label and field selectors, and returns the list of BinaryAuthorizationAttestorIamPolicies that match those selectors.
func (c *FakeBinaryAuthorizationAttestorIamPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BinaryAuthorizationAttestorIamPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(binaryauthorizationattestoriampoliciesResource, binaryauthorizationattestoriampoliciesKind, c.ns, opts), &v1alpha1.BinaryAuthorizationAttestorIamPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BinaryAuthorizationAttestorIamPolicyList{ListMeta: obj.(*v1alpha1.BinaryAuthorizationAttestorIamPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.BinaryAuthorizationAttestorIamPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested binaryAuthorizationAttestorIamPolicies.
func (c *FakeBinaryAuthorizationAttestorIamPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(binaryauthorizationattestoriampoliciesResource, c.ns, opts))

}

// Create takes the representation of a binaryAuthorizationAttestorIamPolicy and creates it.  Returns the server's representation of the binaryAuthorizationAttestorIamPolicy, and an error, if there is any.
func (c *FakeBinaryAuthorizationAttestorIamPolicies) Create(ctx context.Context, binaryAuthorizationAttestorIamPolicy *v1alpha1.BinaryAuthorizationAttestorIamPolicy, opts v1.CreateOptions) (result *v1alpha1.BinaryAuthorizationAttestorIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(binaryauthorizationattestoriampoliciesResource, c.ns, binaryAuthorizationAttestorIamPolicy), &v1alpha1.BinaryAuthorizationAttestorIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BinaryAuthorizationAttestorIamPolicy), err
}

// Update takes the representation of a binaryAuthorizationAttestorIamPolicy and updates it. Returns the server's representation of the binaryAuthorizationAttestorIamPolicy, and an error, if there is any.
func (c *FakeBinaryAuthorizationAttestorIamPolicies) Update(ctx context.Context, binaryAuthorizationAttestorIamPolicy *v1alpha1.BinaryAuthorizationAttestorIamPolicy, opts v1.UpdateOptions) (result *v1alpha1.BinaryAuthorizationAttestorIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(binaryauthorizationattestoriampoliciesResource, c.ns, binaryAuthorizationAttestorIamPolicy), &v1alpha1.BinaryAuthorizationAttestorIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BinaryAuthorizationAttestorIamPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBinaryAuthorizationAttestorIamPolicies) UpdateStatus(ctx context.Context, binaryAuthorizationAttestorIamPolicy *v1alpha1.BinaryAuthorizationAttestorIamPolicy, opts v1.UpdateOptions) (*v1alpha1.BinaryAuthorizationAttestorIamPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(binaryauthorizationattestoriampoliciesResource, "status", c.ns, binaryAuthorizationAttestorIamPolicy), &v1alpha1.BinaryAuthorizationAttestorIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BinaryAuthorizationAttestorIamPolicy), err
}

// Delete takes name of the binaryAuthorizationAttestorIamPolicy and deletes it. Returns an error if one occurs.
func (c *FakeBinaryAuthorizationAttestorIamPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(binaryauthorizationattestoriampoliciesResource, c.ns, name), &v1alpha1.BinaryAuthorizationAttestorIamPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBinaryAuthorizationAttestorIamPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(binaryauthorizationattestoriampoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BinaryAuthorizationAttestorIamPolicyList{})
	return err
}

// Patch applies the patch and returns the patched binaryAuthorizationAttestorIamPolicy.
func (c *FakeBinaryAuthorizationAttestorIamPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BinaryAuthorizationAttestorIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(binaryauthorizationattestoriampoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.BinaryAuthorizationAttestorIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BinaryAuthorizationAttestorIamPolicy), err
}
