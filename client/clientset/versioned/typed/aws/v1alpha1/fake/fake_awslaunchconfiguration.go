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

// FakeAwsLaunchConfigurations implements AwsLaunchConfigurationInterface
type FakeAwsLaunchConfigurations struct {
	Fake *FakeAwsV1alpha1
}

var awslaunchconfigurationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awslaunchconfigurations"}

var awslaunchconfigurationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsLaunchConfiguration"}

// Get takes name of the awsLaunchConfiguration, and returns the corresponding awsLaunchConfiguration object, and an error if there is any.
func (c *FakeAwsLaunchConfigurations) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsLaunchConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awslaunchconfigurationsResource, name), &v1alpha1.AwsLaunchConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLaunchConfiguration), err
}

// List takes label and field selectors, and returns the list of AwsLaunchConfigurations that match those selectors.
func (c *FakeAwsLaunchConfigurations) List(opts v1.ListOptions) (result *v1alpha1.AwsLaunchConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awslaunchconfigurationsResource, awslaunchconfigurationsKind, opts), &v1alpha1.AwsLaunchConfigurationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsLaunchConfigurationList{ListMeta: obj.(*v1alpha1.AwsLaunchConfigurationList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsLaunchConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsLaunchConfigurations.
func (c *FakeAwsLaunchConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awslaunchconfigurationsResource, opts))
}

// Create takes the representation of a awsLaunchConfiguration and creates it.  Returns the server's representation of the awsLaunchConfiguration, and an error, if there is any.
func (c *FakeAwsLaunchConfigurations) Create(awsLaunchConfiguration *v1alpha1.AwsLaunchConfiguration) (result *v1alpha1.AwsLaunchConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awslaunchconfigurationsResource, awsLaunchConfiguration), &v1alpha1.AwsLaunchConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLaunchConfiguration), err
}

// Update takes the representation of a awsLaunchConfiguration and updates it. Returns the server's representation of the awsLaunchConfiguration, and an error, if there is any.
func (c *FakeAwsLaunchConfigurations) Update(awsLaunchConfiguration *v1alpha1.AwsLaunchConfiguration) (result *v1alpha1.AwsLaunchConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awslaunchconfigurationsResource, awsLaunchConfiguration), &v1alpha1.AwsLaunchConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLaunchConfiguration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsLaunchConfigurations) UpdateStatus(awsLaunchConfiguration *v1alpha1.AwsLaunchConfiguration) (*v1alpha1.AwsLaunchConfiguration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awslaunchconfigurationsResource, "status", awsLaunchConfiguration), &v1alpha1.AwsLaunchConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLaunchConfiguration), err
}

// Delete takes name of the awsLaunchConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeAwsLaunchConfigurations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awslaunchconfigurationsResource, name), &v1alpha1.AwsLaunchConfiguration{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsLaunchConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awslaunchconfigurationsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsLaunchConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched awsLaunchConfiguration.
func (c *FakeAwsLaunchConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsLaunchConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awslaunchconfigurationsResource, name, pt, data, subresources...), &v1alpha1.AwsLaunchConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLaunchConfiguration), err
}