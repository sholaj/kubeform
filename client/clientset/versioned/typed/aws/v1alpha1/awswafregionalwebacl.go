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

// AwsWafregionalWebAclsGetter has a method to return a AwsWafregionalWebAclInterface.
// A group's client should implement this interface.
type AwsWafregionalWebAclsGetter interface {
	AwsWafregionalWebAcls() AwsWafregionalWebAclInterface
}

// AwsWafregionalWebAclInterface has methods to work with AwsWafregionalWebAcl resources.
type AwsWafregionalWebAclInterface interface {
	Create(*v1alpha1.AwsWafregionalWebAcl) (*v1alpha1.AwsWafregionalWebAcl, error)
	Update(*v1alpha1.AwsWafregionalWebAcl) (*v1alpha1.AwsWafregionalWebAcl, error)
	UpdateStatus(*v1alpha1.AwsWafregionalWebAcl) (*v1alpha1.AwsWafregionalWebAcl, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsWafregionalWebAcl, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsWafregionalWebAclList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafregionalWebAcl, err error)
	AwsWafregionalWebAclExpansion
}

// awsWafregionalWebAcls implements AwsWafregionalWebAclInterface
type awsWafregionalWebAcls struct {
	client rest.Interface
}

// newAwsWafregionalWebAcls returns a AwsWafregionalWebAcls
func newAwsWafregionalWebAcls(c *AwsV1alpha1Client) *awsWafregionalWebAcls {
	return &awsWafregionalWebAcls{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsWafregionalWebAcl, and returns the corresponding awsWafregionalWebAcl object, and an error if there is any.
func (c *awsWafregionalWebAcls) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsWafregionalWebAcl, err error) {
	result = &v1alpha1.AwsWafregionalWebAcl{}
	err = c.client.Get().
		Resource("awswafregionalwebacls").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsWafregionalWebAcls that match those selectors.
func (c *awsWafregionalWebAcls) List(opts v1.ListOptions) (result *v1alpha1.AwsWafregionalWebAclList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsWafregionalWebAclList{}
	err = c.client.Get().
		Resource("awswafregionalwebacls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsWafregionalWebAcls.
func (c *awsWafregionalWebAcls) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awswafregionalwebacls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsWafregionalWebAcl and creates it.  Returns the server's representation of the awsWafregionalWebAcl, and an error, if there is any.
func (c *awsWafregionalWebAcls) Create(awsWafregionalWebAcl *v1alpha1.AwsWafregionalWebAcl) (result *v1alpha1.AwsWafregionalWebAcl, err error) {
	result = &v1alpha1.AwsWafregionalWebAcl{}
	err = c.client.Post().
		Resource("awswafregionalwebacls").
		Body(awsWafregionalWebAcl).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsWafregionalWebAcl and updates it. Returns the server's representation of the awsWafregionalWebAcl, and an error, if there is any.
func (c *awsWafregionalWebAcls) Update(awsWafregionalWebAcl *v1alpha1.AwsWafregionalWebAcl) (result *v1alpha1.AwsWafregionalWebAcl, err error) {
	result = &v1alpha1.AwsWafregionalWebAcl{}
	err = c.client.Put().
		Resource("awswafregionalwebacls").
		Name(awsWafregionalWebAcl.Name).
		Body(awsWafregionalWebAcl).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsWafregionalWebAcls) UpdateStatus(awsWafregionalWebAcl *v1alpha1.AwsWafregionalWebAcl) (result *v1alpha1.AwsWafregionalWebAcl, err error) {
	result = &v1alpha1.AwsWafregionalWebAcl{}
	err = c.client.Put().
		Resource("awswafregionalwebacls").
		Name(awsWafregionalWebAcl.Name).
		SubResource("status").
		Body(awsWafregionalWebAcl).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsWafregionalWebAcl and deletes it. Returns an error if one occurs.
func (c *awsWafregionalWebAcls) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awswafregionalwebacls").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsWafregionalWebAcls) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awswafregionalwebacls").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsWafregionalWebAcl.
func (c *awsWafregionalWebAcls) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafregionalWebAcl, err error) {
	result = &v1alpha1.AwsWafregionalWebAcl{}
	err = c.client.Patch(pt).
		Resource("awswafregionalwebacls").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}