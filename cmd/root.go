package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "octane",
	Short: "Octane Performance Analyzer",
	Long:  `A comprehensive tool to evaluate the performance of your system.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Default action when no subcommand is provided
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Define global flags
	rootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is $HOME/.octane.yaml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "enable verbose output")

	// Bind flags to viper
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

func initConfig() {
	if cfgFile := viper.GetString("config"); cfgFile != "" {
		viper.SetConfigFile(cfgFile)

		if err := viper.ReadInConfig(); err != nil {
			log.Printf("Error reading config file %s: %v", cfgFile, err)
		}
	} else {
		// 只有在显式指定配置文件时才尝试读取，避免不必要的错误日志
		viper.AddConfigPath("$HOME")
		viper.SetConfigName(".octane")
		viper.SetConfigType("yaml")

		// 静默处理配置文件不存在的情况
		if err := viper.ReadInConfig(); err != nil {
			// 只在verbose模式下显示配置文件相关信息
			if viper.GetBool("verbose") {
				log.Printf("Config file not found (this is optional): %v", err)
			}
		}
	}
}
