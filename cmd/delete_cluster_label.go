package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource"
)

// deleteClusterLabelCmd represents the deleteClusterLabel command
var deleteClusterLabelCmd = &cobra.Command{
	Use:     "cluster-label",
	Short:   "Delete one or more cluster labels",
	Aliases: resource.ClusterLabel().Aliases(),

	Args: cobra.MinimumNArgs(1),

	PreRunE: clusterScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		for _, id := range args {
			err := clientset.API().ClusterLabels(organizationID, clusterID).Delete(id)
			if err != nil {
				return err
			}

			fmt.Printf("Cluster label %s deleted successfully\n", id)
		}

		return nil
	},
}

func init() {
	deleteCmd.AddCommand(deleteClusterLabelCmd)

	bindCommandToClusterScope(deleteClusterLabelCmd, false)
}
