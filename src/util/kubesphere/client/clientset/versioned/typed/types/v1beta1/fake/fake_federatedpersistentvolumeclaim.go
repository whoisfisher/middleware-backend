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
	v1beta1 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/types/v1beta1"
)

// FakeFederatedPersistentVolumeClaims implements FederatedPersistentVolumeClaimInterface
type FakeFederatedPersistentVolumeClaims struct {
	Fake *FakeTypesV1beta1
	ns   string
}

var federatedpersistentvolumeclaimsResource = schema.GroupVersionResource{Group: "types.kubefed.io", Version: "v1beta1", Resource: "federatedpersistentvolumeclaims"}

var federatedpersistentvolumeclaimsKind = schema.GroupVersionKind{Group: "types.kubefed.io", Version: "v1beta1", Kind: "FederatedPersistentVolumeClaim"}

// Get takes name of the federatedPersistentVolumeClaim, and returns the corresponding federatedPersistentVolumeClaim object, and an error if there is any.
func (c *FakeFederatedPersistentVolumeClaims) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.FederatedPersistentVolumeClaim, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(federatedpersistentvolumeclaimsResource, c.ns, name), &v1beta1.FederatedPersistentVolumeClaim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedPersistentVolumeClaim), err
}

// List takes label and field selectors, and returns the list of FederatedPersistentVolumeClaims that match those selectors.
func (c *FakeFederatedPersistentVolumeClaims) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.FederatedPersistentVolumeClaimList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(federatedpersistentvolumeclaimsResource, federatedpersistentvolumeclaimsKind, c.ns, opts), &v1beta1.FederatedPersistentVolumeClaimList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.FederatedPersistentVolumeClaimList{ListMeta: obj.(*v1beta1.FederatedPersistentVolumeClaimList).ListMeta}
	for _, item := range obj.(*v1beta1.FederatedPersistentVolumeClaimList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested federatedPersistentVolumeClaims.
func (c *FakeFederatedPersistentVolumeClaims) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(federatedpersistentvolumeclaimsResource, c.ns, opts))

}

// Create takes the representation of a federatedPersistentVolumeClaim and creates it.  Returns the server's representation of the federatedPersistentVolumeClaim, and an error, if there is any.
func (c *FakeFederatedPersistentVolumeClaims) Create(ctx context.Context, federatedPersistentVolumeClaim *v1beta1.FederatedPersistentVolumeClaim, opts v1.CreateOptions) (result *v1beta1.FederatedPersistentVolumeClaim, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(federatedpersistentvolumeclaimsResource, c.ns, federatedPersistentVolumeClaim), &v1beta1.FederatedPersistentVolumeClaim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedPersistentVolumeClaim), err
}

// Update takes the representation of a federatedPersistentVolumeClaim and updates it. Returns the server's representation of the federatedPersistentVolumeClaim, and an error, if there is any.
func (c *FakeFederatedPersistentVolumeClaims) Update(ctx context.Context, federatedPersistentVolumeClaim *v1beta1.FederatedPersistentVolumeClaim, opts v1.UpdateOptions) (result *v1beta1.FederatedPersistentVolumeClaim, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(federatedpersistentvolumeclaimsResource, c.ns, federatedPersistentVolumeClaim), &v1beta1.FederatedPersistentVolumeClaim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedPersistentVolumeClaim), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFederatedPersistentVolumeClaims) UpdateStatus(ctx context.Context, federatedPersistentVolumeClaim *v1beta1.FederatedPersistentVolumeClaim, opts v1.UpdateOptions) (*v1beta1.FederatedPersistentVolumeClaim, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(federatedpersistentvolumeclaimsResource, "status", c.ns, federatedPersistentVolumeClaim), &v1beta1.FederatedPersistentVolumeClaim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedPersistentVolumeClaim), err
}

// Delete takes name of the federatedPersistentVolumeClaim and deletes it. Returns an error if one occurs.
func (c *FakeFederatedPersistentVolumeClaims) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(federatedpersistentvolumeclaimsResource, c.ns, name), &v1beta1.FederatedPersistentVolumeClaim{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFederatedPersistentVolumeClaims) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(federatedpersistentvolumeclaimsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.FederatedPersistentVolumeClaimList{})
	return err
}

// Patch applies the patch and returns the patched federatedPersistentVolumeClaim.
func (c *FakeFederatedPersistentVolumeClaims) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.FederatedPersistentVolumeClaim, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(federatedpersistentvolumeclaimsResource, c.ns, name, pt, data, subresources...), &v1beta1.FederatedPersistentVolumeClaim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedPersistentVolumeClaim), err
}
