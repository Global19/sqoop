/*
Copyright The Kubernetes Authors.

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
	solo_io_v1 "github.com/solo-io/sqoop/pkg/storage/crd/solo.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeResolverMaps implements ResolverMapInterface
type FakeResolverMaps struct {
	Fake *FakeSqoopV1
	ns   string
}

var resolvermapsResource = schema.GroupVersionResource{Group: "sqoop.solo.io", Version: "v1", Resource: "resolvermaps"}

var resolvermapsKind = schema.GroupVersionKind{Group: "sqoop.solo.io", Version: "v1", Kind: "ResolverMap"}

// Get takes name of the resolverMap, and returns the corresponding resolverMap object, and an error if there is any.
func (c *FakeResolverMaps) Get(name string, options v1.GetOptions) (result *solo_io_v1.ResolverMap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(resolvermapsResource, c.ns, name), &solo_io_v1.ResolverMap{})

	if obj == nil {
		return nil, err
	}
	return obj.(*solo_io_v1.ResolverMap), err
}

// List takes label and field selectors, and returns the list of ResolverMaps that match those selectors.
func (c *FakeResolverMaps) List(opts v1.ListOptions) (result *solo_io_v1.ResolverMapList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(resolvermapsResource, resolvermapsKind, c.ns, opts), &solo_io_v1.ResolverMapList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &solo_io_v1.ResolverMapList{}
	for _, item := range obj.(*solo_io_v1.ResolverMapList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resolverMaps.
func (c *FakeResolverMaps) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(resolvermapsResource, c.ns, opts))

}

// Create takes the representation of a resolverMap and creates it.  Returns the server's representation of the resolverMap, and an error, if there is any.
func (c *FakeResolverMaps) Create(resolverMap *solo_io_v1.ResolverMap) (result *solo_io_v1.ResolverMap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(resolvermapsResource, c.ns, resolverMap), &solo_io_v1.ResolverMap{})

	if obj == nil {
		return nil, err
	}
	return obj.(*solo_io_v1.ResolverMap), err
}

// Update takes the representation of a resolverMap and updates it. Returns the server's representation of the resolverMap, and an error, if there is any.
func (c *FakeResolverMaps) Update(resolverMap *solo_io_v1.ResolverMap) (result *solo_io_v1.ResolverMap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(resolvermapsResource, c.ns, resolverMap), &solo_io_v1.ResolverMap{})

	if obj == nil {
		return nil, err
	}
	return obj.(*solo_io_v1.ResolverMap), err
}

// Delete takes name of the resolverMap and deletes it. Returns an error if one occurs.
func (c *FakeResolverMaps) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(resolvermapsResource, c.ns, name), &solo_io_v1.ResolverMap{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResolverMaps) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(resolvermapsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &solo_io_v1.ResolverMapList{})
	return err
}

// Patch applies the patch and returns the patched resolverMap.
func (c *FakeResolverMaps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *solo_io_v1.ResolverMap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resolvermapsResource, c.ns, name, data, subresources...), &solo_io_v1.ResolverMap{})

	if obj == nil {
		return nil, err
	}
	return obj.(*solo_io_v1.ResolverMap), err
}
