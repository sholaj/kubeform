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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeAwsAppsyncDatasources implements AwsAppsyncDatasourceInterface
type FakeAwsAppsyncDatasources struct {
	Fake *FakeAwsV1alpha1
}

var awsappsyncdatasourcesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsappsyncdatasources"}

var awsappsyncdatasourcesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsAppsyncDatasource"}

// Get takes name of the awsAppsyncDatasource, and returns the corresponding awsAppsyncDatasource object, and an error if there is any.
func (c *FakeAwsAppsyncDatasources) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsAppsyncDatasource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsappsyncdatasourcesResource, name), &v1alpha1.AwsAppsyncDatasource{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAppsyncDatasource), err
}

// List takes label and field selectors, and returns the list of AwsAppsyncDatasources that match those selectors.
func (c *FakeAwsAppsyncDatasources) List(opts v1.ListOptions) (result *v1alpha1.AwsAppsyncDatasourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsappsyncdatasourcesResource, awsappsyncdatasourcesKind, opts), &v1alpha1.AwsAppsyncDatasourceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsAppsyncDatasourceList{ListMeta: obj.(*v1alpha1.AwsAppsyncDatasourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsAppsyncDatasourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsAppsyncDatasources.
func (c *FakeAwsAppsyncDatasources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsappsyncdatasourcesResource, opts))
}

// Create takes the representation of a awsAppsyncDatasource and creates it.  Returns the server's representation of the awsAppsyncDatasource, and an error, if there is any.
func (c *FakeAwsAppsyncDatasources) Create(awsAppsyncDatasource *v1alpha1.AwsAppsyncDatasource) (result *v1alpha1.AwsAppsyncDatasource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsappsyncdatasourcesResource, awsAppsyncDatasource), &v1alpha1.AwsAppsyncDatasource{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAppsyncDatasource), err
}

// Update takes the representation of a awsAppsyncDatasource and updates it. Returns the server's representation of the awsAppsyncDatasource, and an error, if there is any.
func (c *FakeAwsAppsyncDatasources) Update(awsAppsyncDatasource *v1alpha1.AwsAppsyncDatasource) (result *v1alpha1.AwsAppsyncDatasource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsappsyncdatasourcesResource, awsAppsyncDatasource), &v1alpha1.AwsAppsyncDatasource{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAppsyncDatasource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsAppsyncDatasources) UpdateStatus(awsAppsyncDatasource *v1alpha1.AwsAppsyncDatasource) (*v1alpha1.AwsAppsyncDatasource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsappsyncdatasourcesResource, "status", awsAppsyncDatasource), &v1alpha1.AwsAppsyncDatasource{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAppsyncDatasource), err
}

// Delete takes name of the awsAppsyncDatasource and deletes it. Returns an error if one occurs.
func (c *FakeAwsAppsyncDatasources) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsappsyncdatasourcesResource, name), &v1alpha1.AwsAppsyncDatasource{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsAppsyncDatasources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsappsyncdatasourcesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsAppsyncDatasourceList{})
	return err
}

// Patch applies the patch and returns the patched awsAppsyncDatasource.
func (c *FakeAwsAppsyncDatasources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAppsyncDatasource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsappsyncdatasourcesResource, name, pt, data, subresources...), &v1alpha1.AwsAppsyncDatasource{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAppsyncDatasource), err
}