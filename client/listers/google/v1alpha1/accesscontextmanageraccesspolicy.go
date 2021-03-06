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

// AccessContextManagerAccessPolicyLister helps list AccessContextManagerAccessPolicies.
type AccessContextManagerAccessPolicyLister interface {
	// List lists all AccessContextManagerAccessPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AccessContextManagerAccessPolicy, err error)
	// AccessContextManagerAccessPolicies returns an object that can list and get AccessContextManagerAccessPolicies.
	AccessContextManagerAccessPolicies(namespace string) AccessContextManagerAccessPolicyNamespaceLister
	AccessContextManagerAccessPolicyListerExpansion
}

// accessContextManagerAccessPolicyLister implements the AccessContextManagerAccessPolicyLister interface.
type accessContextManagerAccessPolicyLister struct {
	indexer cache.Indexer
}

// NewAccessContextManagerAccessPolicyLister returns a new AccessContextManagerAccessPolicyLister.
func NewAccessContextManagerAccessPolicyLister(indexer cache.Indexer) AccessContextManagerAccessPolicyLister {
	return &accessContextManagerAccessPolicyLister{indexer: indexer}
}

// List lists all AccessContextManagerAccessPolicies in the indexer.
func (s *accessContextManagerAccessPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.AccessContextManagerAccessPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AccessContextManagerAccessPolicy))
	})
	return ret, err
}

// AccessContextManagerAccessPolicies returns an object that can list and get AccessContextManagerAccessPolicies.
func (s *accessContextManagerAccessPolicyLister) AccessContextManagerAccessPolicies(namespace string) AccessContextManagerAccessPolicyNamespaceLister {
	return accessContextManagerAccessPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AccessContextManagerAccessPolicyNamespaceLister helps list and get AccessContextManagerAccessPolicies.
type AccessContextManagerAccessPolicyNamespaceLister interface {
	// List lists all AccessContextManagerAccessPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AccessContextManagerAccessPolicy, err error)
	// Get retrieves the AccessContextManagerAccessPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AccessContextManagerAccessPolicy, error)
	AccessContextManagerAccessPolicyNamespaceListerExpansion
}

// accessContextManagerAccessPolicyNamespaceLister implements the AccessContextManagerAccessPolicyNamespaceLister
// interface.
type accessContextManagerAccessPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AccessContextManagerAccessPolicies in the indexer for a given namespace.
func (s accessContextManagerAccessPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AccessContextManagerAccessPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AccessContextManagerAccessPolicy))
	})
	return ret, err
}

// Get retrieves the AccessContextManagerAccessPolicy from the indexer for a given namespace and name.
func (s accessContextManagerAccessPolicyNamespaceLister) Get(name string) (*v1alpha1.AccessContextManagerAccessPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("accesscontextmanageraccesspolicy"), name)
	}
	return obj.(*v1alpha1.AccessContextManagerAccessPolicy), nil
}
