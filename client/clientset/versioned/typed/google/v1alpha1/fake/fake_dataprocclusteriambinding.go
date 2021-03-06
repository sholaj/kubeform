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

// FakeDataprocClusterIamBindings implements DataprocClusterIamBindingInterface
type FakeDataprocClusterIamBindings struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var dataprocclusteriambindingsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "dataprocclusteriambindings"}

var dataprocclusteriambindingsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "DataprocClusterIamBinding"}

// Get takes name of the dataprocClusterIamBinding, and returns the corresponding dataprocClusterIamBinding object, and an error if there is any.
func (c *FakeDataprocClusterIamBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DataprocClusterIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dataprocclusteriambindingsResource, c.ns, name), &v1alpha1.DataprocClusterIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataprocClusterIamBinding), err
}

// List takes label and field selectors, and returns the list of DataprocClusterIamBindings that match those selectors.
func (c *FakeDataprocClusterIamBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DataprocClusterIamBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dataprocclusteriambindingsResource, dataprocclusteriambindingsKind, c.ns, opts), &v1alpha1.DataprocClusterIamBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DataprocClusterIamBindingList{ListMeta: obj.(*v1alpha1.DataprocClusterIamBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.DataprocClusterIamBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataprocClusterIamBindings.
func (c *FakeDataprocClusterIamBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dataprocclusteriambindingsResource, c.ns, opts))

}

// Create takes the representation of a dataprocClusterIamBinding and creates it.  Returns the server's representation of the dataprocClusterIamBinding, and an error, if there is any.
func (c *FakeDataprocClusterIamBindings) Create(ctx context.Context, dataprocClusterIamBinding *v1alpha1.DataprocClusterIamBinding, opts v1.CreateOptions) (result *v1alpha1.DataprocClusterIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dataprocclusteriambindingsResource, c.ns, dataprocClusterIamBinding), &v1alpha1.DataprocClusterIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataprocClusterIamBinding), err
}

// Update takes the representation of a dataprocClusterIamBinding and updates it. Returns the server's representation of the dataprocClusterIamBinding, and an error, if there is any.
func (c *FakeDataprocClusterIamBindings) Update(ctx context.Context, dataprocClusterIamBinding *v1alpha1.DataprocClusterIamBinding, opts v1.UpdateOptions) (result *v1alpha1.DataprocClusterIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dataprocclusteriambindingsResource, c.ns, dataprocClusterIamBinding), &v1alpha1.DataprocClusterIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataprocClusterIamBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDataprocClusterIamBindings) UpdateStatus(ctx context.Context, dataprocClusterIamBinding *v1alpha1.DataprocClusterIamBinding, opts v1.UpdateOptions) (*v1alpha1.DataprocClusterIamBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dataprocclusteriambindingsResource, "status", c.ns, dataprocClusterIamBinding), &v1alpha1.DataprocClusterIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataprocClusterIamBinding), err
}

// Delete takes name of the dataprocClusterIamBinding and deletes it. Returns an error if one occurs.
func (c *FakeDataprocClusterIamBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dataprocclusteriambindingsResource, c.ns, name), &v1alpha1.DataprocClusterIamBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataprocClusterIamBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dataprocclusteriambindingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DataprocClusterIamBindingList{})
	return err
}

// Patch applies the patch and returns the patched dataprocClusterIamBinding.
func (c *FakeDataprocClusterIamBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DataprocClusterIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dataprocclusteriambindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DataprocClusterIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataprocClusterIamBinding), err
}
