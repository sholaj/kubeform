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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeGoogleDataprocJobs implements GoogleDataprocJobInterface
type FakeGoogleDataprocJobs struct {
	Fake *FakeGoogleV1alpha1
}

var googledataprocjobsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googledataprocjobs"}

var googledataprocjobsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleDataprocJob"}

// Get takes name of the googleDataprocJob, and returns the corresponding googleDataprocJob object, and an error if there is any.
func (c *FakeGoogleDataprocJobs) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleDataprocJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googledataprocjobsResource, name), &v1alpha1.GoogleDataprocJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleDataprocJob), err
}

// List takes label and field selectors, and returns the list of GoogleDataprocJobs that match those selectors.
func (c *FakeGoogleDataprocJobs) List(opts v1.ListOptions) (result *v1alpha1.GoogleDataprocJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googledataprocjobsResource, googledataprocjobsKind, opts), &v1alpha1.GoogleDataprocJobList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleDataprocJobList{ListMeta: obj.(*v1alpha1.GoogleDataprocJobList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleDataprocJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleDataprocJobs.
func (c *FakeGoogleDataprocJobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googledataprocjobsResource, opts))
}

// Create takes the representation of a googleDataprocJob and creates it.  Returns the server's representation of the googleDataprocJob, and an error, if there is any.
func (c *FakeGoogleDataprocJobs) Create(googleDataprocJob *v1alpha1.GoogleDataprocJob) (result *v1alpha1.GoogleDataprocJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googledataprocjobsResource, googleDataprocJob), &v1alpha1.GoogleDataprocJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleDataprocJob), err
}

// Update takes the representation of a googleDataprocJob and updates it. Returns the server's representation of the googleDataprocJob, and an error, if there is any.
func (c *FakeGoogleDataprocJobs) Update(googleDataprocJob *v1alpha1.GoogleDataprocJob) (result *v1alpha1.GoogleDataprocJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googledataprocjobsResource, googleDataprocJob), &v1alpha1.GoogleDataprocJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleDataprocJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleDataprocJobs) UpdateStatus(googleDataprocJob *v1alpha1.GoogleDataprocJob) (*v1alpha1.GoogleDataprocJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googledataprocjobsResource, "status", googleDataprocJob), &v1alpha1.GoogleDataprocJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleDataprocJob), err
}

// Delete takes name of the googleDataprocJob and deletes it. Returns an error if one occurs.
func (c *FakeGoogleDataprocJobs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googledataprocjobsResource, name), &v1alpha1.GoogleDataprocJob{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleDataprocJobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googledataprocjobsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleDataprocJobList{})
	return err
}

// Patch applies the patch and returns the patched googleDataprocJob.
func (c *FakeGoogleDataprocJobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleDataprocJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googledataprocjobsResource, name, pt, data, subresources...), &v1alpha1.GoogleDataprocJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleDataprocJob), err
}