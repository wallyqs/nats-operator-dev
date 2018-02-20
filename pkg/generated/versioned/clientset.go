// Boilerplate
package versioned

import (
	glog "github.com/golang/glog"
	natsv1alpha2 "github.com/nats-io/nats-kubernetes/operators/nats-server/pkg/generated/versioned/typed/nats.io/v1alpha2"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	NatsV1alpha2() natsv1alpha2.NatsV1alpha2Interface
	// Deprecated: please explicitly pick a version if possible.
	Nats() natsv1alpha2.NatsV1alpha2Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	natsV1alpha2 *natsv1alpha2.NatsV1alpha2Client
}

// NatsV1alpha2 retrieves the NatsV1alpha2Client
func (c *Clientset) NatsV1alpha2() natsv1alpha2.NatsV1alpha2Interface {
	return c.natsV1alpha2
}

// Deprecated: Nats retrieves the default version of NatsClient.
// Please explicitly pick a version.
func (c *Clientset) Nats() natsv1alpha2.NatsV1alpha2Interface {
	return c.natsV1alpha2
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.natsV1alpha2, err = natsv1alpha2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		glog.Errorf("failed to create the DiscoveryClient: %v", err)
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.natsV1alpha2 = natsv1alpha2.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.natsV1alpha2 = natsv1alpha2.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
