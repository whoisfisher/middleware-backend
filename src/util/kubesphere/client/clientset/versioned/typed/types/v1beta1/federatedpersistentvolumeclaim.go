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

package v1beta1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1beta1 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/types/v1beta1"
	scheme "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/client/clientset/versioned/scheme"
)

// FederatedPersistentVolumeClaimsGetter has a method to return a FederatedPersistentVolumeClaimInterface.
// A group's client should implement this interface.
type FederatedPersistentVolumeClaimsGetter interface {
	FederatedPersistentVolumeClaims(namespace string) FederatedPersistentVolumeClaimInterface
}

// FederatedPersistentVolumeClaimInterface has methods to work with FederatedPersistentVolumeClaim resources.
type FederatedPersistentVolumeClaimInterface interface {
	Create(ctx context.Context, federatedPersistentVolumeClaim *v1beta1.FederatedPersistentVolumeClaim, opts v1.CreateOptions) (*v1beta1.FederatedPersistentVolumeClaim, error)
	Update(ctx context.Context, federatedPersistentVolumeClaim *v1beta1.FederatedPersistentVolumeClaim, opts v1.UpdateOptions) (*v1beta1.FederatedPersistentVolumeClaim, error)
	UpdateStatus(ctx context.Context, federatedPersistentVolumeClaim *v1beta1.FederatedPersistentVolumeClaim, opts v1.UpdateOptions) (*v1beta1.FederatedPersistentVolumeClaim, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.FederatedPersistentVolumeClaim, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.FederatedPersistentVolumeClaimList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.FederatedPersistentVolumeClaim, err error)
	FederatedPersistentVolumeClaimExpansion
}

// federatedPersistentVolumeClaims implements FederatedPersistentVolumeClaimInterface
type federatedPersistentVolumeClaims struct {
	client rest.Interface
	ns     string
}

// newFederatedPersistentVolumeClaims returns a FederatedPersistentVolumeClaims
func newFederatedPersistentVolumeClaims(c *TypesV1beta1Client, namespace string) *federatedPersistentVolumeClaims {
	return &federatedPersistentVolumeClaims{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the federatedPersistentVolumeClaim, and returns the corresponding federatedPersistentVolumeClaim object, and an error if there is any.
func (c *federatedPersistentVolumeClaims) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.FederatedPersistentVolumeClaim, err error) {
	result = &v1beta1.FederatedPersistentVolumeClaim{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedpersistentvolumeclaims").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FederatedPersistentVolumeClaims that match those selectors.
func (c *federatedPersistentVolumeClaims) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.FederatedPersistentVolumeClaimList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.FederatedPersistentVolumeClaimList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedpersistentvolumeclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested federatedPersistentVolumeClaims.
func (c *federatedPersistentVolumeClaims) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("federatedpersistentvolumeclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a federatedPersistentVolumeClaim and creates it.  Returns the server's representation of the federatedPersistentVolumeClaim, and an error, if there is any.
func (c *federatedPersistentVolumeClaims) Create(ctx context.Context, federatedPersistentVolumeClaim *v1beta1.FederatedPersistentVolumeClaim, opts v1.CreateOptions) (result *v1beta1.FederatedPersistentVolumeClaim, err error) {
	result = &v1beta1.FederatedPersistentVolumeClaim{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("federatedpersistentvolumeclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(federatedPersistentVolumeClaim).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a federatedPersistentVolumeClaim and updates it. Returns the server's representation of the federatedPersistentVolumeClaim, and an error, if there is any.
func (c *federatedPersistentVolumeClaims) Update(ctx context.Context, federatedPersistentVolumeClaim *v1beta1.FederatedPersistentVolumeClaim, opts v1.UpdateOptions) (result *v1beta1.FederatedPersistentVolumeClaim, err error) {
	result = &v1beta1.FederatedPersistentVolumeClaim{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedpersistentvolumeclaims").
		Name(federatedPersistentVolumeClaim.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(federatedPersistentVolumeClaim).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *federatedPersistentVolumeClaims) UpdateStatus(ctx context.Context, federatedPersistentVolumeClaim *v1beta1.FederatedPersistentVolumeClaim, opts v1.UpdateOptions) (result *v1beta1.FederatedPersistentVolumeClaim, err error) {
	result = &v1beta1.FederatedPersistentVolumeClaim{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedpersistentvolumeclaims").
		Name(federatedPersistentVolumeClaim.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(federatedPersistentVolumeClaim).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the federatedPersistentVolumeClaim and deletes it. Returns an error if one occurs.
func (c *federatedPersistentVolumeClaims) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedpersistentvolumeclaims").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *federatedPersistentVolumeClaims) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedpersistentvolumeclaims").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched federatedPersistentVolumeClaim.
func (c *federatedPersistentVolumeClaims) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.FederatedPersistentVolumeClaim, err error) {
	result = &v1beta1.FederatedPersistentVolumeClaim{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("federatedpersistentvolumeclaims").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
