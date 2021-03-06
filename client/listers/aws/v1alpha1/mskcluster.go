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

// MskClusterLister helps list MskClusters.
type MskClusterLister interface {
	// List lists all MskClusters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MskCluster, err error)
	// MskClusters returns an object that can list and get MskClusters.
	MskClusters(namespace string) MskClusterNamespaceLister
	MskClusterListerExpansion
}

// mskClusterLister implements the MskClusterLister interface.
type mskClusterLister struct {
	indexer cache.Indexer
}

// NewMskClusterLister returns a new MskClusterLister.
func NewMskClusterLister(indexer cache.Indexer) MskClusterLister {
	return &mskClusterLister{indexer: indexer}
}

// List lists all MskClusters in the indexer.
func (s *mskClusterLister) List(selector labels.Selector) (ret []*v1alpha1.MskCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MskCluster))
	})
	return ret, err
}

// MskClusters returns an object that can list and get MskClusters.
func (s *mskClusterLister) MskClusters(namespace string) MskClusterNamespaceLister {
	return mskClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MskClusterNamespaceLister helps list and get MskClusters.
type MskClusterNamespaceLister interface {
	// List lists all MskClusters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.MskCluster, err error)
	// Get retrieves the MskCluster from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.MskCluster, error)
	MskClusterNamespaceListerExpansion
}

// mskClusterNamespaceLister implements the MskClusterNamespaceLister
// interface.
type mskClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MskClusters in the indexer for a given namespace.
func (s mskClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MskCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MskCluster))
	})
	return ret, err
}

// Get retrieves the MskCluster from the indexer for a given namespace and name.
func (s mskClusterNamespaceLister) Get(name string) (*v1alpha1.MskCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("mskcluster"), name)
	}
	return obj.(*v1alpha1.MskCluster), nil
}
