// Package created to keep only one DB connection and handle the application settings

package config

import (
  "fmt"
  "github.com/jinzhu/gorm"
  // libraries used for gorm
  _ "github.com/lib/pq"
  _ "github.com/go-sql-driver/mysql"
)

// DB is the database handler for all connections.
var (
  DB gorm.DB
  err error
)

func init(){
  DB, err = gorm.Open("mysql", "root:@/tasks?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    panic("It was not possible to connect database.")
  }
  DB.DB()
  fmt.Printf("Database Started.")
}
