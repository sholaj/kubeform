/*
Copyright The Kubeform Authors.

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
	"context"

	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBigqueryDataTransferConfigs implements BigqueryDataTransferConfigInterface
type FakeBigqueryDataTransferConfigs struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var bigquerydatatransferconfigsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "bigquerydatatransferconfigs"}

var bigquerydatatransferconfigsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "BigqueryDataTransferConfig"}

// Get takes name of the bigqueryDataTransferConfig, and returns the corresponding bigqueryDataTransferConfig object, and an error if there is any.
func (c *FakeBigqueryDataTransferConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BigqueryDataTransferConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(bigquerydatatransferconfigsResource, c.ns, name), &v1alpha1.BigqueryDataTransferConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BigqueryDataTransferConfig), err
}

// List takes label and field selectors, and returns the list of BigqueryDataTransferConfigs that match those selectors.
func (c *FakeBigqueryDataTransferConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BigqueryDataTransferConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(bigquerydatatransferconfigsResource, bigquerydatatransferconfigsKind, c.ns, opts), &v1alpha1.BigqueryDataTransferConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BigqueryDataTransferConfigList{ListMeta: obj.(*v1alpha1.BigqueryDataTransferConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.BigqueryDataTransferConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bigqueryDataTransferConfigs.
func (c *FakeBigqueryDataTransferConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(bigquerydatatransferconfigsResource, c.ns, opts))

}

// Create takes the representation of a bigqueryDataTransferConfig and creates it.  Returns the server's representation of the bigqueryDataTransferConfig, and an error, if there is any.
func (c *FakeBigqueryDataTransferConfigs) Create(ctx context.Context, bigqueryDataTransferConfig *v1alpha1.BigqueryDataTransferConfig, opts v1.CreateOptions) (result *v1alpha1.BigqueryDataTransferConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(bigquerydatatransferconfigsResource, c.ns, bigqueryDataTransferConfig), &v1alpha1.BigqueryDataTransferConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BigqueryDataTransferConfig), err
}

// Update takes the representation of a bigqueryDataTransferConfig and updates it. Returns the server's representation of the bigqueryDataTransferConfig, and an error, if there is any.
func (c *FakeBigqueryDataTransferConfigs) Update(ctx context.Context, bigqueryDataTransferConfig *v1alpha1.BigqueryDataTransferConfig, opts v1.UpdateOptions) (result *v1alpha1.BigqueryDataTransferConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(bigquerydatatransferconfigsResource, c.ns, bigqueryDataTransferConfig), &v1alpha1.BigqueryDataTransferConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BigqueryDataTransferConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBigqueryDataTransferConfigs) UpdateStatus(ctx context.Context, bigqueryDataTransferConfig *v1alpha1.BigqueryDataTransferConfig, opts v1.UpdateOptions) (*v1alpha1.BigqueryDataTransferConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(bigquerydatatransferconfigsResource, "status", c.ns, bigqueryDataTransferConfig), &v1alpha1.BigqueryDataTransferConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BigqueryDataTransferConfig), err
}

// Delete takes name of the bigqueryDataTransferConfig and deletes it. Returns an error if one occurs.
func (c *FakeBigqueryDataTransferConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(bigquerydatatransferconfigsResource, c.ns, name), &v1alpha1.BigqueryDataTransferConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBigqueryDataTransferConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(bigquerydatatransferconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BigqueryDataTransferConfigList{})
	return err
}

// Patch applies the patch and returns the patched bigqueryDataTransferConfig.
func (c *FakeBigqueryDataTransferConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BigqueryDataTransferConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(bigquerydatatransferconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.BigqueryDataTransferConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BigqueryDataTransferConfig), err
}
