package main

import (
	"context"

	"azuki774/bill-manager/internal/factory"

	"github.com/spf13/cobra"
)

var BillingMonth string

// electCmd represents the load command
var electCmd = &cobra.Command{
	Use:          "elect",
	Short:        "bill-manager-db の bill_elect を元に mawinter-api にデータを送信する",
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
		return us.RegistElectBill(ctx, BillingMonth)
	},
}

func init() {
	rootCmd.AddCommand(electCmd)

	electCmd.Flags().StringVar(&BillingMonth, "billing-month", "", "biiling month (YYYYMM)")
	electCmd.Flags().StringVar(&HTTPClientInfo.Scheme, "scheme", "", "http or https")
	electCmd.Flags().StringVar(&HTTPClientInfo.Host, "host", "", "server host")
	electCmd.Flags().StringVar(&HTTPClientInfo.Port, "port", "", "server port")
	electCmd.Flags().StringVar(&HTTPClientInfo.BasicAuth.User, "user", "", ".BasicAuth.User")
	electCmd.Flags().StringVar(&HTTPClientInfo.BasicAuth.Pass, "pass", "", ".BasicAuth.Pass") // overwrited by ENV
}
