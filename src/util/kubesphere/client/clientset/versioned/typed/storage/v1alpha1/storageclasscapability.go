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

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/storage/v1alpha1"
	scheme "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/client/clientset/versioned/scheme"
)

// StorageClassCapabilitiesGetter has a method to return a StorageClassCapabilityInterface.
// A group's client should implement this interface.
type StorageClassCapabilitiesGetter interface {
	StorageClassCapabilities() StorageClassCapabilityInterface
}

// StorageClassCapabilityInterface has methods to work with StorageClassCapability resources.
type StorageClassCapabilityInterface interface {
	Create(ctx context.Context, storageClassCapability *v1alpha1.StorageClassCapability, opts v1.CreateOptions) (*v1alpha1.StorageClassCapability, error)
	Update(ctx context.Context, storageClassCapability *v1alpha1.StorageClassCapability, opts v1.UpdateOptions) (*v1alpha1.StorageClassCapability, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.StorageClassCapability, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.StorageClassCapabilityList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.StorageClassCapability, err error)
	StorageClassCapabilityExpansion
}

// storageClassCapabilities implements StorageClassCapabilityInterface
type storageClassCapabilities struct {
	client rest.Interface
}

// newStorageClassCapabilities returns a StorageClassCapabilities
func newStorageClassCapabilities(c *StorageV1alpha1Client) *storageClassCapabilities {
	return &storageClassCapabilities{
		client: c.RESTClient(),
	}
}

// Get takes name of the storageClassCapability, and returns the corresponding storageClassCapability object, and an error if there is any.
func (c *storageClassCapabilities) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.StorageClassCapability, err error) {
	result = &v1alpha1.StorageClassCapability{}
	err = c.client.Get().
		Resource("storageclasscapabilities").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of StorageClassCapabilities that match those selectors.
func (c *storageClassCapabilities) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.StorageClassCapabilityList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.StorageClassCapabilityList{}
	err = c.client.Get().
		Resource("storageclasscapabilities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested storageClassCapabilities.
func (c *storageClassCapabilities) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("storageclasscapabilities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a storageClassCapability and creates it.  Returns the server's representation of the storageClassCapability, and an error, if there is any.
func (c *storageClassCapabilities) Create(ctx context.Context, storageClassCapability *v1alpha1.StorageClassCapability, opts v1.CreateOptions) (result *v1alpha1.StorageClassCapability, err error) {
	result = &v1alpha1.StorageClassCapability{}
	err = c.client.Post().
		Resource("storageclasscapabilities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(storageClassCapability).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a storageClassCapability and updates it. Returns the server's representation of the storageClassCapability, and an error, if there is any.
func (c *storageClassCapabilities) Update(ctx context.Context, storageClassCapability *v1alpha1.StorageClassCapability, opts v1.UpdateOptions) (result *v1alpha1.StorageClassCapability, err error) {
	result = &v1alpha1.StorageClassCapability{}
	err = c.client.Put().
		Resource("storageclasscapabilities").
		Name(storageClassCapability.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(storageClassCapability).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the storageClassCapability and deletes it. Returns an error if one occurs.
func (c *storageClassCapabilities) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("storageclasscapabilities").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *storageClassCapabilities) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("storageclasscapabilities").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched storageClassCapability.
func (c *storageClassCapabilities) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.StorageClassCapability, err error) {
	result = &v1alpha1.StorageClassCapability{}
	err = c.client.Patch(pt).
		Resource("storageclasscapabilities").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
