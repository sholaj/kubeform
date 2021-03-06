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

// DefaultVpcLister helps list DefaultVpcs.
type DefaultVpcLister interface {
	// List lists all DefaultVpcs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DefaultVpc, err error)
	// DefaultVpcs returns an object that can list and get DefaultVpcs.
	DefaultVpcs(namespace string) DefaultVpcNamespaceLister
	DefaultVpcListerExpansion
}

// defaultVpcLister implements the DefaultVpcLister interface.
type defaultVpcLister struct {
	indexer cache.Indexer
}

// NewDefaultVpcLister returns a new DefaultVpcLister.
func NewDefaultVpcLister(indexer cache.Indexer) DefaultVpcLister {
	return &defaultVpcLister{indexer: indexer}
}

// List lists all DefaultVpcs in the indexer.
func (s *defaultVpcLister) List(selector labels.Selector) (ret []*v1alpha1.DefaultVpc, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DefaultVpc))
	})
	return ret, err
}

// DefaultVpcs returns an object that can list and get DefaultVpcs.
func (s *defaultVpcLister) DefaultVpcs(namespace string) DefaultVpcNamespaceLister {
	return defaultVpcNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DefaultVpcNamespaceLister helps list and get DefaultVpcs.
type DefaultVpcNamespaceLister interface {
	// List lists all DefaultVpcs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DefaultVpc, err error)
	// Get retrieves the DefaultVpc from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DefaultVpc, error)
	DefaultVpcNamespaceListerExpansion
}

// defaultVpcNamespaceLister implements the DefaultVpcNamespaceLister
// interface.
type defaultVpcNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DefaultVpcs in the indexer for a given namespace.
func (s defaultVpcNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DefaultVpc, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DefaultVpc))
	})
	return ret, err
}

// Get retrieves the DefaultVpc from the indexer for a given namespace and name.
func (s defaultVpcNamespaceLister) Get(name string) (*v1alpha1.DefaultVpc, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("defaultvpc"), name)
	}
	return obj.(*v1alpha1.DefaultVpc), nil
}
