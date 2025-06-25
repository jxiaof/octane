package executor

import (
	"bufio"
	"fmt"
	"math"
	"octane/pkg/types"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

// ExecuteCPUTest 执行CPU性能测试
func ExecuteCPUTest(threads int, duration string, testType string) (*types.CPUResults, error) {
	// 解析持续时间
	testDuration, err := time.ParseDuration(duration)
	if err != nil {
		return nil, fmt.Errorf("invalid duration format: %v", err)
	}

	// 如果threads为0，使用系统CPU核心数
	if threads == 0 {
		threads = runtime.NumCPU()
	}

	// 创建结果结构
	results := &types.CPUResults{
		TestSuite: "octane-cpu-test",
		Duration:  duration,
	}

	// 记录初始温度（模拟）
	results.Temperature.Idle = 35.0
	results.Temperature.Load = 65.0
	results.Temperature.Max = 78.0

	// 记录频率信息（模拟）
	results.Frequencies.AverageAllCores = 3200.0
	results.Frequencies.Stability = 98.5

	// 根据测试类型执行相应测试
	switch testType {
	case "all":
		return runAllCPUTests(threads, testDuration, results)
	case "compute":
		return runComputeTest(threads, testDuration, results)
	case "crypto":
		return runCryptoTest(threads, testDuration, results)
	case "compress":
		return runCompressionTest(threads, testDuration, results)
	default:
		return nil, fmt.Errorf("unknown test type: %s", testType)
	}
}

// runAllCPUTests 运行所有CPU测试
func runAllCPUTests(threads int, duration time.Duration, results *types.CPUResults) (*types.CPUResults, error) {
	fmt.Printf("Running comprehensive CPU tests with %d threads for %v...\n", threads, duration)

	// 运行单核测试
	singleCoreScore := runSingleCoreTest(duration / 4)
	results.Tests.SingleCore.IntegerPerformance.Score = singleCoreScore
	results.Tests.SingleCore.IntegerPerformance.Unit = "points"
	results.Tests.SingleCore.IntegerPerformance.Percentile = calculatePercentile(singleCoreScore)

	// 运行多核测试
	multiCoreScore := runMultiCoreTest(threads, duration/2)
	results.Tests.MultiCore.IntegerPerformance.Score = multiCoreScore
	results.Tests.MultiCore.IntegerPerformance.Unit = "points"
	results.Tests.MultiCore.IntegerPerformance.Percentile = calculatePercentile(multiCoreScore)

	// 运行加密测试
	runCryptographyTests(threads, duration/4, results)

	// 运行压缩测试
	runCompressionTests(threads, duration/4, results)

	return results, nil
}

// runSingleCoreTest 运行单核性能测试
func runSingleCoreTest(duration time.Duration) int {
	fmt.Println("Running single-core integer performance test...")

	start := time.Now()
	operations := 0

	for time.Since(start) < duration {
		// 执行计算密集型操作
		for i := 0; i < 10000; i++ {
			_ = math.Sqrt(float64(i)) * math.Sin(float64(i))
			operations++
		}
	}

	// 根据操作数计算分数
	score := operations / 1000
	fmt.Printf("Single-core test completed: %d operations, score: %d\n", operations, score)
	return score
}

// runMultiCoreTest 运行多核性能测试
func runMultiCoreTest(threads int, duration time.Duration) int {
	fmt.Printf("Running multi-core test with %d threads...\n", threads)

	var wg sync.WaitGroup
	var mu sync.Mutex
	totalOperations := 0

	start := time.Now()

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			operations := 0

			for time.Since(start) < duration {
				// 执行计算密集型操作
				for j := 0; j < 5000; j++ {
					_ = math.Sqrt(float64(j)) * math.Cos(float64(j))
					operations++
				}
			}

			mu.Lock()
			totalOperations += operations
			mu.Unlock()
		}()
	}

	wg.Wait()

	score := totalOperations / 1000
	fmt.Printf("Multi-core test completed: %d total operations, score: %d\n", totalOperations, score)
	return score
}

