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

// AzurermRecoveryServicesVaultsGetter has a method to return a AzurermRecoveryServicesVaultInterface.
// A group's client should implement this interface.
type AzurermRecoveryServicesVaultsGetter interface {
	AzurermRecoveryServicesVaults() AzurermRecoveryServicesVaultInterface
}

// AzurermRecoveryServicesVaultInterface has methods to work with AzurermRecoveryServicesVault resources.
type AzurermRecoveryServicesVaultInterface interface {
	Create(*v1alpha1.AzurermRecoveryServicesVault) (*v1alpha1.AzurermRecoveryServicesVault, error)
	Update(*v1alpha1.AzurermRecoveryServicesVault) (*v1alpha1.AzurermRecoveryServicesVault, error)
	UpdateStatus(*v1alpha1.AzurermRecoveryServicesVault) (*v1alpha1.AzurermRecoveryServicesVault, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermRecoveryServicesVault, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermRecoveryServicesVaultList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermRecoveryServicesVault, err error)
	AzurermRecoveryServicesVaultExpansion
}

// azurermRecoveryServicesVaults implements AzurermRecoveryServicesVaultInterface
type azurermRecoveryServicesVaults struct {
	client rest.Interface
}

// newAzurermRecoveryServicesVaults returns a AzurermRecoveryServicesVaults
func newAzurermRecoveryServicesVaults(c *AzurermV1alpha1Client) *azurermRecoveryServicesVaults {
	return &azurermRecoveryServicesVaults{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermRecoveryServicesVault, and returns the corresponding azurermRecoveryServicesVault object, and an error if there is any.
func (c *azurermRecoveryServicesVaults) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermRecoveryServicesVault, err error) {
	result = &v1alpha1.AzurermRecoveryServicesVault{}
	err = c.client.Get().
		Resource("azurermrecoveryservicesvaults").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermRecoveryServicesVaults that match those selectors.
func (c *azurermRecoveryServicesVaults) List(opts v1.ListOptions) (result *v1alpha1.AzurermRecoveryServicesVaultList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermRecoveryServicesVaultList{}
	err = c.client.Get().
		Resource("azurermrecoveryservicesvaults").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermRecoveryServicesVaults.
func (c *azurermRecoveryServicesVaults) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermrecoveryservicesvaults").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermRecoveryServicesVault and creates it.  Returns the server's representation of the azurermRecoveryServicesVault, and an error, if there is any.
func (c *azurermRecoveryServicesVaults) Create(azurermRecoveryServicesVault *v1alpha1.AzurermRecoveryServicesVault) (result *v1alpha1.AzurermRecoveryServicesVault, err error) {
	result = &v1alpha1.AzurermRecoveryServicesVault{}
	err = c.client.Post().
		Resource("azurermrecoveryservicesvaults").
		Body(azurermRecoveryServicesVault).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermRecoveryServicesVault and updates it. Returns the server's representation of the azurermRecoveryServicesVault, and an error, if there is any.
func (c *azurermRecoveryServicesVaults) Update(azurermRecoveryServicesVault *v1alpha1.AzurermRecoveryServicesVault) (result *v1alpha1.AzurermRecoveryServicesVault, err error) {
	result = &v1alpha1.AzurermRecoveryServicesVault{}
	err = c.client.Put().
		Resource("azurermrecoveryservicesvaults").
		Name(azurermRecoveryServicesVault.Name).
		Body(azurermRecoveryServicesVault).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermRecoveryServicesVaults) UpdateStatus(azurermRecoveryServicesVault *v1alpha1.AzurermRecoveryServicesVault) (result *v1alpha1.AzurermRecoveryServicesVault, err error) {
	result = &v1alpha1.AzurermRecoveryServicesVault{}
	err = c.client.Put().
		Resource("azurermrecoveryservicesvaults").
		Name(azurermRecoveryServicesVault.Name).
		SubResource("status").
		Body(azurermRecoveryServicesVault).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermRecoveryServicesVault and deletes it. Returns an error if one occurs.
func (c *azurermRecoveryServicesVaults) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermrecoveryservicesvaults").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermRecoveryServicesVaults) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermrecoveryservicesvaults").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermRecoveryServicesVault.
func (c *azurermRecoveryServicesVaults) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermRecoveryServicesVault, err error) {
	result = &v1alpha1.AzurermRecoveryServicesVault{}
	err = c.client.Patch(pt).
		Resource("azurermrecoveryservicesvaults").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}