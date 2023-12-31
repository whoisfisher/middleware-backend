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

// FederatedSecretsGetter has a method to return a FederatedSecretInterface.
// A group's client should implement this interface.
type FederatedSecretsGetter interface {
	FederatedSecrets(namespace string) FederatedSecretInterface
}

// FederatedSecretInterface has methods to work with FederatedSecret resources.
type FederatedSecretInterface interface {
	Create(ctx context.Context, federatedSecret *v1beta1.FederatedSecret, opts v1.CreateOptions) (*v1beta1.FederatedSecret, error)
	Update(ctx context.Context, federatedSecret *v1beta1.FederatedSecret, opts v1.UpdateOptions) (*v1beta1.FederatedSecret, error)
	UpdateStatus(ctx context.Context, federatedSecret *v1beta1.FederatedSecret, opts v1.UpdateOptions) (*v1beta1.FederatedSecret, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.FederatedSecret, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.FederatedSecretList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.FederatedSecret, err error)
	FederatedSecretExpansion
}

// federatedSecrets implements FederatedSecretInterface
type federatedSecrets struct {
	client rest.Interface
	ns     string
}

// newFederatedSecrets returns a FederatedSecrets
func newFederatedSecrets(c *TypesV1beta1Client, namespace string) *federatedSecrets {
	return &federatedSecrets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the federatedSecret, and returns the corresponding federatedSecret object, and an error if there is any.
func (c *federatedSecrets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.FederatedSecret, err error) {
	result = &v1beta1.FederatedSecret{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedsecrets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FederatedSecrets that match those selectors.
func (c *federatedSecrets) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.FederatedSecretList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.FederatedSecretList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedsecrets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested federatedSecrets.
func (c *federatedSecrets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("federatedsecrets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a federatedSecret and creates it.  Returns the server's representation of the federatedSecret, and an error, if there is any.
func (c *federatedSecrets) Create(ctx context.Context, federatedSecret *v1beta1.FederatedSecret, opts v1.CreateOptions) (result *v1beta1.FederatedSecret, err error) {
	result = &v1beta1.FederatedSecret{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("federatedsecrets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(federatedSecret).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a federatedSecret and updates it. Returns the server's representation of the federatedSecret, and an error, if there is any.
func (c *federatedSecrets) Update(ctx context.Context, federatedSecret *v1beta1.FederatedSecret, opts v1.UpdateOptions) (result *v1beta1.FederatedSecret, err error) {
	result = &v1beta1.FederatedSecret{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedsecrets").
		Name(federatedSecret.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(federatedSecret).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *federatedSecrets) UpdateStatus(ctx context.Context, federatedSecret *v1beta1.FederatedSecret, opts v1.UpdateOptions) (result *v1beta1.FederatedSecret, err error) {
	result = &v1beta1.FederatedSecret{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedsecrets").
		Name(federatedSecret.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(federatedSecret).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the federatedSecret and deletes it. Returns an error if one occurs.
func (c *federatedSecrets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedsecrets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *federatedSecrets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedsecrets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched federatedSecret.
func (c *federatedSecrets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.FederatedSecret, err error) {
	result = &v1beta1.FederatedSecret{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("federatedsecrets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
