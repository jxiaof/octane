package octane

import (
	"math"
	"octane/pkg/types"
)

// OctaneCalculator is responsible for calculating the octane rating.
type OctaneCalculator struct {
	BaselineDB map[string]BaselineData // Baseline database
}

// NewOctaneCalculator creates a new instance of OctaneCalculator
func NewOctaneCalculator() *OctaneCalculator {
	return &OctaneCalculator{
		BaselineDB: BaselineDB,
	}
}

// CalculateOctane calculates the octane rating based on test results.
func (oc *OctaneCalculator) CalculateOctane(results *types.TestResults) *types.OctaneRating {
	cpuOctane := oc.calculateCPUOctane(results.CPU)
	memoryOctane := oc.calculateMemoryOctane(results.Memory)
	storageOctane := oc.calculateStorageOctane(results.Storage)
	gpuOctane := oc.calculateGPUOctane(results.GPU)
	networkOctane := oc.calculateNetworkOctane(results.Network)

	// 权重计算 - 根据 Octane 品牌理念调整权重
	overall := (cpuOctane*0.20 + memoryOctane*0.15 +
		storageOctane*0.15 + gpuOctane*0.25 +
		networkOctane*0.15)

	return &types.OctaneRating{
		RON:         overall,
		Grade:       oc.getGradeFromRON(overall),
		Description: oc.getDescriptionFromRON(overall),
		Color:       oc.getColorFromRON(overall),
	}
}

// CalculateComponentOctanes calculates octane ratings for individual components
func (oc *OctaneCalculator) CalculateComponentOctanes(results *types.TestResults) map[string]types.OctaneRating {
	return map[string]types.OctaneRating{
		"cpu": {
			RON:         oc.calculateCPUOctane(results.CPU),
			Grade:       oc.getGradeFromRON(oc.calculateCPUOctane(results.CPU)),
			Description: oc.getDescriptionFromRON(oc.calculateCPUOctane(results.CPU)),
			Color:       oc.getColorFromRON(oc.calculateCPUOctane(results.CPU)),
		},
		"memory": {
			RON:         oc.calculateMemoryOctane(results.Memory),
			Grade:       oc.getGradeFromRON(oc.calculateMemoryOctane(results.Memory)),
			Description: oc.getDescriptionFromRON(oc.calculateMemoryOctane(results.Memory)),
			Color:       oc.getColorFromRON(oc.calculateMemoryOctane(results.Memory)),
		},
		"storage": {
			RON:         oc.calculateStorageOctane(results.Storage),
			Grade:       oc.getGradeFromRON(oc.calculateStorageOctane(results.Storage)),
			Description: oc.getDescriptionFromRON(oc.calculateStorageOctane(results.Storage)),
			Color:       oc.getColorFromRON(oc.calculateStorageOctane(results.Storage)),
		},
		"gpu": {
			RON:         oc.calculateGPUOctane(results.GPU),
			Grade:       oc.getGradeFromRON(oc.calculateGPUOctane(results.GPU)),
			Description: oc.getDescriptionFromRON(oc.calculateGPUOctane(results.GPU)),
			Color:       oc.getColorFromRON(oc.calculateGPUOctane(results.GPU)),
		},
		"network": {
			RON:         oc.calculateNetworkOctane(results.Network),
			Grade:       oc.getGradeFromRON(oc.calculateNetworkOctane(results.Network)),
			Description: oc.getDescriptionFromRON(oc.calculateNetworkOctane(results.Network)),
			Color:       oc.getColorFromRON(oc.calculateNetworkOctane(results.Network)),
		},
	}
}

// calculateCPUOctane calculates CPU octane rating based on performance results
func (oc *OctaneCalculator) calculateCPUOctane(results types.CPUResults) float64 {
	baseline := oc.BaselineDB["default"].CPU

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
func (oc *OctaneCalculator) calculateMemoryOctane(results types.MemoryResults) float64 {
	baseline := oc.BaselineDB["default"].Memory

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
func (oc *OctaneCalculator) calculateStorageOctane(results types.StorageResults) float64 {
	if len(results.Devices) == 0 {
		return 70.0
	}

	baseline := oc.BaselineDB["default"].Storage
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
func (oc *OctaneCalculator) calculateGPUOctane(results types.GPUResults) float64 {
	baseline := oc.BaselineDB["default"].GPU

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
func (oc *OctaneCalculator) calculateNetworkOctane(results types.NetworkResults) float64 {
	baseline := oc.BaselineDB["default"].Network

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
func (oc *OctaneCalculator) getGradeFromRON(ron float64) string {
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
func (oc *OctaneCalculator) getDescriptionFromRON(ron float64) string {
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
func (oc *OctaneCalculator) getColorFromRON(ron float64) string {
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

// CalculateProfessionalScenarios calculates octane ratings for professional scenarios
func (oc *OctaneCalculator) CalculateProfessionalScenarios(results *types.TestResults) types.ProfessionalScenarios {
	cpuOctane := oc.calculateCPUOctane(results.CPU)
	gpuOctane := oc.calculateGPUOctane(results.GPU)
	memoryOctane := oc.calculateMemoryOctane(results.Memory)
	storageOctane := oc.calculateStorageOctane(results.Storage)

	return types.ProfessionalScenarios{
		Gaming: types.ProfessionalScore{
			Score:       gpuOctane*0.6 + cpuOctane*0.3 + memoryOctane*0.1,
			Grade:       oc.getGradeFromRON(gpuOctane*0.6 + cpuOctane*0.3 + memoryOctane*0.1),
			Description: "Gaming performance rating based on GPU and CPU capabilities",
		},
		AIMachineLearning: types.ProfessionalScore{
			Score:       gpuOctane*0.7 + cpuOctane*0.2 + memoryOctane*0.1,
			Grade:       oc.getGradeFromRON(gpuOctane*0.7 + cpuOctane*0.2 + memoryOctane*0.1),
			Description: "AI/ML performance rating based on compute capabilities",
		},
		ServerWorkload: types.ProfessionalScore{
			Score:       cpuOctane*0.4 + memoryOctane*0.3 + storageOctane*0.3,
			Grade:       oc.getGradeFromRON(cpuOctane*0.4 + memoryOctane*0.3 + storageOctane*0.3),
			Description: "Server workload performance rating",
		},
		Workstation: types.ProfessionalScore{
			Score:       cpuOctane*0.3 + gpuOctane*0.3 + memoryOctane*0.2 + storageOctane*0.2,
			Grade:       oc.getGradeFromRON(cpuOctane*0.3 + gpuOctane*0.3 + memoryOctane*0.2 + storageOctane*0.2),
			Description: "Professional workstation performance rating",
		},
	}
}

// GetOctaneScale returns the octane scale information
func GetOctaneScale() map[string]OctaneGrade {
	scale := make(map[string]OctaneGrade)
	for _, grade := range OctaneGrades {
		scale[grade.Grade] = grade
	}
	return scale
}
