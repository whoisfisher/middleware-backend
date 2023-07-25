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

package v1alpha3

import (
	internalinterfaces "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// DevOpsProjects returns a DevOpsProjectInformer.
	DevOpsProjects() DevOpsProjectInformer
	// Pipelines returns a PipelineInformer.
	Pipelines() PipelineInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// DevOpsProjects returns a DevOpsProjectInformer.
func (v *version) DevOpsProjects() DevOpsProjectInformer {
	return &devOpsProjectInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Pipelines returns a PipelineInformer.
func (v *version) Pipelines() PipelineInformer {
	return &pipelineInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
