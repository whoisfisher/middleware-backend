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

// FederatedApplicationLister helps list FederatedApplications.
// All objects returned here must be treated as read-only.
type FederatedApplicationLister interface {
	// List lists all FederatedApplications in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.FederatedApplication, err error)
	// FederatedApplications returns an object that can list and get FederatedApplications.
	FederatedApplications(namespace string) FederatedApplicationNamespaceLister
	FederatedApplicationListerExpansion
}

// federatedApplicationLister implements the FederatedApplicationLister interface.
type federatedApplicationLister struct {
	indexer cache.Indexer
}

// NewFederatedApplicationLister returns a new FederatedApplicationLister.
func NewFederatedApplicationLister(indexer cache.Indexer) FederatedApplicationLister {
	return &federatedApplicationLister{indexer: indexer}
}

// List lists all FederatedApplications in the indexer.
func (s *federatedApplicationLister) List(selector labels.Selector) (ret []*v1beta1.FederatedApplication, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FederatedApplication))
	})
	return ret, err
}

// FederatedApplications returns an object that can list and get FederatedApplications.
func (s *federatedApplicationLister) FederatedApplications(namespace string) FederatedApplicationNamespaceLister {
	return federatedApplicationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FederatedApplicationNamespaceLister helps list and get FederatedApplications.
// All objects returned here must be treated as read-only.
type FederatedApplicationNamespaceLister interface {
	// List lists all FederatedApplications in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.FederatedApplication, err error)
	// Get retrieves the FederatedApplication from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.FederatedApplication, error)
	FederatedApplicationNamespaceListerExpansion
}

// federatedApplicationNamespaceLister implements the FederatedApplicationNamespaceLister
// interface.
type federatedApplicationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FederatedApplications in the indexer for a given namespace.
func (s federatedApplicationNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.FederatedApplication, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FederatedApplication))
	})
	return ret, err
}

// Get retrieves the FederatedApplication from the indexer for a given namespace and name.
func (s federatedApplicationNamespaceLister) Get(name string) (*v1beta1.FederatedApplication, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("federatedapplication"), name)
	}
	return obj.(*v1beta1.FederatedApplication), nil
}
