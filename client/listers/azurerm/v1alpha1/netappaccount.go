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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NetappAccountLister helps list NetappAccounts.
type NetappAccountLister interface {
	// List lists all NetappAccounts in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.NetappAccount, err error)
	// NetappAccounts returns an object that can list and get NetappAccounts.
	NetappAccounts(namespace string) NetappAccountNamespaceLister
	NetappAccountListerExpansion
}

// netappAccountLister implements the NetappAccountLister interface.
type netappAccountLister struct {
	indexer cache.Indexer
}

// NewNetappAccountLister returns a new NetappAccountLister.
func NewNetappAccountLister(indexer cache.Indexer) NetappAccountLister {
	return &netappAccountLister{indexer: indexer}
}

// List lists all NetappAccounts in the indexer.
func (s *netappAccountLister) List(selector labels.Selector) (ret []*v1alpha1.NetappAccount, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetappAccount))
	})
	return ret, err
}

// NetappAccounts returns an object that can list and get NetappAccounts.
func (s *netappAccountLister) NetappAccounts(namespace string) NetappAccountNamespaceLister {
	return netappAccountNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NetappAccountNamespaceLister helps list and get NetappAccounts.
type NetappAccountNamespaceLister interface {
	// List lists all NetappAccounts in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.NetappAccount, err error)
	// Get retrieves the NetappAccount from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.NetappAccount, error)
	NetappAccountNamespaceListerExpansion
}

// netappAccountNamespaceLister implements the NetappAccountNamespaceLister
// interface.
type netappAccountNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NetappAccounts in the indexer for a given namespace.
func (s netappAccountNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NetappAccount, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetappAccount))
	})
	return ret, err
}

// Get retrieves the NetappAccount from the indexer for a given namespace and name.
func (s netappAccountNamespaceLister) Get(name string) (*v1alpha1.NetappAccount, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("netappaccount"), name)
	}
	return obj.(*v1alpha1.NetappAccount), nil
}