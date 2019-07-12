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
	v1alpha1 "kubeform.dev/kubeform/apis/digitalocean/v1alpha1"
)

// FakeDigitaloceanProjects implements DigitaloceanProjectInterface
type FakeDigitaloceanProjects struct {
	Fake *FakeDigitaloceanV1alpha1
}

var digitaloceanprojectsResource = schema.GroupVersionResource{Group: "digitalocean.kubeform.com", Version: "v1alpha1", Resource: "digitaloceanprojects"}

var digitaloceanprojectsKind = schema.GroupVersionKind{Group: "digitalocean.kubeform.com", Version: "v1alpha1", Kind: "DigitaloceanProject"}

// Get takes name of the digitaloceanProject, and returns the corresponding digitaloceanProject object, and an error if there is any.
func (c *FakeDigitaloceanProjects) Get(name string, options v1.GetOptions) (result *v1alpha1.DigitaloceanProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(digitaloceanprojectsResource, name), &v1alpha1.DigitaloceanProject{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanProject), err
}

// List takes label and field selectors, and returns the list of DigitaloceanProjects that match those selectors.
func (c *FakeDigitaloceanProjects) List(opts v1.ListOptions) (result *v1alpha1.DigitaloceanProjectList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(digitaloceanprojectsResource, digitaloceanprojectsKind, opts), &v1alpha1.DigitaloceanProjectList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DigitaloceanProjectList{ListMeta: obj.(*v1alpha1.DigitaloceanProjectList).ListMeta}
	for _, item := range obj.(*v1alpha1.DigitaloceanProjectList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested digitaloceanProjects.
func (c *FakeDigitaloceanProjects) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(digitaloceanprojectsResource, opts))
}

// Create takes the representation of a digitaloceanProject and creates it.  Returns the server's representation of the digitaloceanProject, and an error, if there is any.
func (c *FakeDigitaloceanProjects) Create(digitaloceanProject *v1alpha1.DigitaloceanProject) (result *v1alpha1.DigitaloceanProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(digitaloceanprojectsResource, digitaloceanProject), &v1alpha1.DigitaloceanProject{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanProject), err
}

// Update takes the representation of a digitaloceanProject and updates it. Returns the server's representation of the digitaloceanProject, and an error, if there is any.
func (c *FakeDigitaloceanProjects) Update(digitaloceanProject *v1alpha1.DigitaloceanProject) (result *v1alpha1.DigitaloceanProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(digitaloceanprojectsResource, digitaloceanProject), &v1alpha1.DigitaloceanProject{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanProject), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDigitaloceanProjects) UpdateStatus(digitaloceanProject *v1alpha1.DigitaloceanProject) (*v1alpha1.DigitaloceanProject, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(digitaloceanprojectsResource, "status", digitaloceanProject), &v1alpha1.DigitaloceanProject{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanProject), err
}

// Delete takes name of the digitaloceanProject and deletes it. Returns an error if one occurs.
func (c *FakeDigitaloceanProjects) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(digitaloceanprojectsResource, name), &v1alpha1.DigitaloceanProject{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDigitaloceanProjects) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(digitaloceanprojectsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DigitaloceanProjectList{})
	return err
}

// Patch applies the patch and returns the patched digitaloceanProject.
func (c *FakeDigitaloceanProjects) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DigitaloceanProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(digitaloceanprojectsResource, name, pt, data, subresources...), &v1alpha1.DigitaloceanProject{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanProject), err
}