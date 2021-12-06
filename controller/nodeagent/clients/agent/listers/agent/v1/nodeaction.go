/*
Copyright 2021.

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

package v1

import (
	v1 "nodeagent/apis/agent/v1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NodeActionLister helps list NodeActions.
// All objects returned here must be treated as read-only.
type NodeActionLister interface {
	// List lists all NodeActions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.NodeAction, err error)
	// NodeActions returns an object that can list and get NodeActions.
	NodeActions(namespace string) NodeActionNamespaceLister
	NodeActionListerExpansion
}

// nodeActionLister implements the NodeActionLister interface.
type nodeActionLister struct {
	indexer cache.Indexer
}

// NewNodeActionLister returns a new NodeActionLister.
func NewNodeActionLister(indexer cache.Indexer) NodeActionLister {
	return &nodeActionLister{indexer: indexer}
}

// List lists all NodeActions in the indexer.
func (s *nodeActionLister) List(selector labels.Selector) (ret []*v1.NodeAction, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.NodeAction))
	})
	return ret, err
}

// NodeActions returns an object that can list and get NodeActions.
func (s *nodeActionLister) NodeActions(namespace string) NodeActionNamespaceLister {
	return nodeActionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NodeActionNamespaceLister helps list and get NodeActions.
// All objects returned here must be treated as read-only.
type NodeActionNamespaceLister interface {
	// List lists all NodeActions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.NodeAction, err error)
	// Get retrieves the NodeAction from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.NodeAction, error)
	NodeActionNamespaceListerExpansion
}

// nodeActionNamespaceLister implements the NodeActionNamespaceLister
// interface.
type nodeActionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NodeActions in the indexer for a given namespace.
func (s nodeActionNamespaceLister) List(selector labels.Selector) (ret []*v1.NodeAction, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.NodeAction))
	})
	return ret, err
}

// Get retrieves the NodeAction from the indexer for a given namespace and name.
func (s nodeActionNamespaceLister) Get(name string) (*v1.NodeAction, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("nodeaction"), name)
	}
	return obj.(*v1.NodeAction), nil
}
