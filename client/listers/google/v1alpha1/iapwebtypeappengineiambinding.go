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

// IapWebTypeAppEngineIamBindingLister helps list IapWebTypeAppEngineIamBindings.
type IapWebTypeAppEngineIamBindingLister interface {
	// List lists all IapWebTypeAppEngineIamBindings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.IapWebTypeAppEngineIamBinding, err error)
	// IapWebTypeAppEngineIamBindings returns an object that can list and get IapWebTypeAppEngineIamBindings.
	IapWebTypeAppEngineIamBindings(namespace string) IapWebTypeAppEngineIamBindingNamespaceLister
	IapWebTypeAppEngineIamBindingListerExpansion
}

// iapWebTypeAppEngineIamBindingLister implements the IapWebTypeAppEngineIamBindingLister interface.
type iapWebTypeAppEngineIamBindingLister struct {
	indexer cache.Indexer
}

// NewIapWebTypeAppEngineIamBindingLister returns a new IapWebTypeAppEngineIamBindingLister.
func NewIapWebTypeAppEngineIamBindingLister(indexer cache.Indexer) IapWebTypeAppEngineIamBindingLister {
	return &iapWebTypeAppEngineIamBindingLister{indexer: indexer}
}

// List lists all IapWebTypeAppEngineIamBindings in the indexer.
func (s *iapWebTypeAppEngineIamBindingLister) List(selector labels.Selector) (ret []*v1alpha1.IapWebTypeAppEngineIamBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IapWebTypeAppEngineIamBinding))
	})
	return ret, err
}

// IapWebTypeAppEngineIamBindings returns an object that can list and get IapWebTypeAppEngineIamBindings.
func (s *iapWebTypeAppEngineIamBindingLister) IapWebTypeAppEngineIamBindings(namespace string) IapWebTypeAppEngineIamBindingNamespaceLister {
	return iapWebTypeAppEngineIamBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IapWebTypeAppEngineIamBindingNamespaceLister helps list and get IapWebTypeAppEngineIamBindings.
type IapWebTypeAppEngineIamBindingNamespaceLister interface {
	// List lists all IapWebTypeAppEngineIamBindings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.IapWebTypeAppEngineIamBinding, err error)
	// Get retrieves the IapWebTypeAppEngineIamBinding from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.IapWebTypeAppEngineIamBinding, error)
	IapWebTypeAppEngineIamBindingNamespaceListerExpansion
}

// iapWebTypeAppEngineIamBindingNamespaceLister implements the IapWebTypeAppEngineIamBindingNamespaceLister
// interface.
type iapWebTypeAppEngineIamBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IapWebTypeAppEngineIamBindings in the indexer for a given namespace.
func (s iapWebTypeAppEngineIamBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IapWebTypeAppEngineIamBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IapWebTypeAppEngineIamBinding))
	})
	return ret, err
}

// Get retrieves the IapWebTypeAppEngineIamBinding from the indexer for a given namespace and name.
func (s iapWebTypeAppEngineIamBindingNamespaceLister) Get(name string) (*v1alpha1.IapWebTypeAppEngineIamBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iapwebtypeappengineiambinding"), name)
	}
	return obj.(*v1alpha1.IapWebTypeAppEngineIamBinding), nil
}
