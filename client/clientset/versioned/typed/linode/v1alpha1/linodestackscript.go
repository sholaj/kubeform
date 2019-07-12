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
	v1alpha1 "kubeform.dev/kubeform/apis/linode/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// LinodeStackscriptsGetter has a method to return a LinodeStackscriptInterface.
// A group's client should implement this interface.
type LinodeStackscriptsGetter interface {
	LinodeStackscripts() LinodeStackscriptInterface
}

// LinodeStackscriptInterface has methods to work with LinodeStackscript resources.
type LinodeStackscriptInterface interface {
	Create(*v1alpha1.LinodeStackscript) (*v1alpha1.LinodeStackscript, error)
	Update(*v1alpha1.LinodeStackscript) (*v1alpha1.LinodeStackscript, error)
	UpdateStatus(*v1alpha1.LinodeStackscript) (*v1alpha1.LinodeStackscript, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.LinodeStackscript, error)
	List(opts v1.ListOptions) (*v1alpha1.LinodeStackscriptList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LinodeStackscript, err error)
	LinodeStackscriptExpansion
}

// linodeStackscripts implements LinodeStackscriptInterface
type linodeStackscripts struct {
	client rest.Interface
}

// newLinodeStackscripts returns a LinodeStackscripts
func newLinodeStackscripts(c *LinodeV1alpha1Client) *linodeStackscripts {
	return &linodeStackscripts{
		client: c.RESTClient(),
	}
}

// Get takes name of the linodeStackscript, and returns the corresponding linodeStackscript object, and an error if there is any.
func (c *linodeStackscripts) Get(name string, options v1.GetOptions) (result *v1alpha1.LinodeStackscript, err error) {
	result = &v1alpha1.LinodeStackscript{}
	err = c.client.Get().
		Resource("linodestackscripts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LinodeStackscripts that match those selectors.
func (c *linodeStackscripts) List(opts v1.ListOptions) (result *v1alpha1.LinodeStackscriptList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LinodeStackscriptList{}
	err = c.client.Get().
		Resource("linodestackscripts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested linodeStackscripts.
func (c *linodeStackscripts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("linodestackscripts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a linodeStackscript and creates it.  Returns the server's representation of the linodeStackscript, and an error, if there is any.
func (c *linodeStackscripts) Create(linodeStackscript *v1alpha1.LinodeStackscript) (result *v1alpha1.LinodeStackscript, err error) {
	result = &v1alpha1.LinodeStackscript{}
	err = c.client.Post().
		Resource("linodestackscripts").
		Body(linodeStackscript).
		Do().
		Into(result)
	return
}

// Update takes the representation of a linodeStackscript and updates it. Returns the server's representation of the linodeStackscript, and an error, if there is any.
func (c *linodeStackscripts) Update(linodeStackscript *v1alpha1.LinodeStackscript) (result *v1alpha1.LinodeStackscript, err error) {
	result = &v1alpha1.LinodeStackscript{}
	err = c.client.Put().
		Resource("linodestackscripts").
		Name(linodeStackscript.Name).
		Body(linodeStackscript).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *linodeStackscripts) UpdateStatus(linodeStackscript *v1alpha1.LinodeStackscript) (result *v1alpha1.LinodeStackscript, err error) {
	result = &v1alpha1.LinodeStackscript{}
	err = c.client.Put().
		Resource("linodestackscripts").
		Name(linodeStackscript.Name).
		SubResource("status").
		Body(linodeStackscript).
		Do().
		Into(result)
	return
}

// Delete takes name of the linodeStackscript and deletes it. Returns an error if one occurs.
func (c *linodeStackscripts) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("linodestackscripts").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *linodeStackscripts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("linodestackscripts").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched linodeStackscript.
func (c *linodeStackscripts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LinodeStackscript, err error) {
	result = &v1alpha1.LinodeStackscript{}
	err = c.client.Patch(pt).
		Resource("linodestackscripts").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}