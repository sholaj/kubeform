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
	v1alpha1 "kubeform.dev/kubeform/apis/linode/v1alpha1"
)

// FakeLinodeDomainRecords implements LinodeDomainRecordInterface
type FakeLinodeDomainRecords struct {
	Fake *FakeLinodeV1alpha1
}

var linodedomainrecordsResource = schema.GroupVersionResource{Group: "linode.kubeform.com", Version: "v1alpha1", Resource: "linodedomainrecords"}

var linodedomainrecordsKind = schema.GroupVersionKind{Group: "linode.kubeform.com", Version: "v1alpha1", Kind: "LinodeDomainRecord"}

// Get takes name of the linodeDomainRecord, and returns the corresponding linodeDomainRecord object, and an error if there is any.
func (c *FakeLinodeDomainRecords) Get(name string, options v1.GetOptions) (result *v1alpha1.LinodeDomainRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(linodedomainrecordsResource, name), &v1alpha1.LinodeDomainRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinodeDomainRecord), err
}

// List takes label and field selectors, and returns the list of LinodeDomainRecords that match those selectors.
func (c *FakeLinodeDomainRecords) List(opts v1.ListOptions) (result *v1alpha1.LinodeDomainRecordList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(linodedomainrecordsResource, linodedomainrecordsKind, opts), &v1alpha1.LinodeDomainRecordList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LinodeDomainRecordList{ListMeta: obj.(*v1alpha1.LinodeDomainRecordList).ListMeta}
	for _, item := range obj.(*v1alpha1.LinodeDomainRecordList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested linodeDomainRecords.
func (c *FakeLinodeDomainRecords) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(linodedomainrecordsResource, opts))
}

// Create takes the representation of a linodeDomainRecord and creates it.  Returns the server's representation of the linodeDomainRecord, and an error, if there is any.
func (c *FakeLinodeDomainRecords) Create(linodeDomainRecord *v1alpha1.LinodeDomainRecord) (result *v1alpha1.LinodeDomainRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(linodedomainrecordsResource, linodeDomainRecord), &v1alpha1.LinodeDomainRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinodeDomainRecord), err
}

// Update takes the representation of a linodeDomainRecord and updates it. Returns the server's representation of the linodeDomainRecord, and an error, if there is any.
func (c *FakeLinodeDomainRecords) Update(linodeDomainRecord *v1alpha1.LinodeDomainRecord) (result *v1alpha1.LinodeDomainRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(linodedomainrecordsResource, linodeDomainRecord), &v1alpha1.LinodeDomainRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinodeDomainRecord), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLinodeDomainRecords) UpdateStatus(linodeDomainRecord *v1alpha1.LinodeDomainRecord) (*v1alpha1.LinodeDomainRecord, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(linodedomainrecordsResource, "status", linodeDomainRecord), &v1alpha1.LinodeDomainRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinodeDomainRecord), err
}

// Delete takes name of the linodeDomainRecord and deletes it. Returns an error if one occurs.
func (c *FakeLinodeDomainRecords) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(linodedomainrecordsResource, name), &v1alpha1.LinodeDomainRecord{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLinodeDomainRecords) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(linodedomainrecordsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LinodeDomainRecordList{})
	return err
}

// Patch applies the patch and returns the patched linodeDomainRecord.
func (c *FakeLinodeDomainRecords) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LinodeDomainRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(linodedomainrecordsResource, name, pt, data, subresources...), &v1alpha1.LinodeDomainRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinodeDomainRecord), err
}