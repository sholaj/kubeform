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

package v1alpha1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// AwsBackupPlansGetter has a method to return a AwsBackupPlanInterface.
// A group's client should implement this interface.
type AwsBackupPlansGetter interface {
	AwsBackupPlans() AwsBackupPlanInterface
}

// AwsBackupPlanInterface has methods to work with AwsBackupPlan resources.
type AwsBackupPlanInterface interface {
	Create(*v1alpha1.AwsBackupPlan) (*v1alpha1.AwsBackupPlan, error)
	Update(*v1alpha1.AwsBackupPlan) (*v1alpha1.AwsBackupPlan, error)
	UpdateStatus(*v1alpha1.AwsBackupPlan) (*v1alpha1.AwsBackupPlan, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsBackupPlan, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsBackupPlanList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsBackupPlan, err error)
	AwsBackupPlanExpansion
}

// awsBackupPlans implements AwsBackupPlanInterface
type awsBackupPlans struct {
	client rest.Interface
}

// newAwsBackupPlans returns a AwsBackupPlans
func newAwsBackupPlans(c *AwsV1alpha1Client) *awsBackupPlans {
	return &awsBackupPlans{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsBackupPlan, and returns the corresponding awsBackupPlan object, and an error if there is any.
func (c *awsBackupPlans) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsBackupPlan, err error) {
	result = &v1alpha1.AwsBackupPlan{}
	err = c.client.Get().
		Resource("awsbackupplans").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsBackupPlans that match those selectors.
func (c *awsBackupPlans) List(opts v1.ListOptions) (result *v1alpha1.AwsBackupPlanList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsBackupPlanList{}
	err = c.client.Get().
		Resource("awsbackupplans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsBackupPlans.
func (c *awsBackupPlans) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsbackupplans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsBackupPlan and creates it.  Returns the server's representation of the awsBackupPlan, and an error, if there is any.
func (c *awsBackupPlans) Create(awsBackupPlan *v1alpha1.AwsBackupPlan) (result *v1alpha1.AwsBackupPlan, err error) {
	result = &v1alpha1.AwsBackupPlan{}
	err = c.client.Post().
		Resource("awsbackupplans").
		Body(awsBackupPlan).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsBackupPlan and updates it. Returns the server's representation of the awsBackupPlan, and an error, if there is any.
func (c *awsBackupPlans) Update(awsBackupPlan *v1alpha1.AwsBackupPlan) (result *v1alpha1.AwsBackupPlan, err error) {
	result = &v1alpha1.AwsBackupPlan{}
	err = c.client.Put().
		Resource("awsbackupplans").
		Name(awsBackupPlan.Name).
		Body(awsBackupPlan).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsBackupPlans) UpdateStatus(awsBackupPlan *v1alpha1.AwsBackupPlan) (result *v1alpha1.AwsBackupPlan, err error) {
	result = &v1alpha1.AwsBackupPlan{}
	err = c.client.Put().
		Resource("awsbackupplans").
		Name(awsBackupPlan.Name).
		SubResource("status").
		Body(awsBackupPlan).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsBackupPlan and deletes it. Returns an error if one occurs.
func (c *awsBackupPlans) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsbackupplans").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsBackupPlans) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsbackupplans").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsBackupPlan.
func (c *awsBackupPlans) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsBackupPlan, err error) {
	result = &v1alpha1.AwsBackupPlan{}
	err = c.client.Patch(pt).
		Resource("awsbackupplans").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}