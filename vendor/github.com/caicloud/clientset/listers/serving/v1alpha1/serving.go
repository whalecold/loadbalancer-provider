/*
Copyright 2019 caicloud authors. All rights reserved.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/caicloud/clientset/pkg/apis/serving/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ServingLister helps list Servings.
type ServingLister interface {
	// List lists all Servings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Serving, err error)
	// Servings returns an object that can list and get Servings.
	Servings(namespace string) ServingNamespaceLister
	ServingListerExpansion
}

// servingLister implements the ServingLister interface.
type servingLister struct {
	indexer cache.Indexer
}

// NewServingLister returns a new ServingLister.
func NewServingLister(indexer cache.Indexer) ServingLister {
	return &servingLister{indexer: indexer}
}

// List lists all Servings in the indexer.
func (s *servingLister) List(selector labels.Selector) (ret []*v1alpha1.Serving, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Serving))
	})
	return ret, err
}

// Servings returns an object that can list and get Servings.
func (s *servingLister) Servings(namespace string) ServingNamespaceLister {
	return servingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServingNamespaceLister helps list and get Servings.
type ServingNamespaceLister interface {
	// List lists all Servings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Serving, err error)
	// Get retrieves the Serving from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Serving, error)
	ServingNamespaceListerExpansion
}

// servingNamespaceLister implements the ServingNamespaceLister
// interface.
type servingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Servings in the indexer for a given namespace.
func (s servingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Serving, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Serving))
	})
	return ret, err
}

// Get retrieves the Serving from the indexer for a given namespace and name.
func (s servingNamespaceLister) Get(name string) (*v1alpha1.Serving, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("serving"), name)
	}
	return obj.(*v1alpha1.Serving), nil
}