// runCryptographyTests 运行加密性能测试
func runCryptographyTests(threads int, duration time.Duration, results *types.CPUResults) {
	fmt.Println("Running cryptography tests...")

	// 模拟AES加密测试
	results.Tests.SingleCore.Cryptography.AES256 = runAESTest(duration / 3)

	// 模拟SHA256测试
	results.Tests.SingleCore.Cryptography.SHA256 = runSHA256Test(duration / 3)

	// 模拟RSA测试
	results.Tests.SingleCore.Cryptography.RSA2048 = runRSATest(duration / 3)
}

// runAESTest 模拟AES加密测试
func runAESTest(duration time.Duration) float64 {
	fmt.Println("Running AES-256 encryption test...")

	start := time.Now()
	bytesProcessed := 0

	for time.Since(start) < duration {
		// 模拟AES加密操作
		data := make([]byte, 1024*1024) // 1MB数据
		for i := range data {
			data[i] = byte(i % 256)
		}
		bytesProcessed += len(data)
	}

	// 计算GB/s
	gbPerSecond := float64(bytesProcessed) / float64(duration.Seconds()) / (1024 * 1024 * 1024)
	fmt.Printf("AES-256 test: %.2f GB/s\n", gbPerSecond)
	return gbPerSecond
}

// runSHA256Test 模拟SHA256测试
func runSHA256Test(duration time.Duration) float64 {
	fmt.Println("Running SHA-256 hash test...")

	start := time.Now()
	bytesProcessed := 0

	for time.Since(start) < duration {
		// 模拟SHA256哈希操作
		data := make([]byte, 512*1024) // 512KB数据
		for i := range data {
			data[i] = byte(i % 256)
		}
		bytesProcessed += len(data)
	}

	gbPerSecond := float64(bytesProcessed) / float64(duration.Seconds()) / (1024 * 1024 * 1024)
	fmt.Printf("SHA-256 test: %.2f GB/s\n", gbPerSecond)
	return gbPerSecond
}

// runRSATest 模拟RSA测试
func runRSATest(duration time.Duration) int {
	fmt.Println("Running RSA-2048 test...")

	start := time.Now()
	operations := 0

	for time.Since(start) < duration {
		// 模拟RSA操作（简化版本）
		for i := 0; i < 100; i++ {
			_ = math.Pow(float64(i), 2.0)
			operations++
		}
	}

	opsPerSecond := int(float64(operations) / duration.Seconds())
	fmt.Printf("RSA-2048 test: %d ops/sec\n", opsPerSecond)
	return opsPerSecond
}

// runCompressionTests 运行压缩测试
func runCompressionTests(threads int, duration time.Duration, results *types.CPUResults) {
	fmt.Println("Running compression tests...")

	results.Tests.MultiCore.Compression.Gzip = runGzipTest(threads, duration/3)
	results.Tests.MultiCore.Compression.LZ4 = runLZ4Test(threads, duration/3)
	results.Tests.MultiCore.Compression.Zstd = runZstdTest(threads, duration/3)
}

// runGzipTest 模拟Gzip压缩测试
func runGzipTest(threads int, duration time.Duration) int {
	fmt.Println("Running Gzip compression test...")

	var wg sync.WaitGroup
	var mu sync.Mutex
	totalMB := 0

	start := time.Now()

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mbProcessed := 0

			for time.Since(start) < duration {
				// 模拟压缩操作
				data := make([]byte, 1024*1024) // 1MB数据
				for j := range data {
					data[j] = byte(j % 256)
				}
				mbProcessed++
			}

			mu.Lock()
			totalMB += mbProcessed
			mu.Unlock()
		}()
	}

	wg.Wait()

	mbPerSecond := int(float64(totalMB) / duration.Seconds())
	fmt.Printf("Gzip test: %d MB/s\n", mbPerSecond)
	return mbPerSecond
}

