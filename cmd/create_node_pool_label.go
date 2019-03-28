package cmd

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/resource"
	"github.com/containership/csctl/resource/label"
)

// createNodePoolLabelCmd represents the createNodePoolLabel command
var createNodePoolLabelCmd = &cobra.Command{
	Use:     "node-pool-label",
	Short:   "Create a node pool label",
	Aliases: resource.NodePoolLabel().Aliases(),

	Args: cobra.MinimumNArgs(1),

	PreRunE: nodePoolScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		for _, labelStr := range args {
			key, val, err := label.ParseString(labelStr)
			if err != nil {
				return errors.Wrapf(err, "parsing label %q", labelStr)
			}

			if strings.Contains(key, "/") {
				if !strings.HasPrefix(key, label.NodePoolPrefix) {
					return errors.Errorf("parsing label %q: key contains a prefix but it is not %q", labelStr, label.NodePoolPrefix)
				}
			} else {
				key = fmt.Sprintf("%s/%s", label.NodePoolPrefix, key)
			}

			req := types.NodePoolLabel{
				Key:   &key,
				Value: &val,
			}

			resp, err := clientset.Provision().NodePoolLabels(organizationID, clusterID, nodePoolID).Create(&req)
			if err != nil {
				return err
			}

			fmt.Printf("Label %s (%s) created successfully\n", resp.ID, labelStr)
		}

		return nil
	},
}

func init() {
	createCmd.AddCommand(createNodePoolLabelCmd)

	bindCommandToNodePoolScope(createNodePoolLabelCmd, false)
}
