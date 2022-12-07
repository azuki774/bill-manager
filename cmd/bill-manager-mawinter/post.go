package main

import (
	"context"
	"os"

	"azuki774/bill-manager/internal/factory"

	"github.com/spf13/cobra"
)

var HTTPClientInfo factory.HTTPClientInfo

// loadCmd represents the load command
var registCmd = &cobra.Command{
	Use:   "post",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args:         cobra.MinimumNArgs(1),
	SilenceUsage: false,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		setInfoFromEnv(&HTTPClientInfo)
		hc := factory.NewHTTPClient(&HTTPClientInfo)
		fl := factory.NewFileLoader()
		us, err := factory.NewUsecaseMawinter(hc, fl)
		if err != nil {
			return err
		}

		ctx := context.Background()
		return us.RegistFromJSON(ctx, args[0])
	},
}

func init() {
	rootCmd.AddCommand(registCmd)

	registCmd.Flags().StringVar(&HTTPClientInfo.Scheme, "scheme", "", "http or https")
	registCmd.Flags().StringVar(&HTTPClientInfo.Host, "host", "", "server host")
	registCmd.Flags().StringVar(&HTTPClientInfo.Port, "port", "", "server port")
	registCmd.Flags().StringVar(&HTTPClientInfo.BasicAuth.User, "user", "", ".BasicAuth.User")
	registCmd.Flags().StringVar(&HTTPClientInfo.BasicAuth.Pass, "pass", "", ".BasicAuth.Pass") // overwrited by ENV
}

func setInfoFromEnv(info *factory.HTTPClientInfo) {
	pass, ok := os.LookupEnv("BASIC_PASS")
	if ok {
		info.BasicAuth.Pass = pass
	}
}
