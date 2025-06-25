package cmd

import (
	"fmt"
	"octane/pkg/executor"
	"octane/pkg/types"

	"github.com/spf13/cobra"
)

// cpuCmd represents the CPU performance test command
var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "Test CPU performance",
	Long:  `Run various CPU performance tests to evaluate the processing power of the system.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Parse command line arguments
		threads, _ := cmd.Flags().GetInt("threads")
		duration, _ := cmd.Flags().GetString("duration")
		testType, _ := cmd.Flags().GetString("test")

		// Execute the CPU test
		results, err := executor.ExecuteCPUTest(threads, duration, testType)
		if err != nil {
			fmt.Printf("Error executing CPU test: %v\n", err)
			return
		}

		// Display results
		displayResults(results)
	},
}

// cpuInfoCmd represents the CPU info command
var cpuInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show CPU information",
	Long:  `Display detailed information about the current CPU platform.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get CPU info
		cpuInfo, err := executor.GetCPUInfo()
		if err != nil {
			fmt.Printf("Error getting CPU info: %v\n", err)
			return
		}

		// Display CPU info
		displayCPUInfo(cpuInfo)
	},
}

func init() {
	// Add flags for the CPU command
	cpuCmd.Flags().IntP("threads", "t", 0, "Number of threads to use (default is auto)")
	cpuCmd.Flags().StringP("duration", "d", "60s", "Duration of the test (e.g., 60s, 2m)")
	cpuCmd.Flags().StringP("test", "T", "all", "Type of test to run (all|compute|crypto|compress)")

	// Add the cpu command to the root command
	rootCmd.AddCommand(cpuCmd)

	// Add the info subcommand to cpu command
	cpuCmd.AddCommand(cpuInfoCmd)
}

// displayCPUInfo formats and prints CPU information
func displayCPUInfo(info *types.CPUInfo) {
	fmt.Println("💻 CPU Platform Information:")
	fmt.Printf("Model Name: %s\n", info.ModelName)
	fmt.Printf("Brand: %s\n", info.Brand)
	fmt.Printf("Architecture: %s\n", info.Architecture)
	fmt.Printf("Physical Cores: %d\n", info.PhysicalCores)
	fmt.Printf("Logical Cores: %d\n", info.LogicalCores)
	fmt.Printf("Base Frequency: %.2f GHz\n", info.BaseFrequency)
	fmt.Printf("Max Frequency: %.2f GHz\n", info.MaxFrequency)

	if len(info.CacheL1Data) > 0 {
		fmt.Println("\n🗄️  Cache Information:")
		fmt.Printf("L1 Data Cache: %s\n", info.CacheL1Data)
		fmt.Printf("L1 Instruction Cache: %s\n", info.CacheL1Instruction)
		fmt.Printf("L2 Cache: %s\n", info.CacheL2)
		fmt.Printf("L3 Cache: %s\n", info.CacheL3)
	}

	if len(info.Features) > 0 {
		fmt.Println("\n⚡ CPU Features:")
		for i, feature := range info.Features {
			if i > 0 && i%8 == 0 {
				fmt.Println()
			}
			fmt.Printf("%-12s", feature)
		}
		fmt.Println()
	}

	if info.TDP > 0 {
		fmt.Printf("\n🔥 Thermal Design Power: %d W\n", info.TDP)
	}
}

// displayResults formats and prints the results of the CPU test
func displayResults(results *types.CPUResults) {
	fmt.Println("🔥 CPU Performance Test Results:")
	fmt.Printf("Test Suite: %s\n", results.TestSuite)
	fmt.Printf("Duration: %s\n", results.Duration)
	fmt.Printf("Temperature: %.1f°C (Idle) / %.1f°C (Load) / %.1f°C (Max)\n",
		results.Temperature.Idle, results.Temperature.Load, results.Temperature.Max)
	fmt.Printf("Average Frequency: %.1f MHz (%.1f%% stable)\n",
		results.Frequencies.AverageAllCores, results.Frequencies.Stability)

	fmt.Println("\n📊 Single-Core Performance:")
	fmt.Printf("  Integer Performance: %d %s (Percentile: %d%%)\n",
		results.Tests.SingleCore.IntegerPerformance.Score,
		results.Tests.SingleCore.IntegerPerformance.Unit,
		results.Tests.SingleCore.IntegerPerformance.Percentile)

	fmt.Println("\n🚀 Multi-Core Performance:")
	fmt.Printf("  Integer Performance: %d %s (Percentile: %d%%)\n",
		results.Tests.MultiCore.IntegerPerformance.Score,
		results.Tests.MultiCore.IntegerPerformance.Unit,
		results.Tests.MultiCore.IntegerPerformance.Percentile)

	if results.Tests.SingleCore.Cryptography.AES256 > 0 {
		fmt.Println("\n🔐 Cryptography Performance:")
		fmt.Printf("  AES-256: %.2f GB/s\n", results.Tests.SingleCore.Cryptography.AES256)
		fmt.Printf("  SHA-256: %.2f GB/s\n", results.Tests.SingleCore.Cryptography.SHA256)
		fmt.Printf("  RSA-2048: %d ops/sec\n", results.Tests.SingleCore.Cryptography.RSA2048)
	}

	if results.Tests.MultiCore.Compression.Gzip > 0 {
		fmt.Println("\n📦 Compression Performance:")
		fmt.Printf("  Gzip: %d MB/s\n", results.Tests.MultiCore.Compression.Gzip)
		fmt.Printf("  LZ4: %d MB/s\n", results.Tests.MultiCore.Compression.LZ4)
		fmt.Printf("  Zstd: %d MB/s\n", results.Tests.MultiCore.Compression.Zstd)
	}
}
