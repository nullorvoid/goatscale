package consulapi

import (
	"os"

	consul "github.com/hashicorp/consul/api"
)

// Client holds an internal reference to the consul api that allows registration and deregistration of a service
type Client struct {
	consul *consul.Client
}

// NewConsulClient returns a Client that allows to register and deregister
func NewConsulClient(addr string) (*Client, error) {
	config := consul.DefaultConfig()
	config.Address = addr
	c, err := consul.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Client{consul: c}, nil
}

// Register a service with consul
func (c *Client) Register(name string, addr string, port int) error {
	host, _ := os.Hostname()

	reg := &consul.AgentServiceRegistration{
		// ID defined has to be unique so we will append the host name.
		ID: name + "_" + host,
		// Name is the name of the service that this client provides
		Name:    name,
		Address: addr,
		Port:    port,
		// Enable service in Traefik using Tags
		Tags: []string{
			"traefik.backend=" + name,
			"traefik.enabled=true",
		},
	}

	return c.consul.Agent().ServiceRegister(reg)
}

// Deregister a service with consul
func (c *Client) Deregister(id string) error {
	return c.consul.Agent().ServiceDeregister(id)
}
