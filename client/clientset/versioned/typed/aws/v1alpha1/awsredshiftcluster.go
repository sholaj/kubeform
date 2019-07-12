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

// AwsRedshiftClustersGetter has a method to return a AwsRedshiftClusterInterface.
// A group's client should implement this interface.
type AwsRedshiftClustersGetter interface {
	AwsRedshiftClusters() AwsRedshiftClusterInterface
}

// AwsRedshiftClusterInterface has methods to work with AwsRedshiftCluster resources.
type AwsRedshiftClusterInterface interface {
	Create(*v1alpha1.AwsRedshiftCluster) (*v1alpha1.AwsRedshiftCluster, error)
	Update(*v1alpha1.AwsRedshiftCluster) (*v1alpha1.AwsRedshiftCluster, error)
	UpdateStatus(*v1alpha1.AwsRedshiftCluster) (*v1alpha1.AwsRedshiftCluster, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsRedshiftCluster, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsRedshiftClusterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsRedshiftCluster, err error)
	AwsRedshiftClusterExpansion
}

// awsRedshiftClusters implements AwsRedshiftClusterInterface
type awsRedshiftClusters struct {
	client rest.Interface
}

// newAwsRedshiftClusters returns a AwsRedshiftClusters
func newAwsRedshiftClusters(c *AwsV1alpha1Client) *awsRedshiftClusters {
	return &awsRedshiftClusters{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsRedshiftCluster, and returns the corresponding awsRedshiftCluster object, and an error if there is any.
func (c *awsRedshiftClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsRedshiftCluster, err error) {
	result = &v1alpha1.AwsRedshiftCluster{}
	err = c.client.Get().
		Resource("awsredshiftclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsRedshiftClusters that match those selectors.
func (c *awsRedshiftClusters) List(opts v1.ListOptions) (result *v1alpha1.AwsRedshiftClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsRedshiftClusterList{}
	err = c.client.Get().
		Resource("awsredshiftclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsRedshiftClusters.
func (c *awsRedshiftClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsredshiftclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsRedshiftCluster and creates it.  Returns the server's representation of the awsRedshiftCluster, and an error, if there is any.
func (c *awsRedshiftClusters) Create(awsRedshiftCluster *v1alpha1.AwsRedshiftCluster) (result *v1alpha1.AwsRedshiftCluster, err error) {
	result = &v1alpha1.AwsRedshiftCluster{}
	err = c.client.Post().
		Resource("awsredshiftclusters").
		Body(awsRedshiftCluster).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsRedshiftCluster and updates it. Returns the server's representation of the awsRedshiftCluster, and an error, if there is any.
func (c *awsRedshiftClusters) Update(awsRedshiftCluster *v1alpha1.AwsRedshiftCluster) (result *v1alpha1.AwsRedshiftCluster, err error) {
	result = &v1alpha1.AwsRedshiftCluster{}
	err = c.client.Put().
		Resource("awsredshiftclusters").
		Name(awsRedshiftCluster.Name).
		Body(awsRedshiftCluster).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsRedshiftClusters) UpdateStatus(awsRedshiftCluster *v1alpha1.AwsRedshiftCluster) (result *v1alpha1.AwsRedshiftCluster, err error) {
	result = &v1alpha1.AwsRedshiftCluster{}
	err = c.client.Put().
		Resource("awsredshiftclusters").
		Name(awsRedshiftCluster.Name).
		SubResource("status").
		Body(awsRedshiftCluster).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsRedshiftCluster and deletes it. Returns an error if one occurs.
func (c *awsRedshiftClusters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsredshiftclusters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsRedshiftClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsredshiftclusters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsRedshiftCluster.
func (c *awsRedshiftClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsRedshiftCluster, err error) {
	result = &v1alpha1.AwsRedshiftCluster{}
	err = c.client.Patch(pt).
		Resource("awsredshiftclusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}