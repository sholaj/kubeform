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

// FakeRuntimeconfigConfigIamMembers implements RuntimeconfigConfigIamMemberInterface
type FakeRuntimeconfigConfigIamMembers struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var runtimeconfigconfigiammembersResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "runtimeconfigconfigiammembers"}

var runtimeconfigconfigiammembersKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "RuntimeconfigConfigIamMember"}

// Get takes name of the runtimeconfigConfigIamMember, and returns the corresponding runtimeconfigConfigIamMember object, and an error if there is any.
func (c *FakeRuntimeconfigConfigIamMembers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RuntimeconfigConfigIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(runtimeconfigconfigiammembersResource, c.ns, name), &v1alpha1.RuntimeconfigConfigIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RuntimeconfigConfigIamMember), err
}

// List takes label and field selectors, and returns the list of RuntimeconfigConfigIamMembers that match those selectors.
func (c *FakeRuntimeconfigConfigIamMembers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RuntimeconfigConfigIamMemberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(runtimeconfigconfigiammembersResource, runtimeconfigconfigiammembersKind, c.ns, opts), &v1alpha1.RuntimeconfigConfigIamMemberList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RuntimeconfigConfigIamMemberList{ListMeta: obj.(*v1alpha1.RuntimeconfigConfigIamMemberList).ListMeta}
	for _, item := range obj.(*v1alpha1.RuntimeconfigConfigIamMemberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested runtimeconfigConfigIamMembers.
func (c *FakeRuntimeconfigConfigIamMembers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(runtimeconfigconfigiammembersResource, c.ns, opts))

}

// Create takes the representation of a runtimeconfigConfigIamMember and creates it.  Returns the server's representation of the runtimeconfigConfigIamMember, and an error, if there is any.
func (c *FakeRuntimeconfigConfigIamMembers) Create(ctx context.Context, runtimeconfigConfigIamMember *v1alpha1.RuntimeconfigConfigIamMember, opts v1.CreateOptions) (result *v1alpha1.RuntimeconfigConfigIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(runtimeconfigconfigiammembersResource, c.ns, runtimeconfigConfigIamMember), &v1alpha1.RuntimeconfigConfigIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RuntimeconfigConfigIamMember), err
}

// Update takes the representation of a runtimeconfigConfigIamMember and updates it. Returns the server's representation of the runtimeconfigConfigIamMember, and an error, if there is any.
func (c *FakeRuntimeconfigConfigIamMembers) Update(ctx context.Context, runtimeconfigConfigIamMember *v1alpha1.RuntimeconfigConfigIamMember, opts v1.UpdateOptions) (result *v1alpha1.RuntimeconfigConfigIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(runtimeconfigconfigiammembersResource, c.ns, runtimeconfigConfigIamMember), &v1alpha1.RuntimeconfigConfigIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RuntimeconfigConfigIamMember), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRuntimeconfigConfigIamMembers) UpdateStatus(ctx context.Context, runtimeconfigConfigIamMember *v1alpha1.RuntimeconfigConfigIamMember, opts v1.UpdateOptions) (*v1alpha1.RuntimeconfigConfigIamMember, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(runtimeconfigconfigiammembersResource, "status", c.ns, runtimeconfigConfigIamMember), &v1alpha1.RuntimeconfigConfigIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RuntimeconfigConfigIamMember), err
}

// Delete takes name of the runtimeconfigConfigIamMember and deletes it. Returns an error if one occurs.
func (c *FakeRuntimeconfigConfigIamMembers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(runtimeconfigconfigiammembersResource, c.ns, name), &v1alpha1.RuntimeconfigConfigIamMember{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRuntimeconfigConfigIamMembers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(runtimeconfigconfigiammembersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RuntimeconfigConfigIamMemberList{})
	return err
}

// Patch applies the patch and returns the patched runtimeconfigConfigIamMember.
func (c *FakeRuntimeconfigConfigIamMembers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RuntimeconfigConfigIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(runtimeconfigconfigiammembersResource, c.ns, name, pt, data, subresources...), &v1alpha1.RuntimeconfigConfigIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RuntimeconfigConfigIamMember), err
}
