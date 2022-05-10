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

package v1

import (
	"context"
	"time"

	v1 "github.com/kubeovn/kube-ovn/pkg/apis/kubeovn/v1"
	scheme "github.com/kubeovn/kube-ovn/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// IptablesDnatRulesGetter has a method to return a IptablesDnatRuleInterface.
// A group's client should implement this interface.
type IptablesDnatRulesGetter interface {
	IptablesDnatRules() IptablesDnatRuleInterface
}

// IptablesDnatRuleInterface has methods to work with IptablesDnatRule resources.
type IptablesDnatRuleInterface interface {
	Create(ctx context.Context, iptablesDnatRule *v1.IptablesDnatRule, opts metav1.CreateOptions) (*v1.IptablesDnatRule, error)
	Update(ctx context.Context, iptablesDnatRule *v1.IptablesDnatRule, opts metav1.UpdateOptions) (*v1.IptablesDnatRule, error)
	UpdateStatus(ctx context.Context, iptablesDnatRule *v1.IptablesDnatRule, opts metav1.UpdateOptions) (*v1.IptablesDnatRule, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.IptablesDnatRule, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.IptablesDnatRuleList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.IptablesDnatRule, err error)
	IptablesDnatRuleExpansion
}

// iptablesDnatRules implements IptablesDnatRuleInterface
type iptablesDnatRules struct {
	client rest.Interface
}

// newIptablesDnatRules returns a IptablesDnatRules
func newIptablesDnatRules(c *KubeovnV1Client) *iptablesDnatRules {
	return &iptablesDnatRules{
		client: c.RESTClient(),
	}
}

// Get takes name of the iptablesDnatRule, and returns the corresponding iptablesDnatRule object, and an error if there is any.
func (c *iptablesDnatRules) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.IptablesDnatRule, err error) {
	result = &v1.IptablesDnatRule{}
	err = c.client.Get().
		Resource("iptables-dnat-rules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IptablesDnatRules that match those selectors.
func (c *iptablesDnatRules) List(ctx context.Context, opts metav1.ListOptions) (result *v1.IptablesDnatRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.IptablesDnatRuleList{}
	err = c.client.Get().
		Resource("iptables-dnat-rules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iptablesDnatRules.
func (c *iptablesDnatRules) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("iptables-dnat-rules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a iptablesDnatRule and creates it.  Returns the server's representation of the iptablesDnatRule, and an error, if there is any.
func (c *iptablesDnatRules) Create(ctx context.Context, iptablesDnatRule *v1.IptablesDnatRule, opts metav1.CreateOptions) (result *v1.IptablesDnatRule, err error) {
	result = &v1.IptablesDnatRule{}
	err = c.client.Post().
		Resource("iptables-dnat-rules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(iptablesDnatRule).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a iptablesDnatRule and updates it. Returns the server's representation of the iptablesDnatRule, and an error, if there is any.
func (c *iptablesDnatRules) Update(ctx context.Context, iptablesDnatRule *v1.IptablesDnatRule, opts metav1.UpdateOptions) (result *v1.IptablesDnatRule, err error) {
	result = &v1.IptablesDnatRule{}
	err = c.client.Put().
		Resource("iptables-dnat-rules").
		Name(iptablesDnatRule.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(iptablesDnatRule).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *iptablesDnatRules) UpdateStatus(ctx context.Context, iptablesDnatRule *v1.IptablesDnatRule, opts metav1.UpdateOptions) (result *v1.IptablesDnatRule, err error) {
	result = &v1.IptablesDnatRule{}
	err = c.client.Put().
		Resource("iptables-dnat-rules").
		Name(iptablesDnatRule.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(iptablesDnatRule).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the iptablesDnatRule and deletes it. Returns an error if one occurs.
func (c *iptablesDnatRules) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("iptables-dnat-rules").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iptablesDnatRules) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("iptables-dnat-rules").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched iptablesDnatRule.
func (c *iptablesDnatRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.IptablesDnatRule, err error) {
	result = &v1.IptablesDnatRule{}
	err = c.client.Patch(pt).
		Resource("iptables-dnat-rules").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
