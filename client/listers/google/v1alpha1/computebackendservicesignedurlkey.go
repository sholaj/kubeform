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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ComputeBackendServiceSignedURLKeyLister helps list ComputeBackendServiceSignedURLKeys.
type ComputeBackendServiceSignedURLKeyLister interface {
	// List lists all ComputeBackendServiceSignedURLKeys in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeBackendServiceSignedURLKey, err error)
	// ComputeBackendServiceSignedURLKeys returns an object that can list and get ComputeBackendServiceSignedURLKeys.
	ComputeBackendServiceSignedURLKeys(namespace string) ComputeBackendServiceSignedURLKeyNamespaceLister
	ComputeBackendServiceSignedURLKeyListerExpansion
}

// computeBackendServiceSignedURLKeyLister implements the ComputeBackendServiceSignedURLKeyLister interface.
type computeBackendServiceSignedURLKeyLister struct {
	indexer cache.Indexer
}

// NewComputeBackendServiceSignedURLKeyLister returns a new ComputeBackendServiceSignedURLKeyLister.
func NewComputeBackendServiceSignedURLKeyLister(indexer cache.Indexer) ComputeBackendServiceSignedURLKeyLister {
	return &computeBackendServiceSignedURLKeyLister{indexer: indexer}
}

// List lists all ComputeBackendServiceSignedURLKeys in the indexer.
func (s *computeBackendServiceSignedURLKeyLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeBackendServiceSignedURLKey, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeBackendServiceSignedURLKey))
	})
	return ret, err
}

// ComputeBackendServiceSignedURLKeys returns an object that can list and get ComputeBackendServiceSignedURLKeys.
func (s *computeBackendServiceSignedURLKeyLister) ComputeBackendServiceSignedURLKeys(namespace string) ComputeBackendServiceSignedURLKeyNamespaceLister {
	return computeBackendServiceSignedURLKeyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ComputeBackendServiceSignedURLKeyNamespaceLister helps list and get ComputeBackendServiceSignedURLKeys.
type ComputeBackendServiceSignedURLKeyNamespaceLister interface {
	// List lists all ComputeBackendServiceSignedURLKeys in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeBackendServiceSignedURLKey, err error)
	// Get retrieves the ComputeBackendServiceSignedURLKey from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ComputeBackendServiceSignedURLKey, error)
	ComputeBackendServiceSignedURLKeyNamespaceListerExpansion
}

// computeBackendServiceSignedURLKeyNamespaceLister implements the ComputeBackendServiceSignedURLKeyNamespaceLister
// interface.
type computeBackendServiceSignedURLKeyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ComputeBackendServiceSignedURLKeys in the indexer for a given namespace.
func (s computeBackendServiceSignedURLKeyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeBackendServiceSignedURLKey, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeBackendServiceSignedURLKey))
	})
	return ret, err
}

// Get retrieves the ComputeBackendServiceSignedURLKey from the indexer for a given namespace and name.
func (s computeBackendServiceSignedURLKeyNamespaceLister) Get(name string) (*v1alpha1.ComputeBackendServiceSignedURLKey, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("computebackendservicesignedurlkey"), name)
	}
	return obj.(*v1alpha1.ComputeBackendServiceSignedURLKey), nil
}
