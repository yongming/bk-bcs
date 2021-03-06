/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "bk-bcs/bcs-k8s/tkex-statefulsetplus-operator/pkg/apis/tkex/v1alpha1"
	scheme "bk-bcs/bcs-k8s/tkex-statefulsetplus-operator/pkg/clientset/internalclientset/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	autoscaling "k8s.io/kubernetes/pkg/apis/autoscaling"
)

// StatefulSetPlusesGetter has a method to return a StatefulSetPlusInterface.
// A group's client should implement this interface.
type StatefulSetPlusesGetter interface {
	StatefulSetPluses(namespace string) StatefulSetPlusInterface
}

// StatefulSetPlusInterface has methods to work with StatefulSetPlus resources.
type StatefulSetPlusInterface interface {
	Create(*v1alpha1.StatefulSetPlus) (*v1alpha1.StatefulSetPlus, error)
	Update(*v1alpha1.StatefulSetPlus) (*v1alpha1.StatefulSetPlus, error)
	UpdateStatus(*v1alpha1.StatefulSetPlus) (*v1alpha1.StatefulSetPlus, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.StatefulSetPlus, error)
	List(opts v1.ListOptions) (*v1alpha1.StatefulSetPlusList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StatefulSetPlus, err error)
	GetScale(statefulSetPlusName string, options v1.GetOptions) (*autoscaling.Scale, error)
	UpdateScale(statefulSetPlusName string, scale *autoscaling.Scale) (*autoscaling.Scale, error)

	StatefulSetPlusExpansion
}

// statefulSetPluses implements StatefulSetPlusInterface
type statefulSetPluses struct {
	client rest.Interface
	ns     string
}

// newStatefulSetPluses returns a StatefulSetPluses
func newStatefulSetPluses(c *TkexV1alpha1Client, namespace string) *statefulSetPluses {
	return &statefulSetPluses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the statefulSetPlus, and returns the corresponding statefulSetPlus object, and an error if there is any.
func (c *statefulSetPluses) Get(name string, options v1.GetOptions) (result *v1alpha1.StatefulSetPlus, err error) {
	result = &v1alpha1.StatefulSetPlus{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("statefulsetpluses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of StatefulSetPluses that match those selectors.
func (c *statefulSetPluses) List(opts v1.ListOptions) (result *v1alpha1.StatefulSetPlusList, err error) {
	result = &v1alpha1.StatefulSetPlusList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("statefulsetpluses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested statefulSetPluses.
func (c *statefulSetPluses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("statefulsetpluses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a statefulSetPlus and creates it.  Returns the server's representation of the statefulSetPlus, and an error, if there is any.
func (c *statefulSetPluses) Create(statefulSetPlus *v1alpha1.StatefulSetPlus) (result *v1alpha1.StatefulSetPlus, err error) {
	result = &v1alpha1.StatefulSetPlus{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("statefulsetpluses").
		Body(statefulSetPlus).
		Do().
		Into(result)
	return
}

// Update takes the representation of a statefulSetPlus and updates it. Returns the server's representation of the statefulSetPlus, and an error, if there is any.
func (c *statefulSetPluses) Update(statefulSetPlus *v1alpha1.StatefulSetPlus) (result *v1alpha1.StatefulSetPlus, err error) {
	result = &v1alpha1.StatefulSetPlus{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("statefulsetpluses").
		Name(statefulSetPlus.Name).
		Body(statefulSetPlus).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *statefulSetPluses) UpdateStatus(statefulSetPlus *v1alpha1.StatefulSetPlus) (result *v1alpha1.StatefulSetPlus, err error) {
	result = &v1alpha1.StatefulSetPlus{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("statefulsetpluses").
		Name(statefulSetPlus.Name).
		SubResource("status").
		Body(statefulSetPlus).
		Do().
		Into(result)
	return
}

// Delete takes name of the statefulSetPlus and deletes it. Returns an error if one occurs.
func (c *statefulSetPluses) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("statefulsetpluses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *statefulSetPluses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("statefulsetpluses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched statefulSetPlus.
func (c *statefulSetPluses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StatefulSetPlus, err error) {
	result = &v1alpha1.StatefulSetPlus{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("statefulsetpluses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

// GetScale takes name of the statefulSetPlus, and returns the corresponding autoscaling.Scale object, and an error if there is any.
func (c *statefulSetPluses) GetScale(statefulSetPlusName string, options v1.GetOptions) (result *autoscaling.Scale, err error) {
	result = &autoscaling.Scale{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("statefulsetpluses").
		Name(statefulSetPlusName).
		SubResource("scale").
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// UpdateScale takes the top resource name and the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *statefulSetPluses) UpdateScale(statefulSetPlusName string, scale *autoscaling.Scale) (result *autoscaling.Scale, err error) {
	result = &autoscaling.Scale{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("statefulsetpluses").
		Name(statefulSetPlusName).
		SubResource("scale").
		Body(scale).
		Do().
		Into(result)
	return
}
