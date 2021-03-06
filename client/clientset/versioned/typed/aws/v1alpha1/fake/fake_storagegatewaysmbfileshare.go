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

// FakeStoragegatewaySmbFileShares implements StoragegatewaySmbFileShareInterface
type FakeStoragegatewaySmbFileShares struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var storagegatewaysmbfilesharesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "storagegatewaysmbfileshares"}

var storagegatewaysmbfilesharesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "StoragegatewaySmbFileShare"}

// Get takes name of the storagegatewaySmbFileShare, and returns the corresponding storagegatewaySmbFileShare object, and an error if there is any.
func (c *FakeStoragegatewaySmbFileShares) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.StoragegatewaySmbFileShare, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(storagegatewaysmbfilesharesResource, c.ns, name), &v1alpha1.StoragegatewaySmbFileShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StoragegatewaySmbFileShare), err
}

// List takes label and field selectors, and returns the list of StoragegatewaySmbFileShares that match those selectors.
func (c *FakeStoragegatewaySmbFileShares) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.StoragegatewaySmbFileShareList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(storagegatewaysmbfilesharesResource, storagegatewaysmbfilesharesKind, c.ns, opts), &v1alpha1.StoragegatewaySmbFileShareList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StoragegatewaySmbFileShareList{ListMeta: obj.(*v1alpha1.StoragegatewaySmbFileShareList).ListMeta}
	for _, item := range obj.(*v1alpha1.StoragegatewaySmbFileShareList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storagegatewaySmbFileShares.
func (c *FakeStoragegatewaySmbFileShares) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(storagegatewaysmbfilesharesResource, c.ns, opts))

}

// Create takes the representation of a storagegatewaySmbFileShare and creates it.  Returns the server's representation of the storagegatewaySmbFileShare, and an error, if there is any.
func (c *FakeStoragegatewaySmbFileShares) Create(ctx context.Context, storagegatewaySmbFileShare *v1alpha1.StoragegatewaySmbFileShare, opts v1.CreateOptions) (result *v1alpha1.StoragegatewaySmbFileShare, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(storagegatewaysmbfilesharesResource, c.ns, storagegatewaySmbFileShare), &v1alpha1.StoragegatewaySmbFileShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StoragegatewaySmbFileShare), err
}

// Update takes the representation of a storagegatewaySmbFileShare and updates it. Returns the server's representation of the storagegatewaySmbFileShare, and an error, if there is any.
func (c *FakeStoragegatewaySmbFileShares) Update(ctx context.Context, storagegatewaySmbFileShare *v1alpha1.StoragegatewaySmbFileShare, opts v1.UpdateOptions) (result *v1alpha1.StoragegatewaySmbFileShare, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(storagegatewaysmbfilesharesResource, c.ns, storagegatewaySmbFileShare), &v1alpha1.StoragegatewaySmbFileShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StoragegatewaySmbFileShare), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStoragegatewaySmbFileShares) UpdateStatus(ctx context.Context, storagegatewaySmbFileShare *v1alpha1.StoragegatewaySmbFileShare, opts v1.UpdateOptions) (*v1alpha1.StoragegatewaySmbFileShare, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(storagegatewaysmbfilesharesResource, "status", c.ns, storagegatewaySmbFileShare), &v1alpha1.StoragegatewaySmbFileShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StoragegatewaySmbFileShare), err
}

// Delete takes name of the storagegatewaySmbFileShare and deletes it. Returns an error if one occurs.
func (c *FakeStoragegatewaySmbFileShares) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(storagegatewaysmbfilesharesResource, c.ns, name), &v1alpha1.StoragegatewaySmbFileShare{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStoragegatewaySmbFileShares) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(storagegatewaysmbfilesharesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.StoragegatewaySmbFileShareList{})
	return err
}

// Patch applies the patch and returns the patched storagegatewaySmbFileShare.
func (c *FakeStoragegatewaySmbFileShares) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.StoragegatewaySmbFileShare, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(storagegatewaysmbfilesharesResource, c.ns, name, pt, data, subresources...), &v1alpha1.StoragegatewaySmbFileShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StoragegatewaySmbFileShare), err
}
