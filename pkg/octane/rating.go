package octane

import (
	"math"
	"octane/pkg/types"
)

// OctaneRating represents the octane rating of a system.
type OctaneRating struct {
	RON         float64 `yaml:"ron"`         // Research Octane Number
	Grade       string  `yaml:"grade"`       // Grade name
	Description string  `yaml:"description"` // Description
	Color       string  `yaml:"color"`       // Display color
}

// CalculateOctane calculates the octane rating based on test results.
func CalculateOctane(results *types.TestResults) *OctaneRating {
	cpuOctane := calculateCPUOctane(results.CPU)
	memoryOctane := calculateMemoryOctane(results.Memory)
	storageOctane := calculateStorageOctane(results.Storage)
	gpuOctane := calculateGPUOctane(results.GPU)
	networkOctane := calculateNetworkOctane(results.Network)

	overall := (cpuOctane*0.20 + memoryOctane*0.15 +
		storageOctane*0.15 + gpuOctane*0.25 +
		networkOctane*0.15)

	return &OctaneRating{
		RON:         overall,
		Grade:       getGradeFromRON(overall),
		Description: getDescriptionFromRON(overall),
		Color:       getColorFromRON(overall),
	}
}

// calculateCPUOctane calculates CPU octane rating based on performance results
func calculateCPUOctane(results types.CPUResults) float64 {
	baseline := GetBaseline("default").CPU

	// 综合考虑单核和多核性能
	singleCoreScore := float64(results.Tests.SingleCore.IntegerPerformance.Score)
	multiCoreScore := float64(results.Tests.MultiCore.IntegerPerformance.Score)

	// 使用对数函数进行评分，确保高端性能的区分度
	singleCoreOctane := 70 + 30*math.Log10(singleCoreScore/baseline)
	multiCoreOctane := 70 + 30*math.Log10(multiCoreScore/(baseline*8)) // 假设基准为8核

	// 综合评分 (单核40%, 多核60%)
	overall := singleCoreOctane*0.4 + multiCoreOctane*0.6

	// 温度惩罚机制
	if results.Temperature.Max > 85 {
		overall *= 0.95 // 高温降低5%评分
	}

	return math.Min(100, math.Max(70, overall))
}

// calculateMemoryOctane calculates memory octane rating
func calculateMemoryOctane(results types.MemoryResults) float64 {
	baseline := GetBaseline("default").Memory

	// 综合带宽和延迟评分
	avgBandwidth := (results.Bandwidth.SequentialRead + results.Bandwidth.SequentialWrite +
		results.Bandwidth.Copy) / 3

	bandwidthOctane := 70 + 30*math.Log10(avgBandwidth/baseline)

	// 延迟评分 (延迟越低越好)
	latencyScore := 100 - results.Latency.MainMemory // 简化计算
	latencyOctane := 70 + 30*(latencyScore/100)

	// 稳定性加分
	stabilityBonus := 0.0
	if results.Stability.ErrorsDetected == 0 {
		stabilityBonus = 5.0
	}

	overall := (bandwidthOctane*0.7 + latencyOctane*0.3) + stabilityBonus

	return math.Min(100, math.Max(70, overall))
}

// calculateStorageOctane calculates storage octane rating
func calculateStorageOctane(results types.StorageResults) float64 {
	if len(results.Devices) == 0 {
		return 70.0
	}

	baseline := GetBaseline("default").Storage
	totalOctane := 0.0

	for _, device := range results.Devices {
		// 顺序读写性能
		seqReadOctane := 70 + 30*math.Log10(device.Tests.Sequential.Read1MB/baseline)
		seqWriteOctane := 70 + 30*math.Log10(device.Tests.Sequential.Write1MB/baseline)

		// 随机性能 (IOPS)
		randomReadOctane := 70 + 30*math.Log10(device.Tests.Random.Read4KIops/(baseline*1000))
		randomWriteOctane := 70 + 30*math.Log10(device.Tests.Random.Write4KIops/(baseline*1000))

		// 延迟评分
		latencyScore := 100 - (device.Tests.Latency.ReadAvg + device.Tests.Latency.WriteAvg)
		latencyOctane := 70 + 30*(latencyScore/100)

		deviceOctane := (seqReadOctane*0.2 + seqWriteOctane*0.2 +
			randomReadOctane*0.2 + randomWriteOctane*0.2 +
			latencyOctane*0.2)

		totalOctane += deviceOctane
	}

	overall := totalOctane / float64(len(results.Devices))
	return math.Min(100, math.Max(70, overall))
}

