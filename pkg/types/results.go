package types

// TestResults 定义测试结果的结构
type TestResults struct {
	CPU     CPUResults     `json:"cpu"`
	Memory  MemoryResults  `json:"memory"`
	Storage StorageResults `json:"storage"`
	GPU     GPUResults     `json:"gpu"`
	Network NetworkResults `json:"network"`
}

// CPUResults 定义CPU测试结果的结构
type CPUResults struct {
	TestSuite string `json:"test_suite"`
	Duration  string `json:"duration"`

	Temperature struct {
		Idle float64 `json:"idle"` // °C
		Load float64 `json:"load"` // °C
		Max  float64 `json:"max"`  // °C
	} `json:"temperature"`

	Frequencies struct {
		AverageAllCores float64 `json:"average_all_cores"` // MHz
		Stability       float64 `json:"stability"`         // %
	} `json:"frequencies"`

	Tests struct {
		SingleCore struct {
			IntegerPerformance struct {
				Score      int    `json:"score"`
				Unit       string `json:"unit"`
				Percentile int    `json:"percentile"`
			} `json:"integer_performance"`
			FloatingPoint struct {
				Score      float64 `json:"score"`
				Unit       string  `json:"unit"`
				Percentile int     `json:"percentile"`
			} `json:"floating_point"`
			Cryptography struct {
				AES256  float64 `json:"aes_256"`  // GB/s
				SHA256  float64 `json:"sha256"`   // GB/s
				RSA2048 int     `json:"rsa_2048"` // ops/sec
			} `json:"cryptography"`
		} `json:"single_core"`

		MultiCore struct {
			IntegerPerformance struct {
				Score      int    `json:"score"`
				Unit       string `json:"unit"`
				Percentile int    `json:"percentile"`
			} `json:"integer_performance"`
			FloatingPoint struct {
				Score      float64 `json:"score"`
				Unit       string  `json:"unit"`
				Percentile int     `json:"percentile"`
			} `json:"floating_point"`
			Compression struct {
				Gzip int `json:"gzip"` // MB/s
				LZ4  int `json:"lz4"`  // MB/s
				Zstd int `json:"zstd"` // MB/s
			} `json:"compression"`
		} `json:"multi_core"`
	} `json:"tests"`
}

// MemoryResults 定义内存测试结果的结构
type MemoryResults struct {
	TestSuite string `json:"test_suite"`
	Duration  string `json:"duration"`

	Bandwidth struct {
		SequentialRead  float64 `json:"sequential_read"`  // MB/s
		SequentialWrite float64 `json:"sequential_write"` // MB/s
		RandomRead      float64 `json:"random_read"`      // MB/s
		RandomWrite     float64 `json:"random_write"`     // MB/s
		Copy            float64 `json:"copy"`             // MB/s
	} `json:"bandwidth"`

	Latency struct {
		L1Cache    float64 `json:"l1_cache"`    // ns
		L2Cache    float64 `json:"l2_cache"`    // ns
		L3Cache    float64 `json:"l3_cache"`    // ns
		MainMemory float64 `json:"main_memory"` // ns
	} `json:"latency"`

	Stability struct {
		ErrorsDetected int     `json:"errors_detected"`
		TestDuration   string  `json:"test_duration"`
		MemoryTested   float64 `json:"memory_tested"` // MB
		Passes         int     `json:"passes"`
	} `json:"stability"`
}

// StorageResults 定义存储测试结果的结构
type StorageResults struct {
	TestSuite string          `json:"test_suite"`
	Duration  string          `json:"duration"`
	Devices   []DeviceResults `json:"devices"`
}

// DeviceResults 定义单个存储设备的测试结果
type DeviceResults struct {
	Name  string `json:"name"`
	Tests struct {
		Sequential struct {
			Read1MB  float64 `json:"read_1mb"`  // MB/s
			Write1MB float64 `json:"write_1mb"` // MB/s
			Read4K   float64 `json:"read_4k"`   // MB/s
			Write4K  float64 `json:"write_4k"`  // MB/s
		} `json:"sequential"`
		Random struct {
			Read4KIops  float64 `json:"read_4k_iops"`  // IOPS
			Write4KIops float64 `json:"write_4k_iops"` // IOPS
			Mixed70_30  float64 `json:"mixed_70_30"`   // IOPS
		} `json:"random"`
		Latency struct {
			ReadAvg  float64 `json:"read_avg"`  // ms
			WriteAvg float64 `json:"write_avg"` // ms
			Read99p  float64 `json:"read_99p"`  // ms
			Write99p float64 `json:"write_99p"` // ms
		} `json:"latency"`
	} `json:"tests"`
}

