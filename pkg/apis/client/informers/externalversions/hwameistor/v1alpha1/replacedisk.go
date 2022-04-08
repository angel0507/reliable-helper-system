/*
Copyright 2022 Contributors to The HwameiStor project.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	versioned "github.com/hwameistor/improved-system/pkg/apis/client/clientset/versioned"
	internalinterfaces "github.com/hwameistor/improved-system/pkg/apis/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/hwameistor/improved-system/pkg/apis/client/listers/hwameistor/v1alpha1"
	hwameistorv1alpha1 "github.com/hwameistor/improved-system/pkg/apis/hwameistor/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ReplaceDiskInformer provides access to a shared informer and lister for
// ReplaceDisks.
type ReplaceDiskInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ReplaceDiskLister
}

type replaceDiskInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewReplaceDiskInformer constructs a new informer for ReplaceDisk type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewReplaceDiskInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredReplaceDiskInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredReplaceDiskInformer constructs a new informer for ReplaceDisk type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredReplaceDiskInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HwameistorV1alpha1().ReplaceDisks().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HwameistorV1alpha1().ReplaceDisks().Watch(context.TODO(), options)
			},
		},
		&hwameistorv1alpha1.ReplaceDisk{},
		resyncPeriod,
		indexers,
	)
}

func (f *replaceDiskInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredReplaceDiskInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *replaceDiskInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&hwameistorv1alpha1.ReplaceDisk{}, f.defaultInformer)
}

func (f *replaceDiskInformer) Lister() v1alpha1.ReplaceDiskLister {
	return v1alpha1.NewReplaceDiskLister(f.Informer().GetIndexer())
}
