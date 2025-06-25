package types

// Report represents the structure of a performance report.
type Report struct {
	Metadata        Metadata        `yaml:"metadata"`
	SystemInfo      SystemInfo      `yaml:"system_info"`
	TestResults     TestResults     `yaml:"test_results"`
	Scores          Scores          `yaml:"scores"`
	Comparisons     Comparisons     `yaml:"comparisons"`
	Recommendations Recommendations `yaml:"recommendations"`
	UploadInfo      UploadInfo      `yaml:"upload_info"`
	OctaneRatings   OctaneRatings   `yaml:"octane_ratings"`
}

// Metadata contains information about the report.
type Metadata struct {
	Version       string   `yaml:"version"`
	TestID        string   `yaml:"test_id"`
	Timestamp     string   `yaml:"timestamp"`
	User          string   `yaml:"user"`
	Hostname      string   `yaml:"hostname"`
	Duration      string   `yaml:"duration"`
	UploadConsent bool     `yaml:"upload_consent"`
	Tags          []string `yaml:"tags"`
}

// SystemInfo contains details about the system being tested.
type SystemInfo struct {
	Host    HostInfo      `yaml:"host"`
	CPU     CPUInfo       `yaml:"cpu"`
	Memory  MemoryInfo    `yaml:"memory"`
	Storage []StorageInfo `yaml:"storage"`
	GPU     []GPUInfo     `yaml:"gpu"`
	Network []NetworkInfo `yaml:"network"`
}

// Scores contains the overall and breakdown scores.
type Scores struct {
	Overall               float64               `yaml:"overall"`
	Breakdown             map[string]float64    `yaml:"breakdown"`
	ProfessionalScenarios ProfessionalScenarios `yaml:"professional_scenarios"`
}

// Comparisons contains information for comparing results.
type Comparisons struct {
	PercentileRanking   int              `yaml:"percentile_ranking"`
	SimilarSystemsCount int              `yaml:"similar_systems_count"`
	SimilarSystems      []SimilarSystem  `yaml:"similar_systems"`
	Recommendations     []Recommendation `yaml:"recommendations"`
}

// Recommendations contains optimization tips.
type Recommendations struct {
	FuelOptimizationTips []FuelOptimizationTip `yaml:"fuel_optimization_tips"`
}

// UploadInfo contains information about the report upload.
type UploadInfo struct {
	Uploaded   bool   `yaml:"uploaded"`
	UploadTime string `yaml:"upload_time"`
	Server     string `yaml:"server"`
	Anonymized bool   `yaml:"anonymized"`
	ReportID   string `yaml:"report_id"`
}

// OctaneRatings contains the octane ratings for the system.
type OctaneRatings struct {
	Overall   OctaneRating            `yaml:"overall"`
	Breakdown map[string]OctaneRating `yaml:"breakdown"`
}

// HostInfo contains host system information.
type HostInfo struct {
	OS           string `yaml:"os"`
	Kernel       string `yaml:"kernel"`
	Architecture string `yaml:"architecture"`
	Hostname     string `yaml:"hostname"`
	Uptime       string `yaml:"uptime"`
	Timezone     string `yaml:"timezone"`
}

// CPUCores contains CPU core information.
type CPUCores struct {
	Physical int `yaml:"physical"`
	Logical  int `yaml:"logical"`
}

// CPUFrequencies contains CPU frequency information.
type CPUFrequencies struct {
	Base  int `yaml:"base"`  // MHz
	Boost int `yaml:"boost"` // MHz
}

// CPUCache contains CPU cache information.
type CPUCache struct {
	L1D string `yaml:"l1d"`
	L1I string `yaml:"l1i"`
	L2  string `yaml:"l2"`
	L3  string `yaml:"l3"`
}

// MemoryInfo contains memory information.
type MemoryInfo struct {
	Total     int            `yaml:"total"`     // MB
	Available int            `yaml:"available"` // MB
	Type      string         `yaml:"type"`
	Frequency int            `yaml:"frequency"` // MHz
	Timing    string         `yaml:"timing"`
	Slots     MemorySlots    `yaml:"slots"`
	Modules   []MemoryModule `yaml:"modules"`
}

// MemorySlots contains memory slot information.
type MemorySlots struct {
	Used  int `yaml:"used"`
	Total int `yaml:"total"`
}

// MemoryModule contains individual memory module information.
type MemoryModule struct {
	Size         int    `yaml:"size"` // MB
	Manufacturer string `yaml:"manufacturer"`
	PartNumber   string `yaml:"part_number"`
}

