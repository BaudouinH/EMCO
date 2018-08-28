// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2020 Intel Corporation
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/open-ness/EMCO/src/monitor/pkg/apis/k8splugin/v1alpha1"
	scheme "github.com/open-ness/EMCO/src/monitor/pkg/generated/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ResourceBundleStatesGetter has a method to return a ResourceBundleStateInterface.
// A group's client should implement this interface.
type ResourceBundleStatesGetter interface {
	ResourceBundleStates(namespace string) ResourceBundleStateInterface
}

// ResourceBundleStateInterface has methods to work with ResourceBundleState resources.
type ResourceBundleStateInterface interface {
	Create(*v1alpha1.ResourceBundleState) (*v1alpha1.ResourceBundleState, error)
	Update(*v1alpha1.ResourceBundleState) (*v1alpha1.ResourceBundleState, error)
	UpdateStatus(*v1alpha1.ResourceBundleState) (*v1alpha1.ResourceBundleState, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ResourceBundleState, error)
	List(opts v1.ListOptions) (*v1alpha1.ResourceBundleStateList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ResourceBundleState, err error)
	ResourceBundleStateExpansion
}

// resourceBundleStates implements ResourceBundleStateInterface
type resourceBundleStates struct {
	client rest.Interface
	ns     string
}

// newResourceBundleStates returns a ResourceBundleStates
func newResourceBundleStates(c *K8spluginV1alpha1Client, namespace string) *resourceBundleStates {
	return &resourceBundleStates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the resourceBundleState, and returns the corresponding resourceBundleState object, and an error if there is any.
func (c *resourceBundleStates) Get(name string, options v1.GetOptions) (result *v1alpha1.ResourceBundleState, err error) {
	result = &v1alpha1.ResourceBundleState{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resourcebundlestates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ResourceBundleStates that match those selectors.
func (c *resourceBundleStates) List(opts v1.ListOptions) (result *v1alpha1.ResourceBundleStateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ResourceBundleStateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resourcebundlestates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.TODO()).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested resourceBundleStates.
func (c *resourceBundleStates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("resourcebundlestates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(context.TODO())
}

// Create takes the representation of a resourceBundleState and creates it.  Returns the server's representation of the resourceBundleState, and an error, if there is any.
func (c *resourceBundleStates) Create(resourceBundleState *v1alpha1.ResourceBundleState) (result *v1alpha1.ResourceBundleState, err error) {
	result = &v1alpha1.ResourceBundleState{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("resourcebundlestates").
		Body(resourceBundleState).
		Do(context.TODO()).
		Into(result)
	return
}

// Update takes the representation of a resourceBundleState and updates it. Returns the server's representation of the resourceBundleState, and an error, if there is any.
func (c *resourceBundleStates) Update(resourceBundleState *v1alpha1.ResourceBundleState) (result *v1alpha1.ResourceBundleState, err error) {
	result = &v1alpha1.ResourceBundleState{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("resourcebundlestates").
		Name(resourceBundleState.Name).
		Body(resourceBundleState).
		Do(context.TODO()).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *resourceBundleStates) UpdateStatus(resourceBundleState *v1alpha1.ResourceBundleState) (result *v1alpha1.ResourceBundleState, err error) {
	result = &v1alpha1.ResourceBundleState{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("resourcebundlestates").
		Name(resourceBundleState.Name).
		SubResource("status").
		Body(resourceBundleState).
		Do(context.TODO()).
		Into(result)
	return
}

// Delete takes name of the resourceBundleState and deletes it. Returns an error if one occurs.
func (c *resourceBundleStates) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resourcebundlestates").
		Name(name).
		Body(options).
		Do(context.TODO()).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *resourceBundleStates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resourcebundlestates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do(context.TODO()).
		Error()
}

// Patch applies the patch and returns the patched resourceBundleState.
func (c *resourceBundleStates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ResourceBundleState, err error) {
	result = &v1alpha1.ResourceBundleState{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("resourcebundlestates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.TODO()).
		Into(result)
	return
}