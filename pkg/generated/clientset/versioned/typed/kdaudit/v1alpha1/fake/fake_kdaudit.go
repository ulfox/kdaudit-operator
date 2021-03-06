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

	v1alpha1 "github.com/ulfox/kdaudit-operator/pkg/apis/kdaudit/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKDAudits implements KDAuditInterface
type FakeKDAudits struct {
	Fake *FakeKdauditV1alpha1
	ns   string
}

var kdauditsResource = schema.GroupVersionResource{Group: "kdaudit.k8s.io", Version: "v1alpha1", Resource: "kdaudits"}

var kdauditsKind = schema.GroupVersionKind{Group: "kdaudit.k8s.io", Version: "v1alpha1", Kind: "KDAudit"}

// Get takes name of the kDAudit, and returns the corresponding kDAudit object, and an error if there is any.
func (c *FakeKDAudits) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KDAudit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kdauditsResource, c.ns, name), &v1alpha1.KDAudit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KDAudit), err
}

// List takes label and field selectors, and returns the list of KDAudits that match those selectors.
func (c *FakeKDAudits) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KDAuditList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kdauditsResource, kdauditsKind, c.ns, opts), &v1alpha1.KDAuditList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KDAuditList{ListMeta: obj.(*v1alpha1.KDAuditList).ListMeta}
	for _, item := range obj.(*v1alpha1.KDAuditList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kDAudits.
func (c *FakeKDAudits) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kdauditsResource, c.ns, opts))

}

// Create takes the representation of a kDAudit and creates it.  Returns the server's representation of the kDAudit, and an error, if there is any.
func (c *FakeKDAudits) Create(ctx context.Context, kDAudit *v1alpha1.KDAudit, opts v1.CreateOptions) (result *v1alpha1.KDAudit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kdauditsResource, c.ns, kDAudit), &v1alpha1.KDAudit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KDAudit), err
}

// Update takes the representation of a kDAudit and updates it. Returns the server's representation of the kDAudit, and an error, if there is any.
func (c *FakeKDAudits) Update(ctx context.Context, kDAudit *v1alpha1.KDAudit, opts v1.UpdateOptions) (result *v1alpha1.KDAudit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kdauditsResource, c.ns, kDAudit), &v1alpha1.KDAudit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KDAudit), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKDAudits) UpdateStatus(ctx context.Context, kDAudit *v1alpha1.KDAudit, opts v1.UpdateOptions) (*v1alpha1.KDAudit, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kdauditsResource, "status", c.ns, kDAudit), &v1alpha1.KDAudit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KDAudit), err
}

// Delete takes name of the kDAudit and deletes it. Returns an error if one occurs.
func (c *FakeKDAudits) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kdauditsResource, c.ns, name), &v1alpha1.KDAudit{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKDAudits) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kdauditsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.KDAuditList{})
	return err
}

// Patch applies the patch and returns the patched kDAudit.
func (c *FakeKDAudits) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KDAudit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kdauditsResource, c.ns, name, pt, data, subresources...), &v1alpha1.KDAudit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KDAudit), err
}
