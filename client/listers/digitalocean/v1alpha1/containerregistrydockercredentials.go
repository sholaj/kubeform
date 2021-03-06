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
	v1alpha1 "kubeform.dev/kubeform/apis/digitalocean/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ContainerRegistryDockerCredentialsLister helps list ContainerRegistryDockerCredentialses.
type ContainerRegistryDockerCredentialsLister interface {
	// List lists all ContainerRegistryDockerCredentialses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ContainerRegistryDockerCredentials, err error)
	// ContainerRegistryDockerCredentialses returns an object that can list and get ContainerRegistryDockerCredentialses.
	ContainerRegistryDockerCredentialses(namespace string) ContainerRegistryDockerCredentialsNamespaceLister
	ContainerRegistryDockerCredentialsListerExpansion
}

// containerRegistryDockerCredentialsLister implements the ContainerRegistryDockerCredentialsLister interface.
type containerRegistryDockerCredentialsLister struct {
	indexer cache.Indexer
}

// NewContainerRegistryDockerCredentialsLister returns a new ContainerRegistryDockerCredentialsLister.
func NewContainerRegistryDockerCredentialsLister(indexer cache.Indexer) ContainerRegistryDockerCredentialsLister {
	return &containerRegistryDockerCredentialsLister{indexer: indexer}
}

// List lists all ContainerRegistryDockerCredentialses in the indexer.
func (s *containerRegistryDockerCredentialsLister) List(selector labels.Selector) (ret []*v1alpha1.ContainerRegistryDockerCredentials, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ContainerRegistryDockerCredentials))
	})
	return ret, err
}

// ContainerRegistryDockerCredentialses returns an object that can list and get ContainerRegistryDockerCredentialses.
func (s *containerRegistryDockerCredentialsLister) ContainerRegistryDockerCredentialses(namespace string) ContainerRegistryDockerCredentialsNamespaceLister {
	return containerRegistryDockerCredentialsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ContainerRegistryDockerCredentialsNamespaceLister helps list and get ContainerRegistryDockerCredentialses.
type ContainerRegistryDockerCredentialsNamespaceLister interface {
	// List lists all ContainerRegistryDockerCredentialses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ContainerRegistryDockerCredentials, err error)
	// Get retrieves the ContainerRegistryDockerCredentials from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ContainerRegistryDockerCredentials, error)
	ContainerRegistryDockerCredentialsNamespaceListerExpansion
}

// containerRegistryDockerCredentialsNamespaceLister implements the ContainerRegistryDockerCredentialsNamespaceLister
// interface.
type containerRegistryDockerCredentialsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ContainerRegistryDockerCredentialses in the indexer for a given namespace.
func (s containerRegistryDockerCredentialsNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ContainerRegistryDockerCredentials, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ContainerRegistryDockerCredentials))
	})
	return ret, err
}

// Get retrieves the ContainerRegistryDockerCredentials from the indexer for a given namespace and name.
func (s containerRegistryDockerCredentialsNamespaceLister) Get(name string) (*v1alpha1.ContainerRegistryDockerCredentials, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("containerregistrydockercredentials"), name)
	}
	return obj.(*v1alpha1.ContainerRegistryDockerCredentials), nil
}
