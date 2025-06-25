package octane

// OctaneGrade 定义辛烷值等级的结构
type OctaneGrade struct {
    RON         float64 `yaml:"ron"`          // Research Octane Number
    Grade       string  `yaml:"grade"`       // 等级名称
    Description string  `yaml:"description"` // 描述
    Color       string  `yaml:"color"`       // 显示颜色
}

// OctaneGrades 定义所有辛烷值等级
var OctaneGrades = []OctaneGrade{
    {
        RON:         95,
        Grade:       "racing_fuel",
        Description: "Ultimate performance for extreme workloads",
        Color:       "🔥 RED",
    },
    {
        RON:         90,
        Grade:       "premium_plus",
        Description: "High performance for demanding applications",
        Color:       "🟠 ORANGE",
    },
    {
        RON:         85,
        Grade:       "premium",
        Description: "Good performance for most applications",
        Color:       "🟡 YELLOW",
    },
    {
        RON:         80,
        Grade:       "regular_plus",
        Description: "Standard performance for regular use",
        Color:       "🟢 GREEN",
    },
    {
        RON:         70,
        Grade:       "regular",
        Description: "Basic performance for light workloads",
        Color:       "🔵 BLUE",
    },
}