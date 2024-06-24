package config

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func Connect(databaseURL string) {
    d, err := gorm.Open("postgres", databaseURL)
    if err != nil {
        panic(err)
    }
    db = d
}

func GetDB() *gorm.DB {
    return db
}

