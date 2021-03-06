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

// FmsAdminAccountLister helps list FmsAdminAccounts.
type FmsAdminAccountLister interface {
	// List lists all FmsAdminAccounts in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.FmsAdminAccount, err error)
	// FmsAdminAccounts returns an object that can list and get FmsAdminAccounts.
	FmsAdminAccounts(namespace string) FmsAdminAccountNamespaceLister
	FmsAdminAccountListerExpansion
}

// fmsAdminAccountLister implements the FmsAdminAccountLister interface.
type fmsAdminAccountLister struct {
	indexer cache.Indexer
}

// NewFmsAdminAccountLister returns a new FmsAdminAccountLister.
func NewFmsAdminAccountLister(indexer cache.Indexer) FmsAdminAccountLister {
	return &fmsAdminAccountLister{indexer: indexer}
}

// List lists all FmsAdminAccounts in the indexer.
func (s *fmsAdminAccountLister) List(selector labels.Selector) (ret []*v1alpha1.FmsAdminAccount, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FmsAdminAccount))
	})
	return ret, err
}

// FmsAdminAccounts returns an object that can list and get FmsAdminAccounts.
func (s *fmsAdminAccountLister) FmsAdminAccounts(namespace string) FmsAdminAccountNamespaceLister {
	return fmsAdminAccountNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FmsAdminAccountNamespaceLister helps list and get FmsAdminAccounts.
type FmsAdminAccountNamespaceLister interface {
	// List lists all FmsAdminAccounts in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.FmsAdminAccount, err error)
	// Get retrieves the FmsAdminAccount from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.FmsAdminAccount, error)
	FmsAdminAccountNamespaceListerExpansion
}

// fmsAdminAccountNamespaceLister implements the FmsAdminAccountNamespaceLister
// interface.
type fmsAdminAccountNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FmsAdminAccounts in the indexer for a given namespace.
func (s fmsAdminAccountNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FmsAdminAccount, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FmsAdminAccount))
	})
	return ret, err
}

// Get retrieves the FmsAdminAccount from the indexer for a given namespace and name.
func (s fmsAdminAccountNamespaceLister) Get(name string) (*v1alpha1.FmsAdminAccount, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("fmsadminaccount"), name)
	}
	return obj.(*v1alpha1.FmsAdminAccount), nil
}
