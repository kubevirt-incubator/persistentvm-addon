/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/kubevirt-incubator/persistentvm-addon/pkg/apis/persistentvm/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PVMLister helps list PVMs.
type PVMLister interface {
	// List lists all PVMs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.PVM, err error)
	// PVMs returns an object that can list and get PVMs.
	PVMs(namespace string) PVMNamespaceLister
	PVMListerExpansion
}

// pVMLister implements the PVMLister interface.
type pVMLister struct {
	indexer cache.Indexer
}

// NewPVMLister returns a new PVMLister.
func NewPVMLister(indexer cache.Indexer) PVMLister {
	return &pVMLister{indexer: indexer}
}

// List lists all PVMs in the indexer.
func (s *pVMLister) List(selector labels.Selector) (ret []*v1alpha1.PVM, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PVM))
	})
	return ret, err
}

// PVMs returns an object that can list and get PVMs.
func (s *pVMLister) PVMs(namespace string) PVMNamespaceLister {
	return pVMNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PVMNamespaceLister helps list and get PVMs.
type PVMNamespaceLister interface {
	// List lists all PVMs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.PVM, err error)
	// Get retrieves the PVM from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.PVM, error)
	PVMNamespaceListerExpansion
}

// pVMNamespaceLister implements the PVMNamespaceLister
// interface.
type pVMNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PVMs in the indexer for a given namespace.
func (s pVMNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PVM, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PVM))
	})
	return ret, err
}

// Get retrieves the PVM from the indexer for a given namespace and name.
func (s pVMNamespaceLister) Get(name string) (*v1alpha1.PVM, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("pvm"), name)
	}
	return obj.(*v1alpha1.PVM), nil
}
