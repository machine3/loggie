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
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1beta1 "loggie.io/loggie/pkg/discovery/kubernetes/apis/loggie/v1beta1"
)

// FakeLogConfigs implements LogConfigInterface
type FakeLogConfigs struct {
	Fake *FakeLoggieV1beta1
	ns   string
}

var logconfigsResource = schema.GroupVersionResource{Group: "loggie.io", Version: "v1beta1", Resource: "logconfigs"}

var logconfigsKind = schema.GroupVersionKind{Group: "loggie.io", Version: "v1beta1", Kind: "LogConfig"}

// Get takes name of the logConfig, and returns the corresponding logConfig object, and an error if there is any.
func (c *FakeLogConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.LogConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(logconfigsResource, c.ns, name), &v1beta1.LogConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.LogConfig), err
}

// List takes label and field selectors, and returns the list of LogConfigs that match those selectors.
func (c *FakeLogConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.LogConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(logconfigsResource, logconfigsKind, c.ns, opts), &v1beta1.LogConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.LogConfigList{ListMeta: obj.(*v1beta1.LogConfigList).ListMeta}
	for _, item := range obj.(*v1beta1.LogConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested logConfigs.
func (c *FakeLogConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(logconfigsResource, c.ns, opts))

}

// Create takes the representation of a logConfig and creates it.  Returns the server's representation of the logConfig, and an error, if there is any.
func (c *FakeLogConfigs) Create(ctx context.Context, logConfig *v1beta1.LogConfig, opts v1.CreateOptions) (result *v1beta1.LogConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(logconfigsResource, c.ns, logConfig), &v1beta1.LogConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.LogConfig), err
}

// Update takes the representation of a logConfig and updates it. Returns the server's representation of the logConfig, and an error, if there is any.
func (c *FakeLogConfigs) Update(ctx context.Context, logConfig *v1beta1.LogConfig, opts v1.UpdateOptions) (result *v1beta1.LogConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(logconfigsResource, c.ns, logConfig), &v1beta1.LogConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.LogConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLogConfigs) UpdateStatus(ctx context.Context, logConfig *v1beta1.LogConfig, opts v1.UpdateOptions) (*v1beta1.LogConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(logconfigsResource, "status", c.ns, logConfig), &v1beta1.LogConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.LogConfig), err
}

// Delete takes name of the logConfig and deletes it. Returns an error if one occurs.
func (c *FakeLogConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(logconfigsResource, c.ns, name), &v1beta1.LogConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLogConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(logconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.LogConfigList{})
	return err
}

// Patch applies the patch and returns the patched logConfig.
func (c *FakeLogConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.LogConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(logconfigsResource, c.ns, name, pt, data, subresources...), &v1beta1.LogConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.LogConfig), err
}
