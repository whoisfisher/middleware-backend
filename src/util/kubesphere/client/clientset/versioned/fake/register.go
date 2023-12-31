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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	alertingv2beta1 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/alerting/v2beta1"
	applicationv1alpha1 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/application/v1alpha1"
	auditingv1alpha1 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/auditing/v1alpha1"
	clusterv1alpha1 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/cluster/v1alpha1"
	devopsv1alpha1 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/devops/v1alpha1"
	devopsv1alpha3 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/devops/v1alpha3"
	iamv1alpha2 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/iam/v1alpha2"
	networkv1alpha1 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/network/v1alpha1"
	notificationv2beta1 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/notification/v2beta1"
	notificationv2beta2 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/notification/v2beta2"
	quotav1alpha2 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/quota/v1alpha2"
	servicemeshv1alpha2 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/servicemesh/v1alpha2"
	storagev1alpha1 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/storage/v1alpha1"
	tenantv1alpha1 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/tenant/v1alpha1"
	tenantv1alpha2 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/tenant/v1alpha2"
	typesv1beta1 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/types/v1beta1"
	typesv1beta2 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/types/v1beta2"
)

var scheme = runtime.NewScheme()
var codecs = serializer.NewCodecFactory(scheme)

var localSchemeBuilder = runtime.SchemeBuilder{
	alertingv2beta1.AddToScheme,
	applicationv1alpha1.AddToScheme,
	auditingv1alpha1.AddToScheme,
	clusterv1alpha1.AddToScheme,
	devopsv1alpha1.AddToScheme,
	devopsv1alpha3.AddToScheme,
	iamv1alpha2.AddToScheme,
	networkv1alpha1.AddToScheme,
	notificationv2beta1.AddToScheme,
	notificationv2beta2.AddToScheme,
	quotav1alpha2.AddToScheme,
	servicemeshv1alpha2.AddToScheme,
	storagev1alpha1.AddToScheme,
	tenantv1alpha1.AddToScheme,
	tenantv1alpha2.AddToScheme,
	typesv1beta1.AddToScheme,
	typesv1beta2.AddToScheme,
}

// AddToScheme adds all types of this clientset into the given scheme. This allows composition
// of clientsets, like in:
//
//	import (
//	  "k8s.io/client-go/kubernetes"
//	  clientsetscheme "k8s.io/client-go/kubernetes/scheme"
//	  aggregatorclientsetscheme "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/scheme"
//	)
//
//	kclientset, _ := kubernetes.NewForConfig(c)
//	_ = aggregatorclientsetscheme.AddToScheme(clientsetscheme.Scheme)
//
// After this, RawExtensions in Kubernetes types will serialize kube-aggregator types
// correctly.
var AddToScheme = localSchemeBuilder.AddToScheme

func init() {
	v1.AddToGroupVersion(scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(AddToScheme(scheme))
}
