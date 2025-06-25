package cmd

// import (
// 	"fmt"
// 	"octane/pkg/octane"
// 	"octane/pkg/utils"

// 	"github.com/spf13/cobra"
// )

// // storageCmd represents the storage command
// var storageCmd = &cobra.Command{
// 	Use:   "storage",
// 	Short: "Run storage performance tests",
// 	Long:  `This command runs various storage performance tests to evaluate the performance of the storage subsystem.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		// Execute storage tests
// 		results, err := octane.RunStorageTests()
// 		if err != nil {
// 			fmt.Println("Error running storage tests:", err)
// 			return
// 		}

// 		// Display results
// 		utils.DisplayResults(results)
// 	},
// }

// func init() {
// 	rootCmd.AddCommand(storageCmd)
// }
