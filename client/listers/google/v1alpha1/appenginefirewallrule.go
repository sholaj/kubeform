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

// AppEngineFirewallRuleLister helps list AppEngineFirewallRules.
type AppEngineFirewallRuleLister interface {
	// List lists all AppEngineFirewallRules in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AppEngineFirewallRule, err error)
	// AppEngineFirewallRules returns an object that can list and get AppEngineFirewallRules.
	AppEngineFirewallRules(namespace string) AppEngineFirewallRuleNamespaceLister
	AppEngineFirewallRuleListerExpansion
}

// appEngineFirewallRuleLister implements the AppEngineFirewallRuleLister interface.
type appEngineFirewallRuleLister struct {
	indexer cache.Indexer
}

// NewAppEngineFirewallRuleLister returns a new AppEngineFirewallRuleLister.
func NewAppEngineFirewallRuleLister(indexer cache.Indexer) AppEngineFirewallRuleLister {
	return &appEngineFirewallRuleLister{indexer: indexer}
}

// List lists all AppEngineFirewallRules in the indexer.
func (s *appEngineFirewallRuleLister) List(selector labels.Selector) (ret []*v1alpha1.AppEngineFirewallRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppEngineFirewallRule))
	})
	return ret, err
}

// AppEngineFirewallRules returns an object that can list and get AppEngineFirewallRules.
func (s *appEngineFirewallRuleLister) AppEngineFirewallRules(namespace string) AppEngineFirewallRuleNamespaceLister {
	return appEngineFirewallRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AppEngineFirewallRuleNamespaceLister helps list and get AppEngineFirewallRules.
type AppEngineFirewallRuleNamespaceLister interface {
	// List lists all AppEngineFirewallRules in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AppEngineFirewallRule, err error)
	// Get retrieves the AppEngineFirewallRule from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AppEngineFirewallRule, error)
	AppEngineFirewallRuleNamespaceListerExpansion
}

// appEngineFirewallRuleNamespaceLister implements the AppEngineFirewallRuleNamespaceLister
// interface.
type appEngineFirewallRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AppEngineFirewallRules in the indexer for a given namespace.
func (s appEngineFirewallRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AppEngineFirewallRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppEngineFirewallRule))
	})
	return ret, err
}

// Get retrieves the AppEngineFirewallRule from the indexer for a given namespace and name.
func (s appEngineFirewallRuleNamespaceLister) Get(name string) (*v1alpha1.AppEngineFirewallRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("appenginefirewallrule"), name)
	}
	return obj.(*v1alpha1.AppEngineFirewallRule), nil
}