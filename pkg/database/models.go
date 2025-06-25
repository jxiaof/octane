package database

import (
    "gorm.io/gorm"
)

// TestResult 定义测试结果模型
type TestResult struct {
    ID        uint   `gorm:"primaryKey"`
    TestType  string `gorm:"not null"` // 测试类型，例如 CPU、内存、存储等
    Score     float64 `gorm:"not null"` // 测试得分
    Timestamp string `gorm:"not null"` // 测试时间戳
}

// SystemInfo 定义系统信息模型
type SystemInfo struct {
    ID        uint   `gorm:"primaryKey"`
    Hostname  string `gorm:"not null"` // 主机名
    OS        string `gorm:"not null"` // 操作系统
    CPU       string `gorm:"not null"` // CPU 信息
    Memory    string `gorm:"not null"` // 内存信息
    Storage   string `gorm:"not null"` // 存储信息
    GPU       string `gorm:"not null"` // GPU 信息
}

// UploadRecord 定义数据上报记录模型
type UploadRecord struct {
    ID        uint   `gorm:"primaryKey"`
    ReportID  string `gorm:"not null"` // 报告 ID
    Status    string `gorm:"not null"` // 上报状态
    Timestamp string `gorm:"not null"` // 上报时间戳
}

// InitializeDatabase 初始化数据库
func InitializeDatabase(db *gorm.DB) error {
    // 自动迁移模型
    err := db.AutoMigrate(&TestResult{}, &SystemInfo{}, &UploadRecord{})
    return err
}