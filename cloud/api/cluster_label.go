package api

import (
	"fmt"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/cloud/rest"
)

// ClusterLabelsGetter is the getter for labels
type ClusterLabelsGetter interface {
	ClusterLabels(organizationID, clusterID string) ClusterLabelInterface
}

// ClusterLabelInterface is the interface for labels
type ClusterLabelInterface interface {
	Create(*types.ClusterLabel) (*types.IDResponse, error)
	Get(id string) (*types.ClusterLabel, error)
	Delete(id string) error
	List() ([]types.ClusterLabel, error)
}

// labels implements ClusterLabelInterface
type clusterLabels struct {
	client         rest.Interface
	organizationID string
	clusterID      string
}

func newClusterLabels(c *Client, organizationID, clusterID string) *clusterLabels {
	return &clusterLabels{
		client:         c.RESTClient(),
		organizationID: organizationID,
		clusterID:      clusterID,
	}
}

// Create creates a label
func (l *clusterLabels) Create(req *types.ClusterLabel) (*types.IDResponse, error) {
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s/labels", l.organizationID, l.clusterID)
	var out types.IDResponse
	return &out, l.client.Post(path, req, &out)
}

// Get gets a label
func (l *clusterLabels) Get(id string) (*types.ClusterLabel, error) {
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s/labels/%s", l.organizationID, l.clusterID, id)
	var out types.ClusterLabel
	return &out, l.client.Get(path, &out)
}

// Delete deletes a label
func (l *clusterLabels) Delete(id string) error {
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s/labels/%s", l.organizationID, l.clusterID, id)
	return l.client.Delete(path)
}

// List lists all labels
func (l *clusterLabels) List() ([]types.ClusterLabel, error) {
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s/labels", l.organizationID, l.clusterID)
	out := make([]types.ClusterLabel, 0)
	return out, l.client.Get(path, &out)
}
