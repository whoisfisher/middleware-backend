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
	v1alpha2 "github.com/mensylisir/kmpp-middleware/src/util/kubesphere/api/iam/v1alpha2"
)

// FakeLoginRecords implements LoginRecordInterface
type FakeLoginRecords struct {
	Fake *FakeIamV1alpha2
}

var loginrecordsResource = schema.GroupVersionResource{Group: "iam.kubesphere.io", Version: "v1alpha2", Resource: "loginrecords"}

var loginrecordsKind = schema.GroupVersionKind{Group: "iam.kubesphere.io", Version: "v1alpha2", Kind: "LoginRecord"}

// Get takes name of the loginRecord, and returns the corresponding loginRecord object, and an error if there is any.
func (c *FakeLoginRecords) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.LoginRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(loginrecordsResource, name), &v1alpha2.LoginRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.LoginRecord), err
}

// List takes label and field selectors, and returns the list of LoginRecords that match those selectors.
func (c *FakeLoginRecords) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.LoginRecordList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(loginrecordsResource, loginrecordsKind, opts), &v1alpha2.LoginRecordList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.LoginRecordList{ListMeta: obj.(*v1alpha2.LoginRecordList).ListMeta}
	for _, item := range obj.(*v1alpha2.LoginRecordList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested loginRecords.
func (c *FakeLoginRecords) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(loginrecordsResource, opts))
}

// Create takes the representation of a loginRecord and creates it.  Returns the server's representation of the loginRecord, and an error, if there is any.
func (c *FakeLoginRecords) Create(ctx context.Context, loginRecord *v1alpha2.LoginRecord, opts v1.CreateOptions) (result *v1alpha2.LoginRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(loginrecordsResource, loginRecord), &v1alpha2.LoginRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.LoginRecord), err
}

// Update takes the representation of a loginRecord and updates it. Returns the server's representation of the loginRecord, and an error, if there is any.
func (c *FakeLoginRecords) Update(ctx context.Context, loginRecord *v1alpha2.LoginRecord, opts v1.UpdateOptions) (result *v1alpha2.LoginRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(loginrecordsResource, loginRecord), &v1alpha2.LoginRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.LoginRecord), err
}

// Delete takes name of the loginRecord and deletes it. Returns an error if one occurs.
func (c *FakeLoginRecords) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(loginrecordsResource, name), &v1alpha2.LoginRecord{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLoginRecords) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(loginrecordsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.LoginRecordList{})
	return err
}

// Patch applies the patch and returns the patched loginRecord.
func (c *FakeLoginRecords) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.LoginRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(loginrecordsResource, name, pt, data, subresources...), &v1alpha2.LoginRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.LoginRecord), err
}
