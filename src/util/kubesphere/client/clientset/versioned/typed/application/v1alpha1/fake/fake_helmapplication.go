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
	v1alpha1 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/application/v1alpha1"
)

// FakeHelmApplications implements HelmApplicationInterface
type FakeHelmApplications struct {
	Fake *FakeApplicationV1alpha1
}

var helmapplicationsResource = schema.GroupVersionResource{Group: "application.kubesphere.io", Version: "v1alpha1", Resource: "helmapplications"}

var helmapplicationsKind = schema.GroupVersionKind{Group: "application.kubesphere.io", Version: "v1alpha1", Kind: "HelmApplication"}

// Get takes name of the helmApplication, and returns the corresponding helmApplication object, and an error if there is any.
func (c *FakeHelmApplications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HelmApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(helmapplicationsResource, name), &v1alpha1.HelmApplication{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HelmApplication), err
}

// List takes label and field selectors, and returns the list of HelmApplications that match those selectors.
func (c *FakeHelmApplications) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HelmApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(helmapplicationsResource, helmapplicationsKind, opts), &v1alpha1.HelmApplicationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HelmApplicationList{ListMeta: obj.(*v1alpha1.HelmApplicationList).ListMeta}
	for _, item := range obj.(*v1alpha1.HelmApplicationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested helmApplications.
func (c *FakeHelmApplications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(helmapplicationsResource, opts))
}

// Create takes the representation of a helmApplication and creates it.  Returns the server's representation of the helmApplication, and an error, if there is any.
func (c *FakeHelmApplications) Create(ctx context.Context, helmApplication *v1alpha1.HelmApplication, opts v1.CreateOptions) (result *v1alpha1.HelmApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(helmapplicationsResource, helmApplication), &v1alpha1.HelmApplication{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HelmApplication), err
}

// Update takes the representation of a helmApplication and updates it. Returns the server's representation of the helmApplication, and an error, if there is any.
func (c *FakeHelmApplications) Update(ctx context.Context, helmApplication *v1alpha1.HelmApplication, opts v1.UpdateOptions) (result *v1alpha1.HelmApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(helmapplicationsResource, helmApplication), &v1alpha1.HelmApplication{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HelmApplication), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHelmApplications) UpdateStatus(ctx context.Context, helmApplication *v1alpha1.HelmApplication, opts v1.UpdateOptions) (*v1alpha1.HelmApplication, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(helmapplicationsResource, "status", helmApplication), &v1alpha1.HelmApplication{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HelmApplication), err
}

// Delete takes name of the helmApplication and deletes it. Returns an error if one occurs.
func (c *FakeHelmApplications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(helmapplicationsResource, name), &v1alpha1.HelmApplication{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHelmApplications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(helmapplicationsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HelmApplicationList{})
	return err
}

// Patch applies the patch and returns the patched helmApplication.
func (c *FakeHelmApplications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HelmApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(helmapplicationsResource, name, pt, data, subresources...), &v1alpha1.HelmApplication{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HelmApplication), err
}