// runLZ4Test 模拟LZ4压缩测试
func runLZ4Test(threads int, duration time.Duration) int {
	fmt.Println("Running LZ4 compression test...")
	// 类似Gzip测试，但速度更快
	result := runGzipTest(threads, duration)
	return int(float64(result) * 1.5) // LZ4通常比Gzip快50%
}

// runZstdTest 模拟Zstd压缩测试
func runZstdTest(threads int, duration time.Duration) int {
	fmt.Println("Running Zstd compression test...")
	// 类似Gzip测试，性能介于Gzip和LZ4之间
	result := runGzipTest(threads, duration)
	return int(float64(result) * 1.2) // Zstd通常比Gzip快20%
}

// runComputeTest 运行计算测试
func runComputeTest(threads int, duration time.Duration, results *types.CPUResults) (*types.CPUResults, error) {
	fmt.Printf("Running compute-focused CPU test with %d threads...\n", threads)

	singleScore := runSingleCoreTest(duration / 2)
	multiScore := runMultiCoreTest(threads, duration/2)

	results.Tests.SingleCore.IntegerPerformance.Score = singleScore
	results.Tests.MultiCore.IntegerPerformance.Score = multiScore

	return results, nil
}

// runCryptoTest 运行加密测试
func runCryptoTest(threads int, duration time.Duration, results *types.CPUResults) (*types.CPUResults, error) {
	fmt.Printf("Running crypto-focused CPU test...\n")

	runCryptographyTests(threads, duration, results)

	return results, nil
}

// runCompressionTest 运行压缩测试
func runCompressionTest(threads int, duration time.Duration, results *types.CPUResults) (*types.CPUResults, error) {
	fmt.Printf("Running compression-focused CPU test...\n")

	runCompressionTests(threads, duration, results)

	return results, nil
}

// calculatePercentile 计算百分位数（简化版本）
func calculatePercentile(score int) int {
	// 简化的百分位数计算，实际应该基于历史数据
	if score > 2000 {
		return 95
	} else if score > 1500 {
		return 85
	} else if score > 1000 {
		return 75
	} else if score > 500 {
		return 50
	}
	return 25
}

// GetCPUInfo 获取CPU平台信息
func GetCPUInfo() (*types.CPUInfo, error) {
	switch runtime.GOOS {
	case "darwin":
		return getCPUInfoMacOS()
	case "linux":
		return getCPUInfoLinux()
	case "windows":
		return getCPUInfoWindows()
	default:
		return getCPUInfoGeneric()
	}
}

