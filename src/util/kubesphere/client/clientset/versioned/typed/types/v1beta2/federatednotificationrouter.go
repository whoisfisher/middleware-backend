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

// FederatedNotificationRoutersGetter has a method to return a FederatedNotificationRouterInterface.
// A group's client should implement this interface.
type FederatedNotificationRoutersGetter interface {
	FederatedNotificationRouters() FederatedNotificationRouterInterface
}

// FederatedNotificationRouterInterface has methods to work with FederatedNotificationRouter resources.
type FederatedNotificationRouterInterface interface {
	Create(ctx context.Context, federatedNotificationRouter *v1beta2.FederatedNotificationRouter, opts v1.CreateOptions) (*v1beta2.FederatedNotificationRouter, error)
	Update(ctx context.Context, federatedNotificationRouter *v1beta2.FederatedNotificationRouter, opts v1.UpdateOptions) (*v1beta2.FederatedNotificationRouter, error)
	UpdateStatus(ctx context.Context, federatedNotificationRouter *v1beta2.FederatedNotificationRouter, opts v1.UpdateOptions) (*v1beta2.FederatedNotificationRouter, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta2.FederatedNotificationRouter, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta2.FederatedNotificationRouterList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.FederatedNotificationRouter, err error)
	FederatedNotificationRouterExpansion
}

// federatedNotificationRouters implements FederatedNotificationRouterInterface
type federatedNotificationRouters struct {
	client rest.Interface
}

// newFederatedNotificationRouters returns a FederatedNotificationRouters
func newFederatedNotificationRouters(c *TypesV1beta2Client) *federatedNotificationRouters {
	return &federatedNotificationRouters{
		client: c.RESTClient(),
	}
}

// Get takes name of the federatedNotificationRouter, and returns the corresponding federatedNotificationRouter object, and an error if there is any.
func (c *federatedNotificationRouters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.FederatedNotificationRouter, err error) {
	result = &v1beta2.FederatedNotificationRouter{}
	err = c.client.Get().
		Resource("federatednotificationrouters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FederatedNotificationRouters that match those selectors.
func (c *federatedNotificationRouters) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.FederatedNotificationRouterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta2.FederatedNotificationRouterList{}
	err = c.client.Get().
		Resource("federatednotificationrouters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested federatedNotificationRouters.
func (c *federatedNotificationRouters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("federatednotificationrouters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a federatedNotificationRouter and creates it.  Returns the server's representation of the federatedNotificationRouter, and an error, if there is any.
func (c *federatedNotificationRouters) Create(ctx context.Context, federatedNotificationRouter *v1beta2.FederatedNotificationRouter, opts v1.CreateOptions) (result *v1beta2.FederatedNotificationRouter, err error) {
	result = &v1beta2.FederatedNotificationRouter{}
	err = c.client.Post().
		Resource("federatednotificationrouters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(federatedNotificationRouter).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a federatedNotificationRouter and updates it. Returns the server's representation of the federatedNotificationRouter, and an error, if there is any.
func (c *federatedNotificationRouters) Update(ctx context.Context, federatedNotificationRouter *v1beta2.FederatedNotificationRouter, opts v1.UpdateOptions) (result *v1beta2.FederatedNotificationRouter, err error) {
	result = &v1beta2.FederatedNotificationRouter{}
	err = c.client.Put().
		Resource("federatednotificationrouters").
		Name(federatedNotificationRouter.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(federatedNotificationRouter).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *federatedNotificationRouters) UpdateStatus(ctx context.Context, federatedNotificationRouter *v1beta2.FederatedNotificationRouter, opts v1.UpdateOptions) (result *v1beta2.FederatedNotificationRouter, err error) {
	result = &v1beta2.FederatedNotificationRouter{}
	err = c.client.Put().
		Resource("federatednotificationrouters").
		Name(federatedNotificationRouter.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(federatedNotificationRouter).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the federatedNotificationRouter and deletes it. Returns an error if one occurs.
func (c *federatedNotificationRouters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("federatednotificationrouters").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *federatedNotificationRouters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("federatednotificationrouters").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched federatedNotificationRouter.
func (c *federatedNotificationRouters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.FederatedNotificationRouter, err error) {
	result = &v1beta2.FederatedNotificationRouter{}
	err = c.client.Patch(pt).
		Resource("federatednotificationrouters").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
