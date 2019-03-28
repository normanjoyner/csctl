package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource"
)

// deleteNodePoolLabelCmd represents the deleteNodePoolLabel command
var deleteNodePoolLabelCmd = &cobra.Command{
	Use:     "node-pool-label",
	Short:   "Delete one or more node pool labels",
	Aliases: resource.NodePoolLabel().Aliases(),

	Args: cobra.MinimumNArgs(1),

	PreRunE: nodePoolScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		for _, id := range args {
			err := clientset.Provision().NodePoolLabels(organizationID, clusterID, nodePoolID).Delete(id)
			if err != nil {
				return err
			}

			fmt.Printf("Node pool label %s deleted successfully\n", id)
		}

		return nil
	},
}

func init() {
	deleteCmd.AddCommand(deleteNodePoolLabelCmd)

	bindCommandToNodePoolScope(deleteNodePoolLabelCmd, false)
}
