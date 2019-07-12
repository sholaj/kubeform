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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// GoogleStorageBucketLister helps list GoogleStorageBuckets.
type GoogleStorageBucketLister interface {
	// List lists all GoogleStorageBuckets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.GoogleStorageBucket, err error)
	// Get retrieves the GoogleStorageBucket from the index for a given name.
	Get(name string) (*v1alpha1.GoogleStorageBucket, error)
	GoogleStorageBucketListerExpansion
}

// googleStorageBucketLister implements the GoogleStorageBucketLister interface.
type googleStorageBucketLister struct {
	indexer cache.Indexer
}

// NewGoogleStorageBucketLister returns a new GoogleStorageBucketLister.
func NewGoogleStorageBucketLister(indexer cache.Indexer) GoogleStorageBucketLister {
	return &googleStorageBucketLister{indexer: indexer}
}

// List lists all GoogleStorageBuckets in the indexer.
func (s *googleStorageBucketLister) List(selector labels.Selector) (ret []*v1alpha1.GoogleStorageBucket, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GoogleStorageBucket))
	})
	return ret, err
}

// Get retrieves the GoogleStorageBucket from the index for a given name.
func (s *googleStorageBucketLister) Get(name string) (*v1alpha1.GoogleStorageBucket, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("googlestoragebucket"), name)
	}
	return obj.(*v1alpha1.GoogleStorageBucket), nil
}