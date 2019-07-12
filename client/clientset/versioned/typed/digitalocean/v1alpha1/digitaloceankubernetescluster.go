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
	v1alpha1 "kubeform.dev/kubeform/apis/digitalocean/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// DigitaloceanKubernetesClustersGetter has a method to return a DigitaloceanKubernetesClusterInterface.
// A group's client should implement this interface.
type DigitaloceanKubernetesClustersGetter interface {
	DigitaloceanKubernetesClusters() DigitaloceanKubernetesClusterInterface
}

// DigitaloceanKubernetesClusterInterface has methods to work with DigitaloceanKubernetesCluster resources.
type DigitaloceanKubernetesClusterInterface interface {
	Create(*v1alpha1.DigitaloceanKubernetesCluster) (*v1alpha1.DigitaloceanKubernetesCluster, error)
	Update(*v1alpha1.DigitaloceanKubernetesCluster) (*v1alpha1.DigitaloceanKubernetesCluster, error)
	UpdateStatus(*v1alpha1.DigitaloceanKubernetesCluster) (*v1alpha1.DigitaloceanKubernetesCluster, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DigitaloceanKubernetesCluster, error)
	List(opts v1.ListOptions) (*v1alpha1.DigitaloceanKubernetesClusterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DigitaloceanKubernetesCluster, err error)
	DigitaloceanKubernetesClusterExpansion
}

// digitaloceanKubernetesClusters implements DigitaloceanKubernetesClusterInterface
type digitaloceanKubernetesClusters struct {
	client rest.Interface
}

// newDigitaloceanKubernetesClusters returns a DigitaloceanKubernetesClusters
func newDigitaloceanKubernetesClusters(c *DigitaloceanV1alpha1Client) *digitaloceanKubernetesClusters {
	return &digitaloceanKubernetesClusters{
		client: c.RESTClient(),
	}
}

// Get takes name of the digitaloceanKubernetesCluster, and returns the corresponding digitaloceanKubernetesCluster object, and an error if there is any.
func (c *digitaloceanKubernetesClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.DigitaloceanKubernetesCluster, err error) {
	result = &v1alpha1.DigitaloceanKubernetesCluster{}
	err = c.client.Get().
		Resource("digitaloceankubernetesclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DigitaloceanKubernetesClusters that match those selectors.
func (c *digitaloceanKubernetesClusters) List(opts v1.ListOptions) (result *v1alpha1.DigitaloceanKubernetesClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DigitaloceanKubernetesClusterList{}
	err = c.client.Get().
		Resource("digitaloceankubernetesclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested digitaloceanKubernetesClusters.
func (c *digitaloceanKubernetesClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("digitaloceankubernetesclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a digitaloceanKubernetesCluster and creates it.  Returns the server's representation of the digitaloceanKubernetesCluster, and an error, if there is any.
func (c *digitaloceanKubernetesClusters) Create(digitaloceanKubernetesCluster *v1alpha1.DigitaloceanKubernetesCluster) (result *v1alpha1.DigitaloceanKubernetesCluster, err error) {
	result = &v1alpha1.DigitaloceanKubernetesCluster{}
	err = c.client.Post().
		Resource("digitaloceankubernetesclusters").
		Body(digitaloceanKubernetesCluster).
		Do().
		Into(result)
	return
}

// Update takes the representation of a digitaloceanKubernetesCluster and updates it. Returns the server's representation of the digitaloceanKubernetesCluster, and an error, if there is any.
func (c *digitaloceanKubernetesClusters) Update(digitaloceanKubernetesCluster *v1alpha1.DigitaloceanKubernetesCluster) (result *v1alpha1.DigitaloceanKubernetesCluster, err error) {
	result = &v1alpha1.DigitaloceanKubernetesCluster{}
	err = c.client.Put().
		Resource("digitaloceankubernetesclusters").
		Name(digitaloceanKubernetesCluster.Name).
		Body(digitaloceanKubernetesCluster).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *digitaloceanKubernetesClusters) UpdateStatus(digitaloceanKubernetesCluster *v1alpha1.DigitaloceanKubernetesCluster) (result *v1alpha1.DigitaloceanKubernetesCluster, err error) {
	result = &v1alpha1.DigitaloceanKubernetesCluster{}
	err = c.client.Put().
		Resource("digitaloceankubernetesclusters").
		Name(digitaloceanKubernetesCluster.Name).
		SubResource("status").
		Body(digitaloceanKubernetesCluster).
		Do().
		Into(result)
	return
}

// Delete takes name of the digitaloceanKubernetesCluster and deletes it. Returns an error if one occurs.
func (c *digitaloceanKubernetesClusters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("digitaloceankubernetesclusters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *digitaloceanKubernetesClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("digitaloceankubernetesclusters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched digitaloceanKubernetesCluster.
func (c *digitaloceanKubernetesClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DigitaloceanKubernetesCluster, err error) {
	result = &v1alpha1.DigitaloceanKubernetesCluster{}
	err = c.client.Patch(pt).
		Resource("digitaloceankubernetesclusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}