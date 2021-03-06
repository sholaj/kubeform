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

// PinpointAppLister helps list PinpointApps.
type PinpointAppLister interface {
	// List lists all PinpointApps in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.PinpointApp, err error)
	// PinpointApps returns an object that can list and get PinpointApps.
	PinpointApps(namespace string) PinpointAppNamespaceLister
	PinpointAppListerExpansion
}

// pinpointAppLister implements the PinpointAppLister interface.
type pinpointAppLister struct {
	indexer cache.Indexer
}

// NewPinpointAppLister returns a new PinpointAppLister.
func NewPinpointAppLister(indexer cache.Indexer) PinpointAppLister {
	return &pinpointAppLister{indexer: indexer}
}

// List lists all PinpointApps in the indexer.
func (s *pinpointAppLister) List(selector labels.Selector) (ret []*v1alpha1.PinpointApp, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PinpointApp))
	})
	return ret, err
}

// PinpointApps returns an object that can list and get PinpointApps.
func (s *pinpointAppLister) PinpointApps(namespace string) PinpointAppNamespaceLister {
	return pinpointAppNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PinpointAppNamespaceLister helps list and get PinpointApps.
type PinpointAppNamespaceLister interface {
	// List lists all PinpointApps in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.PinpointApp, err error)
	// Get retrieves the PinpointApp from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.PinpointApp, error)
	PinpointAppNamespaceListerExpansion
}

// pinpointAppNamespaceLister implements the PinpointAppNamespaceLister
// interface.
type pinpointAppNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PinpointApps in the indexer for a given namespace.
func (s pinpointAppNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PinpointApp, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PinpointApp))
	})
	return ret, err
}

// Get retrieves the PinpointApp from the indexer for a given namespace and name.
func (s pinpointAppNamespaceLister) Get(name string) (*v1alpha1.PinpointApp, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("pinpointapp"), name)
	}
	return obj.(*v1alpha1.PinpointApp), nil
}
