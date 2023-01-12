package main

import (
	"azuki774/bill-manager/internal/factory"
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

type startOption struct {
	Port   string
	DBInfo struct {
		Host string
		Port string
		User string
		Pass string
		Name string
	}
}

var startOpt startOption

// startCmd represents the load command
var startCmd = &cobra.Command{
	Use:          "start",
	Short:        "bill-manager API",
	SilenceUsage: false,
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
		defer d.CloseDB()
		ap := factory.NewAPIService(lg, d)
		s := factory.NewAPIServer(lg, startOpt.Port, ap)
		return s.Start(ctx)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVar(&startOpt.Port, "port", "80", "listen port")
	startCmd.Flags().StringVar(&startOpt.DBInfo.Host, "db-host", "bill-manager-db", "DB Host")
	startCmd.Flags().StringVar(&startOpt.DBInfo.Port, "db-port", "3306", "DB Port")
	startCmd.Flags().StringVar(&startOpt.DBInfo.Name, "db-name", "billmanager", "DB Name")
	startCmd.Flags().StringVar(&startOpt.DBInfo.User, "db-user", "root", "DB User")
	startCmd.Flags().StringVar(&startOpt.DBInfo.Pass, "db-pass", "password", "DB Pass")
}
