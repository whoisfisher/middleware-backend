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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	networkv1alpha1 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/network/v1alpha1"
	versioned "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/client/clientset/versioned"
	internalinterfaces "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/client/listers/network/v1alpha1"
)

// NamespaceNetworkPolicyInformer provides access to a shared informer and lister for
// NamespaceNetworkPolicies.
type NamespaceNetworkPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.NamespaceNetworkPolicyLister
}

type namespaceNetworkPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewNamespaceNetworkPolicyInformer constructs a new informer for NamespaceNetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNamespaceNetworkPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNamespaceNetworkPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredNamespaceNetworkPolicyInformer constructs a new informer for NamespaceNetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNamespaceNetworkPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkV1alpha1().NamespaceNetworkPolicies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkV1alpha1().NamespaceNetworkPolicies(namespace).Watch(context.TODO(), options)
			},
		},
		&networkv1alpha1.NamespaceNetworkPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *namespaceNetworkPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNamespaceNetworkPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *namespaceNetworkPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkv1alpha1.NamespaceNetworkPolicy{}, f.defaultInformer)
}

func (f *namespaceNetworkPolicyInformer) Lister() v1alpha1.NamespaceNetworkPolicyLister {
	return v1alpha1.NewNamespaceNetworkPolicyLister(f.Informer().GetIndexer())
}
