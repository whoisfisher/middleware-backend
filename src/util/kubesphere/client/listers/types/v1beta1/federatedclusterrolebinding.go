/*
Copyright 2020 The KubeSphere Authors.

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

package v1beta1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1beta1 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/types/v1beta1"
)

// FederatedClusterRoleBindingLister helps list FederatedClusterRoleBindings.
// All objects returned here must be treated as read-only.
type FederatedClusterRoleBindingLister interface {
	// List lists all FederatedClusterRoleBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.FederatedClusterRoleBinding, err error)
	// FederatedClusterRoleBindings returns an object that can list and get FederatedClusterRoleBindings.
	FederatedClusterRoleBindings(namespace string) FederatedClusterRoleBindingNamespaceLister
	FederatedClusterRoleBindingListerExpansion
}

// federatedClusterRoleBindingLister implements the FederatedClusterRoleBindingLister interface.
type federatedClusterRoleBindingLister struct {
	indexer cache.Indexer
}

// NewFederatedClusterRoleBindingLister returns a new FederatedClusterRoleBindingLister.
func NewFederatedClusterRoleBindingLister(indexer cache.Indexer) FederatedClusterRoleBindingLister {
	return &federatedClusterRoleBindingLister{indexer: indexer}
}

// List lists all FederatedClusterRoleBindings in the indexer.
func (s *federatedClusterRoleBindingLister) List(selector labels.Selector) (ret []*v1beta1.FederatedClusterRoleBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FederatedClusterRoleBinding))
	})
	return ret, err
}

// FederatedClusterRoleBindings returns an object that can list and get FederatedClusterRoleBindings.
func (s *federatedClusterRoleBindingLister) FederatedClusterRoleBindings(namespace string) FederatedClusterRoleBindingNamespaceLister {
	return federatedClusterRoleBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FederatedClusterRoleBindingNamespaceLister helps list and get FederatedClusterRoleBindings.
// All objects returned here must be treated as read-only.
type FederatedClusterRoleBindingNamespaceLister interface {
	// List lists all FederatedClusterRoleBindings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.FederatedClusterRoleBinding, err error)
	// Get retrieves the FederatedClusterRoleBinding from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.FederatedClusterRoleBinding, error)
	FederatedClusterRoleBindingNamespaceListerExpansion
}

// federatedClusterRoleBindingNamespaceLister implements the FederatedClusterRoleBindingNamespaceLister
// interface.
type federatedClusterRoleBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FederatedClusterRoleBindings in the indexer for a given namespace.
func (s federatedClusterRoleBindingNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.FederatedClusterRoleBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FederatedClusterRoleBinding))
	})
	return ret, err
}

// Get retrieves the FederatedClusterRoleBinding from the indexer for a given namespace and name.
func (s federatedClusterRoleBindingNamespaceLister) Get(name string) (*v1beta1.FederatedClusterRoleBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("federatedclusterrolebinding"), name)
	}
	return obj.(*v1beta1.FederatedClusterRoleBinding), nil
}
