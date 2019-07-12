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

// FakeAwsBackupPlans implements AwsBackupPlanInterface
type FakeAwsBackupPlans struct {
	Fake *FakeAwsV1alpha1
}

var awsbackupplansResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsbackupplans"}

var awsbackupplansKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsBackupPlan"}

// Get takes name of the awsBackupPlan, and returns the corresponding awsBackupPlan object, and an error if there is any.
func (c *FakeAwsBackupPlans) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsBackupPlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsbackupplansResource, name), &v1alpha1.AwsBackupPlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsBackupPlan), err
}

// List takes label and field selectors, and returns the list of AwsBackupPlans that match those selectors.
func (c *FakeAwsBackupPlans) List(opts v1.ListOptions) (result *v1alpha1.AwsBackupPlanList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsbackupplansResource, awsbackupplansKind, opts), &v1alpha1.AwsBackupPlanList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsBackupPlanList{ListMeta: obj.(*v1alpha1.AwsBackupPlanList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsBackupPlanList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsBackupPlans.
func (c *FakeAwsBackupPlans) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsbackupplansResource, opts))
}

// Create takes the representation of a awsBackupPlan and creates it.  Returns the server's representation of the awsBackupPlan, and an error, if there is any.
func (c *FakeAwsBackupPlans) Create(awsBackupPlan *v1alpha1.AwsBackupPlan) (result *v1alpha1.AwsBackupPlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsbackupplansResource, awsBackupPlan), &v1alpha1.AwsBackupPlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsBackupPlan), err
}

// Update takes the representation of a awsBackupPlan and updates it. Returns the server's representation of the awsBackupPlan, and an error, if there is any.
func (c *FakeAwsBackupPlans) Update(awsBackupPlan *v1alpha1.AwsBackupPlan) (result *v1alpha1.AwsBackupPlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsbackupplansResource, awsBackupPlan), &v1alpha1.AwsBackupPlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsBackupPlan), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsBackupPlans) UpdateStatus(awsBackupPlan *v1alpha1.AwsBackupPlan) (*v1alpha1.AwsBackupPlan, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsbackupplansResource, "status", awsBackupPlan), &v1alpha1.AwsBackupPlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsBackupPlan), err
}

// Delete takes name of the awsBackupPlan and deletes it. Returns an error if one occurs.
func (c *FakeAwsBackupPlans) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsbackupplansResource, name), &v1alpha1.AwsBackupPlan{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsBackupPlans) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsbackupplansResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsBackupPlanList{})
	return err
}

// Patch applies the patch and returns the patched awsBackupPlan.
func (c *FakeAwsBackupPlans) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsBackupPlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsbackupplansResource, name, pt, data, subresources...), &v1alpha1.AwsBackupPlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsBackupPlan), err
}