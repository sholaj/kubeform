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

// AwsServiceDiscoveryPublicDnsNamespacesGetter has a method to return a AwsServiceDiscoveryPublicDnsNamespaceInterface.
// A group's client should implement this interface.
type AwsServiceDiscoveryPublicDnsNamespacesGetter interface {
	AwsServiceDiscoveryPublicDnsNamespaces() AwsServiceDiscoveryPublicDnsNamespaceInterface
}

// AwsServiceDiscoveryPublicDnsNamespaceInterface has methods to work with AwsServiceDiscoveryPublicDnsNamespace resources.
type AwsServiceDiscoveryPublicDnsNamespaceInterface interface {
	Create(*v1alpha1.AwsServiceDiscoveryPublicDnsNamespace) (*v1alpha1.AwsServiceDiscoveryPublicDnsNamespace, error)
	Update(*v1alpha1.AwsServiceDiscoveryPublicDnsNamespace) (*v1alpha1.AwsServiceDiscoveryPublicDnsNamespace, error)
	UpdateStatus(*v1alpha1.AwsServiceDiscoveryPublicDnsNamespace) (*v1alpha1.AwsServiceDiscoveryPublicDnsNamespace, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsServiceDiscoveryPublicDnsNamespace, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsServiceDiscoveryPublicDnsNamespaceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsServiceDiscoveryPublicDnsNamespace, err error)
	AwsServiceDiscoveryPublicDnsNamespaceExpansion
}

// awsServiceDiscoveryPublicDnsNamespaces implements AwsServiceDiscoveryPublicDnsNamespaceInterface
type awsServiceDiscoveryPublicDnsNamespaces struct {
	client rest.Interface
}

// newAwsServiceDiscoveryPublicDnsNamespaces returns a AwsServiceDiscoveryPublicDnsNamespaces
func newAwsServiceDiscoveryPublicDnsNamespaces(c *AwsV1alpha1Client) *awsServiceDiscoveryPublicDnsNamespaces {
	return &awsServiceDiscoveryPublicDnsNamespaces{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsServiceDiscoveryPublicDnsNamespace, and returns the corresponding awsServiceDiscoveryPublicDnsNamespace object, and an error if there is any.
func (c *awsServiceDiscoveryPublicDnsNamespaces) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsServiceDiscoveryPublicDnsNamespace, err error) {
	result = &v1alpha1.AwsServiceDiscoveryPublicDnsNamespace{}
	err = c.client.Get().
		Resource("awsservicediscoverypublicdnsnamespaces").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsServiceDiscoveryPublicDnsNamespaces that match those selectors.
func (c *awsServiceDiscoveryPublicDnsNamespaces) List(opts v1.ListOptions) (result *v1alpha1.AwsServiceDiscoveryPublicDnsNamespaceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsServiceDiscoveryPublicDnsNamespaceList{}
	err = c.client.Get().
		Resource("awsservicediscoverypublicdnsnamespaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsServiceDiscoveryPublicDnsNamespaces.
func (c *awsServiceDiscoveryPublicDnsNamespaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsservicediscoverypublicdnsnamespaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsServiceDiscoveryPublicDnsNamespace and creates it.  Returns the server's representation of the awsServiceDiscoveryPublicDnsNamespace, and an error, if there is any.
func (c *awsServiceDiscoveryPublicDnsNamespaces) Create(awsServiceDiscoveryPublicDnsNamespace *v1alpha1.AwsServiceDiscoveryPublicDnsNamespace) (result *v1alpha1.AwsServiceDiscoveryPublicDnsNamespace, err error) {
	result = &v1alpha1.AwsServiceDiscoveryPublicDnsNamespace{}
	err = c.client.Post().
		Resource("awsservicediscoverypublicdnsnamespaces").
		Body(awsServiceDiscoveryPublicDnsNamespace).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsServiceDiscoveryPublicDnsNamespace and updates it. Returns the server's representation of the awsServiceDiscoveryPublicDnsNamespace, and an error, if there is any.
func (c *awsServiceDiscoveryPublicDnsNamespaces) Update(awsServiceDiscoveryPublicDnsNamespace *v1alpha1.AwsServiceDiscoveryPublicDnsNamespace) (result *v1alpha1.AwsServiceDiscoveryPublicDnsNamespace, err error) {
	result = &v1alpha1.AwsServiceDiscoveryPublicDnsNamespace{}
	err = c.client.Put().
		Resource("awsservicediscoverypublicdnsnamespaces").
		Name(awsServiceDiscoveryPublicDnsNamespace.Name).
		Body(awsServiceDiscoveryPublicDnsNamespace).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsServiceDiscoveryPublicDnsNamespaces) UpdateStatus(awsServiceDiscoveryPublicDnsNamespace *v1alpha1.AwsServiceDiscoveryPublicDnsNamespace) (result *v1alpha1.AwsServiceDiscoveryPublicDnsNamespace, err error) {
	result = &v1alpha1.AwsServiceDiscoveryPublicDnsNamespace{}
	err = c.client.Put().
		Resource("awsservicediscoverypublicdnsnamespaces").
		Name(awsServiceDiscoveryPublicDnsNamespace.Name).
		SubResource("status").
		Body(awsServiceDiscoveryPublicDnsNamespace).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsServiceDiscoveryPublicDnsNamespace and deletes it. Returns an error if one occurs.
func (c *awsServiceDiscoveryPublicDnsNamespaces) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsservicediscoverypublicdnsnamespaces").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsServiceDiscoveryPublicDnsNamespaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsservicediscoverypublicdnsnamespaces").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsServiceDiscoveryPublicDnsNamespace.
func (c *awsServiceDiscoveryPublicDnsNamespaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsServiceDiscoveryPublicDnsNamespace, err error) {
	result = &v1alpha1.AwsServiceDiscoveryPublicDnsNamespace{}
	err = c.client.Patch(pt).
		Resource("awsservicediscoverypublicdnsnamespaces").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}