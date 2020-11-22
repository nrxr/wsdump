package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"github.com/nrxr/wsdump/ws"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "wsdump",
	Short: "simple websocket client",
	Long: `lightweight websocket client that connects to any server and sends
the input given by the user on the stdin interface`,
	Run: rootFunc,
}

func rootFunc(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("pass a websocket URL as argument")
		return
	}

	c, _ := ws.New(args[0])
	c.Run()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.wsdump.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".wsdump" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".wsdump")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
