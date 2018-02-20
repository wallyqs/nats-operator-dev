// Boilerplate
package fake

import (
	v1alpha2 "github.com/nats-io/nats-kubernetes/operators/nats-server/pkg/apis/nats.io/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNatsClusters implements NatsClusterInterface
type FakeNatsClusters struct {
	Fake *FakeNatsV1alpha2
	ns   string
}

var natsclustersResource = schema.GroupVersionResource{Group: "nats.io", Version: "v1alpha2", Resource: "natsclusters"}

var natsclustersKind = schema.GroupVersionKind{Group: "nats.io", Version: "v1alpha2", Kind: "NatsCluster"}

// Get takes name of the natsCluster, and returns the corresponding natsCluster object, and an error if there is any.
func (c *FakeNatsClusters) Get(name string, options v1.GetOptions) (result *v1alpha2.NatsCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(natsclustersResource, c.ns, name), &v1alpha2.NatsCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.NatsCluster), err
}

// List takes label and field selectors, and returns the list of NatsClusters that match those selectors.
func (c *FakeNatsClusters) List(opts v1.ListOptions) (result *v1alpha2.NatsClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(natsclustersResource, natsclustersKind, c.ns, opts), &v1alpha2.NatsClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.NatsClusterList{}
	for _, item := range obj.(*v1alpha2.NatsClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested natsClusters.
func (c *FakeNatsClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(natsclustersResource, c.ns, opts))

}

// Create takes the representation of a natsCluster and creates it.  Returns the server's representation of the natsCluster, and an error, if there is any.
func (c *FakeNatsClusters) Create(natsCluster *v1alpha2.NatsCluster) (result *v1alpha2.NatsCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(natsclustersResource, c.ns, natsCluster), &v1alpha2.NatsCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.NatsCluster), err
}

// Update takes the representation of a natsCluster and updates it. Returns the server's representation of the natsCluster, and an error, if there is any.
func (c *FakeNatsClusters) Update(natsCluster *v1alpha2.NatsCluster) (result *v1alpha2.NatsCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(natsclustersResource, c.ns, natsCluster), &v1alpha2.NatsCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.NatsCluster), err
}

// Delete takes name of the natsCluster and deletes it. Returns an error if one occurs.
func (c *FakeNatsClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(natsclustersResource, c.ns, name), &v1alpha2.NatsCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNatsClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(natsclustersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha2.NatsClusterList{})
	return err
}

// Patch applies the patch and returns the patched natsCluster.
func (c *FakeNatsClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.NatsCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(natsclustersResource, c.ns, name, data, subresources...), &v1alpha2.NatsCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.NatsCluster), err
}
