package cmd

import (
	"fmt"
	"github.com/Grivn/go-practice/config"
	practiceServer "github.com/Grivn/go-practice/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strconv"
)

const (
	defConsensusCfgFile = "consensus.yaml"
	defKeysFile         = "keys.yaml"
	defUsigEnclaveFile  = "libusig.signed.so"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [id]",
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

		return run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().Int("id", 0, "ID of replica")
	must(viper.BindPFlag("replica.id",
		runCmd.Flags().Lookup("id")))

	runCmd.Flags().StringP("usig-enclave-file", "u",
		defUsigEnclaveFile, "USIG enclave file")
	must(viper.BindPFlag("usig.enclaveFile",
		runCmd.Flags().Lookup("usig-enclave-file")))
}

func run() error {
	cfg := config.New()
	cfg.LoadConfig(viper.GetString("consensusConf"))

	server := practiceServer.New()
	server.Start()

	return nil
}