package types

type Config struct {
    LogLevel       string `yaml:"log_level"`
    OutputFormat   string `yaml:"output_format"`
    ProgressBar    bool   `yaml:"progress_bar"`
    TempDir        string `yaml:"temp_dir"`
    Theme          string `yaml:"theme"`
    
    Octane struct {
        RatingSystem string  `yaml:"rating_system"`
        ScaleMax     float64 `yaml:"scale_max"`
        ScaleMin     float64 `yaml:"scale_min"`
        Precision     int     `yaml:"precision"`
    } `yaml:"octane"`

    Upload struct {
        Enabled     bool   `yaml:"enabled"`
        ServerURL   string `yaml:"server_url"`
        APIKey      string `yaml:"api_key"`
        Anonymous    bool   `yaml:"anonymous"`
        AutoUpload   bool   `yaml:"auto_upload"`
        Tags        []string `yaml:"tags"`
    } `yaml:"upload"`

    Tests struct {
        BoostMode              bool `yaml:"boost_mode"`
        FuelAnalysis           bool `yaml:"fuel_analysis"`
        TemperatureMonitoring  bool `yaml:"temperature_monitoring"`
        PowerMonitoring        bool `yaml:"power_monitoring"`
    } `yaml:"tests"`
}