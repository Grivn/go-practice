package cmd

import (
	"fmt"
	practiceClient "github.com/Grivn/go-practice/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strconv"
)

// runCmd represents the run command
var requestCmd = &cobra.Command{
	Use:   "request [id]",
	Short: "Run replica instance",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("failed to parse replica ID "+
					"from positional argument: %s", err)
			}
			viper.Set("replica.id", id)
		}

		return request()
	},
}

func init() {
	rootCmd.AddCommand(requestCmd)
}

func request() error {
	client := practiceClient.New()
	client.Start()

	return nil
}
