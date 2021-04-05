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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/ulfox/kdaudit-operator/pkg/apis/kdaudit/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KDAuditLister helps list KDAudits.
// All objects returned here must be treated as read-only.
type KDAuditLister interface {
	// List lists all KDAudits in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.KDAudit, err error)
	// KDAudits returns an object that can list and get KDAudits.
	KDAudits(namespace string) KDAuditNamespaceLister
	KDAuditListerExpansion
}

// kDAuditLister implements the KDAuditLister interface.
type kDAuditLister struct {
	indexer cache.Indexer
}

// NewKDAuditLister returns a new KDAuditLister.
func NewKDAuditLister(indexer cache.Indexer) KDAuditLister {
	return &kDAuditLister{indexer: indexer}
}

// List lists all KDAudits in the indexer.
func (s *kDAuditLister) List(selector labels.Selector) (ret []*v1alpha1.KDAudit, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KDAudit))
	})
	return ret, err
}

// KDAudits returns an object that can list and get KDAudits.
func (s *kDAuditLister) KDAudits(namespace string) KDAuditNamespaceLister {
	return kDAuditNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KDAuditNamespaceLister helps list and get KDAudits.
// All objects returned here must be treated as read-only.
type KDAuditNamespaceLister interface {
	// List lists all KDAudits in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.KDAudit, err error)
	// Get retrieves the KDAudit from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.KDAudit, error)
	KDAuditNamespaceListerExpansion
}

// kDAuditNamespaceLister implements the KDAuditNamespaceLister
// interface.
type kDAuditNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all KDAudits in the indexer for a given namespace.
func (s kDAuditNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.KDAudit, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KDAudit))
	})
	return ret, err
}

// Get retrieves the KDAudit from the indexer for a given namespace and name.
func (s kDAuditNamespaceLister) Get(name string) (*v1alpha1.KDAudit, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("kdaudit"), name)
	}
	return obj.(*v1alpha1.KDAudit), nil
}
