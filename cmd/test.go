package cmd

// import (
//     "fmt"
//     "github.com/spf13/cobra"
//     "octane/pkg/octane"
// )

// // testCmd represents the test command
// var testCmd = &cobra.Command{
//     Use:   "test",
//     Short: "Run the complete test suite",
//     Long:  `Run the complete test suite to evaluate the performance of the system across various components.`,
//     Run: func(cmd *cobra.Command, args []string) {
//         // Initialize the octane calculator
//         calculator := octane.NewCalculator()

//         // Run tests
//         results, err := calculator.RunAllTests()
//         if err != nil {
//             fmt.Println("Error running tests:", err)
//             return
//         }

//         // Generate report
//         report, err := calculator.GenerateReport(results)
//         if err != nil {
//             fmt.Println("Error generating report:", err)
//             return
//         }

//         // Output the report
//         fmt.Println(report)
//     },
// }

// func init() {
//     rootCmd.AddCommand(testCmd)
// }
