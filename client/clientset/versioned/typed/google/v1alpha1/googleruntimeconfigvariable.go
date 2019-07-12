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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// GoogleRuntimeconfigVariablesGetter has a method to return a GoogleRuntimeconfigVariableInterface.
// A group's client should implement this interface.
type GoogleRuntimeconfigVariablesGetter interface {
	GoogleRuntimeconfigVariables() GoogleRuntimeconfigVariableInterface
}

// GoogleRuntimeconfigVariableInterface has methods to work with GoogleRuntimeconfigVariable resources.
type GoogleRuntimeconfigVariableInterface interface {
	Create(*v1alpha1.GoogleRuntimeconfigVariable) (*v1alpha1.GoogleRuntimeconfigVariable, error)
	Update(*v1alpha1.GoogleRuntimeconfigVariable) (*v1alpha1.GoogleRuntimeconfigVariable, error)
	UpdateStatus(*v1alpha1.GoogleRuntimeconfigVariable) (*v1alpha1.GoogleRuntimeconfigVariable, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleRuntimeconfigVariable, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleRuntimeconfigVariableList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleRuntimeconfigVariable, err error)
	GoogleRuntimeconfigVariableExpansion
}

// googleRuntimeconfigVariables implements GoogleRuntimeconfigVariableInterface
type googleRuntimeconfigVariables struct {
	client rest.Interface
}

// newGoogleRuntimeconfigVariables returns a GoogleRuntimeconfigVariables
func newGoogleRuntimeconfigVariables(c *GoogleV1alpha1Client) *googleRuntimeconfigVariables {
	return &googleRuntimeconfigVariables{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleRuntimeconfigVariable, and returns the corresponding googleRuntimeconfigVariable object, and an error if there is any.
func (c *googleRuntimeconfigVariables) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleRuntimeconfigVariable, err error) {
	result = &v1alpha1.GoogleRuntimeconfigVariable{}
	err = c.client.Get().
		Resource("googleruntimeconfigvariables").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleRuntimeconfigVariables that match those selectors.
func (c *googleRuntimeconfigVariables) List(opts v1.ListOptions) (result *v1alpha1.GoogleRuntimeconfigVariableList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleRuntimeconfigVariableList{}
	err = c.client.Get().
		Resource("googleruntimeconfigvariables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleRuntimeconfigVariables.
func (c *googleRuntimeconfigVariables) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googleruntimeconfigvariables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleRuntimeconfigVariable and creates it.  Returns the server's representation of the googleRuntimeconfigVariable, and an error, if there is any.
func (c *googleRuntimeconfigVariables) Create(googleRuntimeconfigVariable *v1alpha1.GoogleRuntimeconfigVariable) (result *v1alpha1.GoogleRuntimeconfigVariable, err error) {
	result = &v1alpha1.GoogleRuntimeconfigVariable{}
	err = c.client.Post().
		Resource("googleruntimeconfigvariables").
		Body(googleRuntimeconfigVariable).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleRuntimeconfigVariable and updates it. Returns the server's representation of the googleRuntimeconfigVariable, and an error, if there is any.
func (c *googleRuntimeconfigVariables) Update(googleRuntimeconfigVariable *v1alpha1.GoogleRuntimeconfigVariable) (result *v1alpha1.GoogleRuntimeconfigVariable, err error) {
	result = &v1alpha1.GoogleRuntimeconfigVariable{}
	err = c.client.Put().
		Resource("googleruntimeconfigvariables").
		Name(googleRuntimeconfigVariable.Name).
		Body(googleRuntimeconfigVariable).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleRuntimeconfigVariables) UpdateStatus(googleRuntimeconfigVariable *v1alpha1.GoogleRuntimeconfigVariable) (result *v1alpha1.GoogleRuntimeconfigVariable, err error) {
	result = &v1alpha1.GoogleRuntimeconfigVariable{}
	err = c.client.Put().
		Resource("googleruntimeconfigvariables").
		Name(googleRuntimeconfigVariable.Name).
		SubResource("status").
		Body(googleRuntimeconfigVariable).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleRuntimeconfigVariable and deletes it. Returns an error if one occurs.
func (c *googleRuntimeconfigVariables) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googleruntimeconfigvariables").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleRuntimeconfigVariables) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googleruntimeconfigvariables").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleRuntimeconfigVariable.
func (c *googleRuntimeconfigVariables) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleRuntimeconfigVariable, err error) {
	result = &v1alpha1.GoogleRuntimeconfigVariable{}
	err = c.client.Patch(pt).
		Resource("googleruntimeconfigvariables").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}