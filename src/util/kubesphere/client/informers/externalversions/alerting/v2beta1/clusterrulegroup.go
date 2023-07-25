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

package v2beta1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	alertingv2beta1 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/alerting/v2beta1"
	versioned "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/client/clientset/versioned"
	internalinterfaces "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/client/informers/externalversions/internalinterfaces"
	v2beta1 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/client/listers/alerting/v2beta1"
)

// ClusterRuleGroupInformer provides access to a shared informer and lister for
// ClusterRuleGroups.
type ClusterRuleGroupInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v2beta1.ClusterRuleGroupLister
}

type clusterRuleGroupInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterRuleGroupInformer constructs a new informer for ClusterRuleGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterRuleGroupInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterRuleGroupInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterRuleGroupInformer constructs a new informer for ClusterRuleGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterRuleGroupInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AlertingV2beta1().ClusterRuleGroups().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AlertingV2beta1().ClusterRuleGroups().Watch(context.TODO(), options)
			},
		},
		&alertingv2beta1.ClusterRuleGroup{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterRuleGroupInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterRuleGroupInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterRuleGroupInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&alertingv2beta1.ClusterRuleGroup{}, f.defaultInformer)
}

func (f *clusterRuleGroupInformer) Lister() v2beta1.ClusterRuleGroupLister {
	return v2beta1.NewClusterRuleGroupLister(f.Informer().GetIndexer())
}
