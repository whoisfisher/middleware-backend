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

package v2beta2

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	notificationv2beta2 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/notification/v2beta2"
	versioned "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/client/clientset/versioned"
	internalinterfaces "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/client/informers/externalversions/internalinterfaces"
	v2beta2 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/client/listers/notification/v2beta2"
)

// ReceiverInformer provides access to a shared informer and lister for
// Receivers.
type ReceiverInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v2beta2.ReceiverLister
}

type receiverInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewReceiverInformer constructs a new informer for Receiver type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewReceiverInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredReceiverInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredReceiverInformer constructs a new informer for Receiver type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredReceiverInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NotificationV2beta2().Receivers().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NotificationV2beta2().Receivers().Watch(context.TODO(), options)
			},
		},
		&notificationv2beta2.Receiver{},
		resyncPeriod,
		indexers,
	)
}

func (f *receiverInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredReceiverInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *receiverInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&notificationv2beta2.Receiver{}, f.defaultInformer)
}

func (f *receiverInformer) Lister() v2beta2.ReceiverLister {
	return v2beta2.NewReceiverLister(f.Informer().GetIndexer())
}
