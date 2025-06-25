package cmd

// import (
// 	"fmt"
// 	"octane/pkg/octane"

// 	"github.com/spf13/cobra"
// )

// // ratingCmd represents the rating command
// var ratingCmd = &cobra.Command{
// 	Use:   "rating",
// 	Short: "查看辛烷值评级",
// 	Long:  `此命令用于查看系统的辛烷值评级，帮助用户了解其系统性能等级。`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		// 获取当前系统的辛烷值评级
// 		rating := octane.GetCurrentRating()
// 		fmt.Printf("当前系统辛烷值评级: %.1f RON\n", rating.RON)
// 		fmt.Printf("评级等级: %s\n", rating.Grade)
// 		fmt.Printf("描述: %s\n", rating.Description)
// 	},
// }

// func init() {
// 	rootCmd.AddCommand(ratingCmd)
// }