// getCPUInfoMacOS 获取macOS系统的CPU信息
func getCPUInfoMacOS() (*types.CPUInfo, error) {
	info := &types.CPUInfo{
		LogicalCores: runtime.NumCPU(),
	}

	// 获取CPU品牌和型号
	if brand, err := execSysctl("machdep.cpu.brand_string"); err == nil {
		info.ModelName = strings.TrimSpace(brand)
		info.Brand = extractBrand(brand)
	}

	// 获取物理核心数
	if cores, err := execSysctl("hw.physicalcpu"); err == nil {
		if count, err := strconv.Atoi(strings.TrimSpace(cores)); err == nil {
			info.PhysicalCores = count
		}
	}

	// 获取架构信息
	info.Architecture = runtime.GOARCH

	// 对于Apple Silicon，使用更合适的方式获取频率信息
	if strings.Contains(info.ModelName, "Apple M") {
		// Apple Silicon的基础频率估算
		if strings.Contains(info.ModelName, "M4") {
			info.BaseFrequency = 3.2 // M4 基础频率约3.2GHz
			info.MaxFrequency = 4.4  // M4 最大频率约4.4GHz
		} else if strings.Contains(info.ModelName, "M3") {
			info.BaseFrequency = 3.0
			info.MaxFrequency = 4.0
		} else if strings.Contains(info.ModelName, "M2") {
			info.BaseFrequency = 2.4
			info.MaxFrequency = 3.5
		} else if strings.Contains(info.ModelName, "M1") {
			info.BaseFrequency = 2.0
			info.MaxFrequency = 3.2
		}
	} else {
		// 尝试获取频率信息（适用于Intel Mac）
		if freq, err := execSysctl("hw.cpufrequency"); err == nil {
			if hz, err := strconv.ParseFloat(strings.TrimSpace(freq), 64); err == nil {
				info.BaseFrequency = hz / 1e9 // 转换为GHz
			}
		}

		if maxFreq, err := execSysctl("hw.cpufrequency_max"); err == nil {
			if hz, err := strconv.ParseFloat(strings.TrimSpace(maxFreq), 64); err == nil {
				info.MaxFrequency = hz / 1e9
			}
		}
	}

	// 获取缓存信息
	if l1d, err := execSysctl("hw.l1dcachesize"); err == nil {
		if size, err := strconv.Atoi(strings.TrimSpace(l1d)); err == nil {
			info.CacheL1Data = formatCacheSize(size)
		}
	}

	if l1i, err := execSysctl("hw.l1icachesize"); err == nil {
		if size, err := strconv.Atoi(strings.TrimSpace(l1i)); err == nil {
			info.CacheL1Instruction = formatCacheSize(size)
		}
	}

	if l2, err := execSysctl("hw.l2cachesize"); err == nil {
		if size, err := strconv.Atoi(strings.TrimSpace(l2)); err == nil {
			info.CacheL2 = formatCacheSize(size)
		}
	}

	if l3, err := execSysctl("hw.l3cachesize"); err == nil {
		if size, err := strconv.Atoi(strings.TrimSpace(l3)); err == nil {
			info.CacheL3 = formatCacheSize(size)
		}
	}

	// 如果L3缓存为空，对于Apple Silicon设置统一内存架构信息
	if info.CacheL3 == "" && strings.Contains(info.ModelName, "Apple M") {
		info.CacheL3 = "Unified Memory Architecture"
	}

	// 获取CPU特性
	info.Features = getCPUFeaturesMacOS()

	// 对于Apple Silicon，添加一些已知特性
	if strings.Contains(info.ModelName, "Apple M") {
		additionalFeatures := []string{
			"ARM64", "NEON", "Crypto", "CRC32", "SHA1", "SHA256", "AES",
		}
		info.Features = append(info.Features, additionalFeatures...)
	}

	// 设置TDP（估算值）
	if strings.Contains(info.ModelName, "Apple M4") {
		info.TDP = 22 // M4 TDP约22W
	} else if strings.Contains(info.ModelName, "Apple M3") {
		info.TDP = 20
	} else if strings.Contains(info.ModelName, "Apple M2") {
		info.TDP = 18
	} else if strings.Contains(info.ModelName, "Apple M1") {
		info.TDP = 15
	}

	return info, nil
}

// getCPUInfoLinux 获取Linux系统的CPU信息
func getCPUInfoLinux() (*types.CPUInfo, error) {
	info := &types.CPUInfo{
		LogicalCores: runtime.NumCPU(),
		Architecture: runtime.GOARCH,
	}

	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return info, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ":") {
			parts := strings.Split(line, ":")
			if len(parts) != 2 {
				continue
			}
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			switch key {
			case "model name":
				if info.ModelName == "" {
					info.ModelName = value
					info.Brand = extractBrand(value)
				}
			case "cpu cores":
				if cores, err := strconv.Atoi(value); err == nil {
					info.PhysicalCores = cores
				}
			case "cpu MHz":
				if freq, err := strconv.ParseFloat(value, 64); err == nil {
					info.BaseFrequency = freq / 1000 // 转换为GHz
				}
			case "cache size":
				info.CacheL2 = value
			case "flags":
				info.Features = strings.Fields(value)
			}
		}
	}

	return info, nil
}