// GPUResults 定义GPU测试结果的结构
type GPUResults struct {
	TestSuite string `json:"test_suite"`
	Duration  string `json:"duration"`

	Temperature struct {
		Idle float64 `json:"idle"` // °C
		Load float64 `json:"load"` // °C
		Max  float64 `json:"max"`  // °C
	} `json:"temperature"`

	PowerConsumption struct {
		Idle    float64 `json:"idle"`    // Watts
		Average float64 `json:"average"` // Watts
		Peak    float64 `json:"peak"`    // Watts
	} `json:"power_consumption"`

	Tests struct {
		Graphics struct {
			OpenGL struct {
				Score    int `json:"score"`
				FPS1080p int `json:"fps_1080p"`
				FPS1440p int `json:"fps_1440p"`
				FPS4K    int `json:"fps_4k"`
			} `json:"opengl"`
			DirectX12 struct {
				Score    int `json:"score"`
				FPS1080p int `json:"fps_1080p"`
				FPS1440p int `json:"fps_1440p"`
				FPS4K    int `json:"fps_4k"`
			} `json:"directx12"`
			Vulkan struct {
				Score    int `json:"score"`
				FPS1080p int `json:"fps_1080p"`
				FPS1440p int `json:"fps_1440p"`
				FPS4K    int `json:"fps_4k"`
			} `json:"vulkan"`
			Score float64 `json:"score"` // 综合图形评分
		} `json:"graphics"`

		Compute struct {
			CUDA struct {
				SinglePrecision float64 `json:"single_precision"` // TFLOPS
				HalfPrecision   float64 `json:"half_precision"`   // TFLOPS
				TensorOps       float64 `json:"tensor_ops"`       // TOPS
			} `json:"cuda"`
			OpenCL struct {
				SinglePrecision float64 `json:"single_precision"` // TFLOPS
				DoublePrecision float64 `json:"double_precision"` // TFLOPS
			} `json:"opencl"`
			SinglePrecision float64 `json:"single_precision"` // 综合计算性能
		} `json:"compute"`

		MachineLearning struct {
			Inference struct {
				ResNet50FP32 struct {
					Batch1  int `json:"batch_1"`  // FPS
					Batch32 int `json:"batch_32"` // FPS
				} `json:"resnet50_fp32"`
				BertBase struct {
					Batch1  int `json:"batch_1"`  // sequences/sec
					Batch16 int `json:"batch_16"` // sequences/sec
				} `json:"bert_base"`
			} `json:"inference"`
			Training struct {
				SimpleCNN struct {
					Batch32  int `json:"batch_32"`  // samples/sec
					Batch128 int `json:"batch_128"` // samples/sec
				} `json:"simple_cnn"`
			} `json:"training"`
		} `json:"machine_learning"`

		VideoEncoding struct {
			H264_1080p int `json:"h264_1080p"` // FPS
			H264_4K    int `json:"h264_4k"`    // FPS
			H265_1080p int `json:"h265_1080p"` // FPS
			H265_4K    int `json:"h265_4k"`    // FPS
			AV1_1080p  int `json:"av1_1080p"`  // FPS
			AV1_4K     int `json:"av1_4k"`     // FPS
		} `json:"video_encoding"`

		Memory struct {
			Bandwidth float64 `json:"bandwidth"` // GB/s
			Latency   float64 `json:"latency"`   // μs
		} `json:"memory"`
	} `json:"tests"`
}

// NetworkResults 定义网络测试结果的结构
type NetworkResults struct {
	TestSuite string `json:"test_suite"`
	Duration  string `json:"duration"`

	Bandwidth struct {
		Domestic      map[string]BandwidthResult `json:"domestic"`
		International map[string]BandwidthResult `json:"international"`
	} `json:"bandwidth"`

	Connectivity struct {
		DNSResolution        map[string]float64 `json:"dns_resolution"` // ms
		ServiceAccessibility map[string]bool    `json:"service_accessibility"`
		PortScan             map[string]string  `json:"port_scan"`
	} `json:"connectivity"`
}

// BandwidthResult 定义带宽测试结果的结构
type BandwidthResult struct {
	Download   float64 `json:"download"`    // Mbps
	Upload     float64 `json:"upload"`      // Mbps
	Latency    float64 `json:"latency"`     // ms
	Jitter     float64 `json:"jitter"`      // ms
	PacketLoss float64 `json:"packet_loss"` // %
}
