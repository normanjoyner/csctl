package provision

import (
	"fmt"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/cloud/rest"
)

// NodePoolLabelsGetter is the getter for labels
type NodePoolLabelsGetter interface {
	NodePoolLabels(organizationID, clusterID, labelPoolID string) NodePoolLabelInterface
}

// NodePoolLabelInterface is the interface for labels
type NodePoolLabelInterface interface {
	Create(*types.NodePoolLabel) (*types.IDResponse, error)
	Get(id string) (*types.NodePoolLabel, error)
	Delete(id string) error
	List() ([]types.NodePoolLabel, error)
}

// labels implements NodePoolLabelInterface
type nodePoolLabels struct {
	client         rest.Interface
	organizationID string
	clusterID      string
	labelPoolID    string
}

func newNodePoolLabels(c *Client, organizationID, clusterID, labelPoolID string) *nodePoolLabels {
	return &nodePoolLabels{
		client:         c.RESTClient(),
		organizationID: organizationID,
		clusterID:      clusterID,
		labelPoolID:    labelPoolID,
	}
}

// Create creates a label
func (l *nodePoolLabels) Create(req *types.NodePoolLabel) (*types.IDResponse, error) {
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s/node-pools/%s/labels", l.organizationID, l.clusterID, l.labelPoolID)
	var out types.IDResponse
	return &out, l.client.Post(path, req, &out)
}

// Get gets a label
func (l *nodePoolLabels) Get(id string) (*types.NodePoolLabel, error) {
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s/node-pools/%s/labels/%s", l.organizationID, l.clusterID, l.labelPoolID, id)
	var out types.NodePoolLabel
	return &out, l.client.Get(path, &out)
}

// Delete deletes a label
func (l *nodePoolLabels) Delete(id string) error {
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s/node-pools/%s/labels/%s", l.organizationID, l.clusterID, l.labelPoolID, id)
	return l.client.Delete(path)
}

// List lists all labels
func (l *nodePoolLabels) List() ([]types.NodePoolLabel, error) {
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s/node-pools/%s/labels", l.organizationID, l.clusterID, l.labelPoolID)
	out := make([]types.NodePoolLabel, 0)
	return out, l.client.Get(path, &out)
}
