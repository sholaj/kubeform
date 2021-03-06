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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LightsailKeyPairLister helps list LightsailKeyPairs.
type LightsailKeyPairLister interface {
	// List lists all LightsailKeyPairs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.LightsailKeyPair, err error)
	// LightsailKeyPairs returns an object that can list and get LightsailKeyPairs.
	LightsailKeyPairs(namespace string) LightsailKeyPairNamespaceLister
	LightsailKeyPairListerExpansion
}

// lightsailKeyPairLister implements the LightsailKeyPairLister interface.
type lightsailKeyPairLister struct {
	indexer cache.Indexer
}

// NewLightsailKeyPairLister returns a new LightsailKeyPairLister.
func NewLightsailKeyPairLister(indexer cache.Indexer) LightsailKeyPairLister {
	return &lightsailKeyPairLister{indexer: indexer}
}

// List lists all LightsailKeyPairs in the indexer.
func (s *lightsailKeyPairLister) List(selector labels.Selector) (ret []*v1alpha1.LightsailKeyPair, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LightsailKeyPair))
	})
	return ret, err
}

// LightsailKeyPairs returns an object that can list and get LightsailKeyPairs.
func (s *lightsailKeyPairLister) LightsailKeyPairs(namespace string) LightsailKeyPairNamespaceLister {
	return lightsailKeyPairNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LightsailKeyPairNamespaceLister helps list and get LightsailKeyPairs.
type LightsailKeyPairNamespaceLister interface {
	// List lists all LightsailKeyPairs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.LightsailKeyPair, err error)
	// Get retrieves the LightsailKeyPair from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.LightsailKeyPair, error)
	LightsailKeyPairNamespaceListerExpansion
}

// lightsailKeyPairNamespaceLister implements the LightsailKeyPairNamespaceLister
// interface.
type lightsailKeyPairNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LightsailKeyPairs in the indexer for a given namespace.
func (s lightsailKeyPairNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LightsailKeyPair, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LightsailKeyPair))
	})
	return ret, err
}

// Get retrieves the LightsailKeyPair from the indexer for a given namespace and name.
func (s lightsailKeyPairNamespaceLister) Get(name string) (*v1alpha1.LightsailKeyPair, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("lightsailkeypair"), name)
	}
	return obj.(*v1alpha1.LightsailKeyPair), nil
}
