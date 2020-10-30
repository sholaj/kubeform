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

// FakeProjectIamAuditConfigs implements ProjectIamAuditConfigInterface
type FakeProjectIamAuditConfigs struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var projectiamauditconfigsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "projectiamauditconfigs"}

var projectiamauditconfigsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ProjectIamAuditConfig"}

// Get takes name of the projectIamAuditConfig, and returns the corresponding projectIamAuditConfig object, and an error if there is any.
func (c *FakeProjectIamAuditConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ProjectIamAuditConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(projectiamauditconfigsResource, c.ns, name), &v1alpha1.ProjectIamAuditConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectIamAuditConfig), err
}

// List takes label and field selectors, and returns the list of ProjectIamAuditConfigs that match those selectors.
func (c *FakeProjectIamAuditConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ProjectIamAuditConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(projectiamauditconfigsResource, projectiamauditconfigsKind, c.ns, opts), &v1alpha1.ProjectIamAuditConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ProjectIamAuditConfigList{ListMeta: obj.(*v1alpha1.ProjectIamAuditConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.ProjectIamAuditConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested projectIamAuditConfigs.
func (c *FakeProjectIamAuditConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(projectiamauditconfigsResource, c.ns, opts))

}

// Create takes the representation of a projectIamAuditConfig and creates it.  Returns the server's representation of the projectIamAuditConfig, and an error, if there is any.
func (c *FakeProjectIamAuditConfigs) Create(ctx context.Context, projectIamAuditConfig *v1alpha1.ProjectIamAuditConfig, opts v1.CreateOptions) (result *v1alpha1.ProjectIamAuditConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(projectiamauditconfigsResource, c.ns, projectIamAuditConfig), &v1alpha1.ProjectIamAuditConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectIamAuditConfig), err
}

// Update takes the representation of a projectIamAuditConfig and updates it. Returns the server's representation of the projectIamAuditConfig, and an error, if there is any.
func (c *FakeProjectIamAuditConfigs) Update(ctx context.Context, projectIamAuditConfig *v1alpha1.ProjectIamAuditConfig, opts v1.UpdateOptions) (result *v1alpha1.ProjectIamAuditConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(projectiamauditconfigsResource, c.ns, projectIamAuditConfig), &v1alpha1.ProjectIamAuditConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectIamAuditConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProjectIamAuditConfigs) UpdateStatus(ctx context.Context, projectIamAuditConfig *v1alpha1.ProjectIamAuditConfig, opts v1.UpdateOptions) (*v1alpha1.ProjectIamAuditConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(projectiamauditconfigsResource, "status", c.ns, projectIamAuditConfig), &v1alpha1.ProjectIamAuditConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectIamAuditConfig), err
}

// Delete takes name of the projectIamAuditConfig and deletes it. Returns an error if one occurs.
func (c *FakeProjectIamAuditConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(projectiamauditconfigsResource, c.ns, name), &v1alpha1.ProjectIamAuditConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProjectIamAuditConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(projectiamauditconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ProjectIamAuditConfigList{})
	return err
}

// Patch applies the patch and returns the patched projectIamAuditConfig.
func (c *FakeProjectIamAuditConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ProjectIamAuditConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(projectiamauditconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ProjectIamAuditConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectIamAuditConfig), err
}