package database

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Database struct to hold the database connection
type Database struct {
    Connection *gorm.DB
}

// NewDatabase initializes a new database connection
func NewDatabase(dataSourceName string) (*Database, error) {
    db, err := gorm.Open(sqlite.Open(dataSourceName), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return &Database{Connection: db}, nil
}

// Migrate runs the database migrations
func (db *Database) Migrate(models ...interface{}) error {
    return db.Connection.AutoMigrate(models...)
}

// Close closes the database connection
func (db *Database) Close() error {
    sqlDB, err := db.Connection.DB()
    if err != nil {
        return err
    }
    return sqlDB.Close()
}