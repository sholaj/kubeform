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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeGoogleComputeRouterInterfaces implements GoogleComputeRouterInterfaceInterface
type FakeGoogleComputeRouterInterfaces struct {
	Fake *FakeGoogleV1alpha1
}

var googlecomputerouterinterfacesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlecomputerouterinterfaces"}

var googlecomputerouterinterfacesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleComputeRouterInterface"}

// Get takes name of the googleComputeRouterInterface, and returns the corresponding googleComputeRouterInterface object, and an error if there is any.
func (c *FakeGoogleComputeRouterInterfaces) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleComputeRouterInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlecomputerouterinterfacesResource, name), &v1alpha1.GoogleComputeRouterInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeRouterInterface), err
}

// List takes label and field selectors, and returns the list of GoogleComputeRouterInterfaces that match those selectors.
func (c *FakeGoogleComputeRouterInterfaces) List(opts v1.ListOptions) (result *v1alpha1.GoogleComputeRouterInterfaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlecomputerouterinterfacesResource, googlecomputerouterinterfacesKind, opts), &v1alpha1.GoogleComputeRouterInterfaceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleComputeRouterInterfaceList{ListMeta: obj.(*v1alpha1.GoogleComputeRouterInterfaceList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleComputeRouterInterfaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleComputeRouterInterfaces.
func (c *FakeGoogleComputeRouterInterfaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlecomputerouterinterfacesResource, opts))
}

// Create takes the representation of a googleComputeRouterInterface and creates it.  Returns the server's representation of the googleComputeRouterInterface, and an error, if there is any.
func (c *FakeGoogleComputeRouterInterfaces) Create(googleComputeRouterInterface *v1alpha1.GoogleComputeRouterInterface) (result *v1alpha1.GoogleComputeRouterInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlecomputerouterinterfacesResource, googleComputeRouterInterface), &v1alpha1.GoogleComputeRouterInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeRouterInterface), err
}

// Update takes the representation of a googleComputeRouterInterface and updates it. Returns the server's representation of the googleComputeRouterInterface, and an error, if there is any.
func (c *FakeGoogleComputeRouterInterfaces) Update(googleComputeRouterInterface *v1alpha1.GoogleComputeRouterInterface) (result *v1alpha1.GoogleComputeRouterInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlecomputerouterinterfacesResource, googleComputeRouterInterface), &v1alpha1.GoogleComputeRouterInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeRouterInterface), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleComputeRouterInterfaces) UpdateStatus(googleComputeRouterInterface *v1alpha1.GoogleComputeRouterInterface) (*v1alpha1.GoogleComputeRouterInterface, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlecomputerouterinterfacesResource, "status", googleComputeRouterInterface), &v1alpha1.GoogleComputeRouterInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeRouterInterface), err
}

// Delete takes name of the googleComputeRouterInterface and deletes it. Returns an error if one occurs.
func (c *FakeGoogleComputeRouterInterfaces) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlecomputerouterinterfacesResource, name), &v1alpha1.GoogleComputeRouterInterface{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleComputeRouterInterfaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlecomputerouterinterfacesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleComputeRouterInterfaceList{})
	return err
}

// Patch applies the patch and returns the patched googleComputeRouterInterface.
func (c *FakeGoogleComputeRouterInterfaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComputeRouterInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlecomputerouterinterfacesResource, name, pt, data, subresources...), &v1alpha1.GoogleComputeRouterInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeRouterInterface), err
}