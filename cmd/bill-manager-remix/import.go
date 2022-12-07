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
	Use:   "import",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

		fl := factory.NewFileLoader()
		ip := factory.NewUsecaseRemix(lg, d, fl)
		return ip.Start(ctx)
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
