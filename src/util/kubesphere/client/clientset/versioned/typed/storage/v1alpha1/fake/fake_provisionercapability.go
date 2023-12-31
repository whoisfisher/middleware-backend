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
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/storage/v1alpha1"
)

// FakeProvisionerCapabilities implements ProvisionerCapabilityInterface
type FakeProvisionerCapabilities struct {
	Fake *FakeStorageV1alpha1
}

var provisionercapabilitiesResource = schema.GroupVersionResource{Group: "storage.kubesphere.io", Version: "v1alpha1", Resource: "provisionercapabilities"}

var provisionercapabilitiesKind = schema.GroupVersionKind{Group: "storage.kubesphere.io", Version: "v1alpha1", Kind: "ProvisionerCapability"}

// Get takes name of the provisionerCapability, and returns the corresponding provisionerCapability object, and an error if there is any.
func (c *FakeProvisionerCapabilities) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ProvisionerCapability, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(provisionercapabilitiesResource, name), &v1alpha1.ProvisionerCapability{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProvisionerCapability), err
}

// List takes label and field selectors, and returns the list of ProvisionerCapabilities that match those selectors.
func (c *FakeProvisionerCapabilities) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ProvisionerCapabilityList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(provisionercapabilitiesResource, provisionercapabilitiesKind, opts), &v1alpha1.ProvisionerCapabilityList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ProvisionerCapabilityList{ListMeta: obj.(*v1alpha1.ProvisionerCapabilityList).ListMeta}
	for _, item := range obj.(*v1alpha1.ProvisionerCapabilityList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested provisionerCapabilities.
func (c *FakeProvisionerCapabilities) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(provisionercapabilitiesResource, opts))
}

// Create takes the representation of a provisionerCapability and creates it.  Returns the server's representation of the provisionerCapability, and an error, if there is any.
func (c *FakeProvisionerCapabilities) Create(ctx context.Context, provisionerCapability *v1alpha1.ProvisionerCapability, opts v1.CreateOptions) (result *v1alpha1.ProvisionerCapability, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(provisionercapabilitiesResource, provisionerCapability), &v1alpha1.ProvisionerCapability{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProvisionerCapability), err
}

// Update takes the representation of a provisionerCapability and updates it. Returns the server's representation of the provisionerCapability, and an error, if there is any.
func (c *FakeProvisionerCapabilities) Update(ctx context.Context, provisionerCapability *v1alpha1.ProvisionerCapability, opts v1.UpdateOptions) (result *v1alpha1.ProvisionerCapability, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(provisionercapabilitiesResource, provisionerCapability), &v1alpha1.ProvisionerCapability{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProvisionerCapability), err
}

// Delete takes name of the provisionerCapability and deletes it. Returns an error if one occurs.
func (c *FakeProvisionerCapabilities) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(provisionercapabilitiesResource, name), &v1alpha1.ProvisionerCapability{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProvisionerCapabilities) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(provisionercapabilitiesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ProvisionerCapabilityList{})
	return err
}

// Patch applies the patch and returns the patched provisionerCapability.
func (c *FakeProvisionerCapabilities) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ProvisionerCapability, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(provisionercapabilitiesResource, name, pt, data, subresources...), &v1alpha1.ProvisionerCapability{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProvisionerCapability), err
}
