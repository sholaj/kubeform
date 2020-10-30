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

// ComputeProjectDefaultNetworkTierLister helps list ComputeProjectDefaultNetworkTiers.
type ComputeProjectDefaultNetworkTierLister interface {
	// List lists all ComputeProjectDefaultNetworkTiers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeProjectDefaultNetworkTier, err error)
	// ComputeProjectDefaultNetworkTiers returns an object that can list and get ComputeProjectDefaultNetworkTiers.
	ComputeProjectDefaultNetworkTiers(namespace string) ComputeProjectDefaultNetworkTierNamespaceLister
	ComputeProjectDefaultNetworkTierListerExpansion
}

// computeProjectDefaultNetworkTierLister implements the ComputeProjectDefaultNetworkTierLister interface.
type computeProjectDefaultNetworkTierLister struct {
	indexer cache.Indexer
}

// NewComputeProjectDefaultNetworkTierLister returns a new ComputeProjectDefaultNetworkTierLister.
func NewComputeProjectDefaultNetworkTierLister(indexer cache.Indexer) ComputeProjectDefaultNetworkTierLister {
	return &computeProjectDefaultNetworkTierLister{indexer: indexer}
}

// List lists all ComputeProjectDefaultNetworkTiers in the indexer.
func (s *computeProjectDefaultNetworkTierLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeProjectDefaultNetworkTier, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeProjectDefaultNetworkTier))
	})
	return ret, err
}

// ComputeProjectDefaultNetworkTiers returns an object that can list and get ComputeProjectDefaultNetworkTiers.
func (s *computeProjectDefaultNetworkTierLister) ComputeProjectDefaultNetworkTiers(namespace string) ComputeProjectDefaultNetworkTierNamespaceLister {
	return computeProjectDefaultNetworkTierNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ComputeProjectDefaultNetworkTierNamespaceLister helps list and get ComputeProjectDefaultNetworkTiers.
type ComputeProjectDefaultNetworkTierNamespaceLister interface {
	// List lists all ComputeProjectDefaultNetworkTiers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeProjectDefaultNetworkTier, err error)
	// Get retrieves the ComputeProjectDefaultNetworkTier from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ComputeProjectDefaultNetworkTier, error)
	ComputeProjectDefaultNetworkTierNamespaceListerExpansion
}

// computeProjectDefaultNetworkTierNamespaceLister implements the ComputeProjectDefaultNetworkTierNamespaceLister
// interface.
type computeProjectDefaultNetworkTierNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ComputeProjectDefaultNetworkTiers in the indexer for a given namespace.
func (s computeProjectDefaultNetworkTierNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeProjectDefaultNetworkTier, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeProjectDefaultNetworkTier))
	})
	return ret, err
}

// Get retrieves the ComputeProjectDefaultNetworkTier from the indexer for a given namespace and name.
func (s computeProjectDefaultNetworkTierNamespaceLister) Get(name string) (*v1alpha1.ComputeProjectDefaultNetworkTier, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("computeprojectdefaultnetworktier"), name)
	}
	return obj.(*v1alpha1.ComputeProjectDefaultNetworkTier), nil
}