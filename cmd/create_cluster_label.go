package cmd

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/resource"
	"github.com/containership/csctl/resource/label"
)

// createClusterLabelCmd represents the createClusterLabel command
var createClusterLabelCmd = &cobra.Command{
	Use:     "cluster-label",
	Short:   "Create a cluster label",
	Aliases: resource.ClusterLabel().Aliases(),

	Args: cobra.MinimumNArgs(1),

	PreRunE: clusterScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		for _, labelStr := range args {
			key, val, err := label.ParseString(labelStr)
			if err != nil {
				return errors.Wrapf(err, "parsing label %q", labelStr)
			}

			if strings.Contains(key, "/") {
				if !strings.HasPrefix(key, label.ClusterPrefix) {
					return errors.Errorf("parsing label %q: key contains a prefix but it is not %q", labelStr, label.ClusterPrefix)
				}
			} else {
				key = fmt.Sprintf("%s/%s", label.ClusterPrefix, key)
			}

			req := types.ClusterLabel{
				Key:   &key,
				Value: &val,
			}

			resp, err := clientset.API().ClusterLabels(organizationID, clusterID).Create(&req)
			if err != nil {
				return err
			}

			fmt.Printf("Label %s (%s) created successfully\n", resp.ID, labelStr)
		}

		return nil
	},
}

func init() {
	createCmd.AddCommand(createClusterLabelCmd)

	bindCommandToClusterScope(createClusterLabelCmd, false)
}
