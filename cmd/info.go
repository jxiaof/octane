package cmd

// import (
// 	"fmt"
// 	"octane/pkg/executor"

// 	"github.com/spf13/cobra"
// )

// // infoCmd represents the info command
// var infoCmd = &cobra.Command{
// 	Use:   "info",
// 	Short: "View system information",
// 	Long:  `Displays detailed information about the system's hardware and software configuration.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		systemInfo, err := executor.ExecutePythonScript("scripts/python/system_info.py")
// 		if err != nil {
// 			fmt.Println("Error retrieving system information:", err)
// 			return
// 		}
// 		fmt.Println(systemInfo)
// 	},
// }

// func init() {
// 	rootCmd.AddCommand(infoCmd)
// }
