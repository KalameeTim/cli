package root

import (
	"github.com/debricked/cli/pkg/client"
	"github.com/debricked/cli/pkg/cmd/files"
	"github.com/debricked/cli/pkg/cmd/report"
	"github.com/debricked/cli/pkg/cmd/scan"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var accessToken string

const AccessTokenFlag = "access-token"

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "debricked",
		Short: "Debricked CLI - Keep track of your dependencies!",
		Long: `A fast and flexible software composition analysis CLI tool, given to you by Debricked.
Complete documentation is available at https://debricked.com/docs/integrations/cli.html#debricked-cli`,
		PreRun: func(cmd *cobra.Command, _ []string) {
			_ = viper.BindPFlags(cmd.PersistentFlags())
		},
	}
	viper.SetEnvPrefix("DEBRICKED")
	viper.MustBindEnv(AccessTokenFlag)
	rootCmd.PersistentFlags().StringVarP(
		&accessToken,
		AccessTokenFlag,
		"t",
		"",
		`Debricked access token. 
Read more: https://debricked.com/docs/administration/access-tokens.html`,
	)

	var debClient client.IDebClient = client.NewDebClient(&accessToken)
	rootCmd.AddCommand(report.NewReportCmd(&debClient))
	rootCmd.AddCommand(files.NewFilesCmd(&debClient))
	rootCmd.AddCommand(scan.NewScanCmd(&debClient))

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	return rootCmd
}