// calculateGPUOctane calculates GPU octane rating
func calculateGPUOctane(results types.GPUResults) float64 {
	baseline := GetBaseline("default").GPU

	// 图形性能评分
	graphicsScore := results.Tests.Graphics.Score
	graphicsOctane := 70 + 30*math.Log10(graphicsScore/baseline)

	// 计算性能评分
	computeOctane := 70 + 30*math.Log10(results.Tests.Compute.SinglePrecision/(baseline*10))

	// 机器学习性能 (如果有的话)
	mlOctane := graphicsOctane // 默认使用图形性能作为备选

	// 温度和功耗惩罚
	tempPenalty := 1.0
	if results.Temperature.Max > 80 {
		tempPenalty = 0.95
	}

	powerPenalty := 1.0
	if results.PowerConsumption.Peak > 400 { // 400W以上功耗惩罚
		powerPenalty = 0.98
	}

	overall := (graphicsOctane*0.4 + computeOctane*0.4 + mlOctane*0.2) * tempPenalty * powerPenalty

	return math.Min(100, math.Max(70, overall))
}

// calculateNetworkOctane calculates network octane rating
func calculateNetworkOctane(results types.NetworkResults) float64 {
	baseline := GetBaseline("default").Network

	// 国内带宽评分
	domesticAvg := 0.0
	domesticCount := 0
	for _, result := range results.Bandwidth.Domestic {
		domesticAvg += result.Download
		domesticCount++
	}
	if domesticCount > 0 {
		domesticAvg /= float64(domesticCount)
	}

	bandwidthOctane := 70 + 30*math.Log10(domesticAvg/baseline)

	// 延迟评分
	latencyScore := 100.0
	for _, result := range results.Bandwidth.Domestic {
		if result.Latency > 50 { // 50ms以上延迟惩罚
			latencyScore -= 10
		}
	}
	latencyOctane := 70 + 30*(latencyScore/100)

	// 连通性评分
	connectivityScore := 0
	totalServices := len(results.Connectivity.ServiceAccessibility)
	for _, accessible := range results.Connectivity.ServiceAccessibility {
		if accessible {
			connectivityScore++
		}
	}

	connectivityOctane := 70 + 30*float64(connectivityScore)/float64(totalServices)

	overall := bandwidthOctane*0.5 + latencyOctane*0.3 + connectivityOctane*0.2

	return math.Min(100, math.Max(70, overall))
}

// getGradeFromRON maps the RON to a grade.
func getGradeFromRON(ron float64) string {
	switch {
	case ron >= 95:
		return "racing_fuel"
	case ron >= 90:
		return "premium_plus"
	case ron >= 85:
		return "premium"
	case ron >= 80:
		return "regular_plus"
	default:
		return "regular"
	}
}

// getDescriptionFromRON provides a description based on RON.
func getDescriptionFromRON(ron float64) string {
	switch {
	case ron >= 95:
		return "Ultimate performance for extreme workloads"
	case ron >= 90:
		return "High performance for demanding applications"
	case ron >= 85:
		return "Good performance for most applications"
	case ron >= 80:
		return "Standard performance for regular use"
	default:
		return "Basic performance for light workloads"
	}
}

// getColorFromRON returns a color representation based on RON.
func getColorFromRON(ron float64) string {
	switch {
	case ron >= 95:
		return "🔥 RED"
	case ron >= 90:
		return "🟠 ORANGE"
	case ron >= 85:
		return "🟡 YELLOW"
	case ron >= 80:
		return "🟢 GREEN"
	default:
		return "🔵 BLUE"
	}
}
