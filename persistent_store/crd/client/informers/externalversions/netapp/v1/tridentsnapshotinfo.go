// Copyright 2021 NetApp, Inc. All Rights Reserved.

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	netappv1 "github.com/netapp/trident/persistent_store/crd/apis/netapp/v1"
	versioned "github.com/netapp/trident/persistent_store/crd/client/clientset/versioned"
	internalinterfaces "github.com/netapp/trident/persistent_store/crd/client/informers/externalversions/internalinterfaces"
	v1 "github.com/netapp/trident/persistent_store/crd/client/listers/netapp/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TridentSnapshotInfoInformer provides access to a shared informer and lister for
// TridentSnapshotInfos.
type TridentSnapshotInfoInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.TridentSnapshotInfoLister
}

type tridentSnapshotInfoInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTridentSnapshotInfoInformer constructs a new informer for TridentSnapshotInfo type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTridentSnapshotInfoInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTridentSnapshotInfoInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTridentSnapshotInfoInformer constructs a new informer for TridentSnapshotInfo type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTridentSnapshotInfoInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TridentV1().TridentSnapshotInfos(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TridentV1().TridentSnapshotInfos(namespace).Watch(context.TODO(), options)
			},
		},
		&netappv1.TridentSnapshotInfo{},
		resyncPeriod,
		indexers,
	)
}

func (f *tridentSnapshotInfoInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTridentSnapshotInfoInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *tridentSnapshotInfoInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&netappv1.TridentSnapshotInfo{}, f.defaultInformer)
}

func (f *tridentSnapshotInfoInformer) Lister() v1.TridentSnapshotInfoLister {
	return v1.NewTridentSnapshotInfoLister(f.Informer().GetIndexer())
}