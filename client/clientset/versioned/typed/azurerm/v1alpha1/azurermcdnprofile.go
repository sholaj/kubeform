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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// AzurermCdnProfilesGetter has a method to return a AzurermCdnProfileInterface.
// A group's client should implement this interface.
type AzurermCdnProfilesGetter interface {
	AzurermCdnProfiles() AzurermCdnProfileInterface
}

// AzurermCdnProfileInterface has methods to work with AzurermCdnProfile resources.
type AzurermCdnProfileInterface interface {
	Create(*v1alpha1.AzurermCdnProfile) (*v1alpha1.AzurermCdnProfile, error)
	Update(*v1alpha1.AzurermCdnProfile) (*v1alpha1.AzurermCdnProfile, error)
	UpdateStatus(*v1alpha1.AzurermCdnProfile) (*v1alpha1.AzurermCdnProfile, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermCdnProfile, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermCdnProfileList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermCdnProfile, err error)
	AzurermCdnProfileExpansion
}

// azurermCdnProfiles implements AzurermCdnProfileInterface
type azurermCdnProfiles struct {
	client rest.Interface
}

// newAzurermCdnProfiles returns a AzurermCdnProfiles
func newAzurermCdnProfiles(c *AzurermV1alpha1Client) *azurermCdnProfiles {
	return &azurermCdnProfiles{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermCdnProfile, and returns the corresponding azurermCdnProfile object, and an error if there is any.
func (c *azurermCdnProfiles) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermCdnProfile, err error) {
	result = &v1alpha1.AzurermCdnProfile{}
	err = c.client.Get().
		Resource("azurermcdnprofiles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermCdnProfiles that match those selectors.
func (c *azurermCdnProfiles) List(opts v1.ListOptions) (result *v1alpha1.AzurermCdnProfileList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermCdnProfileList{}
	err = c.client.Get().
		Resource("azurermcdnprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermCdnProfiles.
func (c *azurermCdnProfiles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermcdnprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermCdnProfile and creates it.  Returns the server's representation of the azurermCdnProfile, and an error, if there is any.
func (c *azurermCdnProfiles) Create(azurermCdnProfile *v1alpha1.AzurermCdnProfile) (result *v1alpha1.AzurermCdnProfile, err error) {
	result = &v1alpha1.AzurermCdnProfile{}
	err = c.client.Post().
		Resource("azurermcdnprofiles").
		Body(azurermCdnProfile).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermCdnProfile and updates it. Returns the server's representation of the azurermCdnProfile, and an error, if there is any.
func (c *azurermCdnProfiles) Update(azurermCdnProfile *v1alpha1.AzurermCdnProfile) (result *v1alpha1.AzurermCdnProfile, err error) {
	result = &v1alpha1.AzurermCdnProfile{}
	err = c.client.Put().
		Resource("azurermcdnprofiles").
		Name(azurermCdnProfile.Name).
		Body(azurermCdnProfile).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermCdnProfiles) UpdateStatus(azurermCdnProfile *v1alpha1.AzurermCdnProfile) (result *v1alpha1.AzurermCdnProfile, err error) {
	result = &v1alpha1.AzurermCdnProfile{}
	err = c.client.Put().
		Resource("azurermcdnprofiles").
		Name(azurermCdnProfile.Name).
		SubResource("status").
		Body(azurermCdnProfile).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermCdnProfile and deletes it. Returns an error if one occurs.
func (c *azurermCdnProfiles) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermcdnprofiles").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermCdnProfiles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermcdnprofiles").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermCdnProfile.
func (c *azurermCdnProfiles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermCdnProfile, err error) {
	result = &v1alpha1.AzurermCdnProfile{}
	err = c.client.Patch(pt).
		Resource("azurermcdnprofiles").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}