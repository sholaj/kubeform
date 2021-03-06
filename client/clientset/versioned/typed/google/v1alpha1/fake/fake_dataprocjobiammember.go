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

// FakeDataprocJobIamMembers implements DataprocJobIamMemberInterface
type FakeDataprocJobIamMembers struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var dataprocjobiammembersResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "dataprocjobiammembers"}

var dataprocjobiammembersKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "DataprocJobIamMember"}

// Get takes name of the dataprocJobIamMember, and returns the corresponding dataprocJobIamMember object, and an error if there is any.
func (c *FakeDataprocJobIamMembers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DataprocJobIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dataprocjobiammembersResource, c.ns, name), &v1alpha1.DataprocJobIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataprocJobIamMember), err
}

// List takes label and field selectors, and returns the list of DataprocJobIamMembers that match those selectors.
func (c *FakeDataprocJobIamMembers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DataprocJobIamMemberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dataprocjobiammembersResource, dataprocjobiammembersKind, c.ns, opts), &v1alpha1.DataprocJobIamMemberList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DataprocJobIamMemberList{ListMeta: obj.(*v1alpha1.DataprocJobIamMemberList).ListMeta}
	for _, item := range obj.(*v1alpha1.DataprocJobIamMemberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataprocJobIamMembers.
func (c *FakeDataprocJobIamMembers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dataprocjobiammembersResource, c.ns, opts))

}

// Create takes the representation of a dataprocJobIamMember and creates it.  Returns the server's representation of the dataprocJobIamMember, and an error, if there is any.
func (c *FakeDataprocJobIamMembers) Create(ctx context.Context, dataprocJobIamMember *v1alpha1.DataprocJobIamMember, opts v1.CreateOptions) (result *v1alpha1.DataprocJobIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dataprocjobiammembersResource, c.ns, dataprocJobIamMember), &v1alpha1.DataprocJobIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataprocJobIamMember), err
}

// Update takes the representation of a dataprocJobIamMember and updates it. Returns the server's representation of the dataprocJobIamMember, and an error, if there is any.
func (c *FakeDataprocJobIamMembers) Update(ctx context.Context, dataprocJobIamMember *v1alpha1.DataprocJobIamMember, opts v1.UpdateOptions) (result *v1alpha1.DataprocJobIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dataprocjobiammembersResource, c.ns, dataprocJobIamMember), &v1alpha1.DataprocJobIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataprocJobIamMember), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDataprocJobIamMembers) UpdateStatus(ctx context.Context, dataprocJobIamMember *v1alpha1.DataprocJobIamMember, opts v1.UpdateOptions) (*v1alpha1.DataprocJobIamMember, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dataprocjobiammembersResource, "status", c.ns, dataprocJobIamMember), &v1alpha1.DataprocJobIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataprocJobIamMember), err
}

// Delete takes name of the dataprocJobIamMember and deletes it. Returns an error if one occurs.
func (c *FakeDataprocJobIamMembers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dataprocjobiammembersResource, c.ns, name), &v1alpha1.DataprocJobIamMember{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataprocJobIamMembers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dataprocjobiammembersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DataprocJobIamMemberList{})
	return err
}

// Patch applies the patch and returns the patched dataprocJobIamMember.
func (c *FakeDataprocJobIamMembers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DataprocJobIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dataprocjobiammembersResource, c.ns, name, pt, data, subresources...), &v1alpha1.DataprocJobIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataprocJobIamMember), err
}
