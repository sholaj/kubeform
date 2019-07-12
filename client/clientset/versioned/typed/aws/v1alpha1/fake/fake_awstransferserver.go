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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeAwsTransferServers implements AwsTransferServerInterface
type FakeAwsTransferServers struct {
	Fake *FakeAwsV1alpha1
}

var awstransferserversResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awstransferservers"}

var awstransferserversKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsTransferServer"}

// Get takes name of the awsTransferServer, and returns the corresponding awsTransferServer object, and an error if there is any.
func (c *FakeAwsTransferServers) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsTransferServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awstransferserversResource, name), &v1alpha1.AwsTransferServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsTransferServer), err
}

// List takes label and field selectors, and returns the list of AwsTransferServers that match those selectors.
func (c *FakeAwsTransferServers) List(opts v1.ListOptions) (result *v1alpha1.AwsTransferServerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awstransferserversResource, awstransferserversKind, opts), &v1alpha1.AwsTransferServerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsTransferServerList{ListMeta: obj.(*v1alpha1.AwsTransferServerList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsTransferServerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsTransferServers.
func (c *FakeAwsTransferServers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awstransferserversResource, opts))
}

// Create takes the representation of a awsTransferServer and creates it.  Returns the server's representation of the awsTransferServer, and an error, if there is any.
func (c *FakeAwsTransferServers) Create(awsTransferServer *v1alpha1.AwsTransferServer) (result *v1alpha1.AwsTransferServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awstransferserversResource, awsTransferServer), &v1alpha1.AwsTransferServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsTransferServer), err
}

// Update takes the representation of a awsTransferServer and updates it. Returns the server's representation of the awsTransferServer, and an error, if there is any.
func (c *FakeAwsTransferServers) Update(awsTransferServer *v1alpha1.AwsTransferServer) (result *v1alpha1.AwsTransferServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awstransferserversResource, awsTransferServer), &v1alpha1.AwsTransferServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsTransferServer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsTransferServers) UpdateStatus(awsTransferServer *v1alpha1.AwsTransferServer) (*v1alpha1.AwsTransferServer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awstransferserversResource, "status", awsTransferServer), &v1alpha1.AwsTransferServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsTransferServer), err
}

// Delete takes name of the awsTransferServer and deletes it. Returns an error if one occurs.
func (c *FakeAwsTransferServers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awstransferserversResource, name), &v1alpha1.AwsTransferServer{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsTransferServers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awstransferserversResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsTransferServerList{})
	return err
}

// Patch applies the patch and returns the patched awsTransferServer.
func (c *FakeAwsTransferServers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsTransferServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awstransferserversResource, name, pt, data, subresources...), &v1alpha1.AwsTransferServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsTransferServer), err
}