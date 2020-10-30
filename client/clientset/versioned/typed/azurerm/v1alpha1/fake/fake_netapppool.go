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

	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetappPools implements NetappPoolInterface
type FakeNetappPools struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var netapppoolsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "netapppools"}

var netapppoolsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "NetappPool"}

// Get takes name of the netappPool, and returns the corresponding netappPool object, and an error if there is any.
func (c *FakeNetappPools) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NetappPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(netapppoolsResource, c.ns, name), &v1alpha1.NetappPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetappPool), err
}

// List takes label and field selectors, and returns the list of NetappPools that match those selectors.
func (c *FakeNetappPools) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NetappPoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(netapppoolsResource, netapppoolsKind, c.ns, opts), &v1alpha1.NetappPoolList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetappPoolList{ListMeta: obj.(*v1alpha1.NetappPoolList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetappPoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested netappPools.
func (c *FakeNetappPools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(netapppoolsResource, c.ns, opts))

}

// Create takes the representation of a netappPool and creates it.  Returns the server's representation of the netappPool, and an error, if there is any.
func (c *FakeNetappPools) Create(ctx context.Context, netappPool *v1alpha1.NetappPool, opts v1.CreateOptions) (result *v1alpha1.NetappPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(netapppoolsResource, c.ns, netappPool), &v1alpha1.NetappPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetappPool), err
}

// Update takes the representation of a netappPool and updates it. Returns the server's representation of the netappPool, and an error, if there is any.
func (c *FakeNetappPools) Update(ctx context.Context, netappPool *v1alpha1.NetappPool, opts v1.UpdateOptions) (result *v1alpha1.NetappPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(netapppoolsResource, c.ns, netappPool), &v1alpha1.NetappPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetappPool), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetappPools) UpdateStatus(ctx context.Context, netappPool *v1alpha1.NetappPool, opts v1.UpdateOptions) (*v1alpha1.NetappPool, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(netapppoolsResource, "status", c.ns, netappPool), &v1alpha1.NetappPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetappPool), err
}

// Delete takes name of the netappPool and deletes it. Returns an error if one occurs.
func (c *FakeNetappPools) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(netapppoolsResource, c.ns, name), &v1alpha1.NetappPool{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetappPools) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(netapppoolsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetappPoolList{})
	return err
}

// Patch applies the patch and returns the patched netappPool.
func (c *FakeNetappPools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NetappPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(netapppoolsResource, c.ns, name, pt, data, subresources...), &v1alpha1.NetappPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetappPool), err
}