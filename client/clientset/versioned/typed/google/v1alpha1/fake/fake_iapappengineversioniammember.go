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

// FakeIapAppEngineVersionIamMembers implements IapAppEngineVersionIamMemberInterface
type FakeIapAppEngineVersionIamMembers struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var iapappengineversioniammembersResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "iapappengineversioniammembers"}

var iapappengineversioniammembersKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "IapAppEngineVersionIamMember"}

// Get takes name of the iapAppEngineVersionIamMember, and returns the corresponding iapAppEngineVersionIamMember object, and an error if there is any.
func (c *FakeIapAppEngineVersionIamMembers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IapAppEngineVersionIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iapappengineversioniammembersResource, c.ns, name), &v1alpha1.IapAppEngineVersionIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IapAppEngineVersionIamMember), err
}

// List takes label and field selectors, and returns the list of IapAppEngineVersionIamMembers that match those selectors.
func (c *FakeIapAppEngineVersionIamMembers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IapAppEngineVersionIamMemberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iapappengineversioniammembersResource, iapappengineversioniammembersKind, c.ns, opts), &v1alpha1.IapAppEngineVersionIamMemberList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IapAppEngineVersionIamMemberList{ListMeta: obj.(*v1alpha1.IapAppEngineVersionIamMemberList).ListMeta}
	for _, item := range obj.(*v1alpha1.IapAppEngineVersionIamMemberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iapAppEngineVersionIamMembers.
func (c *FakeIapAppEngineVersionIamMembers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iapappengineversioniammembersResource, c.ns, opts))

}

// Create takes the representation of a iapAppEngineVersionIamMember and creates it.  Returns the server's representation of the iapAppEngineVersionIamMember, and an error, if there is any.
func (c *FakeIapAppEngineVersionIamMembers) Create(ctx context.Context, iapAppEngineVersionIamMember *v1alpha1.IapAppEngineVersionIamMember, opts v1.CreateOptions) (result *v1alpha1.IapAppEngineVersionIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iapappengineversioniammembersResource, c.ns, iapAppEngineVersionIamMember), &v1alpha1.IapAppEngineVersionIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IapAppEngineVersionIamMember), err
}

// Update takes the representation of a iapAppEngineVersionIamMember and updates it. Returns the server's representation of the iapAppEngineVersionIamMember, and an error, if there is any.
func (c *FakeIapAppEngineVersionIamMembers) Update(ctx context.Context, iapAppEngineVersionIamMember *v1alpha1.IapAppEngineVersionIamMember, opts v1.UpdateOptions) (result *v1alpha1.IapAppEngineVersionIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iapappengineversioniammembersResource, c.ns, iapAppEngineVersionIamMember), &v1alpha1.IapAppEngineVersionIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IapAppEngineVersionIamMember), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIapAppEngineVersionIamMembers) UpdateStatus(ctx context.Context, iapAppEngineVersionIamMember *v1alpha1.IapAppEngineVersionIamMember, opts v1.UpdateOptions) (*v1alpha1.IapAppEngineVersionIamMember, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iapappengineversioniammembersResource, "status", c.ns, iapAppEngineVersionIamMember), &v1alpha1.IapAppEngineVersionIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IapAppEngineVersionIamMember), err
}

// Delete takes name of the iapAppEngineVersionIamMember and deletes it. Returns an error if one occurs.
func (c *FakeIapAppEngineVersionIamMembers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iapappengineversioniammembersResource, c.ns, name), &v1alpha1.IapAppEngineVersionIamMember{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIapAppEngineVersionIamMembers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iapappengineversioniammembersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IapAppEngineVersionIamMemberList{})
	return err
}

// Patch applies the patch and returns the patched iapAppEngineVersionIamMember.
func (c *FakeIapAppEngineVersionIamMembers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IapAppEngineVersionIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iapappengineversioniammembersResource, c.ns, name, pt, data, subresources...), &v1alpha1.IapAppEngineVersionIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IapAppEngineVersionIamMember), err
}
