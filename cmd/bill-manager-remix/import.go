package main

import (
	"azuki774/bill-manager/internal/factory"
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

type startOption struct {
	DBInfo struct {
		Host string
		Port string
		User string
		Pass string
		Name string
	}
	Date string // YYYYMMDD
}

var startOpt startOption

// loadCmd represents the load command
var importCmd = &cobra.Command{
	Use:          "import",
	Short:        "リミックスでんきの請求CSVをDBに登録する",
	SilenceUsage: false,
	Args:         cobra.MinimumNArgs(1), // consume or bill
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		ctx := context.Background()
		lg, err := factory.NewLogger()
		if err != nil {
			fmt.Println(err)
			return err
		}
		d, err := factory.NewDBRepository(startOpt.DBInfo.Host, startOpt.DBInfo.Port, startOpt.DBInfo.User, startOpt.DBInfo.Pass, startOpt.DBInfo.Name)
		if err != nil {
			fmt.Println(err)
			return err
		}

		fl := factory.NewFileLoader()
		ip := factory.NewUsecaseRemix(lg, d, fl)
		return ip.Start(ctx, args[0])
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	importCmd.Flags().StringVar(&startOpt.Date, "date", "", "YYYYMMDD")
	importCmd.Flags().StringVar(&startOpt.DBInfo.Host, "db-host", "bill-manager-db", "DB Host")
	importCmd.Flags().StringVar(&startOpt.DBInfo.Port, "db-port", "3306", "DB Port")
	importCmd.Flags().StringVar(&startOpt.DBInfo.Name, "db-name", "billmanager", "DB Name")
	importCmd.Flags().StringVar(&startOpt.DBInfo.User, "db-user", "root", "DB User")
	importCmd.Flags().StringVar(&startOpt.DBInfo.Pass, "db-pass", "password", "DB Pass")
}
