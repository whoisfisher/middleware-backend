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

package v2beta1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v2beta1 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/alerting/v2beta1"
	scheme "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/client/clientset/versioned/scheme"
)

// ClusterRuleGroupsGetter has a method to return a ClusterRuleGroupInterface.
// A group's client should implement this interface.
type ClusterRuleGroupsGetter interface {
	ClusterRuleGroups() ClusterRuleGroupInterface
}

// ClusterRuleGroupInterface has methods to work with ClusterRuleGroup resources.
type ClusterRuleGroupInterface interface {
	Create(ctx context.Context, clusterRuleGroup *v2beta1.ClusterRuleGroup, opts v1.CreateOptions) (*v2beta1.ClusterRuleGroup, error)
	Update(ctx context.Context, clusterRuleGroup *v2beta1.ClusterRuleGroup, opts v1.UpdateOptions) (*v2beta1.ClusterRuleGroup, error)
	UpdateStatus(ctx context.Context, clusterRuleGroup *v2beta1.ClusterRuleGroup, opts v1.UpdateOptions) (*v2beta1.ClusterRuleGroup, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2beta1.ClusterRuleGroup, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2beta1.ClusterRuleGroupList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2beta1.ClusterRuleGroup, err error)
	ClusterRuleGroupExpansion
}

// clusterRuleGroups implements ClusterRuleGroupInterface
type clusterRuleGroups struct {
	client rest.Interface
}

// newClusterRuleGroups returns a ClusterRuleGroups
func newClusterRuleGroups(c *AlertingV2beta1Client) *clusterRuleGroups {
	return &clusterRuleGroups{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterRuleGroup, and returns the corresponding clusterRuleGroup object, and an error if there is any.
func (c *clusterRuleGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2beta1.ClusterRuleGroup, err error) {
	result = &v2beta1.ClusterRuleGroup{}
	err = c.client.Get().
		Resource("clusterrulegroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterRuleGroups that match those selectors.
func (c *clusterRuleGroups) List(ctx context.Context, opts v1.ListOptions) (result *v2beta1.ClusterRuleGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2beta1.ClusterRuleGroupList{}
	err = c.client.Get().
		Resource("clusterrulegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterRuleGroups.
func (c *clusterRuleGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clusterrulegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterRuleGroup and creates it.  Returns the server's representation of the clusterRuleGroup, and an error, if there is any.
func (c *clusterRuleGroups) Create(ctx context.Context, clusterRuleGroup *v2beta1.ClusterRuleGroup, opts v1.CreateOptions) (result *v2beta1.ClusterRuleGroup, err error) {
	result = &v2beta1.ClusterRuleGroup{}
	err = c.client.Post().
		Resource("clusterrulegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterRuleGroup).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterRuleGroup and updates it. Returns the server's representation of the clusterRuleGroup, and an error, if there is any.
func (c *clusterRuleGroups) Update(ctx context.Context, clusterRuleGroup *v2beta1.ClusterRuleGroup, opts v1.UpdateOptions) (result *v2beta1.ClusterRuleGroup, err error) {
	result = &v2beta1.ClusterRuleGroup{}
	err = c.client.Put().
		Resource("clusterrulegroups").
		Name(clusterRuleGroup.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterRuleGroup).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterRuleGroups) UpdateStatus(ctx context.Context, clusterRuleGroup *v2beta1.ClusterRuleGroup, opts v1.UpdateOptions) (result *v2beta1.ClusterRuleGroup, err error) {
	result = &v2beta1.ClusterRuleGroup{}
	err = c.client.Put().
		Resource("clusterrulegroups").
		Name(clusterRuleGroup.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterRuleGroup).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterRuleGroup and deletes it. Returns an error if one occurs.
func (c *clusterRuleGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusterrulegroups").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterRuleGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clusterrulegroups").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterRuleGroup.
func (c *clusterRuleGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2beta1.ClusterRuleGroup, err error) {
	result = &v2beta1.ClusterRuleGroup{}
	err = c.client.Patch(pt).
		Resource("clusterrulegroups").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
