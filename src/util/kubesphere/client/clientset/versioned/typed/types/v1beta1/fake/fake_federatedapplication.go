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

// FakeFederatedApplications implements FederatedApplicationInterface
type FakeFederatedApplications struct {
	Fake *FakeTypesV1beta1
	ns   string
}

var federatedapplicationsResource = schema.GroupVersionResource{Group: "types.kubefed.io", Version: "v1beta1", Resource: "federatedapplications"}

var federatedapplicationsKind = schema.GroupVersionKind{Group: "types.kubefed.io", Version: "v1beta1", Kind: "FederatedApplication"}

// Get takes name of the federatedApplication, and returns the corresponding federatedApplication object, and an error if there is any.
func (c *FakeFederatedApplications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.FederatedApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(federatedapplicationsResource, c.ns, name), &v1beta1.FederatedApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedApplication), err
}

// List takes label and field selectors, and returns the list of FederatedApplications that match those selectors.
func (c *FakeFederatedApplications) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.FederatedApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(federatedapplicationsResource, federatedapplicationsKind, c.ns, opts), &v1beta1.FederatedApplicationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.FederatedApplicationList{ListMeta: obj.(*v1beta1.FederatedApplicationList).ListMeta}
	for _, item := range obj.(*v1beta1.FederatedApplicationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested federatedApplications.
func (c *FakeFederatedApplications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(federatedapplicationsResource, c.ns, opts))

}

// Create takes the representation of a federatedApplication and creates it.  Returns the server's representation of the federatedApplication, and an error, if there is any.
func (c *FakeFederatedApplications) Create(ctx context.Context, federatedApplication *v1beta1.FederatedApplication, opts v1.CreateOptions) (result *v1beta1.FederatedApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(federatedapplicationsResource, c.ns, federatedApplication), &v1beta1.FederatedApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedApplication), err
}

// Update takes the representation of a federatedApplication and updates it. Returns the server's representation of the federatedApplication, and an error, if there is any.
func (c *FakeFederatedApplications) Update(ctx context.Context, federatedApplication *v1beta1.FederatedApplication, opts v1.UpdateOptions) (result *v1beta1.FederatedApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(federatedapplicationsResource, c.ns, federatedApplication), &v1beta1.FederatedApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedApplication), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFederatedApplications) UpdateStatus(ctx context.Context, federatedApplication *v1beta1.FederatedApplication, opts v1.UpdateOptions) (*v1beta1.FederatedApplication, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(federatedapplicationsResource, "status", c.ns, federatedApplication), &v1beta1.FederatedApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedApplication), err
}

// Delete takes name of the federatedApplication and deletes it. Returns an error if one occurs.
func (c *FakeFederatedApplications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(federatedapplicationsResource, c.ns, name), &v1beta1.FederatedApplication{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFederatedApplications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(federatedapplicationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.FederatedApplicationList{})
	return err
}

// Patch applies the patch and returns the patched federatedApplication.
func (c *FakeFederatedApplications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.FederatedApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(federatedapplicationsResource, c.ns, name, pt, data, subresources...), &v1beta1.FederatedApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedApplication), err
}
