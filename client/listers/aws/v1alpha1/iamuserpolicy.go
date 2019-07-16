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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// IamUserPolicyLister helps list IamUserPolicies.
type IamUserPolicyLister interface {
	// List lists all IamUserPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.IamUserPolicy, err error)
	// Get retrieves the IamUserPolicy from the index for a given name.
	Get(name string) (*v1alpha1.IamUserPolicy, error)
	IamUserPolicyListerExpansion
}

// iamUserPolicyLister implements the IamUserPolicyLister interface.
type iamUserPolicyLister struct {
	indexer cache.Indexer
}

// NewIamUserPolicyLister returns a new IamUserPolicyLister.
func NewIamUserPolicyLister(indexer cache.Indexer) IamUserPolicyLister {
	return &iamUserPolicyLister{indexer: indexer}
}

// List lists all IamUserPolicies in the indexer.
func (s *iamUserPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.IamUserPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamUserPolicy))
	})
	return ret, err
}

// Get retrieves the IamUserPolicy from the index for a given name.
func (s *iamUserPolicyLister) Get(name string) (*v1alpha1.IamUserPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iamuserpolicy"), name)
	}
	return obj.(*v1alpha1.IamUserPolicy), nil
}