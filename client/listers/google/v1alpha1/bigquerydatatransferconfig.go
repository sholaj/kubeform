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

// BigqueryDataTransferConfigLister helps list BigqueryDataTransferConfigs.
type BigqueryDataTransferConfigLister interface {
	// List lists all BigqueryDataTransferConfigs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.BigqueryDataTransferConfig, err error)
	// BigqueryDataTransferConfigs returns an object that can list and get BigqueryDataTransferConfigs.
	BigqueryDataTransferConfigs(namespace string) BigqueryDataTransferConfigNamespaceLister
	BigqueryDataTransferConfigListerExpansion
}

// bigqueryDataTransferConfigLister implements the BigqueryDataTransferConfigLister interface.
type bigqueryDataTransferConfigLister struct {
	indexer cache.Indexer
}

// NewBigqueryDataTransferConfigLister returns a new BigqueryDataTransferConfigLister.
func NewBigqueryDataTransferConfigLister(indexer cache.Indexer) BigqueryDataTransferConfigLister {
	return &bigqueryDataTransferConfigLister{indexer: indexer}
}

// List lists all BigqueryDataTransferConfigs in the indexer.
func (s *bigqueryDataTransferConfigLister) List(selector labels.Selector) (ret []*v1alpha1.BigqueryDataTransferConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BigqueryDataTransferConfig))
	})
	return ret, err
}

// BigqueryDataTransferConfigs returns an object that can list and get BigqueryDataTransferConfigs.
func (s *bigqueryDataTransferConfigLister) BigqueryDataTransferConfigs(namespace string) BigqueryDataTransferConfigNamespaceLister {
	return bigqueryDataTransferConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BigqueryDataTransferConfigNamespaceLister helps list and get BigqueryDataTransferConfigs.
type BigqueryDataTransferConfigNamespaceLister interface {
	// List lists all BigqueryDataTransferConfigs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.BigqueryDataTransferConfig, err error)
	// Get retrieves the BigqueryDataTransferConfig from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.BigqueryDataTransferConfig, error)
	BigqueryDataTransferConfigNamespaceListerExpansion
}

// bigqueryDataTransferConfigNamespaceLister implements the BigqueryDataTransferConfigNamespaceLister
// interface.
type bigqueryDataTransferConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BigqueryDataTransferConfigs in the indexer for a given namespace.
func (s bigqueryDataTransferConfigNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BigqueryDataTransferConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BigqueryDataTransferConfig))
	})
	return ret, err
}

// Get retrieves the BigqueryDataTransferConfig from the indexer for a given namespace and name.
func (s bigqueryDataTransferConfigNamespaceLister) Get(name string) (*v1alpha1.BigqueryDataTransferConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("bigquerydatatransferconfig"), name)
	}
	return obj.(*v1alpha1.BigqueryDataTransferConfig), nil
}
