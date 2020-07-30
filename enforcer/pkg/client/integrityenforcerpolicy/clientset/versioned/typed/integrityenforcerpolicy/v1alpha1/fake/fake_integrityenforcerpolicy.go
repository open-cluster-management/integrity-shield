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

package fake

import (
	v1alpha1 "github.com/IBM/integrity-enforcer/enforcer/pkg/apis/integrityenforcerpolicy/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIntegrityEnforcerPolicies implements IntegrityEnforcerPolicyInterface
type FakeIntegrityEnforcerPolicies struct {
	Fake *FakeResearchV1alpha1
	ns   string
}

var integrityenforcerpoliciesResource = schema.GroupVersionResource{Group: "research.ibm.com", Version: "v1alpha1", Resource: "integrityenforcerpolicies"}

var integrityenforcerpoliciesKind = schema.GroupVersionKind{Group: "research.ibm.com", Version: "v1alpha1", Kind: "IntegrityEnforcerPolicy"}

// Get takes name of the integrityEnforcerPolicy, and returns the corresponding integrityEnforcerPolicy object, and an error if there is any.
func (c *FakeIntegrityEnforcerPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.IntegrityEnforcerPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(integrityenforcerpoliciesResource, c.ns, name), &v1alpha1.IntegrityEnforcerPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IntegrityEnforcerPolicy), err
}

// List takes label and field selectors, and returns the list of IntegrityEnforcerPolicies that match those selectors.
func (c *FakeIntegrityEnforcerPolicies) List(opts v1.ListOptions) (result *v1alpha1.IntegrityEnforcerPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(integrityenforcerpoliciesResource, integrityenforcerpoliciesKind, c.ns, opts), &v1alpha1.IntegrityEnforcerPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IntegrityEnforcerPolicyList{ListMeta: obj.(*v1alpha1.IntegrityEnforcerPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.IntegrityEnforcerPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested integrityEnforcerPolicies.
func (c *FakeIntegrityEnforcerPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(integrityenforcerpoliciesResource, c.ns, opts))

}

// Create takes the representation of a integrityEnforcerPolicy and creates it.  Returns the server's representation of the integrityEnforcerPolicy, and an error, if there is any.
func (c *FakeIntegrityEnforcerPolicies) Create(integrityEnforcerPolicy *v1alpha1.IntegrityEnforcerPolicy) (result *v1alpha1.IntegrityEnforcerPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(integrityenforcerpoliciesResource, c.ns, integrityEnforcerPolicy), &v1alpha1.IntegrityEnforcerPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IntegrityEnforcerPolicy), err
}

// Update takes the representation of a integrityEnforcerPolicy and updates it. Returns the server's representation of the integrityEnforcerPolicy, and an error, if there is any.
func (c *FakeIntegrityEnforcerPolicies) Update(integrityEnforcerPolicy *v1alpha1.IntegrityEnforcerPolicy) (result *v1alpha1.IntegrityEnforcerPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(integrityenforcerpoliciesResource, c.ns, integrityEnforcerPolicy), &v1alpha1.IntegrityEnforcerPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IntegrityEnforcerPolicy), err
}

// Delete takes name of the integrityEnforcerPolicy and deletes it. Returns an error if one occurs.
func (c *FakeIntegrityEnforcerPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(integrityenforcerpoliciesResource, c.ns, name), &v1alpha1.IntegrityEnforcerPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIntegrityEnforcerPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(integrityenforcerpoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.IntegrityEnforcerPolicyList{})
	return err
}

// Patch applies the patch and returns the patched integrityEnforcerPolicy.
func (c *FakeIntegrityEnforcerPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IntegrityEnforcerPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(integrityenforcerpoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.IntegrityEnforcerPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IntegrityEnforcerPolicy), err
}