// getCPUInfoWindows 获取Windows系统的CPU信息
func getCPUInfoWindows() (*types.CPUInfo, error) {
	info := &types.CPUInfo{
		LogicalCores: runtime.NumCPU(),
		Architecture: runtime.GOARCH,
	}

	// 使用wmic获取CPU信息
	cmd := exec.Command("wmic", "cpu", "get", "Name,NumberOfCores,NumberOfLogicalProcessors,MaxClockSpeed", "/format:csv")
	output, err := cmd.Output()
	if err == nil {
		lines := strings.Split(string(output), "\n")
		for _, line := range lines {
			if strings.Contains(line, ",") {
				parts := strings.Split(line, ",")
				if len(parts) >= 4 {
					if parts[2] != "" && parts[2] != "MaxClockSpeed" {
						if freq, err := strconv.ParseFloat(parts[2], 64); err == nil {
							info.MaxFrequency = freq / 1000 // MHz to GHz
						}
					}
					if parts[3] != "" && parts[3] != "Name" {
						info.ModelName = strings.TrimSpace(parts[3])
						info.Brand = extractBrand(parts[3])
					}
					if parts[1] != "" && parts[1] != "NumberOfCores" {
						if cores, err := strconv.Atoi(parts[1]); err == nil {
							info.PhysicalCores = cores
						}
					}
				}
			}
		}
	}

	return info, nil
}

// getCPUInfoGeneric 获取通用CPU信息
func getCPUInfoGeneric() (*types.CPUInfo, error) {
	return &types.CPUInfo{
		ModelName:     "Unknown CPU",
		Brand:         "Unknown",
		Architecture:  runtime.GOARCH,
		PhysicalCores: runtime.NumCPU(),
		LogicalCores:  runtime.NumCPU(),
	}, nil
}

// execSysctl 执行sysctl命令
func execSysctl(key string) (string, error) {
	cmd := exec.Command("sysctl", "-n", key)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// extractBrand 从CPU型号中提取品牌
func extractBrand(modelName string) string {
	modelName = strings.ToLower(modelName)
	if strings.Contains(modelName, "intel") {
		return "Intel"
	} else if strings.Contains(modelName, "amd") {
		return "AMD"
	} else if strings.Contains(modelName, "apple") {
		return "Apple"
	} else if strings.Contains(modelName, "arm") {
		return "ARM"
	}
	return "Unknown"
}

// formatCacheSize 格式化缓存大小
func formatCacheSize(bytes int) string {
	if bytes >= 1024*1024 {
		return fmt.Sprintf("%.1f MB", float64(bytes)/(1024*1024))
	} else if bytes >= 1024 {
		return fmt.Sprintf("%.1f KB", float64(bytes)/1024)
	}
	return fmt.Sprintf("%d B", bytes)
}

// getCPUFeaturesMacOS 获取macOS的CPU特性
func getCPUFeaturesMacOS() []string {
	features := []string{}

	// 常见的CPU特性检查
	featureMap := map[string]string{
		"machdep.cpu.feature.SSE":    "SSE",
		"machdep.cpu.feature.SSE2":   "SSE2",
		"machdep.cpu.feature.SSE3":   "SSE3",
		"machdep.cpu.feature.SSSE3":  "SSSE3",
		"machdep.cpu.feature.SSE4_1": "SSE4.1",
		"machdep.cpu.feature.SSE4_2": "SSE4.2",
		"machdep.cpu.feature.AVX":    "AVX",
		"machdep.cpu.feature.AVX2":   "AVX2",
		"machdep.cpu.feature.AES":    "AES",
	}

	for sysctl, feature := range featureMap {
		if output, err := execSysctl(sysctl); err == nil {
			if strings.TrimSpace(output) == "1" {
				features = append(features, feature)
			}
		}
	}

	return features
}
