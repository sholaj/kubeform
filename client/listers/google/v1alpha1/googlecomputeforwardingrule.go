/*
Copyright 2019 The Kubeform Authors.

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
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// GoogleComputeForwardingRuleLister helps list GoogleComputeForwardingRules.
type GoogleComputeForwardingRuleLister interface {
	// List lists all GoogleComputeForwardingRules in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.GoogleComputeForwardingRule, err error)
	// Get retrieves the GoogleComputeForwardingRule from the index for a given name.
	Get(name string) (*v1alpha1.GoogleComputeForwardingRule, error)
	GoogleComputeForwardingRuleListerExpansion
}

// googleComputeForwardingRuleLister implements the GoogleComputeForwardingRuleLister interface.
type googleComputeForwardingRuleLister struct {
	indexer cache.Indexer
}

// NewGoogleComputeForwardingRuleLister returns a new GoogleComputeForwardingRuleLister.
func NewGoogleComputeForwardingRuleLister(indexer cache.Indexer) GoogleComputeForwardingRuleLister {
	return &googleComputeForwardingRuleLister{indexer: indexer}
}

// List lists all GoogleComputeForwardingRules in the indexer.
func (s *googleComputeForwardingRuleLister) List(selector labels.Selector) (ret []*v1alpha1.GoogleComputeForwardingRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GoogleComputeForwardingRule))
	})
	return ret, err
}

// Get retrieves the GoogleComputeForwardingRule from the index for a given name.
func (s *googleComputeForwardingRuleLister) Get(name string) (*v1alpha1.GoogleComputeForwardingRule, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("googlecomputeforwardingrule"), name)
	}
	return obj.(*v1alpha1.GoogleComputeForwardingRule), nil
}