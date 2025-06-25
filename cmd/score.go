package cmd

// import (
//     "fmt"
//     "github.com/spf13/cobra"
//     "octane/pkg/octane"
// )

// // scoreCmd represents the score command
// var scoreCmd = &cobra.Command{
//     Use:   "score",
//     Short: "查看辛烷值评级",
//     Long:  `此命令用于查看系统的辛烷值评级，帮助用户了解系统性能等级。`,
//     Run: func(cmd *cobra.Command, args []string) {
//         // 获取辛烷值评级
//         rating := octane.GetOctaneRating()
//         fmt.Printf("🏁 OCTANE PERFORMANCE RATING 🏁\n")
//         fmt.Printf("Overall System Octane: %.1f RON\n", rating.RON)
//         fmt.Printf("Performance Grade: %s\n", rating.Grade)
//         fmt.Printf("Description: %s\n", rating.Description)
//     },
// }

// func init() {
//     rootCmd.AddCommand(scoreCmd)
// }
