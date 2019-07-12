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
	v1alpha1 "kubeform.dev/kubeform/apis/linode/v1alpha1"
)

// FakeLinodeNodebalancers implements LinodeNodebalancerInterface
type FakeLinodeNodebalancers struct {
	Fake *FakeLinodeV1alpha1
}

var linodenodebalancersResource = schema.GroupVersionResource{Group: "linode.kubeform.com", Version: "v1alpha1", Resource: "linodenodebalancers"}

var linodenodebalancersKind = schema.GroupVersionKind{Group: "linode.kubeform.com", Version: "v1alpha1", Kind: "LinodeNodebalancer"}

// Get takes name of the linodeNodebalancer, and returns the corresponding linodeNodebalancer object, and an error if there is any.
func (c *FakeLinodeNodebalancers) Get(name string, options v1.GetOptions) (result *v1alpha1.LinodeNodebalancer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(linodenodebalancersResource, name), &v1alpha1.LinodeNodebalancer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinodeNodebalancer), err
}

// List takes label and field selectors, and returns the list of LinodeNodebalancers that match those selectors.
func (c *FakeLinodeNodebalancers) List(opts v1.ListOptions) (result *v1alpha1.LinodeNodebalancerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(linodenodebalancersResource, linodenodebalancersKind, opts), &v1alpha1.LinodeNodebalancerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LinodeNodebalancerList{ListMeta: obj.(*v1alpha1.LinodeNodebalancerList).ListMeta}
	for _, item := range obj.(*v1alpha1.LinodeNodebalancerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested linodeNodebalancers.
func (c *FakeLinodeNodebalancers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(linodenodebalancersResource, opts))
}

// Create takes the representation of a linodeNodebalancer and creates it.  Returns the server's representation of the linodeNodebalancer, and an error, if there is any.
func (c *FakeLinodeNodebalancers) Create(linodeNodebalancer *v1alpha1.LinodeNodebalancer) (result *v1alpha1.LinodeNodebalancer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(linodenodebalancersResource, linodeNodebalancer), &v1alpha1.LinodeNodebalancer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinodeNodebalancer), err
}

// Update takes the representation of a linodeNodebalancer and updates it. Returns the server's representation of the linodeNodebalancer, and an error, if there is any.
func (c *FakeLinodeNodebalancers) Update(linodeNodebalancer *v1alpha1.LinodeNodebalancer) (result *v1alpha1.LinodeNodebalancer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(linodenodebalancersResource, linodeNodebalancer), &v1alpha1.LinodeNodebalancer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinodeNodebalancer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLinodeNodebalancers) UpdateStatus(linodeNodebalancer *v1alpha1.LinodeNodebalancer) (*v1alpha1.LinodeNodebalancer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(linodenodebalancersResource, "status", linodeNodebalancer), &v1alpha1.LinodeNodebalancer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinodeNodebalancer), err
}

// Delete takes name of the linodeNodebalancer and deletes it. Returns an error if one occurs.
func (c *FakeLinodeNodebalancers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(linodenodebalancersResource, name), &v1alpha1.LinodeNodebalancer{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLinodeNodebalancers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(linodenodebalancersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LinodeNodebalancerList{})
	return err
}

// Patch applies the patch and returns the patched linodeNodebalancer.
func (c *FakeLinodeNodebalancers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LinodeNodebalancer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(linodenodebalancersResource, name, pt, data, subresources...), &v1alpha1.LinodeNodebalancer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinodeNodebalancer), err
}