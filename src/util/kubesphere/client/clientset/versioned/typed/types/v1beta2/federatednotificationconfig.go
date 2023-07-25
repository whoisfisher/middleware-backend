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

package v1beta2

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1beta2 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/types/v1beta2"
	scheme "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/client/clientset/versioned/scheme"
)

// FederatedNotificationConfigsGetter has a method to return a FederatedNotificationConfigInterface.
// A group's client should implement this interface.
type FederatedNotificationConfigsGetter interface {
	FederatedNotificationConfigs() FederatedNotificationConfigInterface
}

// FederatedNotificationConfigInterface has methods to work with FederatedNotificationConfig resources.
type FederatedNotificationConfigInterface interface {
	Create(ctx context.Context, federatedNotificationConfig *v1beta2.FederatedNotificationConfig, opts v1.CreateOptions) (*v1beta2.FederatedNotificationConfig, error)
	Update(ctx context.Context, federatedNotificationConfig *v1beta2.FederatedNotificationConfig, opts v1.UpdateOptions) (*v1beta2.FederatedNotificationConfig, error)
	UpdateStatus(ctx context.Context, federatedNotificationConfig *v1beta2.FederatedNotificationConfig, opts v1.UpdateOptions) (*v1beta2.FederatedNotificationConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta2.FederatedNotificationConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta2.FederatedNotificationConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.FederatedNotificationConfig, err error)
	FederatedNotificationConfigExpansion
}

// federatedNotificationConfigs implements FederatedNotificationConfigInterface
type federatedNotificationConfigs struct {
	client rest.Interface
}

// newFederatedNotificationConfigs returns a FederatedNotificationConfigs
func newFederatedNotificationConfigs(c *TypesV1beta2Client) *federatedNotificationConfigs {
	return &federatedNotificationConfigs{
		client: c.RESTClient(),
	}
}

// Get takes name of the federatedNotificationConfig, and returns the corresponding federatedNotificationConfig object, and an error if there is any.
func (c *federatedNotificationConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.FederatedNotificationConfig, err error) {
	result = &v1beta2.FederatedNotificationConfig{}
	err = c.client.Get().
		Resource("federatednotificationconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FederatedNotificationConfigs that match those selectors.
func (c *federatedNotificationConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.FederatedNotificationConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta2.FederatedNotificationConfigList{}
	err = c.client.Get().
		Resource("federatednotificationconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested federatedNotificationConfigs.
func (c *federatedNotificationConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("federatednotificationconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a federatedNotificationConfig and creates it.  Returns the server's representation of the federatedNotificationConfig, and an error, if there is any.
func (c *federatedNotificationConfigs) Create(ctx context.Context, federatedNotificationConfig *v1beta2.FederatedNotificationConfig, opts v1.CreateOptions) (result *v1beta2.FederatedNotificationConfig, err error) {
	result = &v1beta2.FederatedNotificationConfig{}
	err = c.client.Post().
		Resource("federatednotificationconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(federatedNotificationConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a federatedNotificationConfig and updates it. Returns the server's representation of the federatedNotificationConfig, and an error, if there is any.
func (c *federatedNotificationConfigs) Update(ctx context.Context, federatedNotificationConfig *v1beta2.FederatedNotificationConfig, opts v1.UpdateOptions) (result *v1beta2.FederatedNotificationConfig, err error) {
	result = &v1beta2.FederatedNotificationConfig{}
	err = c.client.Put().
		Resource("federatednotificationconfigs").
		Name(federatedNotificationConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(federatedNotificationConfig).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *federatedNotificationConfigs) UpdateStatus(ctx context.Context, federatedNotificationConfig *v1beta2.FederatedNotificationConfig, opts v1.UpdateOptions) (result *v1beta2.FederatedNotificationConfig, err error) {
	result = &v1beta2.FederatedNotificationConfig{}
	err = c.client.Put().
		Resource("federatednotificationconfigs").
		Name(federatedNotificationConfig.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(federatedNotificationConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the federatedNotificationConfig and deletes it. Returns an error if one occurs.
func (c *federatedNotificationConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("federatednotificationconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *federatedNotificationConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("federatednotificationconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched federatedNotificationConfig.
func (c *federatedNotificationConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.FederatedNotificationConfig, err error) {
	result = &v1beta2.FederatedNotificationConfig{}
	err = c.client.Patch(pt).
		Resource("federatednotificationconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
