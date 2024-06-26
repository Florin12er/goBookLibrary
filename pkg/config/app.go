package config

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func Connect() {
    d, err := gorm.Open("postgres", "postgresql://florin:florin12er@localhost:5432/gobooklibrary?sslmode=disable")
    if err != nil {
        panic(err)
    }
    db = d
}

func GetDB() *gorm.DB {
    return db
}

