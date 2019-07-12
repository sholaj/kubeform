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

// AwsSesDomainIdentityLister helps list AwsSesDomainIdentities.
type AwsSesDomainIdentityLister interface {
	// List lists all AwsSesDomainIdentities in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsSesDomainIdentity, err error)
	// Get retrieves the AwsSesDomainIdentity from the index for a given name.
	Get(name string) (*v1alpha1.AwsSesDomainIdentity, error)
	AwsSesDomainIdentityListerExpansion
}

// awsSesDomainIdentityLister implements the AwsSesDomainIdentityLister interface.
type awsSesDomainIdentityLister struct {
	indexer cache.Indexer
}

// NewAwsSesDomainIdentityLister returns a new AwsSesDomainIdentityLister.
func NewAwsSesDomainIdentityLister(indexer cache.Indexer) AwsSesDomainIdentityLister {
	return &awsSesDomainIdentityLister{indexer: indexer}
}

// List lists all AwsSesDomainIdentities in the indexer.
func (s *awsSesDomainIdentityLister) List(selector labels.Selector) (ret []*v1alpha1.AwsSesDomainIdentity, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsSesDomainIdentity))
	})
	return ret, err
}

// Get retrieves the AwsSesDomainIdentity from the index for a given name.
func (s *awsSesDomainIdentityLister) Get(name string) (*v1alpha1.AwsSesDomainIdentity, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awssesdomainidentity"), name)
	}
	return obj.(*v1alpha1.AwsSesDomainIdentity), nil
}