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

package notification

import (
	internalinterfaces "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/client/informers/externalversions/internalinterfaces"
	v2beta1 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/client/informers/externalversions/notification/v2beta1"
	v2beta2 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/client/informers/externalversions/notification/v2beta2"
)

// Interface provides access to each of this group's versions.
type Interface interface {
	// V2beta1 provides access to shared informers for resources in V2beta1.
	V2beta1() v2beta1.Interface
	// V2beta2 provides access to shared informers for resources in V2beta2.
	V2beta2() v2beta2.Interface
}

type group struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &group{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// V2beta1 returns a new v2beta1.Interface.
func (g *group) V2beta1() v2beta1.Interface {
	return v2beta1.New(g.factory, g.namespace, g.tweakListOptions)
}

// V2beta2 returns a new v2beta2.Interface.
func (g *group) V2beta2() v2beta2.Interface {
	return v2beta2.New(g.factory, g.namespace, g.tweakListOptions)
}
