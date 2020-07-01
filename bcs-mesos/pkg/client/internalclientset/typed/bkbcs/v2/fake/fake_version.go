/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v2 "github.com/Tencent/bk-bcs/bcs-mesos/pkg/apis/bkbcs/v2"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVersions implements VersionInterface
type FakeVersions struct {
	Fake *FakeBkbcsV2
	ns   string
}

var versionsResource = schema.GroupVersionResource{Group: "bkbcs.tencent.com", Version: "v2", Resource: "versions"}

var versionsKind = schema.GroupVersionKind{Group: "bkbcs.tencent.com", Version: "v2", Kind: "Version"}

// Get takes name of the version, and returns the corresponding version object, and an error if there is any.
func (c *FakeVersions) Get(name string, options v1.GetOptions) (result *v2.Version, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(versionsResource, c.ns, name), &v2.Version{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.Version), err
}

// List takes label and field selectors, and returns the list of Versions that match those selectors.
func (c *FakeVersions) List(opts v1.ListOptions) (result *v2.VersionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(versionsResource, versionsKind, c.ns, opts), &v2.VersionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2.VersionList{ListMeta: obj.(*v2.VersionList).ListMeta}
	for _, item := range obj.(*v2.VersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested versions.
func (c *FakeVersions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(versionsResource, c.ns, opts))

}

// Create takes the representation of a version and creates it.  Returns the server's representation of the version, and an error, if there is any.
func (c *FakeVersions) Create(version *v2.Version) (result *v2.Version, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(versionsResource, c.ns, version), &v2.Version{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.Version), err
}

// Update takes the representation of a version and updates it. Returns the server's representation of the version, and an error, if there is any.
func (c *FakeVersions) Update(version *v2.Version) (result *v2.Version, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(versionsResource, c.ns, version), &v2.Version{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.Version), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVersions) UpdateStatus(version *v2.Version) (*v2.Version, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(versionsResource, "status", c.ns, version), &v2.Version{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.Version), err
}

// Delete takes name of the version and deletes it. Returns an error if one occurs.
func (c *FakeVersions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(versionsResource, c.ns, name), &v2.Version{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVersions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(versionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v2.VersionList{})
	return err
}

// Patch applies the patch and returns the patched version.
func (c *FakeVersions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2.Version, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(versionsResource, c.ns, name, data, subresources...), &v2.Version{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.Version), err
}
