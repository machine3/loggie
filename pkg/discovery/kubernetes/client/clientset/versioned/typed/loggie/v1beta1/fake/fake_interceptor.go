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

	v1beta1 "github.com/loggie-io/loggie/pkg/discovery/kubernetes/apis/loggie/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeInterceptors implements InterceptorInterface
type FakeInterceptors struct {
	Fake *FakeLoggieV1beta1
}

var interceptorsResource = schema.GroupVersionResource{Group: "loggie.io", Version: "v1beta1", Resource: "interceptors"}

var interceptorsKind = schema.GroupVersionKind{Group: "loggie.io", Version: "v1beta1", Kind: "Interceptor"}

// Get takes name of the interceptor, and returns the corresponding interceptor object, and an error if there is any.
func (c *FakeInterceptors) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Interceptor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(interceptorsResource, name), &v1beta1.Interceptor{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Interceptor), err
}

// List takes label and field selectors, and returns the list of Interceptors that match those selectors.
func (c *FakeInterceptors) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.InterceptorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(interceptorsResource, interceptorsKind, opts), &v1beta1.InterceptorList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.InterceptorList{ListMeta: obj.(*v1beta1.InterceptorList).ListMeta}
	for _, item := range obj.(*v1beta1.InterceptorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested interceptors.
func (c *FakeInterceptors) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(interceptorsResource, opts))
}

// Create takes the representation of a interceptor and creates it.  Returns the server's representation of the interceptor, and an error, if there is any.
func (c *FakeInterceptors) Create(ctx context.Context, interceptor *v1beta1.Interceptor, opts v1.CreateOptions) (result *v1beta1.Interceptor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(interceptorsResource, interceptor), &v1beta1.Interceptor{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Interceptor), err
}

// Update takes the representation of a interceptor and updates it. Returns the server's representation of the interceptor, and an error, if there is any.
func (c *FakeInterceptors) Update(ctx context.Context, interceptor *v1beta1.Interceptor, opts v1.UpdateOptions) (result *v1beta1.Interceptor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(interceptorsResource, interceptor), &v1beta1.Interceptor{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Interceptor), err
}

// Delete takes name of the interceptor and deletes it. Returns an error if one occurs.
func (c *FakeInterceptors) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(interceptorsResource, name, opts), &v1beta1.Interceptor{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInterceptors) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(interceptorsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.InterceptorList{})
	return err
}

// Patch applies the patch and returns the patched interceptor.
func (c *FakeInterceptors) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Interceptor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(interceptorsResource, name, pt, data, subresources...), &v1beta1.Interceptor{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Interceptor), err
}
