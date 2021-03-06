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

// This file was automatically generated by informer-gen

package v1alpha1

import (
	persistentvm_v1alpha1 "github.com/kubevirt-incubator/persistentvm-addon/pkg/apis/persistentvm/v1alpha1"
	versioned "github.com/kubevirt-incubator/persistentvm-addon/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kubevirt-incubator/persistentvm-addon/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kubevirt-incubator/persistentvm-addon/pkg/client/listers/persistentvm/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// PVMInformer provides access to a shared informer and lister for
// PVMs.
type PVMInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PVMLister
}

type pVMInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewPVMInformer constructs a new informer for PVM type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPVMInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.PersistentvmV1alpha1().PVMs(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.PersistentvmV1alpha1().PVMs(namespace).Watch(options)
			},
		},
		&persistentvm_v1alpha1.PVM{},
		resyncPeriod,
		indexers,
	)
}

func defaultPVMInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewPVMInformer(client, v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *pVMInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&persistentvm_v1alpha1.PVM{}, defaultPVMInformer)
}

func (f *pVMInformer) Lister() v1alpha1.PVMLister {
	return v1alpha1.NewPVMLister(f.Informer().GetIndexer())
}
