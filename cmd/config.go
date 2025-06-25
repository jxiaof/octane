package cmd

// import (
//     "fmt"
//     "github.com/spf13/viper"
//     "github.com/spf13/cobra"
// )

// var configCmd = &cobra.Command{
//     Use:   "config",
//     Short: "Manage configuration settings",
//     Long:  `This command allows you to view and modify the configuration settings for the Octane performance testing tool.`,
//     Run: func(cmd *cobra.Command, args []string) {
//         fmt.Println("Current configuration settings:")
//         displayConfig()
//     },
// }

// func init() {
//     rootCmd.AddCommand(configCmd)
// }

// func displayConfig() {
//     fmt.Printf("Log Level: %s\n", viper.GetString("global.log_level"))
//     fmt.Printf("Output Format: %s\n", viper.GetString("global.output_format"))
//     fmt.Printf("Progress Bar: %t\n", viper.GetBool("global.progress_bar"))
//     fmt.Printf("Temporary Directory: %s\n", viper.GetString("global.temp_dir"))
//     fmt.Printf("Theme: %s\n", viper.GetString("global.theme"))
//     fmt.Printf("Upload Enabled: %t\n", viper.GetBool("upload.enabled"))
//     fmt.Printf("Upload Server URL: %s\n", viper.GetString("upload.server_url"))
//     fmt.Printf("API Key: %s\n", viper.GetString("upload.api_key"))
//     fmt.Printf("Anonymous Upload: %t\n", viper.GetBool("upload.anonymous"))
//     fmt.Printf("Auto Upload: %t\n", viper.GetBool("upload.auto_upload"))
//     fmt.Printf("Test Boost Mode: %t\n", viper.GetBool("tests.boost_mode"))
//     fmt.Printf("Fuel Analysis: %t\n", viper.GetBool("tests.fuel_analysis"))
//     fmt.Printf("Temperature Monitoring: %t\n", viper.GetBool("tests.temperature_monitoring"))
//     fmt.Printf("Power Monitoring: %t\n", viper.GetBool("tests.power_monitoring"))
// }
