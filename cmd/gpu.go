package cmd

// import (
//     "fmt"
//     "github.com/spf13/cobra"
//     "octane/pkg/executor"
//     "octane/pkg/yaml"
// )

// // gpuCmd represents the GPU performance test command
// var gpuCmd = &cobra.Command{
//     Use:   "gpu",
//     Short: "Test GPU performance",
//     Long:  `Run a series of tests to evaluate the performance of the GPU in the system.`,
//     Run: func(cmd *cobra.Command, args []string) {
//         // Execute the GPU test script
//         result, err := executor.ExecutePythonScript("scripts/python/gpu_test.py")
//         if err != nil {
//             fmt.Println("Error executing GPU test:", err)
//             return
//         }

//         // Process and display the results
//         report, err := yaml.ParseYAML(result)
//         if err != nil {
//             fmt.Println("Error parsing GPU test results:", err)
//             return
//         }

//         fmt.Println("GPU Performance Test Results:")
//         fmt.Println(report)
//     },
// }

// func init() {
//     rootCmd.AddCommand(gpuCmd)
// }
