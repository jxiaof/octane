package types

// CPUInfo represents detailed CPU platform information
type CPUInfo struct {
	ModelName          string   `json:"model_name"`
	Brand              string   `json:"brand"`
	Architecture       string   `json:"architecture"`
	PhysicalCores      int      `json:"physical_cores"`
	LogicalCores       int      `json:"logical_cores"`
	BaseFrequency      float64  `json:"base_frequency_ghz"`
	MaxFrequency       float64  `json:"max_frequency_ghz"`
	CurrentFrequency   float64  `json:"current_frequency_ghz"`
	CurrentTemperature float64  `json:"current_temperature_celsius"`
	CacheL1Data        string   `json:"cache_l1_data"`
	CacheL1Instruction string   `json:"cache_l1_instruction"`
	CacheL2            string   `json:"cache_l2"`
	CacheL3            string   `json:"cache_l3"`
	Features           []string `json:"features"`
	TDP                int      `json:"tdp_watts"`
	Family             int      `json:"family"`
	Model              int      `json:"model"`
	Stepping           int      `json:"stepping"`
}
