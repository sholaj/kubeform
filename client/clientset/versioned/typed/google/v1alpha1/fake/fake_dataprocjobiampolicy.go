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

// FakeDataprocJobIamPolicies implements DataprocJobIamPolicyInterface
type FakeDataprocJobIamPolicies struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var dataprocjobiampoliciesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "dataprocjobiampolicies"}

var dataprocjobiampoliciesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "DataprocJobIamPolicy"}

// Get takes name of the dataprocJobIamPolicy, and returns the corresponding dataprocJobIamPolicy object, and an error if there is any.
func (c *FakeDataprocJobIamPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DataprocJobIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dataprocjobiampoliciesResource, c.ns, name), &v1alpha1.DataprocJobIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataprocJobIamPolicy), err
}

// List takes label and field selectors, and returns the list of DataprocJobIamPolicies that match those selectors.
func (c *FakeDataprocJobIamPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DataprocJobIamPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dataprocjobiampoliciesResource, dataprocjobiampoliciesKind, c.ns, opts), &v1alpha1.DataprocJobIamPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DataprocJobIamPolicyList{ListMeta: obj.(*v1alpha1.DataprocJobIamPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.DataprocJobIamPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataprocJobIamPolicies.
func (c *FakeDataprocJobIamPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dataprocjobiampoliciesResource, c.ns, opts))

}

// Create takes the representation of a dataprocJobIamPolicy and creates it.  Returns the server's representation of the dataprocJobIamPolicy, and an error, if there is any.
func (c *FakeDataprocJobIamPolicies) Create(ctx context.Context, dataprocJobIamPolicy *v1alpha1.DataprocJobIamPolicy, opts v1.CreateOptions) (result *v1alpha1.DataprocJobIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dataprocjobiampoliciesResource, c.ns, dataprocJobIamPolicy), &v1alpha1.DataprocJobIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataprocJobIamPolicy), err
}

// Update takes the representation of a dataprocJobIamPolicy and updates it. Returns the server's representation of the dataprocJobIamPolicy, and an error, if there is any.
func (c *FakeDataprocJobIamPolicies) Update(ctx context.Context, dataprocJobIamPolicy *v1alpha1.DataprocJobIamPolicy, opts v1.UpdateOptions) (result *v1alpha1.DataprocJobIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dataprocjobiampoliciesResource, c.ns, dataprocJobIamPolicy), &v1alpha1.DataprocJobIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataprocJobIamPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDataprocJobIamPolicies) UpdateStatus(ctx context.Context, dataprocJobIamPolicy *v1alpha1.DataprocJobIamPolicy, opts v1.UpdateOptions) (*v1alpha1.DataprocJobIamPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dataprocjobiampoliciesResource, "status", c.ns, dataprocJobIamPolicy), &v1alpha1.DataprocJobIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataprocJobIamPolicy), err
}

// Delete takes name of the dataprocJobIamPolicy and deletes it. Returns an error if one occurs.
func (c *FakeDataprocJobIamPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dataprocjobiampoliciesResource, c.ns, name), &v1alpha1.DataprocJobIamPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataprocJobIamPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dataprocjobiampoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DataprocJobIamPolicyList{})
	return err
}

// Patch applies the patch and returns the patched dataprocJobIamPolicy.
func (c *FakeDataprocJobIamPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DataprocJobIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dataprocjobiampoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.DataprocJobIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataprocJobIamPolicy), err
}
