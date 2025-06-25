package database

import (
    "gorm.io/gorm"
)

// Migrate 执行数据库迁移
func Migrate(db *gorm.DB) error {
    // 在此处添加需要迁移的模型
    // err := db.AutoMigrate(&YourModel{})
    // if err != nil {
    //     return err
    // }
    return nil
}