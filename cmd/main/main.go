package main

import (
    "fmt"
    "github.com/florin12er/goBookstore/config"
)

func main() {
databaseURL := "postgresql://florin:florin12er@localhost:5432/gobooklibrary"

config.Connect(databaseURL)
}
