// Package created to keep only one DB connection and handle the application settings

package config

import (
  "os"
  "fmt"
  "io/ioutil"
  "encoding/json"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
  _ "github.com/go-sql-driver/mysql"
)

// DB is the database handler for all connections.
var (
  DB gorm.DB
  err error
  Settings map[string]interface{}
  TokenSecret = "weAreTheChampions"
)

// GetLogFile opens the log file
func GetLogFile() (*os.File){
  outputFile, outErr := os.OpenFile(Settings["LogFile"].(string), os.O_WRONLY|os.O_CREATE, 0666)
	if outErr != nil {
		fmt.Println("It was not possible to create the file!")
    return os.Stdout
	}
  return outputFile
}

// getConfigContent reads the settings file and returns its content as string
func getConfigContent(path string) string{
  dataFile, err := ioutil.ReadFile(path)
	if err != nil {
    fmt.Println("Error: %v", err)
		panic("Error reading the file.")
	}
	return string(dataFile)
}

// requireSettingsField validates required settings fields
func requireSettingsField(fieldName string) {
  field := Settings[fieldName]
  if field == nil {
    panic("Settings file shall have the " + fieldName + " parameter")
  }
}

// LoadConfig loads the main config sets
func LoadConfig(path string) {
  jsonContent := ""
  if path == "" {
    jsonContent = getConfigContent("settings.json")
  } else {
    jsonContent = getConfigContent(path)
  }

  json.Unmarshal([]byte(jsonContent), &Settings)
  requireSettingsField("ListenAddress")
  requireSettingsField("DatabaseUri")
  requireSettingsField("LogFile")

  startDB()
}

// startDB opens a connection with Database. It shall be loaded after the settings, since settings contains the needed URLs
func startDB(){
  DB, err = gorm.Open("mysql", Settings["DatabaseUri"])
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    panic("It was not possible to connect database.")
  }
  DB.DB()
  fmt.Printf("Database Started.")
}
