package main

import (
	"azuki774/bill-manager/internal/factory"
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

type importOption struct {
	DBInfo struct {
		Host string
		Port string
		User string
		Pass string
		Name string
	}
	Date string // YYYYMMDD
}

var importOpt importOption

// importCmd represents the load command
var importCmd = &cobra.Command{
	Use:          "import",
	Short:        "水道の請求CSVをDBに登録する",
	SilenceUsage: false,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		ctx := context.Background()
		lg, err := factory.NewLogger()
		if err != nil {
			fmt.Println(err)
			return err
		}
		d, err := factory.NewDBRepository(importOpt.DBInfo.Host, importOpt.DBInfo.Port, importOpt.DBInfo.User, importOpt.DBInfo.Pass, importOpt.DBInfo.Name)
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer d.CloseDB()

		fl := factory.NewFileLoader()
		ap := factory.NewUsecaseWater(lg, d, fl, importOpt.Date)
		return ap.Import(ctx)
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
	importCmd.Flags().StringVar(&importOpt.Date, "date", "", "YYYYMMDD")
	// ENV: SRC_HOST
	// ENV: SRC_REMOTE_DIR
	importCmd.Flags().StringVar(&importOpt.DBInfo.Host, "db-host", "bill-manager-db", "DB Host")
	importCmd.Flags().StringVar(&importOpt.DBInfo.Port, "db-port", "3306", "DB Port")
	importCmd.Flags().StringVar(&importOpt.DBInfo.Name, "db-name", "billmanager", "DB Name")
	importCmd.Flags().StringVar(&importOpt.DBInfo.User, "db-user", "root", "DB User")
	importCmd.Flags().StringVar(&importOpt.DBInfo.Pass, "db-pass", "password", "DB Pass")
}
