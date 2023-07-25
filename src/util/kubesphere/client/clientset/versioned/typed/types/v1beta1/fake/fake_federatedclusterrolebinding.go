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

// FakeFederatedClusterRoleBindings implements FederatedClusterRoleBindingInterface
type FakeFederatedClusterRoleBindings struct {
	Fake *FakeTypesV1beta1
	ns   string
}

var federatedclusterrolebindingsResource = schema.GroupVersionResource{Group: "types.kubefed.io", Version: "v1beta1", Resource: "federatedclusterrolebindings"}

var federatedclusterrolebindingsKind = schema.GroupVersionKind{Group: "types.kubefed.io", Version: "v1beta1", Kind: "FederatedClusterRoleBinding"}

// Get takes name of the federatedClusterRoleBinding, and returns the corresponding federatedClusterRoleBinding object, and an error if there is any.
func (c *FakeFederatedClusterRoleBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.FederatedClusterRoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(federatedclusterrolebindingsResource, c.ns, name), &v1beta1.FederatedClusterRoleBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedClusterRoleBinding), err
}

// List takes label and field selectors, and returns the list of FederatedClusterRoleBindings that match those selectors.
func (c *FakeFederatedClusterRoleBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.FederatedClusterRoleBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(federatedclusterrolebindingsResource, federatedclusterrolebindingsKind, c.ns, opts), &v1beta1.FederatedClusterRoleBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.FederatedClusterRoleBindingList{ListMeta: obj.(*v1beta1.FederatedClusterRoleBindingList).ListMeta}
	for _, item := range obj.(*v1beta1.FederatedClusterRoleBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested federatedClusterRoleBindings.
func (c *FakeFederatedClusterRoleBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(federatedclusterrolebindingsResource, c.ns, opts))

}

// Create takes the representation of a federatedClusterRoleBinding and creates it.  Returns the server's representation of the federatedClusterRoleBinding, and an error, if there is any.
func (c *FakeFederatedClusterRoleBindings) Create(ctx context.Context, federatedClusterRoleBinding *v1beta1.FederatedClusterRoleBinding, opts v1.CreateOptions) (result *v1beta1.FederatedClusterRoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(federatedclusterrolebindingsResource, c.ns, federatedClusterRoleBinding), &v1beta1.FederatedClusterRoleBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedClusterRoleBinding), err
}

// Update takes the representation of a federatedClusterRoleBinding and updates it. Returns the server's representation of the federatedClusterRoleBinding, and an error, if there is any.
func (c *FakeFederatedClusterRoleBindings) Update(ctx context.Context, federatedClusterRoleBinding *v1beta1.FederatedClusterRoleBinding, opts v1.UpdateOptions) (result *v1beta1.FederatedClusterRoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(federatedclusterrolebindingsResource, c.ns, federatedClusterRoleBinding), &v1beta1.FederatedClusterRoleBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedClusterRoleBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFederatedClusterRoleBindings) UpdateStatus(ctx context.Context, federatedClusterRoleBinding *v1beta1.FederatedClusterRoleBinding, opts v1.UpdateOptions) (*v1beta1.FederatedClusterRoleBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(federatedclusterrolebindingsResource, "status", c.ns, federatedClusterRoleBinding), &v1beta1.FederatedClusterRoleBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedClusterRoleBinding), err
}

// Delete takes name of the federatedClusterRoleBinding and deletes it. Returns an error if one occurs.
func (c *FakeFederatedClusterRoleBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(federatedclusterrolebindingsResource, c.ns, name), &v1beta1.FederatedClusterRoleBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFederatedClusterRoleBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(federatedclusterrolebindingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.FederatedClusterRoleBindingList{})
	return err
}

// Patch applies the patch and returns the patched federatedClusterRoleBinding.
func (c *FakeFederatedClusterRoleBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.FederatedClusterRoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(federatedclusterrolebindingsResource, c.ns, name, pt, data, subresources...), &v1beta1.FederatedClusterRoleBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedClusterRoleBinding), err
}