// StorageInfo contains storage device information.
type StorageInfo struct {
	Name        string `yaml:"name"`
	Model       string `yaml:"model"`
	Type        string `yaml:"type"`
	Interface   string `yaml:"interface"`
	Capacity    int    `yaml:"capacity"`    // MB
	Used        int    `yaml:"used"`        // MB
	Health      int    `yaml:"health"`      // %
	Temperature int    `yaml:"temperature"` // °C
}

// GPUInfo contains GPU information.
type GPUInfo struct {
	Index        int            `yaml:"index"`
	Name         string         `yaml:"name"`
	Architecture string         `yaml:"architecture"`
	PCIBus       string         `yaml:"pci_bus"`
	CUDACores    int            `yaml:"cuda_cores"`
	RTCores      int            `yaml:"rt_cores"`
	TensorCores  int            `yaml:"tensor_cores"`
	Memory       GPUMemory      `yaml:"memory"`
	Frequencies  GPUFrequencies `yaml:"frequencies"`
	Power        GPUPower       `yaml:"power"`
	Temperature  int            `yaml:"temperature"` // °C
	Driver       GPUDriver      `yaml:"driver"`
}

// GPUMemory contains GPU memory information.
type GPUMemory struct {
	Total     int    `yaml:"total"` // MB
	Type      string `yaml:"type"`
	Bandwidth int    `yaml:"bandwidth"` // GB/s
	BusWidth  int    `yaml:"bus_width"` // bits
}

// GPUFrequencies contains GPU frequency information.
type GPUFrequencies struct {
	Base   int `yaml:"base"`   // MHz
	Boost  int `yaml:"boost"`  // MHz
	Memory int `yaml:"memory"` // MHz
}

// GPUPower contains GPU power information.
type GPUPower struct {
	TDP     int `yaml:"tdp"`     // Watts
	Current int `yaml:"current"` // Watts
}

// GPUDriver contains GPU driver information.
type GPUDriver struct {
	Version       string `yaml:"version"`
	CUDAVersion   string `yaml:"cuda_version"`
	OpenGLVersion string `yaml:"opengl_version"`
	VulkanVersion string `yaml:"vulkan_version"`
}

// NetworkInfo contains network interface information.
type NetworkInfo struct {
	Name   string `yaml:"name"`
	Type   string `yaml:"type"`
	MAC    string `yaml:"mac"`
	Model  string `yaml:"model"`
	Driver string `yaml:"driver"`
	Speed  int    `yaml:"speed"` // Mbps
	Duplex string `yaml:"duplex"`
	Status string `yaml:"status"`
	IPv4   string `yaml:"ipv4,omitempty"`
	IPv6   string `yaml:"ipv6,omitempty"`
	SSID   string `yaml:"ssid,omitempty"`
	Signal int    `yaml:"signal,omitempty"` // dBm
}

// OctaneRating contains octane rating information.
type OctaneRating struct {
	RON         float64 `yaml:"ron"`
	Grade       string  `yaml:"grade"`
	Description string  `yaml:"description"`
	Color       string  `yaml:"color"`
}

// ProfessionalScenarios contains professional scenario scores.
type ProfessionalScenarios struct {
	Gaming            ProfessionalScore `yaml:"gaming"`
	AIMachineLearning ProfessionalScore `yaml:"ai_machine_learning"`
	ServerWorkload    ProfessionalScore `yaml:"server_workload"`
	Workstation       ProfessionalScore `yaml:"workstation"`
}

// ProfessionalScore contains professional scenario score information.
type ProfessionalScore struct {
	Score       float64 `yaml:"score"`
	Grade       string  `yaml:"grade"`
	Description string  `yaml:"description"`
}

// SimilarSystem contains information about similar systems.
type SimilarSystem struct {
	Hostname     string  `yaml:"hostname"`
	OverallScore float64 `yaml:"overall_score"`
	CPUModel     string  `yaml:"cpu_model"`
	GPUModel     string  `yaml:"gpu_model"`
	Location     string  `yaml:"location"`
	TestDate     string  `yaml:"test_date"`
}

// Recommendation contains optimization recommendations.
type Recommendation struct {
	Category   string `yaml:"category"`
	Suggestion string `yaml:"suggestion"`
	Impact     string `yaml:"impact"`
}

// FuelOptimizationTip contains fuel optimization tips.
type FuelOptimizationTip struct {
	Category    string `yaml:"category"`
	Tip         string `yaml:"tip"`
	OctaneBoost string `yaml:"octane_boost"`
}
