package consul

import (
	"fmt"

	"github.com/abronan/valkeyrie/store"
	"github.com/abronan/valkeyrie/store/consul"
	"github.com/containous/traefik/old/provider"
	"github.com/containous/traefik/old/provider/kv"
	"github.com/containous/traefik/old/types"
	"github.com/containous/traefik/pkg/safe"
)

var _ provider.Provider = (*Provider)(nil)

// Provider holds configurations of the p.
type Provider struct {
	kv.Provider `mapstructure:",squash" export:"true"`
}

// Init the provider
func (p *Provider) Init(constraints types.Constraints) error {
	err := p.Provider.Init(constraints)
	if err != nil {
		return err
	}

	store, err := p.CreateStore()
	if err != nil {
		return fmt.Errorf("failed to Connect to KV store: %v", err)
	}

	p.SetKVClient(store)
	return nil
}

// Provide allows the consul provider to provide configurations to traefik
// using the given configuration channel.
func (p *Provider) Provide(configurationChan chan<- types.ConfigMessage, pool *safe.Pool) error {
	return p.Provider.Provide(configurationChan, pool)
}

// CreateStore creates the KV store
func (p *Provider) CreateStore() (store.Store, error) {
	p.SetStoreType(store.CONSUL)
	consul.Register()
	return p.Provider.CreateStore()
}
