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

// FakeAwsElastictranscoderPipelines implements AwsElastictranscoderPipelineInterface
type FakeAwsElastictranscoderPipelines struct {
	Fake *FakeAwsV1alpha1
}

var awselastictranscoderpipelinesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awselastictranscoderpipelines"}

var awselastictranscoderpipelinesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsElastictranscoderPipeline"}

// Get takes name of the awsElastictranscoderPipeline, and returns the corresponding awsElastictranscoderPipeline object, and an error if there is any.
func (c *FakeAwsElastictranscoderPipelines) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsElastictranscoderPipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awselastictranscoderpipelinesResource, name), &v1alpha1.AwsElastictranscoderPipeline{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElastictranscoderPipeline), err
}

// List takes label and field selectors, and returns the list of AwsElastictranscoderPipelines that match those selectors.
func (c *FakeAwsElastictranscoderPipelines) List(opts v1.ListOptions) (result *v1alpha1.AwsElastictranscoderPipelineList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awselastictranscoderpipelinesResource, awselastictranscoderpipelinesKind, opts), &v1alpha1.AwsElastictranscoderPipelineList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsElastictranscoderPipelineList{ListMeta: obj.(*v1alpha1.AwsElastictranscoderPipelineList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsElastictranscoderPipelineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsElastictranscoderPipelines.
func (c *FakeAwsElastictranscoderPipelines) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awselastictranscoderpipelinesResource, opts))
}

// Create takes the representation of a awsElastictranscoderPipeline and creates it.  Returns the server's representation of the awsElastictranscoderPipeline, and an error, if there is any.
func (c *FakeAwsElastictranscoderPipelines) Create(awsElastictranscoderPipeline *v1alpha1.AwsElastictranscoderPipeline) (result *v1alpha1.AwsElastictranscoderPipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awselastictranscoderpipelinesResource, awsElastictranscoderPipeline), &v1alpha1.AwsElastictranscoderPipeline{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElastictranscoderPipeline), err
}

// Update takes the representation of a awsElastictranscoderPipeline and updates it. Returns the server's representation of the awsElastictranscoderPipeline, and an error, if there is any.
func (c *FakeAwsElastictranscoderPipelines) Update(awsElastictranscoderPipeline *v1alpha1.AwsElastictranscoderPipeline) (result *v1alpha1.AwsElastictranscoderPipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awselastictranscoderpipelinesResource, awsElastictranscoderPipeline), &v1alpha1.AwsElastictranscoderPipeline{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElastictranscoderPipeline), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsElastictranscoderPipelines) UpdateStatus(awsElastictranscoderPipeline *v1alpha1.AwsElastictranscoderPipeline) (*v1alpha1.AwsElastictranscoderPipeline, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awselastictranscoderpipelinesResource, "status", awsElastictranscoderPipeline), &v1alpha1.AwsElastictranscoderPipeline{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElastictranscoderPipeline), err
}

// Delete takes name of the awsElastictranscoderPipeline and deletes it. Returns an error if one occurs.
func (c *FakeAwsElastictranscoderPipelines) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awselastictranscoderpipelinesResource, name), &v1alpha1.AwsElastictranscoderPipeline{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsElastictranscoderPipelines) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awselastictranscoderpipelinesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsElastictranscoderPipelineList{})
	return err
}

// Patch applies the patch and returns the patched awsElastictranscoderPipeline.
func (c *FakeAwsElastictranscoderPipelines) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsElastictranscoderPipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awselastictranscoderpipelinesResource, name, pt, data, subresources...), &v1alpha1.AwsElastictranscoderPipeline{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElastictranscoderPipeline), err
}