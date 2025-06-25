package cmd

// import (
// 	"fmt"
// 	"octane/pkg/executor"
// 	"octane/pkg/types"

// 	"github.com/spf13/cobra"
// )

// // memoryCmd represents the memory command
// var memoryCmd = &cobra.Command{
// 	Use:   "memory",
// 	Short: "Test memory performance",
// 	Long:  `Run various tests to evaluate the memory performance of the system.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		// Define memory test parameters
// 		size, _ := cmd.Flags().GetString("size")
// 		testType, _ := cmd.Flags().GetString("test")

// 		// Execute memory test
// 		results, err := executor.ExecuteMemoryTest(size, testType)
// 		if err != nil {
// 			fmt.Printf("Error executing memory test: %v\n", err)
// 			return
// 		}

// 		// Display results
// 		displayMemoryResults(results)
// 	},
// }

// func init() {
// 	// Add flags for memory command
// 	memoryCmd.Flags().StringP("size", "s", "1GB", "Size of memory to test (e.g., 1GB, 2GB)")
// 	memoryCmd.Flags().StringP("test", "t", "all", "Type of memory test to run (e.g., all, bandwidth, latency, stability)")

// 	// Add memory command to root command
// 	rootCmd.AddCommand(memoryCmd)
// }

// // displayMemoryResults formats and displays the memory test results
// func displayMemoryResults(results *types.MemoryResults) {
// 	fmt.Println("Memory Test Results:")
// 	fmt.Printf("Total Size: %s\n", results.TotalSize)
// 	fmt.Printf("Bandwidth: %s\n", results.Bandwidth)
// 	fmt.Printf("Latency: %s\n", results.Latency)
// 	fmt.Printf("Stability: %s\n", results.Stability)
// }
