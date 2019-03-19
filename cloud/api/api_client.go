package api

import (
	"github.com/pkg/errors"

	"github.com/containership/csctl/cloud/rest"
)

const (
	defaultBaseURL = "https://api.containership.io"
)

// Interface is the interface for API
type Interface interface {
	RESTClient() rest.Interface
	AccessTokensGetter
	AccountGetter
	ClustersGetter
	OrganizationsGetter
	ProvidersGetter
	PluginsGetter
	PluginCatalogGetter
	UsersGetter
	RegistryGetter
}

// Client is the API client
type Client struct {
	name       string
	restClient *rest.Client
}

// New constructs a new API client
func New(cfg rest.Config) (*Client, error) {
	if cfg.BaseURL == "" {
		cfg.BaseURL = defaultBaseURL
	}

	restClient, err := rest.NewClient(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "constructing REST client")
	}

	return &Client{
		name:       "API",
		restClient: restClient,
	}, nil
}

// RESTClient returns the REST client associated with this client
func (c *Client) RESTClient() rest.Interface {
	return c.restClient
}

// AccessTokens returns the access token interface
func (c *Client) AccessTokens() AccessTokenInterface {
	return newAccessTokens(c)
}

// Account returns the account interface
func (c *Client) Account() AccountInterface {
	return newAccount(c)
}

// Clusters returns the clusters interface
func (c *Client) Clusters(organizationID string) ClusterInterface {
	return newClusters(c, organizationID)
}

// Organizations returns the organizations interface
func (c *Client) Organizations() OrganizationInterface {
	return newOrganizations(c)
}

// Providers returns the providers interface
func (c *Client) Providers(organizationID string) ProviderInterface {
	return newProviders(c, organizationID)
}

// Plugins returns the plugins interface
func (c *Client) Plugins(organizationID, clusterID string) PluginInterface {
	return newPlugins(c, organizationID, clusterID)
}

// PluginCatalog returns the pluginsCatalog interface
func (c *Client) PluginCatalog() PluginCatalogInterface {
	return newPluginCatalog(c)
}

// Users returns the users interface
func (c *Client) Users(organizationID string) UserInterface {
	return newUsers(c, organizationID)
}

// Registries returns the registries interface
func (c *Client) Registries(organizationID string) RegistryInterface {
	return newRegistry(c, organizationID)
}
