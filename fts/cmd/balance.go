package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	client "github.com/textileio/powergate/api/client"
)

func init() {
	rootCmd.AddCommand(balanceCmd)
}

var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "the Filecoin wallet address and balance",
	Long:  `Print the address and balance of the Filecoin wallet`,
	PreRun: func(cmd *cobra.Command, args []string) {
		err := viper.BindPFlags(cmd.Flags())
		checkErr(err)
	},
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), cmdTimeout)
		defer cancel()
		ctx = context.WithValue(ctx, client.AuthKey, viper.GetString("token"))
		powClient, err := client.NewClient(viper.GetString("server-address"))
		checkErr(err)

		addr, err := powClient.Wallet.Addresses(ctx)
		checkErr(err)

		if len(addr.Addresses) == 0 {
			checkErr(fmt.Errorf("No address found"))
		}

		res, err := powClient.Wallet.Balance(ctx, addr.Addresses[0].Address)
		checkErr(err)

		combined := map[string]string{
			"address": addr.Addresses[0].Address,
			"balance": res.Balance,
		}

		json, err := json.MarshalIndent(combined, "", "  ")
		checkErr(err)

		fmt.Println(string(json))
	},
}
