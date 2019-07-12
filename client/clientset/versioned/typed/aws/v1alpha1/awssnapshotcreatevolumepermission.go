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

// AwsSnapshotCreateVolumePermissionsGetter has a method to return a AwsSnapshotCreateVolumePermissionInterface.
// A group's client should implement this interface.
type AwsSnapshotCreateVolumePermissionsGetter interface {
	AwsSnapshotCreateVolumePermissions() AwsSnapshotCreateVolumePermissionInterface
}

// AwsSnapshotCreateVolumePermissionInterface has methods to work with AwsSnapshotCreateVolumePermission resources.
type AwsSnapshotCreateVolumePermissionInterface interface {
	Create(*v1alpha1.AwsSnapshotCreateVolumePermission) (*v1alpha1.AwsSnapshotCreateVolumePermission, error)
	Update(*v1alpha1.AwsSnapshotCreateVolumePermission) (*v1alpha1.AwsSnapshotCreateVolumePermission, error)
	UpdateStatus(*v1alpha1.AwsSnapshotCreateVolumePermission) (*v1alpha1.AwsSnapshotCreateVolumePermission, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsSnapshotCreateVolumePermission, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsSnapshotCreateVolumePermissionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSnapshotCreateVolumePermission, err error)
	AwsSnapshotCreateVolumePermissionExpansion
}

// awsSnapshotCreateVolumePermissions implements AwsSnapshotCreateVolumePermissionInterface
type awsSnapshotCreateVolumePermissions struct {
	client rest.Interface
}

// newAwsSnapshotCreateVolumePermissions returns a AwsSnapshotCreateVolumePermissions
func newAwsSnapshotCreateVolumePermissions(c *AwsV1alpha1Client) *awsSnapshotCreateVolumePermissions {
	return &awsSnapshotCreateVolumePermissions{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsSnapshotCreateVolumePermission, and returns the corresponding awsSnapshotCreateVolumePermission object, and an error if there is any.
func (c *awsSnapshotCreateVolumePermissions) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSnapshotCreateVolumePermission, err error) {
	result = &v1alpha1.AwsSnapshotCreateVolumePermission{}
	err = c.client.Get().
		Resource("awssnapshotcreatevolumepermissions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsSnapshotCreateVolumePermissions that match those selectors.
func (c *awsSnapshotCreateVolumePermissions) List(opts v1.ListOptions) (result *v1alpha1.AwsSnapshotCreateVolumePermissionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsSnapshotCreateVolumePermissionList{}
	err = c.client.Get().
		Resource("awssnapshotcreatevolumepermissions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsSnapshotCreateVolumePermissions.
func (c *awsSnapshotCreateVolumePermissions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awssnapshotcreatevolumepermissions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsSnapshotCreateVolumePermission and creates it.  Returns the server's representation of the awsSnapshotCreateVolumePermission, and an error, if there is any.
func (c *awsSnapshotCreateVolumePermissions) Create(awsSnapshotCreateVolumePermission *v1alpha1.AwsSnapshotCreateVolumePermission) (result *v1alpha1.AwsSnapshotCreateVolumePermission, err error) {
	result = &v1alpha1.AwsSnapshotCreateVolumePermission{}
	err = c.client.Post().
		Resource("awssnapshotcreatevolumepermissions").
		Body(awsSnapshotCreateVolumePermission).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsSnapshotCreateVolumePermission and updates it. Returns the server's representation of the awsSnapshotCreateVolumePermission, and an error, if there is any.
func (c *awsSnapshotCreateVolumePermissions) Update(awsSnapshotCreateVolumePermission *v1alpha1.AwsSnapshotCreateVolumePermission) (result *v1alpha1.AwsSnapshotCreateVolumePermission, err error) {
	result = &v1alpha1.AwsSnapshotCreateVolumePermission{}
	err = c.client.Put().
		Resource("awssnapshotcreatevolumepermissions").
		Name(awsSnapshotCreateVolumePermission.Name).
		Body(awsSnapshotCreateVolumePermission).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsSnapshotCreateVolumePermissions) UpdateStatus(awsSnapshotCreateVolumePermission *v1alpha1.AwsSnapshotCreateVolumePermission) (result *v1alpha1.AwsSnapshotCreateVolumePermission, err error) {
	result = &v1alpha1.AwsSnapshotCreateVolumePermission{}
	err = c.client.Put().
		Resource("awssnapshotcreatevolumepermissions").
		Name(awsSnapshotCreateVolumePermission.Name).
		SubResource("status").
		Body(awsSnapshotCreateVolumePermission).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsSnapshotCreateVolumePermission and deletes it. Returns an error if one occurs.
func (c *awsSnapshotCreateVolumePermissions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awssnapshotcreatevolumepermissions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsSnapshotCreateVolumePermissions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awssnapshotcreatevolumepermissions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsSnapshotCreateVolumePermission.
func (c *awsSnapshotCreateVolumePermissions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSnapshotCreateVolumePermission, err error) {
	result = &v1alpha1.AwsSnapshotCreateVolumePermission{}
	err = c.client.Patch(pt).
		Resource("awssnapshotcreatevolumepermissions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}