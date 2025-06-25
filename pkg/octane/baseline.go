package octane

// BaselineData 定义基准数据结构
type BaselineData struct {
	CPU     float64 `yaml:"cpu"`     // CPU基准值
	Memory  float64 `yaml:"memory"`  // 内存基准值
	Storage float64 `yaml:"storage"` // 存储基准值
	GPU     float64 `yaml:"gpu"`     // GPU基准值
	Network float64 `yaml:"network"` // 网络基准值
}

// BaselineDB 定义基准数据库
var BaselineDB = map[string]BaselineData{
	"default": {
		CPU:     1000.0,  // 基准CPU性能分数
		Memory:  25000.0, // 基准内存带宽 MB/s
		Storage: 500.0,   // 基准存储速度 MB/s
		GPU:     10000.0, // 基准GPU性能分数
		Network: 100.0,   // 基准网络速度 Mbps
	},
	"entry_level": {
		CPU:     500.0,
		Memory:  15000.0,
		Storage: 150.0,
		GPU:     3000.0,
		Network: 50.0,
	},
	"mid_range": {
		CPU:     1500.0,
		Memory:  35000.0,
		Storage: 1000.0,
		GPU:     15000.0,
		Network: 200.0,
	},
	"high_end": {
		CPU:     2500.0,
		Memory:  50000.0,
		Storage: 2000.0,
		GPU:     25000.0,
		Network: 500.0,
	},
	"enthusiast": {
		CPU:     4000.0,
		Memory:  70000.0,
		Storage: 5000.0,
		GPU:     40000.0,
		Network: 1000.0,
	},
}

// GetBaseline returns the baseline data for the specified category
func GetBaseline(category string) BaselineData {
	if baseline, exists := BaselineDB[category]; exists {
		return baseline
	}
	return BaselineDB["default"]
}

// UpdateBaseline updates the baseline data for a specific category
func UpdateBaseline(category string, data BaselineData) {
	BaselineDB[category] = data
}

// GetAllBaselines returns all available baseline categories
func GetAllBaselines() map[string]BaselineData {
	return BaselineDB
}
